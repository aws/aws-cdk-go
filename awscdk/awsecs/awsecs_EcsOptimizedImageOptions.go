package awsecs


// Additional configuration properties for EcsOptimizedImage factory functions.
//
// Example:
//   var vpc vpc
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux(&ecsOptimizedImageOptions{
//   		cachedInContext: jsii.Boolean(true),
//   	}),
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   })
//
type EcsOptimizedImageOptions struct {
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
}

