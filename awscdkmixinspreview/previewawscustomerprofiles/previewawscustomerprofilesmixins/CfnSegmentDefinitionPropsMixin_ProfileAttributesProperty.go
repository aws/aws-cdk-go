package previewawscustomerprofilesmixins


// The object used to segment on attributes within the customer profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   profileAttributesProperty := &ProfileAttributesProperty{
//   	AccountNumber: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	AdditionalInformation: &ExtraLengthValueProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Address: &AddressDimensionProperty{
//   		City: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Country: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		County: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PostalCode: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Province: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		State: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Attributes: map[string]interface{}{
//   		"attributesKey": &AttributeDimensionProperty{
//   			"dimensionType": jsii.String("dimensionType"),
//   			"values": []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	BillingAddress: &AddressDimensionProperty{
//   		City: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Country: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		County: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PostalCode: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Province: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		State: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	BirthDate: &DateDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	BusinessEmailAddress: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	BusinessName: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	BusinessPhoneNumber: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	EmailAddress: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	FirstName: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	GenderString: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	HomePhoneNumber: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	LastName: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	MailingAddress: &AddressDimensionProperty{
//   		City: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Country: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		County: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PostalCode: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Province: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		State: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	MiddleName: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	MobilePhoneNumber: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	PartyTypeString: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	PersonalEmailAddress: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	PhoneNumber: &ProfileDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	ProfileType: &ProfileTypeDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	ShippingAddress: &AddressDimensionProperty{
//   		City: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Country: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		County: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		PostalCode: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Province: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		State: &ProfileDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html
//
type CfnSegmentDefinitionPropsMixin_ProfileAttributesProperty struct {
	// A field to describe values to segment on within account number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-accountnumber
	//
	AccountNumber interface{} `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// A field to describe values to segment on within additional information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-additionalinformation
	//
	AdditionalInformation interface{} `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// A field to describe values to segment on within address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-address
	//
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// A field to describe values to segment on within attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// A field to describe values to segment on within billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-billingaddress
	//
	BillingAddress interface{} `field:"optional" json:"billingAddress" yaml:"billingAddress"`
	// A field to describe values to segment on within birthDate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-birthdate
	//
	BirthDate interface{} `field:"optional" json:"birthDate" yaml:"birthDate"`
	// A field to describe values to segment on within business email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessemailaddress
	//
	BusinessEmailAddress interface{} `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// A field to describe values to segment on within business name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessname
	//
	BusinessName interface{} `field:"optional" json:"businessName" yaml:"businessName"`
	// A field to describe values to segment on within business phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-businessphonenumber
	//
	BusinessPhoneNumber interface{} `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// A field to describe values to segment on within email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-emailaddress
	//
	EmailAddress interface{} `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// A field to describe values to segment on within first name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-firstname
	//
	FirstName interface{} `field:"optional" json:"firstName" yaml:"firstName"`
	// A field to describe values to segment on within genderString.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-genderstring
	//
	GenderString interface{} `field:"optional" json:"genderString" yaml:"genderString"`
	// A field to describe values to segment on within home phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-homephonenumber
	//
	HomePhoneNumber interface{} `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// A field to describe values to segment on within last name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-lastname
	//
	LastName interface{} `field:"optional" json:"lastName" yaml:"lastName"`
	// A field to describe values to segment on within mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-mailingaddress
	//
	MailingAddress interface{} `field:"optional" json:"mailingAddress" yaml:"mailingAddress"`
	// A field to describe values to segment on within middle name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-middlename
	//
	MiddleName interface{} `field:"optional" json:"middleName" yaml:"middleName"`
	// A field to describe values to segment on within mobile phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-mobilephonenumber
	//
	MobilePhoneNumber interface{} `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// A field to describe values to segment on within partyTypeString.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-partytypestring
	//
	PartyTypeString interface{} `field:"optional" json:"partyTypeString" yaml:"partyTypeString"`
	// A field to describe values to segment on within personal email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-personalemailaddress
	//
	PersonalEmailAddress interface{} `field:"optional" json:"personalEmailAddress" yaml:"personalEmailAddress"`
	// A field to describe values to segment on within phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-phonenumber
	//
	PhoneNumber interface{} `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The type of profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-profiletype
	//
	ProfileType interface{} `field:"optional" json:"profileType" yaml:"profileType"`
	// A field to describe values to segment on within shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profileattributes.html#cfn-customerprofiles-segmentdefinition-profileattributes-shippingaddress
	//
	ShippingAddress interface{} `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
}

