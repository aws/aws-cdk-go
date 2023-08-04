package awscognito


// Properties to create a new instance of UserPoolIdentityProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//
//   userPoolIdentityProviderProps := &UserPoolIdentityProviderProps{
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	AttributeMapping: &AttributeMapping{
//   		Address: providerAttribute,
//   		Birthdate: providerAttribute,
//   		Custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		Email: providerAttribute,
//   		FamilyName: providerAttribute,
//   		Fullname: providerAttribute,
//   		Gender: providerAttribute,
//   		GivenName: providerAttribute,
//   		LastUpdateTime: providerAttribute,
//   		Locale: providerAttribute,
//   		MiddleName: providerAttribute,
//   		Nickname: providerAttribute,
//   		PhoneNumber: providerAttribute,
//   		PreferredUsername: providerAttribute,
//   		ProfilePage: providerAttribute,
//   		ProfilePicture: providerAttribute,
//   		Timezone: providerAttribute,
//   		Website: providerAttribute,
//   	},
//   }
//
type UserPoolIdentityProviderProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
}

