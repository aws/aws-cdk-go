package awsbedrock


// The configuration of the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowConnectionConfigurationProperty := &FlowConnectionConfigurationProperty{
//   	Conditional: &FlowConditionalConnectionConfigurationProperty{
//   		Condition: jsii.String("condition"),
//   	},
//   	Data: &FlowDataConnectionConfigurationProperty{
//   		SourceOutput: jsii.String("sourceOutput"),
//   		TargetInput: jsii.String("targetInput"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnectionconfiguration.html
//
type CfnFlowVersion_FlowConnectionConfigurationProperty struct {
	// The configuration of a connection originating from a Condition node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnectionconfiguration.html#cfn-bedrock-flowversion-flowconnectionconfiguration-conditional
	//
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// The configuration of a connection originating from a node that isn't a Condition node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnectionconfiguration.html#cfn-bedrock-flowversion-flowconnectionconfiguration-data
	//
	Data interface{} `field:"optional" json:"data" yaml:"data"`
}

