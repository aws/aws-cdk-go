package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs and/ or an Amazon S3 bucket.
// For more information, see [ExecuteCommandLogConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandlogconfiguration.html
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
//   	EncryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &BucketProps{
//   	EncryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   	ExecuteCommandConfiguration: &ExecuteCommandConfiguration{
//   		KmsKey: *KmsKey,
//   		LogConfiguration: &ExecuteCommandLogConfiguration{
//   			CloudWatchLogGroup: logGroup,
//   			CloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			S3Bucket: execBucket,
//   			S3EncryptionEnabled: jsii.Boolean(true),
//   			S3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		Logging: ecs.ExecuteCommandLogging_OVERRIDE,
//   	},
//   })
//
type ExecuteCommandLogConfiguration struct {
	// Whether or not to enable encryption on the CloudWatch logs.
	// Default: - encryption will be disabled.
	//
	CloudWatchEncryptionEnabled *bool `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// The CloudWatch log group must already be created.
	// Default: - none.
	//
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// The name of the S3 bucket to send logs to.
	//
	// The S3 bucket must already be created.
	// Default: - none.
	//
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Whether or not to enable encryption on the S3 bucket.
	// Default: - encryption will be disabled.
	//
	S3EncryptionEnabled *bool `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	// Default: - none.
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

