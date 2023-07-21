package hash_table

import "testing"

func TestGetAdditiveHash(t *testing.T) {

	want1 := "Hello"
	want2 := "World"

	ht := new(AdditiveHasher{}, 8)
	ht.Add(want1)
	ht.Add(want2)

	s1 := ht.Get(want1)
	s2 := ht.Get(want2)

	if s1 != want1 {
		t.Errorf("want %s, got %s", want1, s1)
	}

	if s2 != want2 {
		t.Errorf("want %s, got %s", want2, s2)
	}
}

// TestAdditiveHashUniformity isn't really something we want to test for,
// it just demonstrates that the algorithm isn't uniform such that the
// hash for "foo" is the same as for "oof", for example.
func TestAdditiveHashNotUniform(t *testing.T) {

	want := "Hello, World!"
	get := "World, Hello!"

	ht := new(AdditiveHasher{}, 8)
	ht.Add(want)
	got := ht.Get(get) // get a different string with the same hash

	if want != got {
		t.Errorf("string %s is the same as %s", want, get)
	}
}

func TestGetFoldingHash(t *testing.T) {

	want1 := "Hello"
	want2 := "World"

	ht := new(FoldingHasher{}, 8)
	ht.Add(want1)
	ht.Add(want2)

	s1 := ht.Get(want1)
	s2 := ht.Get(want2)

	if s1 != want1 {
		t.Errorf("want %s, got %s", want1, s1)
	}

	if s2 != want2 {
		t.Errorf("want %s, got %s", want2, s2)
	}
}

func TestFoldingHashUniform(t *testing.T) {

	str := "Hello, World!"
	different := "World, Hello!"

	ht := new(FoldingHasher{}, 8)
	ht.Add(str)
	got := ht.Get(different)

	if different == got {
		t.Errorf("string %s is the same as %s", different, str)
	}
}
