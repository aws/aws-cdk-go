package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portal.
// Experimental.
type IPortalRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Portal resource.
	// Experimental.
	PortalRef() *PortalReference
}

// The jsii proxy for IPortalRef
type jsiiProxy_IPortalRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPortalRef) PortalRef() *PortalReference {
	var returns *PortalReference
	_jsii_.Get(
		j,
		"portalRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortalRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortalRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

