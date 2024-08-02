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
	// Whether updates to an API automatically trigger a new deployment.
	// Default: false.
	//
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The description for the API stage.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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
}

