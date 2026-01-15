package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for EventBus Connections.
type IConnection interface {
	interfacesawsevents.IConnectionRef
	awscdk.IResource
	// The ARN of the connection created.
	ConnectionArn() *string
	// The Name for the connection.
	ConnectionName() *string
	// The ARN for the secret created for the connection.
	ConnectionSecretArn() *string
}

// The jsii proxy for IConnection
type jsiiProxy_IConnection struct {
	internal.Type__interfacesawseventsIConnectionRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionRef() *interfacesawsevents.ConnectionReference {
	var returns *interfacesawsevents.ConnectionReference
	_jsii_.Get(
		j,
		"connectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

