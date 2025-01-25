package awscdkgluealpha


// Python version.
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
type PythonVersion string

const (
	// Python 2 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_TWO PythonVersion = "TWO"
	// Python 3 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_THREE PythonVersion = "THREE"
	// Python 3.9 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_THREE_NINE PythonVersion = "THREE_NINE"
)

