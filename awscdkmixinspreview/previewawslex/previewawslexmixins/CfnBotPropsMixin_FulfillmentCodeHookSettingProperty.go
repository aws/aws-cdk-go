package previewawslexmixins


// Determines if a Lambda function should be invoked for a specific intent.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentcodehooksetting.html
//
type CfnBotPropsMixin_FulfillmentCodeHookSettingProperty struct {
	// Indicates whether a Lambda function should be invoked to fulfill a specific intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentcodehooksetting.html#cfn-lex-bot-fulfillmentcodehooksetting-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Provides settings for update messages sent to the user for long-running Lambda fulfillment functions.
	//
	// Fulfillment updates can be used only with streaming conversations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentcodehooksetting.html#cfn-lex-bot-fulfillmentcodehooksetting-fulfillmentupdatesspecification
	//
	FulfillmentUpdatesSpecification interface{} `field:"optional" json:"fulfillmentUpdatesSpecification" yaml:"fulfillmentUpdatesSpecification"`
	// Determines whether the fulfillment code hook is used.
	//
	// When `active` is false, the code hook doesn't run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentcodehooksetting.html#cfn-lex-bot-fulfillmentcodehooksetting-isactive
	//
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// Provides settings for messages sent to the user for after the Lambda fulfillment function completes.
	//
	// Post-fulfillment messages can be sent for both streaming and non-streaming conversations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentcodehooksetting.html#cfn-lex-bot-fulfillmentcodehooksetting-postfulfillmentstatusspecification
	//
	PostFulfillmentStatusSpecification interface{} `field:"optional" json:"postFulfillmentStatusSpecification" yaml:"postFulfillmentStatusSpecification"`
}

