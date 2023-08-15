package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The details of the execute command configuration.
//
// For more information, see
// [ExecuteCommandConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html
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
type ExecuteCommandConfiguration struct {
	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	// Default: - none.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
	// Default: - none.
	//
	LogConfiguration *ExecuteCommandLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The log settings to use for logging the execute command session.
	// Default: - none.
	//
	Logging ExecuteCommandLogging `field:"optional" json:"logging" yaml:"logging"`
}

