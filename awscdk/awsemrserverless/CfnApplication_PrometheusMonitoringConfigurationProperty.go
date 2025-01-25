package awsemrserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prometheusMonitoringConfigurationProperty := &PrometheusMonitoringConfigurationProperty{
//   	RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-prometheusmonitoringconfiguration.html
//
type CfnApplication_PrometheusMonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-prometheusmonitoringconfiguration.html#cfn-emrserverless-application-prometheusmonitoringconfiguration-remotewriteurl
	//
	RemoteWriteUrl *string `field:"optional" json:"remoteWriteUrl" yaml:"remoteWriteUrl"`
}

