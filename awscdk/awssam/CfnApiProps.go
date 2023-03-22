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
//   cfnApiProps := &CfnApiProps{
//   	StageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	AccessLogSetting: &AccessLogSettingProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	Auth: &AuthProperty{
//   		AddDefaultAuthorizerToCorsPreflight: jsii.Boolean(false),
//   		Authorizers: authorizers,
//   		DefaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	BinaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	CacheClusterEnabled: jsii.Boolean(false),
//   	CacheClusterSize: jsii.String("cacheClusterSize"),
//   	CanarySetting: &CanarySettingProperty{
//   		DeploymentId: jsii.String("deploymentId"),
//   		PercentTraffic: jsii.Number(123),
//   		StageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		UseStageCache: jsii.Boolean(false),
//   	},
//   	Cors: jsii.String("cors"),
//   	DefinitionBody: definitionBody,
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Description: jsii.String("description"),
//   	DisableExecuteApiEndpoint: jsii.Boolean(false),
//   	Domain: &DomainConfigurationProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DomainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		BasePath: []*string{
//   			jsii.String("basePath"),
//   		},
//   		EndpointConfiguration: jsii.String("endpointConfiguration"),
//   		MutualTlsAuthentication: &MutualTlsAuthenticationProperty{
//   			TruststoreUri: jsii.String("truststoreUri"),
//   			TruststoreVersion: jsii.String("truststoreVersion"),
//   		},
//   		OwnershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   		Route53: &Route53ConfigurationProperty{
//   			DistributedDomainName: jsii.String("distributedDomainName"),
//   			EvaluateTargetHealth: jsii.Boolean(false),
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			HostedZoneName: jsii.String("hostedZoneName"),
//   			IpV6: jsii.Boolean(false),
//   		},
//   		SecurityPolicy: jsii.String("securityPolicy"),
//   	},
//   	EndpointConfiguration: jsii.String("endpointConfiguration"),
//   	GatewayResponses: gatewayResponses,
//   	MethodSettings: []interface{}{
//   		methodSettings,
//   	},
//   	MinimumCompressionSize: jsii.Number(123),
//   	Models: models,
//   	Name: jsii.String("name"),
//   	OpenApiVersion: jsii.String("openApiVersion"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TracingEnabled: jsii.Boolean(false),
//   	Variables: map[string]*string{
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
	// `AWS::Serverless::Api.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
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

