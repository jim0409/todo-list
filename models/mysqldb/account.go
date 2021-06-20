package mysqldb

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type AccountImp interface {
	// basic CRUD
	CreateUser(string, string, string, string) error
	GetUserInfo(string) (map[string]interface{}, error)

	UpdateUserInfos(map[string]interface{}) error

	// need to verify auth again ...
	UpdateUserPassword(string, string) error

	// admin action
	ChangeUserStatus(string, uint8) error
	UpdateUserRole(string, string) error
	DeleteUser(string) error
}

// Name 不允許改名，系統上 Name 為 PrimvKey
type Account struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);comment:'使用者姓名'" mapstructure:"name"`
	// NickName string
	Status   uint8  `gorm:"type:tinyint(1);default:1;comment:'用戶狀態(正常/禁用,默認正常)'"`
	Role     string `gorm:"type:varchar(8);default:normal;comment:'角色(admin/normal.預設normal)'"` // 對應 casbin 的 id .. 所屬的 role
	Password string `gorm:"type:varchar(10);comment:'使用者密碼hash'"`
	Phone    string `gorm:"type:varchar(16);comment:'使用者手機號'" mapstructure:"phone,omitempty"`
	Mail     string `gorm:"type:varchar(100);comment:'使用者電郵'" mapstructure:"mail,omitempty"`
	Avatar   string `gorm:"type:varchar(20);comment:'大頭貼'" mapstructure:"avatar,omitempty"`
	Intro    string `gorm:"type:varchar(255);comment:'簡介'" mapstructure:"intro,omitempty"`
}

type Session struct {
	gorm.Model        // create_at = login_time
	UserID     int    `gorm:"type:tinyint(8)"`
	Token      string `gorm:"type:varchar(32)"`
}

var (
	accountTable = "account_table"
	roleMap      = map[string]bool{
		"admin":  true,
		"normal": true,
	}
)

func (db *Account) TableName() string {
	return accountTable
}

func (db *mysqlDBObj) CreateUser(name string, pw string, phone string, mail string) error {
	ac := &Account{
		Name:     name,
		Password: pw,
		Phone:    phone,
		Mail:     mail,
	}
	return db.DB.Create(ac).Error

}

func (db *mysqlDBObj) GetUserInfo(name string) (map[string]interface{}, error) {
	ac := &Account{}
	if err := db.DB.Table(accountTable).Where("name = ? and deleted_at is NULL", name).Scan(ac).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"name":   ac.Name,
		"role":   ac.Role,
		"phone":  ac.Phone,
		"mail":   ac.Mail,
		"status": ac.Status,
		"avatar": ac.Avatar,
		"intro":  ac.Intro,
	}, nil
}

// 更換 phone 及 password 的流程要另外走驗證
// TODO: use map[string]interface and convert with mapstruct
// mapstructure="name,omitempty"
func (db *mysqlDBObj) UpdateUserInfos(m map[string]interface{}) error {
	ac := &Account{}
	name, ok := m["name"]
	if !ok {
		return fmt.Errorf("username %v not found!\n", name)
	}
	if err := mapstructure.Decode(m, ac); err != nil {
		return err
	}

	// 有沒有可能要更新的對象不存在? ... 答案是否定的，因為 name 一定要從 session 過來
	return db.DB.Debug().Updates(ac).Where("name = ?", name).Error
}

func (db *mysqlDBObj) UpdateUserPassword(string, string) error {
	return nil
}

/*
	[admin] 行為
	有沒有可能要更新的對象不存在? ...
	通城能做做 DeleteUser 的情境是 Admin 想要砍掉某些使用者
	使用者即便不存在也只會噴 records not found 故不影響具體業務
*/
func (db *mysqlDBObj) ChangeUserStatus(name string, status uint8) error {
	ac := &Account{
		Status: status,
	}
	return db.DB.Debug().Table(accountTable).Where("name = ? and deleted_at is NULL", name).Updates(ac).Error
}

func (db *mysqlDBObj) DeleteUser(name string) error {
	ac := &Account{
		Name: name,
	}
	return db.DB.Debug().Table(accountTable).Where("name = ?", name).Delete(ac).Error
}

func (db *mysqlDBObj) UpdateUserRole(name string, role string) error {
	if roleMap[role] {
		ac := &Account{
			Role: role,
		}
		return db.DB.Table(accountTable).Where("name = ? and deleted_at is NULL", name).Updates(ac).Error
	}
	return fmt.Errorf("update role not found")
}
