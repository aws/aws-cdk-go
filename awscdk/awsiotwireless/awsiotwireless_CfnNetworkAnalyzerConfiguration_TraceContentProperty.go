package awsiotwireless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   traceContentProperty := &traceContentProperty{
//   	logLevel: jsii.String("logLevel"),
//   	wirelessDeviceFrameInfo: jsii.String("wirelessDeviceFrameInfo"),
//   }
//
type CfnNetworkAnalyzerConfiguration_TraceContentProperty struct {
	// `CfnNetworkAnalyzerConfiguration.TraceContentProperty.LogLevel`.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// `CfnNetworkAnalyzerConfiguration.TraceContentProperty.WirelessDeviceFrameInfo`.
	WirelessDeviceFrameInfo *string `field:"optional" json:"wirelessDeviceFrameInfo" yaml:"wirelessDeviceFrameInfo"`
}

