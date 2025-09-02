package awscdkgluealpha


// The type of predefined worker that is allocated when a job runs.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack stack
//   var role iRole
//   var script code
//
//   glue.NewPySparkEtlJob(stack, jsii.String("PySparkETLJob"), &PySparkEtlJobProps{
//   	JobName: jsii.String("PySparkETLJobCustomName"),
//   	Description: jsii.String("This is a description"),
//   	Role: Role,
//   	Script: Script,
//   	GlueVersion: glue.GlueVersion_V3_0,
//   	ContinuousLogging: &ContinuousLoggingProps{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	WorkerType: glue.WorkerType_G_2X,
//   	MaxConcurrentRuns: jsii.Number(100),
//   	Timeout: cdk.Duration_Hours(jsii.Number(2)),
//   	Connections: []iConnection{
//   		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
//   	},
//   	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
//   	Tags: map[string]*string{
//   		"FirstTagName": jsii.String("FirstTagValue"),
//   		"SecondTagName": jsii.String("SecondTagValue"),
//   		"XTagName": jsii.String("XTagValue"),
//   	},
//   	NumberOfWorkers: jsii.Number(2),
//   	MaxRetries: jsii.Number(2),
//   })
//
// Experimental.
type WorkerType string

const (
	// Standard Worker Type 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// Experimental.
	WorkerType_STANDARD WorkerType = "STANDARD"
	// G.1X Worker Type 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Suitable for memory-intensive jobs.
	// Experimental.
	WorkerType_G_1X WorkerType = "G_1X"
	// G.2X Worker Type 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Suitable for memory-intensive jobs.
	// Experimental.
	WorkerType_G_2X WorkerType = "G_2X"
	// G.4X Worker Type 4 DPU (16 vCPU, 64 GB of memory, 256 GB disk), and provides 1 executor per worker. We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later jobs.
	// Experimental.
	WorkerType_G_4X WorkerType = "G_4X"
	// G.8X Worker Type 8 DPU (32 vCPU, 128 GB of memory, 512 GB disk), and provides 1 executor per worker. We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later jobs.
	// Experimental.
	WorkerType_G_8X WorkerType = "G_8X"
	// G.12X Worker Type 12 DPU (48 vCPU, 192 GB of memory, 768 GB disk), and provides 1 executor per worker. We recommend this worker type for jobs with very large and resource-intensive workloads that require significant compute capacity. This worker type is available only for AWS Glue version 3.0 or later jobs.
	// Experimental.
	WorkerType_G_12X WorkerType = "G_12X"
	// G.16X Worker Type 16 DPU (64 vCPU, 256 GB of memory, 1024 GB disk), and provides 1 executor per worker. We recommend this worker type for jobs with the largest and most resource-intensive workloads that require maximum compute capacity. This worker type is available only for AWS Glue version 3.0 or later jobs.
	// Experimental.
	WorkerType_G_16X WorkerType = "G_16X"
	// G.025X Worker Type 0.25 DPU (2 vCPU, 4 GB of memory, 64 GB disk), and provides 1 executor per worker. Suitable for low volume streaming jobs.
	// Experimental.
	WorkerType_G_025X WorkerType = "G_025X"
	// Z.2X Worker Type.
	// Experimental.
	WorkerType_Z_2X WorkerType = "Z_2X"
	// R.1X Worker Type 1 M-DPU (4 vCPUs, 32 GB memory), We recommend this worker type for memory-intensive workloads that frequently encounter out-of-memory errors or require high memory-to-CPU ratios.
	// Experimental.
	WorkerType_R_1X WorkerType = "R_1X"
	// R.2X Worker Type 2 M-DPU (8 vCPUs, 64 GB memory), We recommend this worker type for memory-intensive workloads that frequently encounter out-of-memory errors or require high memory-to-CPU ratios.
	// Experimental.
	WorkerType_R_2X WorkerType = "R_2X"
	// R.4X Worker Type 4 M-DPU (16 vCPUs, 128 GB memory), We recommend this worker type for large memory-intensive workloads that frequently encounter out-of-memory errors or require high memory-to-CPU ratios.
	// Experimental.
	WorkerType_R_4X WorkerType = "R_4X"
	// R.8X Worker Type 8 M-DPU (32 vCPUs, 256 GB memory), We recommend this worker type for very large memory-intensive workloads that frequently encounter out-of-memory errors or require high memory-to-CPU ratios.
	// Experimental.
	WorkerType_R_8X WorkerType = "R_8X"
)

