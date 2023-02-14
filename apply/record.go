package apply

//记录,发布和回复都是一条记录
type Record struct {
	classify uint8 //0,发布；1，回复。
	//Title, Content string //标题，内容
}

func NewRecord() *Record {
	return &Record{}
}
func (r *Record) Write() {

}
func (r *Record) Open() {

}
