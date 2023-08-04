package awscognito


// Properties to initialize UserPoolIdentityProviderSaml.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   var userPoolIdentityProviderSamlMetadata userPoolIdentityProviderSamlMetadata
//
//   userPoolIdentityProviderSamlProps := &UserPoolIdentityProviderSamlProps{
//   	Metadata: userPoolIdentityProviderSamlMetadata,
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
//   	Identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	IdpSignout: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
type UserPoolIdentityProviderSamlProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The SAML metadata.
	Metadata UserPoolIdentityProviderSamlMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Identifiers.
	//
	// Identifiers can be used to redirect users to the correct IdP in multitenant apps.
	// Default: - no identifiers used.
	//
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Whether to enable the "Sign-out flow" feature.
	// Default: - false.
	//
	IdpSignout *bool `field:"optional" json:"idpSignout" yaml:"idpSignout"`
	// The name of the provider.
	//
	// Must be between 3 and 32 characters.
	// Default: - the unique ID of the construct.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

