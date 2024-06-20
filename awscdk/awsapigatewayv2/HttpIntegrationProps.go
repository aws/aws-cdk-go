package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi httpApi
//   var integrationCredentials integrationCredentials
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpIntegrationProps := &HttpIntegrationProps{
//   	HttpApi: httpApi,
//   	IntegrationType: awscdk.Aws_apigatewayv2.HttpIntegrationType_HTTP_PROXY,
//
//   	// the properties below are optional
//   	ConnectionId: jsii.String("connectionId"),
//   	ConnectionType: awscdk.*Aws_apigatewayv2.HttpConnectionType_VPC_LINK,
//   	Credentials: integrationCredentials,
//   	IntegrationSubtype: awscdk.*Aws_apigatewayv2.HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	IntegrationUri: jsii.String("integrationUri"),
//   	Method: awscdk.*Aws_apigatewayv2.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	PayloadFormatVersion: payloadFormatVersion,
//   	SecureServerName: jsii.String("secureServerName"),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type HttpIntegrationProps struct {
	// The HTTP API to which this integration should be bound.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// Integration type.
	IntegrationType HttpIntegrationType `field:"required" json:"integrationType" yaml:"integrationType"`
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
	// Integration subtype.
	//
	// Used for AWS Service integrations, specifies the target of the integration.
	// Default: - none, required if no `integrationUri` is defined.
	//
	IntegrationSubtype HttpIntegrationSubtype `field:"optional" json:"integrationSubtype" yaml:"integrationSubtype"`
	// Integration URI.
	//
	// This will be the function ARN in the case of `HttpIntegrationType.AWS_PROXY`,
	// or HTTP URL in the case of `HttpIntegrationType.HTTP_PROXY`.
	// Default: - none, required if no `integrationSubtype` is defined.
	//
	IntegrationUri *string `field:"optional" json:"integrationUri" yaml:"integrationUri"`
	// The HTTP method to use when calling the underlying HTTP proxy.
	// Default: - none. required if the integration type is `HttpIntegrationType.HTTP_PROXY`.
	//
	Method HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	ParameterMapping ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The version of the payload format.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Default: - defaults to latest in the case of HttpIntegrationType.AWS_PROXY`, irrelevant otherwise.
	//
	PayloadFormatVersion PayloadFormatVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Specifies the TLS configuration for a private integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

