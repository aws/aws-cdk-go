package awstransfer


// Specifies the S3 details for the file being used, such as bucket, ETag, and so forth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FileLocationProperty := &S3FileLocationProperty{
//   	S3FileLocation: &S3InputFileLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
type CfnWorkflow_S3FileLocationProperty struct {
	// Specifies the details for the file location for the file that's being used in the workflow.
	//
	// Only applicable if you are using Amazon S3 storage.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

