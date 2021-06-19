package ds

import "testing"

func TestLen(t *testing.T)  {
	ll := LinkedList{}

	want := 0

	if got := ll.Len(); got != want {
		t.Errorf("Len() = %v, want %v", got, want)
	}
}

func TestAddHead(t *testing.T)  {
	ll := LinkedList{}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		want := val[i]
		ll.AddHead(want)
		
		if got := ll.head.data; got != want {
			t.Errorf("head.data = %v, want %v", got, want)
		}

		if got := ll.tail.data; got != val[0] {
			t.Errorf("tail.data = %v, want %v", got, val[0])
		}

		if got, want := ll.length, (i+1); got != want {
			t.Errorf("length = %v, want %v", got, want)
		}
	}
}

func TestAddTail(t *testing.T)  {
	ll := LinkedList{}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		want := val[i]
		ll.AddTail(want)
		
		if got := ll.tail.data; got != want {
			t.Errorf("tail.data = %v, want %v", got, want)
		}

		if got := ll.head.data; got != val[0] {
			t.Errorf("head.data = %v, want %v", got, val[0])
		}

		if got, want := ll.length, (i+1); got != want {
			t.Errorf("length = %v, want %v", got, want)
		}
	}
}

func TestRemoveHead(t *testing.T)  {
	ll := LinkedList{}

	if _, got := ll.RemoveHead(); got == nil {
		t.Errorf("RemoveHead() must return an error")
	}

	if _, got := ll.RemoveHead(); got.Error() != "empty list" {
		t.Errorf("RemoveHead() = %v, want \"empty list\"", got.Error())
	}

	val := []int{42, 24}
	for i := 0; i < len(val); i++ {
		ll.AddHead(val[i])
	}

	for i := (len(val) - 1); i >= 0; i-- {
		if got, _ := ll.RemoveHead(); got != val[i] {
			t.Errorf("RemoveHead() = %v, want %v", got, val[i])
		}

		if got := ll.length; got != i {
			t.Errorf("length = %v, want %v", got, i)
		}
	}
}

func TestRemoveTail(t *testing.T)  {
	ll := LinkedList{}

	if _, got := ll.RemoveTail(); got == nil {
		t.Errorf("RemoveTail() must return an error")
	}

	if _, got := ll.RemoveTail(); got.Error() != "empty list" {
		t.Errorf("RemoveTail() = %v, want \"empty list\"", got.Error())
	}

	val := []int{42, 24}
	for i := 0; i < len(val); i++ {
		ll.AddTail(val[i])
	}

	for i := (len(val) - 1); i >= 0; i-- {
		if got, _ := ll.RemoveTail(); got != val[i] {
			t.Errorf("RemoveTail() = %v, want %v", got, val[i])
		}

		if got := ll.length; got != i {
			t.Errorf("length = %v, want %v", got, i)
		}
	}
}

func TestPeekHead(t *testing.T)  {
	ll := LinkedList{}

	if _, got := ll.PeekHead(); got == nil {
		t.Errorf("PeekHead() must return an error")
	}

	if _, got := ll.PeekHead(); got.Error() != "empty list" {
		t.Errorf("PeekHead() = %v, want \"empty list\"", got.Error())
	}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		want := val[i]
		ll.AddHead(want)

		if got, _ := ll.PeekHead(); got != want {
			t.Errorf("PeekHead() = %v, want %v", got, want)
		}
	}
}

func TestPeekTail(t *testing.T)  {
	ll := LinkedList{}

	if _, got := ll.PeekHead(); got == nil {
		t.Errorf("PeekHead() must return an error")
	}

	if _, got := ll.PeekHead(); got.Error() != "empty list" {
		t.Errorf("PeekHead() = %v, want \"empty list\"", got.Error())
	}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		want := val[i]
		ll.AddTail(want)

		if got, _ := ll.PeekTail(); got != want {
			t.Errorf("PeekTail() = %v, want %v", got, want)
		}
	}
}

func TestToString(t *testing.T)  {
	ll := LinkedList{}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		ll.AddTail(val[i])
	}

	want := "42 -> 24"

	if got := ll.ToString(); got != want {
		t.Errorf("ToString() = %v, want %v", got, want)
	}
}

func TestEmpty(t *testing.T)  {
	ll := LinkedList{}

	want := true

	if got := ll.Empty(); got != want {
		t.Errorf("Empty() = %v, want %v", got, want)
	}

	val := []int{42, 24}

	for i := 0; i < len(val); i++ {
		ll.AddTail(val[i])
	}

	want = false

	if got := ll.Empty(); got != want {
		t.Errorf("Empty() = %v, want %v", got, want)
	}
}