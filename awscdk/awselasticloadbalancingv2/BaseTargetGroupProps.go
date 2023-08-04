package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Basic properties of both Application and Network Target Groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   baseTargetGroupProps := &BaseTargetGroupProps{
//   	DeregistrationDelay: cdk.Duration_Minutes(jsii.Number(30)),
//   	HealthCheck: &HealthCheck{
//   		Enabled: jsii.Boolean(false),
//   		HealthyGrpcCodes: jsii.String("healthyGrpcCodes"),
//   		HealthyHttpCodes: jsii.String("healthyHttpCodes"),
//   		HealthyThresholdCount: jsii.Number(123),
//   		Interval: cdk.Duration_*Minutes(jsii.Number(30)),
//   		Path: jsii.String("path"),
//   		Port: jsii.String("port"),
//   		Protocol: awscdk.Aws_elasticloadbalancingv2.Protocol_HTTP,
//   		Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   	TargetGroupName: jsii.String("targetGroupName"),
//   	TargetType: awscdk.*Aws_elasticloadbalancingv2.TargetType_INSTANCE,
//   	Vpc: vpc,
//   }
//
type BaseTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Default: 300.
	//
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Default: - The default value for each property in this configuration varies depending on the target.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Default: - Automatically generated.
	//
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	// Default: - Determined automatically.
	//
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	// Default: - undefined.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

