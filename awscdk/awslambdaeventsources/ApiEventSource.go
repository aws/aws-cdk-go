package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer Authorizer
//   var model Model
//   var requestValidator RequestValidator
//
//   apiEventSource := awscdk.Aws_lambda_event_sources.NewApiEventSource(jsii.String("method"), jsii.String("path"), &MethodOptions{
//   	ApiKeyRequired: jsii.Boolean(false),
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	AuthorizationType: awscdk.Aws_apigateway.AuthorizationType_NONE,
//   	Authorizer: authorizer,
//   	MethodResponses: []MethodResponse{
//   		&MethodResponse{
//   			StatusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			ResponseModels: map[string]IModel{
//   				"responseModelsKey": model,
//   			},
//   			ResponseParameters: map[string]*bool{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	OperationName: jsii.String("operationName"),
//   	RequestModels: map[string]IModel{
//   		"requestModelsKey": model,
//   	},
//   	RequestParameters: map[string]*bool{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	RequestValidator: requestValidator,
//   	RequestValidatorOptions: &RequestValidatorOptions{
//   		RequestValidatorName: jsii.String("requestValidatorName"),
//   		ValidateRequestBody: jsii.Boolean(false),
//   		ValidateRequestParameters: jsii.Boolean(false),
//   	},
//   })
//
type ApiEventSource interface {
	awslambda.IEventSource
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for ApiEventSource
type jsiiProxy_ApiEventSource struct {
	internal.Type__awslambdaIEventSource
}

func NewApiEventSource(method *string, path *string, options *awsapigateway.MethodOptions) ApiEventSource {
	_init_.Initialize()

	if err := validateNewApiEventSourceParameters(method, path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		&j,
	)

	return &j
}

func NewApiEventSource_Override(a ApiEventSource, method *string, path *string, options *awsapigateway.MethodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		a,
	)
}

func (a *jsiiProxy_ApiEventSource) Bind(target awslambda.IFunction) {
	if err := a.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{target},
	)
}

