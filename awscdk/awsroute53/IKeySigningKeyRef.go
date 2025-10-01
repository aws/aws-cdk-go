package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeySigningKey.
// Experimental.
type IKeySigningKeyRef interface {
	constructs.IConstruct
	// A reference to a KeySigningKey resource.
	// Experimental.
	KeySigningKeyRef() *KeySigningKeyReference
}

// The jsii proxy for IKeySigningKeyRef
type jsiiProxy_IKeySigningKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKeySigningKeyRef) KeySigningKeyRef() *KeySigningKeyReference {
	var returns *KeySigningKeyReference
	_jsii_.Get(
		j,
		"keySigningKeyRef",
		&returns,
	)
	return returns
}

