package awsapigatewayv2


// Reference to an WebSocket authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketAuthorizerAttributes := &webSocketAuthorizerAttributes{
//   	authorizerId: jsii.String("authorizerId"),
//   	authorizerType: jsii.String("authorizerType"),
//   }
//
// Experimental.
type WebSocketAuthorizerAttributes struct {
	// Id of the Authorizer.
	// Experimental.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Experimental.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

