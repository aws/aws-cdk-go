package mixinsawsbedrock


// Contains information about a connection between two nodes in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flowConnectionProperty := &FlowConnectionProperty{
//   	Configuration: &FlowConnectionConfigurationProperty{
//   		Conditional: &FlowConditionalConnectionConfigurationProperty{
//   			Condition: jsii.String("condition"),
//   		},
//   		Data: &FlowDataConnectionConfigurationProperty{
//   			SourceOutput: jsii.String("sourceOutput"),
//   			TargetInput: jsii.String("targetInput"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html
//
type CfnFlowPropsMixin_FlowConnectionProperty struct {
	// The configuration of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html#cfn-bedrock-flow-flowconnection-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A name for the connection that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html#cfn-bedrock-flow-flowconnection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The node that the connection starts at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html#cfn-bedrock-flow-flowconnection-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The node that the connection ends at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html#cfn-bedrock-flow-flowconnection-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Whether the source node that the connection begins from is a condition node ( `Conditional` ) or not ( `Data` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowconnection.html#cfn-bedrock-flow-flowconnection-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

