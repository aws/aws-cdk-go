package previewawsmaciemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmacie/previewawsmaciemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Macie::AllowList` resource specifies an allow list.
//
// In Amazon Macie , an allow list defines specific text or a text pattern for Macie to ignore when it inspects data sources for sensitive data. If data matches text or a text pattern in an allow list, Macie doesn’t report the data in sensitive data findings or sensitive data discovery results, even if the data matches the criteria of a custom data identifier or a managed data identifier. You can create and use allow lists in all the AWS Regions where Macie is currently available except the Asia Pacific (Osaka) Region.
//
// Macie supports two types of allow lists:
//
// - *Predefined text* - For this type of list ( `S3WordsList` ), you create a line-delimited plaintext file that lists specific text to ignore, and you store the file in an Amazon Simple Storage Service ( Amazon S3 ) bucket. You then configure settings for Macie to access the list in the bucket.
//
// This type of list typically contains specific words, phrases, and other kinds of character sequences that aren’t sensitive, aren't likely to change, and don’t necessarily adhere to a common pattern. If you use this type of list, Macie doesn't report occurrences of text that exactly match a complete entry in the list. Macie treats each entry in the list as a string literal value. Matches aren't case sensitive.
// - *Regular expression* - For this type of list ( `Regex` ), you specify a regular expression that defines a text pattern to ignore. Unlike an allow list with predefined text, you store the regex and all other list settings in Macie .
//
// This type of list is helpful if you want to specify text that isn’t sensitive but varies or is likely to change while also adhering to a common pattern. If you use this type of list, Macie doesn't report occurrences of text that completely match the pattern defined by the list.
//
// For more information, see [Defining sensitive data exceptions with allow lists](https://docs.aws.amazon.com/macie/latest/user/allow-lists.html) in the *Amazon Macie User Guide* .
//
// An `AWS::Macie::Session` resource must exist for an AWS account before you can create an `AWS::Macie::AllowList` resource for the account. Use a [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to ensure that an `AWS::Macie::Session` resource is created before other Macie resources are created for an account. For example, `"DependsOn": "Session"` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAllowListPropsMixin := awscdkmixinspreview.Mixins.NewCfnAllowListPropsMixin(&CfnAllowListMixinProps{
//   	Criteria: &CriteriaProperty{
//   		Regex: jsii.String("regex"),
//   		S3WordsList: &S3WordsListProperty{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html
//
type CfnAllowListPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAllowListMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAllowListPropsMixin
type jsiiProxy_CfnAllowListPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAllowListPropsMixin) Props() *CfnAllowListMixinProps {
	var returns *CfnAllowListMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAllowListPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Macie::AllowList`.
func NewCfnAllowListPropsMixin(props *CfnAllowListMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAllowListPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAllowListPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAllowListPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnAllowListPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Macie::AllowList`.
func NewCfnAllowListPropsMixin_Override(c CfnAllowListPropsMixin, props *CfnAllowListMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnAllowListPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAllowListPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAllowListPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnAllowListPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAllowListPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnAllowListPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAllowListPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAllowListPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

