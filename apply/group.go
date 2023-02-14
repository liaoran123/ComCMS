package apply

//群组
type Group struct {
	Users []User //授权的用户id列表
}

func NewGroup(Name string) *Group { //Name 应用名称 对应的授权群组
	return &Group{}
}
func (g *Group) Find(id int) bool {
	return false
}
