package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a restore testing plan.
//
// The first of two steps to create a restore testing plan. After this request is successful, finish the procedure using CreateRestoreTestingSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestoreTestingPlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnRestoreTestingPlanPropsMixin(&CfnRestoreTestingPlanMixinProps{
//   	RecoveryPointSelection: &RestoreTestingRecoveryPointSelectionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		ExcludeVaults: []*string{
//   			jsii.String("excludeVaults"),
//   		},
//   		IncludeVaults: []*string{
//   			jsii.String("includeVaults"),
//   		},
//   		RecoveryPointTypes: []*string{
//   			jsii.String("recoveryPointTypes"),
//   		},
//   		SelectionWindowDays: jsii.Number(123),
//   	},
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	StartWindowHours: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html
//
type CfnRestoreTestingPlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRestoreTestingPlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRestoreTestingPlanPropsMixin
type jsiiProxy_CfnRestoreTestingPlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRestoreTestingPlanPropsMixin) Props() *CfnRestoreTestingPlanMixinProps {
	var returns *CfnRestoreTestingPlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestoreTestingPlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::RestoreTestingPlan`.
func NewCfnRestoreTestingPlanPropsMixin(props *CfnRestoreTestingPlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRestoreTestingPlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRestoreTestingPlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRestoreTestingPlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::RestoreTestingPlan`.
func NewCfnRestoreTestingPlanPropsMixin_Override(c CfnRestoreTestingPlanPropsMixin, props *CfnRestoreTestingPlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRestoreTestingPlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRestoreTestingPlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRestoreTestingPlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRestoreTestingPlanPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRestoreTestingPlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

