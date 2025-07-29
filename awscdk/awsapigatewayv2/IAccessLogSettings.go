package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Settings for access logging.
type IAccessLogSettings interface {
	// The destination where to write access logs.
	// Default: - No destination.
	//
	Destination() IAccessLogDestination
	// A single line format of access logs of data, as specified by selected $context variables.
	//
	// The format must include either `AccessLogFormat.contextRequestId()`
	// or `AccessLogFormat.contextExtendedRequestId()`.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-logging-variables.html
	//
	// Default: - Common Log Format.
	//
	Format() awsapigateway.AccessLogFormat
}

// The jsii proxy for IAccessLogSettings
type jsiiProxy_IAccessLogSettings struct {
	_ byte // padding
}

func (j *jsiiProxy_IAccessLogSettings) Destination() IAccessLogDestination {
	var returns IAccessLogDestination
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessLogSettings) Format() awsapigateway.AccessLogFormat {
	var returns awsapigateway.AccessLogFormat
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

