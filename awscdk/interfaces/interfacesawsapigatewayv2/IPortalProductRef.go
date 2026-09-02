package interfacesawsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortalProduct.
// Experimental.
type IPortalProductRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PortalProduct resource.
	// Experimental.
	PortalProductRef() *PortalProductReference
}

// The jsii proxy for IPortalProductRef
type jsiiProxy_IPortalProductRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPortalProductRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPortalProductRef) PortalProductRef() *PortalProductReference {
	var returns *PortalProductReference
	_jsii_.Get(
		j,
		"portalProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortalProductRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortalProductRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

