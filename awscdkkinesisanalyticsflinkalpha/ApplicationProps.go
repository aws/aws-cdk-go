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
//   import path "github.com/aws-samples/dummy/path"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"
//
//   app := core.NewApp()
//   stack := core.NewStack(app, jsii.String("FlinkAppTest"))
//
//   flinkApp := flink.NewApplication(stack, jsii.String("App"), &ApplicationProps{
//   	Code: flink.ApplicationCode_FromAsset(path.join(__dirname, jsii.String("code-asset"))),
//   	Runtime: flink.Runtime_FLINK_1_11(),
//   })
//
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &AlarmProps{
//   	Metric: flinkApp.metricFullRestarts(),
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(3),
//   })
//
//   app.Synth()
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
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Whether the Kinesis Data Analytics service can increase the parallelism of the application in response to resource usage.
	// Experimental.
	AutoScalingEnabled *bool `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Whether checkpointing is enabled while your application runs.
	// Experimental.
	CheckpointingEnabled *bool `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// The interval between checkpoints.
	// Experimental.
	CheckpointInterval awscdk.Duration `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// The log group to send log entries to.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The level of log verbosity from the Flink application.
	// Experimental.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Describes the granularity of the CloudWatch metrics for an application.
	//
	// Use caution with Parallelism level metrics. Parallelism granularity logs
	// metrics for each parallel thread and can quickly become expensive when
	// parallelism is high (e.g. > 64).
	// Experimental.
	MetricsLevel MetricsLevel `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
	// The minimum amount of time in to wait after a checkpoint finishes to start a new checkpoint.
	// Experimental.
	MinPauseBetweenCheckpoints awscdk.Duration `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
	// The initial parallelism for the application.
	//
	// Kinesis Data Analytics can
	// stop the app, increase the parallelism, and start the app again if
	// autoScalingEnabled is true (the default value).
	// Experimental.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// The Flink parallelism allowed per Kinesis Processing Unit (KPU).
	// Experimental.
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
	// Configuration PropertyGroups.
	//
	// You can use these property groups to pass
	// arbitrary runtime configuration values to your Flink app.
	// Experimental.
	PropertyGroups *map[string]*map[string]*string `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// Provide a RemovalPolicy to override the default.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// A role to use to grant permissions to your application.
	//
	// Prefer omitting
	// this property and using the default role.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to use with a provided VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Determines if Flink snapshots are enabled.
	// Experimental.
	SnapshotsEnabled *bool `field:"optional" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
	// Deploy the Flink application in a VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Choose which VPC subnets to use.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

