package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// The properties for adding an AutoScalingGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var key key
//
//   addAutoScalingGroupCapacityOptions := &addAutoScalingGroupCapacityOptions{
//   	canContainersAccessInstanceRole: jsii.Boolean(false),
//   	machineImageType: awscdk.Aws_ecs.machineImageType_AMAZON_LINUX_2,
//   	spotInstanceDraining: jsii.Boolean(false),
//   	taskDrainTime: duration,
//   	topicEncryptionKey: key,
//   }
//
// Experimental.
type AddAutoScalingGroupCapacityOptions struct {
	// Specifies whether the containers can access the container instance role.
	// Experimental.
	CanContainersAccessInstanceRole *bool `field:"optional" json:"canContainersAccessInstanceRole" yaml:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	// Experimental.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	// Experimental.
	SpotInstanceDraining *bool `field:"optional" json:"spotInstanceDraining" yaml:"spotInstanceDraining"`
	// The time period to wait before force terminating an instance that is draining.
	//
	// This creates a Lambda function that is used by a lifecycle hook for the
	// AutoScalingGroup that will delay instance termination until all ECS tasks
	// have drained from the instance. Set to 0 to disable task draining.
	//
	// Set to 0 to disable task draining.
	// Deprecated: The lifecycle draining hook is not configured if using the EC2 Capacity Provider. Enable managed termination protection instead.
	TaskDrainTime awscdk.Duration `field:"optional" json:"taskDrainTime" yaml:"taskDrainTime"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	// Experimental.
	TopicEncryptionKey awskms.IKey `field:"optional" json:"topicEncryptionKey" yaml:"topicEncryptionKey"`
}

