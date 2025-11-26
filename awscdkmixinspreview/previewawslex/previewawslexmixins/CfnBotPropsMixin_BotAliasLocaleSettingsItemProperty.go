package previewawslexmixins


// Specifies locale settings for a single locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   botAliasLocaleSettingsItemProperty := &BotAliasLocaleSettingsItemProperty{
//   	BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   		CodeHookSpecification: &CodeHookSpecificationProperty{
//   			LambdaCodeHook: &LambdaCodeHookProperty{
//   				CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				LambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   	LocaleId: jsii.String("localeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html
//
type CfnBotPropsMixin_BotAliasLocaleSettingsItemProperty struct {
	// Specifies locale settings for a locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html#cfn-lex-bot-botaliaslocalesettingsitem-botaliaslocalesetting
	//
	BotAliasLocaleSetting interface{} `field:"optional" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// Specifies the locale that the settings apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botaliaslocalesettingsitem.html#cfn-lex-bot-botaliaslocalesettingsitem-localeid
	//
	LocaleId *string `field:"optional" json:"localeId" yaml:"localeId"`
}

