package oneLineSolver

func (s *Solver)Run() bool {
	for _, node := range s.NodeSlice {
		if s.Solve(node.Id, node.Id, []int{node.Id}, []int{}) {
			return true
		}
	}
	return false
}

func (s *Solver)Solve(startId int, nowId int, nodeHistory []int, usedLine []int) bool {
	if len(nodeHistory) == len(s.LineSlice) + 1 {
		s.Solution = nodeHistory
		return true
	}
	lineSlice := s.GetLineSliceByNodeId(nowId)
	for _, l := range lineSlice {
		flag := false
		for _, ul := range usedLine {
			if ul == l.Id {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		nId := l.NodeId1
		if l.NodeId1 == nowId {
			nId = l.NodeId2
		}
		return s.Solve(startId, nId, append(nodeHistory, nId), append(usedLine, l.Id))
	}
	return false
}

func (s *Solver)GetLineSliceByNodeId(id int) []*Line {
	result := []*Line{}
	for _, l := range s.LineSlice {
		if l.NodeId1 == id || l.NodeId2 == id {
			result = append(result, l)
		}
	}
	return result
}


/*
穷举
遍历点
	开始点
		遍历线
			分级到各个连接的点
			先一直找到最深的层次
			如果当前层没有正确的解决方案
			退回上一层
方案检测
 点记录=点个数+1
 并且所有线都用了

 点历史  线历史
*/