package awsfis


// Specifies the configuration for experiment logging.
//
// For more information, see [Experiment logging](https://docs.aws.amazon.com/fis/latest/userguide/monitoring-logging.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//
//   experimentTemplateLogConfigurationProperty := &ExperimentTemplateLogConfigurationProperty{
//   	LogSchemaVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	CloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   	S3Configuration: s3Configuration,
//   }
//
type CfnExperimentTemplate_ExperimentTemplateLogConfigurationProperty struct {
	// The schema version.
	LogSchemaVersion *float64 `field:"required" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// The configuration for experiment logging to CloudWatch Logs .
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// The configuration for experiment logging to Amazon S3 .
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

