package mixinsawsmediastore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmediastore/mixinsawsmediastore/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::MediaStore::Container resource specifies a storage container to hold objects.
//
// A container is similar to a bucket in Amazon S3.
//
// When you create a container using CloudFormation , the template manages data for five API actions: creating a container, setting access logging, updating the default container policy, adding a cross-origin resource sharing (CORS) policy, and adding an object lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContainerPropsMixin := awscdkmixinspreview.Mixins.NewCfnContainerPropsMixin(&CfnContainerMixinProps{
//   	AccessLoggingEnabled: jsii.Boolean(false),
//   	ContainerName: jsii.String("containerName"),
//   	CorsPolicy: []interface{}{
//   		&CorsRuleProperty{
//   			AllowedHeaders: []*string{
//   				jsii.String("allowedHeaders"),
//   			},
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			AllowedOrigins: []*string{
//   				jsii.String("allowedOrigins"),
//   			},
//   			ExposeHeaders: []*string{
//   				jsii.String("exposeHeaders"),
//   			},
//   			MaxAgeSeconds: jsii.Number(123),
//   		},
//   	},
//   	LifecyclePolicy: jsii.String("lifecyclePolicy"),
//   	MetricPolicy: &MetricPolicyProperty{
//   		ContainerLevelMetrics: jsii.String("containerLevelMetrics"),
//   		MetricPolicyRules: []interface{}{
//   			&MetricPolicyRuleProperty{
//   				ObjectGroup: jsii.String("objectGroup"),
//   				ObjectGroupName: jsii.String("objectGroupName"),
//   			},
//   		},
//   	},
//   	Policy: jsii.String("policy"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediastore-container.html
//
type CfnContainerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnContainerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContainerPropsMixin
type jsiiProxy_CfnContainerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnContainerPropsMixin) Props() *CfnContainerMixinProps {
	var returns *CfnContainerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaStore::Container`.
func NewCfnContainerPropsMixin(props *CfnContainerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnContainerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContainerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaStore::Container`.
func NewCfnContainerPropsMixin_Override(c CfnContainerPropsMixin, props *CfnContainerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnContainerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnContainerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

