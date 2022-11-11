package main

import (
	_ "database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	mw_sql "github.com/middleware-labs/golang-apm-sql/sql"
	"github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

func main() {
	r := gin.Default()
	config, _ := tracker.Track(
		track.WithConfigTag("service", "Your service name"),
		track.WithConfigTag("projectName", "Your project name"),
	)
	r.Use(g.Middleware(config))
	db, err := mw_sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	r.GET("/todo", func(c *gin.Context) {
		ctx := c.Request.Context()
		tracker.SpanFromContext(ctx)
		var num int
		if err := db.QueryRowContext(ctx, "SELECT 42").Scan(&num); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, "ok")
	})
	_ = r.Run(":8081")
}
