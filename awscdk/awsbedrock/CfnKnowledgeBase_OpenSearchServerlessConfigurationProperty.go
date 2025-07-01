package awsbedrock


// Contains details about the storage configuration of the knowledge base in Amazon OpenSearch Service.
//
// For more information, see [Create a vector index in Amazon OpenSearch Service](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-oss.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openSearchServerlessConfigurationProperty := &OpenSearchServerlessConfigurationProperty{
//   	CollectionArn: jsii.String("collectionArn"),
//   	FieldMapping: &OpenSearchServerlessFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		TextField: jsii.String("textField"),
//   		VectorField: jsii.String("vectorField"),
//   	},
//   	VectorIndexName: jsii.String("vectorIndexName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessconfiguration.html
//
type CfnKnowledgeBase_OpenSearchServerlessConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the OpenSearch Service vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessconfiguration.html#cfn-bedrock-knowledgebase-opensearchserverlessconfiguration-collectionarn
	//
	CollectionArn *string `field:"required" json:"collectionArn" yaml:"collectionArn"`
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessconfiguration.html#cfn-bedrock-knowledgebase-opensearchserverlessconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessconfiguration.html#cfn-bedrock-knowledgebase-opensearchserverlessconfiguration-vectorindexname
	//
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

