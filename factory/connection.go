package factory

import (
	"database/sql"
	"Demo-auth/config"
	"Demo-auth/util"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SqlConnector() (*sql.DB, error) {
	if db == nil {

		d, err := sql.Open(config.SQL, util.GetEnv(config.ENV_MYSQL_USER, config.MYSQL_USER) +
			config.KEY_SEPARATOR +
			util.GetEnv(config.ENV_MYSQL_PASSWORD, config.MYSQL_PASSWORD) +
			config.AT + util.GetEnv(config.ENV_DB_HOST, config.DB_HOST) +
			config.FORWARD_SLASH + util.GetEnv(config.ENV_DB_SCHEMA, config.DB_NAME))

		if err != nil {
			return nil, err
		}
		db = d

	}
	return db, nil
}
