package previewawskmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskms/previewawskmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::KMS::Alias` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys) . You can use an alias to identify a KMS key in the AWS  console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) , such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html) .
//
// > Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for AWS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *AWS Key Management Service Developer Guide* .
//
// Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS Regions . For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *AWS Key Management Service Developer Guide* .
//
// When specifying an alias, observe the following rules.
//
// - Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.
// - The alias and its associated KMS key must be in the same AWS account and Region.
// - The alias name must be unique in the AWS account and Region. However, you can create aliases with the same name in different AWS Regions . For example, you can have an `alias/projectKey` in multiple Regions, each of which is associated with a KMS key in its Region.
// - Each alias name must begin with `alias/` followed by a name, such as `alias/exampleKey` . The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with `alias/aws/` . That alias name prefix is reserved for [AWS managed keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) .
//
// *Regions*
//
// AWS  CloudFormation resources are available in all AWS Regions in which AWS  and CloudFormation are supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAliasPropsMixin := awscdkmixinspreview.Mixins.NewCfnAliasPropsMixin(&CfnAliasMixinProps{
//   	AliasName: jsii.String("aliasName"),
//   	TargetKeyId: jsii.String("targetKeyId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
//
type CfnAliasPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAliasMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAliasPropsMixin
type jsiiProxy_CfnAliasPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAliasPropsMixin) Props() *CfnAliasMixinProps {
	var returns *CfnAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAliasPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KMS::Alias`.
func NewCfnAliasPropsMixin(props *CfnAliasMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.mixins.CfnAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KMS::Alias`.
func NewCfnAliasPropsMixin_Override(c CfnAliasPropsMixin, props *CfnAliasMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.mixins.CfnAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kms.mixins.CfnAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kms.mixins.CfnAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAliasPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

