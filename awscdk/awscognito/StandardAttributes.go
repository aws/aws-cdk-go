package awscognito


// The set of standard attributes that can be marked as required or mutable.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	StandardAttributes: &StandardAttributes{
//   		Fullname: &StandardAttribute{
//   			Required: jsii.Boolean(true),
//   			Mutable: jsii.Boolean(false),
//   		},
//   		Address: &StandardAttribute{
//   			Required: jsii.Boolean(false),
//   			Mutable: jsii.Boolean(true),
//   		},
//   	},
//   	CustomAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes
//
// Experimental.
type StandardAttributes struct {
	// The user's postal address.
	// Experimental.
	Address *StandardAttribute `field:"optional" json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	// Experimental.
	Birthdate *StandardAttribute `field:"optional" json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	// Experimental.
	Email *StandardAttribute `field:"optional" json:"email" yaml:"email"`
	// DEPRECATED.
	// Deprecated: this is not a standard attribute and was incorrectly added to the CDK.
	// It is a Cognito built-in attribute and cannot be controlled as part of user pool creation.
	EmailVerified *StandardAttribute `field:"optional" json:"emailVerified" yaml:"emailVerified"`
	// The surname or last name of the user.
	// Experimental.
	FamilyName *StandardAttribute `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	// Experimental.
	Fullname *StandardAttribute `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Experimental.
	Gender *StandardAttribute `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Experimental.
	GivenName *StandardAttribute `field:"optional" json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	// Experimental.
	LastUpdateTime *StandardAttribute `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	// Experimental.
	Locale *StandardAttribute `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	// Experimental.
	MiddleName *StandardAttribute `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Experimental.
	Nickname *StandardAttribute `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Experimental.
	PhoneNumber *StandardAttribute `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// DEPRECATED.
	// Deprecated: this is not a standard attribute and was incorrectly added to the CDK.
	// It is a Cognito built-in attribute and cannot be controlled as part of user pool creation.
	PhoneNumberVerified *StandardAttribute `field:"optional" json:"phoneNumberVerified" yaml:"phoneNumberVerified"`
	// The user's preffered username, different from the immutable user name.
	// Experimental.
	PreferredUsername *StandardAttribute `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Experimental.
	ProfilePage *StandardAttribute `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Experimental.
	ProfilePicture *StandardAttribute `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Experimental.
	Timezone *StandardAttribute `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Experimental.
	Website *StandardAttribute `field:"optional" json:"website" yaml:"website"`
}

