package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Authorizer.
// Experimental.
type IAuthorizerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Authorizer resource.
	// Experimental.
	AuthorizerRef() *AuthorizerReference
}

// The jsii proxy for IAuthorizerRef
type jsiiProxy_IAuthorizerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAuthorizerRef) AuthorizerRef() *AuthorizerReference {
	var returns *AuthorizerReference
	_jsii_.Get(
		j,
		"authorizerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthorizerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthorizerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

