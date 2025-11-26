package previewawsresourcegroupsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsresourcegroups/previewawsresourcegroupsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Onboards and syncs resources tagged with a specific tag key-value pair to an application.
//
// *Minimum permissions*
//
// To run this command, you must have the following permissions:
//
// - `resource-groups:StartTagSyncTask`
// - `resource-groups:CreateGroup`
// - `iam:PassRole` for the role you provide to create a tag-sync task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTagSyncTaskPropsMixin := awscdkmixinspreview.Mixins.NewCfnTagSyncTaskPropsMixin(&CfnTagSyncTaskMixinProps{
//   	Group: jsii.String("group"),
//   	RoleArn: jsii.String("roleArn"),
//   	TagKey: jsii.String("tagKey"),
//   	TagValue: jsii.String("tagValue"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html
//
type CfnTagSyncTaskPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTagSyncTaskMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTagSyncTaskPropsMixin
type jsiiProxy_CfnTagSyncTaskPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTagSyncTaskPropsMixin) Props() *CfnTagSyncTaskMixinProps {
	var returns *CfnTagSyncTaskMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagSyncTaskPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResourceGroups::TagSyncTask`.
func NewCfnTagSyncTaskPropsMixin(props *CfnTagSyncTaskMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTagSyncTaskPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTagSyncTaskPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTagSyncTaskPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResourceGroups::TagSyncTask`.
func NewCfnTagSyncTaskPropsMixin_Override(c CfnTagSyncTaskPropsMixin, props *CfnTagSyncTaskMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTagSyncTaskPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTagSyncTaskPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTagSyncTaskPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTagSyncTaskPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTagSyncTaskPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

