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
// Experimental.
type StageOptions struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Experimental.
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Experimental.
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Throttle settings for the routes of this stage.
	// Experimental.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
}

