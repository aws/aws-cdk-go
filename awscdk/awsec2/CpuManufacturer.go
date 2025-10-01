package awsec2


// CPU manufacturers supported by EC2 instances.
//
// Restricts the acceptable CPU vendor for selected instance types.
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
type CpuManufacturer string

const (
	// Intel CPUs (e.g., Xeon families).
	CpuManufacturer_INTEL CpuManufacturer = "INTEL"
	// AMD CPUs (e.g., EPYC families).
	CpuManufacturer_AMD CpuManufacturer = "AMD"
	// AWS-designed CPUs (e.g., Graviton families).
	CpuManufacturer_AWS CpuManufacturer = "AWS"
	// Apple CPUs (e.g., M1, M2).
	CpuManufacturer_APPLE CpuManufacturer = "APPLE"
)

