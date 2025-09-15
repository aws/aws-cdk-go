package awscdkgluealpha


// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
//
// A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack stack
//   var role iRole
//   var script code
//
//   glue.NewPythonShellJob(stack, jsii.String("PythonShellJob"), &PythonShellJobProps{
//   	JobName: jsii.String("PythonShellJobCustomName"),
//   	Description: jsii.String("This is a description"),
//   	PythonVersion: glue.PythonVersion_TWO,
//   	MaxCapacity: glue.MaxCapacity_DPU_1,
//   	Role: Role,
//   	Script: Script,
//   	GlueVersion: glue.GlueVersion_V2_0,
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
type MaxCapacity string

const (
	// DPU value of 1/16th.
	// Experimental.
	MaxCapacity_DPU_1_16TH MaxCapacity = "DPU_1_16TH"
	// DPU value of 1.
	// Experimental.
	MaxCapacity_DPU_1 MaxCapacity = "DPU_1"
)

