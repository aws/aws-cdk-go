package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Properties for a knowledge base.
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
type KnowledgeBaseProps struct {
	// The type of the knowledge base and its type-specific configuration.
	// Experimental.
	Type KnowledgeBaseType `field:"required" json:"type" yaml:"type"`
	// A description of the knowledge base.
	//
	// Must be 1-200 characters.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the knowledge base.
	//
	// Must be 1-100 characters of letters, digits, hyphens and underscores, and
	// must not contain consecutive hyphens or underscores.
	// Default: - a unique name is generated.
	//
	// Experimental.
	KnowledgeBaseName *string `field:"optional" json:"knowledgeBaseName" yaml:"knowledgeBaseName"`
	// The service role that Amazon Bedrock assumes to operate the knowledge base.
	//
	// The permissions the knowledge base needs to reach its data store and
	// models are granted to this role.
	// Default: - a new role is created, trusted by `bedrock.amazonaws.com` and
	// scoped to knowledge bases in this account.
	//
	// Experimental.
	Role interface{ interfacesawsiam.IRoleRef;awsiam.IGrantable } `field:"optional" json:"role" yaml:"role"`
	// Tags to apply to the knowledge base.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

