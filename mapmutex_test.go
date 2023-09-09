package mapmutex_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/erwanlbp/mapmutex"
)

func TestMapMutex(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().Unix())

	var mm mapmutex.Mutexs
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			r := resource()

			defer wg.Done()
			mm.Lock(r)
			defer mm.Unlock(r)

			fmt.Printf("[%d] locked %s\n", i, r)
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Millisecond)
			fmt.Printf("[%d] unlock %s\n", i, r)
		}()
	}

	wg.Wait()
}

func resource() string {
	return strconv.Itoa(rand.Intn(4))
}
