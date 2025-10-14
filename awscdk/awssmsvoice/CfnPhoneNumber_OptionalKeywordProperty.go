package awssmsvoice


// The `OptionalKeyword` configuration.
//
// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionalKeywordProperty := &OptionalKeywordProperty{
//   	Action: jsii.String("action"),
//   	Keyword: jsii.String("keyword"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-optionalkeyword.html
//
type CfnPhoneNumber_OptionalKeywordProperty struct {
	// The action to perform when the keyword is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-optionalkeyword.html#cfn-smsvoice-phonenumber-optionalkeyword-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The new keyword to add.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-optionalkeyword.html#cfn-smsvoice-phonenumber-optionalkeyword-keyword
	//
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// The message associated with the keyword.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-optionalkeyword.html#cfn-smsvoice-phonenumber-optionalkeyword-message
	//
	Message *string `field:"required" json:"message" yaml:"message"`
}

