package awsec2


// Supported hardware accelerator manufacturers.
//
// Restricts instance selection to accelerators from a particular vendor.
// Useful for choosing specific ecosystems (e.g., NVIDIA CUDA, AWS chips).
//
// Example:
//   var vpc Vpc
//
//
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		// Required: CPU and memory constraints
//   		VCpuCountMin: jsii.Number(2),
//   		VCpuCountMax: jsii.Number(8),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(4)),
//   		MemoryMax: awscdk.Size_*Gibibytes(jsii.Number(32)),
//
//   		// CPU preferences
//   		CpuManufacturers: []CpuManufacturer{
//   			ec2.CpuManufacturer_INTEL,
//   			ec2.CpuManufacturer_AMD,
//   		},
//   		InstanceGenerations: []InstanceGeneration{
//   			ec2.InstanceGeneration_CURRENT,
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
//   		AcceleratorTypes: []AcceleratorType{
//   			ec2.AcceleratorType_GPU,
//   		},
//   		AcceleratorManufacturers: []AcceleratorManufacturer{
//   			ec2.AcceleratorManufacturer_NVIDIA,
//   		},
//   		AcceleratorNames: []AcceleratorName{
//   			ec2.AcceleratorName_T4,
//   			ec2.AcceleratorName_V100,
//   		},
//   		AcceleratorCountMin: jsii.Number(1),
//
//   		// Storage requirements
//   		LocalStorage: ec2.LocalStorage_REQUIRED,
//   		LocalStorageTypes: []LocalStorageType{
//   			ec2.LocalStorageType_SSD,
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

