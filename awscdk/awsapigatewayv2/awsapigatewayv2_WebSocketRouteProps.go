package awsapigatewayv2


// Properties to initialize a new Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var webSocketApi webSocketApi
//   var webSocketRouteAuthorizer iWebSocketRouteAuthorizer
//   var webSocketRouteIntegration webSocketRouteIntegration
//
//   webSocketRouteProps := &webSocketRouteProps{
//   	integration: webSocketRouteIntegration,
//   	routeKey: jsii.String("routeKey"),
//   	webSocketApi: webSocketApi,
//
//   	// the properties below are optional
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizer: webSocketRouteAuthorizer,
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

