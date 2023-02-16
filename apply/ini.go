package apply

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/liaoran123/xbdb"
)

var (
	table     map[string]*xbdb.Table
	ConfigMap map[string]interface{}
)

func Ini() {
	path := GetCurrentAbPath()                       //守护程序读取文件时需要绝对路径
	text, _ := ioutil.ReadFile(path + "config.json") //打开配置文件
	ConfigMap = make(map[string]interface{})
	json.Unmarshal(text, &ConfigMap)
	//打开或创建数据库
	dbpath := ConfigMap["dbpath"].(string)
	xbdb.OpenDb(dbpath)

	//创建表信息结构
	dbinfo := xbdb.NewTableInfo()

	//删除后添加=修改
	//dbinfo.Del("admin")

	//创建表
	createtbs(dbinfo)
	//打开表操作结构
	table = xbdb.OpenTableStructs()

	//打印数据库//用于测试代码
	//Table["u"].Select.For(Pr0)
	//Table["test"].Select.ForTb(Pr)

}

func Pr(k, v []byte) bool {
	fmt.Println(string(k), string(v))
	return true
}

func createtbs(dbinfo *xbdb.TableInfo) {

	if dbinfo.GetInfo("u").FieldType == nil { //创建用户表
		name := "u" //user                                                              //目录表
		fields := []string{
			"id",      //自动增值编码
			"email",   //邮箱
			"psw",     //密码
			"fahao",   //法号
			"jianjie", //简介
			"sj",      //创建时间
			"pass",    //是否通过。1，是；0否。
		}
		fieldType := []string{
			"int",
			"string",
			"string",
			"string",
			"string",
			"time",
			"bool",
		} //字段类型
		idxs := []string{"1"}  //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
		fullText := []string{} //考据级全文搜索索引字段的下标。
		ftlen := "7"           //全文搜索的长度，中文默认是7
		r := dbinfo.Create(name, ftlen, fields, fieldType, idxs, fullText)
		fmt.Printf("r: %v\n", r)

	}
	if dbinfo.GetInfo("apply").FieldType == nil { //创建应用表
		name := "apply" //应用表名称
		fields := []string{
			"id",          //自动增值编码
			"name",        //应用名称
			"userid",      //创建者id
			"description", //简述
			"permVisible", //可见性权限。0，公共；1，群组；2，自己。
			"permPost",    //发表权限。0，公共；1，群组；2，自己。
			"permReply",   //回复权限。0，公共；1，群组；2，自己。
			"postname",    //发布名称，如问答类，发布，叫做“提问”。默认为“发布”
			"replyname",   //回复名称，如问答类，回复，叫做“回答”。默认为“回复”
			"dtime",       //创建时间
		}
		//一一对应上面字段的类型
		fieldType := []string{
			"int",
			"string",
			"int",
			"string",
			"string",
			"string",
			"string",
			"string",
			"string",
			"time",
		} //字段类型
		idxs := []string{"1", "2"} ////索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
		fullText := []string{}     //考据级全文搜索索引字段的下标。
		ftlen := "7"               //全文搜索的长度，中文默认是7
		r := dbinfo.Create(name, ftlen, fields, fieldType, idxs, fullText)
		fmt.Printf("r: %v\n", r)

	}

	if dbinfo.GetInfo("record").FieldType == nil { //创记录表。所有发布、回复信息都是这个表
		name := "record" //记录表
		fields := []string{
			"id",      //自动增值编码
			"userid",  //发布者id
			"type",    //0,发布；1，回复
			"aid",     //应用id，
			"pid",     //父级id。被回复的记录id。
			"title",   //标题
			"content", //内容
			"views",   //浏览次数
			"dtime",   //时间
		}
		fieldType := []string{
			"int",
			"int",
			"int",
			"int",
			"int",
			"string",
			"string",
			"int",
			"time",
		}
		idxs := []string{"2", "3"} //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
		fullText := []string{"4"}  //考据级全文搜索索引字段的下标。
		ftlen := "7"               //全文搜索的长度，中文默认是7
		r := dbinfo.Create(name, ftlen, fields, fieldType, idxs, fullText)
		fmt.Printf("r: %v\n", r)
	}
	if dbinfo.GetInfo("group").FieldType == nil { //创建群组成员表
		name := "group"
		fields := []string{
			"id",       //自动增值编码
			"aid",      //应用id，
			"classify", ////0,可见性权限；1,发布权限；2,回复权限；3,删除权限；这个用于管理员。
			"userid",   //用户id
			"pass",     //通过。1，是；0，否。
			"dtime",    //时间
		}
		fieldType := []string{
			"int",
			"int",
			"int",
			"int",
			"bool",
			"time",
		} //字段类型
		idxs := []string{"1"}  //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
		fullText := []string{} //考据级全文搜索索引字段的下标。
		ftlen := "7"           //全文搜索的长度，中文默认是7
		r := dbinfo.Create(name, ftlen, fields, fieldType, idxs, fullText)
		fmt.Printf("r: %v\n", r)
	}
}
