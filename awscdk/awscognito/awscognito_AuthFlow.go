package awscognito


// Types of authentication flow.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("pool"))
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	authFlows: &authFlow{
//   		userPassword: jsii.Boolean(true),
//   		userSrp: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
//
// Experimental.
type AuthFlow struct {
	// Enable admin based user password authentication flow.
	// Experimental.
	AdminUserPassword *bool `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// Enable custom authentication flow.
	// Experimental.
	Custom *bool `field:"optional" json:"custom" yaml:"custom"`
	// Enable auth using username & password.
	// Experimental.
	UserPassword *bool `field:"optional" json:"userPassword" yaml:"userPassword"`
	// Enable SRP based authentication.
	// Experimental.
	UserSrp *bool `field:"optional" json:"userSrp" yaml:"userSrp"`
}

