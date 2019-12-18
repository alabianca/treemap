package examples

import (
	"testing"
)

var comp = func(a int, b int) int {
	return b - a
}

func TestTreeIntMap_Insert(t *testing.T) {
	tree := NewIntMap(comp)
	tree.Insert(10, 1)
	tree.Insert(12, 1)
	tree.Insert(15, 4)
	tree.Insert(4, 1)

	v, ok := tree.Get(15)
	if !ok || v != 4{
		t.Fatalf("Expected %d, but got %d\n", 4, v)
	}
}

func Test_IsBlack(t *testing.T) {
	n := &node{
		black: true,
	}

	if !isBlack(n) {
		t.Fatalf("Expected node to be black to be true but got %v\n", isBlack(n))
	}

	if isBlack(nil) != true {
		t.Fatalf("Expected nil node to be black")
	}
}

func Test_Uncle(t *testing.T) {
	n0 := &node{
		parent: nil,
	}

	n1 := &node{parent: n0, isLeftChild: true}
	n2 := &node{parent: n0, isLeftChild: false}

	n0.left = n1
	n0.right = n2

	n3 := &node{parent: n2, isLeftChild: false}

	n2.right = n3


	unc0 := uncle(n0)
	unc1 := uncle(n1)
	unc2 := uncle(n2)
	unc3 := uncle(n3)

	if unc3 != n1 {
		t.Fatalf("Expected uncle to be n1 but got %v\n", unc3)
	}

	if unc0 != nil {
		t.Fatalf("Expected n0's uncle to be nil, but got %v\n", unc0)
	}

	if unc1 != nil {
		t.Fatalf("Expected n0's uncle to be nil, but got %v\n", unc1)
	}

	if unc2 != nil {
		t.Fatalf("Expected n0's uncle to be nil, but got %v\n", unc2)
	}


}

func Test_ToggleColor(t *testing.T) {
	n := &node{
		black: true,
	}

	toggleColor(n)

	if isBlack(n) {
		t.Fatalf("Expected node to be red after toggle, but is still black")
	}
}

func Test_RotateLeft(t *testing.T) {
	/*
		      n0
	         /  \
	            n1
	           /  \
	              n2
	             /  \
	*/

	n0 := &node{}
	n1 := &node{}
	n2 := &node{}

	n0.right = n1
	n1.parent = n0
	n1.right = n2
	n2.parent = n1

	rotateLeft(n0)

	/*
	              n1
	             /  \
	            n0  n2
	           /  \/  \
	*/

	//t.Logf("%s, %v\n", "n0", &n0)
	//t.Logf("%s, %v\n", "n1", &n1)
	//t.Logf("%s, %v\n", "n2", &n2)
	if n1.parent != nil {
		t.Fatalf("Expected n1's parent to be nil, but got %v\n", &n1.parent)
	}

	if n1.left != n0 {
		t.Fatalf("Expected n1's left child to be n0, but got %v\n", &n1.left)
	}

	if n1.right != n2 {
		t.Fatalf("Expected n1's right child to be n2, but got %v\n", &n1.right)
	}

	if n0.parent != n1 {
		t.Fatalf("Expected n0's parent to be n1, but got %v\n", (&n0.parent))
	}

	if n2.parent != n1 {
		t.Fatalf("Expected n2's parent to be n1, but got %v\n", &n2.parent)
	}

	if n0.left != nil {
		t.Fatalf("Expected n0's left to be nil")
	}

	if n0.right != nil {
		t.Fatalf("Expected n0's right to be nil")
	}

	if n2.left != nil {
		t.Fatalf("Expected n0's left to be nil")
	}

	if n2.right != nil {
		t.Fatalf("Expected n0's right to be nil")
	}
}

func Test_RotateRight(t *testing.T) {
	/*
			      n0
		         /  \
	            n1
	           /  \
	          n2
	         /  \
	*/

	n0 := &node{}
	n1 := &node{}
	n2 := &node{}

	n0.left = n1
	n1.parent = n0
	n1.left = n2
	n2.parent = n1

	rotateRight(n0)

	/*
	      n1
	     /  \
	    n2  n0
	   /  \/  \
	*/

	adr := map[interface{}]string {
		&n0: "n0",
		&n1: "n1",
		&n2: "n2",
	}

	if n1.parent != nil {
		v, _ := adr[&n1.parent]
		t.Fatalf("Expected n1's parent to be nil, but got %s\n", v)
	}

	if n1.left != n2 {
		v, _ := adr[&n1.left]
		t.Fatalf("Expected n1's left child to be n2, but got %s\n", v)
	}

	if n1.right != n0 {
		v, _ := adr[&n1.right]
		t.Fatalf("Expected n1's right child to be n0, but got %s\n", v)
	}

	if n0.parent != n1 {
		v, _ := adr[&n0.parent]
		t.Fatalf("Expected n0's parent to be n1, but got %s\n", v)
	}

	if n2.parent != n1 {
		v, _ := adr[&n2.parent]
		t.Fatalf("Expected n2's parent to be n1, but got %s\n", v)
	}

	if n0.left != nil {
		t.Fatalf("Expected n0's left to be nil")
	}

	if n0.right != nil {
		t.Fatalf("Expected n0's right to be nil")
	}

	if n2.left != nil {
		t.Fatalf("Expected n0's left to be nil")
	}

	if n2.right != nil {
		t.Fatalf("Expected n0's right to be nil")
	}
}

func TestTreeIntMap_InsertFixup(t *testing.T) {
	cases := []struct{
		TREE func() *TreeIntMap
		EXPECT []int
	}{
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(5, 5)
				tree.Insert(1, 1)
				return tree
			},
			EXPECT: []int{5, 1, 15},
		},
		{
				TREE: func() *TreeIntMap {
					tree := NewIntMap(comp)
					tree.Insert(15, 15)
					tree.Insert(100, 100)
					tree.Insert(12, 12)
					tree.Insert(10, 10)
					return tree
				},
				EXPECT: []int{15, 12, 100, 10},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				return tree
			},
			EXPECT: []int{15, 10, 100, 5, 12},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				tree.Insert(28, 28)
				return tree
			},
			EXPECT: []int{15, 10, 100, 5, 12, 28},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				tree.Insert(28, 28)
				tree.Insert(200, 200)
				return tree
			},
			EXPECT: []int{15, 10, 100, 5, 12, 28, 200},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				tree.Insert(28, 28)
				tree.Insert(200, 200)
				tree.Insert(16, 16)
				tree.Insert(17, 17)
				return tree
			},
			EXPECT: []int{15, 10, 100, 5, 12, 17, 200, 16, 28},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				tree.Insert(28, 28)
				tree.Insert(200, 200)
				tree.Insert(16, 16)
				tree.Insert(17, 17)
				tree.Insert(1, 1)
				return tree
			},
			EXPECT: []int{15, 10, 100, 5, 12, 17, 200, 1, 16, 28},
		},
		{
			TREE: func() *TreeIntMap {
				tree := NewIntMap(comp)
				tree.Insert(15, 15)
				tree.Insert(100, 100)
				tree.Insert(12, 12)
				tree.Insert(10, 10)
				tree.Insert(5, 5)
				tree.Insert(28, 28)
				tree.Insert(200, 200)
				tree.Insert(16, 16)
				tree.Insert(17, 17)
				tree.Insert(1, 1)
				tree.Insert(2, 2)
				return tree
			},
			EXPECT: []int{15, 10, 100, 2, 12, 17, 200, 1, 5, 16, 28},
		},
	}


	for _, c := range cases {
		tree := c.TREE()
		out := c.EXPECT
		index := 0
		tree.TraverseBF(func(key, val int) bool {
			if out[index] != key {
				t.Fatalf("Expected %d at index %d, but got %d\n", out[index], index, key)
			}
			index++
			return true
		})

	}
}
