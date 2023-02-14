package apply

const (
	VisibleClassify uint8 = 0 //0,可见性权限；
	postClassify    uint8 = 1 //1,发布权限；
	replyClassify   uint8 = 2 //0,回复权限；
	delClassify     uint8 = 3 //0,删除性权限；这个用于管理员。
)

/*
//web应用
通过对发布和回复以及可见性的权限组合成各种应用。
比如：法师专栏问答，

	发布是公共权限；
	回答则是自己权限；

又如：博客，

	发布是自己权限；
	回复是公共权限。

又如：私人日记，

	发布和回复以及可见性都是自己。
*/
type Apply struct {
	Name        string //应用名称
	Description string //应用描述
	Visible     uint8  //可见性权限。0，公共；1，群组；2，自己。
	Post        uint8  //发表权限。0，公共；1，群组；2，自己。
	Reply       uint8  //回复权限。0，公共；1，群组；2，自己。
	VGroup      Group  //可见性权限为群组时的用户列表。其他情况为nil
	PGroup      Group  //发表权限为群组时的用户列表。其他情况为nil
	RGroup      Group  //回复权限为群组时的用户列表。其他情况为nil
	DGroup      Group  //删除、屏蔽记录权限。
	Record      Record
}

func NewApply(name string) *Apply {
	//通过name读取db获取该Apply的对应属性。
	return &Apply{}
}

// 授权控制器
func (a *Apply) empower(userid int, classify uint8) bool { //classify uint8 //1,发布；2，回复,0,可见性；3，删除。
	groups := map[uint8]Group{
		0: a.VGroup,
		1: a.PGroup,
		2: a.RGroup,
		3: a.DGroup,
	}
	if g, ok := groups[classify]; ok {
		return g.Find(userid)
	}
	return false
}

// 添加记录
func (a *Apply) PostRecord(userid int, classify uint8) { //classify uint8 //1,发布；2，回复。
	if !a.empower(userid, classify) {
		return
	}
	a.Record.Write()
}

// 打开记录
func (a *Apply) OpenRecord(userid int) {
	if !a.empower(userid, VisibleClassify) {
		return
	}
	a.Record.Open()
}

// 删除记录
func (a *Apply) DelRecord(userid int) {
	if !a.empower(userid, delClassify) {
		return
	}
	a.Record.Delete()
}
