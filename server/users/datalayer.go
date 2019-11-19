package users

import (
	"database/sql"

	"github.com/masihur1989/masihurs-blog/server/common"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers godoc
func (um UserModel) GetAllUsers(pagination common.Pagination) (*[]User, error) {
	l.Started("GetAllUsers")
	db := common.GetDB()
	users := make([]User, 0)
	query := `SELECT * FROM users`
	// TODO make pagination to work
	// arg := make(map[string]interface{})
	// query, arg = common.ApplyPaginationToQuery(query, arg, pagination)
	l.Info("Query: %s", query)
	results, err := db.Query(query)
	l.Info("DB RESULT: %v", results)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			l.Errorf("ErrorScanning: ", err)
			return nil, common.ErrorScanning
		}
		users = append(users, user)
	}
	l.Debug("USERS: %v", &users)
	l.Completed("GetAllUsers")
	return &users, nil
}

// GetUserByID godoc
func (um UserModel) GetUserByID(userID int) (*User, error) {
	l.Started("GetUserByID")
	db := common.GetDB()
	var user User

	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		l.Errorf("ErrorCreatingStmnt: ", err)
		return nil, common.ErrorCreatingStmnt
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows %d", nil, userID)
		return nil, sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}
	l.Completed("GetUserByID")
	return &user, nil
}

// PostUser godoc
func (um UserModel) PostUser(user User) error {
	l.Started("PostUser")
	l.Info("USER TO POST %v", user)
	db := common.GetDB()
	hPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("HASHING ERROR: ", err)
		return common.ErrorHashing
	}
	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}

	_, err = tx.Exec("INSERT INTO users(name, email, password, remember_token, login_type, active, created_at, updated_at) VALUES(?,?,?,?,?,?,NOW(), NOW());", user.Name, user.Email, hPwd, user.RememberToken, user.LoginType, user.Active)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("PostUser")
	return nil
}

// PostLogin godoc
func (um UserModel) PostLogin(ul UserLogin) (*User, error) {
	l.Started("PostLogin")
	db := common.GetDB()
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE email = ?", ul.Email)
	l.Info("ROW: %v", row)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows: ", err)
		return nil, sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ul.Password))
	if err != nil {
		l.Errorf("PASSWORD VALDIATION FAILED: ", err)
		return nil, common.ErrorPasswordMatching
	}
	l.Info("USER: %v", &user)
	l.Completed("PostLogin")
	return &user, nil
}

// UpdateUserPassword -
func (um UserModel) UpdateUserPassword(userID int, pwd ForgotPassword) error {
	l.Started("UpdateUserPassword")
	db := common.GetDB()
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	l.Info("ROW: %v", row)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RememberToken, &user.LoginType, &user.Active, &user.CreatedAt, &user.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows: ", err)
		return sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return common.ErrorScanning
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.Password))
	if err != nil {
		l.Errorf("PASSWORD VALDIATION FAILED: ", err)
		return err
	}
	q, err := db.Prepare("UPDATE users SET password=? WHERE id = ?")
	if err != nil {
		l.Errorf("ErrorCreatingStmnt: ", err)
		return common.ErrorCreatingStmnt
	}
	defer q.Close()
	hPwd, err := bcrypt.GenerateFromPassword([]byte(pwd.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("HASHING ERROR: ", err)
		return common.ErrorHashing
	}
	_, err = q.Exec(hPwd, userID)
	if err != nil {
		l.Errorf("ERROR UPDATING RECORD: ", err)
		return err
	}
	l.Completed("UpdatePassword")
	return nil
}

// DeleteUser godoc
func (um UserModel) DeleteUser(userID int) error {
	l.Started("DeleteUser")
	db := common.GetDB()
	q := `DELETE FROM users WHERE id = ?;`
	result, err := db.Exec(q, userID)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return common.ErrorQuery
	}
	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	l.Completed("DeleteUser")
	return err
}
