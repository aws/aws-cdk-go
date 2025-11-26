package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This request can be sent after CreateRestoreTestingPlan request returns successfully.
//
// This is the second part of creating a resource testing plan, and it must be completed sequentially.
//
// This consists of `RestoreTestingSelectionName` , `ProtectedResourceType` , and one of the following:
//
// - `ProtectedResourceArns`
// - `ProtectedResourceConditions`
//
// Each protected resource type can have one single value.
//
// A restore testing selection can include a wildcard value ("*") for `ProtectedResourceArns` along with `ProtectedResourceConditions` . Alternatively, you can include up to 30 specific protected resource ARNs in `ProtectedResourceArns` .
//
// Cannot select by both protected resource types AND specific ARNs. Request will fail if both are included.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestoreTestingSelectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnRestoreTestingSelectionPropsMixin(&CfnRestoreTestingSelectionMixinProps{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	ProtectedResourceArns: []*string{
//   		jsii.String("protectedResourceArns"),
//   	},
//   	ProtectedResourceConditions: &ProtectedResourceConditionsProperty{
//   		StringEquals: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		StringNotEquals: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ProtectedResourceType: jsii.String("protectedResourceType"),
//   	RestoreMetadataOverrides: map[string]*string{
//   		"restoreMetadataOverridesKey": jsii.String("restoreMetadataOverrides"),
//   	},
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   	RestoreTestingSelectionName: jsii.String("restoreTestingSelectionName"),
//   	ValidationWindowHours: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html
//
type CfnRestoreTestingSelectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRestoreTestingSelectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRestoreTestingSelectionPropsMixin
type jsiiProxy_CfnRestoreTestingSelectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRestoreTestingSelectionPropsMixin) Props() *CfnRestoreTestingSelectionMixinProps {
	var returns *CfnRestoreTestingSelectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestoreTestingSelectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::RestoreTestingSelection`.
func NewCfnRestoreTestingSelectionPropsMixin(props *CfnRestoreTestingSelectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRestoreTestingSelectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRestoreTestingSelectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRestoreTestingSelectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::RestoreTestingSelection`.
func NewCfnRestoreTestingSelectionPropsMixin_Override(c CfnRestoreTestingSelectionPropsMixin, props *CfnRestoreTestingSelectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRestoreTestingSelectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRestoreTestingSelectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRestoreTestingSelectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRestoreTestingSelectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRestoreTestingSelectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

