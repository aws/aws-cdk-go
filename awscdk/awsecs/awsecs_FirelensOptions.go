package awsecs


// The options for firelens log router https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensOptions := &firelensOptions{
//   	configFileType: awscdk.Aws_ecs.firelensConfigFileType_S3,
//   	configFileValue: jsii.String("configFileValue"),
//   	enableECSLogMetadata: jsii.Boolean(false),
//   }
//
type FirelensOptions struct {
	// Custom configuration file, s3 or file.
	//
	// Both configFileType and configFileValue must be used together
	// to define a custom configuration source.
	ConfigFileType FirelensConfigFileType `field:"optional" json:"configFileType" yaml:"configFileType"`
	// Custom configuration file, S3 ARN or a file path Both configFileType and configFileValue must be used together to define a custom configuration source.
	ConfigFileValue *string `field:"optional" json:"configFileValue" yaml:"configFileValue"`
	// By default, Amazon ECS adds additional fields in your log entries that help identify the source of the logs.
	//
	// You can disable this action by setting enable-ecs-log-metadata to false.
	EnableECSLogMetadata *bool `field:"optional" json:"enableECSLogMetadata" yaml:"enableECSLogMetadata"`
}

