package awscdkapigatewayv2alpha


// Options used to add route to the API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketRouteOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration WebSocketRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The authorize to this route.
	//
	// You can only set authorizer to a $connect route.
	// Experimental.
	Authorizer IWebSocketRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Should the route send a response to the client.
	// Experimental.
	ReturnResponse *bool `field:"optional" json:"returnResponse" yaml:"returnResponse"`
}

