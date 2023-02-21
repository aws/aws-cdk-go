package awss3


// Describes the versioning state of an Amazon S3 bucket.
//
// For more information, see [PUT Bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versioningConfigurationProperty := &versioningConfigurationProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_VersioningConfigurationProperty struct {
	// The versioning state of the bucket.
	Status *string `field:"required" json:"status" yaml:"status"`
}

