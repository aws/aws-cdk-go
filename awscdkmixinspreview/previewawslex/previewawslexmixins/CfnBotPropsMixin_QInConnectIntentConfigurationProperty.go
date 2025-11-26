package previewawslexmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   qInConnectIntentConfigurationProperty := &QInConnectIntentConfigurationProperty{
//   	QInConnectAssistantConfiguration: &QInConnectAssistantConfigurationProperty{
//   		AssistantArn: jsii.String("assistantArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qinconnectintentconfiguration.html
//
type CfnBotPropsMixin_QInConnectIntentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qinconnectintentconfiguration.html#cfn-lex-bot-qinconnectintentconfiguration-qinconnectassistantconfiguration
	//
	QInConnectAssistantConfiguration interface{} `field:"optional" json:"qInConnectAssistantConfiguration" yaml:"qInConnectAssistantConfiguration"`
}

