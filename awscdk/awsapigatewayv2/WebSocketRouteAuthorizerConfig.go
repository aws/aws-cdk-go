package awsapigatewayv2


// Results of binding an authorizer to an WebSocket route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketRouteAuthorizerConfig := &WebSocketRouteAuthorizerConfig{
//   	AuthorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	AuthorizerId: jsii.String("authorizerId"),
//   }
//
type WebSocketRouteAuthorizerConfig struct {
	// The type of authorization.
	//
	// Possible values are:
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// The authorizer id.
	// Default: - No authorizer id (useful for AWS_IAM route authorizer).
	//
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
}

