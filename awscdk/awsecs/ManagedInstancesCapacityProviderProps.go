package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The options for creating a Managed Instances Capacity Provider.
//
// Example:
//   var vpc vpc
//   var infrastructureRole role
//   var instanceProfile instanceProfile
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Create a Managed Instances Capacity Provider
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	InfrastructureRole: InfrastructureRole,
//   	Ec2InstanceProfile: instanceProfile,
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []iSecurityGroup{
//   		ec2.NewSecurityGroup(this, jsii.String("MISecurityGroup"), &SecurityGroupProps{
//   			Vpc: *Vpc,
//   		}),
//   	},
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		VCpuCountMin: jsii.Number(1),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(2)),
//   		CpuManufacturers: []cpuManufacturer{
//   			ec2.*cpuManufacturer_INTEL,
//   		},
//   		AcceleratorManufacturers: []acceleratorManufacturer{
//   			ec2.*acceleratorManufacturer_NVIDIA,
//   		},
//   	},
//   	PropagateTags: ecs.PropagateManagedInstancesTags_CAPACITY_PROVIDER,
//   })
//
//   // Add the capacity provider to the cluster
//   cluster.AddManagedInstancesCapacityProvider(miCapacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: miCapacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type ManagedInstancesCapacityProviderProps struct {
	// The EC2 instance profile that will be attached to instances launched by this capacity provider.
	//
	// This instance profile must contain the necessary IAM permissions for ECS container instances
	// to register with the cluster and run tasks. At minimum, it should include permissions for
	// ECS agent communication, ECR image pulling, and CloudWatch logging.
	Ec2InstanceProfile awsiam.IInstanceProfile `field:"required" json:"ec2InstanceProfile" yaml:"ec2InstanceProfile"`
	// The VPC subnets where EC2 instances will be launched.
	//
	// This array must be non-empty and should contain subnets from the VPC where you want
	// the managed instances to be deployed.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws`, `ecs`, or `fargate`.
	// If no name is specified, a default name in the CFNStackName-CFNResourceName-RandomString format is used.
	// If the stack name starts with `aws`, `ecs`, or `fargate`, a unique resource name
	// is generated that starts with `cp-`.
	// Default: CloudFormation-generated name.
	//
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// The IAM role that ECS uses to manage the infrastructure for the capacity provider.
	//
	// This role is used by ECS to perform actions such as launching and terminating instances,
	// managing Auto Scaling Groups, and other infrastructure operations required for the
	// managed instances capacity provider.
	// Default: - A new role will be created with the AmazonECSInfrastructureRolePolicyForManagedInstances managed policy.
	//
	InfrastructureRole awsiam.IRole `field:"optional" json:"infrastructureRole" yaml:"infrastructureRole"`
	// The instance requirements configuration for EC2 instance selection.
	//
	// This allows you to specify detailed requirements for instance selection including
	// vCPU count ranges, memory ranges, CPU manufacturers (Intel, AMD, AWS Graviton),
	// instance generations, network performance requirements, and many other criteria.
	// ECS will automatically select appropriate instance types that meet these requirements.
	// Default: - no specific instance requirements, ECS will choose appropriate instances.
	//
	InstanceRequirements *awsec2.InstanceRequirementsConfig `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The CloudWatch monitoring configuration for the EC2 instances.
	//
	// Determines the granularity of CloudWatch metrics collection for the instances.
	// Detailed monitoring incurs additional costs but provides better observability.
	// Default: - no enhanced monitoring (basic monitoring only).
	//
	Monitoring InstanceMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Specifies whether to propagate tags from the capacity provider to the launched instances.
	//
	// When set to CAPACITY_PROVIDER, tags applied to the capacity provider resource will be
	// automatically applied to all EC2 instances launched by this capacity provider.
	// Default: PropagateManagedInstancesTags.NONE - no tag propagation
	//
	PropagateTags PropagateManagedInstancesTags `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The security groups to associate with the launched EC2 instances.
	//
	// These security groups control the network traffic allowed to and from the instances.
	// If not specified, the default security group of the VPC containing the subnets will be used.
	// Default: - default security group of the VPC.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The size of the task volume storage attached to each instance.
	//
	// This storage is used for container images, container logs, and temporary files.
	// Larger storage may be needed for workloads with large container images or
	// applications that generate significant temporary data.
	// Default: Size.gibibytes(80)
	//
	TaskVolumeStorage awscdk.Size `field:"optional" json:"taskVolumeStorage" yaml:"taskVolumeStorage"`
}

