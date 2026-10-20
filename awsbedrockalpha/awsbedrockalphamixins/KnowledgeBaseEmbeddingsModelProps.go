package awsbedrockalphamixins

import (
	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2"
)

// Properties for `KnowledgeBaseEmbeddingsModel`.
//
// Example:
//   var role Role
//
//
//   awscdk.Aws_bedrock.NewCfnKnowledgeBase(this, jsii.String("KnowledgeBase"), &CfnKnowledgeBaseProps{
//   	Name: jsii.String("product-documentation"),
//   	RoleArn: role.roleArn,
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("VECTOR"),
//   	},
//   }).With(mixins.NewKnowledgeBaseEmbeddingsModel(&KnowledgeBaseEmbeddingsModelProps{
//   	EmbeddingsModel: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024(),
//   }))
//
// Experimental.
type KnowledgeBaseEmbeddingsModelProps struct {
	// The model used to create vector embeddings for the knowledge base.
	// Experimental.
	EmbeddingsModel awsbedrockalpha.BedrockFoundationModel `field:"required" json:"embeddingsModel" yaml:"embeddingsModel"`
	// The data type of the vector embeddings.
	// Default: VectorType.FLOATING_POINT
	//
	// Experimental.
	VectorType awsbedrockalpha.VectorType `field:"optional" json:"vectorType" yaml:"vectorType"`
}

