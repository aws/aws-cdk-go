package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q Business web experience.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebExperiencePropsMixin := awscdkmixinspreview.Mixins.NewCfnWebExperiencePropsMixin(&CfnWebExperienceMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	BrowserExtensionConfiguration: &BrowserExtensionConfigurationProperty{
//   		EnabledBrowserExtensions: []*string{
//   			jsii.String("enabledBrowserExtensions"),
//   		},
//   	},
//   	CustomizationConfiguration: &CustomizationConfigurationProperty{
//   		CustomCssUrl: jsii.String("customCssUrl"),
//   		FaviconUrl: jsii.String("faviconUrl"),
//   		FontUrl: jsii.String("fontUrl"),
//   		LogoUrl: jsii.String("logoUrl"),
//   	},
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		OpenIdConnectConfiguration: &OpenIDConnectProviderConfigurationProperty{
//   			SecretsArn: jsii.String("secretsArn"),
//   			SecretsRole: jsii.String("secretsRole"),
//   		},
//   		SamlConfiguration: &SamlProviderConfigurationProperty{
//   			AuthenticationUrl: jsii.String("authenticationUrl"),
//   		},
//   	},
//   	Origins: []*string{
//   		jsii.String("origins"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SamplePromptsControlMode: jsii.String("samplePromptsControlMode"),
//   	Subtitle: jsii.String("subtitle"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	WelcomeMessage: jsii.String("welcomeMessage"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html
//
type CfnWebExperiencePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWebExperienceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWebExperiencePropsMixin
type jsiiProxy_CfnWebExperiencePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWebExperiencePropsMixin) Props() *CfnWebExperienceMixinProps {
	var returns *CfnWebExperienceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebExperiencePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::WebExperience`.
func NewCfnWebExperiencePropsMixin(props *CfnWebExperienceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWebExperiencePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWebExperiencePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWebExperiencePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::WebExperience`.
func NewCfnWebExperiencePropsMixin_Override(c CfnWebExperiencePropsMixin, props *CfnWebExperienceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWebExperiencePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWebExperiencePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebExperiencePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebExperiencePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWebExperiencePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

