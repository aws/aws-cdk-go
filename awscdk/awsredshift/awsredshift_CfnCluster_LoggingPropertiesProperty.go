package awsredshift


// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingPropertiesProperty := &loggingPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnCluster_LoggingPropertiesProperty struct {
	// The name of an existing S3 bucket where the log files are to be stored.
	//
	// Constraints:
	//
	// - Must be in the same region as the cluster
	// - The cluster must have read bucket and put object permissions.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The prefix applied to the log file names.
	//
	// Constraints:
	//
	// - Cannot exceed 512 characters
	// - Cannot contain spaces( ), double quotes ("), single quotes ('), a backslash (\), or control characters. The hexadecimal codes for invalid characters are:
	//
	// - x00 to x20
	// - x22
	// - x27
	// - x5c
	// - x7f or larger.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

