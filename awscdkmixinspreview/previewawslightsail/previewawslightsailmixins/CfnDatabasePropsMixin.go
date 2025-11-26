package previewawslightsailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslightsail/previewawslightsailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::Database` resource specifies an Amazon Lightsail database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDatabasePropsMixin := awscdkmixinspreview.Mixins.NewCfnDatabasePropsMixin(&CfnDatabaseMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BackupRetention: jsii.Boolean(false),
//   	CaCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	MasterDatabaseName: jsii.String("masterDatabaseName"),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	RelationalDatabaseBlueprintId: jsii.String("relationalDatabaseBlueprintId"),
//   	RelationalDatabaseBundleId: jsii.String("relationalDatabaseBundleId"),
//   	RelationalDatabaseName: jsii.String("relationalDatabaseName"),
//   	RelationalDatabaseParameters: []interface{}{
//   		&RelationalDatabaseParameterProperty{
//   			AllowedValues: jsii.String("allowedValues"),
//   			ApplyMethod: jsii.String("applyMethod"),
//   			ApplyType: jsii.String("applyType"),
//   			DataType: jsii.String("dataType"),
//   			Description: jsii.String("description"),
//   			IsModifiable: jsii.Boolean(false),
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	RotateMasterUserPassword: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html
//
type CfnDatabasePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDatabaseMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatabasePropsMixin
type jsiiProxy_CfnDatabasePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDatabasePropsMixin) Props() *CfnDatabaseMixinProps {
	var returns *CfnDatabaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabasePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::Database`.
func NewCfnDatabasePropsMixin(props *CfnDatabaseMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDatabasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatabasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatabasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::Database`.
func NewCfnDatabasePropsMixin_Override(c CfnDatabasePropsMixin, props *CfnDatabaseMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDatabasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatabasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatabasePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDatabasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

