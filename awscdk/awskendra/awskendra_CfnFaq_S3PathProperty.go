package awskendra


// Information required to find a specific file in an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3PathProperty := &s3PathProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnFaq_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	Key *string `field:"required" json:"key" yaml:"key"`
}

