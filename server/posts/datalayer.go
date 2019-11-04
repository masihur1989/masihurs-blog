package posts

import (
	"database/sql"

	"github.com/masihur1989/masihurs-blog/server/common"
)

// GetAllPosts godoc
func (pm PostModel) GetAllPosts(pagination common.Pagination) (*[]Post, error) {
	l.Started("GetAllPosts")
	db := common.GetDB()
	posts := make([]Post, 0)
	query := "SELECT * FROM posts"
	// TODO add paginations
	l.Info("Query: %s", query)
	results, err := db.Query(query)
	l.Info("DB RESULT: %v", results)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var post Post
		// for each row, scan the result into our tag composite object
		err = results.Scan(&post.ID, &post.Title, &post.Body, &post.UserID, &post.CategoryID, &post.PostView, &post.Active, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			l.Errorf("ErrorScanning: ", err)
			return nil, common.ErrorScanning
		}
		posts = append(posts, post)
	}
	l.Debug("POSTS: %v", &posts)
	l.Completed("GetAllPosts")
	return &posts, nil
}

// GetPostByID godoc
func (pm PostModel) GetPostByID(postID int) (*Post, error) {
	l.Started("GetPostByID")
	db := common.GetDB()
	var post Post

	stmt, err := db.Prepare("SELECT * FROM posts WHERE id = ?")
	if err != nil {
		l.Errorf("ErrorCreatingStmnt: ", err)
		return nil, common.ErrorCreatingStmnt
	}
	defer stmt.Close()
	err = stmt.QueryRow(postID).Scan(&post.ID, &post.Title, &post.Body, &post.UserID, &post.CategoryID, &post.PostView, &post.Active, &post.CreatedAt, &post.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		l.Errorf("ErrNoRows %d", nil, postID)
		return nil, sql.ErrNoRows
	case err != nil:
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}
	l.Debug("POST: %v", &post)
	l.Completed("GetPostByID")
	return &post, nil
}

// PostPost godoc
func (pm PostModel) PostPost(post Post) error {
	l.Started("PostPost")
	l.Info("POST TO POST %v", post)
	db := common.GetDB()
	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}

	_, err = tx.Exec("INSERT INTO posts(title, body, user_id, category_id, post_view, active, created_at, updated_at) VALUES(?,?,?,?,?,?,NOW(), NOW());", post.Title, post.Body, post.UserID, post.CategoryID, post.PostView, post.Active)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("PostPost")
	return nil
}

// UpdatePost godoc
func (pm PostModel) UpdatePost(ID int, post Post) error {
	l.Started("UpdatePost")
	l.Info("POST TO UPDATE %v", post)
	db := common.GetDB()

	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}
	_, err = tx.Exec("UPDATE posts SET title = ?, body = ?, user_id = ?, category_id = ?, active = ?, updated_at = NOW() WHERE id = ?;", post.Title, post.Body, post.UserID, post.CategoryID, post.Active, ID)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("UpdatePost")
	return nil
}

// DeletePost godoc
func (pm PostModel) DeletePost(postID int) error {
	l.Started("DeletePost")
	db := common.GetDB()
	q := `DELETE FROM posts WHERE id = ?;`
	result, err := db.Exec(q, postID)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return common.ErrorQuery
	}
	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	l.Completed("DeletePost")
	return err
}

// UpdatePostViewByID godoc
func (pm PostModel) UpdatePostViewByID(postID int, postView PostView) error {
	l.Started("UpdatePostViewByID")
	l.Info("POST VIEW: %v", postView)
	db := common.GetDB()

	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}
	_, err = tx.Exec("UPDATE posts SET post_view = ?, updated_at = NOW() WHERE id = ?;", postView.CurrentView+1, postID)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("UpdatePostViewByID")
	return nil
}

// AddPostTags godoc
func (pm PostModel) AddPostTags(postID int, tagIDs PostTags) error {
	l.Started("AddPostTags")
	l.Info("POST TAGS: %v", tagIDs)
	db := common.GetDB()
	for _, tag := range tagIDs.Tags {
		tx, err := db.Begin()
		if err != nil {
			l.Errorf("TRANSACTION ERROR: ", err)
			return common.ErrorTransaction
		}
		_, err = tx.Exec("INSERT INTO post_tags(post_id, tag_id, created_at, updated_at) VALUES(?,?,NOW(),NOW());", postID, tag)
		if err != nil {
			l.Errorf("TX EXECUTION ERROR:", err)
			tx.Rollback()
			return common.ErrorTransaction
		}
		tx.Commit()
	}
	l.Completed("AddPostTags")
	return nil
}

// GetPostTags godoc
func (pm PostModel) GetPostTags(postID int) (*PostTags, error) {
	l.Started("GetPostTags")
	db := common.GetDB()
	pTags := make([]int, 0)
	query := "SELECT tag_id FROM post_tags WHERE post_id = ?"
	l.Info("Query: %s", query)
	results, err := db.Query(query, postID)
	l.Info("DB RESULT: %v", results)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return nil, common.ErrorQuery
	}

	for results.Next() {
		var tagID int
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tagID)
		if err != nil {
			l.Errorf("ErrorScanning: ", err)
			return nil, common.ErrorScanning
		}
		pTags = append(pTags, tagID)
	}
	l.Debug("POST TAGS: %v", &pTags)
	l.Completed("GetPostTags")
	pt := PostTags{pTags}
	return &pt, nil
}

// UpdatePostTags godoc
// TODO not sure about the implementation
func (pm PostModel) UpdatePostTags(postID int, tagIDs []int) error {
	l.Started("UpdatePostTags")
	l.Info("TAG IDs: %v", tagIDs)
	db := common.GetDB()
	for _, tag := range tagIDs {
		tx, err := db.Begin()
		if err != nil {
			l.Errorf("TRANSACTION ERROR: ", err)
			return common.ErrorTransaction
		}
		_, err = tx.Exec("INSERT INTO post_tags(post_id, tag_id, created_at, updated_at) VALUES(?,?,NOW(),NOW());", postID, tag)
		if err != nil {
			l.Errorf("TX EXECUTION ERROR:", err)
			tx.Rollback()
			return common.ErrorTransaction
		}
		tx.Commit()
	}
	l.Completed("UpdatePostTags")
	return nil

}

// GetPostLikes godoc
func (pm PostModel) GetPostLikes(ID int) (*PostLikes, error) {
	l.Started("GetPostLikes")
	var pl PostLikes
	db := common.GetDB()
	q := "SELECT COUNT(*) FROM likes WHERE post_id = ?;"
	err := db.QueryRow(q, ID).Scan(&pl.Likes)
	if err != nil {
		l.Errorf("ErrorScanning: ", err)
		return nil, common.ErrorScanning
	}
	l.Debug("POST LIKES: %v", &pl)
	l.Completed("GetPostLikes")
	return &pl, nil
}

// LikePost godoc
func (pm PostModel) LikePost(ID int, pl PostLike) error {
	l.Started("LikePost")
	db := common.GetDB()
	tx, err := db.Begin()
	if err != nil {
		l.Errorf("TRANSACTION ERROR: ", err)
		return common.ErrorTransaction
	}
	_, err = tx.Exec("INSERT INTO likes(user_id, post_id, created_at, updated_at) VALUES(?,?,NOW(),NOW());", &pl.UserID, ID)
	if err != nil {
		l.Errorf("TX EXECUTION ERROR:", err)
		tx.Rollback()
		return common.ErrorTransaction
	}
	tx.Commit()
	l.Completed("LikePost")
	return nil
}

// DeletePostLike godoc
func (pm PostModel) DeletePostLike(ID int, pl PostLike) error {
	l.Started("LikePost")
	db := common.GetDB()
	q := `DELETE FROM likes WHERE post_id = ? AND user_id = ?;`
	result, err := db.Exec(q, ID, pl.UserID)
	if err != nil {
		l.Errorf("ErrorQuery: ", err)
		return common.ErrorQuery
	}
	// didn't hit any rows, return a 404
	deleteCount, err := result.RowsAffected()

	if deleteCount == 0 {
		return sql.ErrNoRows
	}
	l.Completed("LikePost")
	return nil
}
