package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdaeventsources/internal"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var model model
//   var requestValidator requestValidator
//
//   apiEventSource := awscdk.Aws_lambda_event_sources.NewApiEventSource(jsii.String("method"), jsii.String("path"), &methodOptions{
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   	authorizer: authorizer,
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			responseModels: map[string]iModel{
//   				"responseModelsKey": model,
//   			},
//   			responseParameters: map[string]*bool{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	operationName: jsii.String("operationName"),
//   	requestModels: map[string]*iModel{
//   		"requestModelsKey": model,
//   	},
//   	requestParameters: map[string]*bool{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	requestValidator: requestValidator,
//   	requestValidatorOptions: &requestValidatorOptions{
//   		requestValidatorName: jsii.String("requestValidatorName"),
//   		validateRequestBody: jsii.Boolean(false),
//   		validateRequestParameters: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type ApiEventSource interface {
	awslambda.IEventSource
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	// Experimental.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for ApiEventSource
type jsiiProxy_ApiEventSource struct {
	internal.Type__awslambdaIEventSource
}

// Experimental.
func NewApiEventSource(method *string, path *string, options *awsapigateway.MethodOptions) ApiEventSource {
	_init_.Initialize()

	if err := validateNewApiEventSourceParameters(method, path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiEventSource{}

	_jsii_.Create(
		"monocdk.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewApiEventSource_Override(a ApiEventSource, method *string, path *string, options *awsapigateway.MethodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_event_sources.ApiEventSource",
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

