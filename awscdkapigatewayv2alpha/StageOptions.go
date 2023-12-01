package awscdkapigatewayv2alpha


// Options required to create a new stage.
//
// Options that are common between HTTP and Websocket APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var domainName domainName
//
//   stageOptions := &StageOptions{
//   	AutoDeploy: jsii.Boolean(false),
//   	DomainMapping: &DomainMappingOptions{
//   		DomainName: domainName,
//
//   		// the properties below are optional
//   		MappingKey: jsii.String("mappingKey"),
//   	},
//   	Throttle: &ThrottleSettings{
//   		BurstLimit: jsii.Number(123),
//   		RateLimit: jsii.Number(123),
//   	},
//   }
//
// Deprecated.
type StageOptions struct {
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
}

