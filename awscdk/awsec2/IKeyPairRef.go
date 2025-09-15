package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyPair.
// Experimental.
type IKeyPairRef interface {
	constructs.IConstruct
	// A reference to a KeyPair resource.
	// Experimental.
	KeyPairRef() *KeyPairReference
}

// The jsii proxy for IKeyPairRef
type jsiiProxy_IKeyPairRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKeyPairRef) KeyPairRef() *KeyPairReference {
	var returns *KeyPairReference
	_jsii_.Get(
		j,
		"keyPairRef",
		&returns,
	)
	return returns
}

