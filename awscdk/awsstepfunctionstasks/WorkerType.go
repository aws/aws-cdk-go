package awsstepfunctionstasks


// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`.
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
//   	GlueJobName: jsii.String("my-glue-job"),
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		WorkerType: tasks.WorkerType_G_1X,
//   		 // Worker type
//   		NumberOfWorkers: jsii.Number(2),
//   	},
//   })
//
type WorkerType string

const (
	// Each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	WorkerType_STANDARD WorkerType = "STANDARD"
	// Each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker.
	//
	// Suitable for memory-intensive jobs.
	WorkerType_G_1X WorkerType = "G_1X"
	// Each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker.
	//
	// Suitable for memory-intensive jobs.
	WorkerType_G_2X WorkerType = "G_2X"
	// Each worker maps to 4 DPU (16 vCPU, 64 GB of memory, 256 GB disk), and provides 1 executor per worker.
	//
	// We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later jobs.
	WorkerType_G_4X WorkerType = "G_4X"
	// Each worker maps to 8 DPU (32 vCPU, 128 GB of memory, 512 GB disk), and provides 1 executor per worker.
	//
	// We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later jobs.
	WorkerType_G_8X WorkerType = "G_8X"
	// Each worker maps to 0.25 DPU (2 vCPU, 4 GB of memory, 64 GB disk), and provides 1 executor per worker. Suitable for low volume streaming jobs.
	WorkerType_G_025X WorkerType = "G_025X"
	// Each worker maps to 2 high-memory DPU [M-DPU] (8 vCPU, 64 GB of memory, 128 GB disk).
	//
	// Supported in Ray jobs.
	WorkerType_Z_2X WorkerType = "Z_2X"
)

