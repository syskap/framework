package project

/* Comments on node traversal

When dealing with a given node in a tree, the way you move through the tree is either vertically (moving between siblings) or horizontally (moving between parents and children).

As shown in a compass below:

		prev (sibling above)
				^
				|
parent<- current node ->children
				|
				v
		next (sibling below)

Or as in a tree below:

parent
├──...
├──...
├──prev (sibling above)
├──current node
│  └──children
│     ├──...
│     └──...
├──next (sibling below)
└──...

Walking the whole structure might look like:

func (n *Node) WalkTree(fn func(FileNode)) {
	fn(n)
	for _, c := range n.Children() {
		c.WalkTree(fn)
	}
	n.Next().WalkTree(fn)
}

func CreateDir(node FileNode) {
	// creates directory
}

n := &Node{Name: "adapter"}
n.WalkTree(CreateDir)
WalkTree(n, CreateDir)
*/

type FileNode interface {
	Parent() FileNode
	Children() []FileNode
	Siblings() []FileNode
	Sibling() FileNode
	Prev() FileNode
	Next() FileNode
	WalkTree(func(FileNode))
	AddParent(FileNode)
	AddChild(FileNode)
}

type Node struct {
	Name     string
	parent   FileNode
	children []FileNode
}

func (n *Node) Parent() FileNode {
	if n == nil {
		return nil
	}
	return n.parent
}

func (n *Node) Children() []FileNode {
	if n == nil {
		return nil
	}
	return n.children
}

func (n *Node) AddParent(fn FileNode) bool {
	if fn == nil {
		return false
	}
	return true
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

var (
	ProjectName     string
	BaseDefaultNode = Node{}
)
