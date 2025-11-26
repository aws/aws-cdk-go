package previewawsemrserverlessmixins


// The monitoring configuration object you can configure to send metrics to Amazon Managed Service for Prometheus for a job run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   prometheusMonitoringConfigurationProperty := &PrometheusMonitoringConfigurationProperty{
//   	RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-prometheusmonitoringconfiguration.html
//
type CfnApplicationPropsMixin_PrometheusMonitoringConfigurationProperty struct {
	// The remote write URL in the Amazon Managed Service for Prometheus workspace to send metrics to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-prometheusmonitoringconfiguration.html#cfn-emrserverless-application-prometheusmonitoringconfiguration-remotewriteurl
	//
	RemoteWriteUrl *string `field:"optional" json:"remoteWriteUrl" yaml:"remoteWriteUrl"`
}

