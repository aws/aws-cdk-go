package awsfis


// Describes the configuration for experiment logging.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html
//
type CfnExperimentTemplate_ExperimentTemplateLogConfigurationProperty struct {
	// The schema version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-logschemaversion
	//
	LogSchemaVersion *float64 `field:"required" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// The configuration for experiment logging to Amazon CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-cloudwatchlogsconfiguration
	//
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// The configuration for experiment logging to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

