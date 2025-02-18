package awscognito


// Types of authentication flow.
//
// Example:
//   userPool := cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	SignInPolicy: &SignInPolicy{
//   		AllowedFirstAuthFactors: &AllowedFirstAuthFactors{
//   			Password: jsii.Boolean(true),
//   			 // password authentication must be enabled
//   			EmailOtp: jsii.Boolean(true),
//   			 // enables email message one-time password
//   			SmsOtp: jsii.Boolean(true),
//   			 // enables SMS message one-time password
//   			Passkey: jsii.Boolean(true),
//   		},
//   	},
//   })
//
//   // You should also configure the user pool client with USER_AUTH authentication flow allowed
//   userPool.addClient(jsii.String("myclient"), &UserPoolClientOptions{
//   	AuthFlows: &AuthFlow{
//   		User: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
//
type AuthFlow struct {
	// Enable admin based user password authentication flow.
	// Default: false.
	//
	AdminUserPassword *bool `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// Enable custom authentication flow.
	// Default: false.
	//
	Custom *bool `field:"optional" json:"custom" yaml:"custom"`
	// Enable Choice-based authentication.
	// Default: false.
	//
	User *bool `field:"optional" json:"user" yaml:"user"`
	// Enable auth using username & password.
	// Default: false.
	//
	UserPassword *bool `field:"optional" json:"userPassword" yaml:"userPassword"`
	// Enable SRP based authentication.
	// Default: false.
	//
	UserSrp *bool `field:"optional" json:"userSrp" yaml:"userSrp"`
}

