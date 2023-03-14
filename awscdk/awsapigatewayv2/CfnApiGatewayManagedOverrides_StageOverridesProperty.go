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
//   stageOverridesProperty := &StageOverridesProperty{
//   	AccessLogSettings: &AccessLogSettingsProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	AutoDeploy: jsii.Boolean(false),
//   	DefaultRouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	RouteSettings: routeSettings,
//   	StageVariables: stageVariables,
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

