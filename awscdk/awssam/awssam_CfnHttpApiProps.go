package awssam


// Properties for defining a `CfnHttpApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//
//   cfnHttpApiProps := &cfnHttpApiProps{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &httpApiAuthProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	corsConfiguration: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domain: &httpApiDomainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.Boolean(false),
//   		},
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	stageName: jsii.String("stageName"),
//   	stageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnHttpApiProps struct {
	// `AWS::Serverless::HttpApi.AccessLogSetting`.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::HttpApi.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `AWS::Serverless::HttpApi.CorsConfiguration`.
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// `AWS::Serverless::HttpApi.DefaultRouteSettings`.
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// `AWS::Serverless::HttpApi.DefinitionBody`.
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::HttpApi.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::HttpApi.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::HttpApi.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// `AWS::Serverless::HttpApi.Domain`.
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// `AWS::Serverless::HttpApi.FailOnWarnings`.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// `AWS::Serverless::HttpApi.RouteSettings`.
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// `AWS::Serverless::HttpApi.StageName`.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::HttpApi.StageVariables`.
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// `AWS::Serverless::HttpApi.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

