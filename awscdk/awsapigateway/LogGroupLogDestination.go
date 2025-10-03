package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Use CloudWatch Logs as a custom access log destination for API Gateway.
//
// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		AccessLogFormat: apigateway.AccessLogFormat_Custom(
//   		fmt.Sprintf("%v %v %v\n      %v %v", apigateway.AccessLogField_ContextRequestId(), apigateway.AccessLogField_ContextErrorMessage(), apigateway.AccessLogField_ContextErrorMessageString(), apigateway.AccessLogField_ContextAuthorizerError(), apigateway.AccessLogField_ContextAuthorizerIntegrationStatus())),
//   	},
//   })
//
type LogGroupLogDestination interface {
	IAccessLogDestination
	// Binds this destination to the CloudWatch Logs.
	Bind(_stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy struct for LogGroupLogDestination
type jsiiProxy_LogGroupLogDestination struct {
	jsiiProxy_IAccessLogDestination
}

func NewLogGroupLogDestination(logGroup awslogs.ILogGroup) LogGroupLogDestination {
	_init_.Initialize()

	if err := validateNewLogGroupLogDestinationParameters(logGroup); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogGroupLogDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

func NewLogGroupLogDestination_Override(l LogGroupLogDestination, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		l,
	)
}

func (l *jsiiProxy_LogGroupLogDestination) Bind(_stage IStage) *AccessLogDestinationConfig {
	if err := l.validateBindParameters(_stage); err != nil {
		panic(err)
	}
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_stage},
		&returns,
	)

	return returns
}

