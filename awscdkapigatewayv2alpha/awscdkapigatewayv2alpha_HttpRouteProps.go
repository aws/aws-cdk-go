// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Properties to initialize a new Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var httpApi httpApi
//   var httpRouteAuthorizer iHttpRouteAuthorizer
//   var httpRouteIntegration httpRouteIntegration
//   var httpRouteKey httpRouteKey
//
//   httpRouteProps := &httpRouteProps{
//   	httpApi: httpApi,
//   	integration: httpRouteIntegration,
//   	routeKey: httpRouteKey,
//
//   	// the properties below are optional
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizer: httpRouteAuthorizer,
//   }
//
// Experimental.
type HttpRouteProps struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// the API the route is associated with.
	// Experimental.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// The key to this route.
	//
	// This is a combination of an HTTP method and an HTTP path.
	// Experimental.
	RouteKey HttpRouteKey `field:"required" json:"routeKey" yaml:"routeKey"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will be merged with the scopes from the attached authorizer.
	// Experimental.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Authorizer for a WebSocket API or an HTTP API.
	// Experimental.
	Authorizer IHttpRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
}

