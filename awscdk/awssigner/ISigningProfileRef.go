package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SigningProfile.
// Experimental.
type ISigningProfileRef interface {
	constructs.IConstruct
	// A reference to a SigningProfile resource.
	// Experimental.
	SigningProfileRef() *SigningProfileReference
}

// The jsii proxy for ISigningProfileRef
type jsiiProxy_ISigningProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISigningProfileRef) SigningProfileRef() *SigningProfileReference {
	var returns *SigningProfileReference
	_jsii_.Get(
		j,
		"signingProfileRef",
		&returns,
	)
	return returns
}

