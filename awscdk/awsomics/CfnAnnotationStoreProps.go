package awsomics


// Properties for defining a `CfnAnnotationStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnAnnotationStoreProps := &CfnAnnotationStoreProps{
//   	Name: jsii.String("name"),
//   	StoreFormat: jsii.String("storeFormat"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Reference: &ReferenceItemProperty{
//   		ReferenceArn: jsii.String("referenceArn"),
//   	},
//   	SseConfig: &SseConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	StoreOptions: &StoreOptionsProperty{
//   		TsvStoreOptions: &TsvStoreOptionsProperty{
//   			AnnotationType: jsii.String("annotationType"),
//   			FormatToHeader: map[string]*string{
//   				"formatToHeaderKey": jsii.String("formatToHeader"),
//   			},
//   			Schema: schema,
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnAnnotationStoreProps struct {
	// The name of the Annotation Store.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The annotation file format of the store.
	StoreFormat *string `field:"required" json:"storeFormat" yaml:"storeFormat"`
	// A description for the store.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The genome reference for the store's annotations.
	Reference interface{} `field:"optional" json:"reference" yaml:"reference"`
	// The store's server-side encryption (SSE) settings.
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// File parsing options for the annotation store.
	StoreOptions interface{} `field:"optional" json:"storeOptions" yaml:"storeOptions"`
	// Tags for the store.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

