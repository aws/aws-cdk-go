package awscognito


// The set of standard attributes that can be marked as required or mutable.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
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
	Address *StandardAttribute `field:"optional" json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	Birthdate *StandardAttribute `field:"optional" json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	Email *StandardAttribute `field:"optional" json:"email" yaml:"email"`
	// The surname or last name of the user.
	FamilyName *StandardAttribute `field:"optional" json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	Fullname *StandardAttribute `field:"optional" json:"fullname" yaml:"fullname"`
	// The user's gender.
	Gender *StandardAttribute `field:"optional" json:"gender" yaml:"gender"`
	// The user's first name or give name.
	GivenName *StandardAttribute `field:"optional" json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	LastUpdateTime *StandardAttribute `field:"optional" json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	Locale *StandardAttribute `field:"optional" json:"locale" yaml:"locale"`
	// The user's middle name.
	MiddleName *StandardAttribute `field:"optional" json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	Nickname *StandardAttribute `field:"optional" json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	PhoneNumber *StandardAttribute `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The user's preffered username, different from the immutable user name.
	PreferredUsername *StandardAttribute `field:"optional" json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	ProfilePage *StandardAttribute `field:"optional" json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	ProfilePicture *StandardAttribute `field:"optional" json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	Timezone *StandardAttribute `field:"optional" json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	Website *StandardAttribute `field:"optional" json:"website" yaml:"website"`
}

