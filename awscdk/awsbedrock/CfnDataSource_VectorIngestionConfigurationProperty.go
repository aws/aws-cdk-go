package awsbedrock


// Contains details about how to ingest the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorIngestionConfigurationProperty := &VectorIngestionConfigurationProperty{
//   	ChunkingConfiguration: &ChunkingConfigurationProperty{
//   		ChunkingStrategy: jsii.String("chunkingStrategy"),
//
//   		// the properties below are optional
//   		FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			OverlapPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html
//
type CfnDataSource_VectorIngestionConfigurationProperty struct {
	// Details about how to chunk the documents in the data source.
	//
	// A *chunk* refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-chunkingconfiguration
	//
	ChunkingConfiguration interface{} `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
}

