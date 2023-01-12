// The CDK Construct Library for AWS::APIGatewayv2
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
//   stageOptions := &stageOptions{
//   	autoDeploy: jsii.Boolean(false),
//   	domainMapping: &domainMappingOptions{
//   		domainName: domainName,
//
//   		// the properties below are optional
//   		mappingKey: jsii.String("mappingKey"),
//   	},
//   	throttle: &throttleSettings{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
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

