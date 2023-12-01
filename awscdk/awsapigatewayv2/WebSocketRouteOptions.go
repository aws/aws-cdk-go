package awsapigatewayv2


// Options used to add route to the API.
//
// Example:
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var messageHandler function
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
//   	Integration: awscdklibawsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   	ReturnResponse: jsii.Boolean(true),
//   })
//
type WebSocketRouteOptions struct {
	// The integration to be configured on this route.
	Integration WebSocketRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The authorize to this route.
	//
	// You can only set authorizer to a $connect route.
	// Default: - No Authorizer.
	//
	Authorizer IWebSocketRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Should the route send a response to the client.
	// Default: false.
	//
	ReturnResponse *bool `field:"optional" json:"returnResponse" yaml:"returnResponse"`
}

