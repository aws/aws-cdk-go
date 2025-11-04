package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OIDCProvider.
// Experimental.
type IOIDCProviderRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OIDCProvider resource.
	// Experimental.
	OidcProviderRef() *OIDCProviderReference
}

// The jsii proxy for IOIDCProviderRef
type jsiiProxy_IOIDCProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOIDCProviderRef) OidcProviderRef() *OIDCProviderReference {
	var returns *OIDCProviderReference
	_jsii_.Get(
		j,
		"oidcProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOIDCProviderRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOIDCProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

