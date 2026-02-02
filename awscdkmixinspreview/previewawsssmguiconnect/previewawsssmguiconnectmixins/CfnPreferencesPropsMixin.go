package previewawsssmguiconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssmguiconnect/previewawsssmguiconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify new or changed connection recording preferences for your AWS Systems Manager GUI Connect connections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPreferencesPropsMixin := awscdkmixinspreview.Mixins.NewCfnPreferencesPropsMixin(&CfnPreferencesMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html
//
type CfnPreferencesPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPreferencesMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPreferencesPropsMixin
type jsiiProxy_CfnPreferencesPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnPreferencesPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMGuiConnect::Preferences`.
func NewCfnPreferencesPropsMixin(props *CfnPreferencesMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPreferencesPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPreferencesPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPreferencesPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMGuiConnect::Preferences`.
func NewCfnPreferencesPropsMixin_Override(c CfnPreferencesPropsMixin, props *CfnPreferencesMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPreferencesPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPreferencesPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin",
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

