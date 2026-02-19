package interfacesawsnotificationscontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnotificationscontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailContact.
// Experimental.
type IEmailContactRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EmailContact resource.
	// Experimental.
	EmailContactRef() *EmailContactReference
}

// The jsii proxy for IEmailContactRef
type jsiiProxy_IEmailContactRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEmailContactRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEmailContactRef) EmailContactRef() *EmailContactReference {
	var returns *EmailContactReference
	_jsii_.Get(
		j,
		"emailContactRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailContactRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailContactRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

