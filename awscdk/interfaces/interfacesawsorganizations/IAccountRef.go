package interfacesawsorganizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Account.
// Experimental.
type IAccountRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Account resource.
	// Experimental.
	AccountRef() *AccountReference
}

// The jsii proxy for IAccountRef
type jsiiProxy_IAccountRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAccountRef) AccountRef() *AccountReference {
	var returns *AccountReference
	_jsii_.Get(
		j,
		"accountRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

