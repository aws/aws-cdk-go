package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayVirtualInterface.
// Experimental.
type ILocalGatewayVirtualInterfaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocalGatewayVirtualInterface resource.
	// Experimental.
	LocalGatewayVirtualInterfaceRef() *LocalGatewayVirtualInterfaceReference
}

// The jsii proxy for ILocalGatewayVirtualInterfaceRef
type jsiiProxy_ILocalGatewayVirtualInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceRef) LocalGatewayVirtualInterfaceRef() *LocalGatewayVirtualInterfaceReference {
	var returns *LocalGatewayVirtualInterfaceReference
	_jsii_.Get(
		j,
		"localGatewayVirtualInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

