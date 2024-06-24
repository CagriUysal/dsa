package queue_test

import (
	"dsa/ds/queue"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQueue(t *testing.T) {
	q := queue.Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	item, _ := q.Dequeue()

	if diff := cmp.Diff(1, item); diff != "" {
		t.Error(diff)
	}

	item, _ = q.Dequeue()
	if diff := cmp.Diff(2, item); diff != "" {
		t.Error(diff)
	}
}
