package previewawsodbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsodb/previewawsodbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::CloudExadataInfrastructure` resource creates an Exadata infrastructure.
//
// An Exadata infrastructure provides the underlying compute and storage resources for Oracle Database workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudExadataInfrastructurePropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudExadataInfrastructurePropsMixin(&CfnCloudExadataInfrastructureMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	ComputeCount: jsii.Number(123),
//   	CustomerContactsToSendToOci: []interface{}{
//   		&CustomerContactProperty{
//   			Email: jsii.String("email"),
//   		},
//   	},
//   	DatabaseServerType: jsii.String("databaseServerType"),
//   	DisplayName: jsii.String("displayName"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		CustomActionTimeoutInMins: jsii.Number(123),
//   		DaysOfWeek: []*string{
//   			jsii.String("daysOfWeek"),
//   		},
//   		HoursOfDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IsCustomActionTimeoutEnabled: jsii.Boolean(false),
//   		LeadTimeInWeeks: jsii.Number(123),
//   		Months: []*string{
//   			jsii.String("months"),
//   		},
//   		PatchingMode: jsii.String("patchingMode"),
//   		Preference: jsii.String("preference"),
//   		WeeksOfMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	Shape: jsii.String("shape"),
//   	StorageCount: jsii.Number(123),
//   	StorageServerType: jsii.String("storageServerType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html
//
type CfnCloudExadataInfrastructurePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudExadataInfrastructureMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudExadataInfrastructurePropsMixin
type jsiiProxy_CfnCloudExadataInfrastructurePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudExadataInfrastructurePropsMixin) Props() *CfnCloudExadataInfrastructureMixinProps {
	var returns *CfnCloudExadataInfrastructureMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructurePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ODB::CloudExadataInfrastructure`.
func NewCfnCloudExadataInfrastructurePropsMixin(props *CfnCloudExadataInfrastructureMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudExadataInfrastructurePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudExadataInfrastructurePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudExadataInfrastructurePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ODB::CloudExadataInfrastructure`.
func NewCfnCloudExadataInfrastructurePropsMixin_Override(c CfnCloudExadataInfrastructurePropsMixin, props *CfnCloudExadataInfrastructureMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudExadataInfrastructurePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudExadataInfrastructurePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudExadataInfrastructurePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructurePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructurePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

