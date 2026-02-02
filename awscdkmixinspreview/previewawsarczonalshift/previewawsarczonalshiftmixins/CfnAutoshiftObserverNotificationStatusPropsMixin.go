package previewawsarczonalshiftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsarczonalshift/previewawsarczonalshiftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::ARCZonalShift::AutoshiftObserverNotificationStatus Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutoshiftObserverNotificationStatusPropsMixin := awscdkmixinspreview.Mixins.NewCfnAutoshiftObserverNotificationStatusPropsMixin(&CfnAutoshiftObserverNotificationStatusMixinProps{
//   	Status: jsii.String("status"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-autoshiftobservernotificationstatus.html
//
type CfnAutoshiftObserverNotificationStatusPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAutoshiftObserverNotificationStatusMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutoshiftObserverNotificationStatusPropsMixin
type jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin) Props() *CfnAutoshiftObserverNotificationStatusMixinProps {
	var returns *CfnAutoshiftObserverNotificationStatusMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ARCZonalShift::AutoshiftObserverNotificationStatus`.
func NewCfnAutoshiftObserverNotificationStatusPropsMixin(props *CfnAutoshiftObserverNotificationStatusMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutoshiftObserverNotificationStatusPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutoshiftObserverNotificationStatusPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ARCZonalShift::AutoshiftObserverNotificationStatus`.
func NewCfnAutoshiftObserverNotificationStatusPropsMixin_Override(c CfnAutoshiftObserverNotificationStatusPropsMixin, props *CfnAutoshiftObserverNotificationStatusMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAutoshiftObserverNotificationStatusPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoshiftObserverNotificationStatusPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutoshiftObserverNotificationStatusPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnAutoshiftObserverNotificationStatusPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAutoshiftObserverNotificationStatusPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

