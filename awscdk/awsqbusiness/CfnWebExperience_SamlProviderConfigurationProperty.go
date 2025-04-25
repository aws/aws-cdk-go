package awsqbusiness


// Information about the SAML 2.0-compliant identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samlProviderConfigurationProperty := &SamlProviderConfigurationProperty{
//   	AuthenticationUrl: jsii.String("authenticationUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-samlproviderconfiguration.html
//
type CfnWebExperience_SamlProviderConfigurationProperty struct {
	// The URL where Amazon Q Business end users will be redirected for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-samlproviderconfiguration.html#cfn-qbusiness-webexperience-samlproviderconfiguration-authenticationurl
	//
	AuthenticationUrl *string `field:"required" json:"authenticationUrl" yaml:"authenticationUrl"`
}

