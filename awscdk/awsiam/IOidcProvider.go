package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// Represents an IAM OpenID Connect provider.
type IOidcProvider interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	OidcProviderArn() *string
	// The issuer for OIDC Provider.
	OidcProviderIssuer() *string
}

// The jsii proxy for IOidcProvider
type jsiiProxy_IOidcProvider struct {
	internal.Type__awscdkIResource
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

