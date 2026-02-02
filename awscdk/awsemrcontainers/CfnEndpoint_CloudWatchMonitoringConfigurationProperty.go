package awsemrcontainers


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchMonitoringConfigurationProperty := &CloudWatchMonitoringConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html
//
type CfnEndpoint_CloudWatchMonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html#cfn-emrcontainers-endpoint-cloudwatchmonitoringconfiguration-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html#cfn-emrcontainers-endpoint-cloudwatchmonitoringconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
}

