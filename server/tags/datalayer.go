package tags

import (
	"database/sql"

	"github.com/masihur1989/masihurs-blog/server/common"
)

// GetAllTags godoc
func (cm TagModel) GetAllTags(pagination common.Pagination) (*[]Tag, error) {
	l.Started("GetAllTags")
	db := common.GetDB()
	tags := make([]Tag, 0)
	query := "SELECT * FROM tags"
	// TODO add paginations
	l.Info("Query: %s", query)
	results, err := db.Query(query)
	l.Info("DB RESULT: %v", results)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name, &tag.Active, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			l.Errorf("ErrorScanning: ", err)
			return nil, common.ErrorScanning
		}
		tags = append(tags, tag)
	}
	l.Debug("TAGS: %v", &tags)
	l.Completed("GetAllTags")
	return &tags, nil
}

// GetTagByID godoc
func (cm TagModel) GetTagByID(tagID int) (*Tag, error) {
	l.Started("GetTagByID")
	db := common.GetDB()
	var tag Tag

	stmt, err := db.Prepare("SELECT * FROM tags WHERE id = ?")
	if err != nil {
		l.Errorf("ErrorCreatingStmnt: ", err)
		return nil, common.ErrorCreatingStmnt
	}
	defer stmt.Close()
	err = stmt.QueryRow(tagID).Scan(&tag.ID, &tag.Name, &tag.Active, &tag.CreatedAt, &tag.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows %d", nil, tagID)
		return nil, sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}
	l.Debug("TAG: %v", &tag)
	l.Completed("GetTagByID")
	return &tag, nil
}

// PostTag godoc
func (cm TagModel) PostTag(tag Tag) error {
	l.Started("PostTag")
	l.Info("TAG TO POST %v", tag)
	db := common.GetDB()
	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}

	_, err = tx.Exec("INSERT INTO tags(name, active, created_at, updated_at) VALUES(?,?,NOW(), NOW());", tag.Name, tag.Active)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("PostTag")
	return nil
}

// UpdateTag godoc
func (cm TagModel) UpdateTag(ID int, tag Tag) error {
	l.Started("UpdateTag")
	l.Info("TAG TO UPDATE %v", tag)
	db := common.GetDB()

	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}
	_, err = tx.Exec("UPDATE tags SET name = ?, active = ?, updated_at = NOW() WHERE id = ?;", tag.Name, tag.Active, ID)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("UpdateTag")
	return nil
}

// DeleteTag godoc
func (cm TagModel) DeleteTag(tagID int) error {
	l.Started("DeleteTag")
	db := common.GetDB()
	q := `DELETE FROM tags WHERE id = ?;`
	result, err := db.Exec(q, tagID)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return common.ErrorQuery
	}
	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	l.Completed("DeleteTag")
	return err
}
