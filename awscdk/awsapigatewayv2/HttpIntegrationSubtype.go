package awsapigatewayv2


// Supported integration subtypes.
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
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html
//
type HttpIntegrationSubtype string

const (
	// EventBridge PutEvents integration.
	HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS HttpIntegrationSubtype = "EVENTBRIDGE_PUT_EVENTS"
	// SQS SendMessage integration.
	HttpIntegrationSubtype_SQS_SEND_MESSAGE HttpIntegrationSubtype = "SQS_SEND_MESSAGE"
	// SQS ReceiveMessage integration,.
	HttpIntegrationSubtype_SQS_RECEIVE_MESSAGE HttpIntegrationSubtype = "SQS_RECEIVE_MESSAGE"
	// SQS DeleteMessage integration,.
	HttpIntegrationSubtype_SQS_DELETE_MESSAGE HttpIntegrationSubtype = "SQS_DELETE_MESSAGE"
	// SQS PurgeQueue integration.
	HttpIntegrationSubtype_SQS_PURGE_QUEUE HttpIntegrationSubtype = "SQS_PURGE_QUEUE"
	// AppConfig GetConfiguration integration.
	HttpIntegrationSubtype_APPCONFIG_GET_CONFIGURATION HttpIntegrationSubtype = "APPCONFIG_GET_CONFIGURATION"
	// Kinesis PutRecord integration.
	HttpIntegrationSubtype_KINESIS_PUT_RECORD HttpIntegrationSubtype = "KINESIS_PUT_RECORD"
	// Step Functions StartExecution integration.
	HttpIntegrationSubtype_STEPFUNCTIONS_START_EXECUTION HttpIntegrationSubtype = "STEPFUNCTIONS_START_EXECUTION"
	// Step Functions StartSyncExecution integration.
	HttpIntegrationSubtype_STEPFUNCTIONS_START_SYNC_EXECUTION HttpIntegrationSubtype = "STEPFUNCTIONS_START_SYNC_EXECUTION"
	// Step Functions StopExecution integration.
	HttpIntegrationSubtype_STEPFUNCTIONS_STOP_EXECUTION HttpIntegrationSubtype = "STEPFUNCTIONS_STOP_EXECUTION"
)

