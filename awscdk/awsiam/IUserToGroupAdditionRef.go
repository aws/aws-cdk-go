package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserToGroupAddition.
// Experimental.
type IUserToGroupAdditionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserToGroupAddition resource.
	// Experimental.
	UserToGroupAdditionRef() *UserToGroupAdditionReference
}

// The jsii proxy for IUserToGroupAdditionRef
type jsiiProxy_IUserToGroupAdditionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IUserToGroupAdditionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

