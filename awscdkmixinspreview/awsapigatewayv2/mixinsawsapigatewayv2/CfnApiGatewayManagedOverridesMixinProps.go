package mixinsawsapigatewayv2


// Properties for CfnApiGatewayManagedOverridesPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//
//   cfnApiGatewayManagedOverridesMixinProps := &CfnApiGatewayManagedOverridesMixinProps{
//   	ApiId: jsii.String("apiId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html
//
type CfnApiGatewayManagedOverridesMixinProps struct {
	// The ID of the API for which to override the configuration of API Gateway-managed resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// Overrides the integration configuration for an API Gateway-managed integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-integration
	//
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// Overrides the route configuration for an API Gateway-managed route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-route
	//
	Route interface{} `field:"optional" json:"route" yaml:"route"`
	// Overrides the stage configuration for an API Gateway-managed stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stage
	//
	Stage interface{} `field:"optional" json:"stage" yaml:"stage"`
}

