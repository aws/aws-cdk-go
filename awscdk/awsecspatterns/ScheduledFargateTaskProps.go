package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
)

// The properties for the ScheduledFargateTask task.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//   cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
//   	Cluster: Cluster,
//   	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		MemoryLimitMiB: jsii.Number(512),
//   	},
//   	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
//   	Tags: []tag{
//   		&tag{
//   			Key: jsii.String("my-tag"),
//   			Value: jsii.String("my-tag-value"),
//   		},
//   	},
//   })
//
type ScheduledFargateTaskProps struct {
	// The schedule or rate (frequency) that determines when CloudWatch Events runs the rule.
	//
	// For more information, see
	// [Schedule Expression Syntax for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch User Guide.
	Schedule awsapplicationautoscaling.Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Default: - create a new cluster; if both cluster and vpc are omitted, a new VPC will be created for you.
	//
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Default: 1.
	//
	DesiredTaskCount *float64 `field:"optional" json:"desiredTaskCount" yaml:"desiredTaskCount"`
	// Indicates whether the rule is enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated.
	// Default: - Tags will not be propagated.
	//
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A name for the rule.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that ID
	// for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Existing security groups to use for your service.
	// Default: - a new security group will be created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	// Default: Private subnets.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	// Default: - No tags are applied to the task.
	//
	Tags *[]*awseventstargets.Tag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// 8192 (8 vCPU) - Available memory values: Between 16GB and 60GB in 4GB increments
	//
	// 16384 (16 vCPU) - Available memory values: Between 32GB and 120GB in 8GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 256.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	// Default: Undefined, in which case, the task will receive 20GiB ephemeral storage.
	//
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB) - Available cpu values: 8192 (8 vCPU)
	//
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 512.
	//
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Default: Latest.
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	// Default: - If the property is undefined, `operatingSystemFamily` is LINUX and `cpuArchitecture` is X86_64.
	//
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - none.
	//
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// The properties to define if using an existing TaskDefinition in this construct.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Default: none.
	//
	ScheduledFargateTaskDefinitionOptions *ScheduledFargateTaskDefinitionOptions `field:"optional" json:"scheduledFargateTaskDefinitionOptions" yaml:"scheduledFargateTaskDefinitionOptions"`
	// The properties to define if the construct is to create a TaskDefinition.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Default: none.
	//
	ScheduledFargateTaskImageOptions *ScheduledFargateTaskImageOptions `field:"optional" json:"scheduledFargateTaskImageOptions" yaml:"scheduledFargateTaskImageOptions"`
}

