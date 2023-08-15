package awscognito


// This interface contains standard attributes recognized by Cognito from https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html including built-in attributes `email_verified` and `phone_number_verified`.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   clientWriteAttributes := (cognito.NewClientAttributes()).WithStandardAttributes(&StandardAttributesMask{
//   	Fullname: jsii.Boolean(true),
//   	Email: jsii.Boolean(true),
//   }).WithCustomAttributes(jsii.String("favouritePizza"), jsii.String("favouriteBeverage"))
//
//   clientReadAttributes := clientWriteAttributes.WithStandardAttributes(&StandardAttributesMask{
//   	EmailVerified: jsii.Boolean(true),
//   }).WithCustomAttributes(jsii.String("pointsEarned"))
//
//   pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	// ...
//   	ReadAttributes: clientReadAttributes,
//   	WriteAttributes: clientWriteAttributes,
//   })
//
type StandardAttributesMask struct {
	// The user's postal address.
	// Default: false.
	//
	Address *bool `field:"optional" json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	// Default: false.
	//
	Birthdate *bool `field:"optional" json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	// Default: false.
	//
	Email *bool `field:"optional" json:"email" yaml:"email"`
	// Whether the email address has been verified.
	// Default: false.
	//
	EmailVerified *bool `field:"optional" json:"emailVerified" yaml:"emailVerified"`
	// The surname or last name of the user.
	// Default: false.
	//
	FamilyName *bool `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	// Default: false.
	//
	Fullname *bool `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Default: false.
	//
	Gender *bool `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Default: false.
	//
	GivenName *bool `field:"optional" json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	// Default: false.
	//
	LastUpdateTime *bool `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	// Default: false.
	//
	Locale *bool `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	// Default: false.
	//
	MiddleName *bool `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Default: false.
	//
	Nickname *bool `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Default: false.
	//
	PhoneNumber *bool `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Whether the phone number has been verified.
	// Default: false.
	//
	PhoneNumberVerified *bool `field:"optional" json:"phoneNumberVerified" yaml:"phoneNumberVerified"`
	// The user's preffered username, different from the immutable user name.
	// Default: false.
	//
	PreferredUsername *bool `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Default: false.
	//
	ProfilePage *bool `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Default: false.
	//
	ProfilePicture *bool `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Default: false.
	//
	Timezone *bool `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Default: false.
	//
	Website *bool `field:"optional" json:"website" yaml:"website"`
}

