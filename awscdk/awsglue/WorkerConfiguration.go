package awsglue


// The worker configuration for a Spark job.
//
// The worker type and the number of workers are set together: providing this
// configuration requires both values, so a Spark job can never be given one
// without the other.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack Stack
//   var role IRole
//   var script Code
//
//   glue.NewPySparkEtlJob(stack, jsii.String("PySparkETLJob"), &PySparkEtlJobProps{
//   	JobName: jsii.String("PySparkETLJobCustomName"),
//   	Description: jsii.String("This is a description"),
//   	Role: Role,
//   	Script: Script,
//   	GlueVersion: glue.GlueVersion_V6_0,
//   	ContinuousLogging: &ContinuousLoggingProps{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	WorkerConfiguration: &WorkerConfiguration{
//   		WorkerType: glue.WorkerType_G_2X,
//   		NumberOfWorkers: jsii.Number(2),
//   	},
//   	MaxConcurrentRuns: jsii.Number(100),
//   	Timeout: cdk.Duration_Hours(jsii.Number(2)),
//   	Connections: []IConnection{
//   		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
//   	},
//   	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
//   	Tags: map[string]*string{
//   		"FirstTagName": jsii.String("FirstTagValue"),
//   		"SecondTagName": jsii.String("SecondTagValue"),
//   		"XTagName": jsii.String("XTagValue"),
//   	},
//   	MaxRetries: jsii.Number(2),
//   })
//
type WorkerConfiguration struct {
	// The number of workers of the given `workerType` that are allocated when a job runs.
	NumberOfWorkers *float64 `field:"required" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// See the `WorkerType` enum for the full set of supported values.
	WorkerType WorkerType `field:"required" json:"workerType" yaml:"workerType"`
}

