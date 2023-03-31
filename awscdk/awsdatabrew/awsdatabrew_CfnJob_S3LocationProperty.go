package awsdatabrew


// Represents an Amazon S3 location (bucket name, bucket owner, and object key) where DataBrew can read input data, or write output from a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketOwner: jsii.String("bucketOwner"),
//   	key: jsii.String("key"),
//   }
//
type CfnJob_S3LocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID of the bucket owner.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

