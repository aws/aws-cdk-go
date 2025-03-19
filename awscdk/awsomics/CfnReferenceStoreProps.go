package awsomics


// Properties for defining a `CfnReferenceStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReferenceStoreProps := &CfnReferenceStoreProps{
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-referencestore.html
//
type CfnReferenceStoreProps struct {
	// A name for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-referencestore.html#cfn-omics-referencestore-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-referencestore.html#cfn-omics-referencestore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Server-side encryption (SSE) settings for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-referencestore.html#cfn-omics-referencestore-sseconfig
	//
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-referencestore.html#cfn-omics-referencestore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

