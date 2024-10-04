package awscdkcognitoidentitypoolalpha


// Map roles to users in the Identity Pool based on claims from the Identity Provider.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   var identityPool identityPool
//   var myAddedRoleMapping1 identityPoolRoleMapping
//   var myAddedRoleMapping2 identityPoolRoleMapping
//   var myAddedRoleMapping3 identityPoolRoleMapping
//
//
//   identityPool.AddRoleMappings(myAddedRoleMapping1, myAddedRoleMapping2, myAddedRoleMapping3)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
//
// Experimental.
type IdentityPoolRoleMapping struct {
	// The url of the Provider for which the role is mapped.
	// Experimental.
	ProviderUrl IdentityPoolProviderUrl `field:"required" json:"providerUrl" yaml:"providerUrl"`
	// The key used for the role mapping in the role mapping hash.
	//
	// Required if the providerUrl is a token.
	// Default: - The provided providerUrl.
	//
	// Experimental.
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
	// Allow for role assumption when results of role mapping are ambiguous.
	// Default: false - Ambiguous role resolutions will lead to requester being denied.
	//
	// Experimental.
	ResolveAmbiguousRoles *bool `field:"optional" json:"resolveAmbiguousRoles" yaml:"resolveAmbiguousRoles"`
	// The claim and value that must be matched in order to assume the role.
	//
	// Required if useToken is false.
	// Default: - No role mapping rule.
	//
	// Experimental.
	Rules *[]*RoleMappingRule `field:"optional" json:"rules" yaml:"rules"`
	// If true then mapped roles must be passed through the cognito:roles or cognito:preferred_role claims from Identity Provider.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/role-based-access-control.html#using-tokens-to-assign-roles-to-users
	//
	// Default: false.
	//
	// Experimental.
	UseToken *bool `field:"optional" json:"useToken" yaml:"useToken"`
}

