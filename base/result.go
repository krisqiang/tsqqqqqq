// Package response
// @Title response
// @Description 封装同一返回结果，适用于前后端分离场景
// @AuThor Ken ${DATE}${TIME}
package response

// Result 同一返回结果的封装结构体
type Result struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //操作消息
	Data interface{} `json:"data"` //返回数据
}

// 默认数据常量
const (
	success = 200
	error   = 500
	succMsg = "操作成功"
	errMsg  = "操作失败"
)

// Success 统一返回操作成功，用于简单接口，不需要返回数据
// @return 返回当前结构体指针，用于包内其他方法的调用。
func (res *Result) Success() *Result {
	res.Code = success
	res.Msg = succMsg
	return res
}

// Error 统一返回操作失败，用于简单接口，不需要返回数据
// @return 返回当前结构体指针，用于包内其他方法的调用。
func (res *Result) Error() *Result {
	res.Code = error
	res.Msg = errMsg
	return res
}

// Put  链式函数方式，用于存放返回数据
// @param data是一个空接口 即java语言中的Object类型
// @return 返回当前结构体指针，用于包内其他方法的调用。
func (res *Result) Put(data interface{}) *Result {
	res.Data = data
	return res
}

// Set 链式函数方式，用于修改result结构体的code 和 msg 属性
// params param 是一个空接口类型的值 使用空接口断言判断值的类型，并修改对应属性值
// @return 返回当前结构体指针，用于包内其他方法的调用。
func (res *Result) Set(param interface{}) *Result {
	switch v := param.(type) {
	case int:
		res.Code = v
	case string:
		res.Msg = v
	default:
		panic("传入类型异常,Set方法仅可传入 int 、string 类型参数")
	}
	return res
}
