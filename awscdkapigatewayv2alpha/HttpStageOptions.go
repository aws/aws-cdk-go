// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


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
// Experimental.
type HttpStageOptions struct {
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
}

