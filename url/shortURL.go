package url

import (
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.Mutex
}
