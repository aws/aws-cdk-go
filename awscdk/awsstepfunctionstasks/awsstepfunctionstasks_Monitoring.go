package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Configuration setting for monitoring.
//
// Example:
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
//   	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
//   	jobDriver: &jobDriver{
//   		sparkSubmitJobDriver: &sparkSubmitJobDriver{
//   			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
//   		},
//   	},
//   	monitoring: &monitoring{
//   		logging: jsii.Boolean(true),
//   	},
//   })
//
type Monitoring struct {
	// Amazon S3 Bucket for monitoring log publishing.
	//
	// You can configure your jobs to send log information to Amazon S3.
	LogBucket awss3.IBucket `field:"optional" json:"logBucket" yaml:"logBucket"`
	// Enable logging for this job.
	//
	// If set to true, will automatically create a Cloudwatch Log Group and S3 bucket.
	// This will be set to `true` implicitly if values are provided for `logGroup` or `logBucket`.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// A log group for CloudWatch monitoring.
	//
	// You can configure your jobs to send log information to CloudWatch Logs.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// A log stream name prefix for Cloudwatch monitoring.
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// Monitoring configurations for the persistent application UI.
	PersistentAppUI *bool `field:"optional" json:"persistentAppUI" yaml:"persistentAppUI"`
}

