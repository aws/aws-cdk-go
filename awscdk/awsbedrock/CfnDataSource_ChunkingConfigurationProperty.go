package awsbedrock


// Details about how to chunk the documents in the data source.
//
// A *chunk* refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-chunkingconfiguration.html
//
type CfnDataSource_ChunkingConfigurationProperty struct {
	// Knowledge base can split your source data into chunks.
	//
	// A *chunk* refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for `NONE` , then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.
	//
	// - `FIXED_SIZE` – Amazon Bedrock splits your source data into chunks of the approximate size that you set in the `fixedSizeChunkingConfiguration` .
	// - `HIERARCHICAL` – Split documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.
	// - `SEMANTIC` – Split documents into chunks based on groups of similar content derived with natural language processing.
	// - `NONE` – Amazon Bedrock treats each file as one chunk. If you choose this option, you may want to pre-process your documents by splitting them into separate files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-chunkingconfiguration.html#cfn-bedrock-datasource-chunkingconfiguration-chunkingstrategy
	//
	ChunkingStrategy *string `field:"required" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// Configurations for when you choose fixed-size chunking.
	//
	// If you set the `chunkingStrategy` as `NONE` , exclude this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-chunkingconfiguration.html#cfn-bedrock-datasource-chunkingconfiguration-fixedsizechunkingconfiguration
	//
	FixedSizeChunkingConfiguration interface{} `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// Settings for hierarchical document chunking for a data source.
	//
	// Hierarchical chunking splits documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-chunkingconfiguration.html#cfn-bedrock-datasource-chunkingconfiguration-hierarchicalchunkingconfiguration
	//
	HierarchicalChunkingConfiguration interface{} `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// Settings for semantic document chunking for a data source.
	//
	// Semantic chunking splits a document into into smaller documents based on groups of similar content derived from the text with natural language processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-chunkingconfiguration.html#cfn-bedrock-datasource-chunkingconfiguration-semanticchunkingconfiguration
	//
	SemanticChunkingConfiguration interface{} `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

