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
//   cfnNetworkAnalyzerConfigurationProps := &CfnNetworkAnalyzerConfigurationProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TraceContent: traceContent,
//   	WirelessDevices: []*string{
//   		jsii.String("wirelessDevices"),
//   	},
//   	WirelessGateways: []*string{
//   		jsii.String("wirelessGateways"),
//   	},
//   }
//
type CfnNetworkAnalyzerConfigurationProps struct {
	// Name of the network analyzer configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags to attach to the specified resource.
	//
	// Tags are metadata that you can use to manage a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Trace content for your wireless gateway and wireless device resources.
	TraceContent interface{} `field:"optional" json:"traceContent" yaml:"traceContent"`
	// Wireless device resources to add to the network analyzer configuration.
	//
	// Provide the `WirelessDeviceId` of the resource to add in the input array.
	WirelessDevices *[]*string `field:"optional" json:"wirelessDevices" yaml:"wirelessDevices"`
	// Wireless gateway resources to add to the network analyzer configuration.
	//
	// Provide the `WirelessGatewayId` of the resource to add in the input array.
	WirelessGateways *[]*string `field:"optional" json:"wirelessGateways" yaml:"wirelessGateways"`
}

