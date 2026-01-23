package awsec2


// Specific hardware accelerator models supported by EC2.
//
// Defines exact accelerator models that can be required or excluded
// when selecting instance types.
//
// Example:
//   var vpc Vpc
//
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   	Description: jsii.String("Security group for managed instances"),
//   })
//
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
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
	// NVIDIA A10G GPU.
	AcceleratorName_A10G AcceleratorName = "A10G"
	// NVIDIA H100 GPU.
	AcceleratorName_H100 AcceleratorName = "H100"
	// AWS Inferentia chips.
	AcceleratorName_INFERENTIA AcceleratorName = "INFERENTIA"
	// NVIDIA GRID K520 GPU.
	AcceleratorName_K520 AcceleratorName = "K520"
	// NVIDIA T4G GPUs.
	AcceleratorName_T4G AcceleratorName = "T4G"
	// NVIDIA L40S GPU for AI inference and graphics workloads.
	AcceleratorName_L40S AcceleratorName = "L40S"
	// NVIDIA L4 GPU for AI inference and graphics workloads.
	AcceleratorName_L4 AcceleratorName = "L4"
	// Habana Gaudi HL-205 accelerator for deep learning training.
	AcceleratorName_GAUDI_HL_205 AcceleratorName = "GAUDI_HL_205"
	// AWS Inferentia2 chips for high-performance ML inference.
	AcceleratorName_INFERENTIA2 AcceleratorName = "INFERENTIA2"
	// AWS Trainium chips for high-performance ML training.
	AcceleratorName_TRAINIUM AcceleratorName = "TRAINIUM"
	// AWS Trainium2 chips for high-performance ML training.
	AcceleratorName_TRAINIUM2 AcceleratorName = "TRAINIUM2"
	// Xilinx U30 media transcoding accelerator for video processing.
	AcceleratorName_U30 AcceleratorName = "U30"
)

