package mixinsawsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsappsync/mixinsawsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::ChannelNamespace` resource creates a channel namespace associated with an `Api` .
//
// The `ChannelNamespace` contains the definitions for code handlers for the `Api` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelNamespacePropsMixin := awscdkmixinspreview.Mixins.NewCfnChannelNamespacePropsMixin(&CfnChannelNamespaceMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	CodeHandlers: jsii.String("codeHandlers"),
//   	CodeS3Location: jsii.String("codeS3Location"),
//   	HandlerConfigs: &HandlerConfigsProperty{
//   		OnPublish: &HandlerConfigProperty{
//   			Behavior: jsii.String("behavior"),
//   			Integration: &IntegrationProperty{
//   				DataSourceName: jsii.String("dataSourceName"),
//   				LambdaConfig: &LambdaConfigProperty{
//   					InvokeType: jsii.String("invokeType"),
//   				},
//   			},
//   		},
//   		OnSubscribe: &HandlerConfigProperty{
//   			Behavior: jsii.String("behavior"),
//   			Integration: &IntegrationProperty{
//   				DataSourceName: jsii.String("dataSourceName"),
//   				LambdaConfig: &LambdaConfigProperty{
//   					InvokeType: jsii.String("invokeType"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PublishAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	SubscribeAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html
//
type CfnChannelNamespacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnChannelNamespaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelNamespacePropsMixin
type jsiiProxy_CfnChannelNamespacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnChannelNamespacePropsMixin) Props() *CfnChannelNamespaceMixinProps {
	var returns *CfnChannelNamespaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelNamespacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::ChannelNamespace`.
func NewCfnChannelNamespacePropsMixin(props *CfnChannelNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnChannelNamespacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelNamespacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelNamespacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnChannelNamespacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::ChannelNamespace`.
func NewCfnChannelNamespacePropsMixin_Override(c CfnChannelNamespacePropsMixin, props *CfnChannelNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnChannelNamespacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnChannelNamespacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelNamespacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnChannelNamespacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelNamespacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnChannelNamespacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelNamespacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnChannelNamespacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

