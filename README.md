# golang-apm-sql

go get github.com/middleware-labs/golang-apm-sql

```golang

import (
    mw_sql "github.com/middleware-labs/golang-apm-sql/sql"
	"github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	tracker.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
	)
	db, err := mw_sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
}