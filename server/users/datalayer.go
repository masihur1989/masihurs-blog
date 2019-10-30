package users

import (
	"database/sql"
	"log"

	"github.com/masihur1989/masihurs-blog/server/common"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers godoc
// TODO make pagination to work
func (um UserModel) GetAllUsers(pagination common.Pagination) (*[]User, error) {
	log.Println("GetAllUsers")
	db := common.GetDB()
	defer common.CloseDB()
	users := make([]User, 0)
	query := `SELECT * FROM users`
	// arg := make(map[string]interface{})

	// query, arg = common.ApplyPaginationToQuery(query, arg, pagination)
	results, err := db.Query(query)
	log.Printf("DB RESULT: %v", results)
	if err != nil {
		log.Printf("ErrorQuery: %v\n", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("ErrorScanning: %v", err)
			return nil, common.ErrorScanning
		}
		users = append(users, user)
	}
	log.Printf("USERS: %v\n", &users)
	return &users, nil
}

// GetUserByID godoc
func (um UserModel) GetUserByID(userID int) (*User, error) {
	log.Println("GetUserByID")
	db := common.GetDB()
	var user User

	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		log.Printf("ErrorCreatingStmnt: %v\n", err)
		return nil, common.ErrorCreatingStmnt
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("ErrNoRows %d\n", userID)
		return nil, sql.ErrNoRows
	case err != nil:
		log.Fatalf("ErrorScanning: %v\n", err)
		return nil, common.ErrorScanning
	}
	return &user, nil
}

// PostUser godoc
func (um UserModel) PostUser(user User) error {
	log.Println("PostUser")
	log.Printf("USER TO POST: %v\n", user)
	db := common.GetDB()
	hPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("HASHING ERROR: %v\n", err)
		return common.ErrorHashing
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("TRANSACTION ERROR: %v\n", err)
		return common.ErrorTransaction
	}

	_, err = tx.Exec("INSERT INTO users(name, email, password, remember_token, login_type, active, created_at, updated_at) VALUES(?,?,?,?,?,?,NOW(), NOW());", user.Name, user.Email, hPwd, user.RememberToken, user.LoginType, user.Active)
	if err != nil {
		log.Printf("TX EXECUTION ERROR: %v\n", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	return nil
}

// UpdateUserPassword -
func (um UserModel) UpdateUserPassword(userID int, pwd ForgotPassword) error {
	log.Println("UpdateUserPassword")
	db := common.GetDB()
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	log.Printf("ROW: %v\n", row)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("ErrNoRows: %v\n", err)
		return sql.ErrNoRows
	case err != nil:
		log.Fatalf("ErrorScanning: %v\n", err)
		return common.ErrorScanning
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.Password))
	if err != nil {
		log.Printf("PASSWORD VALDIATION FAILED: %v", err)
		return err
	}
	q, err := db.Prepare("UPDATE users SET password=? WHERE id = ?")
	if err != nil {
		log.Printf("ErrorCreatingStmnt: %v\n", err)
		return common.ErrorCreatingStmnt
	}
	defer q.Close()
	hPwd, err := bcrypt.GenerateFromPassword([]byte(pwd.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("HASHING ERROR: %v\n", err)
		return common.ErrorHashing
	}
	_, err = q.Exec(hPwd, userID)
	if err != nil {
		log.Printf("ERROR UPDATING RECORD: %v\n", err)
		return err
	}
	return nil
}

// DeleteUser -
func (um UserModel) DeleteUser(userID int) error {
	log.Println("DeleteUser")
	db := common.GetDB()
	q := `DELETE FROM users WHERE id = ?;`
	result, err := db.Exec(q, userID)

	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	return err
}
