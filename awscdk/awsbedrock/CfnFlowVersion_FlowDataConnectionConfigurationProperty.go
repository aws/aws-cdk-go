package awsbedrock


// The configuration of a connection originating from a node that isn't a Condition node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowDataConnectionConfigurationProperty := &FlowDataConnectionConfigurationProperty{
//   	SourceOutput: jsii.String("sourceOutput"),
//   	TargetInput: jsii.String("targetInput"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdataconnectionconfiguration.html
//
type CfnFlowVersion_FlowDataConnectionConfigurationProperty struct {
	// The name of the output in the source node that the connection begins from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdataconnectionconfiguration.html#cfn-bedrock-flowversion-flowdataconnectionconfiguration-sourceoutput
	//
	SourceOutput *string `field:"required" json:"sourceOutput" yaml:"sourceOutput"`
	// The name of the input in the target node that the connection ends at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdataconnectionconfiguration.html#cfn-bedrock-flowversion-flowdataconnectionconfiguration-targetinput
	//
	TargetInput *string `field:"required" json:"targetInput" yaml:"targetInput"`
}

