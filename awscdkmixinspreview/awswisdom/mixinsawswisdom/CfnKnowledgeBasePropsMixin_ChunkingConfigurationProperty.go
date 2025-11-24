package mixinsawswisdom


// Details about how to chunk the documents in the data source.
//
// A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   chunkingConfigurationProperty := &ChunkingConfigurationProperty{
//   	ChunkingStrategy: jsii.String("chunkingStrategy"),
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
type CfnKnowledgeBasePropsMixin_ChunkingConfigurationProperty struct {
	// Knowledge base can split your source data into chunks.
	//
	// A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for `NONE` , then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-chunkingstrategy
	//
	ChunkingStrategy *string `field:"optional" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// Configurations for when you choose fixed-size chunking.
	//
	// If you set the `chunkingStrategy` as `NONE` , exclude this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-fixedsizechunkingconfiguration
	//
	FixedSizeChunkingConfiguration interface{} `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// Settings for hierarchical document chunking for a data source.
	//
	// Hierarchical chunking splits documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-hierarchicalchunkingconfiguration
	//
	HierarchicalChunkingConfiguration interface{} `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// Settings for semantic document chunking for a data source.
	//
	// Semantic chunking splits a document into smaller documents based on groups of similar content derived from the text with natural language processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-chunkingconfiguration.html#cfn-wisdom-knowledgebase-chunkingconfiguration-semanticchunkingconfiguration
	//
	SemanticChunkingConfiguration interface{} `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

