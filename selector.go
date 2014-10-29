package goquery

import (
	"container/list"
	"sync"

	"code.google.com/p/cascadia"
)

type selector struct {
	cs  cascadia.Selector
	sel string
	el  *list.Element
}

var (
	selectorLock      sync.Mutex
	selectorCacheSize = 512
	selectorCache     = map[string]*selector{}
	selectorLru       = list.New()
)

// Tune the size of the internal selector cache
func SetSelectorCacheSize(size int) {
	selectorLock.Lock()

	if size >= 0 {
		selectorCacheSize = size
		purgeExtraSelectors()
	}

	selectorLock.Unlock()
}

func purgeExtraSelectors() {
	for selectorLru.Len() > selectorCacheSize {
		os := selectorLru.Remove(selectorLru.Front()).(*selector)
		delete(selectorCache, os.sel)
	}
}

func getSelector(sel string) cascadia.Selector {
	selectorLock.Lock()

	s, ok := selectorCache[sel]

	if !ok {
		// Don't block everything while compiling the selector
		selectorLock.Unlock()

		s = &selector{
			cs:  cascadia.MustCompile(sel),
			sel: sel,
		}

		selectorLock.Lock()

		sn, ok := selectorCache[sel]
		if ok {
			s = sn
		} else {
			selectorCache[sel] = s
			s.el = selectorLru.PushBack(s)
			purgeExtraSelectors()
		}
	}

	selectorLru.MoveToBack(s.el)

	selectorLock.Unlock()

	return s.cs
}
