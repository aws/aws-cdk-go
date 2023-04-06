package awsomics


// Properties for defining a `CfnVariantStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVariantStoreProps := &CfnVariantStoreProps{
//   	Name: jsii.String("name"),
//   	Reference: &ReferenceItemProperty{
//   		ReferenceArn: jsii.String("referenceArn"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
type CfnVariantStoreProps struct {
	// A name for the store.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The genome reference for the store's variants.
	Reference interface{} `field:"required" json:"reference" yaml:"reference"`
	// A description for the store.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Server-side encryption (SSE) settings for the store.
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

