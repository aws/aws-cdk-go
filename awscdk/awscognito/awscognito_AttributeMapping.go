package awscognito


// The mapping of user pool attributes to the attributes provided by the identity providers.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   	userPool: userpool,
//   	attributeMapping: &attributeMapping{
//   		email: cognito.providerAttribute_AMAZON_EMAIL(),
//   		website: cognito.*providerAttribute.other(jsii.String("url")),
//   		 // use other() when an attribute is not pre-defined in the CDK
//   		custom: map[string]*providerAttribute{
//   			// custom user pool attributes go here
//   			"uniqueId": cognito.*providerAttribute_AMAZON_USER_ID(),
//   		},
//   	},
//   })
//
// Experimental.
type AttributeMapping struct {
	// The user's postal address is a required attribute.
	// Experimental.
	Address ProviderAttribute `field:"optional" json:"address" yaml:"address"`
	// The user's birthday.
	// Experimental.
	Birthdate ProviderAttribute `field:"optional" json:"birthdate" yaml:"birthdate"`
	// Specify custom attribute mapping here and mapping for any standard attributes not supported yet.
	// Experimental.
	Custom *map[string]ProviderAttribute `field:"optional" json:"custom" yaml:"custom"`
	// The user's e-mail address.
	// Experimental.
	Email ProviderAttribute `field:"optional" json:"email" yaml:"email"`
	// The surname or last name of user.
	// Experimental.
	FamilyName ProviderAttribute `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form.
	// Experimental.
	Fullname ProviderAttribute `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Experimental.
	Gender ProviderAttribute `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Experimental.
	GivenName ProviderAttribute `field:"optional" json:"givenName" yaml:"givenName"`
	// Time, the user's information was last updated.
	// Experimental.
	LastUpdateTime ProviderAttribute `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale.
	// Experimental.
	Locale ProviderAttribute `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	// Experimental.
	MiddleName ProviderAttribute `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Experimental.
	Nickname ProviderAttribute `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Experimental.
	PhoneNumber ProviderAttribute `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The user's preferred username.
	// Experimental.
	PreferredUsername ProviderAttribute `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Experimental.
	ProfilePage ProviderAttribute `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Experimental.
	ProfilePicture ProviderAttribute `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Experimental.
	Timezone ProviderAttribute `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Experimental.
	Website ProviderAttribute `field:"optional" json:"website" yaml:"website"`
}

