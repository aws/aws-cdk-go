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
//   var vpc vpc
//
//   var sourceInstance databaseInstance
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &databaseInstanceFromSnapshotProps{
//   	snapshotIdentifier: jsii.String("my-snapshot"),
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	// optional, defaults to m5.large
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_LARGE),
//   	vpc: vpc,
//   })
//   rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &databaseInstanceReadReplicaProps{
//   	sourceDatabaseInstance: sourceInstance,
//   	instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE2, ec2.*instanceSize_LARGE),
//   	vpc: vpc,
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
	// Memory optimized instances based on AMD EPYC with local NVME drive, 5th generation.
	InstanceClass_R5AD InstanceClass = "R5AD"
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
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_COMPUTE6_GRAVITON2_NVME_DRIVE InstanceClass = "COMPUTE6_GRAVITON2_NVME_DRIVE"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_C6GD InstanceClass = "C6GD"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and high network bandwidth capabilities.
	InstanceClass_COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWITH InstanceClass = "COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWITH"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and high network bandwidth capabilities.
	InstanceClass_COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH InstanceClass = "COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH"
	// Compute optimized instances for high performance computing, 6th generation with Graviton2 processors and high network bandwidth capabilities.
	InstanceClass_C6GN InstanceClass = "C6GN"
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
	InstanceClass_STORAGE4_GRAVITON_NETWORK_OPTIMIZED InstanceClass = "STORAGE4_GRAVITON_NETWORK_OPTIMIZED"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_IM4GN InstanceClass = "IM4GN"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED InstanceClass = "STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED"
	// Storage optimized instances powered by Graviton2 processor, 4th generation.
	InstanceClass_IS4GEN InstanceClass = "IS4GEN"
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
	// Memory-intensive instances, 1st generation.
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
	// Instances with customizable hardware acceleration, 1st generation.
	InstanceClass_FPGA1 InstanceClass = "FPGA1"
	// Instances with customizable hardware acceleration, 1st generation.
	InstanceClass_F1 InstanceClass = "F1"
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
	// Parallel-processing optimized instances, 2nd generation.
	InstanceClass_PARALLEL2 InstanceClass = "PARALLEL2"
	// Parallel-processing optimized instances, 2nd generation.
	InstanceClass_P2 InstanceClass = "P2"
	// Parallel-processing optimized instances, 3nd generation.
	InstanceClass_PARALLEL3 InstanceClass = "PARALLEL3"
	// Parallel-processing optimized instances, 3rd generation.
	InstanceClass_P3 InstanceClass = "P3"
	// Parallel-processing optimized instances with local NVME drive for high performance computing, 3nd generation.
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
	// Standard instances based on 3rd Gen AMD EPYC processors, 6th generation.
	InstanceClass_STANDARD6_AMD InstanceClass = "STANDARD6_AMD"
	// Standard instances based on 3rd Gen AMD EPYC processors, 6th generation.
	InstanceClass_M6A InstanceClass = "M6A"
	// Standard instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_STANDARD6_GRAVITON2_NVME_DRIVE InstanceClass = "STANDARD6_GRAVITON2_NVME_DRIVE"
	// Standard instances, 6th generation with Graviton2 processors and local NVME drive.
	InstanceClass_M6GD InstanceClass = "M6GD"
	// High memory and compute capacity instances, 1st generation.
	InstanceClass_HIGH_COMPUTE_MEMORY1 InstanceClass = "HIGH_COMPUTE_MEMORY1"
	// High memory and compute capacity instances, 1st generation.
	InstanceClass_Z1D InstanceClass = "Z1D"
	// Inferentia Chips based instances for machine learning inference applications, 1st generation.
	InstanceClass_INFERENCE1 InstanceClass = "INFERENCE1"
	// Inferentia Chips based instances for machine learning inference applications, 1st generation.
	InstanceClass_INF1 InstanceClass = "INF1"
	// Macintosh instances built on Apple Mac mini computers, 1st generation with Intel procesors.
	InstanceClass_MACINTOSH1_INTEL InstanceClass = "MACINTOSH1_INTEL"
	// Macintosh instances built on Apple Mac mini computers, 1st generation with Intel procesors.
	InstanceClass_MAC1 InstanceClass = "MAC1"
	// Multi-stream video transcoding instances for resolutions up to 4K UHD, 1st generation.
	InstanceClass_VIDEO_TRANSCODING1 InstanceClass = "VIDEO_TRANSCODING1"
	// Multi-stream video transcoding instances for resolutions up to 4K UHD, 1st generation.
	InstanceClass_VT1 InstanceClass = "VT1"
	// High performance computing based on AMD EPYC, 6th generation.
	InstanceClass_HIGH_PERFORMANCE_COMPUTING6_AMD InstanceClass = "HIGH_PERFORMANCE_COMPUTING6_AMD"
	// High performance computing based on AMD EPYC, 6th generation.
	InstanceClass_HPC6A InstanceClass = "HPC6A"
	// Deep learning instances powered by Gaudi accelerators from Habana Labs (an Intel company), 1st generation.
	InstanceClass_DEEP_LEARNING1 InstanceClass = "DEEP_LEARNING1"
	// Deep learning instances powered by Gaudi accelerators from Habana Labs (an Intel company), 1st generation.
	InstanceClass_DL1 InstanceClass = "DL1"
)

