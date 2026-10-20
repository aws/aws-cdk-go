package interfacesawsappfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppAuthorization.
// Experimental.
type IAppAuthorizationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AppAuthorization resource.
	// Experimental.
	AppAuthorizationRef() *AppAuthorizationReference
}

// The jsii proxy for IAppAuthorizationRef
type jsiiProxy_IAppAuthorizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAppAuthorizationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAppAuthorizationRef) AppAuthorizationRef() *AppAuthorizationReference {
	var returns *AppAuthorizationReference
	_jsii_.Get(
		j,
		"appAuthorizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppAuthorizationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppAuthorizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

