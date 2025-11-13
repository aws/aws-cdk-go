package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUser.
// Experimental.
type IUserPoolUserRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserPoolUser resource.
	// Experimental.
	UserPoolUserRef() *UserPoolUserReference
}

// The jsii proxy for IUserPoolUserRef
type jsiiProxy_IUserPoolUserRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolUserRef) UserPoolUserRef() *UserPoolUserReference {
	var returns *UserPoolUserReference
	_jsii_.Get(
		j,
		"userPoolUserRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolUserRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolUserRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

