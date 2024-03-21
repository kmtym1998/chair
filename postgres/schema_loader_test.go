package postgres

import (
	"testing"

	"github.com/kmtym1998/chair/testutil"
)

func TestLoadSchema(t *testing.T) {
	container, err := testutil.NewPostgreSQLContainer(
		t,
		"14",
		testutil.TestDBOptions{
			User:     "test",
			Password: "test",
			DBName:   "test",
		},
	)
	if err != nil {
		t.Fatalf("failed to setup postgres container: %v", err)
	}
	defer container.Purge()
	// ddlList := []string{
	// 	"CREATE TABLE numeric_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"smallint_value_nullable SMALLINT," +
	// 		"integer_value_nullable INTEGER," +
	// 		"bigint_value_nullable BIGINT," +
	// 		"decimal_value_nullable DECIMAL," +
	// 		"numeric_value_nullable NUMERIC," +
	// 		"real_value_nullable REAL," +
	// 		"double_precision_value_nullable DOUBLE PRECISION," +
	// 		"smallserial_value_nullable SMALLSERIAL," +
	// 		"serial_value_nullable SERIAL," +
	// 		"bigserial_value_nullable BIGSERIAL," +
	// 		"smallint_value SMALLINT NOT NULL," +
	// 		"integer_value INTEGER NOT NULL," +
	// 		"bigint_value BIGINT NOT NULL," +
	// 		"decimal_value DECIMAL NOT NULL," +
	// 		"numeric_value NUMERIC NOT NULL," +
	// 		"real_value REAL NOT NULL," +
	// 		"double_precision_value DOUBLE PRECISION NOT NULL," +
	// 		"smallserial_value SMALLSERIAL NOT NULL," +
	// 		"serial_value SERIAL NOT NULL," +
	// 		"bigserial_value BIGSERIAL NOT NULL" +
	// 		");",
	// 	"CREATE TABLE money_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"money_value_nullable MONEY," +
	// 		"money_value MONEY NOT NULL" +
	// 		");",
	// 	"CREATE TABLE character_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"character_value_nullable CHAR(1)," +
	// 		"character_varying_value_nullable VARCHAR(255)," +
	// 		"text_value_nullable TEXT," +
	// 		"character_value CHAR(1) NOT NULL," +
	// 		"character_varying_value VARCHAR(255) NOT NULL," +
	// 		"text_value TEXT NOT NULL" +
	// 		");",
	// 	"CREATE TABLE datetime_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"date_value_nullable DATE," +
	// 		"time_value_nullable TIME," +
	// 		"timestamp_value_nullable TIMESTAMP," +
	// 		"timestamptz_value_nullable TIMESTAMPTZ," +
	// 		"interval_value_nullable INTERVAL," +
	// 		"date_value DATE NOT NULL," +
	// 		"time_value TIME NOT NULL," +
	// 		"timestamp_value TIMESTAMP NOT NULL," +
	// 		"timestamptz_value TIMESTAMPTZ NOT NULL," +
	// 		"interval_value INTERVAL NOT NULL" +
	// 		");",
	// 	"CREATE TABLE boolean_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"boolean_value_nullable BOOLEAN," +
	// 		"boolean_value BOOLEAN NOT NULL" +
	// 		");",
	// 	"CREATE TABLE uuid_types (" +
	// 		"id SERIAL PRIMARY KEY," +
	// 		"uuid_value_nullable UUID," +
	// 		"uuid_value UUID NOT NULL" +
	// 		");",
	// }
}
