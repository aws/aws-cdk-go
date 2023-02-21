package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for BottleRocketImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bottleRocketImageProps := &BottleRocketImageProps{
//   	Architecture: awscdk.Aws_ec2.InstanceArchitecture_ARM_64,
//   	CachedInContext: jsii.Boolean(false),
//   	Variant: awscdk.Aws_ecs.BottlerocketEcsVariant_AWS_ECS_1,
//   }
//
type BottleRocketImageProps struct {
	// The CPU architecture.
	Architecture awsec2.InstanceArchitecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// The Amazon ECS variant to use.
	//
	// Only `aws-ecs-1` is currently available.
	Variant BottlerocketEcsVariant `field:"optional" json:"variant" yaml:"variant"`
}

