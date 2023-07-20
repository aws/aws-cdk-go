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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html
//
type CfnSequenceStoreProps struct {
	// A name for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An S3 URI representing the bucket and folder to store failed read set uploads.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-fallbacklocation
	//
	FallbackLocation *string `field:"optional" json:"fallbackLocation" yaml:"fallbackLocation"`
	// Server-side encryption (SSE) settings for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-sseconfig
	//
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

