package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Props for the UserPool construct.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	SelfSignUpEnabled: jsii.Boolean(true),
//   	UserVerification: &UserVerificationConfig{
//   		EmailSubject: jsii.String("Verify your email for our awesome app!"),
//   		EmailBody: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   		EmailStyle: cognito.VerificationEmailStyle_CODE,
//   		SmsMessage: jsii.String("Thanks for signing up to our awesome app! Your verification code is {####}"),
//   	},
//   })
//
type UserPoolProps struct {
	// How will a user be able to recover their account?
	// Default: AccountRecovery.PHONE_WITHOUT_MFA_AND_EMAIL
	//
	AccountRecovery AccountRecovery `field:"optional" json:"accountRecovery" yaml:"accountRecovery"`
	// The user pool's Advanced Security Mode.
	// Default: - no value.
	//
	AdvancedSecurityMode AdvancedSecurityMode `field:"optional" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
	// Attributes which Cognito will look to verify automatically upon user sign up.
	//
	// EMAIL and PHONE are the only available options.
	// Default: - If `signInAlias` includes email and/or phone, they will be included in `autoVerifiedAttributes` by default.
	// If absent, no attributes will be auto-verified.
	//
	AutoVerify *AutoVerifiedAttrs `field:"optional" json:"autoVerify" yaml:"autoVerify"`
	// Define a set of custom attributes that can be configured for each user in the user pool.
	// Default: - No custom attributes.
	//
	CustomAttributes *map[string]ICustomAttribute `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// This key will be used to encrypt temporary passwords and authorization codes that Amazon Cognito generates.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-sender-triggers.html
	//
	// Default: - no key ID configured.
	//
	CustomSenderKmsKey awskms.IKey `field:"optional" json:"customSenderKmsKey" yaml:"customSenderKmsKey"`
	// Indicates whether the user pool should have deletion protection enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Device tracking settings.
	// Default: - see defaults on each property of DeviceTracking.
	//
	DeviceTracking *DeviceTracking `field:"optional" json:"deviceTracking" yaml:"deviceTracking"`
	// Email settings for a user pool.
	// Default: - cognito will use the default email configuration.
	//
	Email UserPoolEmail `field:"optional" json:"email" yaml:"email"`
	// Setting this would explicitly enable or disable SMS role creation.
	//
	// When left unspecified, CDK will determine based on other properties if a role is needed or not.
	// Default: - CDK will determine based on other properties of the user pool if an SMS role should be created or not.
	//
	EnableSmsRole *bool `field:"optional" json:"enableSmsRole" yaml:"enableSmsRole"`
	// Attributes which Cognito will look to handle changes to the value of your users' email address and phone number attributes.
	//
	// EMAIL and PHONE are the only available options.
	// Default: - Nothing is kept.
	//
	KeepOriginal *KeepOriginalAttrs `field:"optional" json:"keepOriginal" yaml:"keepOriginal"`
	// Lambda functions to use for supported Cognito triggers.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	//
	// Default: - No Lambda triggers.
	//
	LambdaTriggers *UserPoolTriggers `field:"optional" json:"lambdaTriggers" yaml:"lambdaTriggers"`
	// Configure whether users of this user pool can or are required use MFA to sign in.
	// Default: Mfa.OFF
	//
	Mfa Mfa `field:"optional" json:"mfa" yaml:"mfa"`
	// The SMS message template sent during MFA verification.
	//
	// Use '{####}' in the template where Cognito should insert the verification code.
	// Default: 'Your authentication code is {####}.'
	//
	MfaMessage *string `field:"optional" json:"mfaMessage" yaml:"mfaMessage"`
	// Configure the MFA types that users can use in this user pool.
	//
	// Ignored if `mfa` is set to `OFF`.
	// Default: - { sms: true, otp: false }, if `mfa` is set to `OPTIONAL` or `REQUIRED`.
	// { sms: false, otp: false }, otherwise.
	//
	MfaSecondFactor *MfaSecondFactor `field:"optional" json:"mfaSecondFactor" yaml:"mfaSecondFactor"`
	// Password policy for this user pool.
	// Default: - see defaults on each property of PasswordPolicy.
	//
	PasswordPolicy *PasswordPolicy `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Policy to apply when the user pool is removed from the stack.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Whether self sign-up should be enabled.
	//
	// To configure self sign-up configuration use the `userVerification` property.
	// Default: - false.
	//
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
	// Default: { username: true }.
	//
	SignInAliases *SignInAliases `field:"optional" json:"signInAliases" yaml:"signInAliases"`
	// Whether sign-in aliases should be evaluated with case sensitivity.
	//
	// For example, when this option is set to false, users will be able to sign in using either `MyUsername` or `myusername`.
	// Default: true.
	//
	SignInCaseSensitive *bool `field:"optional" json:"signInCaseSensitive" yaml:"signInCaseSensitive"`
	// The IAM role that Cognito will assume while sending SMS messages.
	// Default: - a new IAM role is created.
	//
	SmsRole awsiam.IRole `field:"optional" json:"smsRole" yaml:"smsRole"`
	// The 'ExternalId' that Cognito service must be using when assuming the `smsRole`, if the role is restricted with an 'sts:ExternalId' conditional.
	//
	// Learn more about ExternalId here - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	//
	// This property will be ignored if `smsRole` is not specified.
	// Default: - No external id will be configured.
	//
	SmsRoleExternalId *string `field:"optional" json:"smsRoleExternalId" yaml:"smsRoleExternalId"`
	// The region to integrate with SNS to send SMS messages.
	//
	// This property will do nothing if SMS configuration is not configured.
	// Default: - The same region as the user pool, with a few exceptions - https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html#user-pool-sms-settings-first-time
	//
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
	// The set of attributes that are required for every user in the user pool.
	//
	// Read more on attributes here - https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html
	// Default: - All standard attributes are optional and mutable.
	//
	StandardAttributes *StandardAttributes `field:"optional" json:"standardAttributes" yaml:"standardAttributes"`
	// Configuration around admins signing up users into a user pool.
	// Default: - see defaults in UserInvitationConfig.
	//
	UserInvitation *UserInvitationConfig `field:"optional" json:"userInvitation" yaml:"userInvitation"`
	// Name of the user pool.
	// Default: - automatically generated name by CloudFormation at deploy time.
	//
	UserPoolName *string `field:"optional" json:"userPoolName" yaml:"userPoolName"`
	// Configuration around users signing themselves up to the user pool.
	//
	// Enable or disable self sign-up via the `selfSignUpEnabled` property.
	// Default: - see defaults in UserVerificationConfig.
	//
	UserVerification *UserVerificationConfig `field:"optional" json:"userVerification" yaml:"userVerification"`
}

