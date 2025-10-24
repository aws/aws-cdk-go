package awsapigatewayv2


// Properties to initialize a new Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi HttpApi
//   var httpRouteAuthorizer IHttpRouteAuthorizer
//   var httpRouteIntegration HttpRouteIntegration
//   var httpRouteKey HttpRouteKey
//
//   httpRouteProps := &HttpRouteProps{
//   	HttpApi: httpApi,
//   	Integration: httpRouteIntegration,
//   	RouteKey: httpRouteKey,
//
//   	// the properties below are optional
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	Authorizer: httpRouteAuthorizer,
//   }
//
type HttpRouteProps struct {
	// The integration to be configured on this route.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// the API the route is associated with.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// The key to this route.
	//
	// This is a combination of an HTTP method and an HTTP path.
	RouteKey HttpRouteKey `field:"required" json:"routeKey" yaml:"routeKey"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will be merged with the scopes from the attached authorizer.
	// Default: - no additional authorization scopes.
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Authorizer for a WebSocket API or an HTTP API.
	// Default: - No authorizer.
	//
	Authorizer IHttpRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
}

