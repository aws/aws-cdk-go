package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// The properties that define which ECS-optimized AMI is used.
//
// Example:
//   var vpc vpc
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
//   	ComputeResources: &ComputeResources{
//   		Image: ecs.NewEcsOptimizedAmi(&EcsOptimizedAmiProps{
//   			Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		Vpc: *Vpc,
//   	},
//   })
//
// Deprecated: see {@link EcsOptimizedImage}.
type EcsOptimizedAmiProps struct {
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
	// Deprecated: see {@link EcsOptimizedImage}.
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// The Amazon Linux generation to use.
	// Deprecated: see {@link EcsOptimizedImage}.
	Generation awsec2.AmazonLinuxGeneration `field:"optional" json:"generation" yaml:"generation"`
	// The ECS-optimized AMI variant to use.
	// Deprecated: see {@link EcsOptimizedImage}.
	HardwareType AmiHardwareType `field:"optional" json:"hardwareType" yaml:"hardwareType"`
	// The Windows Server version to use.
	// Deprecated: see {@link EcsOptimizedImage}.
	WindowsVersion WindowsOptimizedVersion `field:"optional" json:"windowsVersion" yaml:"windowsVersion"`
}

