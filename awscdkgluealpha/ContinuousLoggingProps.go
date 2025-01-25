package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for enabling Continuous Logging for Glue Jobs.
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
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type ContinuousLoggingProps struct {
	// Enable continouous logging.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// Apply the provided conversion pattern.
	//
	// This is a Log4j Conversion Pattern to customize driver and executor logs.
	// Default: `%d{yy/MM/dd HH:mm:ss} %p %c{1}: %m%n`.
	//
	// Experimental.
	ConversionPattern *string `field:"optional" json:"conversionPattern" yaml:"conversionPattern"`
	// Specify a custom CloudWatch log group name.
	// Default: - a log group is created with name `/aws-glue/jobs/logs-v2/`.
	//
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	// Default: - the job run ID.
	//
	// Experimental.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	// Default: true.
	//
	// Experimental.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

