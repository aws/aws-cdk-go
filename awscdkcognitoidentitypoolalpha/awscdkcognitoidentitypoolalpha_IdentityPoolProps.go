// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for the IdentityPool construct.
//
// Example:
//   var openIdConnectProvider openIdConnectProvider
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &identityPoolProps{
//   	identityPoolName: jsii.String("myidentitypool"),
//   	authenticationProviders: &identityPoolAuthenticationProviders{
//   		google: &identityPoolGoogleLoginProvider{
//   			clientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		openIdConnectProviders: []iOpenIdConnectProvider{
//   			openIdConnectProvider,
//   		},
//   		customProvider: jsii.String("my-custom-provider.example.com"),
//   	},
//   })
//
// Experimental.
type IdentityPoolProps struct {
	// Enables the Basic (Classic) authentication flow.
	// Experimental.
	AllowClassicFlow *bool `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Wwhether the identity pool supports unauthenticated logins.
	// Experimental.
	AllowUnauthenticatedIdentities *bool `field:"optional" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// The Default Role to be assumed by Authenticated Users.
	// Experimental.
	AuthenticatedRole awsiam.IRole `field:"optional" json:"authenticatedRole" yaml:"authenticatedRole"`
	// Authentication providers for using in identity pool.
	// Experimental.
	AuthenticationProviders *IdentityPoolAuthenticationProviders `field:"optional" json:"authenticationProviders" yaml:"authenticationProviders"`
	// The name of the Identity Pool.
	// Experimental.
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// Rules for mapping roles to users.
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// The Default Role to be assumed by Unauthenticated Users.
	// Experimental.
	UnauthenticatedRole awsiam.IRole `field:"optional" json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

