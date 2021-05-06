package go_pg_canceling_bug_test

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	go_pg_canceling_bug "github.com/Barben360/go-pg-conn-canceling-bug"
)

func TestConcurrencyAndUserCancel(t *testing.T) {
	const poolSize = 10
	const concurrency = 10
	const iterations = 100
	const cancelProbability = 0.1

	db, err := go_pg_canceling_bug.Init(poolSize)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctxBkg := context.Background()

	cancelWaitMin := 1 * time.Microsecond
	cancelWaitMax := 5 * time.Millisecond

	cancelProbabilityScaled := int(cancelProbability * 1000000)

	obj := go_pg_canceling_bug.Foo{}

	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// Are we gonna cancel our call ?
				cancelWait := time.Duration(0)
				val := rand.Intn(1000000)
				if val <= cancelProbabilityScaled {
					cancelWait = time.Duration(rand.Int63n(int64(cancelWaitMax-cancelWaitMin))) + cancelWaitMin
					t.Logf("Canceling add statement in %v", cancelWait)
				}

				// Adding object
				ctx := ctxBkg
				cancel := func() {}
				if cancelWait != 0 {
					ctx, cancel = context.WithTimeout(ctx, cancelWait)
				}
				objCopy := obj
				err := (&objCopy).Insert(ctx, db)
				cancel()
				if err != nil {
					if cancelWait == 0 {
						t.Errorf("An error occurred for concurrent %d - iteration %d - call 1: %v", i, j, err)
					}
					continue
				}

				// Editing object
				objCopy.Comments = "some string"
				err = (&objCopy).Update(ctx, db)
				if err != nil {
					t.Errorf("An error occurred for concurrent %d - iteration %d - call 3: %v", i, j, err)
				}

				// deleting object
				err = (&objCopy).Delete(ctx, db)
				if err != nil {
					t.Error(err)
				}
			}
		}(&wg, i)
	}
	doneChan := make(chan bool)
	go func(doneChan chan bool) {
		wg.Wait()
		close(doneChan)
	}(doneChan)

	select {
	case <-doneChan:
	case <-time.After(1 * time.Minute):
		t.Error("Test timeout")
	}
}
