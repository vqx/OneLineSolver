package oneLineSolver

type Node struct {
	Id int

	//other info like x,y
}

type Line struct {
	Id      int
	NodeId1 int
	NodeId2 int
}

type Solver struct {
	NodeSlice []*Node
	LineSlice []*Line
	Solution  []int //node id slice
}