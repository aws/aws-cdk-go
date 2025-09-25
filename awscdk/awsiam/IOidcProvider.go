package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM OpenID Connect provider.
type IOidcProvider interface {
	IOIDCProviderRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	OidcProviderArn() *string
	// The issuer for OIDC Provider.
	OidcProviderIssuer() *string
}

// The jsii proxy for IOidcProvider
type jsiiProxy_IOidcProvider struct {
	jsiiProxy_IOIDCProviderRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOidcProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IOidcProvider) OidcProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) OidcProviderIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcProviderIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) OidcProviderRef() *OIDCProviderReference {
	var returns *OIDCProviderReference
	_jsii_.Get(
		j,
		"oidcProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

