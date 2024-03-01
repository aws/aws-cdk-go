package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The options for creating an Auto Scaling Group Capacity Provider.
//
// Example:
//   var vpc vpc
//
//   launchTemplate := ec2.NewLaunchTemplate(this, jsii.String("ASG-LaunchTemplate"), &LaunchTemplateProps{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.medium")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
//   	UserData: ec2.UserData_ForLinux(),
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	MixedInstancesPolicy: &MixedInstancesPolicy{
//   		InstancesDistribution: &InstancesDistribution{
//   			OnDemandPercentageAboveBaseCapacity: jsii.Number(50),
//   		},
//   		LaunchTemplate: launchTemplate,
//   	},
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
//   	AutoScalingGroup: AutoScalingGroup,
//   	MachineImageType: ecs.MachineImageType_AMAZON_LINUX_2,
//   })
//
//   cluster.AddAsgCapacityProvider(capacityProvider)
//
type AsgCapacityProviderProps struct {
	// Specifies whether the containers can access the container instance role.
	// Default: false.
	//
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
	// Default: - Automatically determined from `machineImage`, if available, otherwise `MachineImageType.AMAZON_LINUX_2`.
	//
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	// Default: false.
	//
	SpotInstanceDraining *bool `field:"optional" json:"spotInstanceDraining" yaml:"spotInstanceDraining"`
	// If `AddAutoScalingGroupCapacityOptions.taskDrainTime` is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	// Default: The SNS Topic will not be encrypted.
	//
	TopicEncryptionKey awskms.IKey `field:"optional" json:"topicEncryptionKey" yaml:"topicEncryptionKey"`
	// The autoscaling group to add as a Capacity Provider.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// The name of the capacity provider.
	//
	// If a name is specified,
	// it cannot start with `aws`, `ecs`, or `fargate`. If no name is specified,
	// a default name in the CFNStackName-CFNResourceName-RandomString format is used.
	// If the stack name starts with `aws`, `ecs`, or `fargate`, a unique resource name
	// is generated that starts with `cp-`.
	// Default: CloudFormation-generated name.
	//
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// Managed instance draining facilitates graceful termination of Amazon ECS instances.
	//
	// This allows your service workloads to stop safely and be rescheduled to non-terminating instances.
	// Infrastructure maintenance and updates are preformed without disruptions to workloads.
	// To use managed instance draining, set enableManagedDraining to true.
	// Default: true.
	//
	EnableManagedDraining *bool `field:"optional" json:"enableManagedDraining" yaml:"enableManagedDraining"`
	// When enabled the scale-in and scale-out actions of the cluster's Auto Scaling Group will be managed for you.
	//
	// This means your cluster will automatically scale instances based on the load your tasks put on the cluster.
	// For more information, see [Using Managed Scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling) in the ECS Developer Guide.
	// Default: true.
	//
	EnableManagedScaling *bool `field:"optional" json:"enableManagedScaling" yaml:"enableManagedScaling"`
	// When enabled the Auto Scaling Group will only terminate EC2 instances that no longer have running non-daemon tasks.
	//
	// Scale-in protection will be automatically enabled on instances. When all non-daemon tasks are
	// stopped on an instance, ECS initiates the scale-in process and turns off scale-in protection for the
	// instance. The Auto Scaling Group can then terminate the instance. For more information see [Managed termination
	//  protection](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-auto-scaling.html#managed-termination-protection)
	// in the ECS Developer Guide.
	//
	// Managed scaling must also be enabled.
	// Default: true.
	//
	EnableManagedTerminationProtection *bool `field:"optional" json:"enableManagedTerminationProtection" yaml:"enableManagedTerminationProtection"`
	// The period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group.
	//
	// Must be between 0 and 10000.
	// Default: 300.
	//
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// Maximum scaling step size.
	//
	// In most cases this should be left alone.
	// Default: 1000.
	//
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// Minimum scaling step size.
	//
	// In most cases this should be left alone.
	// Default: 1.
	//
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Target capacity percent.
	//
	// In most cases this should be left alone.
	// Default: 100.
	//
	TargetCapacityPercent *float64 `field:"optional" json:"targetCapacityPercent" yaml:"targetCapacityPercent"`
}

