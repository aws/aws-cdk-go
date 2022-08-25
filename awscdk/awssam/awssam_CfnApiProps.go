package awssam


// Properties for defining a `CfnApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   var gatewayResponses interface{}
//   var methodSettings interface{}
//   var models interface{}
//
//   cfnApiProps := &cfnApiProps{
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &authProperty{
//   		addDefaultAuthorizerToCorsPreflight: jsii.Boolean(false),
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	cors: jsii.String("cors"),
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	domain: &domainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: []*string{
//   			jsii.String("basePath"),
//   		},
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.String("truststoreVersion"),
//   		},
//   		ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	gatewayResponses: gatewayResponses,
//   	methodSettings: []interface{}{
//   		methodSettings,
//   	},
//   	minimumCompressionSize: jsii.Number(123),
//   	models: models,
//   	name: jsii.String("name"),
//   	openApiVersion: jsii.String("openApiVersion"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnApiProps struct {
	// `AWS::Serverless::Api.StageName`.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::Api.AccessLogSetting`.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::Api.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `AWS::Serverless::Api.BinaryMediaTypes`.
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// `AWS::Serverless::Api.CacheClusterEnabled`.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// `AWS::Serverless::Api.CacheClusterSize`.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// `AWS::Serverless::Api.CanarySetting`.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// `AWS::Serverless::Api.Cors`.
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// `AWS::Serverless::Api.DefinitionBody`.
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::Api.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::Api.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::Api.Domain`.
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// `AWS::Serverless::Api.EndpointConfiguration`.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `AWS::Serverless::Api.GatewayResponses`.
	GatewayResponses interface{} `field:"optional" json:"gatewayResponses" yaml:"gatewayResponses"`
	// `AWS::Serverless::Api.MethodSettings`.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// `AWS::Serverless::Api.MinimumCompressionSize`.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// `AWS::Serverless::Api.Models`.
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// `AWS::Serverless::Api.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion *string `field:"optional" json:"openApiVersion" yaml:"openApiVersion"`
	// `AWS::Serverless::Api.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// `AWS::Serverless::Api.Variables`.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

