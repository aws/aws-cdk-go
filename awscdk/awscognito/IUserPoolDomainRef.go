package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolDomain.
// Experimental.
type IUserPoolDomainRef interface {
	constructs.IConstruct
	// A reference to a UserPoolDomain resource.
	// Experimental.
	UserPoolDomainRef() *UserPoolDomainReference
}

// The jsii proxy for IUserPoolDomainRef
type jsiiProxy_IUserPoolDomainRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPoolDomainRef) UserPoolDomainRef() *UserPoolDomainReference {
	var returns *UserPoolDomainReference
	_jsii_.Get(
		j,
		"userPoolDomainRef",
		&returns,
	)
	return returns
}

