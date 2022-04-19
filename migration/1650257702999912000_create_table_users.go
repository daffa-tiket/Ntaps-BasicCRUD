package migration

import (
	migrator "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/migration"
)

func init() {

	MigrationScript[1650257702999912000] = &migrator.Script{
		Up: `CREATE TABLE IF NOT EXISTS users(
							id SERIAL PRIMARY KEY,
							username VARCHAR,
							password VARCHAR);`,
		Down:             `DROP TABLE IF EXISTS users;`,
		UsingTransaction: true,
	}

}
