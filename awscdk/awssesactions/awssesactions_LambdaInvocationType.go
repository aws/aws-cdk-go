package awssesactions


// The type of invocation to use for a Lambda Action.
// Experimental.
type LambdaInvocationType string

const (
	// The function will be invoked asynchronously.
	// Experimental.
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	// The function will be invoked sychronously.
	//
	// Use RequestResponse only when
	// you want to make a mail flow decision, such as whether to stop the receipt
	// rule or the receipt rule set.
	// Experimental.
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
)

