package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM OpenID Connect provider.
type IOidcProvider interface {
	interfacesawsiam.IOIDCProviderRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	OidcProviderArn() *string
	// The issuer for OIDC Provider.
	OidcProviderIssuer() *string
	// Alias for `oidcProviderArn` to maintain backwards compatibility for constructs which accept `iam.IOpenIdConnectProvider`.
	//
	// Use `oidcProviderArn` instead. This property exists for backward compatibility with existing constructs as migrating between the 2 constructs (OpenIdConnectProvider and OidcProviderNative) is not reasonably feasible as it requires a manual step (cdk import) since the resource type is changing between OpenIdConnectProvider and OidcProviderNative.
	OpenIdConnectProviderArn() *string
	// Alias for `oidcProviderIssuer` to maintain backwards compatibility for constructs which accept `iam.IOpenIdConnectProvider.
	//
	// Use `oidcProviderIssuer` instead. This property exists for backward compatibility with existing constructs as migrating between the 2 constructs (OpenIdConnectProvider and OidcProviderNative) is not reasonably feasible as it requires a manual step (cdk import) since the resource type is changing between OpenIdConnectProvider and OidcProviderNative.
	OpenIdConnectProviderIssuer() *string
}

// The jsii proxy for IOidcProvider
type jsiiProxy_IOidcProvider struct {
	internal.Type__interfacesawsiamIOIDCProviderRef
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

func (j *jsiiProxy_IOidcProvider) OpenIdConnectProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) OpenIdConnectProviderIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOidcProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

func (j *jsiiProxy_IOidcProvider) OidcProviderRef() *interfacesawsiam.OIDCProviderReference {
	var returns *interfacesawsiam.OIDCProviderReference
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

