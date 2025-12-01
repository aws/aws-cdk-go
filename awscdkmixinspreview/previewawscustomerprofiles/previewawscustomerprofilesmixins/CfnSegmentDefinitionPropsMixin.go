package previewawscustomerprofilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscustomerprofiles/previewawscustomerprofilesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A segment definition resource of Amazon Connect Customer Profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSegmentDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnSegmentDefinitionPropsMixin(&CfnSegmentDefinitionMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DomainName: jsii.String("domainName"),
//   	SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   	SegmentGroups: &SegmentGroupProperty{
//   		Groups: []interface{}{
//   			&GroupProperty{
//   				Dimensions: []interface{}{
//   					&DimensionProperty{
//   						CalculatedAttributes: map[string]interface{}{
//   							"calculatedAttributesKey": &CalculatedAttributeDimensionProperty{
//   								"conditionOverrides": &ConditionOverridesProperty{
//   									"range": &RangeOverrideProperty{
//   										"end": jsii.Number(123),
//   										"start": jsii.Number(123),
//   										"unit": jsii.String("unit"),
//   									},
//   								},
//   								"dimensionType": jsii.String("dimensionType"),
//   								"values": []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						ProfileAttributes: &ProfileAttributesProperty{
//   							AccountNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							AdditionalInformation: &ExtraLengthValueProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Address: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							Attributes: map[string]interface{}{
//   								"attributesKey": &AttributeDimensionProperty{
//   									"dimensionType": jsii.String("dimensionType"),
//   									"values": []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							BillingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							BirthDate: &DateDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessEmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							BusinessPhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							EmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							FirstName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							GenderString: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							HomePhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							LastName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							MailingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   							MiddleName: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							MobilePhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PartyTypeString: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PersonalEmailAddress: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							PhoneNumber: &ProfileDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							ProfileType: &ProfileTypeDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							ShippingAddress: &AddressDimensionProperty{
//   								City: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Country: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								County: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								PostalCode: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								Province: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   								State: &ProfileDimensionProperty{
//   									DimensionType: jsii.String("dimensionType"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				SourceSegments: []interface{}{
//   					&SourceSegmentProperty{
//   						SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   					},
//   				},
//   				SourceType: jsii.String("sourceType"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Include: jsii.String("include"),
//   	},
//   	SegmentSqlQuery: jsii.String("segmentSqlQuery"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html
//
type CfnSegmentDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSegmentDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSegmentDefinitionPropsMixin
type jsiiProxy_CfnSegmentDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSegmentDefinitionPropsMixin) Props() *CfnSegmentDefinitionMixinProps {
	var returns *CfnSegmentDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegmentDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CustomerProfiles::SegmentDefinition`.
func NewCfnSegmentDefinitionPropsMixin(props *CfnSegmentDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSegmentDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSegmentDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSegmentDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnSegmentDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CustomerProfiles::SegmentDefinition`.
func NewCfnSegmentDefinitionPropsMixin_Override(c CfnSegmentDefinitionPropsMixin, props *CfnSegmentDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnSegmentDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSegmentDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSegmentDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnSegmentDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSegmentDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnSegmentDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSegmentDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSegmentDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

