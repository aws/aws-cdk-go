package awsefs


// EFS Performance mode.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &fileSystemProps{
//   	vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	lifecyclePolicy: efs.lifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	performanceMode: efs.performanceMode_GENERAL_PURPOSE,
//   	 // default
//   	outOfInfrequentAccessPolicy: efs.outOfInfrequentAccessPolicy_AFTER_1_ACCESS,
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

