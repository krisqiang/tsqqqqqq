package impl

import (
	"tsqqqqqq/login/module"
	module2 "tsqqqqqq/rule/module"
	"tsqqqqqq/util"
)

type LoginImpl struct {
}

func (lm *LoginImpl) GetToken(u *module.Users) (res string) {
	rows, err := util.Query("select r.id,r.rule,r.created,r.updated from rule r "+
		"inner join rule_user ru on r.id = ru.ruleId "+
		"inner join users u on ru.usersId = u.id "+
		"where username = ? and password = ?", u.Username, u.Password)
	if err != nil {
		panic(err)
	}
	r := new(module2.Rule)
	if rows.Next() {
		var id int
		var rule string
		var created string
		var updated string
		err := rows.Scan(&id, &rule, &created, &updated)
		if err != nil {
			panic(err)
		}
		r.Id = id
		r.Rules = rule
		r.Created = created
		r.Updated = updated
	}
	res = util.GenerateToken(r)
	return
}
