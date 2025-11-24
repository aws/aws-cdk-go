package mixinsawswisdom


// The customer profile attributes that are used with the message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerProfileAttributesProperty := &CustomerProfileAttributesProperty{
//   	AccountNumber: jsii.String("accountNumber"),
//   	AdditionalInformation: jsii.String("additionalInformation"),
//   	Address1: jsii.String("address1"),
//   	Address2: jsii.String("address2"),
//   	Address3: jsii.String("address3"),
//   	Address4: jsii.String("address4"),
//   	BillingAddress1: jsii.String("billingAddress1"),
//   	BillingAddress2: jsii.String("billingAddress2"),
//   	BillingAddress3: jsii.String("billingAddress3"),
//   	BillingAddress4: jsii.String("billingAddress4"),
//   	BillingCity: jsii.String("billingCity"),
//   	BillingCountry: jsii.String("billingCountry"),
//   	BillingCounty: jsii.String("billingCounty"),
//   	BillingPostalCode: jsii.String("billingPostalCode"),
//   	BillingProvince: jsii.String("billingProvince"),
//   	BillingState: jsii.String("billingState"),
//   	BirthDate: jsii.String("birthDate"),
//   	BusinessEmailAddress: jsii.String("businessEmailAddress"),
//   	BusinessName: jsii.String("businessName"),
//   	BusinessPhoneNumber: jsii.String("businessPhoneNumber"),
//   	City: jsii.String("city"),
//   	Country: jsii.String("country"),
//   	County: jsii.String("county"),
//   	Custom: map[string]*string{
//   		"customKey": jsii.String("custom"),
//   	},
//   	EmailAddress: jsii.String("emailAddress"),
//   	FirstName: jsii.String("firstName"),
//   	Gender: jsii.String("gender"),
//   	HomePhoneNumber: jsii.String("homePhoneNumber"),
//   	LastName: jsii.String("lastName"),
//   	MailingAddress1: jsii.String("mailingAddress1"),
//   	MailingAddress2: jsii.String("mailingAddress2"),
//   	MailingAddress3: jsii.String("mailingAddress3"),
//   	MailingAddress4: jsii.String("mailingAddress4"),
//   	MailingCity: jsii.String("mailingCity"),
//   	MailingCountry: jsii.String("mailingCountry"),
//   	MailingCounty: jsii.String("mailingCounty"),
//   	MailingPostalCode: jsii.String("mailingPostalCode"),
//   	MailingProvince: jsii.String("mailingProvince"),
//   	MailingState: jsii.String("mailingState"),
//   	MiddleName: jsii.String("middleName"),
//   	MobilePhoneNumber: jsii.String("mobilePhoneNumber"),
//   	PartyType: jsii.String("partyType"),
//   	PhoneNumber: jsii.String("phoneNumber"),
//   	PostalCode: jsii.String("postalCode"),
//   	ProfileArn: jsii.String("profileArn"),
//   	ProfileId: jsii.String("profileId"),
//   	Province: jsii.String("province"),
//   	ShippingAddress1: jsii.String("shippingAddress1"),
//   	ShippingAddress2: jsii.String("shippingAddress2"),
//   	ShippingAddress3: jsii.String("shippingAddress3"),
//   	ShippingAddress4: jsii.String("shippingAddress4"),
//   	ShippingCity: jsii.String("shippingCity"),
//   	ShippingCountry: jsii.String("shippingCountry"),
//   	ShippingCounty: jsii.String("shippingCounty"),
//   	ShippingPostalCode: jsii.String("shippingPostalCode"),
//   	ShippingProvince: jsii.String("shippingProvince"),
//   	ShippingState: jsii.String("shippingState"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html
//
type CfnMessageTemplatePropsMixin_CustomerProfileAttributesProperty struct {
	// A unique account number that you have given to the customer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-accountnumber
	//
	AccountNumber *string `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// Any additional information relevant to the customer's profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-additionalinformation
	//
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// The first line of a customer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-address1
	//
	Address1 *string `field:"optional" json:"address1" yaml:"address1"`
	// The second line of a customer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-address2
	//
	Address2 *string `field:"optional" json:"address2" yaml:"address2"`
	// The third line of a customer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-address3
	//
	Address3 *string `field:"optional" json:"address3" yaml:"address3"`
	// The fourth line of a customer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-address4
	//
	Address4 *string `field:"optional" json:"address4" yaml:"address4"`
	// The first line of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingaddress1
	//
	BillingAddress1 *string `field:"optional" json:"billingAddress1" yaml:"billingAddress1"`
	// The second line of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingaddress2
	//
	BillingAddress2 *string `field:"optional" json:"billingAddress2" yaml:"billingAddress2"`
	// The third line of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingaddress3
	//
	BillingAddress3 *string `field:"optional" json:"billingAddress3" yaml:"billingAddress3"`
	// The fourth line of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingaddress4
	//
	BillingAddress4 *string `field:"optional" json:"billingAddress4" yaml:"billingAddress4"`
	// The city of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingcity
	//
	BillingCity *string `field:"optional" json:"billingCity" yaml:"billingCity"`
	// The country of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingcountry
	//
	BillingCountry *string `field:"optional" json:"billingCountry" yaml:"billingCountry"`
	// The county of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingcounty
	//
	BillingCounty *string `field:"optional" json:"billingCounty" yaml:"billingCounty"`
	// The postal code of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingpostalcode
	//
	BillingPostalCode *string `field:"optional" json:"billingPostalCode" yaml:"billingPostalCode"`
	// The province of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingprovince
	//
	BillingProvince *string `field:"optional" json:"billingProvince" yaml:"billingProvince"`
	// The state of a customer’s billing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-billingstate
	//
	BillingState *string `field:"optional" json:"billingState" yaml:"billingState"`
	// The customer's birth date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-birthdate
	//
	BirthDate *string `field:"optional" json:"birthDate" yaml:"birthDate"`
	// The customer's business email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-businessemailaddress
	//
	BusinessEmailAddress *string `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// The name of the customer's business.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-businessname
	//
	BusinessName *string `field:"optional" json:"businessName" yaml:"businessName"`
	// The customer's business phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-businessphonenumber
	//
	BusinessPhoneNumber *string `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// The city in which a customer lives.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-city
	//
	City *string `field:"optional" json:"city" yaml:"city"`
	// The country in which a customer lives.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-country
	//
	Country *string `field:"optional" json:"country" yaml:"country"`
	// The county in which a customer lives.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-county
	//
	County *string `field:"optional" json:"county" yaml:"county"`
	// The custom attributes in customer profile attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-custom
	//
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// The customer's email address, which has not been specified as a personal or business address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-emailaddress
	//
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The customer's first name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The customer's gender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-gender
	//
	Gender *string `field:"optional" json:"gender" yaml:"gender"`
	// The customer's mobile phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-homephonenumber
	//
	HomePhoneNumber *string `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// The customer's last name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The first line of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingaddress1
	//
	MailingAddress1 *string `field:"optional" json:"mailingAddress1" yaml:"mailingAddress1"`
	// The second line of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingaddress2
	//
	MailingAddress2 *string `field:"optional" json:"mailingAddress2" yaml:"mailingAddress2"`
	// The third line of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingaddress3
	//
	MailingAddress3 *string `field:"optional" json:"mailingAddress3" yaml:"mailingAddress3"`
	// The fourth line of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingaddress4
	//
	MailingAddress4 *string `field:"optional" json:"mailingAddress4" yaml:"mailingAddress4"`
	// The city of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingcity
	//
	MailingCity *string `field:"optional" json:"mailingCity" yaml:"mailingCity"`
	// The country of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingcountry
	//
	MailingCountry *string `field:"optional" json:"mailingCountry" yaml:"mailingCountry"`
	// The county of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingcounty
	//
	MailingCounty *string `field:"optional" json:"mailingCounty" yaml:"mailingCounty"`
	// The postal code of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingpostalcode
	//
	MailingPostalCode *string `field:"optional" json:"mailingPostalCode" yaml:"mailingPostalCode"`
	// The province of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingprovince
	//
	MailingProvince *string `field:"optional" json:"mailingProvince" yaml:"mailingProvince"`
	// The state of a customer’s mailing address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mailingstate
	//
	MailingState *string `field:"optional" json:"mailingState" yaml:"mailingState"`
	// The customer's middle name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-middlename
	//
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
	// The customer's mobile phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-mobilephonenumber
	//
	MobilePhoneNumber *string `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// The customer's party type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-partytype
	//
	PartyType *string `field:"optional" json:"partyType" yaml:"partyType"`
	// The customer's phone number, which has not been specified as a mobile, home, or business number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-phonenumber
	//
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The postal code of a customer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-postalcode
	//
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The ARN of a customer profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-profilearn
	//
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
	// The unique identifier of a customer profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-profileid
	//
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// The province in which a customer lives.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-province
	//
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The first line of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingaddress1
	//
	ShippingAddress1 *string `field:"optional" json:"shippingAddress1" yaml:"shippingAddress1"`
	// The second line of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingaddress2
	//
	ShippingAddress2 *string `field:"optional" json:"shippingAddress2" yaml:"shippingAddress2"`
	// The third line of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingaddress3
	//
	ShippingAddress3 *string `field:"optional" json:"shippingAddress3" yaml:"shippingAddress3"`
	// The fourth line of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingaddress4
	//
	ShippingAddress4 *string `field:"optional" json:"shippingAddress4" yaml:"shippingAddress4"`
	// The city of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingcity
	//
	ShippingCity *string `field:"optional" json:"shippingCity" yaml:"shippingCity"`
	// The country of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingcountry
	//
	ShippingCountry *string `field:"optional" json:"shippingCountry" yaml:"shippingCountry"`
	// The county of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingcounty
	//
	ShippingCounty *string `field:"optional" json:"shippingCounty" yaml:"shippingCounty"`
	// The postal code of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingpostalcode
	//
	ShippingPostalCode *string `field:"optional" json:"shippingPostalCode" yaml:"shippingPostalCode"`
	// The province of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingprovince
	//
	ShippingProvince *string `field:"optional" json:"shippingProvince" yaml:"shippingProvince"`
	// The state of a customer’s shipping address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-shippingstate
	//
	ShippingState *string `field:"optional" json:"shippingState" yaml:"shippingState"`
	// The state in which a customer lives.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-customerprofileattributes.html#cfn-wisdom-messagetemplate-customerprofileattributes-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

