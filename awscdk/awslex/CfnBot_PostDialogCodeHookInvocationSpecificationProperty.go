package awslex


// Specifies next steps to run after the dialog code hook finishes.
//
// Example:
//
//
type CfnBot_PostDialogCodeHookInvocationSpecificationProperty struct {
	// A list of conditional branches to evaluate after the dialog code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step the bot runs after the dialog code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the code hook fails.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// A list of conditional branches to evaluate after the dialog code hook finishes successfully.
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// Specifics the next step the bot runs after the dialog code hook finishes successfully.
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the code hook succeeds.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// A list of conditional branches to evaluate if the code hook times out.
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// Specifies the next step that the bot runs when the code hook times out.
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond to the user input when the code hook times out.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

