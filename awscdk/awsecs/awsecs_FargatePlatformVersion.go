package awsecs


// The platform version on which to run your service.
//
// Example:
//   var cluster cluster
//
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
//   	cluster: cluster,
//   	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		memoryLimitMiB: jsii.Number(512),
//   	},
//   	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
//   	platformVersion: ecs.fargatePlatformVersion_LATEST,
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
//
// Experimental.
type FargatePlatformVersion string

const (
	// The latest, recommended platform version.
	// Experimental.
	FargatePlatformVersion_LATEST FargatePlatformVersion = "LATEST"
	// Version 1.4.0.
	//
	// Supports EFS endpoints, CAP_SYS_PTRACE Linux capability,
	// network performance metrics in CloudWatch Container Insights,
	// consolidated 20 GB ephemeral volume.
	// Experimental.
	FargatePlatformVersion_VERSION1_4 FargatePlatformVersion = "VERSION1_4"
	// Version 1.3.0.
	//
	// Supports secrets, task recycling.
	// Experimental.
	FargatePlatformVersion_VERSION1_3 FargatePlatformVersion = "VERSION1_3"
	// Version 1.2.0.
	//
	// Supports private registries.
	// Experimental.
	FargatePlatformVersion_VERSION1_2 FargatePlatformVersion = "VERSION1_2"
	// Version 1.1.0.
	//
	// Supports task metadata, health checks, service discovery.
	// Experimental.
	FargatePlatformVersion_VERSION1_1 FargatePlatformVersion = "VERSION1_1"
	// Initial release.
	//
	// Based on Amazon Linux 2017.09.
	// Experimental.
	FargatePlatformVersion_VERSION1_0 FargatePlatformVersion = "VERSION1_0"
)

