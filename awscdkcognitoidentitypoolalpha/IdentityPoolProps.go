package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for the Identity Pool construct.
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
// Deprecated.
type IdentityPoolProps struct {
	// Enables the Basic (Classic) authentication flow.
	// Default: - Classic Flow not allowed.
	//
	// Deprecated.
	AllowClassicFlow *bool `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Whether the Identity Pool supports unauthenticated logins.
	// Default: - false.
	//
	// Deprecated.
	AllowUnauthenticatedIdentities *bool `field:"optional" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// The default Role to be assumed by authenticated users.
	// Default: - A default authenticated Role will be added.
	//
	// Deprecated.
	AuthenticatedRole awsiam.IRole `field:"optional" json:"authenticatedRole" yaml:"authenticatedRole"`
	// Authentication Providers for using in Identity Pool.
	// Default: - No Authentication Providers passed directly to Identity Pool.
	//
	// Deprecated.
	AuthenticationProviders *IdentityPoolAuthenticationProviders `field:"optional" json:"authenticationProviders" yaml:"authenticationProviders"`
	// The name of the Identity Pool.
	// Default: - Automatically generated name by CloudFormation at deploy time.
	//
	// Deprecated.
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// Rules for mapping roles to users.
	// Default: - no role mappings.
	//
	// Deprecated.
	RoleMappings *[]*IdentityPoolRoleMapping `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// The default Role to be assumed by unauthenticated users.
	// Default: - A default unauthenticated Role will be added.
	//
	// Deprecated.
	UnauthenticatedRole awsiam.IRole `field:"optional" json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

