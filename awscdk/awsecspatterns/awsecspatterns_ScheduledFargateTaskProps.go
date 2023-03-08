package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// The properties for the ScheduledFargateTask task.
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
// Experimental.
type ScheduledFargateTaskProps struct {
	// The schedule or rate (frequency) that determines when CloudWatch Events runs the rule.
	//
	// For more information, see
	// [Schedule Expression Syntax for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch User Guide.
	// Experimental.
	Schedule awsapplicationautoscaling.Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredTaskCount *float64 `field:"optional" json:"desiredTaskCount" yaml:"desiredTaskCount"`
	// Indicates whether the rule is enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A name for the rule.
	// Experimental.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Existing security groups to use for your service.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The properties to define if using an existing TaskDefinition in this construct.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledFargateTaskDefinitionOptions *ScheduledFargateTaskDefinitionOptions `field:"optional" json:"scheduledFargateTaskDefinitionOptions" yaml:"scheduledFargateTaskDefinitionOptions"`
	// The properties to define if the construct is to create a TaskDefinition.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledFargateTaskImageOptions *ScheduledFargateTaskImageOptions `field:"optional" json:"scheduledFargateTaskImageOptions" yaml:"scheduledFargateTaskImageOptions"`
}

