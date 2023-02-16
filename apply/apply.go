package apply

import (
	"fmt"
	"strconv"

	"github.com/liaoran123/xbdb"
)

const (
	permVisible uint8 = 0 //0,可见性权限；
	permPost    uint8 = 1 //1,发布权限；
	permReply   uint8 = 2 //2,回复权限；
	permdel     uint8 = 3 //3,删除权限；这个用于管理员。
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
	table        map[string]*xbdb.Table
	name         string //应用名称
	creater      int    //创建者id
	description  string //应用描述
	permVisible  uint8  //可见性权限。0，公共；1，群组；2，自己。
	permPost     uint8  //发表权限。0，公共；1，群组；2，自己。
	permReply    uint8  //回复权限。0，公共；1，群组；2，自己。
	groupVisible Group  //可见性权限为群组时的用户列表。其他情况为nil
	groupPost    Group  //发表权限为群组时的用户列表。其他情况为nil
	groupReply   Group  //回复权限为群组时的用户列表。其他情况为nil
	groupDel     Group  //删除、屏蔽记录权限。即管理员。
	record       Record
}

func NewApply(name string) *Apply {
	if table == nil {
		table = xbdb.OpenTableStructs()
	}
	fsidx := table[name].Ifo.GetFieldIds([]string{"name"})
	tbd := table[name].Select.WhereIdx([]byte("name"), []byte(name), false, 0, 1, fsidx, false)
	tbmap := table[name].RDtoMap(tbd.Rd[0])
	fmt.Printf("tbmap: %v\n", tbmap)

	//通过name读取db获取该Apply的对应属性。
	return &Apply{
		name: name,
	}
}

// 授权控制器
func (a *Apply) empower(userid int, classify uint8) bool { //classify uint8 //1,发布；2，回复,0,可见性；3，删除。
	groups := map[uint8]Group{
		0: a.groupVisible,
		1: a.groupPost,
		2: a.groupReply,
		3: a.groupDel,
	}
	if g, ok := groups[classify]; ok {
		return g.Find(userid)
	}
	return false
}

// 添加记录
func (a *Apply) PostRec(paras map[string]string) { //0，公共；1，群组；2，自己
	userid, _ := strconv.Atoi(paras["userid"])
	if a.permPost == 1 {

		if !a.empower(userid, permPost) {
			return
		}
	}
	if a.permPost == 2 {
		if a.creater != userid {
			return
		}
	}
	a.record.Write(paras)

}

// 打开记录
func (a *Apply) OpenRec(paras map[string]string) { //0，公共；1，群组；2，自己
	userid, _ := strconv.Atoi(paras["userid"])
	if a.permVisible == 1 {
		if !a.empower(userid, permVisible) {
			return
		}
	}
	if a.permVisible == 2 {
		if a.creater != userid {
			return
		}
	}
	a.record.Open(paras)
}

// 删除记录只有群组权限。
// 删除、屏蔽记录权限。即管理员，也是一个群组。
func (a *Apply) DelRec(paras map[string]string) {
	userid, _ := strconv.Atoi(paras["userid"])
	if !a.empower(userid, permdel) {
		return
	}
	a.record.Del(paras)
}
