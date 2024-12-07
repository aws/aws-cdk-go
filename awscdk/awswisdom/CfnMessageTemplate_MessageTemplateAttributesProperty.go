package awswisdom


// An object that specifies the default values to use for variables in the message template.
//
// This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageTemplateAttributesProperty := &MessageTemplateAttributesProperty{
//   	AgentAttributes: &AgentAttributesProperty{
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   	},
//   	CustomAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	CustomerProfileAttributes: &CustomerProfileAttributesProperty{
//   		AccountNumber: jsii.String("accountNumber"),
//   		AdditionalInformation: jsii.String("additionalInformation"),
//   		Address1: jsii.String("address1"),
//   		Address2: jsii.String("address2"),
//   		Address3: jsii.String("address3"),
//   		Address4: jsii.String("address4"),
//   		BillingAddress1: jsii.String("billingAddress1"),
//   		BillingAddress2: jsii.String("billingAddress2"),
//   		BillingAddress3: jsii.String("billingAddress3"),
//   		BillingAddress4: jsii.String("billingAddress4"),
//   		BillingCity: jsii.String("billingCity"),
//   		BillingCountry: jsii.String("billingCountry"),
//   		BillingCounty: jsii.String("billingCounty"),
//   		BillingPostalCode: jsii.String("billingPostalCode"),
//   		BillingProvince: jsii.String("billingProvince"),
//   		BillingState: jsii.String("billingState"),
//   		BirthDate: jsii.String("birthDate"),
//   		BusinessEmailAddress: jsii.String("businessEmailAddress"),
//   		BusinessName: jsii.String("businessName"),
//   		BusinessPhoneNumber: jsii.String("businessPhoneNumber"),
//   		City: jsii.String("city"),
//   		Country: jsii.String("country"),
//   		County: jsii.String("county"),
//   		Custom: map[string]*string{
//   			"customKey": jsii.String("custom"),
//   		},
//   		EmailAddress: jsii.String("emailAddress"),
//   		FirstName: jsii.String("firstName"),
//   		Gender: jsii.String("gender"),
//   		HomePhoneNumber: jsii.String("homePhoneNumber"),
//   		LastName: jsii.String("lastName"),
//   		MailingAddress1: jsii.String("mailingAddress1"),
//   		MailingAddress2: jsii.String("mailingAddress2"),
//   		MailingAddress3: jsii.String("mailingAddress3"),
//   		MailingAddress4: jsii.String("mailingAddress4"),
//   		MailingCity: jsii.String("mailingCity"),
//   		MailingCountry: jsii.String("mailingCountry"),
//   		MailingCounty: jsii.String("mailingCounty"),
//   		MailingPostalCode: jsii.String("mailingPostalCode"),
//   		MailingProvince: jsii.String("mailingProvince"),
//   		MailingState: jsii.String("mailingState"),
//   		MiddleName: jsii.String("middleName"),
//   		MobilePhoneNumber: jsii.String("mobilePhoneNumber"),
//   		PartyType: jsii.String("partyType"),
//   		PhoneNumber: jsii.String("phoneNumber"),
//   		PostalCode: jsii.String("postalCode"),
//   		ProfileArn: jsii.String("profileArn"),
//   		ProfileId: jsii.String("profileId"),
//   		Province: jsii.String("province"),
//   		ShippingAddress1: jsii.String("shippingAddress1"),
//   		ShippingAddress2: jsii.String("shippingAddress2"),
//   		ShippingAddress3: jsii.String("shippingAddress3"),
//   		ShippingAddress4: jsii.String("shippingAddress4"),
//   		ShippingCity: jsii.String("shippingCity"),
//   		ShippingCountry: jsii.String("shippingCountry"),
//   		ShippingCounty: jsii.String("shippingCounty"),
//   		ShippingPostalCode: jsii.String("shippingPostalCode"),
//   		ShippingProvince: jsii.String("shippingProvince"),
//   		ShippingState: jsii.String("shippingState"),
//   		State: jsii.String("state"),
//   	},
//   	SystemAttributes: &SystemAttributesProperty{
//   		CustomerEndpoint: &SystemEndpointAttributesProperty{
//   			Address: jsii.String("address"),
//   		},
//   		Name: jsii.String("name"),
//   		SystemEndpoint: &SystemEndpointAttributesProperty{
//   			Address: jsii.String("address"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattributes.html
//
type CfnMessageTemplate_MessageTemplateAttributesProperty struct {
	// The agent attributes that are used with the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattributes.html#cfn-wisdom-messagetemplate-messagetemplateattributes-agentattributes
	//
	AgentAttributes interface{} `field:"optional" json:"agentAttributes" yaml:"agentAttributes"`
	// The custom attributes that are used with the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattributes.html#cfn-wisdom-messagetemplate-messagetemplateattributes-customattributes
	//
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The customer profile attributes that are used with the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattributes.html#cfn-wisdom-messagetemplate-messagetemplateattributes-customerprofileattributes
	//
	CustomerProfileAttributes interface{} `field:"optional" json:"customerProfileAttributes" yaml:"customerProfileAttributes"`
	// The system attributes that are used with the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattributes.html#cfn-wisdom-messagetemplate-messagetemplateattributes-systemattributes
	//
	SystemAttributes interface{} `field:"optional" json:"systemAttributes" yaml:"systemAttributes"`
}

