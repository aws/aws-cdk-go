package awsec2


// Supported hardware accelerator manufacturers.
//
// Restricts instance selection to accelerators from a particular vendor.
// Useful for choosing specific ecosystems (e.g., NVIDIA CUDA, AWS chips).
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
type AcceleratorManufacturer string

const (
	// Amazon Web Services (e.g., Inferentia, Trainium accelerators).
	AcceleratorManufacturer_AWS AcceleratorManufacturer = "AWS"
	// AMD (e.g., Radeon Pro V520 GPU).
	AcceleratorManufacturer_AMD AcceleratorManufacturer = "AMD"
	// NVIDIA (e.g., A100, V100, T4, K80, M60 GPUs).
	AcceleratorManufacturer_NVIDIA AcceleratorManufacturer = "NVIDIA"
	// Xilinx (e.g., VU9P FPGA).
	AcceleratorManufacturer_XILINX AcceleratorManufacturer = "XILINX"
)

