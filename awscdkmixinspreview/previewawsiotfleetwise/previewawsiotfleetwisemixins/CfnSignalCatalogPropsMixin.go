package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotfleetwise/previewawsiotfleetwisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a collection of standardized signals that can be reused to create vehicle models.
//
// For more information, see [Signal catalogs](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/signal-catalogs.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSignalCatalogPropsMixin := awscdkmixinspreview.Mixins.NewCfnSignalCatalogPropsMixin(&CfnSignalCatalogMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NodeCounts: &NodeCountsProperty{
//   		TotalActuators: jsii.Number(123),
//   		TotalAttributes: jsii.Number(123),
//   		TotalBranches: jsii.Number(123),
//   		TotalNodes: jsii.Number(123),
//   		TotalSensors: jsii.Number(123),
//   	},
//   	Nodes: []interface{}{
//   		&NodeProperty{
//   			Actuator: &ActuatorProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				DataType: jsii.String("dataType"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Attribute: &AttributeProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				DataType: jsii.String("dataType"),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Branch: &BranchProperty{
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			},
//   			Sensor: &SensorProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DataType: jsii.String("dataType"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html
//
type CfnSignalCatalogPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSignalCatalogMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSignalCatalogPropsMixin
type jsiiProxy_CfnSignalCatalogPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSignalCatalogPropsMixin) Props() *CfnSignalCatalogMixinProps {
	var returns *CfnSignalCatalogMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalCatalogPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTFleetWise::SignalCatalog`.
func NewCfnSignalCatalogPropsMixin(props *CfnSignalCatalogMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSignalCatalogPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSignalCatalogPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSignalCatalogPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTFleetWise::SignalCatalog`.
func NewCfnSignalCatalogPropsMixin_Override(c CfnSignalCatalogPropsMixin, props *CfnSignalCatalogMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSignalCatalogPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSignalCatalogPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSignalCatalogPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSignalCatalogPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSignalCatalogPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

