package awsecs


// The log settings to use to for logging the execute command session.
//
// For more information, see
// [Logging] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logging
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
type ExecuteCommandLogging string

const (
	// The execute command session is not logged.
	ExecuteCommandLogging_NONE ExecuteCommandLogging = "NONE"
	// The awslogs configuration in the task definition is used.
	//
	// If no logging parameter is specified, it defaults to this value. If no awslogs log driver is configured in the task definition, the output won't be logged.
	ExecuteCommandLogging_DEFAULT ExecuteCommandLogging = "DEFAULT"
	// Specify the logging details as a part of logConfiguration.
	ExecuteCommandLogging_OVERRIDE ExecuteCommandLogging = "OVERRIDE"
)

