// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// Map roles to users in the identity pool based on claims from the Identity Provider.
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
//   identityPool.addRoleMappings(myAddedRoleMapping1, myAddedRoleMapping2, myAddedRoleMapping3)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
//
// Experimental.
type IdentityPoolRoleMapping struct {
	// The url of the provider of for which the role is mapped.
	// Experimental.
	ProviderUrl IdentityPoolProviderUrl `field:"required" json:"providerUrl" yaml:"providerUrl"`
	// The key used for the role mapping in the role mapping hash.
	//
	// Required if the providerUrl is a token.
	// Experimental.
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
	// Allow for role assumption when results of role mapping are ambiguous.
	// Experimental.
	ResolveAmbiguousRoles *bool `field:"optional" json:"resolveAmbiguousRoles" yaml:"resolveAmbiguousRoles"`
	// The claim and value that must be matched in order to assume the role.
	//
	// Required if useToken is false.
	// Experimental.
	Rules *[]*RoleMappingRule `field:"optional" json:"rules" yaml:"rules"`
	// If true then mapped roles must be passed through the cognito:roles or cognito:preferred_role claims from identity provider.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/role-based-access-control.html#using-tokens-to-assign-roles-to-users
	//
	// Experimental.
	UseToken *bool `field:"optional" json:"useToken" yaml:"useToken"`
}

