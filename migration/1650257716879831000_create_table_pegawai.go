package migration

import (
	migrator "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/migration"
)

func init() {

	MigrationScript[1650257716879831000] = &migrator.Script{
		Up: `CREATE TABLE IF NOT EXISTS pegawai(
							id SERIAL PRIMARY KEY,
							nama VARCHAR,
							alamat TEXT,
							telepon VARCHAR,
							created_date timestamp with time zone NOT NULL DEFAULT now(),
							created_by VARCHAR(255) NOT NULL,
							updated_date timestamp with time zone NOT NULL DEFAULT now(),
							updated_by VARCHAR(255) NOT NULL,
							version INT NOT NULL,
							is_deleted BOOLEAN NOT NULL
);`,
		Down:             `DROP TABLE IF EXISTS pegawai;`,
		UsingTransaction: true,
	}

}
