package response

type Service interface {
	Insert() bool
	QueryList(params *map[string]interface{}) *[]interface{}
	Update(param *interface{}) bool
	Delete(id int) bool
	QueryById(id int) *interface{}
	Count() int
}
