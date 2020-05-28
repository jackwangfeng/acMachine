package acMachine

type AcNode struct {
	tireNum   int
	value     rune
	isPattern bool
	father    *AcNode
	fail      *AcNode
	next      map[rune]*AcNode
}

func NewAcNode() *AcNode {
	return &AcNode{0, 0, false, nil, nil, make(map[rune]*AcNode)}
}

type AcMachine struct {
	root *AcNode
}

func NewAcMachine() *AcMachine {
	return &AcMachine{root: NewAcNode()}
}

func (a *AcMachine) AddPattern(p string) {
	chars := []rune(p)
	if a.root == nil {
		a.root = NewAcNode()
	}
	f := a.root
	pLen := len(chars)
	for i := 0; i < pLen; i++ {
		var ok bool
		var tmp *AcNode
		tmp, ok = f.next[chars[i]]
		if !ok {
			tmp = NewAcNode()
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

func (a *AcMachine) getFail(node *AcNode) {
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

func (a *AcMachine) Build() {
	//build tired tree
	for _, v := range a.root.next {
		a.getFail(v)
	}
}

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
	m.Build()
	results, pos := m.Match("abcdefabcdef")
	cLen := len(results)
	for i := 0; i < cLen; i++ {
		fmt.Printf("%d:%s\n", pos[i], results[i])
	}
}
*/
