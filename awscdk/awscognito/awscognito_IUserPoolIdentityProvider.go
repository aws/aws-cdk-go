package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito/internal"
)

// Represents a UserPoolIdentityProvider.
// Experimental.
type IUserPoolIdentityProvider interface {
	awscdk.IResource
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
}

// The jsii proxy for IUserPoolIdentityProvider
type jsiiProxy_IUserPoolIdentityProvider struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolIdentityProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

