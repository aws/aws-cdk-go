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
//   cfnJobProps := &CfnJobProps{
//   	Command: &JobCommandProperty{
//   		Name: jsii.String("name"),
//   		PythonVersion: jsii.String("pythonVersion"),
//   		Runtime: jsii.String("runtime"),
//   		ScriptLocation: jsii.String("scriptLocation"),
//   	},
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	AllocatedCapacity: jsii.Number(123),
//   	Connections: &ConnectionsListProperty{
//   		Connections: []*string{
//   			jsii.String("connections"),
//   		},
//   	},
//   	DefaultArguments: defaultArguments,
//   	Description: jsii.String("description"),
//   	ExecutionClass: jsii.String("executionClass"),
//   	ExecutionProperty: &ExecutionPropertyProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   	},
//   	GlueVersion: jsii.String("glueVersion"),
//   	JobMode: jsii.String("jobMode"),
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	LogUri: jsii.String("logUri"),
//   	MaintenanceWindow: jsii.String("maintenanceWindow"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NonOverridableArguments: nonOverridableArguments,
//   	NotificationProperty: &NotificationPropertyProperty{
//   		NotifyDelayAfter: jsii.Number(123),
//   	},
//   	NumberOfWorkers: jsii.Number(123),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	Tags: tags,
//   	Timeout: jsii.Number(123),
//   	WorkerType: jsii.String("workerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html
//
type CfnJobProps struct {
	// The code that executes a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-command
	//
	Command interface{} `field:"required" json:"command" yaml:"command"`
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// This parameter is no longer supported. Use `MaxCapacity` instead.
	//
	// The number of capacity units that are allocated to this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-allocatedcapacity
	//
	AllocatedCapacity *float64 `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// The connections used for this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-connections
	//
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// The default arguments for this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) in the *AWS Glue Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-defaultarguments
	//
	DefaultArguments interface{} `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// A description of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the job is run with a standard or flexible execution class.
	//
	// The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources.
	//
	// The flexible execution class is appropriate for time-insensitive jobs whose start and completion times may vary.
	//
	// Only jobs with AWS Glue version 3.0 and above and command type `glueetl` will be allowed to set `ExecutionClass` to `FLEX` . The flexible execution class is available for Spark jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-executionclass
	//
	ExecutionClass *string `field:"optional" json:"executionClass" yaml:"executionClass"`
	// The maximum number of concurrent runs that are allowed for this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-executionproperty
	//
	ExecutionProperty interface{} `field:"optional" json:"executionProperty" yaml:"executionProperty"`
	// Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for jobs of type Spark.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to the latest Glue version available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-glueversion
	//
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// A mode that describes how a job was created. Valid values are:.
	//
	// - `SCRIPT` - The job was created using the AWS Glue Studio script editor.
	// - `VISUAL` - The job was created using the AWS Glue Studio visual editor.
	// - `NOTEBOOK` - The job was created using an interactive sessions notebook.
	//
	// When the `JobMode` field is missing or null, `SCRIPT` is assigned as the default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-jobmode
	//
	JobMode *string `field:"optional" json:"jobMode" yaml:"jobMode"`
	// Specifies whether job run queuing is enabled for the job runs for this job.
	//
	// A value of true means job run queuing is enabled for the job runs. If false or not populated, the job runs will not be considered for queueing.
	//
	// If this field does not match the value set in the job run, then the value from the job run field will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-jobrunqueuingenabled
	//
	JobRunQueuingEnabled interface{} `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
	// This field is reserved for future use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// This field specifies a day of the week and hour for a maintenance window for streaming jobs.
	//
	// AWS Glue periodically performs maintenance activities. During these maintenance windows, AWS Glue will need to restart your streaming jobs.
	//
	// AWS Glue will restart the job within 3 hours of the specified maintenance window. For instance, if you set up the maintenance window for Monday at 10:00AM GMT, your jobs will be restarted between 10:00AM GMT to 1:00PM GMT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-maintenancewindow
	//
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry this job after a JobRun fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-maxretries
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The name you assign to this job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-nonoverridablearguments
	//
	NonOverridableArguments interface{} `field:"optional" json:"nonOverridableArguments" yaml:"nonOverridableArguments"`
	// Specifies configuration properties of a notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-notificationproperty
	//
	NotificationProperty interface{} `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// The number of workers of a defined `workerType` that are allocated when a job runs.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-numberofworkers
	//
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The name of the `SecurityConfiguration` structure to be used with this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-securityconfiguration
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags to use with this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The job timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// Accepts a value of G.1X, G.2X, G.4X, G.8X or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
	//
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPUs, 16 GB of memory) with 94GB disk, and provides 1 executor per worker. We recommend this worker type for workloads such as data transforms, joins, and queries, to offers a scalable and cost effective way to run most jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPUs, 32 GB of memory) with 138GB disk, and provides 1 executor per worker. We recommend this worker type for workloads such as data transforms, joins, and queries, to offers a scalable and cost effective way to run most jobs.
	// - For the `G.4X` worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB of memory) with 256GB disk, and provides 1 executor per worker. We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later Spark ETL jobs in the following AWS Regions: US East (Ohio), US East (N. Virginia), US West (N. California), US West (Oregon), Asia Pacific (Mumbai), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), Canada (Central), Europe (Frankfurt), Europe (Ireland), Europe (London), Europe (Spain), Europe (Stockholm), and South America (São Paulo).
	// - For the `G.8X` worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB of memory) with 512GB disk, and provides 1 executor per worker. We recommend this worker type for jobs whose workloads contain your most demanding transforms, aggregations, joins, and queries. This worker type is available only for AWS Glue version 3.0 or later Spark ETL jobs, in the same AWS Regions as supported for the `G.4X` worker type.
	// - For the `G.025X` worker type, each worker maps to 0.25 DPU (2 vCPUs, 4 GB of memory) with 84GB disk, and provides 1 executor per worker. We recommend this worker type for low volume streaming jobs. This worker type is only available for AWS Glue version 3.0 or later streaming jobs.
	// - For the `Z.2X` worker type, each worker maps to 2 M-DPU (8vCPUs, 64 GB of memory) with 128 GB disk, and provides up to 8 Ray workers based on the autoscaler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-workertype
	//
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

