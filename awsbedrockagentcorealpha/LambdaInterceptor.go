package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda-based interceptor for Gateway.
//
// Interceptors allow you to run custom code during each invocation of your gateway:
// - REQUEST interceptors execute before calling the target
// - RESPONSE interceptors execute after the target responds.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ Function
//
//   lambdaInterceptor := bedrock_agentcore_alpha.LambdaInterceptor_ForRequest(function_, &InterceptorOptions{
//   	PassRequestHeaders: jsii.Boolean(false),
//   })
//
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-interceptors.html
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type LambdaInterceptor interface {
	IInterceptor
	// The interception point (REQUEST or RESPONSE).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	InterceptionPoint() InterceptionPoint
	// Binds this Lambda interceptor to a Gateway.
	//
	// This method:
	// 1. Grants the Gateway's IAM role permission to invoke the Lambda function
	// 2. Returns the CloudFormation configuration for this interceptor
	//
	// Returns: Configuration for CloudFormation rendering.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig
}

// The jsii proxy struct for LambdaInterceptor
type jsiiProxy_LambdaInterceptor struct {
	jsiiProxy_IInterceptor
}

func (j *jsiiProxy_LambdaInterceptor) InterceptionPoint() InterceptionPoint {
	var returns InterceptionPoint
	_jsii_.Get(
		j,
		"interceptionPoint",
		&returns,
	)
	return returns
}


// Create a REQUEST interceptor that executes before the gateway calls the target.
//
// **Important:** When this interceptor is added to a gateway, the gateway's IAM role
// will automatically be granted `lambda:InvokeFunction` permission on the specified
// Lambda function.
//
// Returns: A configured LambdaInterceptor for request interception.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func LambdaInterceptor_ForRequest(lambdaFunction awslambda.IFunction, options *InterceptorOptions) LambdaInterceptor {
	_init_.Initialize()

	if err := validateLambdaInterceptor_ForRequestParameters(lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaInterceptor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaInterceptor",
		"forRequest",
		[]interface{}{lambdaFunction, options},
		&returns,
	)

	return returns
}

// Create a RESPONSE interceptor that executes after the target responds.
//
// **Important:** When this interceptor is added to a gateway, the gateway's IAM role
// will automatically be granted `lambda:InvokeFunction` permission on the specified
// Lambda function.
//
// Returns: A configured LambdaInterceptor for response interception.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func LambdaInterceptor_ForResponse(lambdaFunction awslambda.IFunction, options *InterceptorOptions) LambdaInterceptor {
	_init_.Initialize()

	if err := validateLambdaInterceptor_ForResponseParameters(lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaInterceptor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaInterceptor",
		"forResponse",
		[]interface{}{lambdaFunction, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInterceptor) Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig {
	if err := l.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *InterceptorBindConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

