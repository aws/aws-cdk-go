package awss3


// This resource contains the details of the buckets and Regions for the Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketsAndRegionsProperty := &BucketsAndRegionsProperty{
//   	Buckets: []*string{
//   		jsii.String("buckets"),
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketsandregions.html
//
type CfnStorageLens_BucketsAndRegionsProperty struct {
	// This property contains the details of the buckets for the Amazon S3 Storage Lens configuration.
	//
	// This should be the bucket Amazon Resource Name(ARN). For valid values, see [Buckets ARN format here](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_Include.html#API_control_Include_Contents) in the *Amazon S3 API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketsandregions.html#cfn-s3-storagelens-bucketsandregions-buckets
	//
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// This property contains the details of the Regions for the S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketsandregions.html#cfn-s3-storagelens-bucketsandregions-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

