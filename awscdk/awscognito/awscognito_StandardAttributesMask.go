package awscognito


// This interface contains standard attributes recognized by Cognito from https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html including built-in attributes `email_verified` and `phone_number_verified`.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   clientWriteAttributes := (cognito.NewClientAttributes()).withStandardAttributes(&standardAttributesMask{
//   	fullname: jsii.Boolean(true),
//   	email: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("favouritePizza"), jsii.String("favouriteBeverage"))
//
//   clientReadAttributes := clientWriteAttributes.withStandardAttributes(&standardAttributesMask{
//   	emailVerified: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("pointsEarned"))
//
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	readAttributes: clientReadAttributes,
//   	writeAttributes: clientWriteAttributes,
//   })
//
type StandardAttributesMask struct {
	// The user's postal address.
	Address *bool `field:"optional" json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	Birthdate *bool `field:"optional" json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether the email address has been verified.
	EmailVerified *bool `field:"optional" json:"emailVerified" yaml:"emailVerified"`
	// The surname or last name of the user.
	FamilyName *bool `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	Fullname *bool `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	Gender *bool `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	GivenName *bool `field:"optional" json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	LastUpdateTime *bool `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	Locale *bool `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	MiddleName *bool `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	Nickname *bool `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	PhoneNumber *bool `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Whether the phone number has been verified.
	PhoneNumberVerified *bool `field:"optional" json:"phoneNumberVerified" yaml:"phoneNumberVerified"`
	// The user's preffered username, different from the immutable user name.
	PreferredUsername *bool `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	ProfilePage *bool `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	ProfilePicture *bool `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	Timezone *bool `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	Website *bool `field:"optional" json:"website" yaml:"website"`
}

