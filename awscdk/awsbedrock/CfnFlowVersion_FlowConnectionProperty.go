package awsbedrock


// Contains information about a connection between two nodes in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowConnectionProperty := &FlowConnectionProperty{
//   	Name: jsii.String("name"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Configuration: &FlowConnectionConfigurationProperty{
//   		Conditional: &FlowConditionalConnectionConfigurationProperty{
//   			Condition: jsii.String("condition"),
//   		},
//   		Data: &FlowDataConnectionConfigurationProperty{
//   			SourceOutput: jsii.String("sourceOutput"),
//   			TargetInput: jsii.String("targetInput"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html
//
type CfnFlowVersion_FlowConnectionProperty struct {
	// A name for the connection that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html#cfn-bedrock-flowversion-flowconnection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The node that the connection starts at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html#cfn-bedrock-flowversion-flowconnection-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The node that the connection ends at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html#cfn-bedrock-flowversion-flowconnection-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// Whether the source node that the connection begins from is a condition node ( `Conditional` ) or not ( `Data` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html#cfn-bedrock-flowversion-flowconnection-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowconnection.html#cfn-bedrock-flowversion-flowconnection-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
}

