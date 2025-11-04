package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPool.
// Experimental.
type IIdentityPoolRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdentityPool resource.
	// Experimental.
	IdentityPoolRef() *IdentityPoolReference
}

// The jsii proxy for IIdentityPoolRef
type jsiiProxy_IIdentityPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIdentityPoolRef) IdentityPoolRef() *IdentityPoolReference {
	var returns *IdentityPoolReference
	_jsii_.Get(
		j,
		"identityPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

