package awsecs


// The details of the execute command configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executeCommandConfigurationProperty := &ExecuteCommandConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogConfiguration: &ExecuteCommandLogConfigurationProperty{
//   		CloudWatchEncryptionEnabled: jsii.Boolean(false),
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3EncryptionEnabled: jsii.Boolean(false),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	Logging: jsii.String("logging"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html
//
type CfnCluster_ExecuteCommandConfigurationProperty struct {
	// Specify an AWS Key Management Service key ID to encrypt the data between the local client and the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When `logging=OVERRIDE` is specified, a `logConfiguration` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The log setting to use for redirecting logs for your execute command results. The following log settings are available.
	//
	// - `NONE` : The execute command session is not logged.
	// - `DEFAULT` : The `awslogs` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no `awslogs` log driver is configured in the task definition, the output won't be logged.
	// - `OVERRIDE` : Specify the logging details as a part of `logConfiguration` . If the `OVERRIDE` logging option is specified, the `logConfiguration` is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logging
	//
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
}

