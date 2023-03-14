package awsdatabrew


// The location in Amazon S3 or AWS Glue Data Catalog where the job writes its output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputLocationProperty := &outputLocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketOwner: jsii.String("bucketOwner"),
//   	key: jsii.String("key"),
//   }
//
type CfnJob_OutputLocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnJob.OutputLocationProperty.BucketOwner`.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

