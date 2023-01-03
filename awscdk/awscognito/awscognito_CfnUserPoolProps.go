package awscognito


// Properties for defining a `CfnUserPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPoolTags interface{}
//
//   cfnUserPoolProps := &cfnUserPoolProps{
//   	accountRecoverySetting: &accountRecoverySettingProperty{
//   		recoveryMechanisms: []interface{}{
//   			&recoveryOptionProperty{
//   				name: jsii.String("name"),
//   				priority: jsii.Number(123),
//   			},
//   		},
//   	},
//   	adminCreateUserConfig: &adminCreateUserConfigProperty{
//   		allowAdminCreateUserOnly: jsii.Boolean(false),
//   		inviteMessageTemplate: &inviteMessageTemplateProperty{
//   			emailMessage: jsii.String("emailMessage"),
//   			emailSubject: jsii.String("emailSubject"),
//   			smsMessage: jsii.String("smsMessage"),
//   		},
//   		unusedAccountValidityDays: jsii.Number(123),
//   	},
//   	aliasAttributes: []*string{
//   		jsii.String("aliasAttributes"),
//   	},
//   	autoVerifiedAttributes: []*string{
//   		jsii.String("autoVerifiedAttributes"),
//   	},
//   	deletionProtection: jsii.String("deletionProtection"),
//   	deviceConfiguration: &deviceConfigurationProperty{
//   		challengeRequiredOnNewDevice: jsii.Boolean(false),
//   		deviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   	},
//   	emailConfiguration: &emailConfigurationProperty{
//   		configurationSet: jsii.String("configurationSet"),
//   		emailSendingAccount: jsii.String("emailSendingAccount"),
//   		from: jsii.String("from"),
//   		replyToEmailAddress: jsii.String("replyToEmailAddress"),
//   		sourceArn: jsii.String("sourceArn"),
//   	},
//   	emailVerificationMessage: jsii.String("emailVerificationMessage"),
//   	emailVerificationSubject: jsii.String("emailVerificationSubject"),
//   	enabledMfas: []*string{
//   		jsii.String("enabledMfas"),
//   	},
//   	lambdaConfig: &lambdaConfigProperty{
//   		createAuthChallenge: jsii.String("createAuthChallenge"),
//   		customEmailSender: &customEmailSenderProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			lambdaVersion: jsii.String("lambdaVersion"),
//   		},
//   		customMessage: jsii.String("customMessage"),
//   		customSmsSender: &customSMSSenderProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			lambdaVersion: jsii.String("lambdaVersion"),
//   		},
//   		defineAuthChallenge: jsii.String("defineAuthChallenge"),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		postAuthentication: jsii.String("postAuthentication"),
//   		postConfirmation: jsii.String("postConfirmation"),
//   		preAuthentication: jsii.String("preAuthentication"),
//   		preSignUp: jsii.String("preSignUp"),
//   		preTokenGeneration: jsii.String("preTokenGeneration"),
//   		userMigration: jsii.String("userMigration"),
//   		verifyAuthChallengeResponse: jsii.String("verifyAuthChallengeResponse"),
//   	},
//   	mfaConfiguration: jsii.String("mfaConfiguration"),
//   	policies: &policiesProperty{
//   		passwordPolicy: &passwordPolicyProperty{
//   			minimumLength: jsii.Number(123),
//   			requireLowercase: jsii.Boolean(false),
//   			requireNumbers: jsii.Boolean(false),
//   			requireSymbols: jsii.Boolean(false),
//   			requireUppercase: jsii.Boolean(false),
//   			temporaryPasswordValidityDays: jsii.Number(123),
//   		},
//   	},
//   	schema: []interface{}{
//   		&schemaAttributeProperty{
//   			attributeDataType: jsii.String("attributeDataType"),
//   			developerOnlyAttribute: jsii.Boolean(false),
//   			mutable: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   			numberAttributeConstraints: &numberAttributeConstraintsProperty{
//   				maxValue: jsii.String("maxValue"),
//   				minValue: jsii.String("minValue"),
//   			},
//   			required: jsii.Boolean(false),
//   			stringAttributeConstraints: &stringAttributeConstraintsProperty{
//   				maxLength: jsii.String("maxLength"),
//   				minLength: jsii.String("minLength"),
//   			},
//   		},
//   	},
//   	smsAuthenticationMessage: jsii.String("smsAuthenticationMessage"),
//   	smsConfiguration: &smsConfigurationProperty{
//   		externalId: jsii.String("externalId"),
//   		snsCallerArn: jsii.String("snsCallerArn"),
//   		snsRegion: jsii.String("snsRegion"),
//   	},
//   	smsVerificationMessage: jsii.String("smsVerificationMessage"),
//   	userAttributeUpdateSettings: &userAttributeUpdateSettingsProperty{
//   		attributesRequireVerificationBeforeUpdate: []*string{
//   			jsii.String("attributesRequireVerificationBeforeUpdate"),
//   		},
//   	},
//   	usernameAttributes: []*string{
//   		jsii.String("usernameAttributes"),
//   	},
//   	usernameConfiguration: &usernameConfigurationProperty{
//   		caseSensitive: jsii.Boolean(false),
//   	},
//   	userPoolAddOns: &userPoolAddOnsProperty{
//   		advancedSecurityMode: jsii.String("advancedSecurityMode"),
//   	},
//   	userPoolName: jsii.String("userPoolName"),
//   	userPoolTags: userPoolTags,
//   	verificationMessageTemplate: &verificationMessageTemplateProperty{
//   		defaultEmailOption: jsii.String("defaultEmailOption"),
//   		emailMessage: jsii.String("emailMessage"),
//   		emailMessageByLink: jsii.String("emailMessageByLink"),
//   		emailSubject: jsii.String("emailSubject"),
//   		emailSubjectByLink: jsii.String("emailSubjectByLink"),
//   		smsMessage: jsii.String("smsMessage"),
//   	},
//   }
//
type CfnUserPoolProps struct {
	// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
	//
	// It allows you to define a preferred method when a user has more than one method available. With this setting, SMS does not qualify for a valid password recovery mechanism if the user also has SMS MFA enabled. In the absence of this setting, Cognito uses the legacy behavior to determine the recovery method where SMS is preferred over email.
	AccountRecoverySetting interface{} `field:"optional" json:"accountRecoverySetting" yaml:"accountRecoverySetting"`
	// The configuration for creating a new user profile.
	AdminCreateUserConfig interface{} `field:"optional" json:"adminCreateUserConfig" yaml:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
	//
	// > This user pool property cannot be updated.
	AliasAttributes *[]*string `field:"optional" json:"aliasAttributes" yaml:"aliasAttributes"`
	// The attributes to be auto-verified.
	//
	// Possible values: *email* , *phone_number* .
	AutoVerifiedAttributes *[]*string `field:"optional" json:"autoVerifiedAttributes" yaml:"autoVerifiedAttributes"`
	// `AWS::Cognito::UserPool.DeletionProtection`.
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The device configuration.
	DeviceConfiguration interface{} `field:"optional" json:"deviceConfiguration" yaml:"deviceConfiguration"`
	// The email configuration of your user pool.
	//
	// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
	EmailConfiguration interface{} `field:"optional" json:"emailConfiguration" yaml:"emailConfiguration"`
	// A string representing the email verification message.
	//
	// EmailVerificationMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationMessage *string `field:"optional" json:"emailVerificationMessage" yaml:"emailVerificationMessage"`
	// A string representing the email verification subject.
	//
	// EmailVerificationSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationSubject *string `field:"optional" json:"emailVerificationSubject" yaml:"emailVerificationSubject"`
	// Enables MFA on a specified user pool.
	//
	// To disable all MFAs after it has been enabled, set MfaConfiguration to “OFF” and remove EnabledMfas. MFAs can only be all disabled if MfaConfiguration is OFF. Once SMS_MFA is enabled, SMS_MFA can only be disabled by setting MfaConfiguration to “OFF”. Can be one of the following values:
	//
	// - `SMS_MFA` - Enables SMS MFA for the user pool. SMS_MFA can only be enabled if SMS configuration is provided.
	// - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
	//
	// Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA`.
	EnabledMfas *[]*string `field:"optional" json:"enabledMfas" yaml:"enabledMfas"`
	// The Lambda trigger configuration information for the new user pool.
	//
	// > In a push model, event sources (such as Amazon S3 and custom applications) need permission to invoke a function. So you must make an extra call to add permission for these event sources to invoke your Lambda function.
	// >
	// > For more information on using the Lambda API to add permission, see [AddPermission](https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html) .
	// >
	// > For adding permission using the AWS CLI , see [add-permission](https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html) .
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// The multi-factor (MFA) configuration. Valid values include:.
	//
	// - `OFF` MFA won't be used for any users.
	// - `ON` MFA is required for all users to sign in.
	// - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
	MfaConfiguration *string `field:"optional" json:"mfaConfiguration" yaml:"mfaConfiguration"`
	// The policy associated with a user pool.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// The schema attributes for the new user pool. These attributes can be standard or custom attributes.
	//
	// > During a user pool update, you can add new schema attributes but you cannot modify or delete an existing schema attribute.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string `field:"optional" json:"smsAuthenticationMessage" yaml:"smsAuthenticationMessage"`
	// The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service.
	//
	// To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
	SmsConfiguration interface{} `field:"optional" json:"smsConfiguration" yaml:"smsConfiguration"`
	// A string representing the SMS verification message.
	SmsVerificationMessage *string `field:"optional" json:"smsVerificationMessage" yaml:"smsVerificationMessage"`
	// The settings for updates to user attributes.
	//
	// These settings include the property `AttributesRequireVerificationBeforeUpdate` ,
	// a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
	// more information, see [Verifying updates to to email addresses and phone numbers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates) .
	UserAttributeUpdateSettings interface{} `field:"optional" json:"userAttributeUpdateSettings" yaml:"userAttributeUpdateSettings"`
	// Determines whether email addresses or phone numbers can be specified as user names when a user signs up.
	//
	// Possible values: `phone_number` or `email` .
	//
	// This user pool property cannot be updated.
	UsernameAttributes *[]*string `field:"optional" json:"usernameAttributes" yaml:"usernameAttributes"`
	// You can choose to set case sensitivity on the username input for the selected sign-in option.
	//
	// For example, when this is set to `False` , users will be able to sign in using either "username" or "Username". This configuration is immutable once it has been set.
	UsernameConfiguration interface{} `field:"optional" json:"usernameConfiguration" yaml:"usernameConfiguration"`
	// Enables advanced security risk detection.
	//
	// Set the key `AdvancedSecurityMode` to the value "AUDIT".
	UserPoolAddOns interface{} `field:"optional" json:"userPoolAddOns" yaml:"userPoolAddOns"`
	// A string used to name the user pool.
	UserPoolName *string `field:"optional" json:"userPoolName" yaml:"userPoolName"`
	// The tag keys and values to assign to the user pool.
	//
	// A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
	UserPoolTags interface{} `field:"optional" json:"userPoolTags" yaml:"userPoolTags"`
	// The template for the verification message that the user sees when the app requests permission to access the user's information.
	VerificationMessageTemplate interface{} `field:"optional" json:"verificationMessageTemplate" yaml:"verificationMessageTemplate"`
}

