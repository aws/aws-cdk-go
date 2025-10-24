package awscdkkinesisanalyticsflinkalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Props for creating an Application construct.
//
// Example:
//   var bucket Bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
//   	PropertyGroups: map[string]map[string]*string{
//   		"FlinkApplicationProperties": map[string]*string{
//   			"inputStreamName": jsii.String("my-input-kinesis-stream"),
//   			"outputStreamName": jsii.String("my-output-kinesis-stream"),
//   		},
//   	},
//   	// ...
//   	Runtime: flink.Runtime_FLINK_1_20(),
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
//   })
//
// Experimental.
type ApplicationProps struct {
	// The Flink code asset to run.
	// Experimental.
	Code ApplicationCode `field:"required" json:"code" yaml:"code"`
	// The Flink version to use for this application.
	// Experimental.
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// A name for your Application that is unique to an AWS account.
	// Default: - CloudFormation-generated name.
	//
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Whether the Kinesis Data Analytics service can increase the parallelism of the application in response to resource usage.
	// Default: true.
	//
	// Experimental.
	AutoScalingEnabled *bool `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Whether checkpointing is enabled while your application runs.
	// Default: true.
	//
	// Experimental.
	CheckpointingEnabled *bool `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// The interval between checkpoints.
	// Default: - 1 minute.
	//
	// Experimental.
	CheckpointInterval awscdk.Duration `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// The log group to send log entries to.
	// Default: - CDK's default LogGroup.
	//
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The level of log verbosity from the Flink application.
	// Default: FlinkLogLevel.INFO
	//
	// Experimental.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Describes the granularity of the CloudWatch metrics for an application.
	//
	// Use caution with Parallelism level metrics. Parallelism granularity logs
	// metrics for each parallel thread and can quickly become expensive when
	// parallelism is high (e.g. > 64).
	// Default: MetricsLevel.APPLICATION
	//
	// Experimental.
	MetricsLevel MetricsLevel `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
	// The minimum amount of time in to wait after a checkpoint finishes to start a new checkpoint.
	// Default: - 5 seconds.
	//
	// Experimental.
	MinPauseBetweenCheckpoints awscdk.Duration `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
	// The initial parallelism for the application.
	//
	// Kinesis Data Analytics can
	// stop the app, increase the parallelism, and start the app again if
	// autoScalingEnabled is true (the default value).
	// Default: 1.
	//
	// Experimental.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// The Flink parallelism allowed per Kinesis Processing Unit (KPU).
	// Default: 1.
	//
	// Experimental.
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
	// Configuration PropertyGroups.
	//
	// You can use these property groups to pass
	// arbitrary runtime configuration values to your Flink app.
	// Default: - No property group configuration provided to the Flink app.
	//
	// Experimental.
	PropertyGroups *map[string]*map[string]*string `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// Provide a RemovalPolicy to override the default.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// A role to use to grant permissions to your application.
	//
	// Prefer omitting
	// this property and using the default role.
	// Default: - a new Role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to use with a provided VPC.
	// Default: - a new security group is created for this application.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Determines if Flink snapshots are enabled.
	// Default: true.
	//
	// Experimental.
	SnapshotsEnabled *bool `field:"optional" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
	// Deploy the Flink application in a VPC.
	// Default: - no VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Choose which VPC subnets to use.
	// Default: - SubnetType.PRIVATE_WITH_EGRESS subnets
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

