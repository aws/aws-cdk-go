package awsglobalaccelerator


// Properties for defining a `CfnEndpointGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointGroupProps := &CfnEndpointGroupProps{
//   	EndpointGroupRegion: jsii.String("endpointGroupRegion"),
//   	ListenerArn: jsii.String("listenerArn"),
//
//   	// the properties below are optional
//   	EndpointConfigurations: []interface{}{
//   		&EndpointConfigurationProperty{
//   			EndpointId: jsii.String("endpointId"),
//
//   			// the properties below are optional
//   			AttachmentArn: jsii.String("attachmentArn"),
//   			ClientIpPreservationEnabled: jsii.Boolean(false),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	HealthCheckPort: jsii.Number(123),
//   	HealthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	PortOverrides: []interface{}{
//   		&PortOverrideProperty{
//   			EndpointPort: jsii.Number(123),
//   			ListenerPort: jsii.Number(123),
//   		},
//   	},
//   	ThresholdCount: jsii.Number(123),
//   	TrafficDialPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html
//
type CfnEndpointGroupProps struct {
	// The AWS Regions where the endpoint group is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
	//
	EndpointGroupRegion *string `field:"required" json:"endpointGroupRegion" yaml:"endpointGroupRegion"`
	// The Amazon Resource Name (ARN) of the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
	//
	ListenerArn interface{} `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The list of endpoint objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
	//
	EndpointConfigurations interface{} `field:"optional" json:"endpointConfigurations" yaml:"endpointConfigurations"`
	// The time—10 seconds or 30 seconds—between health checks for each endpoint.
	//
	// The default value is 30.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
	//
	// Default: - 30.
	//
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this value provides the ping path that Global Accelerator uses for the destination on the endpoints for health checks.
	//
	// The default is slash (/).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
	//
	// Default: - "/".
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.
	//
	// The default port is the port for the listener that this endpoint group is associated with. If the listener port is a list, Global Accelerator uses the first specified port in the list of ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
	//
	// Default: - -1.
	//
	HealthCheckPort *float64 `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.
	//
	// The default value is TCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
	//
	// Default: - "TCP".
	//
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// Allows you to override the destination ports used to route traffic to an endpoint.
	//
	// Using a port override lets you map a list of external destination ports (that your users send traffic to) to a list of internal destination ports that you want an application endpoint to receive traffic on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
	//
	PortOverrides interface{} `field:"optional" json:"portOverrides" yaml:"portOverrides"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	//
	// The default value is 3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
	//
	// Default: - 3.
	//
	ThresholdCount *float64 `field:"optional" json:"thresholdCount" yaml:"thresholdCount"`
	// The percentage of traffic to send to an AWS Regions .
	//
	// Additional traffic is distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to a specific Region. The percentage is applied to the traffic that would otherwise have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
	//
	// Default: - 100.
	//
	TrafficDialPercentage *float64 `field:"optional" json:"trafficDialPercentage" yaml:"trafficDialPercentage"`
}

