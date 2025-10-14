package awsec2


// Hardware accelerator categories available for EC2 instances.
//
// Defines the general type of hardware accelerator that can be attached
// to an instance, typically used in instance requirement specifications
// (e.g., GPUs for compute-intensive tasks, FPGAs for custom logic, or
// inference chips for ML workloads).
//
// Example:
//   var infrastructureRole role
//   var instanceProfile instanceProfile
//   var vpc vpc
//
//
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	InfrastructureRole: InfrastructureRole,
//   	Ec2InstanceProfile: instanceProfile,
//   	Subnets: vpc.PrivateSubnets,
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		// Required: CPU and memory constraints
//   		VCpuCountMin: jsii.Number(2),
//   		VCpuCountMax: jsii.Number(8),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(4)),
//   		MemoryMax: awscdk.Size_*Gibibytes(jsii.Number(32)),
//
//   		// CPU preferences
//   		CpuManufacturers: []cpuManufacturer{
//   			ec2.*cpuManufacturer_INTEL,
//   			ec2.*cpuManufacturer_AMD,
//   		},
//   		InstanceGenerations: []instanceGeneration{
//   			ec2.*instanceGeneration_CURRENT,
//   		},
//
//   		// Instance type filtering
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("m5.*"),
//   			jsii.String("c5.*"),
//   		},
//
//   		// Performance characteristics
//   		BurstablePerformance: ec2.BurstablePerformance_EXCLUDED,
//   		BareMetal: ec2.BareMetal_EXCLUDED,
//
//   		// Accelerator requirements (for ML/AI workloads)
//   		AcceleratorTypes: []acceleratorType{
//   			ec2.*acceleratorType_GPU,
//   		},
//   		AcceleratorManufacturers: []acceleratorManufacturer{
//   			ec2.*acceleratorManufacturer_NVIDIA,
//   		},
//   		AcceleratorNames: []acceleratorName{
//   			ec2.*acceleratorName_T4,
//   			ec2.*acceleratorName_V100,
//   		},
//   		AcceleratorCountMin: jsii.Number(1),
//
//   		// Storage requirements
//   		LocalStorage: ec2.LocalStorage_REQUIRED,
//   		LocalStorageTypes: []localStorageType{
//   			ec2.*localStorageType_SSD,
//   		},
//   		TotalLocalStorageGBMin: jsii.Number(100),
//
//   		// Network requirements
//   		NetworkInterfaceCountMin: jsii.Number(2),
//   		NetworkBandwidthGbpsMin: jsii.Number(10),
//
//   		// Cost optimization
//   		OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(10),
//   	},
//   })
//
type AcceleratorType string

const (
	// Graphics Processing Unit accelerators, such as NVIDIA GPUs.
	//
	// Commonly used for machine learning training, graphics rendering,
	// or high-performance parallel computing.
	AcceleratorType_GPU AcceleratorType = "GPU"
	// Field Programmable Gate Array accelerators, such as Xilinx FPGAs.
	//
	// Used for hardware-level customization and specialized workloads.
	AcceleratorType_FPGA AcceleratorType = "FPGA"
	// Inference accelerators, such as AWS Inferentia.
	//
	// Purpose-built for efficient machine learning inference.
	AcceleratorType_INFERENCE AcceleratorType = "INFERENCE"
)

