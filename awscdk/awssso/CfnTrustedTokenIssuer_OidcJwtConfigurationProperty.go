package awssso


// The configuration settings for a trusted token issuer that uses OIDC with JWTs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcJwtConfigurationProperty := &OidcJwtConfigurationProperty{
//   	ClaimAttributePath: jsii.String("claimAttributePath"),
//   	IdentityStoreAttributePath: jsii.String("identityStoreAttributePath"),
//   	IssuerUrl: jsii.String("issuerUrl"),
//   	JwksRetrievalOption: jsii.String("jwksRetrievalOption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-oidcjwtconfiguration.html
//
type CfnTrustedTokenIssuer_OidcJwtConfigurationProperty struct {
	// The path of the source attribute in the JWT from the trusted token issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-oidcjwtconfiguration.html#cfn-sso-trustedtokenissuer-oidcjwtconfiguration-claimattributepath
	//
	ClaimAttributePath *string `field:"required" json:"claimAttributePath" yaml:"claimAttributePath"`
	// The path of the destination attribute in a JWT from IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-oidcjwtconfiguration.html#cfn-sso-trustedtokenissuer-oidcjwtconfiguration-identitystoreattributepath
	//
	IdentityStoreAttributePath *string `field:"required" json:"identityStoreAttributePath" yaml:"identityStoreAttributePath"`
	// The URL that IAM Identity Center uses for OpenID Discovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-oidcjwtconfiguration.html#cfn-sso-trustedtokenissuer-oidcjwtconfiguration-issuerurl
	//
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The method that the trusted token issuer can use to retrieve the JSON Web Key Set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-trustedtokenissuer-oidcjwtconfiguration.html#cfn-sso-trustedtokenissuer-oidcjwtconfiguration-jwksretrievaloption
	//
	JwksRetrievalOption *string `field:"required" json:"jwksRetrievalOption" yaml:"jwksRetrievalOption"`
}

