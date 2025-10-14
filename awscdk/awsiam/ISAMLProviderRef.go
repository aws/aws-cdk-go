package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SAMLProvider.
// Experimental.
type ISAMLProviderRef interface {
	constructs.IConstruct
	// A reference to a SAMLProvider resource.
	// Experimental.
	SamlProviderRef() *SAMLProviderReference
}

// The jsii proxy for ISAMLProviderRef
type jsiiProxy_ISAMLProviderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISAMLProviderRef) SamlProviderRef() *SAMLProviderReference {
	var returns *SAMLProviderReference
	_jsii_.Get(
		j,
		"samlProviderRef",
		&returns,
	)
	return returns
}

