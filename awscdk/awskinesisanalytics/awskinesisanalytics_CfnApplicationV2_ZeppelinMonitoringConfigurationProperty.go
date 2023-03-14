package awskinesisanalytics


// Describes configuration parameters for Amazon CloudWatch logging for a Kinesis Data Analytics Studio notebook.
//
// For more information about CloudWatch logging, see [Monitoring](https://docs.aws.amazon.com/kinesisanalytics/latest/java/monitoring-overview.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zeppelinMonitoringConfigurationProperty := &zeppelinMonitoringConfigurationProperty{
//   	logLevel: jsii.String("logLevel"),
//   }
//
type CfnApplicationV2_ZeppelinMonitoringConfigurationProperty struct {
	// The verbosity of the CloudWatch Logs for an application.
	//
	// You can set it to `INFO` , `WARN` , `ERROR` , or `DEBUG` .
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

