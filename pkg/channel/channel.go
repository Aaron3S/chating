package channel

import (
	api "charting/pkg/api"
	"container/ring"
	"sort"
	"sync"
)

const historySize = 10

type Channel struct {
	api.Channel
	sessions  map[string]int32
	history   *ring.Ring
	password  string
	currentId int32
	mutex     sync.Mutex
}

type Option func(c *Channel)

func (c *Channel) Write(message *api.Message) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	message.Id = c.currentId
	c.currentId++
	c.history = c.history.Next()
	c.history.Value = message
}

type historyList []*api.Message

func (h historyList) Len() int {
	return len(h)
}

func (h historyList) Less(i, j int) bool {
	return h[i].Id < h[j].Id
}

func (h historyList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (c *Channel) Read(name string) []*api.Message {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ms := historyList(make([]*api.Message, 0))
	if c.currentId == 0 {
		return ms
	}
	lastMsgId, hasSession := c.sessions[name]
	if hasSession {
		if lastMsgId >= c.history.Value.(*api.Message).Id {
			return ms
		}
	}
	h := c.history
	for i := 0; i < h.Len(); i++ {
		if h.Value != nil {
			m := h.Value.(*api.Message)
			if c.sessions[name] <= m.Id {
				c.sessions[name] = m.Id
			}
			if !hasSession {
				ms = append(ms, m)
			} else {
				if m.Id > lastMsgId {
					ms = append(ms, m)
				}
			}
		}
		h = h.Next()
	}
	sort.Sort(ms)
	return ms
}

func WithPassword(password string) Option {
	return func(c *Channel) {
		c.password = password
	}
}

func NewChannel(name string, opts ...Option) *Channel {
	ch := &Channel{
		Channel: api.Channel{
			Name: name,
		},
		history:  ring.New(historySize),
		sessions: map[string]int32{},
	}
	for _, op := range opts {
		op(ch)
	}
	if ch.password != "" {
		ch.Private = false
	}
	return ch
}
