package orm

const (
	// Date / Time
	timestampType   = "timestamp"           // Timestamp without a time zone
	timestamptzType = "timestamptz"         // Timestamp with a time zone
	dateType        = "date"                // Date
	timeType        = "time"                // Time without a time zone
	timetzType      = "time with time zone" // Time with a time zone
	intervalType    = "interval"            // Time Interval

	// Network Addresses
	inetType    = "inet"    // IPv4 or IPv6 hosts and networks
	cidrType    = "cidr"    // IPv4 or IPv6 networks
	macaddrType = "macaddr" // MAC adresses

	// Boolean
	booleanType = "boolean"

	// Numeric Types

	// Floating Point Types
	realType            = "real"             // 4 byte floating point (6 digit precision)
	doublePrecisionType = "double precision" // 8 byte floating point (15 digit precision)

	// Integer Types
	smallintType = "smallint" // 2 byte integer
	integerType  = "integer"  // 4 byte integer
	bigintType   = "bigint"   // 8 byte integer

	// Serial Types
	smallserialType = "smallserial" // 2 byte autoincrementing integer
	serialType      = "serial"      // 4 byte autoincrementing integer
	bigserialType   = "bigserial"   // 8 byte autoincrementing integer

	// Character Types
	varcharType = "varchar" // variable length string with limit
	charType    = "char"    // fixed length string (blank padded)
	textType    = "text"    // variable length string without limit

	// JSON Types
	jsonType  = "json"  // text representation of json data
	jsonbType = "jsonb" // binary representation of json data

	// Binary Data Types
	byteaType = "bytea" // binary string
)
