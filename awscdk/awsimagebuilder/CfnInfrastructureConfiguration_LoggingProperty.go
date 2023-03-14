package awsimagebuilder


// Logging configuration defines where Image Builder uploads your logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &LoggingProperty{
//   	S3Logs: &S3LogsProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   }
//
type CfnInfrastructureConfiguration_LoggingProperty struct {
	// The Amazon S3 logging configuration.
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

