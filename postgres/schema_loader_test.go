package postgres

import (
	"context"
	"testing"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/testutil"
	"github.com/stretchr/testify/assert"
)

var ddlList = []string{
	"CREATE TABLE public.numeric_types (" +
		"id SERIAL PRIMARY KEY," +
		"smallint_value_nullable SMALLINT," +
		"integer_value_nullable INTEGER," +
		"bigint_value_nullable BIGINT," +
		"decimal_value_nullable DECIMAL," +
		"numeric_value_nullable NUMERIC," +
		"real_value_nullable REAL," +
		"double_precision_value_nullable DOUBLE PRECISION," +
		"smallint_value SMALLINT NOT NULL," +
		"integer_value INTEGER NOT NULL," +
		"bigint_value BIGINT NOT NULL," +
		"decimal_value DECIMAL NOT NULL," +
		"numeric_value NUMERIC NOT NULL," +
		"real_value REAL NOT NULL," +
		"double_precision_value DOUBLE PRECISION NOT NULL," +
		"smallserial_value SMALLSERIAL NOT NULL," +
		"serial_value SERIAL NOT NULL," +
		"bigserial_value BIGSERIAL NOT NULL" +
		");",
	"COMMENT ON TABLE public.numeric_types IS 'numeric types';",
	"COMMENT ON COLUMN public.numeric_types.smallint_value_nullable IS 'smallint value nullable';",
	"COMMENT ON COLUMN public.numeric_types.integer_value_nullable IS 'integer value nullable';",
	"CREATE TABLE public.money_types (" +
		"id SERIAL PRIMARY KEY," +
		"money_value_nullable MONEY," +
		"money_value MONEY NOT NULL" +
		");",
	"CREATE TABLE public.character_types (" +
		"id SERIAL PRIMARY KEY," +
		"character_value_nullable CHAR(1)," +
		"character_varying_value_nullable VARCHAR(255)," +
		"text_value_nullable TEXT," +
		"character_value CHAR(1) NOT NULL," +
		"character_varying_value VARCHAR(255) NOT NULL," +
		"text_value TEXT NOT NULL" +
		");",
	"CREATE TABLE public.datetime_types (" +
		"id SERIAL PRIMARY KEY," +
		"date_value_nullable DATE," +
		"time_value_nullable TIME," +
		"timestamp_value_nullable TIMESTAMP," +
		"timestamptz_value_nullable TIMESTAMPTZ," +
		"interval_value_nullable INTERVAL," +
		"date_value DATE NOT NULL," +
		"time_value TIME NOT NULL," +
		"timestamp_value TIMESTAMP NOT NULL," +
		"timestamptz_value TIMESTAMPTZ NOT NULL," +
		"interval_value INTERVAL NOT NULL" +
		");",
	"CREATE TABLE public.boolean_types (" +
		"id SERIAL PRIMARY KEY," +
		"boolean_value_nullable BOOLEAN," +
		"boolean_value BOOLEAN NOT NULL" +
		");",
	"CREATE TABLE public.uuid_types (" +
		"id SERIAL PRIMARY KEY," +
		"uuid_value_nullable UUID," +
		"uuid_value UUID NOT NULL" +
		");",
}

func TestLoadTableSchemas(t *testing.T) {
	for ver, name := range map[string]string{
		// List of supported versions: https://endoflife.date/postgresql
		"12": "PostgreSQL v12", // EOL: 2024-11
		"13": "PostgreSQL v13", // EOL: 2025-11
		"14": "PostgreSQL v14", // EOL: 2026-11
		"15": "PostgreSQL v15", // EOL: 2027-11
		"16": "PostgreSQL v16", // EOL: 2028-11
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			container, err := testutil.NewPostgreSQLContainer(
				t,
				ver,
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

			db, err := container.ConnectDB()
			if err != nil {
				t.Fatalf("failed to connect to postgres container: %v", err)
			}
			defer db.Close()

			for _, ddl := range ddlList {
				_, err := db.Exec(ddl)
				if err != nil {
					t.Fatalf("failed to create table: %v", err)
				}
			}

			ldr := NewSchemaLoader(db, "public")
			actual, err := ldr.LoadTableSchemas(context.Background())
			if err != nil {
				t.Fatalf("failed to load table schemas: %v", err)
			}

			expected := []generator.Table{
				{
					Name: "character_types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "character_value_nullable", Type: "character", IsNullable: true, OrderAsc: 2},
						{Name: "character_varying_value_nullable", Type: "character varying", IsNullable: true, OrderAsc: 3},
						{Name: "text_value_nullable", Type: "text", IsNullable: true, OrderAsc: 4},
						{Name: "character_value", Type: "character", IsNullable: false, OrderAsc: 5},
						{Name: "character_varying_value", Type: "character varying", IsNullable: false, OrderAsc: 6},
						{Name: "text_value", Type: "text", IsNullable: false, OrderAsc: 7},
					},
				},
				{
					Name:    "numeric_types",
					Comment: "numeric types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "smallint_value_nullable", Type: "smallint", IsNullable: true, OrderAsc: 2, Comment: "smallint value nullable"},
						{Name: "integer_value_nullable", Type: "integer", IsNullable: true, OrderAsc: 3, Comment: "integer value nullable"},
						{Name: "bigint_value_nullable", Type: "bigint", IsNullable: true, OrderAsc: 4},
						{Name: "decimal_value_nullable", Type: "numeric", IsNullable: true, OrderAsc: 5},
						{Name: "numeric_value_nullable", Type: "numeric", IsNullable: true, OrderAsc: 6},
						{Name: "real_value_nullable", Type: "real", IsNullable: true, OrderAsc: 7},
						{Name: "double_precision_value_nullable", Type: "double precision", IsNullable: true, OrderAsc: 8},
						{Name: "smallint_value", Type: "smallint", IsNullable: false, OrderAsc: 9},
						{Name: "integer_value", Type: "integer", IsNullable: false, OrderAsc: 10},
						{Name: "bigint_value", Type: "bigint", IsNullable: false, OrderAsc: 11},
						{Name: "decimal_value", Type: "numeric", IsNullable: false, OrderAsc: 12},
						{Name: "numeric_value", Type: "numeric", IsNullable: false, OrderAsc: 13},
						{Name: "real_value", Type: "real", IsNullable: false, OrderAsc: 14},
						{Name: "double_precision_value", Type: "double precision", IsNullable: false, OrderAsc: 15},
						{Name: "smallserial_value", Type: "smallint", IsNullable: false, OrderAsc: 16},
						{Name: "serial_value", Type: "integer", IsNullable: false, OrderAsc: 17},
						{Name: "bigserial_value", Type: "bigint", IsNullable: false, OrderAsc: 18},
					},
				},
				{
					Name: "datetime_types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "date_value_nullable", Type: "date", IsNullable: true, OrderAsc: 2},
						{Name: "time_value_nullable", Type: "time without time zone", IsNullable: true, OrderAsc: 3},
						{Name: "timestamp_value_nullable", Type: "timestamp without time zone", IsNullable: true, OrderAsc: 4},
						{Name: "timestamptz_value_nullable", Type: "timestamp with time zone", IsNullable: true, OrderAsc: 5},
						{Name: "interval_value_nullable", Type: "interval", IsNullable: true, OrderAsc: 6},
						{Name: "date_value", Type: "date", IsNullable: false, OrderAsc: 7},
						{Name: "time_value", Type: "time without time zone", IsNullable: false, OrderAsc: 8},
						{Name: "timestamp_value", Type: "timestamp without time zone", IsNullable: false, OrderAsc: 9},
						{Name: "timestamptz_value", Type: "timestamp with time zone", IsNullable: false, OrderAsc: 10},
						{Name: "interval_value", Type: "interval", IsNullable: false, OrderAsc: 11},
					},
				},
				{
					Name: "uuid_types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "uuid_value_nullable", Type: "uuid", IsNullable: true, OrderAsc: 2},
						{Name: "uuid_value", Type: "uuid", IsNullable: false, OrderAsc: 3},
					},
				},
				{
					Name: "money_types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "money_value_nullable", Type: "money", IsNullable: true, OrderAsc: 2},
						{Name: "money_value", Type: "money", IsNullable: false, OrderAsc: 3},
					},
				},
				{
					Name: "boolean_types",
					Columns: []generator.Column{
						{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
						{Name: "boolean_value_nullable", Type: "boolean", IsNullable: true, OrderAsc: 2},
						{Name: "boolean_value", Type: "boolean", IsNullable: false, OrderAsc: 3},
					},
				},
			}

			t.Run("assert table length", func(t *testing.T) {
				assert.Len(t, actual, 6)
			})
			t.Run("assert table column length", func(t *testing.T) {
				assertTableColumnLength := func(t *testing.T, table string, expected int) {
					for _, tbl := range actual {
						if tbl.Name == table {
							assert.Len(t, tbl.Columns, expected)
							return
						}
					}

					t.Fatalf("table not found: %s", table)
				}

				assertTableColumnLength(t, "character_types", 7)
				assertTableColumnLength(t, "numeric_types", 18)
				assertTableColumnLength(t, "datetime_types", 11)
				assertTableColumnLength(t, "uuid_types", 3)
				assertTableColumnLength(t, "money_types", 3)
				assertTableColumnLength(t, "boolean_types", 3)
			})

			t.Run("assert table schema content", func(t *testing.T) {
				for _, exp := range expected {
					for _, act := range actual {
						if exp.Name == act.Name && assert.Equal(t, exp, act) {
							continue
						}
					}
				}
			})
		})
	}
}
