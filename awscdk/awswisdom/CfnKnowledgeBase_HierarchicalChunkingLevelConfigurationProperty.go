package awswisdom


// Token settings for each layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hierarchicalChunkingLevelConfigurationProperty := &HierarchicalChunkingLevelConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkinglevelconfiguration.html
//
type CfnKnowledgeBase_HierarchicalChunkingLevelConfigurationProperty struct {
	// The maximum number of tokens that a chunk can contain in this layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkinglevelconfiguration.html#cfn-wisdom-knowledgebase-hierarchicalchunkinglevelconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
}

