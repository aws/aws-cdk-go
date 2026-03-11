package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a version for the specified customer-managed view within the specified instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnViewVersionPropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnViewVersionPropsMixin(&CfnViewVersionMixinProps{
//   	VersionDescription: jsii.String("versionDescription"),
//   	ViewArn: jsii.String("viewArn"),
//   	ViewContentSha256: jsii.String("viewContentSha256"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-viewversion.html
//
type CfnViewVersionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnViewVersionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnViewVersionPropsMixin
type jsiiProxy_CfnViewVersionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnViewVersionPropsMixin) Props() *CfnViewVersionMixinProps {
	var returns *CfnViewVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnViewVersionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::ViewVersion`.
func NewCfnViewVersionPropsMixin(props *CfnViewVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnViewVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnViewVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnViewVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnViewVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::ViewVersion`.
func NewCfnViewVersionPropsMixin_Override(c CfnViewVersionPropsMixin, props *CfnViewVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnViewVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnViewVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnViewVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnViewVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnViewVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnViewVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnViewVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnViewVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

