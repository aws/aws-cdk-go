package previewawsglobalacceleratormixins


// Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
//
// For example, you can create a port override in which the listener receives user traffic on ports 80 and 443, but your accelerator routes that traffic to ports 1080 and 1443, respectively, on the endpoints.
//
// For more information, see [Port overrides](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoint-groups-port-override.html) in the *AWS Global Accelerator Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portOverrideProperty := &PortOverrideProperty{
//   	EndpointPort: jsii.Number(123),
//   	ListenerPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html
//
type CfnEndpointGroupPropsMixin_PortOverrideProperty struct {
	// The endpoint port that you want a listener port to be mapped to.
	//
	// This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-endpointport
	//
	EndpointPort *float64 `field:"optional" json:"endpointPort" yaml:"endpointPort"`
	// The listener port that you want to map to a specific endpoint port.
	//
	// This is the port that user traffic arrives to the Global Accelerator on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-listenerport
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
}

