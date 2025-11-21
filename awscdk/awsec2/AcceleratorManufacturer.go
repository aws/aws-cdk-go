package awsec2


// Supported hardware accelerator manufacturers.
//
// Restricts instance selection to accelerators from a particular vendor.
// Useful for choosing specific ecosystems (e.g., NVIDIA CUDA, AWS chips).
//
// Example:
//   var vpc Vpc
//   var infrastructureRole Role
//   var instanceProfile InstanceProfile
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
//   	SecurityGroups: []ISecurityGroup{
//   		ec2.NewSecurityGroup(this, jsii.String("MISecurityGroup"), &SecurityGroupProps{
//   			Vpc: *Vpc,
//   		}),
//   	},
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		VCpuCountMin: jsii.Number(1),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(2)),
//   		CpuManufacturers: []CpuManufacturer{
//   			ec2.CpuManufacturer_INTEL,
//   		},
//   		AcceleratorManufacturers: []AcceleratorManufacturer{
//   			ec2.AcceleratorManufacturer_NVIDIA,
//   		},
//   	},
//   	PropagateTags: ecs.PropagateManagedInstancesTags_CAPACITY_PROVIDER,
//   })
//
//   // Optionally configure security group rules using IConnectable interface
//   miCapacityProvider.Connections.AllowFrom(ec2.Peer_Ipv4(vpc.VpcCidrBlock), ec2.Port_Tcp(jsii.Number(80)))
//
//   // Add the capacity provider to the cluster
//   cluster.AddManagedInstancesCapacityProvider(miCapacityProvider)
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &TaskDefinitionProps{
//   	MemoryMiB: jsii.String("512"),
//   	Cpu: jsii.String("256"),
//   	NetworkMode: ecs.NetworkMode_AWS_VPC,
//   	Compatibility: ecs.Compatibility_MANAGED_INSTANCES,
//   })
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	CapacityProviderStrategies: []CapacityProviderStrategy{
//   		&CapacityProviderStrategy{
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
	// Habana Labs(e.g, Gaudi accelerator).
	AcceleratorManufacturer_HABANA AcceleratorManufacturer = "HABANA"
)

