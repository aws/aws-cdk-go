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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3filelocation.html
//
type CfnWorkflow_S3FileLocationProperty struct {
	// Specifies the details for the file location for the file that's being used in the workflow.
	//
	// Only applicable if you are using Amazon S3 storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3filelocation.html#cfn-transfer-workflow-s3filelocation-s3filelocation
	//
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

