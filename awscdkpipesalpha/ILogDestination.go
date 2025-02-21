package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Log destination base class.
// Experimental.
type ILogDestination interface {
	// Bind the log destination to the pipe.
	// Experimental.
	Bind(pipe IPipe) *LogDestinationConfig
	// Grant the pipe role to push to the log destination.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy for ILogDestination
type jsiiProxy_ILogDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogDestination) Bind(pipe IPipe) *LogDestinationConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *LogDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogDestination) GrantPush(grantee awsiam.IRole) {
	if err := i.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantPush",
		[]interface{}{grantee},
	)
}

