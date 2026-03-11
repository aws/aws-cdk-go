package awsiotfleetwise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a vehicle model (model manifest) that specifies signals (attributes, branches, sensors, and actuators).
//
// For more information, see [Vehicle models](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/vehicle-models.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnModelManifestPropsMixin := awscdkcfnpropertymixins.Aws_iotfleetwise.NewCfnModelManifestPropsMixin(&CfnModelManifestMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Nodes: []*string{
//   		jsii.String("nodes"),
//   	},
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html
//
type CfnModelManifestPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnModelManifestMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelManifestPropsMixin
type jsiiProxy_CfnModelManifestPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnModelManifestPropsMixin) Props() *CfnModelManifestMixinProps {
	var returns *CfnModelManifestMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelManifestPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTFleetWise::ModelManifest`.
func NewCfnModelManifestPropsMixin(props *CfnModelManifestMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnModelManifestPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelManifestPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelManifestPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTFleetWise::ModelManifest`.
func NewCfnModelManifestPropsMixin_Override(c CfnModelManifestPropsMixin, props *CfnModelManifestMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnModelManifestPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelManifestPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelManifestPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelManifestPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnModelManifestPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

