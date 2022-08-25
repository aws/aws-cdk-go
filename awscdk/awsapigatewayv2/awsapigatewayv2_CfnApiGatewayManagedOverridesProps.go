package awsapigatewayv2


// Properties for defining a `CfnApiGatewayManagedOverrides`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//
//   cfnApiGatewayManagedOverridesProps := &cfnApiGatewayManagedOverridesProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	integration: &integrationOverridesProperty{
//   		description: jsii.String("description"),
//   		integrationMethod: jsii.String("integrationMethod"),
//   		payloadFormatVersion: jsii.String("payloadFormatVersion"),
//   		timeoutInMillis: jsii.Number(123),
//   	},
//   	route: &routeOverridesProperty{
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: jsii.String("authorizationType"),
//   		authorizerId: jsii.String("authorizerId"),
//   		operationName: jsii.String("operationName"),
//   		target: jsii.String("target"),
//   	},
//   	stage: &stageOverridesProperty{
//   		accessLogSettings: &accessLogSettingsProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			format: jsii.String("format"),
//   		},
//   		autoDeploy: jsii.Boolean(false),
//   		defaultRouteSettings: &routeSettingsProperty{
//   			dataTraceEnabled: jsii.Boolean(false),
//   			detailedMetricsEnabled: jsii.Boolean(false),
//   			loggingLevel: jsii.String("loggingLevel"),
//   			throttlingBurstLimit: jsii.Number(123),
//   			throttlingRateLimit: jsii.Number(123),
//   		},
//   		description: jsii.String("description"),
//   		routeSettings: routeSettings,
//   		stageVariables: stageVariables,
//   	},
//   }
//
type CfnApiGatewayManagedOverridesProps struct {
	// The ID of the API for which to override the configuration of API Gateway-managed resources.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Overrides the integration configuration for an API Gateway-managed integration.
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// Overrides the route configuration for an API Gateway-managed route.
	Route interface{} `field:"optional" json:"route" yaml:"route"`
	// Overrides the stage configuration for an API Gateway-managed stage.
	Stage interface{} `field:"optional" json:"stage" yaml:"stage"`
}

