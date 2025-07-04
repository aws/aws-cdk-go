package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// An EC2 Key Pair.
type IKeyPair interface {
	awscdk.IResource
	// The name of the key pair.
	KeyPairName() *string
	// The type of the key pair.
	Type() KeyPairType
}

// The jsii proxy for IKeyPair
type jsiiProxy_IKeyPair struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IKeyPair) KeyPairName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyPair) Type() KeyPairType {
	var returns KeyPairType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

