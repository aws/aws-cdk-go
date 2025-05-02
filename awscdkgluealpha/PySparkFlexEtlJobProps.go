package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for PySparkFlexEtlJob.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack stack
//   var role iRole
//   var script code
//
//   glue.NewPySparkFlexEtlJob(stack, jsii.String("ImportedJob"), &PySparkFlexEtlJobProps{
//   	Role: Role,
//   	Script: Script,
//   })
//
// Experimental.
type PySparkFlexEtlJobProps struct {
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
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - Spark UI debugging and monitoring is disabled.
	//
	// Experimental.
	SparkUI *SparkUIProps `field:"optional" json:"sparkUI" yaml:"sparkUI"`
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
	// Extra Python Files S3 URL (optional) S3 URL where additional python dependencies are located.
	// Default: - no extra files.
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// Specifies configuration properties of a notification (optional).
	//
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// Default: - undefined.
	//
	// Experimental.
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

