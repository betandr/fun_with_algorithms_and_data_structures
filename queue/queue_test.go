package queue

import "testing"

func TestEmptyQueue(t *testing.T) {
	queue := new[string]()
	got := queue.String()
	want := "START END"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEnqueue(t *testing.T) {
	queue := new[string]()
	queue.Enqueue("MIDDLE")

	got := queue.String()
	want := "START MIDDLE END"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDequeue(t *testing.T) {
	queue := new[string]()
	queue.Enqueue("ONE")
	queue.Enqueue("TWO")
	queue.Enqueue("THREE")

	queue.Dequeue()

	got := queue.String()
	want := "START TWO THREE END"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	queue := new[string]()

	err := queue.Dequeue()
	if err == nil {
		t.Error("error: no error dequeuing from empty queue")
	}
}

func TestPeek(t *testing.T) {
	queue := new[string]()
	want := "ONE"
	queue.Enqueue(want)

	got, _ := queue.Peek()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	queue := new[string]()

	_, err := queue.Peek()
	if err == nil {
		t.Error("error: no error peeking empty queue")
	}
}
