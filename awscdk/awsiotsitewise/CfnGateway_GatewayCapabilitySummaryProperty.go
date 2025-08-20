package awsiotsitewise


// Contains a summary of a gateway capability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayCapabilitySummaryProperty := &GatewayCapabilitySummaryProperty{
//   	CapabilityNamespace: jsii.String("capabilityNamespace"),
//
//   	// the properties below are optional
//   	CapabilityConfiguration: jsii.String("capabilityConfiguration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html
//
type CfnGateway_GatewayCapabilitySummaryProperty struct {
	// The namespace of the capability configuration.
	//
	// For example, if you configure OPC UA sources for an MQTT-enabled gateway, your OPC-UA capability configuration has the namespace `iotsitewise:opcuacollector:3` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilitynamespace
	//
	CapabilityNamespace *string `field:"required" json:"capabilityNamespace" yaml:"capabilityNamespace"`
	// The JSON document that defines the configuration for the gateway capability.
	//
	// For more information, see [Configuring data sources (CLI)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/configure-sources.html#configure-source-cli) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilityconfiguration
	//
	CapabilityConfiguration *string `field:"optional" json:"capabilityConfiguration" yaml:"capabilityConfiguration"`
}

