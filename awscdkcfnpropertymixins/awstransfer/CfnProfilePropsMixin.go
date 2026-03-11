package awstransfer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates the local or partner profile to use for AS2 transfers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProfilePropsMixin := awscdkcfnpropertymixins.Aws_transfer.NewCfnProfilePropsMixin(&CfnProfileMixinProps{
//   	As2Id: jsii.String("as2Id"),
//   	CertificateIds: []*string{
//   		jsii.String("certificateIds"),
//   	},
//   	ProfileType: jsii.String("profileType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-profile.html
//
type CfnProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProfilePropsMixin
type jsiiProxy_CfnProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProfilePropsMixin) Props() *CfnProfileMixinProps {
	var returns *CfnProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transfer::Profile`.
func NewCfnProfilePropsMixin(props *CfnProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transfer.CfnProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::Profile`.
func NewCfnProfilePropsMixin_Override(c CfnProfilePropsMixin, props *CfnProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transfer.CfnProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_transfer.CfnProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_transfer.CfnProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

