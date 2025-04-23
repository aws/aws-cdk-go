package awsbedrock


// Contains details about the storage configuration of the knowledge base in MongoDB Atlas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDbAtlasConfigurationProperty := &MongoDbAtlasConfigurationProperty{
//   	CollectionName: jsii.String("collectionName"),
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Endpoint: jsii.String("endpoint"),
//   	FieldMapping: &MongoDbAtlasFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		TextField: jsii.String("textField"),
//   		VectorField: jsii.String("vectorField"),
//   	},
//   	VectorIndexName: jsii.String("vectorIndexName"),
//
//   	// the properties below are optional
//   	EndpointServiceName: jsii.String("endpointServiceName"),
//   	TextIndexName: jsii.String("textIndexName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html
//
type CfnKnowledgeBase_MongoDbAtlasConfigurationProperty struct {
	// The collection name of the knowledge base in MongoDB Atlas.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-collectionname
	//
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that contains user credentials for your MongoDB Atlas cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The database name in your MongoDB Atlas cluster for your knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The endpoint URL of your MongoDB Atlas cluster for your knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the MongoDB Atlas vector search index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-vectorindexname
	//
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
	// The name of the VPC endpoint service in your account that is connected to your MongoDB Atlas cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-endpointservicename
	//
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	// Name of a MongoDB Atlas text index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.html#cfn-bedrock-knowledgebase-mongodbatlasconfiguration-textindexname
	//
	TextIndexName *string `field:"optional" json:"textIndexName" yaml:"textIndexName"`
}

