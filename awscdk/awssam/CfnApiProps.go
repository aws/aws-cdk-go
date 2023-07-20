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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html
//
type CfnApiProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-stagename
	//
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-accesslogsetting
	//
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-binarymediatypes
	//
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-cacheclusterenabled
	//
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-cacheclustersize
	//
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-canarysetting
	//
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-cors
	//
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-definitionbody
	//
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-definitionuri
	//
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-disableexecuteapiendpoint
	//
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-domain
	//
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-endpointconfiguration
	//
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-gatewayresponses
	//
	GatewayResponses interface{} `field:"optional" json:"gatewayResponses" yaml:"gatewayResponses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-methodsettings
	//
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-minimumcompressionsize
	//
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-models
	//
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-openapiversion
	//
	OpenApiVersion *string `field:"optional" json:"openApiVersion" yaml:"openApiVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-tracingenabled
	//
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-api.html#cfn-serverless-api-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

