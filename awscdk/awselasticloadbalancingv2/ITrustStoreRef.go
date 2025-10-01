package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustStore.
// Experimental.
type ITrustStoreRef interface {
	constructs.IConstruct
	// A reference to a TrustStore resource.
	// Experimental.
	TrustStoreRef() *TrustStoreReference
}

// The jsii proxy for ITrustStoreRef
type jsiiProxy_ITrustStoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrustStoreRef) TrustStoreRef() *TrustStoreReference {
	var returns *TrustStoreReference
	_jsii_.Get(
		j,
		"trustStoreRef",
		&returns,
	)
	return returns
}

