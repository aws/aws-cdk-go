package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// The StepFunctions integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import sfn "github.com/aws/aws-cdk-go/awscdk"
//
//   var stateMachine StateMachine
//   var httpApi HttpApi
//
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/start"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpStepFunctionsIntegration(jsii.String("StartExecutionIntegration"), &HttpStepFunctionsIntegrationProps{
//   		StateMachine: *StateMachine,
//   		Subtype: apigwv2.HttpIntegrationSubtype_STEPFUNCTIONS_START_EXECUTION,
//   	}),
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/start-sync"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpStepFunctionsIntegration(jsii.String("StartSyncExecutionIntegration"), &HttpStepFunctionsIntegrationProps{
//   		StateMachine: *StateMachine,
//   		Subtype: apigwv2.HttpIntegrationSubtype_STEPFUNCTIONS_START_SYNC_EXECUTION,
//   	}),
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/stop"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpStepFunctionsIntegration(jsii.String("StopExecutionIntegration"), &HttpStepFunctionsIntegrationProps{
//   		StateMachine: *StateMachine,
//   		Subtype: apigwv2.HttpIntegrationSubtype_STEPFUNCTIONS_STOP_EXECUTION,
//   		// For the `STOP_EXECUTION` subtype, it is necessary to specify the `executionArn`.
//   		ParameterMapping: apigwv2.NewParameterMapping().Custom(jsii.String("ExecutionArn"), jsii.String("$request.querystring.executionArn")),
//   	}),
//   })
//
type HttpStepFunctionsIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Bind this integration to the route.
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpStepFunctionsIntegration
type jsiiProxy_HttpStepFunctionsIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func NewHttpStepFunctionsIntegration(id *string, props *HttpStepFunctionsIntegrationProps) HttpStepFunctionsIntegration {
	_init_.Initialize()

	if err := validateNewHttpStepFunctionsIntegrationParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpStepFunctionsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpStepFunctionsIntegration",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewHttpStepFunctionsIntegration_Override(h HttpStepFunctionsIntegration, id *string, props *HttpStepFunctionsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpStepFunctionsIntegration",
		[]interface{}{id, props},
		h,
	)
}

func (h *jsiiProxy_HttpStepFunctionsIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpStepFunctionsIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

