package awscdkappconfigalpha


// The validator type.
// Experimental.
type ValidatorType string

const (
	// JSON Scema validator.
	// Experimental.
	ValidatorType_JSON_SCHEMA ValidatorType = "JSON_SCHEMA"
	// Validate using a Lambda function.
	// Experimental.
	ValidatorType_LAMBDA ValidatorType = "LAMBDA"
)

