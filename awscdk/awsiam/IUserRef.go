package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a User.
// Experimental.
type IUserRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a User resource.
	// Experimental.
	UserRef() *UserReference
}

// The jsii proxy for IUserRef
type jsiiProxy_IUserRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserRef) UserRef() *UserReference {
	var returns *UserReference
	_jsii_.Get(
		j,
		"userRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

