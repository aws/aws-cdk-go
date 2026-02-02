package previewawsneptunemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsneptune/previewawsneptunemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Neptune::DBInstance` type creates an Amazon Neptune DB instance.
//
// *Updating DB Instances*
//
// You can set a deletion policy for your DB instance to control how CloudFormation handles the instance when the stack is deleted. For Neptune DB instances, you can choose to *retain* the instance, to *delete* the instance, or to *create a snapshot* of the instance. The default CloudFormation behavior depends on the `DBClusterIdentifier` property:
//
// - For `AWS::Neptune::DBInstance` resources that don't specify the `DBClusterIdentifier` property, CloudFormation saves a snapshot of the DB instance.
// - For `AWS::Neptune::DBInstance` resources that do specify the `DBClusterIdentifier` property, CloudFormation deletes the DB instance.
//
// *Deleting DB Instances*
//
// > If a DB instance is deleted or replaced during an update, CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced.
//
// When properties labeled *Update requires: Replacement* are updated, CloudFormation first creates a replacement DB instance, changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.
// > - Create a snapshot of the DB instance.
// > - If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the `DBSnapshotIdentifier` property with the ID of the DB snapshot that you want to use.
// > - Update the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnDBInstancePropsMixin(&CfnDBInstanceMixinProps{
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbInstanceClass: jsii.String("dbInstanceClass"),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	DbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html
//
type CfnDBInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBInstancePropsMixin
type jsiiProxy_CfnDBInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBInstancePropsMixin) Props() *CfnDBInstanceMixinProps {
	var returns *CfnDBInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Neptune::DBInstance`.
func NewCfnDBInstancePropsMixin(props *CfnDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Neptune::DBInstance`.
func NewCfnDBInstancePropsMixin_Override(c CfnDBInstancePropsMixin, props *CfnDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_neptune.mixins.CfnDBInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBInstancePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDBInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

