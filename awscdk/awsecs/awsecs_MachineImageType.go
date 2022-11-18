package awsecs


// The machine image type.
//
// Example:
//   var vpc vpc
//
//   launchTemplate := ec2.NewLaunchTemplate(this, jsii.String("ASG-LaunchTemplate"), &launchTemplateProps{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.medium")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
//   	userData: ec2.userData.forLinux(),
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	mixedInstancesPolicy: &mixedInstancesPolicy{
//   		instancesDistribution: &instancesDistribution{
//   			onDemandPercentageAboveBaseCapacity: jsii.Number(50),
//   		},
//   		launchTemplate: launchTemplate,
//   	},
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   	machineImageType: ecs.machineImageType_AMAZON_LINUX_2,
//   })
//
//   cluster.addAsgCapacityProvider(capacityProvider)
//
type MachineImageType string

const (
	// Amazon ECS-optimized Amazon Linux 2 AMI.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

