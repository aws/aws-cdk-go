package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::Standard` resource specifies the enablement of a security standard.
//
// The standard is identified by the `StandardsArn` property. To view a list of Security Hub standards and their Amazon Resource Names (ARNs), use the [`DescribeStandards`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
//
// You must create a separate `AWS::SecurityHub::Standard` resource for each standard that you want to enable.
//
// For more information about Security Hub standards, see [Security Hub standards reference](https://docs.aws.amazon.com/securityhub/latest/userguide/standards-reference.html) in the *Security Hub User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStandardPropsMixin := awscdkmixinspreview.Mixins.NewCfnStandardPropsMixin(&CfnStandardMixinProps{
//   	DisabledStandardsControls: []interface{}{
//   		&StandardsControlProperty{
//   			Reason: jsii.String("reason"),
//   			StandardsControlArn: jsii.String("standardsControlArn"),
//   		},
//   	},
//   	StandardsArn: jsii.String("standardsArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html
//
type CfnStandardPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStandardMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStandardPropsMixin
type jsiiProxy_CfnStandardPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStandardPropsMixin) Props() *CfnStandardMixinProps {
	var returns *CfnStandardMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStandardPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::Standard`.
func NewCfnStandardPropsMixin(props *CfnStandardMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStandardPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStandardPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStandardPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnStandardPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::Standard`.
func NewCfnStandardPropsMixin_Override(c CfnStandardPropsMixin, props *CfnStandardMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnStandardPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStandardPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStandardPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnStandardPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStandardPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnStandardPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStandardPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStandardPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

