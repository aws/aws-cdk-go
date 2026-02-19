package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents gateway response resource.
type IGatewayResponse interface {
	interfacesawsapigateway.IGatewayResponseRef
	awscdk.IResource
}

// The jsii proxy for IGatewayResponse
type jsiiProxy_IGatewayResponse struct {
	internal.Type__interfacesawsapigatewayIGatewayResponseRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGatewayResponse) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGatewayResponse) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGatewayResponse) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayResponse) GatewayResponseRef() *interfacesawsapigateway.GatewayResponseReference {
	var returns *interfacesawsapigateway.GatewayResponseReference
	_jsii_.Get(
		j,
		"gatewayResponseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayResponse) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

