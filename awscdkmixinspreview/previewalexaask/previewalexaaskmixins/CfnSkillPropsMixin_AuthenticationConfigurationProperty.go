package previewalexaaskmixins


// The `AuthenticationConfiguration` property type specifies the Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
//
// Only Login with Amazon security profiles created through the  are supported for authentication. A client ID, client secret, and refresh token are required. You can generate a client ID and client secret by creating a new  on the Amazon Developer Portal or you can retrieve them from an existing profile. You can then retrieve the refresh token using the Alexa Skills Kit CLI. For instructions, see  in the  .
//
// `AuthenticationConfiguration` is a property of the `Alexa::ASK::Skill` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationConfigurationProperty := &AuthenticationConfigurationProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html
//
type CfnSkillPropsMixin_AuthenticationConfigurationProperty struct {
	// Client ID from Login with Amazon (LWA).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret from Login with Amazon (LWA).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Refresh token from Login with Amazon (LWA).
	//
	// This token is secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-refreshtoken
	//
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

