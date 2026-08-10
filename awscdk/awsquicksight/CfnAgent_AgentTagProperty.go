package awsquicksight


// A key-value pair to associate with the agent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentTagProperty := &AgentTagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-agenttag.html
//
type CfnAgent_AgentTagProperty struct {
	// The key name of the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-agenttag.html#cfn-quicksight-agent-agenttag-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-agenttag.html#cfn-quicksight-agent-agenttag-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

