package awsappconfig


// The validator type.
type ValidatorType string

const (
	// JSON Scema validator.
	ValidatorType_JSON_SCHEMA ValidatorType = "JSON_SCHEMA"
	// Validate using a Lambda function.
	ValidatorType_LAMBDA ValidatorType = "LAMBDA"
)

