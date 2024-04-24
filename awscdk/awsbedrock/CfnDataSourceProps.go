package awsbedrock


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//
//   			// the properties below are optional
//   			InclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	VectorIngestionConfiguration: &VectorIngestionConfigurationProperty{
//   		ChunkingConfiguration: &ChunkingConfigurationProperty{
//   			ChunkingStrategy: jsii.String("chunkingStrategy"),
//
//   			// the properties below are optional
//   			FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				OverlapPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html
//
type CfnDataSourceProps struct {
	// Contains details about how the data source is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"required" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The unique identifier of the knowledge base to which the data source belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the configuration of the server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Contains details about how to ingest the documents in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-vectoringestionconfiguration
	//
	VectorIngestionConfiguration interface{} `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}

