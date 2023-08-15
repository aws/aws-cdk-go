package awscdkapigatewayv2alpha


// Config returned back as a result of the bind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var integrationCredentials integrationCredentials
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpRouteIntegrationConfig := &HttpRouteIntegrationConfig{
//   	PayloadFormatVersion: payloadFormatVersion,
//   	Type: apigatewayv2_alpha.HttpIntegrationType_HTTP_PROXY,
//
//   	// the properties below are optional
//   	ConnectionId: jsii.String("connectionId"),
//   	ConnectionType: apigatewayv2_alpha.HttpConnectionType_VPC_LINK,
//   	Credentials: integrationCredentials,
//   	Method: apigatewayv2_alpha.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	SecureServerName: jsii.String("secureServerName"),
//   	Subtype: apigatewayv2_alpha.HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	Uri: jsii.String("uri"),
//   }
//
// Experimental.
type HttpRouteIntegrationConfig struct {
	// Payload format version in the case of lambda proxy integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Default: - undefined.
	//
	// Experimental.
	PayloadFormatVersion PayloadFormatVersion `field:"required" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Integration type.
	// Experimental.
	Type HttpIntegrationType `field:"required" json:"type" yaml:"type"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	// Default: - undefined.
	//
	// Experimental.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	// Default: HttpConnectionType.INTERNET
	//
	// Experimental.
	ConnectionType HttpConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The credentials with which to invoke the integration.
	// Default: - no credentials, use resource-based permissions on supported AWS services.
	//
	// Experimental.
	Credentials IntegrationCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The HTTP method that must be used to invoke the underlying proxy.
	//
	// Required for `HttpIntegrationType.HTTP_PROXY`
	// Default: - undefined.
	//
	// Experimental.
	Method HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	// Experimental.
	ParameterMapping ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// Integration subtype.
	// Default: - none, required if no `integrationUri` is defined.
	//
	// Experimental.
	Subtype HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
	// Integration URI.
	// Default: - none, required if no `integrationSubtype` is defined.
	//
	// Experimental.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

