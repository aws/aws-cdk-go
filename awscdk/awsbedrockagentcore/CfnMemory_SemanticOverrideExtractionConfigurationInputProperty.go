package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   semanticOverrideExtractionConfigurationInputProperty := &SemanticOverrideExtractionConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput.html
//
type CfnMemory_SemanticOverrideExtractionConfigurationInputProperty struct {
	// Text prompt for model instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
}

