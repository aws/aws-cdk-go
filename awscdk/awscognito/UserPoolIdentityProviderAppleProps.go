package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties to initialize UserPoolAppleIdentityProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var secretValue secretValue
//   var userPool userPool
//
//   userPoolIdentityProviderAppleProps := &UserPoolIdentityProviderAppleProps{
//   	ClientId: jsii.String("clientId"),
//   	KeyId: jsii.String("keyId"),
//   	TeamId: jsii.String("teamId"),
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
//   		EmailVerified: providerAttribute,
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
//   	PrivateKey: jsii.String("privateKey"),
//   	PrivateKeyValue: secretValue,
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
type UserPoolIdentityProviderAppleProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Apple APIs.
	// See: https://developer.apple.com/documentation/sign_in_with_apple/clientconfigi/3230948-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The keyId (of the same key, which content has to be later supplied as `privateKey`) for Apple APIs to authenticate the client.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The teamId for Apple APIs to authenticate the client.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The privateKey content for Apple APIs to authenticate the client.
	// Default: none.
	//
	// Deprecated: use privateKeyValue.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The privateKey content for Apple APIs to authenticate the client.
	// Default: none.
	//
	PrivateKeyValue awscdk.SecretValue `field:"optional" json:"privateKeyValue" yaml:"privateKeyValue"`
	// The list of apple permissions to obtain for getting access to the apple profile.
	// See: https://developer.apple.com/documentation/sign_in_with_apple/clientconfigi/3230955-scope
	//
	// Default: [ name ].
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

