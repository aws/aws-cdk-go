package awsapigatewayv2


// Properties to initialize an instance of `HttpStage`.
//
// Example:
//   var api httpApi
//
//
//   apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
//   	HttpApi: api,
//   	StageName: jsii.String("beta"),
//   	Description: jsii.String("My Stage"),
//   })
//
type HttpStageProps struct {
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
	// The HTTP API to which this stage is associated.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
}

