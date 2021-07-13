package service

import "tsqqqqqq/blog/module"

type BlogService interface {
	Save(blog *module.Blog) bool
	QueryList(params *map[string]interface{}) *[]module.Blog
	Update(blog *module.Blog) bool
	Delete(id int) bool
	QueryById(id int) *module.Blog
	Count() int
}
