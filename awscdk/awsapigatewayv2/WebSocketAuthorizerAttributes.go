package awsapigatewayv2


// Reference to an WebSocket authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketAuthorizerAttributes := &WebSocketAuthorizerAttributes{
//   	AuthorizerId: jsii.String("authorizerId"),
//   	AuthorizerType: jsii.String("authorizerType"),
//   }
//
type WebSocketAuthorizerAttributes struct {
	// Id of the Authorizer.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

