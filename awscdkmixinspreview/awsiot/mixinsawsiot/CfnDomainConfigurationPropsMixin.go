package mixinsawsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiot/mixinsawsiot/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a domain configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnDomainConfigurationPropsMixin(&CfnDomainConfigurationMixinProps{
//   	ApplicationProtocol: jsii.String("applicationProtocol"),
//   	AuthenticationType: jsii.String("authenticationType"),
//   	AuthorizerConfig: &AuthorizerConfigProperty{
//   		AllowAuthorizerOverride: jsii.Boolean(false),
//   		DefaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   	},
//   	ClientCertificateConfig: &ClientCertificateConfigProperty{
//   		ClientCertificateCallbackArn: jsii.String("clientCertificateCallbackArn"),
//   	},
//   	DomainConfigurationName: jsii.String("domainConfigurationName"),
//   	DomainConfigurationStatus: jsii.String("domainConfigurationStatus"),
//   	DomainName: jsii.String("domainName"),
//   	ServerCertificateArns: []*string{
//   		jsii.String("serverCertificateArns"),
//   	},
//   	ServerCertificateConfig: &ServerCertificateConfigProperty{
//   		EnableOcspCheck: jsii.Boolean(false),
//   		OcspAuthorizedResponderArn: jsii.String("ocspAuthorizedResponderArn"),
//   		OcspLambdaArn: jsii.String("ocspLambdaArn"),
//   	},
//   	ServiceType: jsii.String("serviceType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsConfig: &TlsConfigProperty{
//   		SecurityPolicy: jsii.String("securityPolicy"),
//   	},
//   	ValidationCertificateArn: jsii.String("validationCertificateArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html
//
type CfnDomainConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDomainConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainConfigurationPropsMixin
type jsiiProxy_CfnDomainConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDomainConfigurationPropsMixin) Props() *CfnDomainConfigurationMixinProps {
	var returns *CfnDomainConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::DomainConfiguration`.
func NewCfnDomainConfigurationPropsMixin(props *CfnDomainConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDomainConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnDomainConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::DomainConfiguration`.
func NewCfnDomainConfigurationPropsMixin_Override(c CfnDomainConfigurationPropsMixin, props *CfnDomainConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnDomainConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDomainConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnDomainConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnDomainConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDomainConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

