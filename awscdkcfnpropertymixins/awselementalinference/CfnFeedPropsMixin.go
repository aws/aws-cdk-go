package awselementalinference

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awselementalinference/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a feed that receives media for inference processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cropping interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnFeedPropsMixin := awscdkcfnpropertymixins.Aws_elementalinference.NewCfnFeedPropsMixin(&CfnFeedMixinProps{
//   	Name: jsii.String("name"),
//   	Outputs: []interface{}{
//   		&GetOutputProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			OutputConfig: &OutputConfigProperty{
//   				Clipping: &ClippingConfigProperty{
//   					CallbackMetadata: jsii.String("callbackMetadata"),
//   				},
//   				Cropping: cropping,
//   			},
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html
//
type CfnFeedPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFeedMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFeedPropsMixin
type jsiiProxy_CfnFeedPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFeedPropsMixin) Props() *CfnFeedMixinProps {
	var returns *CfnFeedMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeedPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElementalInference::Feed`.
func NewCfnFeedPropsMixin(props *CfnFeedMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFeedPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFeedPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFeedPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElementalInference::Feed`.
func NewCfnFeedPropsMixin_Override(c CfnFeedPropsMixin, props *CfnFeedMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFeedPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFeedPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeedPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeedPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFeedPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

