package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
)

// The properties for the ScheduledEc2Task task.
//
// Example:
//   // Instantiate an Amazon EC2 Task to run at a scheduled interval
//   var cluster cluster
//
//   ecsScheduledTask := ecsPatterns.NewScheduledEc2Task(this, jsii.String("ScheduledTask"), &ScheduledEc2TaskProps{
//   	Cluster: Cluster,
//   	ScheduledEc2TaskImageOptions: &ScheduledEc2TaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		MemoryLimitMiB: jsii.Number(256),
//   		Environment: map[string]*string{
//   			"name": jsii.String("TRIGGER"),
//   			"value": jsii.String("CloudWatch Events"),
//   		},
//   	},
//   	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
//   	Enabled: jsii.Boolean(true),
//   	RuleName: jsii.String("sample-scheduled-task-rule"),
//   })
//
type ScheduledEc2TaskProps struct {
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
	// The properties to define if using an existing TaskDefinition in this construct.
	//
	// ScheduledEc2TaskDefinitionOptions or ScheduledEc2TaskImageOptions must be defined, but not both.
	// Default: none.
	//
	ScheduledEc2TaskDefinitionOptions *ScheduledEc2TaskDefinitionOptions `field:"optional" json:"scheduledEc2TaskDefinitionOptions" yaml:"scheduledEc2TaskDefinitionOptions"`
	// The properties to define if the construct is to create a TaskDefinition.
	//
	// ScheduledEc2TaskDefinitionOptions or ScheduledEc2TaskImageOptions must be defined, but not both.
	// Default: none.
	//
	ScheduledEc2TaskImageOptions *ScheduledEc2TaskImageOptions `field:"optional" json:"scheduledEc2TaskImageOptions" yaml:"scheduledEc2TaskImageOptions"`
}

