package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"os"
)

var (
	Enforcer *casbin.Enforcer
)

//创建casbin的enforcer
func SetupCasbinEnforcer() (error) {
	dir, _ := os.Getwd()
	modelPath := dir + "/config/rbac_model.conf"
	csvPath := dir + "/config/rbac2.csv"
	fmt.Println("modelPath:"+modelPath);
	fmt.Println("csvPath:"+csvPath);
	var errC error
	Enforcer, errC = casbin.NewEnforcer(modelPath, csvPath)
	fmt.Printf("RBAC test start\n") // output for debug
	if (errC != nil) {
		fmt.Println(errC)
		return errC
	} else {
		Enforcer.EnableLog(true)
		return nil
	}
}
