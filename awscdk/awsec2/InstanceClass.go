package awsec2


// What class and generation of instance to use.
//
// We have both symbolic and concrete enums for every type.
//
// The first are for people that want to specify by purpose,
// the second one are for people who already know exactly what
// 'R4' means.
//
// Example:
//   // Creates a distribution from an EC2 instance
//   var vpc vpc
//
//   // Create an EC2 instance in a VPC. 'subnetType' can be private.
//   instance := ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.VpcOrigin_WithEc2Instance(instance),
//   	},
//   })
//
type InstanceClass string

const (
	// Standard instances, 3rd generation.
	InstanceClass_STANDARD3 InstanceClass = "STANDARD3"
	// Standard instances, 3rd generation.
	InstanceClass_M3 InstanceClass = "M3"
	// Standard instances, 4th generation.
	InstanceClass_STANDARD4 InstanceClass = "STANDARD4"
	// Standard instances, 4th generation.
	InstanceClass_M4 InstanceClass = "M4"
	// Standard instances, 5th generation.
	InstanceClass_STANDARD5 InstanceClass = "STANDARD5"
	// Standard instances, 5th generation.
	InstanceClass_M5 InstanceClass = "M5"
	// Standard instances with local NVME drive, 5th generation.
	InstanceClass_STANDARD5_NVME_DRIVE InstanceClass = "STANDARD5_NVME_DRIVE"
	// Standard instances with local NVME drive, 5th generation.
	InstanceClass_M5D InstanceClass = "M5D"
	// Standard instances based on AMD EPYC, 5th generation.
	InstanceClass_STANDARD5_AMD InstanceClass = "STANDARD5_AMD"
	// Standard instances based on AMD EPYC, 5th generation.
	InstanceClass_M5A InstanceClass = "M5A"
	// Standard instances based on AMD EPYC with local NVME drive, 5th generation.
	InstanceClass_STANDARD5_AMD_NVME_DRIVE InstanceClass = "STANDARD5_AMD_NVME_DRIVE"
	// Standard instances based on AMD EPYC with local NVME drive, 5th generation.
	InstanceClass_M5AD InstanceClass = "M5AD"
	// Standard instances for high performance computing, 5th generation.
	InstanceClass_STANDARD5_HIGH_PERFORMANCE InstanceClass = "STANDARD5_HIGH_PERFORMANCE"
	// Standard instances for high performance computing, 5th generation.
	InstanceClass_M5N InstanceClass = "M5N"
	// Standard instances with local NVME drive for high performance computing, 5th generation.
	InstanceClass_STANDARD5_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "STANDARD5_NVME_DRIVE_HIGH_PERFORMANCE"
	// Standard instances with local NVME drive for high performance computing, 5th generation.
	InstanceClass_M5DN InstanceClass = "M5DN"
	// Standard instances with high memory and compute capacity based on Intel Xeon Scalable (Cascade Lake) processors, 5nd generation.
	InstanceClass_STANDARD5_HIGH_COMPUTE InstanceClass = "STANDARD5_HIGH_COMPUTE"
	// Standard instances with high memory and compute capacity based on Intel Xeon Scalable (Cascade Lake) processors, 5nd generation.
	InstanceClass_M5ZN InstanceClass = "M5ZN"
	// Memory optimized instances, 3rd generation.
	InstanceClass_MEMORY3 InstanceClass = "MEMORY3"
	// Memory optimized instances, 3rd generation.
	InstanceClass_R3 InstanceClass = "R3"
	// Memory optimized instances, 4th generation.
	InstanceClass_MEMORY4 InstanceClass = "MEMORY4"
	// Memory optimized instances, 4th generation.
	InstanceClass_R4 InstanceClass = "R4"
	// Memory optimized instances, 5th generation.
	InstanceClass_MEMORY5 InstanceClass = "MEMORY5"
	// Memory optimized instances, 5th generation.
	InstanceClass_R5 InstanceClass = "R5"
	// Memory optimized instances based on AMD EPYC, 6th generation.
	InstanceClass_MEMORY6_AMD InstanceClass = "MEMORY6_AMD"
	// Memory optimized instances based on AMD EPYC, 6th generation.
	InstanceClass_R6A InstanceClass = "R6A"
	// Memory optimized instances, 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake).
	InstanceClass_MEMORY6_INTEL InstanceClass = "MEMORY6_INTEL"
	// Memory optimized instances, 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake).
	InstanceClass_R6I InstanceClass = "R6I"
	// Memory optimized instances with local NVME drive, 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake).
	InstanceClass_MEMORY6_INTEL_NVME_DRIVE InstanceClass = "MEMORY6_INTEL_NVME_DRIVE"
	// Memory optimized instances with local NVME drive, 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake).
	InstanceClass_R6ID InstanceClass = "R6ID"
	// Memory optimized instances for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_MEMORY6_INTEL_HIGH_PERFORMANCE InstanceClass = "MEMORY6_INTEL_HIGH_PERFORMANCE"
	// Memory optimized instances for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_R6IN InstanceClass = "R6IN"
	// Memory optimized instances with local NVME drive for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_MEMORY6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "MEMORY6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE"
	// Memory optimized instances with local NVME drive for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_R6IDN InstanceClass = "R6IDN"
	// Memory optimized instances for high performance computing, 5th generation.
	InstanceClass_MEMORY5_HIGH_PERFORMANCE InstanceClass = "MEMORY5_HIGH_PERFORMANCE"
	// Memory optimized instances for high performance computing, 5th generation.
	InstanceClass_R5N InstanceClass = "R5N"
	// Memory optimized instances with local NVME drive, 5th generation.
	InstanceClass_MEMORY5_NVME_DRIVE InstanceClass = "MEMORY5_NVME_DRIVE"
	// Memory optimized instances with local NVME drive, 5th generation.
	InstanceClass_R5D InstanceClass = "R5D"
	// Memory optimized instances with local NVME drive for high performance computing, 5th generation.
	InstanceClass_MEMORY5_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "MEMORY5_NVME_DRIVE_HIGH_PERFORMANCE"
	// Memory optimized instances with local NVME drive for high performance computing, 5th generation.
	InstanceClass_R5DN InstanceClass = "R5DN"
	// Memory optimized instances based on AMD EPYC, 5th generation.
	InstanceClass_MEMORY5_AMD InstanceClass = "MEMORY5_AMD"
	// Memory optimized instances based on AMD EPYC, 5th generation.
	InstanceClass_R5A InstanceClass = "R5A"
	// Memory optimized instances based on AMD EPYC with local NVME drive, 5th generation.
	InstanceClass_MEMORY5_AMD_NVME_DRIVE InstanceClass = "MEMORY5_AMD_NVME_DRIVE"
	// Memory optimized instances based on AMD EPYC with local NVME drive, 5th generation.
	InstanceClass_R5AD InstanceClass = "R5AD"
	// High memory instances (3TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_3TB_1 InstanceClass = "HIGH_MEMORY_3TB_1"
	// High memory instances (3TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_U_3TB1 InstanceClass = "U_3TB1"
	// High memory instances (6TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_6TB_1 InstanceClass = "HIGH_MEMORY_6TB_1"
	// High memory instances (6TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_U_6TB1 InstanceClass = "U_6TB1"
	// High memory instances (9TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_9TB_1 InstanceClass = "HIGH_MEMORY_9TB_1"
	// High memory instances (9TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_U_9TB1 InstanceClass = "U_9TB1"
	// High memory instances (12TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_12TB_1 InstanceClass = "HIGH_MEMORY_12TB_1"
	// High memory instances (12TB) based on Intel Xeon Platinum 8176M (Skylake) processors, 1st generation.
	InstanceClass_U_12TB1 InstanceClass = "U_12TB1"
	// High memory instances (18TB) based on Intel Xeon Scalable (Cascade Lake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_18TB_1 InstanceClass = "HIGH_MEMORY_18TB_1"
	// High memory instances (18TB) based on Intel Xeon Scalable (Cascade Lake) processors, 1st generation.
	InstanceClass_U_18TB1 InstanceClass = "U_18TB1"
	// High memory instances (24TB) based on Intel Xeon Scalable (Cascade Lake) processors, 1st generation.
	InstanceClass_HIGH_MEMORY_24TB_1 InstanceClass = "HIGH_MEMORY_24TB_1"
	// High memory instances (24TB) based on Intel Xeon Scalable (Cascade Lake) processors, 1st generation.
	InstanceClass_U_24TB1 InstanceClass = "U_24TB1"
	// High memory instances (6TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_6TB_7 InstanceClass = "HIGH_MEMORY_6TB_7"
	// High memory instances (6TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7I_6TB InstanceClass = "U7I_6TB"
	// High memory instances (8TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_8TB_7 InstanceClass = "HIGH_MEMORY_8TB_7"
	// High memory instances (8TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7I_8TB InstanceClass = "U7I_8TB"
	// High memory instances (12TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_12TB_7 InstanceClass = "HIGH_MEMORY_12TB_7"
	// High memory instances (12TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7I_12TB InstanceClass = "U7I_12TB"
	// High memory, network-intensive instances (16TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_HIGH_NETWORK_16TB_7 InstanceClass = "HIGH_MEMORY_HIGH_NETWORK_16TB_7"
	// High memory, network-intensive instances (16TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7IN_16TB InstanceClass = "U7IN_16TB"
	// High memory, network-intensive instances (24TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_HIGH_NETWORK_24TB_7 InstanceClass = "HIGH_MEMORY_HIGH_NETWORK_24TB_7"
	// High memory, network-intensive instances (24TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7IN_24TB InstanceClass = "U7IN_24TB"
	// High memory, network-intensive instances (32TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_HIGH_MEMORY_HIGH_NETWORK_32TB_7 InstanceClass = "HIGH_MEMORY_HIGH_NETWORK_32TB_7"
	// High memory, network-intensive instances (32TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation.
	InstanceClass_U7IN_32TB InstanceClass = "U7IN_32TB"
	// High memory, network-intensive instances (32TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation U7inh instances use Hewlett Packard Enterprise (HPE) Compute Scale Up Server 3200.
	InstanceClass_HIGH_MEMORY_HIGH_NETWORK_HPE_32TB_7 InstanceClass = "HIGH_MEMORY_HIGH_NETWORK_HPE_32TB_7"
	// High memory, network-intensive instances (32TB) based on 4th Generation Intel Xeon Scalable processors (Sapphire Rapids), 7th generation U7inh instances use Hewlett Packard Enterprise (HPE) Compute Scale Up Server 3200.
	InstanceClass_U7INH_32TB InstanceClass = "U7INH_32TB"
	// Memory optimized instances that are also EBS-optimized, 5th generation.
	InstanceClass_MEMORY5_EBS_OPTIMIZED InstanceClass = "MEMORY5_EBS_OPTIMIZED"
	// Memory optimized instances that are also EBS-optimized, 5th generation.
	InstanceClass_R5B InstanceClass = "R5B"
	// Memory optimized instances, 6th generation with Graviton2 processors.
	InstanceClass_MEMORY6_GRAVITON InstanceClass = "MEMORY6_GRAVITON"
	// Memory optimized instances, 6th generation with Graviton2 processors.
	InstanceClass_R6G InstanceClass = "R6G"
	// Memory optimized instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_MEMORY6_GRAVITON2_NVME_DRIVE InstanceClass = "MEMORY6_GRAVITON2_NVME_DRIVE"
	// Memory optimized instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_R6GD InstanceClass = "R6GD"
	// Memory optimized instances, 7th generation with Graviton3 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_MEMORY7_GRAVITON InstanceClass = "MEMORY7_GRAVITON"
	// Memory optimized instances, 7th generation with Graviton3 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_R7G InstanceClass = "R7G"
	// Memory optimized instances, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_MEMORY7_GRAVITON3_NVME_DRIVE InstanceClass = "MEMORY7_GRAVITON3_NVME_DRIVE"
	// Memory optimized instances, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_R7GD InstanceClass = "R7GD"
	// Memory optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation w/  3.2GHz turbo frequency.
	InstanceClass_MEMORY7_INTEL_BASE InstanceClass = "MEMORY7_INTEL_BASE"
	// Memory optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation w/  3.2GHz turbo frequency.
	InstanceClass_R7I InstanceClass = "R7I"
	// Memory optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation, with sustained 3.9GHz turbo frequency.
	InstanceClass_MEMORY7_INTEL InstanceClass = "MEMORY7_INTEL"
	// Memory optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation, with sustained 3.9GHz turbo frequency.
	InstanceClass_R7IZ InstanceClass = "R7IZ"
	// Memory optimized instances based on custom Intel Xeon Scalable 6 processors, available only on AWS.
	InstanceClass_MEMORY8_INTEL_BASE InstanceClass = "MEMORY8_INTEL_BASE"
	// Memory optimized instances based on custom Intel Xeon Scalable 6 processors, available only on AWS.
	InstanceClass_R8I InstanceClass = "R8I"
	// Memory optimized flexible instances based on custom Intel Xeon Scalable 6 processors, available only on AWS Provide up to 20% higher performance, 2.5x more memory throughput, and can scale up to full CPU performance 95% of the time.
	InstanceClass_MEMORY8_INTEL_FLEX InstanceClass = "MEMORY8_INTEL_FLEX"
	// Memory optimized flexible instances based on custom Intel Xeon Scalable 6 processors, available only on AWS Provide up to 20% higher performance, 2.5x more memory throughput, and can scale up to full CPU performance 95% of the time.
	InstanceClass_R8I_FLEX InstanceClass = "R8I_FLEX"
	// Memory optimized instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_MEMORY7_AMD InstanceClass = "MEMORY7_AMD"
	// Memory optimized instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_R7A InstanceClass = "R7A"
	// Memory optimized instances with Graviton4 processors.
	InstanceClass_MEMORY8_GRAVITON InstanceClass = "MEMORY8_GRAVITON"
	// Memory optimized instances with Graviton4 processors.
	InstanceClass_R8G InstanceClass = "R8G"
	// Memory optimized instances, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_MEMORY8_GRAVITON4_NVME_DRIVE InstanceClass = "MEMORY8_GRAVITON4_NVME_DRIVE"
	// Memory optimized instances, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_R8GD InstanceClass = "R8GD"
	// Memory optimized instances, 8th generation with Graviton4 processors and high network bandwidth capabilities.
	InstanceClass_MEMORY8_GRAVITON4_HIGH_NETWORK_BANDWIDTH InstanceClass = "MEMORY8_GRAVITON4_HIGH_NETWORK_BANDWIDTH"
	// Memory optimized instances, 8th generation with Graviton4 processors and high network bandwidth capabilities.
	InstanceClass_R8GN InstanceClass = "R8GN"
	// Memory and EBS optimized instances, 8th generation with Graviton4 processors.
	InstanceClass_MEMORY8_GRAVITON4_EBS_OPTIMIZED InstanceClass = "MEMORY8_GRAVITON4_EBS_OPTIMIZED"
	// Memory and EBS optimized instances, 8th generation with Graviton4 processors.
	InstanceClass_R8GB InstanceClass = "R8GB"
	// Compute optimized instances, 3rd generation.
	InstanceClass_COMPUTE3 InstanceClass = "COMPUTE3"
	// Compute optimized instances, 3rd generation.
	InstanceClass_C3 InstanceClass = "C3"
	// Compute optimized instances, 4th generation.
	InstanceClass_COMPUTE4 InstanceClass = "COMPUTE4"
	// Compute optimized instances, 4th generation.
	InstanceClass_C4 InstanceClass = "C4"
	// Compute optimized instances, 5th generation.
	InstanceClass_COMPUTE5 InstanceClass = "COMPUTE5"
	// Compute optimized instances, 5th generation.
	InstanceClass_C5 InstanceClass = "C5"
	// Compute optimized instances with local NVME drive, 5th generation.
	InstanceClass_COMPUTE5_NVME_DRIVE InstanceClass = "COMPUTE5_NVME_DRIVE"
	// Compute optimized instances with local NVME drive, 5th generation.
	InstanceClass_C5D InstanceClass = "C5D"
	// Compute optimized instances based on AMD EPYC, 5th generation.
	InstanceClass_COMPUTE5_AMD InstanceClass = "COMPUTE5_AMD"
	// Compute optimized instances based on AMD EPYC, 5th generation.
	InstanceClass_C5A InstanceClass = "C5A"
	// Compute optimized instances with local NVME drive based on AMD EPYC, 5th generation.
	InstanceClass_COMPUTE5_AMD_NVME_DRIVE InstanceClass = "COMPUTE5_AMD_NVME_DRIVE"
	// Compute optimized instances with local NVME drive based on AMD EPYC, 5th generation.
	InstanceClass_C5AD InstanceClass = "C5AD"
	// Compute optimized instances for high performance computing, 5th generation.
	InstanceClass_COMPUTE5_HIGH_PERFORMANCE InstanceClass = "COMPUTE5_HIGH_PERFORMANCE"
	// Compute optimized instances for high performance computing, 5th generation.
	InstanceClass_C5N InstanceClass = "C5N"
	// Compute optimized instances, 6th generation.
	InstanceClass_COMPUTE6_INTEL InstanceClass = "COMPUTE6_INTEL"
	// Compute optimized instances, 6th generation.
	InstanceClass_C6I InstanceClass = "C6I"
	// Compute optimized instances with local NVME drive, 6th generation.
	InstanceClass_COMPUTE6_INTEL_NVME_DRIVE InstanceClass = "COMPUTE6_INTEL_NVME_DRIVE"
	// Compute optimized instances with local NVME drive, 6th generation.
	InstanceClass_C6ID InstanceClass = "C6ID"
	// Compute optimized instances for high performance computing, 6th generation.
	InstanceClass_COMPUTE6_INTEL_HIGH_PERFORMANCE InstanceClass = "COMPUTE6_INTEL_HIGH_PERFORMANCE"
	// Compute optimized instances for high performance computing, 6th generation.
	InstanceClass_C6IN InstanceClass = "C6IN"
	// Compute optimized instances based on AMD EPYC (codename Milan), 6th generation.
	InstanceClass_COMPUTE6_AMD InstanceClass = "COMPUTE6_AMD"
	// Compute optimized instances based on AMD EPYC (codename Milan), 6th generation.
	InstanceClass_C6A InstanceClass = "C6A"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors.
	InstanceClass_COMPUTE6_GRAVITON2 InstanceClass = "COMPUTE6_GRAVITON2"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors.
	InstanceClass_C6G InstanceClass = "C6G"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors.
	InstanceClass_COMPUTE7_GRAVITON3 InstanceClass = "COMPUTE7_GRAVITON3"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors.
	InstanceClass_C7G InstanceClass = "C7G"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Frankfurt).
	InstanceClass_COMPUTE8_GRAVITON4 InstanceClass = "COMPUTE8_GRAVITON4"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Frankfurt).
	InstanceClass_C8G InstanceClass = "C8G"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_COMPUTE6_GRAVITON2_NVME_DRIVE InstanceClass = "COMPUTE6_GRAVITON2_NVME_DRIVE"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_C6GD InstanceClass = "C6GD"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_COMPUTE7_GRAVITON3_NVME_DRIVE InstanceClass = "COMPUTE7_GRAVITON3_NVME_DRIVE"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_C7GD InstanceClass = "C7GD"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_COMPUTE8_GRAVITON4_NVME_DRIVE InstanceClass = "COMPUTE8_GRAVITON4_NVME_DRIVE"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_C8GD InstanceClass = "C8GD"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and high network bandwidth capabilities.
	InstanceClass_COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH InstanceClass = "COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and high network bandwidth capabilities.
	InstanceClass_C6GN InstanceClass = "C6GN"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors and high network bandwidth capabilities.
	InstanceClass_COMPUTE7_GRAVITON3_HIGH_NETWORK_BANDWIDTH InstanceClass = "COMPUTE7_GRAVITON3_HIGH_NETWORK_BANDWIDTH"
	// Compute optimized instances for high performance computing, 7th generation with Graviton3 processors and high network bandwidth capabilities.
	InstanceClass_C7GN InstanceClass = "C7GN"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors and high network bandwidth capabilities.
	InstanceClass_COMPUTE8_GRAVITON4_HIGH_NETWORK_BANDWIDTH InstanceClass = "COMPUTE8_GRAVITON4_HIGH_NETWORK_BANDWIDTH"
	// Compute optimized instances for high performance computing, 8th generation with Graviton4 processors and high network bandwidth capabilities.
	InstanceClass_C8GN InstanceClass = "C8GN"
	// Compute optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation.
	InstanceClass_COMPUTE7_INTEL InstanceClass = "COMPUTE7_INTEL"
	// Compute optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation.
	InstanceClass_C7I InstanceClass = "C7I"
	// Compute optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation C7i-flex instances efficiently use compute resources to deliver a baseline level of performance with the ability to scale up to the full compute performance a majority of the time.
	InstanceClass_COMPUTE7_INTEL_FLEX InstanceClass = "COMPUTE7_INTEL_FLEX"
	// Compute optimized instances based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation C7i-flex instances efficiently use compute resources to deliver a baseline level of performance with the ability to scale up to the full compute performance a majority of the time.
	InstanceClass_C7I_FLEX InstanceClass = "C7I_FLEX"
	// Compute optimized instances based on custom Intel Xeon 6 processors, available only on AWS, 8th generation.
	InstanceClass_COMPUTE8_INTEL InstanceClass = "COMPUTE8_INTEL"
	// Compute optimized instances based on custom Intel Xeon 6 processors, available only on AWS, 8th generation.
	InstanceClass_C8I InstanceClass = "C8I"
	// Compute optimized instances based on custom Intel Xeon 6 processors, available only on AWS, 8th generation C8i-flex instances efficiently use compute resources to deliver a baseline level of performance with the ability to scale up to the full compute performance a majority of the time.
	InstanceClass_COMPUTE8_INTEL_FLEX InstanceClass = "COMPUTE8_INTEL_FLEX"
	// Compute optimized instances based on custom Intel Xeon 6 processors, available only on AWS, 8th generation C8i-flex instances efficiently use compute resources to deliver a baseline level of performance with the ability to scale up to the full compute performance a majority of the time.
	InstanceClass_C8I_FLEX InstanceClass = "C8I_FLEX"
	// Compute optimized instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_COMPUTE7_AMD InstanceClass = "COMPUTE7_AMD"
	// Compute optimized instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_C7A InstanceClass = "C7A"
	// Storage-optimized instances, 2nd generation.
	InstanceClass_STORAGE2 InstanceClass = "STORAGE2"
	// Storage-optimized instances, 2nd generation.
	InstanceClass_D2 InstanceClass = "D2"
	// Storage-optimized instances, 3rd generation.
	InstanceClass_STORAGE3 InstanceClass = "STORAGE3"
	// Storage-optimized instances, 3rd generation.
	InstanceClass_D3 InstanceClass = "D3"
	// Storage-optimized instances, 3rd generation.
	InstanceClass_STORAGE3_ENHANCED_NETWORK InstanceClass = "STORAGE3_ENHANCED_NETWORK"
	// Storage-optimized instances, 3rd generation.
	InstanceClass_D3EN InstanceClass = "D3EN"
	// Storage/compute balanced instances, 1st generation.
	InstanceClass_STORAGE_COMPUTE_1 InstanceClass = "STORAGE_COMPUTE_1"
	// Storage/compute balanced instances, 1st generation.
	InstanceClass_H1 InstanceClass = "H1"
	// High performance computing powered by AWS Trainium.
	InstanceClass_TRAINING_ACCELERATOR1 InstanceClass = "TRAINING_ACCELERATOR1"
	// High performance computing powered by AWS Trainium.
	InstanceClass_TRN1 InstanceClass = "TRN1"
	// Network-optimized high performance computing powered by AWS Trainium.
	InstanceClass_TRAINING_ACCELERATOR1_ENHANCED_NETWORK InstanceClass = "TRAINING_ACCELERATOR1_ENHANCED_NETWORK"
	// Network-optimized high performance computing powered by AWS Trainium.
	InstanceClass_TRN1N InstanceClass = "TRN1N"
	// High performance computing powered by AWS Trainium2, 2nd generation.
	InstanceClass_TRAINING_ACCELERATOR2 InstanceClass = "TRAINING_ACCELERATOR2"
	// High performance computing powered by AWS Trainium2, 2nd generation.
	InstanceClass_TRN2 InstanceClass = "TRN2"
	// High performance computing powered by AWS Trainium2 and EC2 Ultra Servers, 2nd generation UltraServers connect multiple EC2 instances using a dedicated, high-bandwidth, low-latency accelerator interconnect.
	InstanceClass_TRAINING_ACCELERATOR2_ULTRASERVER InstanceClass = "TRAINING_ACCELERATOR2_ULTRASERVER"
	// High performance computing powered by AWS Trainium2 and EC2 Ultra Servers, 2nd generation UltraServers connect multiple EC2 instances using a dedicated, high-bandwidth, low-latency accelerator interconnect.
	InstanceClass_TRN2U InstanceClass = "TRN2U"
	// I/O-optimized instances, 3rd generation.
	InstanceClass_IO3 InstanceClass = "IO3"
	// I/O-optimized instances, 3rd generation.
	InstanceClass_I3 InstanceClass = "I3"
	// I/O-optimized instances with local NVME drive, 3rd generation.
	InstanceClass_IO3_DENSE_NVME_DRIVE InstanceClass = "IO3_DENSE_NVME_DRIVE"
	// I/O-optimized instances with local NVME drive, 3rd generation.
	InstanceClass_I3EN InstanceClass = "I3EN"
	// I/O-optimized instances with local NVME drive powered by Intel Xeon Scalable processors (code named Ice Lake), 4th generation.
	InstanceClass_IO4_INTEL InstanceClass = "IO4_INTEL"
	// I/O-optimized instances with local NVME drive powered by Intel Xeon Scalable processors (code named Ice Lake), 4th generation.
	InstanceClass_I4I InstanceClass = "I4I"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_STORAGE4_GRAVITON InstanceClass = "STORAGE4_GRAVITON"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_I4G InstanceClass = "I4G"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_STORAGE4_GRAVITON_NETWORK_OPTIMIZED InstanceClass = "STORAGE4_GRAVITON_NETWORK_OPTIMIZED"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_IM4GN InstanceClass = "IM4GN"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED InstanceClass = "STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_IS4GEN InstanceClass = "IS4GEN"
	// Storage optimized instances powered by 5th generation Intel Xeon Scalable processors, 7th generation.
	InstanceClass_STORAGE7_INTEL_STORAGE_OPTIMIZED InstanceClass = "STORAGE7_INTEL_STORAGE_OPTIMIZED"
	// Storage optimized instances powered by 5th generation Intel Xeon Scalable processors, 7th generation.
	InstanceClass_I7IE InstanceClass = "I7IE"
	// I/O-optimized instances with local NVME drive powered by 5th generation Intel Xeon Scalable processors, 7th generation.
	InstanceClass_IO7_INTEL InstanceClass = "IO7_INTEL"
	// I/O-optimized instances with local NVME drive powered by 5th generation Intel Xeon Scalable processors, 7th generation.
	InstanceClass_I7I InstanceClass = "I7I"
	// Storage optimized instances powered by Graviton4 processor, 8th generation.
	InstanceClass_STORAGE8_GRAVITON InstanceClass = "STORAGE8_GRAVITON"
	// Storage optimized instances powered by Graviton4 processor, 8th generation.
	InstanceClass_I8G InstanceClass = "I8G"
	// Storage optimized instances powered by Graviton4 processor, 8th generation.
	InstanceClass_STORAGE8_GRAVITON_STORAGE_OPTIMIZED InstanceClass = "STORAGE8_GRAVITON_STORAGE_OPTIMIZED"
	// Storage optimized instances powered by Graviton4 processor, 8th generation.
	InstanceClass_I8GE InstanceClass = "I8GE"
	// Burstable instances, 2nd generation.
	InstanceClass_BURSTABLE2 InstanceClass = "BURSTABLE2"
	// Burstable instances, 2nd generation.
	InstanceClass_T2 InstanceClass = "T2"
	// Burstable instances, 3rd generation.
	InstanceClass_BURSTABLE3 InstanceClass = "BURSTABLE3"
	// Burstable instances, 3rd generation.
	InstanceClass_T3 InstanceClass = "T3"
	// Burstable instances based on AMD EPYC, 3rd generation.
	InstanceClass_BURSTABLE3_AMD InstanceClass = "BURSTABLE3_AMD"
	// Burstable instances based on AMD EPYC, 3rd generation.
	InstanceClass_T3A InstanceClass = "T3A"
	// Burstable instances, 4th generation with Graviton2 processors.
	InstanceClass_BURSTABLE4_GRAVITON InstanceClass = "BURSTABLE4_GRAVITON"
	// Burstable instances, 4th generation with Graviton2 processors.
	InstanceClass_T4G InstanceClass = "T4G"
	// Memory-intensive instances, 1st generation.
	InstanceClass_MEMORY_INTENSIVE_1 InstanceClass = "MEMORY_INTENSIVE_1"
	// Memory-intensive instances, 1st generation.
	InstanceClass_X1 InstanceClass = "X1"
	// Memory-intensive instances, extended, 1st generation.
	InstanceClass_MEMORY_INTENSIVE_1_EXTENDED InstanceClass = "MEMORY_INTENSIVE_1_EXTENDED"
	// Memory-intensive instances, extended, 1st generation.
	InstanceClass_X1E InstanceClass = "X1E"
	// Memory-intensive instances, 2nd generation with Graviton2 processors.
	//
	// This instance type can be used only in RDS. It is not supported in EC2.
	InstanceClass_MEMORY_INTENSIVE_2_GRAVITON2 InstanceClass = "MEMORY_INTENSIVE_2_GRAVITON2"
	// Memory-intensive instances, 2nd generation with Graviton2 processors.
	//
	// This instance type can be used only in RDS. It is not supported in EC2.
	InstanceClass_X2G InstanceClass = "X2G"
	// Memory-intensive instances, 2nd generation with Graviton2 processors and local NVME drive.
	InstanceClass_MEMORY_INTENSIVE_2_GRAVITON2_NVME_DRIVE InstanceClass = "MEMORY_INTENSIVE_2_GRAVITON2_NVME_DRIVE"
	// Memory-intensive instances, 2nd generation with Graviton2 processors and local NVME drive.
	InstanceClass_X2GD InstanceClass = "X2GD"
	// Memory-intensive instances with higher network bandwith, local NVME drive, and extended memory.
	//
	// Intel Xeon Scalable (Ice Lake) processors.
	InstanceClass_MEMORY_INTENSIVE_2_XT_INTEL InstanceClass = "MEMORY_INTENSIVE_2_XT_INTEL"
	// Memory-intensive instances with higher network bandwith, local NVME drive, and extended memory.
	//
	// Intel Xeon Scalable (Ice Lake) processors.
	InstanceClass_X2IEDN InstanceClass = "X2IEDN"
	// Memory-intensive instances with higher network bandwith and local NVME drive, Intel Xeon Scalable (Ice Lake) processors.
	InstanceClass_MEMORY_INTENSIVE_2_INTEL InstanceClass = "MEMORY_INTENSIVE_2_INTEL"
	// Memory-intensive instances with higher network bandwith and local NVME drive, Intel Xeon Scalable (Ice Lake) processors.
	InstanceClass_X2IDN InstanceClass = "X2IDN"
	// Memory-intensive instances with higher network bandwith and single-threaded performance, Intel Xeon Scalable (Cascade Lake) processors.
	InstanceClass_MEMORY_INTENSIVE_2_XTZ_INTEL InstanceClass = "MEMORY_INTENSIVE_2_XTZ_INTEL"
	// Memory-intensive instances with higher network bandwith and single-threaded performance, Intel Xeon Scalable (Cascade Lake) processors.
	InstanceClass_X2IEZN InstanceClass = "X2IEZN"
	// Memory-intensive instances powered by Graviton4 processors, 8th generation.
	InstanceClass_MEMORY_INTENSIVE_8_GRAVITON InstanceClass = "MEMORY_INTENSIVE_8_GRAVITON"
	// Memory-intensive instances powered by Graviton4 processors, 8th generation.
	InstanceClass_X8G InstanceClass = "X8G"
	// Instances with customizable hardware acceleration, 1st generation.
	InstanceClass_FPGA1 InstanceClass = "FPGA1"
	// Instances with customizable hardware acceleration, 1st generation.
	InstanceClass_F1 InstanceClass = "F1"
	// Instances with customizable hardware acceleration, 2nd generation.
	InstanceClass_FPGA2 InstanceClass = "FPGA2"
	// Instances with customizable hardware acceleration, 2nd generation.
	InstanceClass_F2 InstanceClass = "F2"
	// Graphics-optimized instances, 3rd generation.
	InstanceClass_GRAPHICS3_SMALL InstanceClass = "GRAPHICS3_SMALL"
	// Graphics-optimized instances, 3rd generation.
	InstanceClass_G3S InstanceClass = "G3S"
	// Graphics-optimized instances, 3rd generation.
	InstanceClass_GRAPHICS3 InstanceClass = "GRAPHICS3"
	// Graphics-optimized instances, 3rd generation.
	InstanceClass_G3 InstanceClass = "G3"
	// Graphics-optimized instances with NVME drive for high performance computing, 4th generation.
	InstanceClass_GRAPHICS4_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "GRAPHICS4_NVME_DRIVE_HIGH_PERFORMANCE"
	// Graphics-optimized instances with NVME drive for high performance computing, 4th generation.
	InstanceClass_G4DN InstanceClass = "G4DN"
	// Graphics-optimized instances based on AMD EPYC And Radeon Pro GPU (NAVI) with local NVME drive, 4th generation.
	InstanceClass_GRAPHICS4_AMD_NVME_DRIVE InstanceClass = "GRAPHICS4_AMD_NVME_DRIVE"
	// Graphics-optimized instances based on AMD EPYC And Radeon Pro GPU (NAVI) with local NVME drive, 4th generation.
	InstanceClass_G4AD InstanceClass = "G4AD"
	// Graphics-optimized instances, 5th generation.
	InstanceClass_GRAPHICS5 InstanceClass = "GRAPHICS5"
	// Graphics-optimized instances, 5th generation.
	InstanceClass_G5 InstanceClass = "G5"
	// Graphics-optimized instances powered by AWS Graviton2 Processors and NVIDIA T4G Tensor Core GPUs, 5th generation.
	InstanceClass_GRAPHICS5_GRAVITON2 InstanceClass = "GRAPHICS5_GRAVITON2"
	// Graphics-optimized instances powered by AWS Graviton2 Processors and NVIDIA T4G Tensor Core GPUs, 5th generation.
	InstanceClass_G5G InstanceClass = "G5G"
	// Graphics-optimized instances, 6th generation.
	InstanceClass_GRAPHICS6 InstanceClass = "GRAPHICS6"
	// Graphics-optimized instances, 6th generation.
	InstanceClass_G6 InstanceClass = "G6"
	// Cost-efficient GPU-based instances for AI inference and spatial computing workloads, 6th generation.
	InstanceClass_GRAPHICS6_EFFICIENT InstanceClass = "GRAPHICS6_EFFICIENT"
	// Cost-efficient GPU-based instances for AI inference and spatial computing workloads, 6th generation.
	InstanceClass_G6E InstanceClass = "G6E"
	// Graphics-optimized instances, 6th generation Gr6 instances offer a 1:8 vCPU to RAM ratio, making them better suited for graphics workloads with higher memory requirements.
	InstanceClass_GRAPHICS_RAM_6 InstanceClass = "GRAPHICS_RAM_6"
	// Graphics-optimized instances, 6th generation Gr6 instances offer a 1:8 vCPU to RAM ratio, making them better suited for graphics workloads with higher memory requirements.
	InstanceClass_GR6 InstanceClass = "GR6"
	// Parallel-processing optimized instances, 2nd generation.
	InstanceClass_PARALLEL2 InstanceClass = "PARALLEL2"
	// Parallel-processing optimized instances, 2nd generation.
	InstanceClass_P2 InstanceClass = "P2"
	// Parallel-processing optimized instances, 3rd generation.
	InstanceClass_PARALLEL3 InstanceClass = "PARALLEL3"
	// Parallel-processing optimized instances, 3rd generation.
	InstanceClass_P3 InstanceClass = "P3"
	// Parallel-processing optimized instances with local NVME drive for high performance computing, 3rd generation.
	InstanceClass_PARALLEL3_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "PARALLEL3_NVME_DRIVE_HIGH_PERFORMANCE"
	// Parallel-processing optimized instances with local NVME drive for high performance computing, 3rd generation.
	InstanceClass_P3DN InstanceClass = "P3DN"
	// Parallel-processing optimized instances with local NVME drive, extended, 4th generation (in developer preview).
	InstanceClass_PARALLEL4_NVME_DRIVE_EXTENDED InstanceClass = "PARALLEL4_NVME_DRIVE_EXTENDED"
	// Parallel-processing optimized instances with local NVME drive, extended, 4th generation (in developer preview).
	InstanceClass_P4DE InstanceClass = "P4DE"
	// Parallel-processing optimized instances, 4th generation.
	InstanceClass_PARALLEL4 InstanceClass = "PARALLEL4"
	// Parallel-processing optimized instances, 4th generation.
	InstanceClass_P4D InstanceClass = "P4D"
	// Parallel-processing optimized instances powered by NVIDIA H100 Tensor Core GPUs, 5th generation.
	InstanceClass_PARALLEL5 InstanceClass = "PARALLEL5"
	// Parallel-processing optimized instances powered by NVIDIA H100 Tensor Core GPUs, 5th generation.
	InstanceClass_P5 InstanceClass = "P5"
	// Parallel-processing optimized instances powered by NVIDIA H200 Tensor Core GPUs, 5th generation.
	InstanceClass_PARALLEL5_EXTENDED InstanceClass = "PARALLEL5_EXTENDED"
	// Parallel-processing optimized instances, 5th generation.
	InstanceClass_P5E InstanceClass = "P5E"
	// Parallel-processing, network-optimized instances powered by NVIDIA H200 Tensor Core GPUs, 5th generation P5en instances pair NVIDIA H200 Tensor Core GPUs with Intel Sapphire Rapids CPU, enabling Gen5 PCIe between CPU and GPU.
	InstanceClass_PARALLEL5_EXTENDED_NETWORK InstanceClass = "PARALLEL5_EXTENDED_NETWORK"
	// Parallel-processing optimized instances, 5th generation.
	InstanceClass_P5EN InstanceClass = "P5EN"
	// Arm processor based instances, 1st generation.
	InstanceClass_ARM1 InstanceClass = "ARM1"
	// Arm processor based instances, 1st generation.
	InstanceClass_A1 InstanceClass = "A1"
	// Arm processor based instances, 2nd generation.
	InstanceClass_STANDARD6_GRAVITON InstanceClass = "STANDARD6_GRAVITON"
	// Arm processor based instances, 2nd generation.
	InstanceClass_M6G InstanceClass = "M6G"
	// Standard instances based on Intel (Ice Lake), 6th generation.
	InstanceClass_STANDARD6_INTEL InstanceClass = "STANDARD6_INTEL"
	// Standard instances based on Intel (Ice Lake), 6th generation.
	InstanceClass_M6I InstanceClass = "M6I"
	// Standard instances based on Intel (Ice Lake) with local NVME drive, 6th generation.
	InstanceClass_STANDARD6_INTEL_NVME_DRIVE InstanceClass = "STANDARD6_INTEL_NVME_DRIVE"
	// Standard instances based on Intel (Ice Lake) with local NVME drive, 6th generation.
	InstanceClass_M6ID InstanceClass = "M6ID"
	// Standard instances for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_STANDARD6_INTEL_HIGH_PERFORMANCE InstanceClass = "STANDARD6_INTEL_HIGH_PERFORMANCE"
	// Standard instances for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_M6IN InstanceClass = "M6IN"
	// Standard instances with local NVME drive for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_STANDARD6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE InstanceClass = "STANDARD6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE"
	// Standard instances with local NVME drive for high performance computing powered by Intel Xeon Scalable processors (code named Ice Lake), 6th generation.
	InstanceClass_M6IDN InstanceClass = "M6IDN"
	// Standard instances based on 3rd Gen AMD EPYC processors, 6th generation.
	InstanceClass_STANDARD6_AMD InstanceClass = "STANDARD6_AMD"
	// Standard instances based on 3rd Gen AMD EPYC processors, 6th generation.
	InstanceClass_M6A InstanceClass = "M6A"
	// Standard instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_STANDARD6_GRAVITON2_NVME_DRIVE InstanceClass = "STANDARD6_GRAVITON2_NVME_DRIVE"
	// Standard instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_M6GD InstanceClass = "M6GD"
	// Standard instances, 7th generation with Graviton3 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_STANDARD7_GRAVITON InstanceClass = "STANDARD7_GRAVITON"
	// Standard instances, 7th generation with Graviton3 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_M7G InstanceClass = "M7G"
	// Standard instances, 8th generation with Graviton4 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Frankfurt).
	InstanceClass_STANDARD8_GRAVITON InstanceClass = "STANDARD8_GRAVITON"
	// Standard instances, 8th generation with Graviton4 processors.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Frankfurt).
	InstanceClass_M8G InstanceClass = "M8G"
	// Standard instances, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_STANDARD7_GRAVITON3_NVME_DRIVE InstanceClass = "STANDARD7_GRAVITON3_NVME_DRIVE"
	// Standard instances, 7th generation with Graviton3 processors and local NVME drive.
	//
	// This instance class is currently only available in US East (Ohio), US East (N. Virginia), US West (Oregon), and Europe (Ireland).
	InstanceClass_M7GD InstanceClass = "M7GD"
	// Standard instances, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_STANDARD8_GRAVITON4_NVME_DRIVE InstanceClass = "STANDARD8_GRAVITON4_NVME_DRIVE"
	// Standard instances, 8th generation with Graviton4 processors and local NVME drive.
	InstanceClass_M8GD InstanceClass = "M8GD"
	// Standard instances based on custom Intel Xeon Scalable 6 processors, available only on AWS.
	InstanceClass_STANDARD8_INTEL InstanceClass = "STANDARD8_INTEL"
	// Standard instances based on custom Intel Xeon Scalable 6 processors, available only on AWS.
	InstanceClass_M8I InstanceClass = "M8I"
	// Flexible instances based on custom Intel Xeon Scalable 6 processors, available only on AWS Provide up to 20% higher performance, 2.5x more memory throughput, and can scale up to full CPU performance 95% of the time.
	InstanceClass_STANDARD8_INTEL_FLEX InstanceClass = "STANDARD8_INTEL_FLEX"
	// Flexible instances based on custom Intel Xeon Scalable 6 processors, available only on AWS Provide up to 20% higher performance, 2.5x more memory throughput, and can scale up to full CPU performance 95% of the time.
	InstanceClass_M8I_FLEX InstanceClass = "M8I_FLEX"
	// Standard instances with high memory and compute capacity based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation.
	InstanceClass_STANDARD7_INTEL InstanceClass = "STANDARD7_INTEL"
	// Standard instances with high memory and compute capacity based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation.
	InstanceClass_M7I InstanceClass = "M7I"
	// Flexible instances with high memory and compute capacity based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation The M7i-Flex instances deliver a baseline of 40% CPU performance, and can scale up to full CPU performance 95% of the time.
	InstanceClass_STANDARD7_INTEL_FLEX InstanceClass = "STANDARD7_INTEL_FLEX"
	// Flexible instances with high memory and compute capacity based on Intel Xeon Scalable (Sapphire Rapids) processors, 7th generation The M7i-Flex instances deliver a baseline of 40% CPU performance, and can scale up to full CPU performance 95% of the time.
	InstanceClass_M7I_FLEX InstanceClass = "M7I_FLEX"
	// Standard instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_STANDARD7_AMD InstanceClass = "STANDARD7_AMD"
	// Standard instances based on 4th generation AMD EPYC (codename Genoa), 7th generation.
	InstanceClass_M7A InstanceClass = "M7A"
	// Standard instances based on 5th generation AMD EPYC (formerly code named Turin), 8th generation.
	InstanceClass_STANDARD8_AMD InstanceClass = "STANDARD8_AMD"
	// Standard instances based on 5th generation AMD EPYC (formerly code named Turin), 8th generation.
	InstanceClass_M8A InstanceClass = "M8A"
	// High memory and compute capacity instances, 1st generation.
	InstanceClass_HIGH_COMPUTE_MEMORY1 InstanceClass = "HIGH_COMPUTE_MEMORY1"
	// High memory and compute capacity instances, 1st generation.
	InstanceClass_Z1D InstanceClass = "Z1D"
	// Inferentia Chips based instances for machine learning inference applications, 1st generation.
	InstanceClass_INFERENCE1 InstanceClass = "INFERENCE1"
	// Inferentia Chips based instances for machine learning inference applications, 1st generation.
	InstanceClass_INF1 InstanceClass = "INF1"
	// Inferentia Chips based instances for machine learning inference applications, 2nd generation.
	InstanceClass_INFERENCE2 InstanceClass = "INFERENCE2"
	// Inferentia Chips based instances for machine learning inference applications, 2nd generation.
	InstanceClass_INF2 InstanceClass = "INF2"
	// Macintosh instances built on Apple Mac mini computers, 1st generation with Intel procesors.
	InstanceClass_MACINTOSH1_INTEL InstanceClass = "MACINTOSH1_INTEL"
	// Macintosh instances built on Apple Mac mini computers, 1st generation with Intel procesors.
	InstanceClass_MAC1 InstanceClass = "MAC1"
	// Macintosh instances built on Apple Mac mini 2020 computers, 2nd generation with Apple silicon M1 processors.
	InstanceClass_MACINTOSH2_M1 InstanceClass = "MACINTOSH2_M1"
	// Macintosh instances built on Apple Mac mini 2020 computers, 2nd generation with Apple silicon M1 processors.
	InstanceClass_MAC2 InstanceClass = "MAC2"
	// Macintosh instances built on Apple Mac mini 2023 computers, 2nd generation with Apple silicon M2 processors.
	InstanceClass_MACINTOSH2_M2 InstanceClass = "MACINTOSH2_M2"
	// Macintosh instances built on Apple Mac mini 2023 computers, 2nd generation with Apple silicon M2 processors.
	InstanceClass_MAC2_M2 InstanceClass = "MAC2_M2"
	// Macintosh instances built on Apple Mac mini 2023 computers, 2nd generation with Apple silicon M2 Pro processors.
	InstanceClass_MACINTOSH2_M2_PRO InstanceClass = "MACINTOSH2_M2_PRO"
	// Macintosh instances built on Apple Mac mini 2023 computers, 2nd generation with Apple silicon M2 Pro processors.
	InstanceClass_MAC2_M2PRO InstanceClass = "MAC2_M2PRO"
	// Macintosh instances built on 2022 Mac Studio hardware powered by Apple silicon M1 Ultra processors.
	InstanceClass_MACINTOSH2_M1_ULTRA InstanceClass = "MACINTOSH2_M1_ULTRA"
	// Macintosh instances built on 2022 Mac Studio hardware powered by Apple silicon M1 Ultra processors.
	InstanceClass_MAC2_M1ULTRA InstanceClass = "MAC2_M1ULTRA"
	// Multi-stream video transcoding instances for resolutions up to 4K UHD, 1st generation.
	InstanceClass_VIDEO_TRANSCODING1 InstanceClass = "VIDEO_TRANSCODING1"
	// Multi-stream video transcoding instances for resolutions up to 4K UHD, 1st generation.
	InstanceClass_VT1 InstanceClass = "VT1"
	// High performance computing based on AMD EPYC, 6th generation.
	InstanceClass_HIGH_PERFORMANCE_COMPUTING6_AMD InstanceClass = "HIGH_PERFORMANCE_COMPUTING6_AMD"
	// High performance computing based on AMD EPYC, 6th generation.
	InstanceClass_HPC6A InstanceClass = "HPC6A"
	// High performance computing with local NVME drive based on 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake), 6th generation.
	InstanceClass_HIGH_PERFORMANCE_COMPUTING6_INTEL_NVME_DRIVE InstanceClass = "HIGH_PERFORMANCE_COMPUTING6_INTEL_NVME_DRIVE"
	// High performance computing with local NVME drive based on 6th generation with Intel Xeon Scalable processors (3rd generation processors code named Ice Lake), 6th generation.
	InstanceClass_HPC6ID InstanceClass = "HPC6ID"
	// High performance computing based on AMD EPYC, 7th generation.
	InstanceClass_HIGH_PERFORMANCE_COMPUTING7_AMD InstanceClass = "HIGH_PERFORMANCE_COMPUTING7_AMD"
	// High performance computing based on AMD EPYC, 7th generation.
	InstanceClass_HPC7A InstanceClass = "HPC7A"
	// High performance computing based on Graviton, 7th generation.
	InstanceClass_HIGH_PERFORMANCE_COMPUTING7_GRAVITON InstanceClass = "HIGH_PERFORMANCE_COMPUTING7_GRAVITON"
	// High performance computing based on Graviton, 7th generation.
	InstanceClass_HPC7G InstanceClass = "HPC7G"
	// Deep learning instances powered by Gaudi accelerators from Habana Labs (an Intel company), 1st generation.
	InstanceClass_DEEP_LEARNING1 InstanceClass = "DEEP_LEARNING1"
	// Deep learning instances powered by Gaudi accelerators from Habana Labs (an Intel company), 1st generation.
	InstanceClass_DL1 InstanceClass = "DL1"
	// Deep learning instances powered by Qualcomm AI 100 Standard accelerators, 2nd generation.
	InstanceClass_DEEP_LEARNING2_QUALCOMM InstanceClass = "DEEP_LEARNING2_QUALCOMM"
	// Deep learning instances powered by Qualcomm AI 100 Standard accelerators, 2nd generation.
	InstanceClass_DL2Q InstanceClass = "DL2Q"
)

