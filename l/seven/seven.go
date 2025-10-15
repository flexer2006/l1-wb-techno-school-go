package seven

import "sync"

type SyncData struct {
	m  map[any]any
	mu sync.RWMutex
}

func NewSyncData() *SyncData {
	return &SyncData{
		m: make(map[any]any),
	}
}

func (s *SyncData) SetSyncData(k, v any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

func concurrencyWriting() {
	var wg sync.WaitGroup
	data := NewSyncData()

	for i := range 10 {
		ii := i
		wg.Go(func() {
			data.SetSyncData(ii, ii)
		})
	}
	wg.Wait()
}
