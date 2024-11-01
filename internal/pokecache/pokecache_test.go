package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = time.Second * 5
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://testurl.com",
			val: []byte("testdata"),
		}, {
			key: "https://testurl.com/path",
			val: []byte("testdata2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("testKey", []byte("testVal"))

	_, ok := cache.Get("testKey")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	_, ok = cache.Get("testKey")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}