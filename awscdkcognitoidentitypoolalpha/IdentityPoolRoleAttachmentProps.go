package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for an Identity Pool Role Attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var identityPool identityPool
//   var identityPoolProviderUrl identityPoolProviderUrl
//   var role role
//
//   identityPoolRoleAttachmentProps := &IdentityPoolRoleAttachmentProps{
//   	IdentityPool: identityPool,
//
//   	// the properties below are optional
//   	AuthenticatedRole: role,
//   	RoleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			ProviderUrl: identityPoolProviderUrl,
//
//   			// the properties below are optional
//   			MappingKey: jsii.String("mappingKey"),
//   			ResolveAmbiguousRoles: jsii.Boolean(false),
//   			Rules: []roleMappingRule{
//   				&roleMappingRule{
//   					Claim: jsii.String("claim"),
//   					ClaimValue: jsii.String("claimValue"),
//   					MappedRole: role,
//
//   					// the properties below are optional
//   					MatchType: cognito_identitypool_alpha.RoleMappingMatchType_EQUALS,
//   				},
//   			},
//   			UseToken: jsii.Boolean(false),
//   		},
//   	},
//   	UnauthenticatedRole: role,
//   }
//
// Experimental.
type IdentityPoolRoleAttachmentProps struct {
	// ID of the Attachment's underlying Identity Pool.
	// Experimental.
	IdentityPool IIdentityPool `field:"required" json:"identityPool" yaml:"identityPool"`
	// Default authenticated (User) Role.
	// Default: - No default authenticated Role will be added.
	//
	// Experimental.
	AuthenticatedRole awsiam.IRole `field:"optional" json:"authenticatedRole" yaml:"authenticatedRole"`
	// Rules for mapping roles to users.
	// Default: - No role mappings.
	//
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// Default unauthenticated (Guest) Role.
	// Default: - No default unauthenticated Role will be added.
	//
	// Experimental.
	UnauthenticatedRole awsiam.IRole `field:"optional" json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

