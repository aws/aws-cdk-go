package awssso


// The configuration of a trusted token issuer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trustedTokenIssuerConfigurationProperty := &TrustedTokenIssuerConfigurationProperty{
//   	OidcJwtConfiguration: &OidcJwtConfigurationProperty{
//   		ClaimAttributePath: jsii.String("claimAttributePath"),
//   		IdentityStoreAttributePath: jsii.String("identityStoreAttributePath"),
//   		IssuerUrl: jsii.String("issuerUrl"),
//   		JwksRetrievalOption: jsii.String("jwksRetrievalOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-trustedtokenissuerconfiguration.html
//
type CfnTrustedTokenIssuerPropsMixin_TrustedTokenIssuerConfigurationProperty struct {
	// The configuration settings for a trusted token issuer that uses OIDC with JWTs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-trustedtokenissuerconfiguration.html#cfn-sso-trustedtokenissuer-trustedtokenissuerconfiguration-oidcjwtconfiguration
	//
	OidcJwtConfiguration interface{} `field:"optional" json:"oidcJwtConfiguration" yaml:"oidcJwtConfiguration"`
}

