package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Comment struct {
	Id         int64     `json:"id" db:"id,primarykey,autoincrement"`
	Name       string    `json:"name" db:"name,notnull,size: 200"`
	Text       string    `json:"text" db:"text,notnull"`
	Created_at time.Time `json:"created_at" db:"created_at,notnull"`
	Update_at  time.Time `json:"updated_at" db:"updated_at,notnull"`
}

func Route(e *echo.Echo) {
	// データベースの初期化
	dbmap := initDb()

	// 静的ファイルの保存先指定
	e.Static("/", "static/")

	// コメント一覧
	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment
		_, err := dbmap.Select(&comments, "SELECT * FROM comments ORDER BY created_at desc LIMIT 10")
		if err != nil {
			c.Logger().Error("Select : ", err)
			return c.String(http.StatusBadRequest, "Select :"+err.Error())
		}

		return c.JSON(http.StatusOK, comments)
	})

	// コメント登録
	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			c.Logger().Error("Bind :", err)
			return c.String(http.StatusBadRequest, "Bind :"+err.Error())
		}

		if err := dbmap.Insert(&comment); err != nil {
			c.Logger().Error("Insert :", err)
			return c.String(http.StatusBadRequest, "Insert :"+err.Error())
		}

		if err := c.Validate(&comment); err != nil {
			c.Logger().Error("Validate :", err)
			return c.String(http.StatusBadRequest, "Validate :"+err.Error())
		}

		c.Logger().Infof("ADDED : %v", comment.Id)
		return c.JSON(http.StatusCreated, "")
	})
}

// DB初期設定
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_echo?parseTime=true")
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Comment{}, "comments")

	// テーブルがない場合作成
	err = dbmap.CreateTablesIfNotExists()
	return dbmap
}

// フック関数
// インサート
func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	c.Created_at = time.Now()
	c.Update_at = c.Created_at
	return nil
}

// アップデート
func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.Update_at = time.Now()
	return nil
}
