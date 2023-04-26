package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
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
//   flowLogDestinationConfig := &FlowLogDestinationConfig{
//   	LogDestinationType: awscdk.Aws_ec2.FlowLogDestinationType_CLOUD_WATCH_LOGS,
//
//   	// the properties below are optional
//   	IamRole: role,
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	LogGroup: logGroup,
//   	S3Bucket: bucket,
//   }
//
// Experimental.
type FlowLogDestinationConfig struct {
	// The type of destination to publish the flow logs to.
	// Experimental.
	LogDestinationType FlowLogDestinationType `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// The IAM Role that has access to publish to CloudWatch logs.
	// Experimental.
	IamRole awsiam.IRole `field:"optional" json:"iamRole" yaml:"iamRole"`
	// S3 bucket key prefix to publish the flow logs to.
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The CloudWatch Logs Log Group to publish the flow logs to.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// S3 bucket to publish the flow logs to.
	// Experimental.
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

