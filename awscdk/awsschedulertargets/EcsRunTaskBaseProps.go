package awsschedulertargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Parameters for scheduling ECS Run Task (common to EC2 and Fargate launch types).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue Queue
//   var role Role
//   var scheduleTargetInput ScheduleTargetInput
//   var securityGroup SecurityGroup
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var taskDefinition TaskDefinition
//
//   ecsRunTaskBaseProps := &EcsRunTaskBaseProps{
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	CapacityProviderStrategies: []CapacityProviderStrategy{
//   		&CapacityProviderStrategy{
//   			CapacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			Base: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	DeadLetterQueue: queue,
//   	EnableEcsManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	Group: jsii.String("group"),
//   	Input: scheduleTargetInput,
//   	MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	PropagateTags: jsii.Boolean(false),
//   	ReferenceId: jsii.String("referenceId"),
//   	RetryAttempts: jsii.Number(123),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Tags: []Tag{
//   		&Tag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskCount: jsii.Number(123),
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []SubnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []ISubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type EcsRunTaskBaseProps struct {
	// The SQS queue to be used as deadLetterQueue.
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Input passed to the target.
	// Default: - no input.
	//
	Input awsscheduler.ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The maximum age of a request that Scheduler sends to a target for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the target returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// An execution role is an IAM role that EventBridge Scheduler assumes in order to interact with other AWS services on your behalf.
	//
	// If none provided templates target will automatically create an IAM role with all the minimum necessary
	// permissions to interact with the templated target. If you wish you may specify your own IAM role, then the templated targets
	// will grant minimal required permissions.
	// Default: - created by target.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The task definition to use for scheduled tasks.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// If you want to run a RunTask with an imported task definition,
	// consider using a Universal target.
	TaskDefinition awsecs.TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// The capacity provider strategy to use for the task.
	// Default: - No capacity provider strategy.
	//
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Specifies whether to enable Amazon ECS managed tags for the task.
	// Default: - false.
	//
	EnableEcsManagedTags *bool `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Whether to enable execute command functionality for the containers in this task.
	//
	// If true, this enables execute command functionality on all containers in the task.
	// Default: - false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies an ECS task group for the task.
	// Default: - No group.
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated.
	// Default: - No tag propagation.
	//
	PropagateTags *bool `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The reference ID to use for the task.
	// Default: - No reference ID.
	//
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// The security groups associated with the task.
	//
	// These security groups must all be in the same VPC.
	// Controls inbound and outbound network access for the task.
	// Default: - The security group for the VPC is used.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	// Default: - No tags.
	//
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
	// The number of tasks to create based on TaskDefinition.
	// Default: 1.
	//
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
	// The subnets associated with the task.
	//
	// These subnets must all be in the same VPC.
	// The task will be launched in these subnets.
	// Default: - all private subnets of the VPC are selected.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

