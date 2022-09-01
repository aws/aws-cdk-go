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
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
//   	encryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
//   	encryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: kmsKey,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchLogGroup: logGroup,
//   			cloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			s3Bucket: execBucket,
//   			s3EncryptionEnabled: jsii.Boolean(true),
//   			s3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		logging: ecs.executeCommandLogging_OVERRIDE,
//   	},
//   })
//
type ExecuteCommandConfiguration struct {
	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
	LogConfiguration *ExecuteCommandLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The log settings to use for logging the execute command session.
	Logging ExecuteCommandLogging `field:"optional" json:"logging" yaml:"logging"`
}

