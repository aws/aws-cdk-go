package previewawssammixins


// Properties for CfnHttpApiPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//
//   cfnHttpApiMixinProps := &CfnHttpApiMixinProps{
//   	AccessLogSetting: &AccessLogSettingProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	Auth: &HttpApiAuthProperty{
//   		Authorizers: authorizers,
//   		DefaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	CorsConfiguration: jsii.Boolean(false),
//   	DefaultRouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	DefinitionBody: definitionBody,
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Description: jsii.String("description"),
//   	DisableExecuteApiEndpoint: jsii.Boolean(false),
//   	Domain: &HttpApiDomainConfigurationProperty{
//   		BasePath: jsii.String("basePath"),
//   		CertificateArn: jsii.String("certificateArn"),
//   		DomainName: jsii.String("domainName"),
//   		EndpointConfiguration: jsii.String("endpointConfiguration"),
//   		MutualTlsAuthentication: &MutualTlsAuthenticationProperty{
//   			TruststoreUri: jsii.String("truststoreUri"),
//   			TruststoreVersion: jsii.Boolean(false),
//   		},
//   		Route53: &Route53ConfigurationProperty{
//   			DistributedDomainName: jsii.String("distributedDomainName"),
//   			EvaluateTargetHealth: jsii.Boolean(false),
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			HostedZoneName: jsii.String("hostedZoneName"),
//   			IpV6: jsii.Boolean(false),
//   		},
//   		SecurityPolicy: jsii.String("securityPolicy"),
//   	},
//   	FailOnWarnings: jsii.Boolean(false),
//   	RouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	StageName: jsii.String("stageName"),
//   	StageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html
//
type CfnHttpApiMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-accesslogsetting
	//
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-corsconfiguration
	//
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-defaultroutesettings
	//
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-definitionbody
	//
	DefinitionBody interface{} `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-definitionuri
	//
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-disableexecuteapiendpoint
	//
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-domain
	//
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-failonwarnings
	//
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-routesettings
	//
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-stagename
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-stagevariables
	//
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html#cfn-serverless-httpapi-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

