package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolClient.
// Experimental.
type IUserPoolClientRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserPoolClient resource.
	// Experimental.
	UserPoolClientRef() *UserPoolClientReference
}

// The jsii proxy for IUserPoolClientRef
type jsiiProxy_IUserPoolClientRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserPoolClientRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserPoolClientRef) UserPoolClientRef() *UserPoolClientReference {
	var returns *UserPoolClientReference
	_jsii_.Get(
		j,
		"userPoolClientRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClientRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClientRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

