package interfacesawsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcConnection.
// Experimental.
type IVpcConnectionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VpcConnection resource.
	// Experimental.
	VpcConnectionRef() *VpcConnectionReference
}

// The jsii proxy for IVpcConnectionRef
type jsiiProxy_IVpcConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVpcConnectionRef) VpcConnectionRef() *VpcConnectionReference {
	var returns *VpcConnectionReference
	_jsii_.Get(
		j,
		"vpcConnectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnectionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

