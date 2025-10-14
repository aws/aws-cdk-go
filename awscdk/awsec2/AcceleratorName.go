package awsec2


// Specific hardware accelerator models supported by EC2.
//
// Defines exact accelerator models that can be required or excluded
// when selecting instance types.
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
type AcceleratorName string

const (
	// NVIDIA A100 GPU.
	AcceleratorName_A100 AcceleratorName = "A100"
	// NVIDIA K80 GPU.
	AcceleratorName_K80 AcceleratorName = "K80"
	// NVIDIA M60 GPU.
	AcceleratorName_M60 AcceleratorName = "M60"
	// AMD Radeon Pro V520 GPU.
	AcceleratorName_RADEON_PRO_V520 AcceleratorName = "RADEON_PRO_V520"
	// NVIDIA T4 GPU.
	AcceleratorName_T4 AcceleratorName = "T4"
	// NVIDIA V100 GPU.
	AcceleratorName_V100 AcceleratorName = "V100"
	// Xilinx VU9P FPGA.
	AcceleratorName_VU9P AcceleratorName = "VU9P"
)

