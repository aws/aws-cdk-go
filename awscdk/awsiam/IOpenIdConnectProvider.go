package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// Represents an IAM OpenID Connect provider.
type IOpenIdConnectProvider interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	OpenIdConnectProviderArn() *string
	// The issuer for OIDC Provider.
	OpenIdConnectProviderIssuer() *string
}

// The jsii proxy for IOpenIdConnectProvider
type jsiiProxy_IOpenIdConnectProvider struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IOpenIdConnectProvider) OpenIdConnectProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenIdConnectProvider) OpenIdConnectProviderIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderIssuer",
		&returns,
	)
	return returns
}

