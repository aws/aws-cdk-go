package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties to initialize `HttpStepFunctionsIntegration`.
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
type HttpStepFunctionsIntegrationProps struct {
	// Statemachine that Integrates with API Gateway.
	StateMachine awsstepfunctions.StateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	//
	// When the subtype is either `START_EXECUTION` or `START_SYNC_EXECUTION`,
	// it is necessary to specify the `StateMachineArn`.
	// Conversely, when the subtype is `STOP_EXECUTION`, the `ExecutionArn` must be specified.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: - specify only `StateMachineArn`.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The subtype of the HTTP integration.
	//
	// Only subtypes starting with STEPFUNCTIONS_ can be specified.
	// Default: HttpIntegrationSubtype.STEPFUNCTIONS_START_EXECUTION
	//
	Subtype awsapigatewayv2.HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
}

