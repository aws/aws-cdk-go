package awsbedrockalphamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchserverless"
)

// Properties for `KnowledgeBaseOpenSearchServerlessStorage`.
//
// Example:
//   var collection ICollectionRef
//   var role Role
//
//
//   awscdk.Aws_bedrock.NewCfnKnowledgeBase(this, jsii.String("KnowledgeBase"), &CfnKnowledgeBaseProps{
//   	Name: jsii.String("product-documentation"),
//   	RoleArn: role.roleArn,
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("VECTOR"),
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024().ModelArn,
//   		},
//   	},
//   }).With(mixins.NewKnowledgeBaseOpenSearchServerlessStorage(&KnowledgeBaseOpenSearchServerlessStorageProps{
//   	Collection: Collection,
//   	VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   	VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   	TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   	MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   }))
//
// Experimental.
type KnowledgeBaseOpenSearchServerlessStorageProps struct {
	// The OpenSearch Serverless collection that contains the vector index.
	// Experimental.
	Collection interfacesawsopensearchserverless.ICollectionRef `field:"required" json:"collection" yaml:"collection"`
	// The name of the index field that stores the metadata of each chunk.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The name of the index field that stores the raw text of each chunk.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// The name of the index field that stores the vector embeddings.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
	// The name of the vector index in the collection.
	// Experimental.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

