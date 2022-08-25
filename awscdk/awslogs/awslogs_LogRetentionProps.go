package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a LogRetention.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logRetentionProps := &logRetentionProps{
//   	logGroupName: jsii.String("logGroupName"),
//   	retention: awscdk.Aws_logs.retentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	logGroupRegion: jsii.String("logGroupRegion"),
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: cdk.duration.minutes(jsii.Number(30)),
//   		maxRetries: jsii.Number(123),
//   	},
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	role: role,
//   }
//
type LogRetentionProps struct {
	// The log group name.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The number of days log events are kept in CloudWatch Logs.
	Retention RetentionDays `field:"required" json:"retention" yaml:"retention"`
	// The region where the log group should be created.
	LogGroupRegion *string `field:"optional" json:"logGroupRegion" yaml:"logGroupRegion"`
	// Retry options for all AWS API calls.
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The removalPolicy for the log group when the stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The IAM role for the Lambda function associated with the custom resource.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

