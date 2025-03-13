package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
)

// A Key Signing Key for a Route 53 Hosted Zone.
type IKeySigningKey interface {
	awscdk.IResource
	// The hosted zone that the key signing key signs.
	HostedZone() IHostedZone
	// The ID of the key signing key, derived from the hosted zone ID and its name.
	KeySigningKeyId() *string
	// The name of the key signing key.
	KeySigningKeyName() *string
}

// The jsii proxy for IKeySigningKey
type jsiiProxy_IKeySigningKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IKeySigningKey) HostedZone() IHostedZone {
	var returns IHostedZone
	_jsii_.Get(
		j,
		"hostedZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeySigningKey) KeySigningKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySigningKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeySigningKey) KeySigningKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySigningKeyName",
		&returns,
	)
	return returns
}

