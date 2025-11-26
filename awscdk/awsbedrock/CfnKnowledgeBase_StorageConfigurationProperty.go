package awsbedrock


// Contains the storage configuration of the knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConfigurationProperty := &StorageConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	MongoDbAtlasConfiguration: &MongoDbAtlasConfigurationProperty{
//   		CollectionName: jsii.String("collectionName"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Endpoint: jsii.String("endpoint"),
//   		FieldMapping: &MongoDbAtlasFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//   		},
//   		VectorIndexName: jsii.String("vectorIndexName"),
//
//   		// the properties below are optional
//   		EndpointServiceName: jsii.String("endpointServiceName"),
//   		TextIndexName: jsii.String("textIndexName"),
//   	},
//   	NeptuneAnalyticsConfiguration: &NeptuneAnalyticsConfigurationProperty{
//   		FieldMapping: &NeptuneAnalyticsFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   		},
//   		GraphArn: jsii.String("graphArn"),
//   	},
//   	OpensearchManagedClusterConfiguration: &OpenSearchManagedClusterConfigurationProperty{
//   		DomainArn: jsii.String("domainArn"),
//   		DomainEndpoint: jsii.String("domainEndpoint"),
//   		FieldMapping: &OpenSearchManagedClusterFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//   		},
//   		VectorIndexName: jsii.String("vectorIndexName"),
//   	},
//   	OpensearchServerlessConfiguration: &OpenSearchServerlessConfigurationProperty{
//   		CollectionArn: jsii.String("collectionArn"),
//   		FieldMapping: &OpenSearchServerlessFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//   		},
//   		VectorIndexName: jsii.String("vectorIndexName"),
//   	},
//   	PineconeConfiguration: &PineconeConfigurationProperty{
//   		ConnectionString: jsii.String("connectionString"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		FieldMapping: &PineconeFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   		},
//
//   		// the properties below are optional
//   		Namespace: jsii.String("namespace"),
//   	},
//   	RdsConfiguration: &RdsConfigurationProperty{
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		FieldMapping: &RdsFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			PrimaryKeyField: jsii.String("primaryKeyField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//
//   			// the properties below are optional
//   			CustomMetadataField: jsii.String("customMetadataField"),
//   		},
//   		ResourceArn: jsii.String("resourceArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	S3VectorsConfiguration: &S3VectorsConfigurationProperty{
//   		IndexArn: jsii.String("indexArn"),
//   		IndexName: jsii.String("indexName"),
//   		VectorBucketArn: jsii.String("vectorBucketArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html
//
type CfnKnowledgeBase_StorageConfigurationProperty struct {
	// The vector store service in which the knowledge base is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Contains the storage configuration of the knowledge base in MongoDB Atlas.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-mongodbatlasconfiguration
	//
	MongoDbAtlasConfiguration interface{} `field:"optional" json:"mongoDbAtlasConfiguration" yaml:"mongoDbAtlasConfiguration"`
	// Contains details about the Neptune Analytics configuration of the knowledge base in Amazon Neptune.
	//
	// For more information, see [Create a vector index in Amazon Neptune Analytics.](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-neptune.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-neptuneanalyticsconfiguration
	//
	NeptuneAnalyticsConfiguration interface{} `field:"optional" json:"neptuneAnalyticsConfiguration" yaml:"neptuneAnalyticsConfiguration"`
	// Contains details about the storage configuration of the knowledge base in OpenSearch Managed Cluster.
	//
	// For more information, see [Create a vector index in Amazon OpenSearch Service](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-osm.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-opensearchmanagedclusterconfiguration
	//
	OpensearchManagedClusterConfiguration interface{} `field:"optional" json:"opensearchManagedClusterConfiguration" yaml:"opensearchManagedClusterConfiguration"`
	// Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-opensearchserverlessconfiguration
	//
	OpensearchServerlessConfiguration interface{} `field:"optional" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// Contains the storage configuration of the knowledge base in Pinecone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-pineconeconfiguration
	//
	PineconeConfiguration interface{} `field:"optional" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// Contains details about the storage configuration of the knowledge base in Amazon RDS.
	//
	// For more information, see [Create a vector index in Amazon RDS](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-rds.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-rdsconfiguration
	//
	RdsConfiguration interface{} `field:"optional" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// Contains the storage configuration of the knowledge base for S3 vectors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-s3vectorsconfiguration
	//
	S3VectorsConfiguration interface{} `field:"optional" json:"s3VectorsConfiguration" yaml:"s3VectorsConfiguration"`
}

