package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The vector store that backs a vector knowledge base.
//
// Use one of the static factory methods to reference an existing store.
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
// Experimental.
type VectorStore interface {
}

// The jsii proxy struct for VectorStore
type jsiiProxy_VectorStore struct {
	_ byte // padding
}

// An existing vector index in an Amazon OpenSearch Serverless collection.
// Experimental.
func VectorStore_OpenSearchServerless(props *OpenSearchServerlessVectorStoreProps) VectorStore {
	_init_.Initialize()

	if err := validateVectorStore_OpenSearchServerlessParameters(props); err != nil {
		panic(err)
	}
	var returns VectorStore

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.VectorStore",
		"openSearchServerless",
		[]interface{}{props},
		&returns,
	)

	return returns
}

