package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subnet.
// Experimental.
type ISubnetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Subnet resource.
	// Experimental.
	SubnetRef() *SubnetReference
}

// The jsii proxy for ISubnetRef
type jsiiProxy_ISubnetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISubnetRef) SubnetRef() *SubnetReference {
	var returns *SubnetReference
	_jsii_.Get(
		j,
		"subnetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

