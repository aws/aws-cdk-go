package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface Bridge output.
// Experimental.
type IBridgeOutput interface {
	interfacesawsmediaconnect.IBridgeOutputRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the bridge that owns this output.
	// Experimental.
	BridgeArn() *string
	// The name of the bridge output.
	// Experimental.
	BridgeOutputName() *string
}

// The jsii proxy for IBridgeOutput
type jsiiProxy_IBridgeOutput struct {
	internal.Type__interfacesawsmediaconnectIBridgeOutputRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBridgeOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IBridgeOutput) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBridgeOutput) BridgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeOutput) BridgeOutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeOutputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeOutput) BridgeOutputRef() *interfacesawsmediaconnect.BridgeOutputReference {
	var returns *interfacesawsmediaconnect.BridgeOutputReference
	_jsii_.Get(
		j,
		"bridgeOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeOutput) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeOutput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

