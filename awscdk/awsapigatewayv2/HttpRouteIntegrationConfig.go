package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Config returned back as a result of the bind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var integrationCredentials IntegrationCredentials
//   var parameterMapping ParameterMapping
//   var payloadFormatVersion PayloadFormatVersion
//
//   httpRouteIntegrationConfig := &HttpRouteIntegrationConfig{
//   	PayloadFormatVersion: payloadFormatVersion,
//   	Type: awscdk.Aws_apigatewayv2.HttpIntegrationType_HTTP_PROXY,
//
//   	// the properties below are optional
//   	ConnectionId: jsii.String("connectionId"),
//   	ConnectionType: awscdk.*Aws_apigatewayv2.HttpConnectionType_VPC_LINK,
//   	Credentials: integrationCredentials,
//   	Method: awscdk.*Aws_apigatewayv2.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	SecureServerName: jsii.String("secureServerName"),
//   	Subtype: awscdk.*Aws_apigatewayv2.HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	Uri: jsii.String("uri"),
//   }
//
type HttpRouteIntegrationConfig struct {
	// Payload format version in the case of lambda proxy integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Default: - undefined.
	//
	PayloadFormatVersion PayloadFormatVersion `field:"required" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Integration type.
	Type HttpIntegrationType `field:"required" json:"type" yaml:"type"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	// Default: - undefined.
	//
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	// Default: HttpConnectionType.INTERNET
	//
	ConnectionType HttpConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The credentials with which to invoke the integration.
	// Default: - no credentials, use resource-based permissions on supported AWS services.
	//
	Credentials IntegrationCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The HTTP method that must be used to invoke the underlying proxy.
	//
	// Required for `HttpIntegrationType.HTTP_PROXY`
	// Default: - undefined.
	//
	Method HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	ParameterMapping ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// Integration subtype.
	// Default: - none, required if no `integrationUri` is defined.
	//
	Subtype HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Integration URI.
	// Default: - none, required if no `integrationSubtype` is defined.
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

