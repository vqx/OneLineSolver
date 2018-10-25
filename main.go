package main

import (
	"oneLineSolver"
	"fmt"
)

func main() {
	s := oneLineSolver.Solver{
		NodeSlice:[]*oneLineSolver.Node{
			{
				Id:1,
			},
			{
				Id:2,
			},
			{
				Id:3,
			},
			{
				Id:4,
			},
			{
				Id:5,
			},
		},
		LineSlice:[]*oneLineSolver.Line{
			{
				Id:11,
				NodeId1:1,
				NodeId2:2,
			},
			{
				Id:12,
				NodeId1:1,
				NodeId2:3,
			},
			{
				Id:13,
				NodeId1:2,
				NodeId2:3,
			},
			{
				Id:14,
				NodeId1:2,
				NodeId2:4,
			},
			{
				Id:15,
				NodeId1:3,
				NodeId2:5,
			},
			{
				Id:16,
				NodeId1:4,
				NodeId2:5,
			},

		},
		Solution:[]int{},
	}
	if s.Run() {
		fmt.Println(s.Solution)
	}
}
