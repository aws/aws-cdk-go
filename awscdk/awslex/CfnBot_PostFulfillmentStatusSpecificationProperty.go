package awslex


// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [](https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete)
//
// Example:
//
//
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// A list of conditional branches to evaluate after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step the bot runs after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// A list of conditional branches to evaluate after the fulfillment code hook finishes successfully.
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// Specifies the next step in the conversation that Amazon Lex invokes when the fulfillment code hook completes successfully.
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// A list of conditional branches to evaluate if the fulfillment code hook times out.
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// Specifies the next step that the bot runs when the fulfillment code hook times out.
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

