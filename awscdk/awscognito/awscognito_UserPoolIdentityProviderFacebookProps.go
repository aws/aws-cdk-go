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
//   userPoolIdentityProviderFacebookProps := &userPoolIdentityProviderFacebookProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
type UserPoolIdentityProviderFacebookProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Facebook APIs.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientUd for Facebook to authenticate the client.
	// See: https://developers.facebook.com/docs/facebook-login/security#appsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The Facebook API version to use.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The list of facebook permissions to obtain for getting access to the Facebook profile.
	// See: https://developers.facebook.com/docs/facebook-login/permissions
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

