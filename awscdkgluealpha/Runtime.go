package awscdkgluealpha


// AWS Glue runtime determines the runtime engine of the job.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack stack
//   var role iRole
//   var script code
//
//   glue.NewRayJob(stack, jsii.String("ImportedJob"), &RayJobProps{
//   	Role: Role,
//   	Script: Script,
//   	JobName: jsii.String("RayCustomJobName"),
//   	Description: jsii.String("This is a description"),
//   	WorkerType: glue.WorkerType_Z_2X,
//   	NumberOfWorkers: jsii.Number(5),
//   	Runtime: glue.Runtime_RAY_TWO_FOUR,
//   	MaxRetries: jsii.Number(3),
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
//   })
//
// Experimental.
type Runtime string

const (
	// Runtime for a Glue for Ray 2.4.
	// Experimental.
	Runtime_RAY_TWO_FOUR Runtime = "RAY_TWO_FOUR"
)

