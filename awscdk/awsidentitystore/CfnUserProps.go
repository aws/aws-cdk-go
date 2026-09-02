package awsidentitystore


// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProps := &CfnUserProps{
//   	IdentityStoreId: jsii.String("identityStoreId"),
//
//   	// the properties below are optional
//   	Addresses: []interface{}{
//   		&AddressesItemsProperty{
//   			Country: jsii.String("country"),
//   			Formatted: jsii.String("formatted"),
//   			Locality: jsii.String("locality"),
//   			PostalCode: jsii.String("postalCode"),
//   			Primary: jsii.Boolean(false),
//   			Region: jsii.String("region"),
//   			StreetAddress: jsii.String("streetAddress"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Birthdate: jsii.String("birthdate"),
//   	DisplayName: jsii.String("displayName"),
//   	Emails: []interface{}{
//   		&EmailsItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Locale: jsii.String("locale"),
//   	Name: &NameProperty{
//   		FamilyName: jsii.String("familyName"),
//   		Formatted: jsii.String("formatted"),
//   		GivenName: jsii.String("givenName"),
//   		HonorificPrefix: jsii.String("honorificPrefix"),
//   		HonorificSuffix: jsii.String("honorificSuffix"),
//   		MiddleName: jsii.String("middleName"),
//   	},
//   	NickName: jsii.String("nickName"),
//   	PhoneNumbers: []interface{}{
//   		&PhoneNumbersItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Photos: []interface{}{
//   		&PhotosItemsProperty{
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Display: jsii.String("display"),
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	PreferredLanguage: jsii.String("preferredLanguage"),
//   	ProfileUrl: jsii.String("profileUrl"),
//   	Roles: []interface{}{
//   		&RolesItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timezone: jsii.String("timezone"),
//   	Title: jsii.String("title"),
//   	UserName: jsii.String("userName"),
//   	UserType: jsii.String("userType"),
//   	Website: jsii.String("website"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html
//
type CfnUserProps struct {
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-identitystoreid
	//
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// A list of addresses associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-addresses
	//
	Addresses interface{} `field:"optional" json:"addresses" yaml:"addresses"`
	// The user's birthdate in YYYY-MM-DD format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-birthdate
	//
	Birthdate *string `field:"optional" json:"birthdate" yaml:"birthdate"`
	// A string containing the name of the user for display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of email addresses associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-emails
	//
	Emails interface{} `field:"optional" json:"emails" yaml:"emails"`
	// The geographical region or location of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// The name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-name
	//
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// An alternate name for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-nickname
	//
	NickName *string `field:"optional" json:"nickName" yaml:"nickName"`
	// A list of phone numbers associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-phonenumbers
	//
	PhoneNumbers interface{} `field:"optional" json:"phoneNumbers" yaml:"phoneNumbers"`
	// A list of photos associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-photos
	//
	Photos interface{} `field:"optional" json:"photos" yaml:"photos"`
	// The preferred language of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-preferredlanguage
	//
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// A URL associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-profileurl
	//
	ProfileUrl *string `field:"optional" json:"profileUrl" yaml:"profileUrl"`
	// A list of roles associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-roles
	//
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The time zone for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// The title of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A unique string used to identify the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
	// A string indicating the type of user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-usertype
	//
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
	// The user's personal website or blog URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html#cfn-identitystore-user-website
	//
	Website *string `field:"optional" json:"website" yaml:"website"`
}

