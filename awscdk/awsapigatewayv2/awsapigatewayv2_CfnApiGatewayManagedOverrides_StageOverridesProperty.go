package awsapigatewayv2


// The `StageOverrides` property overrides the stage configuration for an API Gateway-managed stage.
//
// If you remove this property, API Gateway restores the default values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//
//   stageOverridesProperty := &stageOverridesProperty{
//   	accessLogSettings: &accessLogSettingsProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	autoDeploy: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	routeSettings: routeSettings,
//   	stageVariables: stageVariables,
//   }
//
type CfnApiGatewayManagedOverrides_StageOverridesProperty struct {
	// Settings for logging access in a stage.
	AccessLogSettings interface{} `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Specifies whether updates to an API automatically trigger a new deployment.
	//
	// The default value is `true` .
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The default route settings for the stage.
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// The description for the API stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Route settings for the stage.
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// A map that defines the stage variables for a `Stage` .
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
}

