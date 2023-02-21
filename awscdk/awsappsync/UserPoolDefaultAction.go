package awsappsync


// enum with all possible values for Cognito user-pool default actions.
type UserPoolDefaultAction string

const (
	// ALLOW access to API.
	UserPoolDefaultAction_ALLOW UserPoolDefaultAction = "ALLOW"
	// DENY access to API.
	UserPoolDefaultAction_DENY UserPoolDefaultAction = "DENY"
)

