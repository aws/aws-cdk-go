package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator"
	"github.com/aws/constructs-go/constructs/v10"
)

// The interface of the EndpointGroup.
type IEndpointGroup interface {
	interfacesawsglobalaccelerator.IEndpointGroupRef
	awscdk.IResource
	// EndpointGroup ARN.
	EndpointGroupArn() *string
}

// The jsii proxy for IEndpointGroup
type jsiiProxy_IEndpointGroup struct {
	internal.Type__interfacesawsglobalacceleratorIEndpointGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IEndpointGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroup) EndpointGroupRef() *interfacesawsglobalaccelerator.EndpointGroupReference {
	var returns *interfacesawsglobalaccelerator.EndpointGroupReference
	_jsii_.Get(
		j,
		"endpointGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

