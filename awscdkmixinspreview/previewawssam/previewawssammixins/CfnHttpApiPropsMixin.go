package previewawssammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssam/previewawssammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesshttpapi.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//
//   cfnHttpApiPropsMixin := awscdkmixinspreview.Mixins.NewCfnHttpApiPropsMixin(&CfnHttpApiMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-httpapi.html
//
type CfnHttpApiPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHttpApiMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHttpApiPropsMixin
type jsiiProxy_CfnHttpApiPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHttpApiPropsMixin) Props() *CfnHttpApiMixinProps {
	var returns *CfnHttpApiMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApiPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Serverless::HttpApi`.
func NewCfnHttpApiPropsMixin(props *CfnHttpApiMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHttpApiPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHttpApiPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHttpApiPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Serverless::HttpApi`.
func NewCfnHttpApiPropsMixin_Override(c CfnHttpApiPropsMixin, props *CfnHttpApiMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHttpApiPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHttpApiPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHttpApiPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHttpApiPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApiPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

