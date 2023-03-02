package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito/internal"
)

// Represents a user pool domain.
// Experimental.
type IUserPoolDomain interface {
	awscdk.IResource
	// The domain that was specified to be created.
	//
	// If `customDomain` was selected, this holds the full domain name that was specified.
	// If the `cognitoDomain` was used, it contains the prefix to the Cognito hosted domain.
	// Experimental.
	DomainName() *string
}

// The jsii proxy for IUserPoolDomain
type jsiiProxy_IUserPoolDomain struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

