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
//   	key: jsii.String("key"),
//   }
//
type CfnRecipe_S3LocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

