package awscognito


// The mapping of user pool attributes to the attributes provided by the identity providers.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &UserPoolIdentityProviderAmazonProps{
//   	ClientId: jsii.String("amzn-client-id"),
//   	ClientSecret: jsii.String("amzn-client-secret"),
//   	UserPool: userpool,
//   	AttributeMapping: &AttributeMapping{
//   		Email: cognito.ProviderAttribute_AMAZON_EMAIL(),
//   		Website: cognito.ProviderAttribute_Other(jsii.String("url")),
//   		 // use other() when an attribute is not pre-defined in the CDK
//   		Custom: map[string]providerAttribute{
//   			// custom user pool attributes go here
//   			"uniqueId": cognito.*providerAttribute_AMAZON_USER_ID(),
//   		},
//   	},
//   })
//
type AttributeMapping struct {
	// The user's postal address is a required attribute.
	// Default: - not mapped.
	//
	Address ProviderAttribute `field:"optional" json:"address" yaml:"address"`
	// The user's birthday.
	// Default: - not mapped.
	//
	Birthdate ProviderAttribute `field:"optional" json:"birthdate" yaml:"birthdate"`
	// Specify custom attribute mapping here and mapping for any standard attributes not supported yet.
	// Default: - no custom attribute mapping.
	//
	Custom *map[string]ProviderAttribute `field:"optional" json:"custom" yaml:"custom"`
	// The user's e-mail address.
	// Default: - not mapped.
	//
	Email ProviderAttribute `field:"optional" json:"email" yaml:"email"`
	// The surname or last name of user.
	// Default: - not mapped.
	//
	FamilyName ProviderAttribute `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form.
	// Default: - not mapped.
	//
	Fullname ProviderAttribute `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Default: - not mapped.
	//
	Gender ProviderAttribute `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Default: - not mapped.
	//
	GivenName ProviderAttribute `field:"optional" json:"givenName" yaml:"givenName"`
	// Time, the user's information was last updated.
	// Default: - not mapped.
	//
	LastUpdateTime ProviderAttribute `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale.
	// Default: - not mapped.
	//
	Locale ProviderAttribute `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	// Default: - not mapped.
	//
	MiddleName ProviderAttribute `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Default: - not mapped.
	//
	Nickname ProviderAttribute `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Default: - not mapped.
	//
	PhoneNumber ProviderAttribute `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The user's preferred username.
	// Default: - not mapped.
	//
	PreferredUsername ProviderAttribute `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Default: - not mapped.
	//
	ProfilePage ProviderAttribute `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Default: - not mapped.
	//
	ProfilePicture ProviderAttribute `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Default: - not mapped.
	//
	Timezone ProviderAttribute `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Default: - not mapped.
	//
	Website ProviderAttribute `field:"optional" json:"website" yaml:"website"`
}

