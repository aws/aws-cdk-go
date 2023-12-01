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
// Deprecated.
type WebSocketRouteProps struct {
	// The integration to be configured on this route.
	// Deprecated.
	Integration WebSocketRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The authorize to this route.
	//
	// You can only set authorizer to a $connect route.
	// Default: - No Authorizer.
	//
	// Deprecated.
	Authorizer IWebSocketRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Should the route send a response to the client.
	// Default: false.
	//
	// Deprecated.
	ReturnResponse *bool `field:"optional" json:"returnResponse" yaml:"returnResponse"`
	// The key to this route.
	// Deprecated.
	RouteKey *string `field:"required" json:"routeKey" yaml:"routeKey"`
	// The API the route is associated with.
	// Deprecated.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
	// Whether the route requires an API Key to be provided.
	// Default: false.
	//
	// Deprecated.
	ApiKeyRequired *bool `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
}

