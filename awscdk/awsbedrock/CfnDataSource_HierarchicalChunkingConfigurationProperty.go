package awsbedrock


// Settings for hierarchical document chunking for a data source.
//
// Hierarchical chunking splits documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.
//
// You configure the number of tokens to overlap, or repeat across adjacent chunks. For example, if you set overlap tokens to 60, the last 60 tokens in the first chunk are also included at the beginning of the second chunk. For each layer, you must also configure the maximum number of tokens in a chunk.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-hierarchicalchunkingconfiguration.html
//
type CfnDataSource_HierarchicalChunkingConfigurationProperty struct {
	// Token settings for each layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-hierarchicalchunkingconfiguration.html#cfn-bedrock-datasource-hierarchicalchunkingconfiguration-levelconfigurations
	//
	LevelConfigurations interface{} `field:"required" json:"levelConfigurations" yaml:"levelConfigurations"`
	// The number of tokens to repeat across chunks in the same layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-hierarchicalchunkingconfiguration.html#cfn-bedrock-datasource-hierarchicalchunkingconfiguration-overlaptokens
	//
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
}

