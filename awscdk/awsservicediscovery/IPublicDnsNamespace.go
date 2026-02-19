package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

type IPublicDnsNamespace interface {
	INamespace
	interfacesawsservicediscovery.IPublicDnsNamespaceRef
}

// The jsii proxy for IPublicDnsNamespace
type jsiiProxy_IPublicDnsNamespace struct {
	jsiiProxy_INamespace
	internal.Type__interfacesawsservicediscoveryIPublicDnsNamespaceRef
}

func (i *jsiiProxy_IPublicDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPublicDnsNamespace) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPublicDnsNamespace) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) PublicDnsNamespaceRef() *interfacesawsservicediscovery.PublicDnsNamespaceReference {
	var returns *interfacesawsservicediscovery.PublicDnsNamespaceReference
	_jsii_.Get(
		j,
		"publicDnsNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

