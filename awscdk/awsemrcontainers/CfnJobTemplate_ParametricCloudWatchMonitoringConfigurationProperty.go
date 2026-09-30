package awsemrcontainers


// A configuration for CloudWatch monitoring.
//
// This data type allows job template parameters to be specified within.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametricCloudWatchMonitoringConfigurationProperty := &ParametricCloudWatchMonitoringConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametriccloudwatchmonitoringconfiguration.html
//
type CfnJobTemplate_ParametricCloudWatchMonitoringConfigurationProperty struct {
	// The name of the log group for log publishing, or a template parameter reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametriccloudwatchmonitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametriccloudwatchmonitoringconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The specified name prefix for log streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametriccloudwatchmonitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametriccloudwatchmonitoringconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
}

