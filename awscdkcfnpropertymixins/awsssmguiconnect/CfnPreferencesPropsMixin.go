package awsssmguiconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsssmguiconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify new or changed connection recording preferences for your AWS Systems Manager GUI Connect connections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPreferencesPropsMixin := awscdkcfnpropertymixins.Aws_ssmguiconnect.NewCfnPreferencesPropsMixin(&CfnPreferencesMixinProps{
//   	ConnectionRecordingPreferences: &ConnectionRecordingPreferencesProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		RecordingDestinations: &RecordingDestinationsProperty{
//   			S3Buckets: []interface{}{
//   				&S3BucketProperty{
//   					BucketName: jsii.String("bucketName"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html
//
type CfnPreferencesPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPreferencesMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPreferencesPropsMixin
type jsiiProxy_CfnPreferencesPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPreferencesPropsMixin) Props() *CfnPreferencesMixinProps {
	var returns *CfnPreferencesMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPreferencesPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMGuiConnect::Preferences`.
func NewCfnPreferencesPropsMixin(props *CfnPreferencesMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPreferencesPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPreferencesPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPreferencesPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMGuiConnect::Preferences`.
func NewCfnPreferencesPropsMixin_Override(c CfnPreferencesPropsMixin, props *CfnPreferencesMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPreferencesPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPreferencesPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPreferencesPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPreferencesPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPreferencesPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

