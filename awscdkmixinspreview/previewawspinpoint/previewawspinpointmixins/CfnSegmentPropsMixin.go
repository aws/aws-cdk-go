package previewawspinpointmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpoint/previewawspinpointmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Updates the configuration, dimension, and other settings for an existing segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributes interface{}
//   var metrics interface{}
//   var tags interface{}
//   var userAttributes interface{}
//
//   cfnSegmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnSegmentPropsMixin(&CfnSegmentMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Dimensions: &SegmentDimensionsProperty{
//   		Attributes: attributes,
//   		Behavior: &BehaviorProperty{
//   			Recency: &RecencyProperty{
//   				Duration: jsii.String("duration"),
//   				RecencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		Demographic: &DemographicProperty{
//   			AppVersion: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Channel: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			DeviceType: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Make: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Model: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Platform: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Location: &LocationProperty{
//   			Country: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			GpsPoint: &GPSPointProperty{
//   				Coordinates: &CoordinatesProperty{
//   					Latitude: jsii.Number(123),
//   					Longitude: jsii.Number(123),
//   				},
//   				RangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		Metrics: metrics,
//   		UserAttributes: userAttributes,
//   	},
//   	Name: jsii.String("name"),
//   	SegmentGroups: &SegmentGroupsProperty{
//   		Groups: []interface{}{
//   			&GroupsProperty{
//   				Dimensions: []interface{}{
//   					&SegmentDimensionsProperty{
//   						Attributes: attributes,
//   						Behavior: &BehaviorProperty{
//   							Recency: &RecencyProperty{
//   								Duration: jsii.String("duration"),
//   								RecencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						Demographic: &DemographicProperty{
//   							AppVersion: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Channel: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							DeviceType: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Make: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Model: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Platform: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						Location: &LocationProperty{
//   							Country: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							GpsPoint: &GPSPointProperty{
//   								Coordinates: &CoordinatesProperty{
//   									Latitude: jsii.Number(123),
//   									Longitude: jsii.Number(123),
//   								},
//   								RangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						Metrics: metrics,
//   						UserAttributes: userAttributes,
//   					},
//   				},
//   				SourceSegments: []interface{}{
//   					&SourceSegmentsProperty{
//   						Id: jsii.String("id"),
//   						Version: jsii.Number(123),
//   					},
//   				},
//   				SourceType: jsii.String("sourceType"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Include: jsii.String("include"),
//   	},
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html
//
type CfnSegmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSegmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSegmentPropsMixin
type jsiiProxy_CfnSegmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSegmentPropsMixin) Props() *CfnSegmentMixinProps {
	var returns *CfnSegmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::Segment`.
func NewCfnSegmentPropsMixin(props *CfnSegmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSegmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSegmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSegmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnSegmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::Segment`.
func NewCfnSegmentPropsMixin_Override(c CfnSegmentPropsMixin, props *CfnSegmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnSegmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSegmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnSegmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSegmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnSegmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSegmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSegmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

