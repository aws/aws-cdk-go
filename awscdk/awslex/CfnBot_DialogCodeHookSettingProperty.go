package awslex


// Settings that determine the Lambda function that Amazon Lex uses for processing user responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogCodeHookSettingProperty := &DialogCodeHookSettingProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehooksetting.html
//
type CfnBot_DialogCodeHookSettingProperty struct {
	// Enables the dialog code hook so that it processes user requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehooksetting.html#cfn-lex-bot-dialogcodehooksetting-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

