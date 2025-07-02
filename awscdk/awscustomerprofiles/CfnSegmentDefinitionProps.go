package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSegmentDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSegmentDefinitionProps := &CfnSegmentDefinitionProps{
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
//   								"dimensionType": jsii.String("dimensionType"),
//   								"values": []*string{
//   									jsii.String("values"),
//   								},
//
//   								// the properties below are optional
//   								"conditionOverrides": &ConditionOverridesProperty{
//   									"range": &RangeOverrideProperty{
//   										"start": jsii.Number(123),
//   										"unit": jsii.String("unit"),
//
//   										// the properties below are optional
//   										"end": jsii.Number(123),
//   									},
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html
//
type CfnSegmentDefinitionProps struct {
	// Display name of the segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Name of the segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-segmentdefinitionname
	//
	SegmentDefinitionName *string `field:"required" json:"segmentDefinitionName" yaml:"segmentDefinitionName"`
	// Contains all groups of the segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-segmentgroups
	//
	SegmentGroups interface{} `field:"required" json:"segmentGroups" yaml:"segmentGroups"`
	// The description of the segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags belonging to the segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-segmentdefinition.html#cfn-customerprofiles-segmentdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

