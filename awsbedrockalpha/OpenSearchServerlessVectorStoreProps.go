package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchserverless"
)

// Properties for an Amazon OpenSearch Serverless vector store.
//
// The collection and its vector index must already exist. Amazon Bedrock
// validates at creation time that the index exists in the collection and
// that its vector field matches the embeddings model. Specify the field
// names exactly as they are defined in the index.
//
// Example:
//   var collection ICollectionRef
//
//
//   knowledgeBase := bedrock.NewKnowledgeBase(this, jsii.String("KnowledgeBase"), &KnowledgeBaseProps{
//   	Type: bedrock.KnowledgeBaseType_Vector(&VectorKnowledgeBaseTypeProps{
//   		EmbeddingsModel: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024(),
//   		VectorStore: bedrock.VectorStore_OpenSearchServerless(&OpenSearchServerlessVectorStoreProps{
//   			Collection: *Collection,
//   			VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   			VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   			TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   			MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   		}),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-oss.html
//
// Experimental.
type OpenSearchServerlessVectorStoreProps struct {
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

