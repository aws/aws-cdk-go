package awsec2


// CPU manufacturers supported by EC2 instances.
//
// Restricts the acceptable CPU vendor for selected instance types.
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

