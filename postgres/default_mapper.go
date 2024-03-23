package postgres

import (
	"github.com/kmtym1998/chair/generator/config"
)

// https://www.postgresql.org/docs/current/datatype.html
func DefaultMappers() []config.TypeMapping {
	// https://www.postgresql.org/docs/current/datatype-numeric.html
	numericTypes := []config.TypeMapping{
		// non-nullable numeric types
		{
			DBType:     "smallint",
			GoType:     "int",
			IsNullable: false,
		},
		{
			DBType:     "integer",
			GoType:     "int",
			IsNullable: false,
		},
		{
			DBType:     "bigint",
			GoType:     "int64",
			IsNullable: false,
		},
		{
			DBType:     "decimal",
			GoType:     "float64",
			IsNullable: false,
		},
		{
			DBType:     "numeric",
			GoType:     "float64",
			IsNullable: false,
		},
		{
			DBType:     "real",
			GoType:     "float32",
			IsNullable: false,
		},
		{
			DBType:     "double precision",
			GoType:     "float64",
			IsNullable: false,
		},
		{
			DBType:     "smallserial",
			GoType:     "int",
			IsNullable: false,
		},
		{
			DBType:     "serial",
			GoType:     "int",
			IsNullable: false,
		},
		{
			DBType:     "bigserial",
			GoType:     "int64",
			IsNullable: false,
		},
		// nullable numeric types
		{
			DBType:     "smallint",
			GoType:     "NullInt32",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "integer",
			GoType:     "NullInt32",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "bigint",
			GoType:     "NullInt64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "decimal",
			GoType:     "NullFloat64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "numeric",
			GoType:     "NullFloat64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "real",
			GoType:     "NullFloat64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "double precision",
			GoType:     "NullFloat64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	// https://www.postgresql.org/docs/current/datatype-money.html
	moneyTypes := []config.TypeMapping{
		// non-nullable
		{
			DBType:     "money",
			GoType:     "float64",
			IsNullable: false,
		},
		// nullable
		{
			DBType:     "money",
			GoType:     "NullFloat64",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	// https://www.postgresql.org/docs/current/datatype-character.html
	characterTypes := []config.TypeMapping{
		// non-nullable
		{
			DBType:     "character",
			GoType:     "string",
			IsNullable: false,
		},
		{
			DBType:     "char",
			GoType:     "string",
			IsNullable: false,
		},
		{
			DBType:     "character varying",
			GoType:     "string",
			IsNullable: false,
		},
		{
			DBType:     "varchar",
			GoType:     "string",
			IsNullable: false,
		},
		{
			DBType:     "bpchar",
			GoType:     "string",
			IsNullable: false,
		},
		{
			DBType:     "text",
			GoType:     "string",
			IsNullable: false,
		},
		// nullable
		{
			DBType:     "character",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "char",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "character varying",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "varchar",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "bpchar",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "text",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	// https://www.postgresql.org/docs/current/datatype-datetime.html
	datetimeTypes := []config.TypeMapping{
		// non-nullable
		{
			DBType:     "timestamp",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "timestamp with time zone",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "timestamp without time zone",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "date",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "time",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "time with time zone",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		{
			DBType:     "time without time zone",
			GoType:     "Time",
			GoPkg:      "time",
			IsNullable: false,
		},
		// nullable
		{
			DBType:     "timestamp",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "timestamp with time zone",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "timestamp without time zone",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "date",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "time",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "time with time zone",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
		{
			DBType:     "time without time zone",
			GoType:     "NullTime",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	// https://www.postgresql.org/docs/current/datatype-boolean.html
	booleanTypes := []config.TypeMapping{
		// non-nullable
		{
			DBType:     "boolean",
			GoType:     "bool",
			IsNullable: false,
		},
		// nullable
		{
			DBType:     "boolean",
			GoType:     "NullBool",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	// https://www.postgresql.org/docs/current/datatype-uuid.html
	uuidTypes := []config.TypeMapping{
		// non-nullable
		{
			DBType:     "uuid",
			GoType:     "string",
			IsNullable: false,
		},
		// nullable
		{
			DBType:     "uuid",
			GoType:     "NullString",
			GoPkg:      "database/sql",
			IsNullable: true,
		},
	}

	merge := func(arrList ...[]config.TypeMapping) []config.TypeMapping {
		var itemsCount int
		for _, arr := range arrList {
			itemsCount += len(arr)
		}

		merged := make([]config.TypeMapping, 0, itemsCount)
		for _, arr := range arrList {
			merged = append(merged, arr...)
		}

		return merged
	}

	return merge(
		numericTypes,
		moneyTypes,
		characterTypes,
		datetimeTypes,
		booleanTypes,
		uuidTypes,
	)
}
