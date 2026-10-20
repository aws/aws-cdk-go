package awsbedrockalpha


// The data type for the vectors when using a model to convert text into vector embeddings.
//
// The model must support the specified data type for vector embeddings. Floating-point (float32)
// is the default data type, and is supported by most models for vector embeddings. See Supported
// embeddings models for information on the available models and their vector data types.
//
// Example:
//   var collection ICollectionRef
//   var role Role
//   var multimodalBucket IBucket
//
//
//   bedrock.NewKnowledgeBase(this, jsii.String("KnowledgeBase"), &KnowledgeBaseProps{
//   	KnowledgeBaseName: jsii.String("product-documentation"),
//   	Description: jsii.String("Product manuals and release notes"),
//   	Type: bedrock.KnowledgeBaseType_Vector(&VectorKnowledgeBaseTypeProps{
//   		EmbeddingsModel: bedrock.BedrockFoundationModel_COHERE_EMBED_MULTILINGUAL_V3(),
//   		// Store binary vectors: less precise, but cheaper to store and search
//   		VectorType: bedrock.VectorType_BINARY,
//   		VectorStore: bedrock.VectorStore_OpenSearchServerless(&OpenSearchServerlessVectorStoreProps{
//   			Collection: *Collection,
//   			VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   			VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   			TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   			MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   		}),
//   		// Where images, audio and video extracted from multimodal documents are stored
//   		SupplementalDataStorageBucket: multimodalBucket,
//   	}),
//   	// Bring your own service role; the model and vector store permissions are added to it
//   	Role: Role,
//   	Tags: map[string]*string{
//   		"Team": jsii.String("Documentation"),
//   	},
//   })
//
// Experimental.
type VectorType string

const (
	// `FLOATING_POINT` convert the data to floating-point (float32) vector embeddings (more precise, but more costly).
	// Experimental.
	VectorType_FLOATING_POINT VectorType = "FLOATING_POINT"
	// `BINARY` convert the data to binary vector embeddings (less precise, but less costly).
	// Experimental.
	VectorType_BINARY VectorType = "BINARY"
)

