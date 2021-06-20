package mysqldb

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type Auth struct {
	Enforcer *casbin.SyncedEnforcer
}

type AuthImp interface {
	CheckPolicy(string, string, string) (bool, error)
}

var (
	authTable    = "auth"
	syncEnforcer *casbin.SyncedEnforcer
)

func (db *mysqlDBObj) CheckPolicy(sub, obj, act string) (bool, error) {
	return syncEnforcer.Enforce(sub, obj, act)
}

func (db *mysqlDBObj) GetRolesForUser(usr string) ([]string, error) {
	return syncEnforcer.GetRolesForUser(usr)
}

func initAuth(db *gorm.DB, path string) error {
	a, err := gormadapter.NewAdapterByDBUseTableName(db, authTable, "")
	if err != nil {
		return err
	}

	e, err := casbin.NewSyncedEnforcer(path, a)
	if err != nil {
		return err
	}

	err = e.LoadPolicy()
	if err != nil {
		return err
	}

	syncEnforcer = e
	injectDefaultData(e)

	return nil
}

func injectDefaultData(e *casbin.SyncedEnforcer) {
	var ok bool
	var err error

	// TODO : suppose to construct admin policy alone with router-metrics?
	// AdminRouter := map[string][]string{
	// 	"admin": []string{"/v1/note/add", "POST"},
	// }

	// 宣告 user/normal 的權限
	ok, err = e.AddRoleForUser("user", "normal")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	ok, err = e.AddPolicy("normal", "/v1/note/add", "POST")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	ok, err = e.AddPolicy("normal", "/v1/note", "GET")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	ok, err = e.AddPolicy("normal", "/v1/note/lists", "GET")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	ok, err = e.AddPolicy("normal", "/v1/note/totalpages", "GET")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}

	// 宣告 root/admin 的權限
	ok, err = e.AddRoleForUser("root", "admin")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	// root 也具備 normal 的權限
	ok, err = e.AddRoleForUser("admin", "normal")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	// 正則表達式
	ok, err = e.AddPolicy("admin", "/v1/note/update/", "PUT")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
	// 正則表達式
	ok, err = e.AddPolicy("admin", "/v1/note/delete/", "DELETE")
	if err != nil && !ok {
		log.Fatalf("init error %v__%v\n", ok, err)
	}
}
