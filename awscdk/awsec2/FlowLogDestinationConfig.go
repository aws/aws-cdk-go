package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//   var deliveryStreamRef IDeliveryStreamRef
//   var logGroupRef ILogGroupRef
//   var role Role
//
//   flowLogDestinationConfig := &FlowLogDestinationConfig{
//   	LogDestinationType: awscdk.Aws_ec2.FlowLogDestinationType_CLOUD_WATCH_LOGS,
//
//   	// the properties below are optional
//   	DeliveryStream: deliveryStreamRef,
//   	DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	DestinationOptions: &DestinationOptions{
//   		FileFormat: awscdk.*Aws_ec2.FlowLogFileFormat_PLAIN_TEXT,
//   		HiveCompatiblePartitions: jsii.Boolean(false),
//   		PerHourPartition: jsii.Boolean(false),
//   	},
//   	IamRole: role,
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	LogGroup: logGroupRef,
//   	S3Bucket: bucket,
//   }
//
type FlowLogDestinationConfig struct {
	// The type of destination to publish the flow logs to.
	// Default: - CLOUD_WATCH_LOGS.
	//
	LogDestinationType FlowLogDestinationType `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// The Amazon Data Firehose delivery stream to publish the flow logs to.
	// Default: - undefined.
	//
	DeliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// The ARN of Amazon Data Firehose delivery stream to publish the flow logs to.
	// Default: - undefined.
	//
	// Deprecated: use deliveryStream.
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// Options for writing flow logs to a supported destination.
	// Default: - undefined.
	//
	DestinationOptions *DestinationOptions `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// The IAM role that allows Amazon EC2 to publish flow logs to the log destination.
	//
	// Required if the destination type is CloudWatch logs, or if the destination type is Amazon Data Firehose delivery stream and the delivery stream and the VPC are in different accounts.
	// Default: - default IAM role is created for you if the destination type is CloudWatch logs.
	//
	IamRole awsiam.IRole `field:"optional" json:"iamRole" yaml:"iamRole"`
	// S3 bucket key prefix to publish the flow logs to.
	// Default: - undefined.
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The CloudWatch Logs Log Group to publish the flow logs to.
	// Default: - default log group is created for you.
	//
	LogGroup interfacesawslogs.ILogGroupRef `field:"optional" json:"logGroup" yaml:"logGroup"`
	// S3 bucket to publish the flow logs to.
	// Default: - undefined.
	//
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

