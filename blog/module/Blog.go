package module

// Blog 用户结构体
type Blog struct {
	Id      int    `json:"id" field:"id"`           //主键
	Title   string `json:"title" field:"title"`     //标题
	Content string `json:"content" field:"content"` //内容
	Created string `json:"created" field:"created"` //创建时间
	Updated string `json:"updated" field:"updated"` //修改时间
	UsersId int    `json:"usersId" field:"usersId"` //用户ID 外键关联
}
