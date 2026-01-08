package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   episodicOverrideReflectionConfigurationInputProperty := &EpisodicOverrideReflectionConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html
//
type CfnMemory_EpisodicOverrideReflectionConfigurationInputProperty struct {
	// Text prompt for model instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

