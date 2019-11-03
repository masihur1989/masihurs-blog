package categories

import (
	"database/sql"

	"github.com/masihur1989/masihurs-blog/server/common"
)

// GetAllCategories godoc
func (cm CategoryModel) GetAllCategories(pagination common.Pagination) (*[]Category, error) {
	l.Started("GetAllCategories")
	db := common.GetDB()
	categories := make([]Category, 0)
	query := "SELECT * FROM categories"
	// TODO add paginations
	l.Info("Query: %s", query)
	results, err := db.Query(query)
	l.Info("DB RESULT: %v", results)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var category Category
		// for each row, scan the result into our tag composite object
		err = results.Scan(&category.ID, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			l.Errorf("ErrorScanning: ", err)
			return nil, common.ErrorScanning
		}
		categories = append(categories, category)
	}
	l.Debug("CATEGORIES: %v", &categories)
	l.Completed("GetAllCategories")
	return &categories, nil
}

// GetCategoryByID godoc
func (cm CategoryModel) GetCategoryByID(categoryID int) (*Category, error) {
	l.Started("GetCategoryByID")
	db := common.GetDB()
	var category Category

	stmt, err := db.Prepare("SELECT * FROM categories WHERE id = ?")
	if err != nil {
		l.Errorf("ErrorCreatingStmnt: ", err)
		return nil, common.ErrorCreatingStmnt
	}
	defer stmt.Close()
	err = stmt.QueryRow(categoryID).Scan(&category.ID, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows %d", nil, categoryID)
		return nil, sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}
	l.Debug("CATEGORY: %v", &category)
	l.Completed("GetCategoryByID")
	return &category, nil
}

// PostCategory godoc
func (cm CategoryModel) PostCategory(category Category) error {
	l.Started("PostCategory")
	l.Info("CATEGORY TO POST %v", category)
	db := common.GetDB()
	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}

	_, err = tx.Exec("INSERT INTO categories(name, active, created_at, updated_at) VALUES(?,?,NOW(), NOW());", category.Name, category.Active)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("PostCategory")
	return nil
}

// UpdateCategory godoc
func (cm CategoryModel) UpdateCategory(ID int, category Category) error {
	l.Started("UpdateCategory")
	l.Info("CATEGORY TO UPDATE %v", category)
	db := common.GetDB()

	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}
	_, err = tx.Exec("UPDATE categories SET name = ?, active = ?, updated_at = NOW() WHERE id = ?;", category.Name, category.Active, ID)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("UpdateCategory")
	return nil
}

// DeleteCategory godoc
func (cm CategoryModel) DeleteCategory(categoryID int) error {
	l.Started("DeleteCategory")
	db := common.GetDB()
	q := `DELETE FROM categories WHERE id = ?;`
	result, err := db.Exec(q, categoryID)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return common.ErrorQuery
	}
	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	l.Completed("DeleteCategory")
	return err
}
