package util

type Node struct {
	Id       int
	ParentId int
	Name     string
	Children []*Node
}

func GetTree(data []Node, pid int) []*Node {
	var list []*Node
	for _, v := range data {
		if v.ParentId == pid {
			v.Children = GetTree(data, v.Id)
			list = append(list, &v)
		}
	}
	return list

}
