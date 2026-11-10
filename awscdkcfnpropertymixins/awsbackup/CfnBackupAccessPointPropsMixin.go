package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a backup access point for an Amazon S3 recovery point.
//
// A backup access point provides on-demand, read-only access to the backup data in a recovery point through an Amazon S3 access point, without initiating a restore. While a backup access point is active for a recovery point, Backup pauses lifecycle transitions and blocks deletion of that recovery point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPointPolicy interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnBackupAccessPointPropsMixin := awscdkcfnpropertymixins.Aws_backup.NewCfnBackupAccessPointPropsMixin(&CfnBackupAccessPointMixinProps{
//   	AccessPointInTime: jsii.String("accessPointInTime"),
//   	AccessPointPolicy: accessPointPolicy,
//   	Name: jsii.String("name"),
//   	RecoveryPointArn: jsii.String("recoveryPointArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html
//
type CfnBackupAccessPointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBackupAccessPointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBackupAccessPointPropsMixin
type jsiiProxy_CfnBackupAccessPointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnBackupAccessPointPropsMixin) Props() *CfnBackupAccessPointMixinProps {
	var returns *CfnBackupAccessPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBackupAccessPointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::BackupAccessPoint`.
func NewCfnBackupAccessPointPropsMixin(props *CfnBackupAccessPointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBackupAccessPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBackupAccessPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBackupAccessPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupAccessPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::BackupAccessPoint`.
func NewCfnBackupAccessPointPropsMixin_Override(c CfnBackupAccessPointPropsMixin, props *CfnBackupAccessPointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupAccessPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBackupAccessPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBackupAccessPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupAccessPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBackupAccessPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupAccessPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBackupAccessPointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBackupAccessPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

