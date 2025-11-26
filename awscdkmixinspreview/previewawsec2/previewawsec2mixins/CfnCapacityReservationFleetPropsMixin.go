package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new Capacity Reservation Fleet with the specified attributes.
//
// For more information, see [Capacity Reservation Fleets](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cr-fleets.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapacityReservationFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnCapacityReservationFleetPropsMixin(&CfnCapacityReservationFleetMixinProps{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	EndDate: jsii.String("endDate"),
//   	InstanceMatchCriteria: jsii.String("instanceMatchCriteria"),
//   	InstanceTypeSpecifications: []interface{}{
//   		&InstanceTypeSpecificationProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   			EbsOptimized: jsii.Boolean(false),
//   			InstancePlatform: jsii.String("instancePlatform"),
//   			InstanceType: jsii.String("instanceType"),
//   			Priority: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	NoRemoveEndDate: jsii.Boolean(false),
//   	RemoveEndDate: jsii.Boolean(false),
//   	TagSpecifications: []TagSpecificationProperty{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Tenancy: jsii.String("tenancy"),
//   	TotalTargetCapacity: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservationfleet.html
//
type CfnCapacityReservationFleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCapacityReservationFleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCapacityReservationFleetPropsMixin
type jsiiProxy_CfnCapacityReservationFleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCapacityReservationFleetPropsMixin) Props() *CfnCapacityReservationFleetMixinProps {
	var returns *CfnCapacityReservationFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityReservationFleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::CapacityReservationFleet`.
func NewCfnCapacityReservationFleetPropsMixin(props *CfnCapacityReservationFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCapacityReservationFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCapacityReservationFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCapacityReservationFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnCapacityReservationFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::CapacityReservationFleet`.
func NewCfnCapacityReservationFleetPropsMixin_Override(c CfnCapacityReservationFleetPropsMixin, props *CfnCapacityReservationFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnCapacityReservationFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCapacityReservationFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapacityReservationFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnCapacityReservationFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapacityReservationFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnCapacityReservationFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCapacityReservationFleetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCapacityReservationFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

