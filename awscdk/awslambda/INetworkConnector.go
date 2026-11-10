package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda Network Connector.
type INetworkConnector interface {
	interfacesawslambda.INetworkConnectorRef
	awscdk.IResource
	// The qualified ARN of the connector's latest version.
	//
	// WebFunction pins a revision to this. For a connector defined in this app the value changes
	// on every update to the connector, which replaces the revision that references it.
	LatestVersionArn() *string
	// The unqualified ARN of the network connector.
	//
	// Identifies the connector as a whole, across its versions: use it in IAM policies and for
	// consumers that resolve the current version at run time. It does not name a version, so
	// WebFunction — which pins a revision to one version — does not accept it.
	NetworkConnectorArn() *string
}

// The jsii proxy for INetworkConnector
type jsiiProxy_INetworkConnector struct {
	internal.Type__interfacesawslambdaINetworkConnectorRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INetworkConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_INetworkConnector) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkConnector) LatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnector) NetworkConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnector) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnector) NetworkConnectorRef() *interfacesawslambda.NetworkConnectorReference {
	var returns *interfacesawslambda.NetworkConnectorReference
	_jsii_.Get(
		j,
		"networkConnectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

