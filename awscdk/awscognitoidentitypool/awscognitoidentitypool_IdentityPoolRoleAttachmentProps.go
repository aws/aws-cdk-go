package awscognitoidentitypool

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Props for an Identity Pool Role Attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var identityPool identityPool
//   var identityPoolProviderUrl identityPoolProviderUrl
//   var role role
//
//   identityPoolRoleAttachmentProps := &identityPoolRoleAttachmentProps{
//   	identityPool: identityPool,
//
//   	// the properties below are optional
//   	authenticatedRole: role,
//   	roleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			providerUrl: identityPoolProviderUrl,
//
//   			// the properties below are optional
//   			resolveAmbiguousRoles: jsii.Boolean(false),
//   			rules: []roleMappingRule{
//   				&roleMappingRule{
//   					claim: jsii.String("claim"),
//   					claimValue: jsii.String("claimValue"),
//   					mappedRole: role,
//
//   					// the properties below are optional
//   					matchType: awscdk.Aws_cognito_identitypool.roleMappingMatchType_EQUALS,
//   				},
//   			},
//   			useToken: jsii.Boolean(false),
//   		},
//   	},
//   	unauthenticatedRole: role,
//   }
//
// Experimental.
type IdentityPoolRoleAttachmentProps struct {
	// Id of the Attachments Underlying Identity Pool.
	// Experimental.
	IdentityPool IIdentityPool `field:"required" json:"identityPool" yaml:"identityPool"`
	// Default Authenticated (User) Role.
	// Experimental.
	AuthenticatedRole awsiam.IRole `field:"optional" json:"authenticatedRole" yaml:"authenticatedRole"`
	// Rules for mapping roles to users.
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// Default Unauthenticated (Guest) Role.
	// Experimental.
	UnauthenticatedRole awsiam.IRole `field:"optional" json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

