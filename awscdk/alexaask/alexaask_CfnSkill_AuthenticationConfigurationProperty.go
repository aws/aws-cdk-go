package alexaask


// The `AuthenticationConfiguration` property type specifies the Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
//
// Only Login with Amazon security profiles created through the  are supported for authentication. A client ID, client secret, and refresh token are required. You can generate a client ID and client secret by creating a new  on the Amazon Developer Portal or you can retrieve them from an existing profile. You can then retrieve the refresh token using the Alexa Skills Kit CLI. For instructions, see  in the  .
//
// `AuthenticationConfiguration` is a property of the `Alexa::ASK::Skill` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationProperty := &authenticationConfigurationProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnSkill_AuthenticationConfigurationProperty struct {
	// Client ID from Login with Amazon (LWA).
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Client secret from Login with Amazon (LWA).
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Refresh token from Login with Amazon (LWA).
	//
	// This token is secret.
	RefreshToken *string `field:"required" json:"refreshToken" yaml:"refreshToken"`
}

