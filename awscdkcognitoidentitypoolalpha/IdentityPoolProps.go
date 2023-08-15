package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for the IdentityPool construct.
//
// Example:
//   var openIdConnectProvider openIdConnectProvider
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	AuthenticationProviders: &IdentityPoolAuthenticationProviders{
//   		Google: &IdentityPoolGoogleLoginProvider{
//   			ClientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		OpenIdConnectProviders: []iOpenIdConnectProvider{
//   			openIdConnectProvider,
//   		},
//   		CustomProvider: jsii.String("my-custom-provider.example.com"),
//   	},
//   })
//
// Experimental.
type IdentityPoolProps struct {
	// Enables the Basic (Classic) authentication flow.
	// Default: - Classic Flow not allowed.
	//
	// Experimental.
	AllowClassicFlow *bool `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Wwhether the identity pool supports unauthenticated logins.
	// Default: - false.
	//
	// Experimental.
	AllowUnauthenticatedIdentities *bool `field:"optional" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// The Default Role to be assumed by Authenticated Users.
	// Default: - A Default Authenticated Role will be added.
	//
	// Experimental.
	AuthenticatedRole awsiam.IRole `field:"optional" json:"authenticatedRole" yaml:"authenticatedRole"`
	// Authentication providers for using in identity pool.
	// Default: - No Authentication Providers passed directly to Identity Pool.
	//
	// Experimental.
	AuthenticationProviders *IdentityPoolAuthenticationProviders `field:"optional" json:"authenticationProviders" yaml:"authenticationProviders"`
	// The name of the Identity Pool.
	// Default: - automatically generated name by CloudFormation at deploy time.
	//
	// Experimental.
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// Rules for mapping roles to users.
	// Default: - no Role Mappings.
	//
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// The Default Role to be assumed by Unauthenticated Users.
	// Default: - A Default Unauthenticated Role will be added.
	//
	// Experimental.
	UnauthenticatedRole awsiam.IRole `field:"optional" json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

