package treemap

type Key interface {
	String() string
}

type Data interface {
	String() string
}

type nodeData struct {
	key Key
	data Data
}

type node struct {
	right *node
	left *node
	parent *node
	data nodeData
	comp CompareFunc
}

func newNode(parent *node, c CompareFunc, key Key, data Data) *node {
	return &node{
		right:  nil,
		left:   nil,
		parent: parent,
		data:   nodeData{
			data: data,
			key: key,
		},
		comp:   c,
	}
}

func (n *node) insert(key Key, data Data) bool {
	res := n.comp(n.data.key, key)
	if res < 0 {
		if n.left != nil {
			return n.left.insert(key, data)
		}
		n.left = newNode(n, n.comp, key, data)
		return true
	}

	if res > 0 {
		if n.right != nil {
			return n.right.insert(key, data)
		}
		n.right = newNode(n, n.comp, key, data)
		return true
	}

	return false
}
