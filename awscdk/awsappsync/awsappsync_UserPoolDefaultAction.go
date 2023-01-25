package awsappsync


// enum with all possible values for Cognito user-pool default actions.
// Experimental.
type UserPoolDefaultAction string

const (
	// ALLOW access to API.
	// Experimental.
	UserPoolDefaultAction_ALLOW UserPoolDefaultAction = "ALLOW"
	// DENY access to API.
	// Experimental.
	UserPoolDefaultAction_DENY UserPoolDefaultAction = "DENY"
)

