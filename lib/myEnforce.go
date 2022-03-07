package lib

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/dengh57/gorm-adapter/v3"
	"gorm.io/gorm"
)

const tableRuleName = "casbin_rule"

func GetEnforce(id int64, DB *gorm.DB) (*casbin.Enforcer, error) {

	a, err := gormadapter.NewFilteredAdapter("mysql", "root:1234@tcp(127.0.0.1:3306)/", "auth", "policies")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	m := Model{}
	DB.Model(&Model{}).First(&m, id)

	text := "[request_definition]\n" + m.RequestDefinition + "\n[policy_definition]\n" + m.PolicyDefinition +
		"\n[role_definition]\n" + m.RoleDefinition + "\n[policy_effect]\n" + m.PolicyEffect + "\n[matchers]\n" + m.Matchers

	modelOfCasbin, _ := model.NewModelFromString(text)
	e, _ := casbin.NewEnforcer(modelOfCasbin, a)
	return e, nil
}
