package awscdkapigatewayv2alpha


// Reference to an WebSocket authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   webSocketAuthorizerAttributes := &WebSocketAuthorizerAttributes{
//   	AuthorizerId: jsii.String("authorizerId"),
//   	AuthorizerType: jsii.String("authorizerType"),
//   }
//
// Deprecated.
type WebSocketAuthorizerAttributes struct {
	// Id of the Authorizer.
	// Deprecated.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Deprecated.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

