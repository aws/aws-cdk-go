package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for `Job`.
//
// Example:
//   glue.NewJob(this, jsii.String("EnableSparkUI"), &JobProps{
//   	JobName: jsii.String("EtlJobWithSparkUIPrefix"),
//   	SparkUI: &SparkUIProps{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	Executable: glue.JobExecutable_PythonEtl(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V3_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   })
//
// Experimental.
type JobProps struct {
	// The job's executable properties.
	// Experimental.
	Executable JobExecutable `field:"required" json:"executable" yaml:"executable"`
	// The `Connection`s used for this job.
	//
	// Connections are used to connect to other AWS Service or resources within a VPC.
	// Default: [] - no connections are added to the job.
	//
	// Experimental.
	Connections *[]IConnection `field:"optional" json:"connections" yaml:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - continuous logging is disabled.
	//
	// Experimental.
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// The default arguments for this job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html for a list of reserved parameters
	//
	// Default: - no arguments.
	//
	// Experimental.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// The description of the job.
	// Default: - no value.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables the collection of metrics for job profiling.
	// See:  `--enable-metrics` at https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no profiling metrics emitted.
	//
	// Experimental.
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// The ExecutionClass whether the job is run with a standard or flexible execution class.
	// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html
	//
	// Default: - STANDARD.
	//
	// Experimental.
	ExecutionClass ExecutionClass `field:"optional" json:"executionClass" yaml:"executionClass"`
	// The name of the job.
	// Default: - a name is automatically generated.
	//
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// Cannot be used for Glue version 2.0 and later - workerType and workerCount should be used instead.
	// Default: - 10 when job type is Apache Spark ETL or streaming, 0.0625 when job type is Python shell
	//
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of concurrent runs allowed for the job.
	//
	// An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	// Default: 1.
	//
	// Experimental.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The maximum number of times to retry this job after a job run fails.
	// Default: 0.
	//
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The number of minutes to wait after a job run starts, before sending a job run delay notification.
	// Default: - no delay notifications.
	//
	// Experimental.
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The IAM role assumed by Glue to run this job.
	//
	// If providing a custom role, it needs to trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Default: - a role is automatically generated.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The `SecurityConfiguration` to use for this job.
	// Default: - no security configuration.
	//
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - Spark UI debugging and monitoring is disabled.
	//
	// Experimental.
	SparkUI *SparkUIProps `field:"optional" json:"sparkUI" yaml:"sparkUI"`
	// The tags to add to the resources on which the job runs.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Default: cdk.Duration.hours(48)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of workers of a defined `WorkerType` that are allocated when a job runs.
	// Default: - differs based on specific Glue version/worker type.
	//
	// Experimental.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
	// The type of predefined worker that is allocated when a job runs.
	// Default: - differs based on specific Glue version.
	//
	// Experimental.
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
}

