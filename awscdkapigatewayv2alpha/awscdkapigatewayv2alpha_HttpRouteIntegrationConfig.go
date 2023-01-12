// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Config returned back as a result of the bind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var integrationCredentials integrationCredentials
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpRouteIntegrationConfig := &httpRouteIntegrationConfig{
//   	payloadFormatVersion: payloadFormatVersion,
//   	type: apigatewayv2_alpha.httpIntegrationType_HTTP_PROXY,
//
//   	// the properties below are optional
//   	connectionId: jsii.String("connectionId"),
//   	connectionType: apigatewayv2_alpha.httpConnectionType_VPC_LINK,
//   	credentials: integrationCredentials,
//   	method: apigatewayv2_alpha.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   	secureServerName: jsii.String("secureServerName"),
//   	subtype: apigatewayv2_alpha.httpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	uri: jsii.String("uri"),
//   }
//
// Experimental.
type HttpRouteIntegrationConfig struct {
	// Payload format version in the case of lambda proxy integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion PayloadFormatVersion `field:"required" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Integration type.
	// Experimental.
	Type HttpIntegrationType `field:"required" json:"type" yaml:"type"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	// Experimental.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	// Experimental.
	ConnectionType HttpConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The credentials with which to invoke the integration.
	// Experimental.
	Credentials IntegrationCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The HTTP method that must be used to invoke the underlying proxy.
	//
	// Required for `HttpIntegrationType.HTTP_PROXY`
	// Experimental.
	Method HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// Integration subtype.
	// Experimental.
	Subtype HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
	// Integration URI.
	// Experimental.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

