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
type StandardAttributes struct {
	// The user's postal address.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Address *StandardAttribute `field:"optional" json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Birthdate *StandardAttribute `field:"optional" json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Email *StandardAttribute `field:"optional" json:"email" yaml:"email"`
	// The surname or last name of the user.
	// Default: - see the defaults under `StandardAttribute`.
	//
	FamilyName *StandardAttribute `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Fullname *StandardAttribute `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Gender *StandardAttribute `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Default: - see the defaults under `StandardAttribute`.
	//
	GivenName *StandardAttribute `field:"optional" json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	// Default: - see the defaults under `StandardAttribute`.
	//
	LastUpdateTime *StandardAttribute `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Locale *StandardAttribute `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	// Default: - see the defaults under `StandardAttribute`.
	//
	MiddleName *StandardAttribute `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Nickname *StandardAttribute `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Default: - see the defaults under `StandardAttribute`.
	//
	PhoneNumber *StandardAttribute `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The user's preffered username, different from the immutable user name.
	// Default: - see the defaults under `StandardAttribute`.
	//
	PreferredUsername *StandardAttribute `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Default: - see the defaults under `StandardAttribute`.
	//
	ProfilePage *StandardAttribute `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Default: - see the defaults under `StandardAttribute`.
	//
	ProfilePicture *StandardAttribute `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Timezone *StandardAttribute `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Default: - see the defaults under `StandardAttribute`.
	//
	Website *StandardAttribute `field:"optional" json:"website" yaml:"website"`
}

