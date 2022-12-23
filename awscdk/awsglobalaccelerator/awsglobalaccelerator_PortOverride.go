package awsglobalaccelerator


// Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portOverride := &portOverride{
//   	endpointPort: jsii.Number(123),
//   	listenerPort: jsii.Number(123),
//   }
//
type PortOverride struct {
	// The endpoint port that you want a listener port to be mapped to.
	//
	// This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	EndpointPort *float64 `field:"required" json:"endpointPort" yaml:"endpointPort"`
	// The listener port that you want to map to a specific endpoint port.
	//
	// This is the port that user traffic arrives to the Global Accelerator on.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
}

