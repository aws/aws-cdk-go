package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Network.
// Experimental.
type INetwork interface {
	interfacesawsmedialive.INetworkRef
	awscdk.IResource
	// The ARN of the network.
	// Experimental.
	NetworkArn() *string
	// The ID of the network.
	// Experimental.
	NetworkId() *string
}

// The jsii proxy for INetwork
type jsiiProxy_INetwork struct {
	internal.Type__interfacesawsmedialiveINetworkRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INetwork) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_INetwork) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetwork) NetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetwork) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetwork) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetwork) NetworkRef() *interfacesawsmedialive.NetworkReference {
	var returns *interfacesawsmedialive.NetworkReference
	_jsii_.Get(
		j,
		"networkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetwork) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

