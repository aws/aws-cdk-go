package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Flow Log Destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var logGroup logGroup
//   var role role
//
//   flowLogDestinationConfig := &flowLogDestinationConfig{
//   	logDestinationType: awscdk.Aws_ec2.flowLogDestinationType_CLOUD_WATCH_LOGS,
//
//   	// the properties below are optional
//   	destinationOptions: &destinationOptions{
//   		fileFormat: awscdk.*Aws_ec2.flowLogFileFormat_PLAIN_TEXT,
//   		hiveCompatiblePartitions: jsii.Boolean(false),
//   		perHourPartition: jsii.Boolean(false),
//   	},
//   	iamRole: role,
//   	keyPrefix: jsii.String("keyPrefix"),
//   	logGroup: logGroup,
//   	s3Bucket: bucket,
//   }
//
type FlowLogDestinationConfig struct {
	// The type of destination to publish the flow logs to.
	LogDestinationType FlowLogDestinationType `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// Options for writing flow logs to a supported destination.
	DestinationOptions *DestinationOptions `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// The IAM Role that has access to publish to CloudWatch logs.
	IamRole awsiam.IRole `field:"optional" json:"iamRole" yaml:"iamRole"`
	// S3 bucket key prefix to publish the flow logs to.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The CloudWatch Logs Log Group to publish the flow logs to.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// S3 bucket to publish the flow logs to.
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

