package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolResourceServer.
// Experimental.
type IUserPoolResourceServerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPoolResourceServer resource.
	// Experimental.
	UserPoolResourceServerRef() *UserPoolResourceServerReference
}

// The jsii proxy for IUserPoolResourceServerRef
type jsiiProxy_IUserPoolResourceServerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolResourceServerRef) UserPoolResourceServerRef() *UserPoolResourceServerReference {
	var returns *UserPoolResourceServerReference
	_jsii_.Get(
		j,
		"userPoolResourceServerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

