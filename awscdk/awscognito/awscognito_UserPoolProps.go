package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Props for the UserPool construct.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	selfSignUpEnabled: jsii.Boolean(true),
//   	userVerification: &userVerificationConfig{
//   		emailSubject: jsii.String("Verify your email for our awesome app!"),
//   		emailBody: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   		emailStyle: cognito.verificationEmailStyle_CODE,
//   		smsMessage: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   	},
//   })
//
type UserPoolProps struct {
	// How will a user be able to recover their account?
	AccountRecovery AccountRecovery `field:"optional" json:"accountRecovery" yaml:"accountRecovery"`
	// Attributes which Cognito will look to verify automatically upon user sign up.
	//
	// EMAIL and PHONE are the only available options.
	AutoVerify *AutoVerifiedAttrs `field:"optional" json:"autoVerify" yaml:"autoVerify"`
	// Define a set of custom attributes that can be configured for each user in the user pool.
	CustomAttributes *map[string]ICustomAttribute `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// This key will be used to encrypt temporary passwords and authorization codes that Amazon Cognito generates.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-sender-triggers.html
	//
	CustomSenderKmsKey awskms.IKey `field:"optional" json:"customSenderKmsKey" yaml:"customSenderKmsKey"`
	// Device tracking settings.
	DeviceTracking *DeviceTracking `field:"optional" json:"deviceTracking" yaml:"deviceTracking"`
	// Email settings for a user pool.
	Email UserPoolEmail `field:"optional" json:"email" yaml:"email"`
	// Setting this would explicitly enable or disable SMS role creation.
	//
	// When left unspecified, CDK will determine based on other properties if a role is needed or not.
	EnableSmsRole *bool `field:"optional" json:"enableSmsRole" yaml:"enableSmsRole"`
	// Attributes which Cognito will look to handle changes to the value of your users' email address and phone number attributes.
	//
	// EMAIL and PHONE are the only available options.
	KeepOriginal *KeepOriginalAttrs `field:"optional" json:"keepOriginal" yaml:"keepOriginal"`
	// Lambda functions to use for supported Cognito triggers.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	//
	LambdaTriggers *UserPoolTriggers `field:"optional" json:"lambdaTriggers" yaml:"lambdaTriggers"`
	// Configure whether users of this user pool can or are required use MFA to sign in.
	Mfa Mfa `field:"optional" json:"mfa" yaml:"mfa"`
	// The SMS message template sent during MFA verification.
	//
	// Use '{####}' in the template where Cognito should insert the verification code.
	MfaMessage *string `field:"optional" json:"mfaMessage" yaml:"mfaMessage"`
	// Configure the MFA types that users can use in this user pool.
	//
	// Ignored if `mfa` is set to `OFF`.
	MfaSecondFactor *MfaSecondFactor `field:"optional" json:"mfaSecondFactor" yaml:"mfaSecondFactor"`
	// Password policy for this user pool.
	PasswordPolicy *PasswordPolicy `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Policy to apply when the user pool is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Whether self sign up should be enabled.
	//
	// This can be further configured via the `selfSignUp` property.
	SelfSignUpEnabled *bool `field:"optional" json:"selfSignUpEnabled" yaml:"selfSignUpEnabled"`
	// Methods in which a user registers or signs in to a user pool.
	//
	// Allows either username with aliases OR sign in with email, phone, or both.
	//
	// Read the sections on usernames and aliases to learn more -
	// https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html
	//
	// To match with 'Option 1' in the above link, with a verified email, this property should be set to
	// `{ username: true, email: true }`. To match with 'Option 2' in the above link with both a verified email and phone
	// number, this property should be set to `{ email: true, phone: true }`.
	SignInAliases *SignInAliases `field:"optional" json:"signInAliases" yaml:"signInAliases"`
	// Whether sign-in aliases should be evaluated with case sensitivity.
	//
	// For example, when this option is set to false, users will be able to sign in using either `MyUsername` or `myusername`.
	SignInCaseSensitive *bool `field:"optional" json:"signInCaseSensitive" yaml:"signInCaseSensitive"`
	// The IAM role that Cognito will assume while sending SMS messages.
	SmsRole awsiam.IRole `field:"optional" json:"smsRole" yaml:"smsRole"`
	// The 'ExternalId' that Cognito service must using when assuming the `smsRole`, if the role is restricted with an 'sts:ExternalId' conditional.
	//
	// Learn more about ExternalId here - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	//
	// This property will be ignored if `smsRole` is not specified.
	SmsRoleExternalId *string `field:"optional" json:"smsRoleExternalId" yaml:"smsRoleExternalId"`
	// The region to integrate with SNS to send SMS messages.
	//
	// This property will do nothing if SMS configuration is not configured.
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
	// The set of attributes that are required for every user in the user pool.
	//
	// Read more on attributes here - https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html
	StandardAttributes *StandardAttributes `field:"optional" json:"standardAttributes" yaml:"standardAttributes"`
	// Configuration around admins signing up users into a user pool.
	UserInvitation *UserInvitationConfig `field:"optional" json:"userInvitation" yaml:"userInvitation"`
	// Name of the user pool.
	UserPoolName *string `field:"optional" json:"userPoolName" yaml:"userPoolName"`
	// Configuration around users signing themselves up to the user pool.
	//
	// Enable or disable self sign-up via the `selfSignUpEnabled` property.
	UserVerification *UserVerificationConfig `field:"optional" json:"userVerification" yaml:"userVerification"`
}

