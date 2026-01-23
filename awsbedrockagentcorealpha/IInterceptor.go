package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an interceptor that can be bound to a Gateway.
//
// Interceptors allow custom code execution at specific points in the gateway request/response flow.
// Experimental.
type IInterceptor interface {
	// Binds this interceptor to a Gateway.
	//
	// This method is called when the interceptor is added to a gateway. It should:
	// 1. Grant any necessary permissions (e.g., Lambda invoke permissions)
	// 2. Perform any required setup
	// 3. Return the CloudFormation configuration
	//
	// Returns: Configuration that will be rendered to CloudFormation.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig
	// The interception point where this interceptor will be invoked.
	// Experimental.
	InterceptionPoint() InterceptionPoint
}

// The jsii proxy for IInterceptor
type jsiiProxy_IInterceptor struct {
	_ byte // padding
}

func (i *jsiiProxy_IInterceptor) Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig {
	if err := i.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *InterceptorBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInterceptor) InterceptionPoint() InterceptionPoint {
	var returns InterceptionPoint
	_jsii_.Get(
		j,
		"interceptionPoint",
		&returns,
	)
	return returns
}

