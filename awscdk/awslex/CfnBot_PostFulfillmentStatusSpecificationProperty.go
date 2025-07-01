package awslex


// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [](https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete)
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html
//
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// A list of conditional branches to evaluate after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-failureconditional
	//
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step the bot runs after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-failurenextstep
	//
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-failureresponse
	//
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// A list of conditional branches to evaluate after the fulfillment code hook finishes successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-successconditional
	//
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// Specifies the next step in the conversation that Amazon Lex invokes when the fulfillment code hook completes successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-successnextstep
	//
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-successresponse
	//
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// A list of conditional branches to evaluate if the fulfillment code hook times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-timeoutconditional
	//
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// Specifies the next step that the bot runs when the fulfillment code hook times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-timeoutnextstep
	//
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-postfulfillmentstatusspecification.html#cfn-lex-bot-postfulfillmentstatusspecification-timeoutresponse
	//
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

