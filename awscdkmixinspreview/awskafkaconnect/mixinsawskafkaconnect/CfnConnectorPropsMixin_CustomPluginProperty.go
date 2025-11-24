package mixinsawskafkaconnect


// A plugin is an AWS resource that contains the code that defines a connector's logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customPluginProperty := &CustomPluginProperty{
//   	CustomPluginArn: jsii.String("customPluginArn"),
//   	Revision: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-customplugin.html
//
type CfnConnectorPropsMixin_CustomPluginProperty struct {
	// The Amazon Resource Name (ARN) of the custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-customplugin.html#cfn-kafkaconnect-connector-customplugin-custompluginarn
	//
	CustomPluginArn *string `field:"optional" json:"customPluginArn" yaml:"customPluginArn"`
	// The revision of the custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-customplugin.html#cfn-kafkaconnect-connector-customplugin-revision
	//
	Revision *float64 `field:"optional" json:"revision" yaml:"revision"`
}

