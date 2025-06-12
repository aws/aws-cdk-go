package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qInConnectAssistantConfigurationProperty := &QInConnectAssistantConfigurationProperty{
//   	AssistantArn: jsii.String("assistantArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qinconnectassistantconfiguration.html
//
type CfnBot_QInConnectAssistantConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qinconnectassistantconfiguration.html#cfn-lex-bot-qinconnectassistantconfiguration-assistantarn
	//
	AssistantArn *string `field:"required" json:"assistantArn" yaml:"assistantArn"`
}

