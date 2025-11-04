package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPool.
// Experimental.
type IUserPoolRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPool resource.
	// Experimental.
	UserPoolRef() *UserPoolReference
}

// The jsii proxy for IUserPoolRef
type jsiiProxy_IUserPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolRef) UserPoolRef() *UserPoolReference {
	var returns *UserPoolReference
	_jsii_.Get(
		j,
		"userPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

