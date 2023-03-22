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
//   cfnApiGatewayManagedOverridesProps := &CfnApiGatewayManagedOverridesProps{
//   	ApiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	Integration: &IntegrationOverridesProperty{
//   		Description: jsii.String("description"),
//   		IntegrationMethod: jsii.String("integrationMethod"),
//   		PayloadFormatVersion: jsii.String("payloadFormatVersion"),
//   		TimeoutInMillis: jsii.Number(123),
//   	},
//   	Route: &RouteOverridesProperty{
//   		AuthorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		AuthorizationType: jsii.String("authorizationType"),
//   		AuthorizerId: jsii.String("authorizerId"),
//   		OperationName: jsii.String("operationName"),
//   		Target: jsii.String("target"),
//   	},
//   	Stage: &StageOverridesProperty{
//   		AccessLogSettings: &AccessLogSettingsProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   			Format: jsii.String("format"),
//   		},
//   		AutoDeploy: jsii.Boolean(false),
//   		DefaultRouteSettings: &RouteSettingsProperty{
//   			DataTraceEnabled: jsii.Boolean(false),
//   			DetailedMetricsEnabled: jsii.Boolean(false),
//   			LoggingLevel: jsii.String("loggingLevel"),
//   			ThrottlingBurstLimit: jsii.Number(123),
//   			ThrottlingRateLimit: jsii.Number(123),
//   		},
//   		Description: jsii.String("description"),
//   		RouteSettings: routeSettings,
//   		StageVariables: stageVariables,
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

