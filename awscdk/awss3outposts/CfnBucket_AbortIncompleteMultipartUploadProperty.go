package awss3outposts


// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 on Outposts waits before permanently removing all parts of the upload.
//
// For more information, see [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html
//
type CfnBucket_AbortIncompleteMultipartUploadProperty struct {
	// Specifies the number of days after initiation that Amazon S3 on Outposts aborts an incomplete multipart upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html#cfn-s3outposts-bucket-abortincompletemultipartupload-daysafterinitiation
	//
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

