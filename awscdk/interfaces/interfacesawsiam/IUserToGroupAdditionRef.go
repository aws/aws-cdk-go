package interfacesawsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserToGroupAddition.
// Experimental.
type IUserToGroupAdditionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserToGroupAddition resource.
	// Experimental.
	UserToGroupAdditionRef() *UserToGroupAdditionReference
}

// The jsii proxy for IUserToGroupAdditionRef
type jsiiProxy_IUserToGroupAdditionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserToGroupAdditionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserToGroupAdditionRef) UserToGroupAdditionRef() *UserToGroupAdditionReference {
	var returns *UserToGroupAdditionReference
	_jsii_.Get(
		j,
		"userToGroupAdditionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserToGroupAdditionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserToGroupAdditionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

