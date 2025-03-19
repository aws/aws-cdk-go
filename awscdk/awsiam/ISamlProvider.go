package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// A SAML provider.
type ISamlProvider interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the provider.
	SamlProviderArn() *string
}

// The jsii proxy for ISamlProvider
type jsiiProxy_ISamlProvider struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISamlProvider) SamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArn",
		&returns,
	)
	return returns
}

