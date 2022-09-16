package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNetworkAnalyzerConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var traceContent interface{}
//
//   cfnNetworkAnalyzerConfigurationProps := &cfnNetworkAnalyzerConfigurationProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceContent: traceContent,
//   	wirelessDevices: []*string{
//   		jsii.String("wirelessDevices"),
//   	},
//   	wirelessGateways: []*string{
//   		jsii.String("wirelessGateways"),
//   	},
//   }
//
type CfnNetworkAnalyzerConfigurationProps struct {
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.TraceContent`.
	TraceContent interface{} `field:"optional" json:"traceContent" yaml:"traceContent"`
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.WirelessDevices`.
	WirelessDevices *[]*string `field:"optional" json:"wirelessDevices" yaml:"wirelessDevices"`
	// `AWS::IoTWireless::NetworkAnalyzerConfiguration.WirelessGateways`.
	WirelessGateways *[]*string `field:"optional" json:"wirelessGateways" yaml:"wirelessGateways"`
}

