package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Properties for a vector knowledge base.
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
//   	},
//   }).With(mixins.NewKnowledgeBaseTypeConfiguration(bedrock.KnowledgeBaseType_Vector(&VectorKnowledgeBaseTypeProps{
//   	EmbeddingsModel: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024(),
//   	VectorStore: bedrock.VectorStore_OpenSearchServerless(&OpenSearchServerlessVectorStoreProps{
//   		Collection: *Collection,
//   		VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   		VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   		TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   		MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   	}),
//   })))
//
// Experimental.
type VectorKnowledgeBaseTypeProps struct {
	// The model used to create vector embeddings for the knowledge base.
	// Experimental.
	EmbeddingsModel BedrockFoundationModel `field:"required" json:"embeddingsModel" yaml:"embeddingsModel"`
	// The vector store in which the embeddings are stored and searched.
	// Experimental.
	VectorStore VectorStore `field:"required" json:"vectorStore" yaml:"vectorStore"`
	// The Amazon S3 bucket in which multimedia content extracted from multimodal documents (images, audio, video) is stored.
	//
	// Required when the knowledge base ingests multimodal content.
	// Default: - no supplemental data storage; only text content is ingested.
	//
	// Experimental.
	SupplementalDataStorageBucket interfacesawss3.IBucketRef `field:"optional" json:"supplementalDataStorageBucket" yaml:"supplementalDataStorageBucket"`
	// The data type of the vector embeddings.
	// Default: VectorType.FLOATING_POINT
	//
	// Experimental.
	VectorType VectorType `field:"optional" json:"vectorType" yaml:"vectorType"`
}

