package mixinsawslex


// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   botAliasLocaleSettingsProperty := &BotAliasLocaleSettingsProperty{
//   	CodeHookSpecification: &CodeHookSpecificationProperty{
//   		LambdaCodeHook: &LambdaCodeHookProperty{
//   			CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettings.html
//
type CfnBotPropsMixin_BotAliasLocaleSettingsProperty struct {
	// Specifies the Lambda function that should be used in the locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettings.html#cfn-lex-bot-botaliaslocalesettings-codehookspecification
	//
	CodeHookSpecification interface{} `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
	// Determines whether the locale is enabled for the bot.
	//
	// If the value is `false` , the locale isn't available for use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettings.html#cfn-lex-bot-botaliaslocalesettings-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

