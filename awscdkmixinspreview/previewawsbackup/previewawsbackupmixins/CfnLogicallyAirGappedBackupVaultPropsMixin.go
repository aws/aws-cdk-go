package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a logical container to where backups may be copied.
//
// This request includes a name, the Region, the maximum number of retention days, the minimum number of retention days, and optionally can include tags and a creator request ID.
//
// > Do not include sensitive data, such as passport numbers, in the name of a backup vault.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var accessPolicy interface{}
//
//   cfnLogicallyAirGappedBackupVaultPropsMixin := awscdkmixinspreview.Mixins.NewCfnLogicallyAirGappedBackupVaultPropsMixin(&CfnLogicallyAirGappedBackupVaultMixinProps{
//   	AccessPolicy: accessPolicy,
//   	BackupVaultName: jsii.String("backupVaultName"),
//   	BackupVaultTags: map[string]*string{
//   		"backupVaultTagsKey": jsii.String("backupVaultTags"),
//   	},
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	MaxRetentionDays: jsii.Number(123),
//   	MinRetentionDays: jsii.Number(123),
//   	MpaApprovalTeamArn: jsii.String("mpaApprovalTeamArn"),
//   	Notifications: &NotificationObjectTypeProperty{
//   		BackupVaultEvents: []*string{
//   			jsii.String("backupVaultEvents"),
//   		},
//   		SnsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html
//
type CfnLogicallyAirGappedBackupVaultPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLogicallyAirGappedBackupVaultMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLogicallyAirGappedBackupVaultPropsMixin
type jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin) Props() *CfnLogicallyAirGappedBackupVaultMixinProps {
	var returns *CfnLogicallyAirGappedBackupVaultMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::LogicallyAirGappedBackupVault`.
func NewCfnLogicallyAirGappedBackupVaultPropsMixin(props *CfnLogicallyAirGappedBackupVaultMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLogicallyAirGappedBackupVaultPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLogicallyAirGappedBackupVaultPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::LogicallyAirGappedBackupVault`.
func NewCfnLogicallyAirGappedBackupVaultPropsMixin_Override(c CfnLogicallyAirGappedBackupVaultPropsMixin, props *CfnLogicallyAirGappedBackupVaultMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLogicallyAirGappedBackupVaultPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogicallyAirGappedBackupVaultPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogicallyAirGappedBackupVaultPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

