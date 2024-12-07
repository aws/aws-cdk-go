package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chunkingConfigurationProperty := &ChunkingConfigurationProperty{
//   	ChunkingStrategy: jsii.String("chunkingStrategy"),
//
//   	// the properties below are optional
//   	FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   		MaxTokens: jsii.Number(123),
//   		OverlapPercentage: jsii.Number(123),
//   	},
//   	HierarchicalChunkingConfiguration: &HierarchicalChunkingConfigurationProperty{
//   		LevelConfigurations: []interface{}{
//   			&HierarchicalChunkingLevelConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   			},
//   		},
//   		OverlapTokens: jsii.Number(123),
//   	},
//   	SemanticChunkingConfiguration: &SemanticChunkingConfigurationProperty{
//   		BreakpointPercentileThreshold: jsii.Number(123),
//   		BufferSize: jsii.Number(123),
//   		MaxTokens: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html
//
type CfnKnowledgeBase_ChunkingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-chunkingstrategy
	//
	ChunkingStrategy *string `field:"required" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-fixedsizechunkingconfiguration
	//
	FixedSizeChunkingConfiguration interface{} `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-hierarchicalchunkingconfiguration
	//
	HierarchicalChunkingConfiguration interface{} `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-semanticchunkingconfiguration
	//
	SemanticChunkingConfiguration interface{} `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

