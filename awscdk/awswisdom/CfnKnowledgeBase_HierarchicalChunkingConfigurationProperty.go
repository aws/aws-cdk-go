package awswisdom


// Settings for hierarchical document chunking for a data source.
//
// Hierarchical chunking splits documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hierarchicalChunkingConfigurationProperty := &HierarchicalChunkingConfigurationProperty{
//   	LevelConfigurations: []interface{}{
//   		&HierarchicalChunkingLevelConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   		},
//   	},
//   	OverlapTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html
//
type CfnKnowledgeBase_HierarchicalChunkingConfigurationProperty struct {
	// Token settings for each layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html#cfn-wisdom-knowledgebase-hierarchicalchunkingconfiguration-levelconfigurations
	//
	LevelConfigurations interface{} `field:"required" json:"levelConfigurations" yaml:"levelConfigurations"`
	// The number of tokens to repeat across chunks in the same layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html#cfn-wisdom-knowledgebase-hierarchicalchunkingconfiguration-overlaptokens
	//
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
}

