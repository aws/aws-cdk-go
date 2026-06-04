package awsbedrockagentcorealpha


// Custom claim match operator.
//
// Shared by Runtime and Gateway custom claim implementations.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CustomClaimOperator string

const (
	// Equals operator - used for STRING type claims.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomClaimOperator_EQUALS CustomClaimOperator = "EQUALS"
	// Contains operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains a specific string value.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomClaimOperator_CONTAINS CustomClaimOperator = "CONTAINS"
	// ContainsAny operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains any of the provided string values.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomClaimOperator_CONTAINS_ANY CustomClaimOperator = "CONTAINS_ANY"
)

