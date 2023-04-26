// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Properties to initialize a new Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var webSocketApi webSocketApi
//   var webSocketRouteAuthorizer iWebSocketRouteAuthorizer
//   var webSocketRouteIntegration webSocketRouteIntegration
//
//   webSocketRouteProps := &WebSocketRouteProps{
//   	Integration: webSocketRouteIntegration,
//   	RouteKey: jsii.String("routeKey"),
//   	WebSocketApi: webSocketApi,
//
//   	// the properties below are optional
//   	ApiKeyRequired: jsii.Boolean(false),
//   	Authorizer: webSocketRouteAuthorizer,
//   	ReturnResponse: jsii.Boolean(false),
//   }
//
// Experimental.
type WebSocketRouteProps struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration WebSocketRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The authorize to this route.
	//
	// You can only set authorizer to a $connect route.
	// Experimental.
	Authorizer IWebSocketRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Should the route send a response to the client.
	// Experimental.
	ReturnResponse *bool `field:"optional" json:"returnResponse" yaml:"returnResponse"`
	// The key to this route.
	// Experimental.
	RouteKey *string `field:"required" json:"routeKey" yaml:"routeKey"`
	// The API the route is associated with.
	// Experimental.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
	// Whether the route requires an API Key to be provided.
	// Experimental.
	ApiKeyRequired *bool `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
}

