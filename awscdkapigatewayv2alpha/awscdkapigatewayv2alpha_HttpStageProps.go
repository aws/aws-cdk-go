// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Properties to initialize an instance of `HttpStage`.
//
// Example:
//   var api httpApi
//
//
//   apigwv2.NewHttpStage(this, jsii.String("Stage"), &httpStageProps{
//   	httpApi: api,
//   	stageName: jsii.String("beta"),
//   })
//
// Experimental.
type HttpStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Throttle settings for the routes of this stage.
	// Experimental.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of the stage.
	//
	// See `StageName` class for more details.
	// Experimental.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// The HTTP API to which this stage is associated.
	// Experimental.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
}

