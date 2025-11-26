package previewawsbedrockmixins


// The configuration of a connection between a condition node and another node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flowConditionalConnectionConfigurationProperty := &FlowConditionalConnectionConfigurationProperty{
//   	Condition: jsii.String("condition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconditionalconnectionconfiguration.html
//
type CfnFlowVersionPropsMixin_FlowConditionalConnectionConfigurationProperty struct {
	// The condition that triggers this connection.
	//
	// For more information about how to write conditions, see the *Condition* node type in the [Node types](https://docs.aws.amazon.com/bedrock/latest/userguide/node-types.html) topic in the Amazon Bedrock User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconditionalconnectionconfiguration.html#cfn-bedrock-flowversion-flowconditionalconnectionconfiguration-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
}

