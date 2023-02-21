package awssesactions


// The type of invocation to use for a Lambda Action.
type LambdaInvocationType string

const (
	// The function will be invoked asynchronously.
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	// The function will be invoked sychronously.
	//
	// Use RequestResponse only when
	// you want to make a mail flow decision, such as whether to stop the receipt
	// rule or the receipt rule set.
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
)

