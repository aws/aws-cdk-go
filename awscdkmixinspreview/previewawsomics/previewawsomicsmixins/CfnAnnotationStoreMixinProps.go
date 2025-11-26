package previewawsomicsmixins


// Properties for CfnAnnotationStorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schema interface{}
//
//   cfnAnnotationStoreMixinProps := &CfnAnnotationStoreMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Reference: &ReferenceItemProperty{
//   		ReferenceArn: jsii.String("referenceArn"),
//   	},
//   	SseConfig: &SseConfigProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	StoreFormat: jsii.String("storeFormat"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html
//
type CfnAnnotationStoreMixinProps struct {
	// A description for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Annotation Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The genome reference for the store's annotations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-reference
	//
	Reference interface{} `field:"optional" json:"reference" yaml:"reference"`
	// The store's server-side encryption (SSE) settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-sseconfig
	//
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// The annotation file format of the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-storeformat
	//
	StoreFormat *string `field:"optional" json:"storeFormat" yaml:"storeFormat"`
	// File parsing options for the annotation store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-storeoptions
	//
	StoreOptions interface{} `field:"optional" json:"storeOptions" yaml:"storeOptions"`
	// Tags for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-annotationstore.html#cfn-omics-annotationstore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

