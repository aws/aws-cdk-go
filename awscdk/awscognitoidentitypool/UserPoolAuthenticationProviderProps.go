package awscognitoidentitypool

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
)

// Props for the User Pool Authentication Provider.
//
// Example:
//   var identityPool IdentityPool
//
//   userPool := cognito.NewUserPool(this, jsii.String("Pool"))
//   identityPool.AddUserPoolAuthentication(awscdk.NewUserPoolAuthenticationProvider(&UserPoolAuthenticationProviderProps{
//   	UserPool: UserPool,
//   	DisableServerSideTokenCheck: jsii.Boolean(true),
//   }))
//
type UserPoolAuthenticationProviderProps struct {
	// The User Pool of the Associated Identity Providers.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Setting this to true turns off identity pool checks for this user pool to make sure the user has not been globally signed out or deleted before the identity pool provides an OIDC token or AWS credentials for the user.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
	//
	// Default: false.
	//
	DisableServerSideTokenCheck *bool `field:"optional" json:"disableServerSideTokenCheck" yaml:"disableServerSideTokenCheck"`
	// The User Pool Client for the provided User Pool.
	// Default: - A default user pool client will be added to User Pool.
	//
	UserPoolClient interfacesawscognito.IUserPoolClientRef `field:"optional" json:"userPoolClient" yaml:"userPoolClient"`
}

