package goquery

import (
	"fmt"
	"testing"
)

var (
	origSelectorCacheSize = selectorCacheSize
)

func resetSelector(t *testing.T) {
	SetSelectorCacheSize(0)
	SetSelectorCacheSize(origSelectorCacheSize)
}

func TestSelectorCache(t *testing.T) {
	resetSelector(t)

	resized := false

	for i := 1; i < selectorCacheSize*2; i++ {
		if i > selectorCacheSize && !resized {
			SetSelectorCacheSize(selectorCacheSize + 10)
			resized = true
		}

		getSelector(fmt.Sprintf(".i%d", i))
	}

	if selectorLru.Len() != selectorCacheSize {
		t.Errorf("Expected selectorLru to have %d selectors, has %d",
			selectorLru.Len(),
			selectorCacheSize)
	}
}

func TestSelectorLru(t *testing.T) {
	resetSelector(t)

	sels := []string{
		".a1",
		".a2",
		".a3",
		".a4",
	}

	for i := 0; i < len(sels)*2; i++ {
		sel := sels[i%len(sels)]
		getSelector(sel)

		s := selectorLru.Back().Value.(*selector)
		if s.sel != sel {
			t.Fatal("Selector not moved to back of LRU cache")
		}
	}
}
