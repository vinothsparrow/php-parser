package linkedtree

type Visitor interface {
	VisitNode(n Node, depth int) bool
}