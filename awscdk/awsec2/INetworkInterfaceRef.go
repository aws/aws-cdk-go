package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterface.
// Experimental.
type INetworkInterfaceRef interface {
	constructs.IConstruct
	// A reference to a NetworkInterface resource.
	// Experimental.
	NetworkInterfaceRef() *NetworkInterfaceReference
}

// The jsii proxy for INetworkInterfaceRef
type jsiiProxy_INetworkInterfaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkInterfaceRef) NetworkInterfaceRef() *NetworkInterfaceReference {
	var returns *NetworkInterfaceReference
	_jsii_.Get(
		j,
		"networkInterfaceRef",
		&returns,
	)
	return returns
}

