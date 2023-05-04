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
type CfnBot_DialogCodeHookSettingProperty struct {
	// Enables the dialog code hook so that it processes user requests.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

