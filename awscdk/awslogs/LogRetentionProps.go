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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logRetentionProps := &LogRetentionProps{
//   	LogGroupName: jsii.String("logGroupName"),
//   	Retention: awscdk.Aws_logs.RetentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	LogGroupRegion: jsii.String("logGroupRegion"),
//   	LogRetentionRetryOptions: &LogRetentionRetryOptions{
//   		Base: cdk.Duration_Minutes(jsii.Number(30)),
//   		MaxRetries: jsii.Number(123),
//   	},
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	Role: role,
//   }
//
type LogRetentionProps struct {
	// The log group name.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The number of days log events are kept in CloudWatch Logs.
	Retention RetentionDays `field:"required" json:"retention" yaml:"retention"`
	// The region where the log group should be created.
	// Default: - same region as the stack.
	//
	LogGroupRegion *string `field:"optional" json:"logGroupRegion" yaml:"logGroupRegion"`
	// Retry options for all AWS API calls.
	// Default: - AWS SDK default retry options.
	//
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The removalPolicy for the log group when the stack is deleted.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The IAM role for the Lambda function associated with the custom resource.
	// Default: - A new role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

