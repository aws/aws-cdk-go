package awscdkapigatewayv2alpha


// Results of binding an authorizer to an WebSocket route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   webSocketRouteAuthorizerConfig := &WebSocketRouteAuthorizerConfig{
//   	AuthorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	AuthorizerId: jsii.String("authorizerId"),
//   }
//
// Deprecated.
type WebSocketRouteAuthorizerConfig struct {
	// The type of authorization.
	//
	// Possible values are:
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Deprecated.
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// The authorizer id.
	// Default: - No authorizer id (useful for AWS_IAM route authorizer).
	//
	// Deprecated.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
}

