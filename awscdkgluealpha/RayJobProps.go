package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Ray Glue job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code Code
//   var connection Connection
//   var logGroup LogGroup
//   var role Role
//   var securityConfiguration SecurityConfiguration
//
//   rayJobProps := &RayJobProps{
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
//   	GlueVersion: glue_alpha.GlueVersion_V0_9,
//   	JobName: jsii.String("jobName"),
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	MaxConcurrentRuns: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	Runtime: glue_alpha.Runtime_RAY_TWO_FOUR,
//   	SecurityConfiguration: securityConfiguration,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	WorkerType: glue_alpha.WorkerType_STANDARD,
//   }
//
// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
// Migrate to Amazon EKS with KubeRay Operator. See
// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
type RayJobProps struct {
	// IAM Role (required) IAM Role to use for Glue job execution Must be specified by the developer because the L2 doesn't have visibility into the actions the script(s) takes during the job execution The role must trust the Glue service principal (glue.amazonaws.com) and be granted sufficient permissions.
	// See: https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Script Code Location (required) Script to run when the Glue job executes.
	//
	// Can be uploaded
	// from the local directory structure using fromAsset
	// or referenced via S3 location using fromBucket.
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Script Code `field:"required" json:"script" yaml:"script"`
	// Connections (optional) List of connections to use for this Glue job Connections are used to connect to other AWS Service or resources within a VPC.
	// Default: [] - no connections are added to the job.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Connections *[]IConnection `field:"optional" json:"connections" yaml:"connections"`
	// Enables continuous logging with the specified props.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - continuous logging is enabled.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// Default Arguments (optional) The default arguments for every run of this Glue job, specified as name-value pairs.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	// for a list of reserved parameters.
	//
	// Default: - no arguments.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// Description (optional) Developer-specified description of the Glue job.
	// Default: - no value.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables the collection of metrics for job profiling.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no profiling metrics emitted.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// Glue Version The version of Glue to use to execute this job.
	// Default: 3.0 for ETL
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	GlueVersion GlueVersion `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Name of the Glue job (optional) Developer-specified name of the Glue job.
	// Default: - a name is automatically generated.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Max Concurrent Runs (optional) The maximum number of runs this Glue job can concurrently run.
	//
	// An error is returned when this threshold is reached. The maximum value
	// you can specify is controlled by a service limit.
	// Default: 1.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Max Retries (optional) Maximum number of retry attempts Glue performs if the job fails.
	// Default: 0.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Number of Workers (optional) Number of workers for Glue to use during job execution.
	// Default: 10.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Security Configuration (optional) Defines the encryption options for the Glue job.
	// Default: - no security configuration.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Glue job resources.
	// Default: {} - no tags.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Timeout (optional) The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	//
	// Specified in minutes.
	// Default: 2880 (2 days for non-streaming).
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Worker Type (optional) Type of Worker for Glue to use during job execution Enum options: Standard, G_1X, G_2X, G_025X.
	//
	// G_4X, G_8X, Z_2X.
	// Default: WorkerType.G_1X
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
	// Enable profiling metrics for the Glue job.
	//
	// When enabled, adds '--enable-metrics' to job arguments.
	// Default: true.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	EnableMetrics *bool `field:"optional" json:"enableMetrics" yaml:"enableMetrics"`
	// Enable observability metrics for the Glue job.
	//
	// When enabled, adds '--enable-observability-metrics': 'true' to job arguments.
	// Default: true.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	EnableObservabilityMetrics *bool `field:"optional" json:"enableObservabilityMetrics" yaml:"enableObservabilityMetrics"`
	// Specifies whether job run queuing is enabled for the job runs for this job.
	//
	// A value of true means job run queuing is enabled for the job runs.
	// If false or not populated, the job runs will not be considered for queueing.
	// If this field does not match the value set in the job run, then the value from
	// the job run field will be used. This property must be set to false for flex jobs.
	// If this property is enabled, maxRetries must be set to zero.
	// Default: - no job run queuing.
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	JobRunQueuingEnabled *bool `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
	// Sets the Ray runtime environment version.
	// Default: - Runtime version will default to Ray2.4
	//
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Runtime Runtime `field:"optional" json:"runtime" yaml:"runtime"`
}

