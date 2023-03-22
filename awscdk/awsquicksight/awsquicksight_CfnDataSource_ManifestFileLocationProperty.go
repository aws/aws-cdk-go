package awsquicksight


// Amazon S3 manifest file location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestFileLocationProperty := &manifestFileLocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnDataSource_ManifestFileLocationProperty struct {
	// Amazon S3 bucket.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Amazon S3 key that identifies an object.
	Key *string `field:"required" json:"key" yaml:"key"`
}

