package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Use CloudWatch Logs as a custom access log destination for API Gateway.
//
// Example:
//   import apigw "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//   var api httpApi
//   var logGroup logGroup
//
//
//   stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
//   	HttpApi: api,
//   	AccessLogSettings: map[string]interface{}{
//   		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
//   		"format": apigw.AccessLogFormat_clf(),
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
		"aws-cdk-lib.aws_apigatewayv2.LogGroupLogDestination",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

func NewLogGroupLogDestination_Override(l LogGroupLogDestination, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.LogGroupLogDestination",
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

