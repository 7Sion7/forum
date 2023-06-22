package models

import "time"

type Content struct {
	ID         int
	Title      string
	Content    string
	CategoryID int
	UserID     int
	CreatedAt  time.Time
	UpdatedAt  string
}

type Posts struct {
	content []Content
}

var AllPosts *Posts

func getUserContent(ID int) (Content, error) {
	var content Content

	stmt, err := db.Prepare("SELECT title FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var title string

	err = stmt.QueryRow(ID).Scan(&title)
	if err != nil {
		return Content{}, err
	}

	stmt, err = db.Prepare("SELECT content FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var post string

	err = stmt.QueryRow(ID).Scan(&post)
	if err != nil {
		return Content{}, err
	}

	stmt, err = db.Prepare("SELECT created_at FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var timeCreated time.Time

	err = stmt.QueryRow(ID).Scan(&timeCreated)
	if err != nil {
		return Content{}, err
	}

	content = Content{
		Title: title,
		Content: post,
		CreatedAt: timeCreated,
	}

	return content, nil
}

func GetAllContent() (interface{}, error){
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Content
	for rows.Next() {
		var post Content
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CategoryID, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	AllPosts = &Posts{content: posts}
	return AllPosts, nil
}