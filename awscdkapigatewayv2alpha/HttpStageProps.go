package awscdkapigatewayv2alpha


// Properties to initialize an instance of `HttpStage`.
//
// Example:
//   var api httpApi
//
//
//   apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
//   	HttpApi: api,
//   	StageName: jsii.String("beta"),
//   })
//
// Deprecated.
type HttpStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Default: false.
	//
	// Deprecated.
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Default: - no custom domain and api mapping configuration.
	//
	// Deprecated.
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Throttle settings for the routes of this stage.
	// Default: - no throttling configuration.
	//
	// Deprecated.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of the stage.
	//
	// See `StageName` class for more details.
	// Default: '$default' the default stage of the API. This stage will have the URL at the root of the API endpoint.
	//
	// Deprecated.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// The HTTP API to which this stage is associated.
	// Deprecated.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
}

