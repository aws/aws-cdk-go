package awsglue


// Properties for defining a `CfnJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var defaultArguments interface{}
//   var nonOverridableArguments interface{}
//   var tags interface{}
//
//   cfnJobProps := &cfnJobProps{
//   	command: &jobCommandProperty{
//   		name: jsii.String("name"),
//   		pythonVersion: jsii.String("pythonVersion"),
//   		scriptLocation: jsii.String("scriptLocation"),
//   	},
//   	role: jsii.String("role"),
//
//   	// the properties below are optional
//   	allocatedCapacity: jsii.Number(123),
//   	connections: &connectionsListProperty{
//   		connections: []*string{
//   			jsii.String("connections"),
//   		},
//   	},
//   	defaultArguments: defaultArguments,
//   	description: jsii.String("description"),
//   	executionClass: jsii.String("executionClass"),
//   	executionProperty: &executionPropertyProperty{
//   		maxConcurrentRuns: jsii.Number(123),
//   	},
//   	glueVersion: jsii.String("glueVersion"),
//   	logUri: jsii.String("logUri"),
//   	maxCapacity: jsii.Number(123),
//   	maxRetries: jsii.Number(123),
//   	name: jsii.String("name"),
//   	nonOverridableArguments: nonOverridableArguments,
//   	notificationProperty: &notificationPropertyProperty{
//   		notifyDelayAfter: jsii.Number(123),
//   	},
//   	numberOfWorkers: jsii.Number(123),
//   	securityConfiguration: jsii.String("securityConfiguration"),
//   	tags: tags,
//   	timeout: jsii.Number(123),
//   	workerType: jsii.String("workerType"),
//   }
//
type CfnJobProps struct {
	// The code that executes a job.
	Command interface{} `field:"required" json:"command" yaml:"command"`
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	Role *string `field:"required" json:"role" yaml:"role"`
	// The number of capacity units that are allocated to this job.
	AllocatedCapacity *float64 `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// The connections used for this job.
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// The default arguments for this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) in the *AWS Glue Developer Guide* .
	DefaultArguments interface{} `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// A description of the job.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Glue::Job.ExecutionClass`.
	ExecutionClass *string `field:"optional" json:"executionClass" yaml:"executionClass"`
	// The maximum number of concurrent runs that are allowed for this job.
	ExecutionProperty interface{} `field:"optional" json:"executionProperty" yaml:"executionProperty"`
	// Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for jobs of type Spark.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// This field is reserved for future use.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	//
	// Do not set `Max Capacity` if using `WorkerType` and `NumberOfWorkers` .
	//
	// The value that can be allocated for `MaxCapacity` depends on whether you are running a Python shell job or an Apache Spark ETL job:
	//
	// - When you specify a Python shell job ( `JobCommand.Name` ="pythonshell"), you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	// - When you specify an Apache Spark ETL job ( `JobCommand.Name` ="glueetl"), you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU allocation.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry this job after a JobRun fails.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The name you assign to this job definition.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Glue::Job.NonOverridableArguments`.
	NonOverridableArguments interface{} `field:"optional" json:"nonOverridableArguments" yaml:"nonOverridableArguments"`
	// Specifies configuration properties of a notification.
	NotificationProperty interface{} `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// The number of workers of a defined `workerType` that are allocated when a job runs.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The name of the `SecurityConfiguration` structure to be used with this job.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags to use with this job.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The job timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

