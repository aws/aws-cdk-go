package awssmsvoice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A destination phone number that has been registered for verification with AWS End User Messaging SMS.
//
// A newly created number is in PENDING status; it becomes VERIFIED only after the recipient supplies the one-time code out of band.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVerifiedDestinationNumberPropsMixin := awscdkcfnpropertymixins.Aws_smsvoice.NewCfnVerifiedDestinationNumberPropsMixin(&CfnVerifiedDestinationNumberMixinProps{
//   	DestinationPhoneNumber: jsii.String("destinationPhoneNumber"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-verifieddestinationnumber.html
//
type CfnVerifiedDestinationNumberPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVerifiedDestinationNumberMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVerifiedDestinationNumberPropsMixin
type jsiiProxy_CfnVerifiedDestinationNumberPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVerifiedDestinationNumberPropsMixin) Props() *CfnVerifiedDestinationNumberMixinProps {
	var returns *CfnVerifiedDestinationNumberMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedDestinationNumberPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::VerifiedDestinationNumber`.
func NewCfnVerifiedDestinationNumberPropsMixin(props *CfnVerifiedDestinationNumberMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVerifiedDestinationNumberPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVerifiedDestinationNumberPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVerifiedDestinationNumberPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnVerifiedDestinationNumberPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::VerifiedDestinationNumber`.
func NewCfnVerifiedDestinationNumberPropsMixin_Override(c CfnVerifiedDestinationNumberPropsMixin, props *CfnVerifiedDestinationNumberMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnVerifiedDestinationNumberPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVerifiedDestinationNumberPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedDestinationNumberPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnVerifiedDestinationNumberPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVerifiedDestinationNumberPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnVerifiedDestinationNumberPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVerifiedDestinationNumberPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVerifiedDestinationNumberPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

