package shortcut

import "github.com/huandu/skiplist"

func Intersection(lists ...*skiplist.SkipList) *skiplist.SkipList {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	result := skiplist.New(skiplist.Uint64)
	nodes := make([]*skiplist.Element, len(lists)) //给每条SkipList分配一个指针，从前往后遍历

	for i, list := range lists {
		if list == nil || list.Len() == 0 { //只要lists中有一条是空链，则交集为空
			return nil
		}
		nodes[i] = list.Front()
	}
	for {
		maxList := make(map[int]struct{}, len(nodes)) //此刻，哪个指针对应的值最大（最大者可能存在多个，所以用map）
		var maxValue uint64 = 0
		for i, node := range nodes {
			if node.Key().(uint64) > maxValue {
				maxValue = node.Key().(uint64)
				maxList = map[int]struct{}{i: {}} //可以用一对大括号表示空结构体实例
			} else if node.Key().(uint64) == maxValue {
				maxList[i] = struct{}{}
			}
		}
		if len(maxList) == len(nodes) { //所有node的值都一样大，则新诞生一个交集
			result.Set(nodes[0].Key(), nodes[0].Value)
			for i, node := range nodes { //所有node均需往后移
				nodes[i] = node.Next()
				if nodes[i] == nil {
					return result
				}
			}
		} else {
			for i, node := range nodes {
				if _, exists := maxList[i]; !exists { //值大的不动，小的往后移
					nodes[i] = node.Next() //不能用node=node.Next()，因为for range取得的是值拷贝
					if nodes[i] == nil {   //只要有一条SkipList已走到最后，则说明不会再有新的交集诞生，可以return了
						return result
					}
				}
			}
		}
	}
}

func UnionSet(lists ...*skiplist.SkipList) *skiplist.SkipList {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	result := skiplist.New(skiplist.Uint64)
	keySet := make(map[any]struct{}, 1000)
	for _, list := range lists {
		if list == nil {
			continue
		}
		node := list.Front()
		for node != nil {
			if _, exists := keySet[node.Key()]; !exists {
				result.Set(node.Key(), node.Value)
				keySet[node.Key()] = struct{}{}
			}
			node = node.Next()
		}
	}
	return result
}
