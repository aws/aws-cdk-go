package awschime


// The configuration for an Amazon Lex V2 bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lexConfigurationProperty := &LexConfigurationProperty{
//   	LexBotAliasArn: jsii.String("lexBotAliasArn"),
//   	LocaleId: jsii.String("localeId"),
//
//   	// the properties below are optional
//   	InvokedBy: &InvokedByProperty{
//   		StandardMessages: jsii.String("standardMessages"),
//   		TargetedMessages: jsii.String("targetedMessages"),
//   	},
//   	RespondsTo: jsii.String("respondsTo"),
//   	WelcomeIntent: jsii.String("welcomeIntent"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html
//
type CfnAppInstanceBot_LexConfigurationProperty struct {
	// The ARN of the Amazon Lex V2 bot's alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html#cfn-chime-appinstancebot-lexconfiguration-lexbotaliasarn
	//
	LexBotAliasArn *string `field:"required" json:"lexBotAliasArn" yaml:"lexBotAliasArn"`
	// Identifies the Amazon Lex V2 bot's language and locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html#cfn-chime-appinstancebot-lexconfiguration-localeid
	//
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Specifies the type of message that triggers a bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html#cfn-chime-appinstancebot-lexconfiguration-invokedby
	//
	InvokedBy interface{} `field:"optional" json:"invokedBy" yaml:"invokedBy"`
	// Determines whether the Amazon Lex V2 bot responds to all standard messages.
	//
	// Control messages are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html#cfn-chime-appinstancebot-lexconfiguration-respondsto
	//
	RespondsTo *string `field:"optional" json:"respondsTo" yaml:"respondsTo"`
	// The name of the welcome intent configured in the Amazon Lex V2 bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-lexconfiguration.html#cfn-chime-appinstancebot-lexconfiguration-welcomeintent
	//
	WelcomeIntent *string `field:"optional" json:"welcomeIntent" yaml:"welcomeIntent"`
}

