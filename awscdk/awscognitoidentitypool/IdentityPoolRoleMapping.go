package awscognitoidentitypool


// Map roles to users in the Identity Pool based on claims from the Identity Provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var identityPoolProviderUrl identityPoolProviderUrl
//   var role role
//
//   identityPoolRoleMapping := &IdentityPoolRoleMapping{
//   	ProviderUrl: identityPoolProviderUrl,
//
//   	// the properties below are optional
//   	MappingKey: jsii.String("mappingKey"),
//   	ResolveAmbiguousRoles: jsii.Boolean(false),
//   	Rules: []roleMappingRule{
//   		&roleMappingRule{
//   			Claim: jsii.String("claim"),
//   			ClaimValue: jsii.String("claimValue"),
//   			MappedRole: role,
//
//   			// the properties below are optional
//   			MatchType: awscdk.Aws_cognito_identitypool.RoleMappingMatchType_EQUALS,
//   		},
//   	},
//   	UseToken: jsii.Boolean(false),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
//
type IdentityPoolRoleMapping struct {
	// The url of the Provider for which the role is mapped.
	ProviderUrl IdentityPoolProviderUrl `field:"required" json:"providerUrl" yaml:"providerUrl"`
	// The key used for the role mapping in the role mapping hash.
	//
	// Required if the providerUrl is a token.
	// Default: - The provided providerUrl.
	//
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
	// Allow for role assumption when results of role mapping are ambiguous.
	// Default: false - Ambiguous role resolutions will lead to requester being denied.
	//
	ResolveAmbiguousRoles *bool `field:"optional" json:"resolveAmbiguousRoles" yaml:"resolveAmbiguousRoles"`
	// The claim and value that must be matched in order to assume the role.
	//
	// Required if useToken is false.
	// Default: - No role mapping rule.
	//
	Rules *[]*RoleMappingRule `field:"optional" json:"rules" yaml:"rules"`
	// If true then mapped roles must be passed through the cognito:roles or cognito:preferred_role claims from Identity Provider.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/role-based-access-control.html#using-tokens-to-assign-roles-to-users
	//
	// Default: false.
	//
	UseToken *bool `field:"optional" json:"useToken" yaml:"useToken"`
}

