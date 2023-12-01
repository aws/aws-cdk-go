package awscdkapigatewayv2alpha


// Supported integration types.
// Deprecated.
type HttpIntegrationType string

const (
	// Integration type is an HTTP proxy.
	//
	// For integrating the route or method request with an HTTP endpoint, with the
	// client request passed through as-is. This is also referred to as HTTP proxy
	// integration. For HTTP API private integrations, use an HTTP_PROXY integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-http.html
	//
	// Deprecated.
	HttpIntegrationType_HTTP_PROXY HttpIntegrationType = "HTTP_PROXY"
	// Integration type is an AWS proxy.
	//
	// For integrating the route or method request with a Lambda function or other
	// AWS service action. This integration is also referred to as a Lambda proxy
	// integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Deprecated.
	HttpIntegrationType_AWS_PROXY HttpIntegrationType = "AWS_PROXY"
)

