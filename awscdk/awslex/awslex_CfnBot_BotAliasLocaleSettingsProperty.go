package awslex


// Specifies settings that are unique to a locale.
//
// For example, you can use a different Lambda function for each locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsProperty := &botAliasLocaleSettingsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	codeHookSpecification: &codeHookSpecificationProperty{
//   		lambdaCodeHook: &lambdaCodeHookProperty{
//   			codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   			lambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   }
//
type CfnBot_BotAliasLocaleSettingsProperty struct {
	// Specifies whether the locale is enabled for the bot.
	//
	// If the value is false, the locale isn't available for use.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the Lambda function to use in this locale.
	CodeHookSpecification interface{} `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

