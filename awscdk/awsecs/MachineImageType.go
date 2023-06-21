package awsecs


// The machine image type.
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
type MachineImageType string

const (
	// Amazon ECS-optimized Amazon Linux 2 AMI.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

