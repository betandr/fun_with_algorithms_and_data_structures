package stack

import "testing"

func TestEmptyStack(t *testing.T) {
	stack := new[string]()
	got := stack.String()
	want := "TOP"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPush(t *testing.T) {
	stack := new[string]()
	stack.Push("ONE")
	stack.Push("TWO")

	want := "TOP\nTWO\nONE"
	got := stack.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestPeek(t *testing.T) {
	stack := new[string]()
	stack.Push("ONE")
	stack.Push("TWO")

	itemGot, _ := stack.Peek()
	itemWant := "TWO"

	if itemGot != itemWant {
		t.Errorf("got item %s, want item %s", itemGot, itemWant)
	}

	wantStack := "TOP\nTWO\nONE"
	gotStack := stack.String()

	if gotStack != wantStack {
		t.Errorf("got %s, want %s", gotStack, wantStack)
	}
}

func TestPeekFromEmpty(t *testing.T) {
	stack := new[string]()

	_, err := stack.Peek()
	if err == nil {
		t.Errorf("pop from empty stack should be error")
	}
}

func TestPop(t *testing.T) {
	stack := new[string]()
	stack.Push("ONE")
	stack.Push("TWO")

	itemGot, _ := stack.Pop()
	itemWant := "TWO"

	if itemGot != itemWant {
		t.Errorf("got item %s, want item %s", itemGot, itemWant)
	}

	wantStack := "TOP\nONE"
	gotStack := stack.String()

	if gotStack != wantStack {
		t.Errorf("got %s, want %s", gotStack, wantStack)
	}
}

func TestPopFromEmpty(t *testing.T) {
	stack := new[string]()

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("pop from empty stack should be error")
	}
}
