package awsomics


// Properties for defining a `CfnSequenceStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSequenceStoreProps := &CfnSequenceStoreProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FallbackLocation: jsii.String("fallbackLocation"),
//   	SseConfig: &SseConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSequenceStoreProps struct {
	// A name for the store.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Omics::SequenceStore.FallbackLocation`.
	FallbackLocation *string `field:"optional" json:"fallbackLocation" yaml:"fallbackLocation"`
	// Server-side encryption (SSE) settings for the store.
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

