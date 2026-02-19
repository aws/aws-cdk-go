package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// An EC2 Key Pair.
type IKeyPair interface {
	interfacesawsec2.IKeyPairRef
	awscdk.IResource
	// The name of the key pair.
	KeyPairName() *string
	// The type of the key pair.
	Type() KeyPairType
}

// The jsii proxy for IKeyPair
type jsiiProxy_IKeyPair struct {
	internal.Type__interfacesawsec2IKeyPairRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKeyPair) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IKeyPair) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IKeyPair) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyPair) KeyPairRef() *interfacesawsec2.KeyPairReference {
	var returns *interfacesawsec2.KeyPairReference
	_jsii_.Get(
		j,
		"keyPairRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyPair) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyPair) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

