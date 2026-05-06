package interfacesawschime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppInstance.
// Experimental.
type IAppInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AppInstance resource.
	// Experimental.
	AppInstanceRef() *AppInstanceReference
}

// The jsii proxy for IAppInstanceRef
type jsiiProxy_IAppInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAppInstanceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAppInstanceRef) AppInstanceRef() *AppInstanceReference {
	var returns *AppInstanceReference
	_jsii_.Get(
		j,
		"appInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

