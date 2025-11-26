package previewawsdatabrewmixins


// The location in Amazon S3 or AWS Glue Data Catalog where the job writes its output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputLocationProperty := &OutputLocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	BucketOwner: jsii.String("bucketOwner"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputlocation.html
//
type CfnJobPropsMixin_OutputLocationProperty struct {
	// The Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputlocation.html#cfn-databrew-job-outputlocation-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputlocation.html#cfn-databrew-job-outputlocation-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The unique name of the object in the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputlocation.html#cfn-databrew-job-outputlocation-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

