package store

import (
	"sync"

	"github.com/ihatiko/go-chef-core-sdk/iface"
)

var LivenessStore = store{
	mt:         sync.Mutex{},
	components: map[string]iface.ILive{},
}

type store struct {
	mt         sync.Mutex
	components map[string]iface.ILive
}

func (s *store) Load(live iface.ILive) {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.components[live.Name()] = live
}

func (s *store) Get() []iface.ILive {
	var result []iface.ILive
	for _, v := range s.components {
		result = append(result, v)
	}
	return result
}
