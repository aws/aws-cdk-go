package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigateway/previewawsapigatewaymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::Deployment` resource deploys an API Gateway `RestApi` resource to a stage so that clients can call the API over the internet.
//
// The stage acts as an environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeploymentPropsMixin(&CfnDeploymentMixinProps{
//   	DeploymentCanarySettings: &DeploymentCanarySettingsProperty{
//   		PercentTraffic: jsii.Number(123),
//   		StageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		UseStageCache: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	RestApiId: jsii.String("restApiId"),
//   	StageDescription: &StageDescriptionProperty{
//   		AccessLogSetting: &AccessLogSettingProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   			Format: jsii.String("format"),
//   		},
//   		CacheClusterEnabled: jsii.Boolean(false),
//   		CacheClusterSize: jsii.String("cacheClusterSize"),
//   		CacheDataEncrypted: jsii.Boolean(false),
//   		CacheTtlInSeconds: jsii.Number(123),
//   		CachingEnabled: jsii.Boolean(false),
//   		CanarySetting: &CanarySettingProperty{
//   			PercentTraffic: jsii.Number(123),
//   			StageVariableOverrides: map[string]*string{
//   				"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   			},
//   			UseStageCache: jsii.Boolean(false),
//   		},
//   		ClientCertificateId: jsii.String("clientCertificateId"),
//   		DataTraceEnabled: jsii.Boolean(false),
//   		Description: jsii.String("description"),
//   		DocumentationVersion: jsii.String("documentationVersion"),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		MethodSettings: []interface{}{
//   			&MethodSettingProperty{
//   				CacheDataEncrypted: jsii.Boolean(false),
//   				CacheTtlInSeconds: jsii.Number(123),
//   				CachingEnabled: jsii.Boolean(false),
//   				DataTraceEnabled: jsii.Boolean(false),
//   				HttpMethod: jsii.String("httpMethod"),
//   				LoggingLevel: jsii.String("loggingLevel"),
//   				MetricsEnabled: jsii.Boolean(false),
//   				ResourcePath: jsii.String("resourcePath"),
//   				ThrottlingBurstLimit: jsii.Number(123),
//   				ThrottlingRateLimit: jsii.Number(123),
//   			},
//   		},
//   		MetricsEnabled: jsii.Boolean(false),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   		TracingEnabled: jsii.Boolean(false),
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
//
type CfnDeploymentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeploymentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeploymentPropsMixin
type jsiiProxy_CfnDeploymentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Props() *CfnDeploymentMixinProps {
	var returns *CfnDeploymentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGateway::Deployment`.
func NewCfnDeploymentPropsMixin(props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeploymentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeploymentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::Deployment`.
func NewCfnDeploymentPropsMixin_Override(c CfnDeploymentPropsMixin, props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeploymentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnDeploymentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnDeploymentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDeploymentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

