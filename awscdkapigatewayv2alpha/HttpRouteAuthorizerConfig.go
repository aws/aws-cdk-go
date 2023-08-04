package awscdkapigatewayv2alpha


// Results of binding an authorizer to an http route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   httpRouteAuthorizerConfig := &HttpRouteAuthorizerConfig{
//   	AuthorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	AuthorizerId: jsii.String("authorizerId"),
//   }
//
// Experimental.
type HttpRouteAuthorizerConfig struct {
	// The type of authorization.
	//
	// Possible values are:
	// - AWS_IAM - IAM Authorizer
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Experimental.
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// The list of OIDC scopes to include in the authorization.
	// Default: - no authorization scopes.
	//
	// Experimental.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The authorizer id.
	// Default: - No authorizer id (useful for AWS_IAM route authorizer).
	//
	// Experimental.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
}

