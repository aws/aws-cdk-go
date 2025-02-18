package awsapigatewayv2


// Options required to create a new stage.
//
// Options that are common between HTTP and Websocket APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainName domainName
//
//   stageOptions := &StageOptions{
//   	AutoDeploy: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DetailedMetricsEnabled: jsii.Boolean(false),
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
type StageOptions struct {
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
}

