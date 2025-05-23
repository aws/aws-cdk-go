package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for BottleRocketImage.
//
// Example:
//   var cluster cluster
//
//
//   cluster.AddCapacity(jsii.String("bottlerocket-asg"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("p3.2xlarge")),
//   	MachineImage: ecs.NewBottleRocketImage(&BottleRocketImageProps{
//   		Variant: ecs.BottlerocketEcsVariant_AWS_ECS_2_NVIDIA,
//   	}),
//   })
//
type BottleRocketImageProps struct {
	// The CPU architecture.
	// Default: - x86_64.
	//
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
	// Default: false.
	//
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// The Amazon ECS variant to use.
	// Default: - BottlerocketEcsVariant.AWS_ECS_1
	//
	Variant BottlerocketEcsVariant `field:"optional" json:"variant" yaml:"variant"`
}

