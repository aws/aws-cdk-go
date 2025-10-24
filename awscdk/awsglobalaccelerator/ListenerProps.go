package awsglobalaccelerator


// Construct properties for Listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accelerator Accelerator
//
//   listenerProps := &ListenerProps{
//   	Accelerator: accelerator,
//   	PortRanges: []PortRange{
//   		&PortRange{
//   			FromPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ClientAffinity: awscdk.Aws_globalaccelerator.ClientAffinity_NONE,
//   	ListenerName: jsii.String("listenerName"),
//   	Protocol: awscdk.*Aws_globalaccelerator.ConnectionProtocol_TCP,
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
	// By default, each connection from each client is routed to separate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	// Default: ClientAffinity.NONE
	//
	ClientAffinity ClientAffinity `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// Name of the listener.
	// Default: - logical ID of the resource.
	//
	ListenerName *string `field:"optional" json:"listenerName" yaml:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	// Default: ConnectionProtocol.TCP
	//
	Protocol ConnectionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The accelerator for this listener.
	Accelerator IAccelerator `field:"required" json:"accelerator" yaml:"accelerator"`
}

