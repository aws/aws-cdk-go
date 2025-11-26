package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotfleetwise/previewawsiotfleetwisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a vehicle, which is an instance of a vehicle model (model manifest).
//
// Vehicles created from the same vehicle model consist of the same signals inherited from the vehicle model.
//
// > If you have an existing AWS IoT thing, you can use AWS IoT FleetWise to create a vehicle and collect data from your thing.
//
// For more information, see [Vehicles](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/vehicles.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var onChange interface{}
//
//   cfnVehiclePropsMixin := awscdkmixinspreview.Mixins.NewCfnVehiclePropsMixin(&CfnVehicleMixinProps{
//   	AssociationBehavior: jsii.String("associationBehavior"),
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	DecoderManifestArn: jsii.String("decoderManifestArn"),
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//   	StateTemplates: []interface{}{
//   		&StateTemplateAssociationProperty{
//   			Identifier: jsii.String("identifier"),
//   			StateTemplateUpdateStrategy: &StateTemplateUpdateStrategyProperty{
//   				OnChange: onChange,
//   				Periodic: &PeriodicStateTemplateUpdateStrategyProperty{
//   					StateTemplateUpdateRate: &TimePeriodProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html
//
type CfnVehiclePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVehicleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVehiclePropsMixin
type jsiiProxy_CfnVehiclePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVehiclePropsMixin) Props() *CfnVehicleMixinProps {
	var returns *CfnVehicleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVehiclePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTFleetWise::Vehicle`.
func NewCfnVehiclePropsMixin(props *CfnVehicleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVehiclePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVehiclePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVehiclePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTFleetWise::Vehicle`.
func NewCfnVehiclePropsMixin_Override(c CfnVehiclePropsMixin, props *CfnVehicleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVehiclePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVehiclePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVehiclePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVehiclePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVehiclePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

