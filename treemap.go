package treemap

type CompareFunc func(a Key, b Key) int
type TraverseProjection func(k Key, v Data) bool

type TreeMap struct {
	root *node
	comp CompareFunc
}

func New(compFn CompareFunc) *TreeMap {
	return &TreeMap{
		comp: compFn,
	}
}

func (t *TreeMap) Insert(key Key, data Data) bool {
	if t.root == nil {
		t.root = newNode(nil, t.comp, key, data)
		return true
	}

	return t.root.insert(key, data)
}

func (t *TreeMap) Get(key Key) (Data, bool) {
	var d Data
	var ok bool

	t.Traverse(func(k Key, v Data) bool {
		if t.comp(key, key) == 0 {
			ok = true
			d = v
			return false
		}

		return true
	})

	return d, ok
}

func (t *TreeMap) Traverse(proj TraverseProjection) {
	t.traverse(t.root, proj)
}

func (t *TreeMap) traverse(n *node, proj TraverseProjection) {
	if n == nil {
		return
	}

	t.traverse(n.left, proj)
	cont := proj(n.data.key, n.data.data)
	if !cont {
		return
	}

	t.traverse(n.right, proj)
}
