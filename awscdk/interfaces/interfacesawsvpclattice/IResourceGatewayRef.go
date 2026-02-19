package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceGateway.
// Experimental.
type IResourceGatewayRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResourceGateway resource.
	// Experimental.
	ResourceGatewayRef() *ResourceGatewayReference
}

// The jsii proxy for IResourceGatewayRef
type jsiiProxy_IResourceGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IResourceGatewayRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IResourceGatewayRef) ResourceGatewayRef() *ResourceGatewayReference {
	var returns *ResourceGatewayReference
	_jsii_.Get(
		j,
		"resourceGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceGatewayRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

