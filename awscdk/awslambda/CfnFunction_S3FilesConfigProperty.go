package awslambda


// Setting controls how your function accesses data from an Amazon S3 file system.
//
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
	//
	// Valid values:
	//   +  ``AUTO`` (default) – Direct reads are active for functions you configure with 512 MB or more of memory.
	//   +  ``ENABLED`` – Enforces all reads are directly from the Amazon S3 bucket, regardless of available memory (less than 512 MB).
	//   +  ``DISABLED`` – Routes all reads through the file system, regardless of memory configuration.
	//
	//  To use direct reads, you must grant the execution role the ``s3:GetObject`` and ``s3:GetObjectVersion`` permissions. If a direct read fails, Lambda automatically falls back to reading through the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-s3filesconfig.html#cfn-lambda-function-s3filesconfig-directs3read
	//
	DirectS3Read *string `field:"optional" json:"directS3Read" yaml:"directS3Read"`
}

