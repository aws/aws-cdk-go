// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha


// Properties to initialize HttpJwtAuthorizer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := awscdkapigatewayv2authorizersalpha.NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &HttpJwtAuthorizerProps{
//   	JwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
// Experimental.
type HttpJwtAuthorizerProps struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `field:"required" json:"jwtAudience" yaml:"jwtAudience"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
}

