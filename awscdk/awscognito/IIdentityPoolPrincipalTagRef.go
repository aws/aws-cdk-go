package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPoolPrincipalTag.
// Experimental.
type IIdentityPoolPrincipalTagRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdentityPoolPrincipalTag resource.
	// Experimental.
	IdentityPoolPrincipalTagRef() *IdentityPoolPrincipalTagReference
}

// The jsii proxy for IIdentityPoolPrincipalTagRef
type jsiiProxy_IIdentityPoolPrincipalTagRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIdentityPoolPrincipalTagRef) IdentityPoolPrincipalTagRef() *IdentityPoolPrincipalTagReference {
	var returns *IdentityPoolPrincipalTagReference
	_jsii_.Get(
		j,
		"identityPoolPrincipalTagRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolPrincipalTagRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolPrincipalTagRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

