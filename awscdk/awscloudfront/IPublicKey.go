package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Public Key.
type IPublicKey interface {
	interfacesawscloudfront.IPublicKeyRef
	awscdk.IResource
	// The ID of the key group.
	PublicKeyId() *string
}

// The jsii proxy for IPublicKey
type jsiiProxy_IPublicKey struct {
	internal.Type__interfacesawscloudfrontIPublicKeyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPublicKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPublicKey) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPublicKey) PublicKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKey) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKey) PublicKeyRef() *interfacesawscloudfront.PublicKeyReference {
	var returns *interfacesawscloudfront.PublicKeyReference
	_jsii_.Get(
		j,
		"publicKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

