package acmachine

import "container/list"

type acNode struct {
	tireNum   int
	value     rune
	isPattern bool
	father    *acNode
	fail      *acNode
	next      map[rune]*acNode
}

//NewAcNode create a AcNode pointer.
func newAcNode() *acNode {
	return &acNode{0, 0, false, nil, nil, make(map[rune]*acNode)}
}

//AcMachine AcAutoMachine.
type AcMachine struct {
	root *acNode
}

//NewAcMachine create a acMachine.
func NewAcMachine() *AcMachine {
	return &AcMachine{root: newAcNode()}
}

//AddPattern add a pattern
func (a *AcMachine) AddPattern(p string) {
	chars := []rune(p)
	if a.root == nil {
		a.root = newAcNode()
	}
	f := a.root
	pLen := len(chars)
	for i := 0; i < pLen; i++ {
		var ok bool
		var tmp *acNode
		tmp, ok = f.next[chars[i]]
		if !ok {
			tmp = newAcNode()
			tmp.tireNum = i
			tmp.value = chars[i]
			tmp.father = f
			if i == 0 {
				tmp.fail = a.root
			}
			f.next[chars[i]] = tmp
		}
		if i == pLen-1 {
			tmp.isPattern = true
		}
		f = tmp
	}
}

func (a *AcMachine) getFail(node *acNode) {
	if node.father != a.root {
		tmpNode, ok := node.father.fail.next[node.value]
		if ok {
			node.fail = tmpNode
		} else {
			node.fail = a.root
		}
	}
	for _, v := range node.next {
		a.getFail(v)
	}
}

//Build 用递归深度搜索
func (a *AcMachine) Build() {
	//build tired tree
	for _, v := range a.root.next {
		a.getFail(v)
	}
}

//Build2 使用压栈方式实现深度遍历
func (a *AcMachine) Build2() {
	stack := list.New()
	stack.PushFront(a.root)
	for stack.Len() > 0 {
		tmp := stack.Front()
		node := tmp.Value.(*acNode)
		stack.Remove(tmp)
		if node != a.root && node.father != a.root {
			tmpNode, ok := node.father.fail.next[node.value]
			if ok {
				node.fail = tmpNode
			} else {
				node.fail = a.root
			}
		}
		for _, v := range node.next {
			stack.PushFront(v)
		}
	}
}

//Build1 使用广度搜索
func (a *AcMachine) Build1() {
	queue := []*acNode{}
	queue = append(queue, a.root)
	for len(queue) > 0 {
		//把quue的节点都拿出来, 求每个节点的fail节点
		tmpLen := len(queue)
		tmpQueue := make([]*acNode, tmpLen)
		copy(tmpQueue, queue)
		queue = queue[0:0]
		for i := 0; i < tmpLen; i++ {
			if tmpQueue[i] != a.root && tmpQueue[i].father != a.root {
				tmpNode, ok := tmpQueue[i].father.fail.next[tmpQueue[i].value]
				if ok {
					tmpQueue[i].fail = tmpNode
				} else {
					tmpQueue[i].fail = a.root
				}
			}
			for _, v := range tmpQueue[i].next {
				queue = append(queue, v)
			}
		}
	}
}

//Match 匹配接口
func (a *AcMachine) Match(con string) (results []string, pos []int) {
	chars := []rune(con)
	cLen := len(chars)
	var i int
	f := a.root
	for {
		if i >= cLen {
			break
		}
		v, ok := f.next[chars[i]]
		if ok {
			if v.isPattern {
				start := i - v.tireNum
				str := string([]rune(con)[start : i+1])
				pos = append(pos, start)
				results = append(results, str)
			}
			i++
			f = v
		} else {
			if f == a.root {
				i++
			} else {
				f = f.fail
			}
		}
	}
	return
}

/*
func main() {
	m := AcMachine{NewAcNode()}
	m.AddPattern("abc")
	m.AddPattern("cde")
	m.Build2()
	results, pos := m.Match("abcdefabcdef")
	cLen := len(results)
	for i := 0; i < cLen; i++ {
		fmt.Printf("%d:%s\n", pos[i], results[i])
	}
}
*/
