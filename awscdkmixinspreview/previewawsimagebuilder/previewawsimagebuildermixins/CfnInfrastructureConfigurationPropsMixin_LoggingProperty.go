package previewawsimagebuildermixins


// Logging configuration defines where Image Builder uploads your logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingProperty := &LoggingProperty{
//   	S3Logs: &S3LogsProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-logging.html
//
type CfnInfrastructureConfigurationPropsMixin_LoggingProperty struct {
	// The Amazon S3 logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-logging.html#cfn-imagebuilder-infrastructureconfiguration-logging-s3logs
	//
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

