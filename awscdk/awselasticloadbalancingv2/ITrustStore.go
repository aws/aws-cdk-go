package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Trust Store.
type ITrustStore interface {
	awscdk.IResource
	interfacesawselasticloadbalancingv2.ITrustStoreRef
	// The ARN of the trust store.
	TrustStoreArn() *string
	// The name of the trust store.
	TrustStoreName() *string
}

// The jsii proxy for ITrustStore
type jsiiProxy_ITrustStore struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawselasticloadbalancingv2ITrustStoreRef
}

func (i *jsiiProxy_ITrustStore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ITrustStore) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITrustStore) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) TrustStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) TrustStoreRef() *interfacesawselasticloadbalancingv2.TrustStoreReference {
	var returns *interfacesawselasticloadbalancingv2.TrustStoreReference
	_jsii_.Get(
		j,
		"trustStoreRef",
		&returns,
	)
	return returns
}

