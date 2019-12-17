package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Key   string
	Name  string
	Value string
}

func main() {
	var d data
	var out string
	flag.StringVar(&d.Key, "key", "", "The subtype used for the treemap being generated")
	flag.StringVar(&d.Value, "value", "", "The value used for the treemap being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the treemap being generated")
	flag.StringVar(&out, "out", "out.go", "The filename of the generated treemap")

	flag.Parse()
	t := template.Must(template.New("treemap").Parse(treeTemplate))

	fout, err := os.Create(out)
	if err != nil {
		panic(err)
	}

	defer fout.Close()
	t.Execute(fout, &d)
}

var treeTemplate = `
package treemap{{.Name}}

type CompareFunc func(a {{.Key}}, b {{.Key}}) int
type TraverseFunc func(k {{.Key}}, v {{.Value}}) bool

// TreeMap implementation

type TreeMap struct {
	root *node
	comp CompareFunc
}

func New{{.Name}}Map(compFn CompareFunc) *TreeMap {
	return &TreeMap{
		comp: compFn,
	}
}

func (t *TreeMap) Insert(key {{.Key}}, data {{.Value}}) bool {
	if t.root == nil {
		t.root = newNode(nil, t.comp, key, data)
		return true
	}

	return t.root.insert(key, data)
}

func (t *TreeMap) Get(key {{.Key}}) ({{.Value}}, bool) {
	var d {{.Value}}
	var ok bool

	t.Traverse(func(k {{.Key}}, v {{.Value}}) bool {
		if t.comp(key, key) == 0 {
			ok = true
			d = v
			return false
		}

		return true
	})

	return d, ok
}

func (t TreeMap) String() string {
	buf := new(bytes.Buffer)
	t.Traverse(func(k {{.Key}}, v {{.Value}}) bool {
		buf.WriteString(fmt.Sprintf("%s(%s)->", k, v))
		return true
	})

	return buf.String()
}

func (t *TreeMap) Traverse(proj Traverse{{.Name}}Func) {
	t.traverse(t.root, proj)
}

func (t *TreeMap) traverse(n *node, proj Traverse{{.Name}}Func) {
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

// Node Implementation. Everything must be private at this point
type nodeData struct {
	key  {{.Key}}
	data {{.Value}}
}

type node struct {
	right  *node
	left   *node
	parent *node
	data   nodeData
	comp   CompareFunc
}

func newNode(parent *node, c CompareFunc, key {{.Key}}, data {{.Value}}) *node {
	return &node{
		right:  nil,
		left:   nil,
		parent: parent,
		data: nodeData{
			data: data,
			key:  key,
		},
		comp: c,
	}
}

func (n *node) insert(key {{.Key}}, data {{.Value}}) bool {
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

`
