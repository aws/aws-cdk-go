package awsec2


// Options for writing logs to a S3 destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationOptions := &S3DestinationOptions{
//   	FileFormat: awscdk.Aws_ec2.FlowLogFileFormat_PLAIN_TEXT,
//   	HiveCompatiblePartitions: jsii.Boolean(false),
//   	PerHourPartition: jsii.Boolean(false),
//   }
//
type S3DestinationOptions struct {
	// The format for the flow log.
	FileFormat FlowLogFileFormat `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Use Hive-compatible prefixes for flow logs stored in Amazon S3.
	HiveCompatiblePartitions *bool `field:"optional" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Partition the flow log per hour.
	PerHourPartition *bool `field:"optional" json:"perHourPartition" yaml:"perHourPartition"`
}

