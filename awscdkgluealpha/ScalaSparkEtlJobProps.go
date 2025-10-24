package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Scala Spark ETL job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//   var code Code
//   var connection Connection
//   var logGroup LogGroup
//   var role Role
//   var securityConfiguration SecurityConfiguration
//
//   scalaSparkEtlJobProps := &ScalaSparkEtlJobProps{
//   	ClassName: jsii.String("className"),
//   	Role: role,
//   	Script: code,
//
//   	// the properties below are optional
//   	Connections: []IConnection{
//   		connection,
//   	},
//   	ContinuousLogging: &ContinuousLoggingProps{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ConversionPattern: jsii.String("conversionPattern"),
//   		LogGroup: logGroup,
//   		LogStreamPrefix: jsii.String("logStreamPrefix"),
//   		Quiet: jsii.Boolean(false),
//   	},
//   	DefaultArguments: map[string]*string{
//   		"defaultArgumentsKey": jsii.String("defaultArguments"),
//   	},
//   	Description: jsii.String("description"),
//   	EnableMetrics: jsii.Boolean(false),
//   	EnableObservabilityMetrics: jsii.Boolean(false),
//   	EnableProfilingMetrics: jsii.Boolean(false),
//   	ExtraFiles: []Code{
//   		code,
//   	},
//   	ExtraJars: []Code{
//   		code,
//   	},
//   	ExtraJarsFirst: jsii.Boolean(false),
//   	GlueVersion: glue_alpha.GlueVersion_V0_9,
//   	JobName: jsii.String("jobName"),
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	MaxConcurrentRuns: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	SecurityConfiguration: securityConfiguration,
//   	SparkUI: &SparkUIProps{
//   		Bucket: bucket,
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	WorkerType: glue_alpha.WorkerType_STANDARD,
//   }
//
// Experimental.
type ScalaSparkEtlJobProps struct {
	// IAM Role (required) IAM Role to use for Glue job execution Must be specified by the developer because the L2 doesn't have visibility into the actions the script(s) takes during the job execution The role must trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Script Code Location (required) Script to run when the Glue job executes.
	//
	// Can be uploaded
	// from the local directory structure using fromAsset
	// or referenced via S3 location using fromBucket.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Connections (optional) List of connections to use for this Glue job Connections are used to connect to other AWS Service or resources within a VPC.
	// Default: [] - no connections are added to the job.
	//
	// Experimental.
	Connections *[]IConnection `field:"optional" json:"connections" yaml:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - continuous logging is enabled.
	//
	// Experimental.
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// Default Arguments (optional) The default arguments for every run of this Glue job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	// for a list of reserved parameters.
	//
	// Default: - no arguments.
	//
	// Experimental.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// Description (optional) Developer-specified description of the Glue job.
	// Default: - no value.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables the collection of metrics for job profiling.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no profiling metrics emitted.
	//
	// Experimental.
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// Glue Version The version of Glue to use to execute this job.
	// Default: 3.0 for ETL
	//
	// Experimental.
	GlueVersion GlueVersion `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Name of the Glue job (optional) Developer-specified name of the Glue job.
	// Default: - a name is automatically generated.
	//
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Max Concurrent Runs (optional) The maximum number of runs this Glue job can concurrently run.
	//
	// An error is returned when this threshold is reached. The maximum value
	// you can specify is controlled by a service limit.
	// Default: 1.
	//
	// Experimental.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Max Retries (optional) Maximum number of retry attempts Glue performs if the job fails.
	// Default: 0.
	//
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Number of Workers (optional) Number of workers for Glue to use during job execution.
	// Default: 10.
	//
	// Experimental.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Security Configuration (optional) Defines the encryption options for the Glue job.
	// Default: - no security configuration.
	//
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Glue job resources.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Timeout (optional) The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	//
	// Specified in minutes.
	// Default: 2880 (2 days for non-streaming).
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Worker Type (optional) Type of Worker for Glue to use during job execution Enum options: Standard, G_1X, G_2X, G_025X.
	//
	// G_4X, G_8X, Z_2X.
	// Default: WorkerType.G_1X
	//
	// Experimental.
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
	// Enable profiling metrics for the Glue job.
	//
	// When enabled, adds '--enable-metrics' to job arguments.
	// Default: true.
	//
	// Experimental.
	EnableMetrics *bool `field:"optional" json:"enableMetrics" yaml:"enableMetrics"`
	// Enable observability metrics for the Glue job.
	//
	// When enabled, adds '--enable-observability-metrics': 'true' to job arguments.
	// Default: true.
	//
	// Experimental.
	EnableObservabilityMetrics *bool `field:"optional" json:"enableObservabilityMetrics" yaml:"enableObservabilityMetrics"`
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - Spark UI debugging and monitoring is disabled.
	//
	// Experimental.
	SparkUI *SparkUIProps `field:"optional" json:"sparkUI" yaml:"sparkUI"`
	// Class name (required for Scala scripts) Package and class name for the entry point of Glue job execution for Java scripts.
	// Experimental.
	ClassName *string `field:"required" json:"className" yaml:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra files specified.
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Extra Jars S3 URL (optional) S3 URL where additional jar dependencies are located.
	// Default: - no extra jar files.
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See:  `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: false - priority is not given to user-provided jars.
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Specifies whether job run queuing is enabled for the job runs for this job.
	//
	// A value of true means job run queuing is enabled for the job runs.
	// If false or not populated, the job runs will not be considered for queueing.
	// If this field does not match the value set in the job run, then the value from
	// the job run field will be used. This property must be set to false for flex jobs.
	// If this property is enabled, maxRetries must be set to zero.
	// Default: - no job run queuing.
	//
	// Experimental.
	JobRunQueuingEnabled *bool `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
}

