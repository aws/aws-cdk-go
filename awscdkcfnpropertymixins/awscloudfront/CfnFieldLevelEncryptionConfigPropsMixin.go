package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudFront::FieldLevelEncryptionConfig.
//
// A field-level encryption configuration, which specifies the query arg profiles and content type profiles CloudFront uses to encrypt sensitive fields in viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFieldLevelEncryptionConfigPropsMixin := awscdkcfnpropertymixins.Aws_cloudfront.NewCfnFieldLevelEncryptionConfigPropsMixin(&CfnFieldLevelEncryptionConfigMixinProps{
//   	CallerReference: jsii.String("callerReference"),
//   	Comment: jsii.String("comment"),
//   	ContentTypeProfileConfig: &ContentTypeProfileConfigProperty{
//   		ContentTypeProfiles: []interface{}{
//   			&ContentTypeProfileProperty{
//   				ContentType: jsii.String("contentType"),
//   				Format: jsii.String("format"),
//   				ProfileId: jsii.String("profileId"),
//   			},
//   		},
//   		ForwardWhenContentTypeIsUnknown: jsii.Boolean(false),
//   	},
//   	QueryArgProfileConfig: &QueryArgProfileConfigProperty{
//   		ForwardWhenQueryArgProfileIsUnknown: jsii.Boolean(false),
//   		QueryArgProfiles: []interface{}{
//   			&QueryArgProfileProperty{
//   				ProfileId: jsii.String("profileId"),
//   				QueryArg: jsii.String("queryArg"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html
//
type CfnFieldLevelEncryptionConfigPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFieldLevelEncryptionConfigMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFieldLevelEncryptionConfigPropsMixin
type jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin) Props() *CfnFieldLevelEncryptionConfigMixinProps {
	var returns *CfnFieldLevelEncryptionConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::FieldLevelEncryptionConfig`.
func NewCfnFieldLevelEncryptionConfigPropsMixin(props *CfnFieldLevelEncryptionConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFieldLevelEncryptionConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFieldLevelEncryptionConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFieldLevelEncryptionConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::FieldLevelEncryptionConfig`.
func NewCfnFieldLevelEncryptionConfigPropsMixin_Override(c CfnFieldLevelEncryptionConfigPropsMixin, props *CfnFieldLevelEncryptionConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFieldLevelEncryptionConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFieldLevelEncryptionConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFieldLevelEncryptionConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFieldLevelEncryptionConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFieldLevelEncryptionConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFieldLevelEncryptionConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFieldLevelEncryptionConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

