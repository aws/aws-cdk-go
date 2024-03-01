package awscognito


// Properties to initialize UserPoolFacebookIdentityProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//
//   userPoolIdentityProviderFacebookProps := &UserPoolIdentityProviderFacebookProps{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
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
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
type UserPoolIdentityProviderFacebookProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Facebook APIs.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for Facebook to authenticate the client.
	// See: https://developers.facebook.com/docs/facebook-login/security#appsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The Facebook API version to use.
	// Default: - to the oldest version supported by Facebook.
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The list of Facebook permissions to obtain for getting access to the Facebook profile.
	// See: https://developers.facebook.com/docs/facebook-login/permissions
	//
	// Default: [ public_profile ].
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

