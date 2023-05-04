// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Payload format version for lambda authorizers.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
//
// Experimental.
type AuthorizerPayloadVersion string

const (
	// Version 1.0.
	// Experimental.
	AuthorizerPayloadVersion_VERSION_1_0 AuthorizerPayloadVersion = "VERSION_1_0"
	// Version 2.0.
	// Experimental.
	AuthorizerPayloadVersion_VERSION_2_0 AuthorizerPayloadVersion = "VERSION_2_0"
)

