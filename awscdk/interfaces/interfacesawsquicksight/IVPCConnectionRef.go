package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCConnection.
// Experimental.
type IVPCConnectionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VPCConnection resource.
	// Experimental.
	VpcConnectionRef() *VPCConnectionReference
}

// The jsii proxy for IVPCConnectionRef
type jsiiProxy_IVPCConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVPCConnectionRef) VpcConnectionRef() *VPCConnectionReference {
	var returns *VPCConnectionReference
	_jsii_.Get(
		j,
		"vpcConnectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCConnectionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

