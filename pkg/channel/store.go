package channel

import (
	"sync"
)

type Store struct {
	m     map[string]*Channel
	mutex sync.Mutex
}

func NewStore() *Store {
	return &Store{
		m: map[string]*Channel{},
	}
}

func (s *Store) Get(name string) *Channel {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.m[name]
}

func (s *Store) Put(ch *Channel) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[ch.Name] = ch
}

func (s *Store) List() []*Channel {
	cs := make([]*Channel, 0)
	for k := range s.m {
		cs = append(cs, s.m[k])
	}
	return cs
}

func (s *Store) Size() int {
	return len(s.m)
}
