package awsapigatewayv2


// The options to create a new Stage for an HTTP API.
//
// Example:
//   var api httpApi
//   var dn domainName
//
//
//   api.AddStage(jsii.String("beta"), &HttpStageOptions{
//   	StageName: jsii.String("beta"),
//   	AutoDeploy: jsii.Boolean(true),
//   	// https://${dn.domainName}/bar goes to the beta stage
//   	DomainMapping: &DomainMappingOptions{
//   		DomainName: dn,
//   		MappingKey: jsii.String("bar"),
//   	},
//   })
//
type HttpStageOptions struct {
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
	//
	// See `StageName` class for more details.
	// Default: '$default' the default stage of the API. This stage will have the URL at the root of the API endpoint.
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

