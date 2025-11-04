package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyPair.
// Experimental.
type IKeyPairRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a KeyPair resource.
	// Experimental.
	KeyPairRef() *KeyPairReference
}

// The jsii proxy for IKeyPairRef
type jsiiProxy_IKeyPairRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IKeyPairRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyPairRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

