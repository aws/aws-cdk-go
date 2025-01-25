package awscdkgluealpha


// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`
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
	// G.025X Worker Type 0.25 DPU (2 vCPU, 4 GB of memory, 64 GB disk), and provides 1 executor per worker. Suitable for low volume streaming jobs.
	// Experimental.
	WorkerType_G_025X WorkerType = "G_025X"
	// Z.2X Worker Type.
	// Experimental.
	WorkerType_Z_2X WorkerType = "Z_2X"
)

