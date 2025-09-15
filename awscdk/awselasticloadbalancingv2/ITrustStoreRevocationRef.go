package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustStoreRevocation.
// Experimental.
type ITrustStoreRevocationRef interface {
	constructs.IConstruct
	// A reference to a TrustStoreRevocation resource.
	// Experimental.
	TrustStoreRevocationRef() *TrustStoreRevocationReference
}

// The jsii proxy for ITrustStoreRevocationRef
type jsiiProxy_ITrustStoreRevocationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrustStoreRevocationRef) TrustStoreRevocationRef() *TrustStoreRevocationReference {
	var returns *TrustStoreRevocationReference
	_jsii_.Get(
		j,
		"trustStoreRevocationRef",
		&returns,
	)
	return returns
}

