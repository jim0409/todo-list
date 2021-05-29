package mysqldb

import (
	"github.com/jinzhu/gorm"
)

/*
	宣告 NoteImp，透過 require 同一份檔案，告訴 mysqlDBObj 去實現對應的 interface
	實作方法時需要使用的還是 mysqlDBObj
*/
type NoteImp interface {
	CreateNotes(string, string) error
	ReadAllNotes() (map[string]string, error)
	UpdateNotes(string, string) error
	DeleteNotes(string) error

	ReadNoteByPage(int, int) (map[uint]interface{}, error)
}

type NoteTable struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	Title   string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:varchar(255)"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
}

var (
	NoteTableName = "note_table"
)

func (db *mysqlDBObj) CreateNotes(title string, content string) error {
	var newNots = &NoteTable{
		Title:   title,
		Content: content,
	}

	return db.DB.Create(newNots).Error
}

func (db *mysqlDBObj) ReadAllNotes() (map[string]string, error) {
	var notes = []NoteTable{}
	if err := db.DB.Find(&notes).Error; err != nil {
		return nil, err
	}

	results := map[string]string{}
	for _, j := range notes {
		results[j.Title] = j.Content
	}

	return results, nil
}

func (db *mysqlDBObj) UpdateNotes(title string, content string) error {
	var updateNote = &NoteTable{
		Title:   title,
		Content: content,
	}

	return db.DB.First(&NoteTable{Title: title}).Update(updateNote).Error
}

func (db *mysqlDBObj) DeleteNotes(title string) error {
	return db.DB.Where("Title = ?", title).Delete(&NoteTable{}).Error
}

func (db *mysqlDBObj) ReadNoteByPage(id int, limit int) (map[uint]interface{}, error) {
	var notes = []NoteTable{}

	if err := db.DB.Table(NoteTableName).Limit(limit).Find(&notes).Error; err != nil {
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
