package mixinsawsqbusiness


// Provides information about the identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   identityProviderConfigurationProperty := &IdentityProviderConfigurationProperty{
//   	OpenIdConnectConfiguration: &OpenIDConnectProviderConfigurationProperty{
//   		SecretsArn: jsii.String("secretsArn"),
//   		SecretsRole: jsii.String("secretsRole"),
//   	},
//   	SamlConfiguration: &SamlProviderConfigurationProperty{
//   		AuthenticationUrl: jsii.String("authenticationUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-identityproviderconfiguration.html
//
type CfnWebExperiencePropsMixin_IdentityProviderConfigurationProperty struct {
	// The OIDC-compliant identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-identityproviderconfiguration.html#cfn-qbusiness-webexperience-identityproviderconfiguration-openidconnectconfiguration
	//
	OpenIdConnectConfiguration interface{} `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
	// The SAML 2.0-compliant identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-identityproviderconfiguration.html#cfn-qbusiness-webexperience-identityproviderconfiguration-samlconfiguration
	//
	SamlConfiguration interface{} `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
}

