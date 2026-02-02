package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Contains an optional backup plan display name and an array of `BackupRule` objects, each of which specifies a backup rule.
//
// Each rule in a backup plan is a separate scheduled task and can back up a different selection of AWS resources.
//
// For a sample CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var backupOptions interface{}
//
//   cfnBackupPlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnBackupPlanPropsMixin(&CfnBackupPlanMixinProps{
//   	BackupPlan: &BackupPlanResourceTypeProperty{
//   		AdvancedBackupSettings: []interface{}{
//   			&AdvancedBackupSettingResourceTypeProperty{
//   				BackupOptions: backupOptions,
//   				ResourceType: jsii.String("resourceType"),
//   			},
//   		},
//   		BackupPlanName: jsii.String("backupPlanName"),
//   		BackupPlanRule: []interface{}{
//   			&BackupRuleResourceTypeProperty{
//   				CompletionWindowMinutes: jsii.Number(123),
//   				CopyActions: []interface{}{
//   					&CopyActionResourceTypeProperty{
//   						DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//   						Lifecycle: &LifecycleResourceTypeProperty{
//   							DeleteAfterDays: jsii.Number(123),
//   							MoveToColdStorageAfterDays: jsii.Number(123),
//   							OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				EnableContinuousBackup: jsii.Boolean(false),
//   				IndexActions: []interface{}{
//   					&IndexActionsResourceTypeProperty{
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   				Lifecycle: &LifecycleResourceTypeProperty{
//   					DeleteAfterDays: jsii.Number(123),
//   					MoveToColdStorageAfterDays: jsii.Number(123),
//   					OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   				},
//   				RecoveryPointTags: map[string]*string{
//   					"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   				},
//   				RuleName: jsii.String("ruleName"),
//   				ScanActions: []interface{}{
//   					&ScanActionResourceTypeProperty{
//   						MalwareScanner: jsii.String("malwareScanner"),
//   						ScanMode: jsii.String("scanMode"),
//   					},
//   				},
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   				ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   				StartWindowMinutes: jsii.Number(123),
//   				TargetBackupVault: jsii.String("targetBackupVault"),
//   				TargetLogicallyAirGappedBackupVaultArn: jsii.String("targetLogicallyAirGappedBackupVaultArn"),
//   			},
//   		},
//   		ScanSettings: []interface{}{
//   			&ScanSettingResourceTypeProperty{
//   				MalwareScanner: jsii.String("malwareScanner"),
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   				ScannerRoleArn: jsii.String("scannerRoleArn"),
//   			},
//   		},
//   	},
//   	BackupPlanTags: map[string]*string{
//   		"backupPlanTagsKey": jsii.String("backupPlanTags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html
//
type CfnBackupPlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBackupPlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBackupPlanPropsMixin
type jsiiProxy_CfnBackupPlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBackupPlanPropsMixin) Props() *CfnBackupPlanMixinProps {
	var returns *CfnBackupPlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupPlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::BackupPlan`.
func NewCfnBackupPlanPropsMixin(props *CfnBackupPlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBackupPlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBackupPlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBackupPlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::BackupPlan`.
func NewCfnBackupPlanPropsMixin_Override(c CfnBackupPlanPropsMixin, props *CfnBackupPlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBackupPlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBackupPlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupPlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBackupPlanPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBackupPlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

