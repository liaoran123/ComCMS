package apply

//群组
type Group struct {
	Users []User //授权的用户id列表	
}

func NewGroup(Name string) *Group { //Name 应用名称 对应的授权群组
	return &Group{}
}

//查看用户是否存在
func (g *Group) Find(id int) bool {
	for _, u := range g.Users {
		if u.id == id {
			return true
		}
	}
	return false
}
