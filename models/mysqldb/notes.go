package mysqldb

type NoteTable struct {
	Title   string `gorm:"primary_key"`
	Content string
}

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
