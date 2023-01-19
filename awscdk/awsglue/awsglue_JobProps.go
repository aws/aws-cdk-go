package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for {@link Job}.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type JobProps struct {
	// The job's executable properties.
	// Experimental.
	Executable JobExecutable `field:"required" json:"executable" yaml:"executable"`
	// The {@link Connection}s used for this job.
	//
	// Connections are used to connect to other AWS Service or resources within a VPC.
	// Experimental.
	Connections *[]IConnection `field:"optional" json:"connections" yaml:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// The default arguments for this job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html for a list of reserved parameters
	//
	// Experimental.
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// The description of the job.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables the collection of metrics for job profiling.
	// See: `--enable-metrics` at https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// The name of the job.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// Cannot be used for Glue version 2.0 and later - workerType and workerCount should be used instead.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of concurrent runs allowed for the job.
	//
	// An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	// Experimental.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The maximum number of times to retry this job after a job run fails.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The number of minutes to wait after a job run starts, before sending a job run delay notification.
	// Experimental.
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The IAM role assumed by Glue to run this job.
	//
	// If providing a custom role, it needs to trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The {@link SecurityConfiguration} to use for this job.
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Enables the Spark UI debugging and monitoring with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	SparkUI *SparkUIProps `field:"optional" json:"sparkUI" yaml:"sparkUI"`
	// The tags to add to the resources on which the job runs.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of workers of a defined {@link WorkerType} that are allocated when a job runs.
	// Experimental.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
	// The type of predefined worker that is allocated when a job runs.
	// Experimental.
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
}

