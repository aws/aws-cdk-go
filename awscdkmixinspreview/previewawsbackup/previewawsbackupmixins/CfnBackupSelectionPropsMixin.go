package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a set of resources to assign to a backup plan.
//
// For a sample CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var conditions interface{}
//
//   cfnBackupSelectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnBackupSelectionPropsMixin(&CfnBackupSelectionMixinProps{
//   	BackupPlanId: jsii.String("backupPlanId"),
//   	BackupSelection: &BackupSelectionResourceTypeProperty{
//   		Conditions: conditions,
//   		IamRoleArn: jsii.String("iamRoleArn"),
//   		ListOfTags: []interface{}{
//   			&ConditionResourceTypeProperty{
//   				ConditionKey: jsii.String("conditionKey"),
//   				ConditionType: jsii.String("conditionType"),
//   				ConditionValue: jsii.String("conditionValue"),
//   			},
//   		},
//   		NotResources: []*string{
//   			jsii.String("notResources"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		SelectionName: jsii.String("selectionName"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html
//
type CfnBackupSelectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBackupSelectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBackupSelectionPropsMixin
type jsiiProxy_CfnBackupSelectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBackupSelectionPropsMixin) Props() *CfnBackupSelectionMixinProps {
	var returns *CfnBackupSelectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupSelectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::BackupSelection`.
func NewCfnBackupSelectionPropsMixin(props *CfnBackupSelectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBackupSelectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBackupSelectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBackupSelectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::BackupSelection`.
func NewCfnBackupSelectionPropsMixin_Override(c CfnBackupSelectionPropsMixin, props *CfnBackupSelectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBackupSelectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBackupSelectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupSelectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBackupSelectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBackupSelectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

