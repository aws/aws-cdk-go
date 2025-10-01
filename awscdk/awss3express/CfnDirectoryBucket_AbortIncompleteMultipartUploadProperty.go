package awss3express


// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
//
// For more information, see [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortIncompleteMultipartUploadProperty := &AbortIncompleteMultipartUploadProperty{
//   	DaysAfterInitiation: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-abortincompletemultipartupload.html
//
type CfnDirectoryBucket_AbortIncompleteMultipartUploadProperty struct {
	// Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-abortincompletemultipartupload.html#cfn-s3express-directorybucket-abortincompletemultipartupload-daysafterinitiation
	//
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

