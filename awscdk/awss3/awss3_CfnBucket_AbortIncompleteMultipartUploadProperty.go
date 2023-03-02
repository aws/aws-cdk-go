package awss3


// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
//
// For more information, see [Stopping Incomplete Multipart Uploads Using a Bucket Lifecycle Policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortIncompleteMultipartUploadProperty := &abortIncompleteMultipartUploadProperty{
//   	daysAfterInitiation: jsii.Number(123),
//   }
//
type CfnBucket_AbortIncompleteMultipartUploadProperty struct {
	// Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

