package awsemrcontainers


// Configuration setting for monitoring.
//
// This data type allows job template parameters to be specified within.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametricMonitoringConfigurationProperty := &ParametricMonitoringConfigurationProperty{
//   	CloudWatchMonitoringConfiguration: &ParametricCloudWatchMonitoringConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	},
//   	PersistentAppUi: jsii.String("persistentAppUi"),
//   	S3MonitoringConfiguration: &ParametricS3MonitoringConfigurationProperty{
//   		LogUri: jsii.String("logUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricmonitoringconfiguration.html
//
type CfnJobTemplate_ParametricMonitoringConfigurationProperty struct {
	// A configuration for CloudWatch monitoring.
	//
	// This data type allows job template parameters to be specified within.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricmonitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametricmonitoringconfiguration-cloudwatchmonitoringconfiguration
	//
	CloudWatchMonitoringConfiguration interface{} `field:"optional" json:"cloudWatchMonitoringConfiguration" yaml:"cloudWatchMonitoringConfiguration"`
	// Monitoring configurations for the persistent application UI, or a template parameter reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricmonitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametricmonitoringconfiguration-persistentappui
	//
	PersistentAppUi *string `field:"optional" json:"persistentAppUi" yaml:"persistentAppUi"`
	// Amazon S3 configuration for monitoring log publishing.
	//
	// This data type allows job template parameters to be specified within.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricmonitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametricmonitoringconfiguration-s3monitoringconfiguration
	//
	S3MonitoringConfiguration interface{} `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

