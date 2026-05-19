package awsbedrockagentcorealpha


// Custom claim match operator.
//
// Shared by Runtime and Gateway custom claim implementations.
// Experimental.
type CustomClaimOperator string

const (
	// Equals operator - used for STRING type claims.
	// Experimental.
	CustomClaimOperator_EQUALS CustomClaimOperator = "EQUALS"
	// Contains operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains a specific string value.
	// Experimental.
	CustomClaimOperator_CONTAINS CustomClaimOperator = "CONTAINS"
	// ContainsAny operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains any of the provided string values.
	// Experimental.
	CustomClaimOperator_CONTAINS_ANY CustomClaimOperator = "CONTAINS_ANY"
)

