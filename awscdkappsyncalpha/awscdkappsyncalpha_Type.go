// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Enum containing the Types that can be used to define ObjectTypes.
// Experimental.
type Type string

const (
	// `ID` scalar type is a unique identifier. `ID` type is serialized similar to `String`.
	//
	// Often used as a key for a cache and not intended to be human-readable.
	// Experimental.
	Type_ID Type = "ID"
	// `String` scalar type is a free-form human-readable text.
	// Experimental.
	Type_STRING Type = "STRING"
	// `Int` scalar type is a signed non-fractional numerical value.
	// Experimental.
	Type_INT Type = "INT"
	// `Float` scalar type is a signed double-precision fractional value.
	// Experimental.
	Type_FLOAT Type = "FLOAT"
	// `Boolean` scalar type is a boolean value: true or false.
	// Experimental.
	Type_BOOLEAN Type = "BOOLEAN"
	// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
	//
	// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
	// See: https://en.wikipedia.org/wiki/ISO_8601#Calendar_dates
	//
	// Experimental.
	Type_AWS_DATE Type = "AWS_DATE"
	// `AWSTime` scalar type represents a valid extended `ISO 8601 Time` string.
	//
	// In other words, accepts date strings in the form of `hh:mm:ss.sss`. It accepts time zone offsets.
	// See: https://en.wikipedia.org/wiki/ISO_8601#Times
	//
	// Experimental.
	Type_AWS_TIME Type = "AWS_TIME"
	// `AWSDateTime` scalar type represents a valid extended `ISO 8601 DateTime` string.
	//
	// In other words, accepts date strings in the form of `YYYY-MM-DDThh:mm:ss.sssZ`. It accepts time zone offsets.
	// See: https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations
	//
	// Experimental.
	Type_AWS_DATE_TIME Type = "AWS_DATE_TIME"
	// `AWSTimestamp` scalar type represents the number of seconds since `1970-01-01T00:00Z`.
	//
	// Timestamps are serialized and deserialized as numbers.
	// Experimental.
	Type_AWS_TIMESTAMP Type = "AWS_TIMESTAMP"
	// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
	// Experimental.
	Type_AWS_EMAIL Type = "AWS_EMAIL"
	// `AWSJson` scalar type represents a JSON string.
	// Experimental.
	Type_AWS_JSON Type = "AWS_JSON"
	// `AWSURL` scalar type represetns a valid URL string.
	//
	// URLs wihtout schemes or contain double slashes are considered invalid.
	// Experimental.
	Type_AWS_URL Type = "AWS_URL"
	// `AWSPhone` scalar type represents a valid phone number. Phone numbers maybe be whitespace delimited or hyphenated.
	//
	// The number can specify a country code at the beginning, but is not required for US phone numbers.
	// Experimental.
	Type_AWS_PHONE Type = "AWS_PHONE"
	// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
	// Experimental.
	Type_AWS_IP_ADDRESS Type = "AWS_IP_ADDRESS"
	// Type used for Intermediate Types (i.e. an interface or an object type).
	// Experimental.
	Type_INTERMEDIATE Type = "INTERMEDIATE"
)

