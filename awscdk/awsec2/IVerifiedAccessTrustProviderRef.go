package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessTrustProvider.
// Experimental.
type IVerifiedAccessTrustProviderRef interface {
	constructs.IConstruct
	// A reference to a VerifiedAccessTrustProvider resource.
	// Experimental.
	VerifiedAccessTrustProviderRef() *VerifiedAccessTrustProviderReference
}

// The jsii proxy for IVerifiedAccessTrustProviderRef
type jsiiProxy_IVerifiedAccessTrustProviderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVerifiedAccessTrustProviderRef) VerifiedAccessTrustProviderRef() *VerifiedAccessTrustProviderReference {
	var returns *VerifiedAccessTrustProviderReference
	_jsii_.Get(
		j,
		"verifiedAccessTrustProviderRef",
		&returns,
	)
	return returns
}

