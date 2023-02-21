package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A connection handler for client VPN endpoints.
type IClientVpnConnectionHandler interface {
	// The ARN of the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
}

// The jsii proxy for IClientVpnConnectionHandler
type jsiiProxy_IClientVpnConnectionHandler struct {
	_ byte // padding
}

func (j *jsiiProxy_IClientVpnConnectionHandler) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnConnectionHandler) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

