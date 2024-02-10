package awsefs


// EFS Performance mode.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
//   	 // default
//   	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   	 // files are not transitioned back from (infrequent access) IA to primary storage by default
//   	TransitionToArchivePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to Archive by default
//   	ReplicationOverwriteProtection: efs.ReplicationOverwriteProtection_ENABLED,
//   })
//
// See: https://docs.aws.amazon.com/efs/latest/ug/performance.html#performancemodes
//
type PerformanceMode string

const (
	// General Purpose is ideal for latency-sensitive use cases, like web serving environments, content management systems, home directories, and general file serving.
	//
	// Recommended for the majority of Amazon EFS file systems.
	PerformanceMode_GENERAL_PURPOSE PerformanceMode = "GENERAL_PURPOSE"
	// File systems in the Max I/O mode can scale to higher levels of aggregate throughput and operations per second.
	//
	// This scaling is done with a tradeoff
	// of slightly higher latencies for file metadata operations.
	// Highly parallelized applications and workloads, such as big data analysis,
	// media processing, and genomics analysis, can benefit from this mode.
	PerformanceMode_MAX_IO PerformanceMode = "MAX_IO"
)

