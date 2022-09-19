// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// The integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var httpApi httpApi
//   var integrationCredentials integrationCredentials
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpIntegrationProps := &httpIntegrationProps{
//   	httpApi: httpApi,
//   	integrationType: apigatewayv2_alpha.httpIntegrationType_HTTP_PROXY,
//
//   	// the properties below are optional
//   	connectionId: jsii.String("connectionId"),
//   	connectionType: apigatewayv2_alpha.httpConnectionType_VPC_LINK,
//   	credentials: integrationCredentials,
//   	integrationSubtype: apigatewayv2_alpha.httpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	integrationUri: jsii.String("integrationUri"),
//   	method: apigatewayv2_alpha.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   	payloadFormatVersion: payloadFormatVersion,
//   	secureServerName: jsii.String("secureServerName"),
//   }
//
// Experimental.
type HttpIntegrationProps struct {
	// The HTTP API to which this integration should be bound.
	// Experimental.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// Integration type.
	// Experimental.
	IntegrationType HttpIntegrationType `field:"required" json:"integrationType" yaml:"integrationType"`
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
	// Integration subtype.
	//
	// Used for AWS Service integrations, specifies the target of the integration.
	// Experimental.
	IntegrationSubtype HttpIntegrationSubtype `field:"optional" json:"integrationSubtype" yaml:"integrationSubtype"`
	// Integration URI.
	//
	// This will be the function ARN in the case of `HttpIntegrationType.AWS_PROXY`,
	// or HTTP URL in the case of `HttpIntegrationType.HTTP_PROXY`.
	// Experimental.
	IntegrationUri *string `field:"optional" json:"integrationUri" yaml:"integrationUri"`
	// The HTTP method to use when calling the underlying HTTP proxy.
	// Experimental.
	Method HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The version of the payload format.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion PayloadFormatVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Specifies the TLS configuration for a private integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
}

