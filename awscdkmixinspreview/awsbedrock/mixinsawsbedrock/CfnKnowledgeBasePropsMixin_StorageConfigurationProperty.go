package mixinsawsbedrock


// Contains the storage configuration of the knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageConfigurationProperty := &StorageConfigurationProperty{
//   	MongoDbAtlasConfiguration: &MongoDbAtlasConfigurationProperty{
//   		CollectionName: jsii.String("collectionName"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Endpoint: jsii.String("endpoint"),
//   		EndpointServiceName: jsii.String("endpointServiceName"),
//   		FieldMapping: &MongoDbAtlasFieldMappingProperty{
//   			MetadataField: jsii.String("metadataField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//   		},
//   		TextIndexName: jsii.String("textIndexName"),
//   		VectorIndexName: jsii.String("vectorIndexName"),
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
//   		Namespace: jsii.String("namespace"),
//   	},
//   	RdsConfiguration: &RdsConfigurationProperty{
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		FieldMapping: &RdsFieldMappingProperty{
//   			CustomMetadataField: jsii.String("customMetadataField"),
//   			MetadataField: jsii.String("metadataField"),
//   			PrimaryKeyField: jsii.String("primaryKeyField"),
//   			TextField: jsii.String("textField"),
//   			VectorField: jsii.String("vectorField"),
//   		},
//   		ResourceArn: jsii.String("resourceArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html
//
type CfnKnowledgeBasePropsMixin_StorageConfigurationProperty struct {
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
	// The vector store service in which the knowledge base is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

