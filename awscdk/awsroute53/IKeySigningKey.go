package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Key Signing Key for a Route 53 Hosted Zone.
type IKeySigningKey interface {
	interfacesawsroute53.IKeySigningKeyRef
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
	internal.Type__interfacesawsroute53IKeySigningKeyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKeySigningKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IKeySigningKey) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeySigningKey) KeySigningKeyRef() *interfacesawsroute53.KeySigningKeyReference {
	var returns *interfacesawsroute53.KeySigningKeyReference
	_jsii_.Get(
		j,
		"keySigningKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeySigningKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeySigningKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

