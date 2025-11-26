package previewawsopsworksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsopsworks/previewawsopsworksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppPropsMixin(&CfnAppMixinProps{
//   	AppSource: &SourceProperty{
//   		Password: jsii.String("password"),
//   		Revision: jsii.String("revision"),
//   		SshKey: jsii.String("sshKey"),
//   		Type: jsii.String("type"),
//   		Url: jsii.String("url"),
//   		Username: jsii.String("username"),
//   	},
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	DataSources: []interface{}{
//   		&DataSourceProperty{
//   			Arn: jsii.String("arn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	EnableSsl: jsii.Boolean(false),
//   	Environment: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Key: jsii.String("key"),
//   			Secure: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Shortname: jsii.String("shortname"),
//   	SslConfiguration: &SslConfigurationProperty{
//   		Certificate: jsii.String("certificate"),
//   		Chain: jsii.String("chain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   	StackId: jsii.String("stackId"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
//
type CfnAppPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppPropsMixin
type jsiiProxy_CfnAppPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppPropsMixin) Props() *CfnAppMixinProps {
	var returns *CfnAppMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpsWorks::App`.
func NewCfnAppPropsMixin(props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpsWorks::App`.
func NewCfnAppPropsMixin_Override(c CfnAppPropsMixin, props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnAppPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnAppPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAppPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

