package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// The options for creating an Auto Scaling Group Capacity Provider.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
//   	minCapacity: jsii.Number(0),
//   	maxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   })
//   cluster.addAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: capacityProvider.capacityProviderName,
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type AsgCapacityProviderProps struct {
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
	// The autoscaling group to add as a Capacity Provider.
	// Experimental.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// The name of the capacity provider.
	//
	// If a name is specified,
	// it cannot start with `aws`, `ecs`, or `fargate`. If no name is specified,
	// a default name in the CFNStackName-CFNResourceName-RandomString format is used.
	// Experimental.
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// Whether to enable managed scaling.
	// Experimental.
	EnableManagedScaling *bool `field:"optional" json:"enableManagedScaling" yaml:"enableManagedScaling"`
	// Whether to enable managed termination protection.
	// Experimental.
	EnableManagedTerminationProtection *bool `field:"optional" json:"enableManagedTerminationProtection" yaml:"enableManagedTerminationProtection"`
	// Maximum scaling step size.
	//
	// In most cases this should be left alone.
	// Experimental.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// Minimum scaling step size.
	//
	// In most cases this should be left alone.
	// Experimental.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Target capacity percent.
	//
	// In most cases this should be left alone.
	// Experimental.
	TargetCapacityPercent *float64 `field:"optional" json:"targetCapacityPercent" yaml:"targetCapacityPercent"`
}

