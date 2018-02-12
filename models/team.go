package models

import (
	"database/sql"
	"fmt"
	"hash/crc32"
)

type (
	Team struct {
		ID          int    `json:"id,omitempty" sql:"AUTO_INCREMENT"`
		sid         int    `json:-`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		deleted     int    `json:-`
	}

	TeamList struct {
		teams []*Team
	}
)

// отображаем имя команды
func (t *Team) String() string {
	return t.Name
}

// проверяем заполнение команды
func (t *Team) Validate() error {

	if t.Name == "" {
		return fmt.Errorf("Title of command is empty")
	}

	return nil
}

// определяем номер шарда
func (t *Team) ShardNumber() int {

	crc := crc32.ChecksumIEEE([]byte(t.Name))

	if crc%2 == 0 {
		return 0
	}

	return 1

}

// определяем номер шарда
func (t *Team) Create(db *sql.DB) error {

	result, err := db.Exec(
		"INSERT INTO teams (`name`, `description`, `sid`, `deleted`) VALUES (?, ?, 0, 0)",
		t.Name,
		t.Description,
	)

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	t.ID = int(lastID)

	if err != nil {
		return err
	}

	return nil
}

// определяем номер шарда
func (t *Team) Update(db *sql.DB) error {

	result, err := db.Exec(
		"UPDATE teams SET"+
			"`description` = ?"+
			//			",`updated` = ?"+
			"WHERE id = ?",
		t.ID,
	)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func (tl *TeamList) List(db *sql.DB) error {

	items := []*Team{}

	rows, err := db.Query("SELECT TOP 10 id, name, description FROM teams")

	for rows.Next() {
		t := &Team{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description)

		if err != nil {
			return err
		}

		items = append(items, t)
	}
	// надо закрывать соединение, иначе будет течь
	rows.Close()

	return nil

}
