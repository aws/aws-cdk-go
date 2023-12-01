package awsapigatewayv2


// Supported integration subtypes.
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

