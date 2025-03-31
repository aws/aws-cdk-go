package awsomics


// Properties for defining a `CfnSequenceStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var s3AccessPolicy interface{}
//
//   cfnSequenceStoreProps := &CfnSequenceStoreProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AccessLogLocation: jsii.String("accessLogLocation"),
//   	Description: jsii.String("description"),
//   	ETagAlgorithmFamily: jsii.String("eTagAlgorithmFamily"),
//   	FallbackLocation: jsii.String("fallbackLocation"),
//   	PropagatedSetLevelTags: []*string{
//   		jsii.String("propagatedSetLevelTags"),
//   	},
//   	S3AccessPolicy: s3AccessPolicy,
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

