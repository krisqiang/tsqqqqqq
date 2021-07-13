package impl

import (
	response "tsqqqqqq/base"
	"tsqqqqqq/blog/module"
)

type BlogServiceImpl struct {
}

func (p *BlogServiceImpl) Save(blog *module.Blog) (b bool) {
	service := new(response.Impl)
	service.P = blog
	res := service.Insert()
	if index, err := res.RowsAffected(); err == nil {
		if index > 0 {
			return true
		}
	} else {
		panic(err)
	}

	return
}

func (p *BlogServiceImpl) QueryList(params *map[string]interface{}) (res *[]module.Blog) {
	return
}

func (p *BlogServiceImpl) Update(blog *module.Blog) (res bool) {
	return
}

func (p *BlogServiceImpl) Delete(id int) (res bool) {
	return
}
func (p *BlogServiceImpl) QueryById(id int) (res *module.Blog) {
	return
}

func (p *BlogServiceImpl) Count() (res int) {
	return
}
