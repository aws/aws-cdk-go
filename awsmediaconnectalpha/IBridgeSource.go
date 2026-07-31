package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface Bridge source.
// Experimental.
type IBridgeSource interface {
	interfacesawsmediaconnect.IBridgeSourceRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the bridge that owns this source.
	// Experimental.
	BridgeArn() *string
	// The name of the bridge source.
	// Experimental.
	BridgeSourceName() *string
}

// The jsii proxy for IBridgeSource
type jsiiProxy_IBridgeSource struct {
	internal.Type__interfacesawsmediaconnectIBridgeSourceRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBridgeSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IBridgeSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBridgeSource) BridgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSource) BridgeSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSource) BridgeSourceRef() *interfacesawsmediaconnect.BridgeSourceReference {
	var returns *interfacesawsmediaconnect.BridgeSourceReference
	_jsii_.Get(
		j,
		"bridgeSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSource) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

