package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FilesConfigProperty := &S3FilesConfigProperty{
//   	DirectS3Read: jsii.String("directS3Read"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-s3filesconfig.html
//
type CfnFunction_S3FilesConfigProperty struct {
	// Specifies if a function reads from the file system for the lowest latency, or through Amazon S3 Files feature "direct Amazon S3 bucket reads" for the highest throughput.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-s3filesconfig.html#cfn-lambda-function-s3filesconfig-directs3read
	//
	DirectS3Read *string `field:"optional" json:"directS3Read" yaml:"directS3Read"`
}

