package main

import "testing"

func TestResult(t *testing.T) {
	order, err := topoSort(prereqs)
	if err != nil {
		t.Errorf("get toposort error %v", err)
	}
	// 1, 判断所有nodes仅出现一次
	nodes := make(map[string]bool)
	for k, v := range prereqs {
		nodes[k] = false
		for _, s := range v {
			nodes[s] = false
		}
	}

	if len(order) != len(nodes) {
		t.Errorf("toposort len error(%d), want %d", len(order), len(nodes))
	}
	for _, v := range order {
		if nodes[v] == true {
			t.Errorf("find twice %q node", v)
		}
		nodes[v] = true
	}

	// 2,如果a有到b的路径，则a排在b的前面
	pos := make(map[string]int)
	for k, v := range order {
		pos[v] = k
	}

	for a, v := range prereqs {
		for _, b := range v {
			if pos[a] < pos[b] {
				t.Errorf("node position error %q(%d) > %q(%d)", a, pos[a], b, pos[b])
			}
		}
	}
}
