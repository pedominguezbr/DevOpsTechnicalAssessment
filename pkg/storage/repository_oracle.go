package storage

import (
	"database/sql"
	"fmt"
	"framework-go/pkg/config"

	_ "github.com/godror/godror"
)

func NewDbOracle(c *config.Config) ([]*sql.DB, error) {

	arrayConnections := make([]*sql.DB, 0)

	var db *sql.DB

	var err error

	if len(c.DB.Oracle) > 0 {
		for _, v := range c.DB.Oracle {
			psqlConnStr := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=%s)(PORT=%s)))(CONNECT_DATA=(SERVICE_NAME=%s)))",
				v.Username,
				v.Password,
				v.Host,
				v.Port,
				v.Database)

			db, err = sql.Open("godror", psqlConnStr)

			if err != nil {
				fmt.Println(err)
			}

			arrayConnections = append(arrayConnections, db)

			db = &sql.DB{}
		}
	}

	return arrayConnections, nil
}
