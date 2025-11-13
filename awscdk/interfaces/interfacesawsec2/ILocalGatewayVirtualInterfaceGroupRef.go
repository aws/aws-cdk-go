package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayVirtualInterfaceGroup.
// Experimental.
type ILocalGatewayVirtualInterfaceGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocalGatewayVirtualInterfaceGroup resource.
	// Experimental.
	LocalGatewayVirtualInterfaceGroupRef() *LocalGatewayVirtualInterfaceGroupReference
}

// The jsii proxy for ILocalGatewayVirtualInterfaceGroupRef
type jsiiProxy_ILocalGatewayVirtualInterfaceGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceGroupRef) LocalGatewayVirtualInterfaceGroupRef() *LocalGatewayVirtualInterfaceGroupReference {
	var returns *LocalGatewayVirtualInterfaceGroupReference
	_jsii_.Get(
		j,
		"localGatewayVirtualInterfaceGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

