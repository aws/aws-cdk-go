package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Basic properties of both Application and Network Target Groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var vpc vpc
//
//   baseTargetGroupProps := &BaseTargetGroupProps{
//   	DeregistrationDelay: duration,
//   	HealthCheck: &HealthCheck{
//   		Enabled: jsii.Boolean(false),
//   		HealthyGrpcCodes: jsii.String("healthyGrpcCodes"),
//   		HealthyHttpCodes: jsii.String("healthyHttpCodes"),
//   		HealthyThresholdCount: jsii.Number(123),
//   		Interval: duration,
//   		Path: jsii.String("path"),
//   		Port: jsii.String("port"),
//   		Protocol: awscdk.Aws_elasticloadbalancingv2.Protocol_HTTP,
//   		Timeout: duration,
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   	TargetGroupName: jsii.String("targetGroupName"),
//   	TargetType: awscdk.*Aws_elasticloadbalancingv2.TargetType_INSTANCE,
//   	Vpc: vpc,
//   }
//
// Experimental.
type BaseTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Experimental.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Experimental.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Experimental.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	// Experimental.
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

