package awsapigatewayv2


// Properties to initialize an instance of `WebSocketStage`.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi HttpApi
//   var webSocketApi WebSocketApi
//   var logGroup LogGroup
//
//
//   apigwv2.NewHttpStage(this, jsii.String("HttpStage"), &HttpStageProps{
//   	HttpApi: HttpApi,
//   	AccessLogSettings: map[string]IAccessLogDestination{
//   		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
//   	},
//   })
//
//   apigwv2.NewWebSocketStage(this, jsii.String("WebSocketStage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AccessLogSettings: map[string]IAccessLogDestination{
//   		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
//   	},
//   })
//
type WebSocketStageProps struct {
	// Settings for access logging.
	// Default: - No access logging.
	//
	AccessLogSettings IAccessLogSettings `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Whether updates to an API automatically trigger a new deployment.
	// Default: false.
	//
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The description for the API stage.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether detailed metrics are enabled.
	// Default: false.
	//
	DetailedMetricsEnabled *bool `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// The options for custom domain and api mapping.
	// Default: - no custom domain and api mapping configuration.
	//
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Stage variables for the stage. These are key-value pairs that you can define and use in your API routes.
	//
	// The allowed characters for variable names and the required pattern for variable values are specified here: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-apigateway-stage.html#cfn-apigateway-stage-variables
	// Default: - No stage variables.
	//
	StageVariables *map[string]*string `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// Throttle settings for the routes of this stage.
	// Default: - no throttling configuration.
	//
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The WebSocket API to which this stage is associated.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
}

