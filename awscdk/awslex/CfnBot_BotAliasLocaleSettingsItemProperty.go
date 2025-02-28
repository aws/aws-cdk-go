package awslex


// Specifies locale settings for a single locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsItemProperty := &BotAliasLocaleSettingsItemProperty{
//   	BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CodeHookSpecification: &CodeHookSpecificationProperty{
//   			LambdaCodeHook: &LambdaCodeHookProperty{
//   				CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				LambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   	},
//   	LocaleId: jsii.String("localeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html
//
type CfnBot_BotAliasLocaleSettingsItemProperty struct {
	// Specifies locale settings for a locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html#cfn-lex-bot-botaliaslocalesettingsitem-botaliaslocalesetting
	//
	BotAliasLocaleSetting interface{} `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// Specifies the locale that the settings apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html#cfn-lex-bot-botaliaslocalesettingsitem-localeid
	//
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

