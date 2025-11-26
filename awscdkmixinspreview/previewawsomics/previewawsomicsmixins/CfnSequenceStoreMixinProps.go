package previewawsomicsmixins


// Properties for CfnSequenceStorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var s3AccessPolicy interface{}
//
//   cfnSequenceStoreMixinProps := &CfnSequenceStoreMixinProps{
//   	AccessLogLocation: jsii.String("accessLogLocation"),
//   	Description: jsii.String("description"),
//   	ETagAlgorithmFamily: jsii.String("eTagAlgorithmFamily"),
//   	FallbackLocation: jsii.String("fallbackLocation"),
//   	Name: jsii.String("name"),
//   	PropagatedSetLevelTags: []*string{
//   		jsii.String("propagatedSetLevelTags"),
//   	},
//   	S3AccessPolicy: s3AccessPolicy,
//   	SseConfig: &SseConfigProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html
//
type CfnSequenceStoreMixinProps struct {
	// Location of the access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-accessloglocation
	//
	AccessLogLocation *string `field:"optional" json:"accessLogLocation" yaml:"accessLogLocation"`
	// A description for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The algorithm family of the ETag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-etagalgorithmfamily
	//
	ETagAlgorithmFamily *string `field:"optional" json:"eTagAlgorithmFamily" yaml:"eTagAlgorithmFamily"`
	// An S3 location that is used to store files that have failed a direct upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-fallbacklocation
	//
	FallbackLocation *string `field:"optional" json:"fallbackLocation" yaml:"fallbackLocation"`
	// A name for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags keys to propagate to the S3 objects associated with read sets in the sequence store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-propagatedsetleveltags
	//
	PropagatedSetLevelTags *[]*string `field:"optional" json:"propagatedSetLevelTags" yaml:"propagatedSetLevelTags"`
	// The resource policy that controls S3 access on the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-s3accesspolicy
	//
	S3AccessPolicy interface{} `field:"optional" json:"s3AccessPolicy" yaml:"s3AccessPolicy"`
	// Server-side encryption (SSE) settings for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-sseconfig
	//
	SseConfig interface{} `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Tags for the store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html#cfn-omics-sequencestore-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

