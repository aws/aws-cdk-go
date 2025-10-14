package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicKey.
// Experimental.
type IPublicKeyRef interface {
	constructs.IConstruct
	// A reference to a PublicKey resource.
	// Experimental.
	PublicKeyRef() *PublicKeyReference
}

// The jsii proxy for IPublicKeyRef
type jsiiProxy_IPublicKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPublicKeyRef) PublicKeyRef() *PublicKeyReference {
	var returns *PublicKeyReference
	_jsii_.Get(
		j,
		"publicKeyRef",
		&returns,
	)
	return returns
}

