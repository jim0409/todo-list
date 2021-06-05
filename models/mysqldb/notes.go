package mysqldb

import (
	"gorm.io/gorm"
)

/*
	宣告 NoteImp，透過 require 同一份檔案，告訴 mysqlDBObj 去實現對應的 interface
	實作方法時需要使用的還是 mysqlDBObj
*/
type NoteImp interface {
	// basic CRUD
	CreateNotes(string, string) error
	ReadAllNotes() (map[uint]interface{}, error)
	UpdateNotes(string, string, string) error
	DeleteNote(string) error

	// Big Func
	ReadNoteByPage(int, int) (map[uint]interface{}, error)
	CountPage(int64) (int64, error)
}

type NoteTable struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:varchar(255)"`
}

var (
	noteTable = "note_table"
)

func (db *NoteTable) TableName() string {
	return noteTable
}

func (db *mysqlDBObj) CreateNotes(title string, content string) error {
	var newNots = &NoteTable{
		Title:   title,
		Content: content,
	}

	return db.DB.Create(newNots).Error
}

func (db *mysqlDBObj) ReadAllNotes() (map[uint]interface{}, error) {
	var notes = []NoteTable{}
	if err := db.DB.Find(&notes).Error; err != nil {
		return nil, err
	}

	results := make(map[uint]interface{})
	for _, j := range notes {
		results[j.ID] = map[string]interface{}{
			"Title":   j.Title,
			"Content": j.Content,
		}
	}

	return results, nil
}

func (db *mysqlDBObj) UpdateNotes(id string, title string, content string) error {
	var updateNote = &NoteTable{
		Title:   title,
		Content: content,
	}

	return db.DB.Table(noteTable).Where("id = ?", id).Updates(updateNote).Error
}

func (db *mysqlDBObj) DeleteNote(id string) error {
	return db.DB.Where("id = ?", id).Delete(&NoteTable{}).Error
}

func (db *mysqlDBObj) ReadNoteByPage(page int, limit int) (map[uint]interface{}, error) {
	var notes = []NoteTable{}
	offset := page * limit

	if err := db.DB.Table(noteTable).Order("id").Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
		return nil, err
	}

	results := make(map[uint]interface{})
	for _, j := range notes {
		results[j.ID] = map[string]string{
			"Title":   j.Title,
			"Content": j.Content,
		}
	}

	return results, nil
}

// CountPage would return the total of pages
func (db *mysqlDBObj) CountPage(pageSize int64) (int64, error) {
	var value int64
	// 去計算所有的資料筆數不需要拿取全部資料，僅僅只要拿取 id 即可
	if err := db.DB.Table(noteTable).Select("id").Where("created_at is NOT NULL").Count(&value).Error; err != nil {
		return 0, err
	}

	totalPage := value / pageSize
	if value%pageSize != 0 {
		totalPage = totalPage + 1
	}

	return totalPage, nil
}
