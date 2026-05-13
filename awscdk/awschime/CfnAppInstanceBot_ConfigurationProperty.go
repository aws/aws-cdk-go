package awschime


// A structure that contains configuration data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationProperty := &ConfigurationProperty{
//   	Lex: &LexConfigurationProperty{
//   		LexBotAliasArn: jsii.String("lexBotAliasArn"),
//   		LocaleId: jsii.String("localeId"),
//
//   		// the properties below are optional
//   		InvokedBy: &InvokedByProperty{
//   			StandardMessages: jsii.String("standardMessages"),
//   			TargetedMessages: jsii.String("targetedMessages"),
//   		},
//   		RespondsTo: jsii.String("respondsTo"),
//   		WelcomeIntent: jsii.String("welcomeIntent"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-configuration.html
//
type CfnAppInstanceBot_ConfigurationProperty struct {
	// The configuration for an Amazon Lex V2 bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-configuration.html#cfn-chime-appinstancebot-configuration-lex
	//
	Lex interface{} `field:"required" json:"lex" yaml:"lex"`
}

