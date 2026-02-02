package previewawsopensearchserverlessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsopensearchserverless/previewawsopensearchserverlessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a security configuration for OpenSearch Serverless.
//
// For more information, see [SAML authentication for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityConfigPropsMixin(&CfnSecurityConfigMixinProps{
//   	Description: jsii.String("description"),
//   	IamFederationOptions: &IamFederationConfigOptionsProperty{
//   		GroupAttribute: jsii.String("groupAttribute"),
//   		UserAttribute: jsii.String("userAttribute"),
//   	},
//   	IamIdentityCenterOptions: &IamIdentityCenterConfigOptionsProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		ApplicationDescription: jsii.String("applicationDescription"),
//   		ApplicationName: jsii.String("applicationName"),
//   		GroupAttribute: jsii.String("groupAttribute"),
//   		InstanceArn: jsii.String("instanceArn"),
//   		UserAttribute: jsii.String("userAttribute"),
//   	},
//   	Name: jsii.String("name"),
//   	SamlOptions: &SamlConfigOptionsProperty{
//   		GroupAttribute: jsii.String("groupAttribute"),
//   		Metadata: jsii.String("metadata"),
//   		OpenSearchServerlessEntityId: jsii.String("openSearchServerlessEntityId"),
//   		SessionTimeout: jsii.Number(123),
//   		UserAttribute: jsii.String("userAttribute"),
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html
//
type CfnSecurityConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityConfigPropsMixin
type jsiiProxy_CfnSecurityConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityConfigPropsMixin) Props() *CfnSecurityConfigMixinProps {
	var returns *CfnSecurityConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpenSearchServerless::SecurityConfig`.
func NewCfnSecurityConfigPropsMixin(props *CfnSecurityConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnSecurityConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpenSearchServerless::SecurityConfig`.
func NewCfnSecurityConfigPropsMixin_Override(c CfnSecurityConfigPropsMixin, props *CfnSecurityConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnSecurityConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnSecurityConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnSecurityConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSecurityConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

