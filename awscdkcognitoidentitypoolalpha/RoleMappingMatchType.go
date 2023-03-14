// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// Types of matches allowed for Role Mapping.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   var adminRole role
//   var nonAdminRole role
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	// Assign specific roles to users based on whether or not the custom admin claim is passed from the identity provider
//   	RoleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			ProviderUrl: awscdkcognitoidentitypoolalpha.IdentityPoolProviderUrl_AMAZON(),
//   			Rules: []roleMappingRule{
//   				&roleMappingRule{
//   					Claim: jsii.String("custom:admin"),
//   					ClaimValue: jsii.String("admin"),
//   					MappedRole: adminRole,
//   				},
//   				&roleMappingRule{
//   					Claim: jsii.String("custom:admin"),
//   					ClaimValue: jsii.String("admin"),
//   					MatchType: awscdkcognitoidentitypoolalpha.RoleMappingMatchType_NOTEQUAL,
//   					MappedRole: nonAdminRole,
//   				},
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type RoleMappingMatchType string

const (
	// The Claim from the token must equal the given value in order for a match.
	// Experimental.
	RoleMappingMatchType_EQUALS RoleMappingMatchType = "EQUALS"
	// The Claim from the token must contain the given value in order for a match.
	// Experimental.
	RoleMappingMatchType_CONTAINS RoleMappingMatchType = "CONTAINS"
	// The Claim from the token must start with the given value in order for a match.
	// Experimental.
	RoleMappingMatchType_STARTS_WITH RoleMappingMatchType = "STARTS_WITH"
	// The Claim from the token must not equal the given value in order for a match.
	// Experimental.
	RoleMappingMatchType_NOTEQUAL RoleMappingMatchType = "NOTEQUAL"
)

