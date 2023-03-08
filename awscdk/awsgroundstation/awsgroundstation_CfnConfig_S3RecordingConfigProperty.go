package awsgroundstation


// Provides information about how AWS Ground Station should save downlink data to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3RecordingConfigProperty := &s3RecordingConfigProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	prefix: jsii.String("prefix"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnConfig_S3RecordingConfigProperty struct {
	// S3 Bucket where the data is written.
	//
	// The name of the S3 Bucket provided must begin with `aws-groundstation` .
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The prefix of the S3 data object.
	//
	// If you choose to use any optional keys for substitution, these values will be replaced with the corresponding information from your contact details. For example, a prefix of `{satellite_id}/{year}/{month}/{day}/` will replaced with `fake_satellite_id/2021/01/10/`
	//
	// *Optional keys for substitution* : `{satellite_id}` | `{config-name}` | `{config-id}` | `{year}` | `{month}` | `{day}`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Defines the ARN of the role assumed for putting archives to S3.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

