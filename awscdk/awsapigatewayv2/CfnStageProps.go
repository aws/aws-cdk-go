package awsapigatewayv2


// Properties for defining a `CfnStage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//   var tags interface{}
//
//   cfnStageProps := &CfnStageProps{
//   	ApiId: jsii.String("apiId"),
//   	StageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	AccessLogSettings: &AccessLogSettingsProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	AccessPolicyId: jsii.String("accessPolicyId"),
//   	AutoDeploy: jsii.Boolean(false),
//   	ClientCertificateId: jsii.String("clientCertificateId"),
//   	DefaultRouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	DeploymentId: jsii.String("deploymentId"),
//   	Description: jsii.String("description"),
//   	RouteSettings: routeSettings,
//   	StageVariables: stageVariables,
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html
//
type CfnStageProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The stage name.
	//
	// Stage names can contain only alphanumeric characters, hyphens, and underscores, or be `$default` . Maximum length is 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-stagename
	//
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// Settings for logging access in this stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-accesslogsettings
	//
	AccessLogSettings interface{} `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// This parameter is not currently supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-accesspolicyid
	//
	AccessPolicyId *string `field:"optional" json:"accessPolicyId" yaml:"accessPolicyId"`
	// Specifies whether updates to an API automatically trigger a new deployment.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-autodeploy
	//
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The identifier of a client certificate for a `Stage` .
	//
	// Supported only for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-clientcertificateid
	//
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The default route settings for the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-defaultroutesettings
	//
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// The deployment identifier for the API stage.
	//
	// Can't be updated if `autoDeploy` is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-deploymentid
	//
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The description for the API stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Route settings for the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-routesettings
	//
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// A map that defines the stage variables for a `Stage` .
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-stagevariables
	//
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

