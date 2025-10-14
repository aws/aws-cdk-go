package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustedEntitySet.
// Experimental.
type ITrustedEntitySetRef interface {
	constructs.IConstruct
	// A reference to a TrustedEntitySet resource.
	// Experimental.
	TrustedEntitySetRef() *TrustedEntitySetReference
}

// The jsii proxy for ITrustedEntitySetRef
type jsiiProxy_ITrustedEntitySetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrustedEntitySetRef) TrustedEntitySetRef() *TrustedEntitySetReference {
	var returns *TrustedEntitySetReference
	_jsii_.Get(
		j,
		"trustedEntitySetRef",
		&returns,
	)
	return returns
}

