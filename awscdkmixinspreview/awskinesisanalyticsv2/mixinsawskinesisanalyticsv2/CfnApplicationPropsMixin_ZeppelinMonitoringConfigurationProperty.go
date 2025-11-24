package mixinsawskinesisanalyticsv2


// Describes configuration parameters for Amazon CloudWatch logging for a Kinesis Data Analytics Studio notebook.
//
// For more information about CloudWatch logging, see [Monitoring](https://docs.aws.amazon.com/managed-flink/latest/java/monitoring-overview.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   zeppelinMonitoringConfigurationProperty := &ZeppelinMonitoringConfigurationProperty{
//   	LogLevel: jsii.String("logLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinmonitoringconfiguration.html
//
type CfnApplicationPropsMixin_ZeppelinMonitoringConfigurationProperty struct {
	// The verbosity of the CloudWatch Logs for an application.
	//
	// You can set it to `INFO` , `WARN` , `ERROR` , or `DEBUG` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-zeppelinmonitoringconfiguration.html#cfn-kinesisanalyticsv2-application-zeppelinmonitoringconfiguration-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

