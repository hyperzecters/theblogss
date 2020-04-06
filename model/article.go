package model

import (
	"database/sql"
	"fmt"
	db "go-project/database"
	"time"
)

// Article model
type Article struct {
	ID          int
	Title       string
	Content     string
	CreatedAt   sql.NullTime
	IsPublished bool
	PublishedAt sql.NullTime
	UpdatedAt   sql.NullTime
	User        User
}

// Create : Create Article
func (a *Article) Create() (int, error) {
	var id int
	conn, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	row := conn.QueryRow("INSERT INTO articles (title, content, created_at, id_user, is_published) VALUES ($1, $2, $3, $4, $5) RETURNING id", a.Title, a.Content, time.Now(), a.User.ID, false)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update : Update Article
func (a *Article) Update() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `UPDATE articles SET 
		title = $1,
		content = $2,
		updated_at = $3
	WHERE id = $4`
	conn.QueryRow(sql, a.Title, a.Content, time.Now(), a.ID)

	return nil
}

// Delete : Delete Article
func (a *Article) Delete() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `DELETE FROM articles WHERE id = $1`
	conn.QueryRow(sql, a.ID)

	return nil
}

// GetArticle : Get Article
func (a *Article) GetArticle(id int) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM articles WHERE id = $1", id)
	err = row.Scan(&a.ID, &a.Title, &a.Content, &a.CreatedAt, &a.IsPublished, &a.PublishedAt, &a.UpdatedAt, &a.User.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = a.User.GetUser(a.User.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// Publish : Publish article
func (a *Article) Publish() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.QueryRow("UPDATE articles SET is_published = $1, published_at = $2 WHERE id = $3", true, time.Now(), a.ID)

	return nil
}

// Unpublish : Unpublish article
func (a *Article) Unpublish() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.QueryRow("UPDATE articles SET is_published = $1 WHERE id = $2", false, a.ID)

	return nil
}

// Validate : Article validation
func (a *Article) Validate() (map[string]string, bool) {
	errMessages := map[string]string{}
	err := false

	return errMessages, err
}

// Get All Articles
func GetAllArticles() ([]Article, error) {
	var articles []Article
	var article Article
	var userID int

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT 
		id,
		title,
		LEFT (content, 100) AS content,
		created_at,
		is_published,
		published_at,
		updated_at,
		id_user
	FROM articles WHERE is_published IS true ORDER BY id DESC`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.IsPublished, &article.PublishedAt, &article.UpdatedAt, &userID)
		article.User.GetUser(userID)
		articles = append(articles, article)
	}

	return articles, nil
}

// Get All Articles
func GetAllArticlesByUser(idUser int) ([]Article, error) {
	var articles []Article
	var article Article

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT 
		id,
		title,
		LEFT (content, 50) AS content,
		created_at,
		is_published,
		published_at,
		updated_at,
		id_user
	FROM articles WHERE 
		id_user = $1
	ORDER BY id DESC`, idUser)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.IsPublished, &article.PublishedAt, &article.UpdatedAt, &idUser)
		article.User.GetUser(idUser)
		articles = append(articles, article)
	}

	return articles, nil
}
