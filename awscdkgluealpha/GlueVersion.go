package awscdkgluealpha


// AWS Glue version determines the versions of Apache Spark and Python that are available to the job.
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
//   	GlueVersion: glue.GlueVersion_V5_1,
//   	ContinuousLogging: &ContinuousLoggingProps{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	WorkerType: glue.WorkerType_G_2X,
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
//   	NumberOfWorkers: jsii.Number(2),
//   	MaxRetries: jsii.Number(2),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html.
//
// Experimental.
type GlueVersion string

const (
	// Glue version using Spark 2.2.1 and Python 2.7.
	// Experimental.
	GlueVersion_V0_9 GlueVersion = "V0_9"
	// Glue version using Spark 2.4.3, Python 2.7 and Python 3.6.
	// Experimental.
	GlueVersion_V1_0 GlueVersion = "V1_0"
	// Glue version using Spark 2.4.3 and Python 3.7.
	// Experimental.
	GlueVersion_V2_0 GlueVersion = "V2_0"
	// Glue version using Spark 3.1.1 and Python 3.7.
	// Experimental.
	GlueVersion_V3_0 GlueVersion = "V3_0"
	// Glue version using Spark 3.3.0 and Python 3.10.
	// Experimental.
	GlueVersion_V4_0 GlueVersion = "V4_0"
	// Glue version using Spark 3.5.4, Python 3.11, and Scala 2.12.18.
	// Experimental.
	GlueVersion_V5_0 GlueVersion = "V5_0"
	// Glue version using Spark 3.5.6, Python 3.11, and Scala 2.12.18.
	// Experimental.
	GlueVersion_V5_1 GlueVersion = "V5_1"
)

