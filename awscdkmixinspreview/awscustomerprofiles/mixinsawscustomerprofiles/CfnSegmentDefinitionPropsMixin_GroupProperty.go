package mixinsawscustomerprofiles


// Contains dimensions that determine what to segment on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupProperty := &GroupProperty{
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			CalculatedAttributes: map[string]interface{}{
//   				"calculatedAttributesKey": &CalculatedAttributeDimensionProperty{
//   					"conditionOverrides": &ConditionOverridesProperty{
//   						"range": &RangeOverrideProperty{
//   							"end": jsii.Number(123),
//   							"start": jsii.Number(123),
//   							"unit": jsii.String("unit"),
//   						},
//   					},
//   					"dimensionType": jsii.String("dimensionType"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			ProfileAttributes: &ProfileAttributesProperty{
//   				AccountNumber: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				AdditionalInformation: &ExtraLengthValueProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Address: &AddressDimensionProperty{
//   					City: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Country: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					County: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					PostalCode: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Province: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					State: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				Attributes: map[string]interface{}{
//   					"attributesKey": &AttributeDimensionProperty{
//   						"dimensionType": jsii.String("dimensionType"),
//   						"values": []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				BillingAddress: &AddressDimensionProperty{
//   					City: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Country: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					County: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					PostalCode: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Province: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					State: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				BirthDate: &DateDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				BusinessEmailAddress: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				BusinessName: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				BusinessPhoneNumber: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				EmailAddress: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				FirstName: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				GenderString: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				HomePhoneNumber: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				LastName: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				MailingAddress: &AddressDimensionProperty{
//   					City: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Country: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					County: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					PostalCode: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Province: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					State: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				MiddleName: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				MobilePhoneNumber: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				PartyTypeString: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				PersonalEmailAddress: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				PhoneNumber: &ProfileDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				ProfileType: &ProfileTypeDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				ShippingAddress: &AddressDimensionProperty{
//   					City: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Country: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					County: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					PostalCode: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Province: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					State: &ProfileDimensionProperty{
//   						DimensionType: jsii.String("dimensionType"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SourceSegments: []interface{}{
//   		&SourceSegmentProperty{
//   			SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   		},
//   	},
//   	SourceType: jsii.String("sourceType"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-group.html
//
type CfnSegmentDefinitionPropsMixin_GroupProperty struct {
	// Defines the attributes to segment on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-group.html#cfn-customerprofiles-segmentdefinition-group-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Defines the starting source of data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-group.html#cfn-customerprofiles-segmentdefinition-group-sourcesegments
	//
	SourceSegments interface{} `field:"optional" json:"sourceSegments" yaml:"sourceSegments"`
	// Defines how to interact with the source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-group.html#cfn-customerprofiles-segmentdefinition-group-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Defines how to interact with the profiles found in the current filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-group.html#cfn-customerprofiles-segmentdefinition-group-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

