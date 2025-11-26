package previewawscustomerprofilesmixins


// Defines the attribute to segment on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dimensionProperty := &DimensionProperty{
//   	CalculatedAttributes: map[string]interface{}{
//   		"calculatedAttributesKey": &CalculatedAttributeDimensionProperty{
//   			"conditionOverrides": &ConditionOverridesProperty{
//   				"range": &RangeOverrideProperty{
//   					"end": jsii.Number(123),
//   					"start": jsii.Number(123),
//   					"unit": jsii.String("unit"),
//   				},
//   			},
//   			"dimensionType": jsii.String("dimensionType"),
//   			"values": []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	ProfileAttributes: &ProfileAttributesProperty{
//   		AccountNumber: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		AdditionalInformation: &ExtraLengthValueProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Address: &AddressDimensionProperty{
//   			City: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Country: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			County: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PostalCode: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Province: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			State: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Attributes: map[string]interface{}{
//   			"attributesKey": &AttributeDimensionProperty{
//   				"dimensionType": jsii.String("dimensionType"),
//   				"values": []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		BillingAddress: &AddressDimensionProperty{
//   			City: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Country: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			County: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PostalCode: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Province: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			State: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		BirthDate: &DateDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		BusinessEmailAddress: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		BusinessName: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		BusinessPhoneNumber: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		EmailAddress: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		FirstName: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		GenderString: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		HomePhoneNumber: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		LastName: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		MailingAddress: &AddressDimensionProperty{
//   			City: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Country: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			County: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PostalCode: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Province: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			State: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		MiddleName: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		MobilePhoneNumber: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PartyTypeString: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PersonalEmailAddress: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PhoneNumber: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		ProfileType: &ProfileTypeDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		ShippingAddress: &AddressDimensionProperty{
//   			City: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Country: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			County: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PostalCode: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Province: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			State: &ProfileDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-dimension.html
//
type CfnSegmentDefinitionPropsMixin_DimensionProperty struct {
	// Object that holds the calculated attributes to segment on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-dimension.html#cfn-customerprofiles-segmentdefinition-dimension-calculatedattributes
	//
	CalculatedAttributes interface{} `field:"optional" json:"calculatedAttributes" yaml:"calculatedAttributes"`
	// Object that holds the profile attributes to segment on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-dimension.html#cfn-customerprofiles-segmentdefinition-dimension-profileattributes
	//
	ProfileAttributes interface{} `field:"optional" json:"profileAttributes" yaml:"profileAttributes"`
}

