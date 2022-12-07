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
type ExecuteCommandLogConfiguration struct {
	// Whether or not to enable encryption on the CloudWatch logs.
	CloudWatchEncryptionEnabled *bool `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// The CloudWatch log group must already be created.
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// The name of the S3 bucket to send logs to.
	//
	// The S3 bucket must already be created.
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Whether or not to enable encryption on the CloudWatch logs.
	S3EncryptionEnabled *bool `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

