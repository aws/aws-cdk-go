package awsapigatewayv2


// Options for the Route with Integration resource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bookStoreDefaultFn Function
//
//
//   getBooksIntegration := awscdk.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
//   bookStoreDefaultIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), bookStoreDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_ANY,
//   	},
//   	Integration: bookStoreDefaultIntegration,
//   })
//
type AddRoutesOptions struct {
	// The integration to be configured on this route.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The path at which all of these routes are configured.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will override the default authorization scopes on the gateway.
	// Set to [] to remove default scopes.
	// Default: - uses defaultAuthorizationScopes if configured on the API, otherwise none.
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Authorizer to be associated to these routes.
	//
	// Use NoneAuthorizer to remove the default authorizer for the api.
	// Default: - uses the default authorizer if one is specified on the HttpApi.
	//
	Authorizer IHttpRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// The HTTP methods to be configured.
	// Default: HttpMethod.ANY
	//
	Methods *[]HttpMethod `field:"optional" json:"methods" yaml:"methods"`
}

