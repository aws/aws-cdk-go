package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of a knowledge base, which determines how its data is stored and queried.
//
// Use one of the static factory methods to configure a type. To configure a
// `CfnKnowledgeBase` as a type, apply `mixins.KnowledgeBaseTypeConfiguration`.
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
type KnowledgeBaseType interface {
}

// The jsii proxy struct for KnowledgeBaseType
type jsiiProxy_KnowledgeBaseType struct {
	_ byte // padding
}

// A knowledge base that stores vector embeddings of ingested documents in a vector store and retrieves them by similarity search.
// Experimental.
func KnowledgeBaseType_Vector(props *VectorKnowledgeBaseTypeProps) KnowledgeBaseType {
	_init_.Initialize()

	if err := validateKnowledgeBaseType_VectorParameters(props); err != nil {
		panic(err)
	}
	var returns KnowledgeBaseType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.KnowledgeBaseType",
		"vector",
		[]interface{}{props},
		&returns,
	)

	return returns
}

