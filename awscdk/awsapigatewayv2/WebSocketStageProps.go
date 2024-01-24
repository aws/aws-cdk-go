package awsapigatewayv2


// Properties to initialize an instance of `WebSocketStage`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectHandler function
//   var disconnectHandler function
//   var defaultHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
//   	ConnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("ConnectIntegration"), connectHandler),
//   	},
//   	DisconnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("DisconnectIntegration"), disconnectHandler),
//   	},
//   	DefaultRouteOptions: &WebSocketRouteOptions{
//   		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("DefaultIntegration"), defaultHandler),
//   	},
//   })
//
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//
type WebSocketStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Default: false.
	//
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Default: - no custom domain and api mapping configuration.
	//
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Throttle settings for the routes of this stage.
	// Default: - no throttling configuration.
	//
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The WebSocket API to which this stage is associated.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
}

