package awsapigatewayv2


// Options required to create a new stage.
//
// Options that are common between HTTP and Websocket APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessLogSettings IAccessLogSettings
//   var domainNameRef IDomainNameRef
//
//   stageOptions := &StageOptions{
//   	AccessLogSettings: accessLogSettings,
//   	AutoDeploy: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DetailedMetricsEnabled: jsii.Boolean(false),
//   	DomainMapping: &DomainMappingOptions{
//   		DomainName: domainNameRef,
//
//   		// the properties below are optional
//   		MappingKey: jsii.String("mappingKey"),
//   	},
//   	StageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	Throttle: &ThrottleSettings{
//   		BurstLimit: jsii.Number(123),
//   		RateLimit: jsii.Number(123),
//   	},
//   }
//
type StageOptions struct {
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
}

