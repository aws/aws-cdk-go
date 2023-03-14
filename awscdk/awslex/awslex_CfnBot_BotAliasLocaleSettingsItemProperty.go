package awslex


// Specifies locale settings for a single locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsItemProperty := &botAliasLocaleSettingsItemProperty{
//   	botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		codeHookSpecification: &codeHookSpecificationProperty{
//   			lambdaCodeHook: &lambdaCodeHookProperty{
//   				codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				lambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   	},
//   	localeId: jsii.String("localeId"),
//   }
//
type CfnBot_BotAliasLocaleSettingsItemProperty struct {
	// Specifies locale settings for a locale.
	BotAliasLocaleSetting interface{} `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// Specifies the locale that the settings apply to.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

