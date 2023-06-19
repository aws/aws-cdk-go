package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for a LogRetention.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var role role
//
//   logRetentionProps := &LogRetentionProps{
//   	LogGroupName: jsii.String("logGroupName"),
//   	Retention: awscdk.Aws_logs.RetentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	LogGroupRegion: jsii.String("logGroupRegion"),
//   	LogRetentionRetryOptions: &LogRetentionRetryOptions{
//   		Base: duration,
//   		MaxRetries: jsii.Number(123),
//   	},
//   	Role: role,
//   }
//
// Experimental.
type LogRetentionProps struct {
	// The log group name.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The number of days log events are kept in CloudWatch Logs.
	// Experimental.
	Retention RetentionDays `field:"required" json:"retention" yaml:"retention"`
	// The region where the log group should be created.
	// Experimental.
	LogGroupRegion *string `field:"optional" json:"logGroupRegion" yaml:"logGroupRegion"`
	// Retry options for all AWS API calls.
	// Experimental.
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

