package awsglobalaccelerator


// Construct properties for Listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accelerator accelerator
//
//   listenerProps := &listenerProps{
//   	accelerator: accelerator,
//   	portRanges: []portRange{
//   		&portRange{
//   			fromPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			toPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	clientAffinity: awscdk.Aws_globalaccelerator.clientAffinity_NONE,
//   	listenerName: jsii.String("listenerName"),
//   	protocol: awscdk.*Aws_globalaccelerator.connectionProtocol_TCP,
//   }
//
type ListenerProps struct {
	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges *[]*PortRange `field:"required" json:"portRanges" yaml:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to seperate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	ClientAffinity ClientAffinity `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// Name of the listener.
	ListenerName *string `field:"optional" json:"listenerName" yaml:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	Protocol ConnectionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The accelerator for this listener.
	Accelerator IAccelerator `field:"required" json:"accelerator" yaml:"accelerator"`
}

