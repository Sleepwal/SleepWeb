package frame

import "strings"

type node struct {
	pattern  string  // 待匹配路由，login/:name
	part     string  // 路由中一部分，:name
	children []*node // 子节点
	isWild   bool    // 是否精确匹配，含有 : 或 *
}

// matchChild
// @Description: 第一个匹配的节点，用于插入
// @receiver n
// @param part
// @return *node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren
// @Description: 所有匹配的节点，用于查找
// @receiver n
// @param part
// @return []*node
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert
// @Description: 递归查找每一层的节点，若无，则新建
// @receiver n
// @param pattern
// @param parts
// @param height
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search
// @Description: 递归搜索节点，遇到 * 退出
// @receiver n
// @param parts
// @param height
// @return *node
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
