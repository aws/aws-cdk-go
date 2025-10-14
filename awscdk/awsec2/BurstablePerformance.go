package awsec2


// Burstable CPU performance requirements for EC2 instances.
//
// Controls whether selected instance types must, may, or must not
// support burstable vCPU performance (e.g., T3, T4g families).
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
type BurstablePerformance string

const (
	// Burstable-performance instance types are allowed, but non-burstable types may also be selected.
	BurstablePerformance_INCLUDED BurstablePerformance = "INCLUDED"
	// Only burstable-performance instance types are allowed.
	//
	// Non-burstable types will be excluded from selection.
	BurstablePerformance_REQUIRED BurstablePerformance = "REQUIRED"
	// Burstable-performance instance types are disallowed.
	//
	// Only non-burstable types may be selected.
	BurstablePerformance_EXCLUDED BurstablePerformance = "EXCLUDED"
)

