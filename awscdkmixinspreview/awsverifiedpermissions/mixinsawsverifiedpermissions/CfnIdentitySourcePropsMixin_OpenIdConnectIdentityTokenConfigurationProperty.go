package mixinsawsverifiedpermissions


// The configuration of an OpenID Connect (OIDC) identity source for handling identity (ID) token claims.
//
// Contains the claim that you want to identify as the principal in an authorization request, and the values of the `aud` claim, or audiences, that you want to accept.
//
// This data type is part of a [OpenIdConnectTokenSelection](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_OpenIdConnectTokenSelection.html) structure, which is a parameter of [CreateIdentitySource](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIdConnectIdentityTokenConfigurationProperty := &OpenIdConnectIdentityTokenConfigurationProperty{
//   	ClientIds: []*string{
//   		jsii.String("clientIds"),
//   	},
//   	PrincipalIdClaim: jsii.String("principalIdClaim"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration.html
//
type CfnIdentitySourcePropsMixin_OpenIdConnectIdentityTokenConfigurationProperty struct {
	// The ID token audience, or client ID, claim values that you want to accept in your policy store from an OIDC identity provider.
	//
	// For example, `1example23456789, 2example10111213` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration-clientids
	//
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The claim that determines the principal in OIDC access tokens.
	//
	// For example, `sub` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration-principalidclaim
	//
	// Default: - "sub".
	//
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

