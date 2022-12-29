package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
)

// Interface for EventBus Connections.
type IConnection interface {
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
	internal.Type__awscdkIResource
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

