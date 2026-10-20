package awsbedrockalphamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Properties for `KnowledgeBaseSupplementalDataStorage`.
//
// Example:
//   var knowledgeBase CfnKnowledgeBase
//   var bucket IBucket
//
//
//   knowledgeBase.With(mixins.NewKnowledgeBaseSupplementalDataStorage(&KnowledgeBaseSupplementalDataStorageProps{
//   	Bucket: Bucket,
//   }))
//
// Experimental.
type KnowledgeBaseSupplementalDataStorageProps struct {
	// The S3 bucket in which multimedia content extracted from multimodal documents is stored.
	//
	// Amazon Bedrock manages the layout of the extracted content within the
	// bucket; a key prefix cannot be specified.
	// Experimental.
	Bucket interfacesawss3.IBucketRef `field:"required" json:"bucket" yaml:"bucket"`
}

