// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Props for the User Pool Authentication Provider.
//
// Example:
//   var identityPool identityPool
//
//   userPool := cognito.NewUserPool(this, jsii.String("Pool"))
//   identityPool.addUserPoolAuthentication(awscdkcognitoidentitypoolalpha.NewUserPoolAuthenticationProvider(&userPoolAuthenticationProviderProps{
//   	userPool: userPool,
//   	disableServerSideTokenCheck: jsii.Boolean(true),
//   }))
//
// Experimental.
type UserPoolAuthenticationProviderProps struct {
	// The User Pool of the Associated Identity Providers.
	// Experimental.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Setting this to true turns off identity pool checks for this user pool to make sure the user has not been globally signed out or deleted before the identity pool provides an OIDC token or AWS credentials for the user.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
	//
	// Experimental.
	DisableServerSideTokenCheck *bool `field:"optional" json:"disableServerSideTokenCheck" yaml:"disableServerSideTokenCheck"`
	// The User Pool Client for the provided User Pool.
	// Experimental.
	UserPoolClient awscognito.IUserPoolClient `field:"optional" json:"userPoolClient" yaml:"userPoolClient"`
}

