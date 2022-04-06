package bintree

import (
	"strconv"
	"strings"
)

type BinTree struct {
	Val      int
	Left     *BinTree
	Right    *BinTree
	Parent   *BinTree
	offsetX  int
	rendered bool
}

func BintreeFromSlice(vals []int) *BinTree {
	if len(vals) == 0 {
		return nil
	}
	root := &BinTree{
		Val: vals[0],
	}
	for i := 1; i < len(vals); i++ {
		root.insert(vals[i])
	}
	return root
}

func (b *BinTree) insert(i int) {
	if i > b.Val {
		if b.Right == nil {
			b.Right = &BinTree{Val: i, Parent: b}
		} else {
			b.Right.insert(i)
		}
	} else {
		if b.Left == nil {
			b.Left = &BinTree{Val: i, Parent: b}
		} else {
			b.Left.insert(i)
		}
	}
}

func (b *BinTree) String() string {
	return strconv.Itoa(b.Val)
}

func (b *BinTree) Render() string {
	var s strings.Builder
	b.render(&s, 0)
	return s.String()
}

func (b *BinTree) render(sb *strings.Builder, offset int) {
	if b == nil {
		return
	}
	offset = b.printNode(sb, offset)
	b.Right.render(sb, offset)
	if b.Left != nil {
		sb.WriteString("\n")
		p := b
		offsets := map[int]struct{}{}
		for p != nil {
			if p.Left != nil && !p.Left.rendered {
				offsets[p.offsetX] = struct{}{}
			}
			p = p.Parent
		}
		b.Left.rendered = true
		for n := 1; n <= offset; n++ {
			if n == offset {
				sb.WriteString("\\")
			} else if _, ok := offsets[n]; ok {
				sb.WriteString("|")
			} else {
				sb.WriteString(" ")
			}
		}
		b.Left.render(sb, offset)
	}
}

func (b *BinTree) printNode(sb *strings.Builder, offsetX int) int {
	var rep string
	if b.Parent == nil {
		rep = b.String()
	} else {
		rep = " -- " + b.String()
	}
	sb.WriteString(rep)
	b.offsetX = offsetX + len(rep)
	return b.offsetX
}
