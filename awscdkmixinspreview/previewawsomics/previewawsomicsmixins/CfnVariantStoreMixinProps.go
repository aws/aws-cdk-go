package previewawsomicsmixins


// Properties for CfnVariantStorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVariantStoreMixinProps := &CfnVariantStoreMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Reference: &ReferenceItemProperty{
//   		ReferenceArn: jsii.String("referenceArn"),
//   	},
//   	SseConfig: &SseConfigProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html
//
type CfnVariantStoreMixinProps struct {
	// A description for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The genome reference for the store's variants.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-reference
	//
	Reference interface{} `field:"optional" json:"reference" yaml:"reference"`
	// Server-side encryption (SSE) settings for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-sseconfig
	//
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

