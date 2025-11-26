package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Information about an Amazon Q Business plugin and its configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var noAuthConfiguration interface{}
//
//   cfnPluginPropsMixin := awscdkmixinspreview.Mixins.NewCfnPluginPropsMixin(&CfnPluginMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	AuthConfiguration: &PluginAuthConfigurationProperty{
//   		BasicAuthConfiguration: &BasicAuthConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		NoAuthConfiguration: noAuthConfiguration,
//   		OAuth2ClientCredentialConfiguration: &OAuth2ClientCredentialConfigurationProperty{
//   			AuthorizationUrl: jsii.String("authorizationUrl"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   		},
//   	},
//   	CustomPluginConfiguration: &CustomPluginConfigurationProperty{
//   		ApiSchema: &APISchemaProperty{
//   			Payload: jsii.String("payload"),
//   			S3: &S3Property{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		ApiSchemaType: jsii.String("apiSchemaType"),
//   		Description: jsii.String("description"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	ServerUrl: jsii.String("serverUrl"),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html
//
type CfnPluginPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPluginMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPluginPropsMixin
type jsiiProxy_CfnPluginPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPluginPropsMixin) Props() *CfnPluginMixinProps {
	var returns *CfnPluginMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPluginPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::Plugin`.
func NewCfnPluginPropsMixin(props *CfnPluginMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPluginPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPluginPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPluginPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::Plugin`.
func NewCfnPluginPropsMixin_Override(c CfnPluginPropsMixin, props *CfnPluginMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPluginPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPluginPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPluginPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPluginPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPluginPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

