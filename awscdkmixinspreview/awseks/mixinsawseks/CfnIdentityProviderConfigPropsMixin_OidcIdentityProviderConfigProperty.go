package mixinsawseks


// An object representing the configuration for an OpenID Connect (OIDC) identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oidcIdentityProviderConfigProperty := &OidcIdentityProviderConfigProperty{
//   	ClientId: jsii.String("clientId"),
//   	GroupsClaim: jsii.String("groupsClaim"),
//   	GroupsPrefix: jsii.String("groupsPrefix"),
//   	IssuerUrl: jsii.String("issuerUrl"),
//   	RequiredClaims: []interface{}{
//   		&RequiredClaimProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UsernameClaim: jsii.String("usernameClaim"),
//   	UsernamePrefix: jsii.String("usernamePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html
//
type CfnIdentityProviderConfigPropsMixin_OidcIdentityProviderConfigProperty struct {
	// This is also known as *audience* .
	//
	// The ID of the client application that makes authentication requests to the OIDC identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The JSON web token (JWT) claim that the provider uses to return your groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-groupsclaim
	//
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// The prefix that is prepended to group claims to prevent clashes with existing names (such as `system:` groups).
	//
	// For example, the value `oidc:` creates group names like `oidc:engineering` and `oidc:infra` . The prefix can't contain `system:`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-groupsprefix
	//
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// The URL of the OIDC identity provider that allows the API server to discover public signing keys for verifying tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-issuerurl
	//
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// The key-value pairs that describe required claims in the identity token.
	//
	// If set, each claim is verified to be present in the token with a matching value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-requiredclaims
	//
	RequiredClaims interface{} `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// The JSON Web token (JWT) claim that is used as the username.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-usernameclaim
	//
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// The prefix that is prepended to username claims to prevent clashes with existing names.
	//
	// The prefix can't contain `system:`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html#cfn-eks-identityproviderconfig-oidcidentityproviderconfig-usernameprefix
	//
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

