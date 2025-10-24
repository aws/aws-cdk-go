package awscognitoidentitypool


// Types of matches allowed for role mapping.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var adminRole Role
//   var nonAdminRole Role
//
//   awscdk.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	// Assign specific roles to users based on whether or not the custom admin claim is passed from the identity provider
//   	RoleMappings: []IdentityPoolRoleMapping{
//   		&IdentityPoolRoleMapping{
//   			ProviderUrl: awscdk.IdentityPoolProviderUrl_AMAZON(),
//   			Rules: []RoleMappingRule{
//   				&RoleMappingRule{
//   					Claim: jsii.String("custom:admin"),
//   					ClaimValue: jsii.String("admin"),
//   					MappedRole: adminRole,
//   				},
//   				&RoleMappingRule{
//   					Claim: jsii.String("custom:admin"),
//   					ClaimValue: jsii.String("admin"),
//   					MatchType: awscdk.RoleMappingMatchType_NOTEQUAL,
//   					MappedRole: nonAdminRole,
//   				},
//   			},
//   		},
//   	},
//   })
//
type RoleMappingMatchType string

const (
	// The claim from the token must equal the given value in order for a match.
	RoleMappingMatchType_EQUALS RoleMappingMatchType = "EQUALS"
	// The claim from the token must contain the given value in order for a match.
	RoleMappingMatchType_CONTAINS RoleMappingMatchType = "CONTAINS"
	// The claim from the token must start with the given value in order for a match.
	RoleMappingMatchType_STARTS_WITH RoleMappingMatchType = "STARTS_WITH"
	// The claim from the token must not equal the given value in order for a match.
	RoleMappingMatchType_NOTEQUAL RoleMappingMatchType = "NOTEQUAL"
)

