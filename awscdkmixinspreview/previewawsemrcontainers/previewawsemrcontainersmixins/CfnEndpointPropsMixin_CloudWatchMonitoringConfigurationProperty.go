package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchMonitoringConfigurationProperty := &CloudWatchMonitoringConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html
//
type CfnEndpointPropsMixin_CloudWatchMonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html#cfn-emrcontainers-endpoint-cloudwatchmonitoringconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.html#cfn-emrcontainers-endpoint-cloudwatchmonitoringconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
}

