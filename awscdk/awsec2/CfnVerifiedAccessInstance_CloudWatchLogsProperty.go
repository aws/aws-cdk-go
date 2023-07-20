package awsec2


// Options for CloudWatch Logs as a logging destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsProperty := &CloudWatchLogsProperty{
//   	Enabled: jsii.Boolean(false),
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-cloudwatchlogs.html
//
type CfnVerifiedAccessInstance_CloudWatchLogsProperty struct {
	// Indicates whether logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-cloudwatchlogs.html#cfn-ec2-verifiedaccessinstance-cloudwatchlogs-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the CloudWatch Logs log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-cloudwatchlogs.html#cfn-ec2-verifiedaccessinstance-cloudwatchlogs-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

