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
//   cfnUserPoolProps := &CfnUserPoolProps{
//   	AccountRecoverySetting: &AccountRecoverySettingProperty{
//   		RecoveryMechanisms: []interface{}{
//   			&RecoveryOptionProperty{
//   				Name: jsii.String("name"),
//   				Priority: jsii.Number(123),
//   			},
//   		},
//   	},
//   	AdminCreateUserConfig: &AdminCreateUserConfigProperty{
//   		AllowAdminCreateUserOnly: jsii.Boolean(false),
//   		InviteMessageTemplate: &InviteMessageTemplateProperty{
//   			EmailMessage: jsii.String("emailMessage"),
//   			EmailSubject: jsii.String("emailSubject"),
//   			SmsMessage: jsii.String("smsMessage"),
//   		},
//   		UnusedAccountValidityDays: jsii.Number(123),
//   	},
//   	AliasAttributes: []*string{
//   		jsii.String("aliasAttributes"),
//   	},
//   	AutoVerifiedAttributes: []*string{
//   		jsii.String("autoVerifiedAttributes"),
//   	},
//   	DeletionProtection: jsii.String("deletionProtection"),
//   	DeviceConfiguration: &DeviceConfigurationProperty{
//   		ChallengeRequiredOnNewDevice: jsii.Boolean(false),
//   		DeviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   	},
//   	EmailConfiguration: &EmailConfigurationProperty{
//   		ConfigurationSet: jsii.String("configurationSet"),
//   		EmailSendingAccount: jsii.String("emailSendingAccount"),
//   		From: jsii.String("from"),
//   		ReplyToEmailAddress: jsii.String("replyToEmailAddress"),
//   		SourceArn: jsii.String("sourceArn"),
//   	},
//   	EmailVerificationMessage: jsii.String("emailVerificationMessage"),
//   	EmailVerificationSubject: jsii.String("emailVerificationSubject"),
//   	EnabledMfas: []*string{
//   		jsii.String("enabledMfas"),
//   	},
//   	LambdaConfig: &LambdaConfigProperty{
//   		CreateAuthChallenge: jsii.String("createAuthChallenge"),
//   		CustomEmailSender: &CustomEmailSenderProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			LambdaVersion: jsii.String("lambdaVersion"),
//   		},
//   		CustomMessage: jsii.String("customMessage"),
//   		CustomSmsSender: &CustomSMSSenderProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			LambdaVersion: jsii.String("lambdaVersion"),
//   		},
//   		DefineAuthChallenge: jsii.String("defineAuthChallenge"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		PostAuthentication: jsii.String("postAuthentication"),
//   		PostConfirmation: jsii.String("postConfirmation"),
//   		PreAuthentication: jsii.String("preAuthentication"),
//   		PreSignUp: jsii.String("preSignUp"),
//   		PreTokenGeneration: jsii.String("preTokenGeneration"),
//   		PreTokenGenerationConfig: &PreTokenGenerationConfigProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			LambdaVersion: jsii.String("lambdaVersion"),
//   		},
//   		UserMigration: jsii.String("userMigration"),
//   		VerifyAuthChallengeResponse: jsii.String("verifyAuthChallengeResponse"),
//   	},
//   	MfaConfiguration: jsii.String("mfaConfiguration"),
//   	Policies: &PoliciesProperty{
//   		PasswordPolicy: &PasswordPolicyProperty{
//   			MinimumLength: jsii.Number(123),
//   			RequireLowercase: jsii.Boolean(false),
//   			RequireNumbers: jsii.Boolean(false),
//   			RequireSymbols: jsii.Boolean(false),
//   			RequireUppercase: jsii.Boolean(false),
//   			TemporaryPasswordValidityDays: jsii.Number(123),
//   		},
//   	},
//   	Schema: []interface{}{
//   		&SchemaAttributeProperty{
//   			AttributeDataType: jsii.String("attributeDataType"),
//   			DeveloperOnlyAttribute: jsii.Boolean(false),
//   			Mutable: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			NumberAttributeConstraints: &NumberAttributeConstraintsProperty{
//   				MaxValue: jsii.String("maxValue"),
//   				MinValue: jsii.String("minValue"),
//   			},
//   			Required: jsii.Boolean(false),
//   			StringAttributeConstraints: &StringAttributeConstraintsProperty{
//   				MaxLength: jsii.String("maxLength"),
//   				MinLength: jsii.String("minLength"),
//   			},
//   		},
//   	},
//   	SmsAuthenticationMessage: jsii.String("smsAuthenticationMessage"),
//   	SmsConfiguration: &SmsConfigurationProperty{
//   		ExternalId: jsii.String("externalId"),
//   		SnsCallerArn: jsii.String("snsCallerArn"),
//   		SnsRegion: jsii.String("snsRegion"),
//   	},
//   	SmsVerificationMessage: jsii.String("smsVerificationMessage"),
//   	UserAttributeUpdateSettings: &UserAttributeUpdateSettingsProperty{
//   		AttributesRequireVerificationBeforeUpdate: []*string{
//   			jsii.String("attributesRequireVerificationBeforeUpdate"),
//   		},
//   	},
//   	UsernameAttributes: []*string{
//   		jsii.String("usernameAttributes"),
//   	},
//   	UsernameConfiguration: &UsernameConfigurationProperty{
//   		CaseSensitive: jsii.Boolean(false),
//   	},
//   	UserPoolAddOns: &UserPoolAddOnsProperty{
//   		AdvancedSecurityMode: jsii.String("advancedSecurityMode"),
//   	},
//   	UserPoolName: jsii.String("userPoolName"),
//   	UserPoolTags: userPoolTags,
//   	VerificationMessageTemplate: &VerificationMessageTemplateProperty{
//   		DefaultEmailOption: jsii.String("defaultEmailOption"),
//   		EmailMessage: jsii.String("emailMessage"),
//   		EmailMessageByLink: jsii.String("emailMessageByLink"),
//   		EmailSubject: jsii.String("emailSubject"),
//   		EmailSubjectByLink: jsii.String("emailSubjectByLink"),
//   		SmsMessage: jsii.String("smsMessage"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
//
type CfnUserPoolProps struct {
	// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
	//
	// It allows you to define a preferred method when a user has more than one method available. With this setting, SMS does not qualify for a valid password recovery mechanism if the user also has SMS MFA enabled. In the absence of this setting, Cognito uses the legacy behavior to determine the recovery method where SMS is preferred over email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-accountrecoverysetting
	//
	AccountRecoverySetting interface{} `field:"optional" json:"accountRecoverySetting" yaml:"accountRecoverySetting"`
	// The configuration for creating a new user profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-admincreateuserconfig
	//
	AdminCreateUserConfig interface{} `field:"optional" json:"adminCreateUserConfig" yaml:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
	//
	// > This user pool property cannot be updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-aliasattributes
	//
	AliasAttributes *[]*string `field:"optional" json:"aliasAttributes" yaml:"aliasAttributes"`
	// The attributes to be auto-verified.
	//
	// Possible values: *email* , *phone_number* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-autoverifiedattributes
	//
	AutoVerifiedAttributes *[]*string `field:"optional" json:"autoVerifiedAttributes" yaml:"autoVerifiedAttributes"`
	// When active, `DeletionProtection` prevents accidental deletion of your user pool.
	//
	// Before you can delete a user pool that you have protected against deletion, you must deactivate this feature.
	//
	// When you try to delete a protected user pool in a `DeleteUserPool` API request, Amazon Cognito returns an `InvalidParameterException` error. To delete a protected user pool, send a new `DeleteUserPool` request after you deactivate deletion protection in an `UpdateUserPool` API request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deletionprotection
	//
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The device-remembering configuration for a user pool.
	//
	// A null value indicates that you have deactivated device remembering in your user pool.
	//
	// > When you provide a value for any `DeviceConfiguration` field, you activate the Amazon Cognito device-remembering feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deviceconfiguration
	//
	DeviceConfiguration interface{} `field:"optional" json:"deviceConfiguration" yaml:"deviceConfiguration"`
	// The email configuration of your user pool.
	//
	// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailconfiguration
	//
	EmailConfiguration interface{} `field:"optional" json:"emailConfiguration" yaml:"emailConfiguration"`
	// This parameter is no longer used.
	//
	// See [VerificationMessageTemplateType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationmessage
	//
	EmailVerificationMessage *string `field:"optional" json:"emailVerificationMessage" yaml:"emailVerificationMessage"`
	// This parameter is no longer used.
	//
	// See [VerificationMessageTemplateType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationsubject
	//
	EmailVerificationSubject *string `field:"optional" json:"emailVerificationSubject" yaml:"emailVerificationSubject"`
	// Enables MFA on a specified user pool.
	//
	// To disable all MFAs after it has been enabled, set MfaConfiguration to “OFF” and remove EnabledMfas. MFAs can only be all disabled if MfaConfiguration is OFF. Once SMS_MFA is enabled, SMS_MFA can only be disabled by setting MfaConfiguration to “OFF”. Can be one of the following values:
	//
	// - `SMS_MFA` - Enables SMS MFA for the user pool. SMS_MFA can only be enabled if SMS configuration is provided.
	// - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
	//
	// Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-enabledmfas
	//
	EnabledMfas *[]*string `field:"optional" json:"enabledMfas" yaml:"enabledMfas"`
	// The Lambda trigger configuration information for the new user pool.
	//
	// > In a push model, event sources (such as Amazon S3 and custom applications) need permission to invoke a function. So you must make an extra call to add permission for these event sources to invoke your Lambda function.
	// >
	// > For more information on using the Lambda API to add permission, see [AddPermission](https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html) .
	// >
	// > For adding permission using the AWS CLI , see [add-permission](https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-lambdaconfig
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// The multi-factor authentication (MFA) configuration. Valid values include:.
	//
	// - `OFF` MFA won't be used for any users.
	// - `ON` MFA is required for all users to sign in.
	// - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-mfaconfiguration
	//
	MfaConfiguration *string `field:"optional" json:"mfaConfiguration" yaml:"mfaConfiguration"`
	// The policy associated with a user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// The schema attributes for the new user pool. These attributes can be standard or custom attributes.
	//
	// > During a user pool update, you can add new schema attributes but you cannot modify or delete an existing schema attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// A string representing the SMS authentication message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsauthenticationmessage
	//
	SmsAuthenticationMessage *string `field:"optional" json:"smsAuthenticationMessage" yaml:"smsAuthenticationMessage"`
	// The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service.
	//
	// To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsconfiguration
	//
	SmsConfiguration interface{} `field:"optional" json:"smsConfiguration" yaml:"smsConfiguration"`
	// This parameter is no longer used.
	//
	// See [VerificationMessageTemplateType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsverificationmessage
	//
	SmsVerificationMessage *string `field:"optional" json:"smsVerificationMessage" yaml:"smsVerificationMessage"`
	// The settings for updates to user attributes.
	//
	// These settings include the property `AttributesRequireVerificationBeforeUpdate` ,
	// a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
	// more information, see [Verifying updates to email addresses and phone numbers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userattributeupdatesettings
	//
	UserAttributeUpdateSettings interface{} `field:"optional" json:"userAttributeUpdateSettings" yaml:"userAttributeUpdateSettings"`
	// Determines whether email addresses or phone numbers can be specified as user names when a user signs up.
	//
	// Possible values: `phone_number` or `email` .
	//
	// This user pool property cannot be updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameattributes
	//
	UsernameAttributes *[]*string `field:"optional" json:"usernameAttributes" yaml:"usernameAttributes"`
	// You can choose to set case sensitivity on the username input for the selected sign-in option.
	//
	// For example, when this is set to `False` , users will be able to sign in using either "username" or "Username". This configuration is immutable once it has been set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameconfiguration
	//
	UsernameConfiguration interface{} `field:"optional" json:"usernameConfiguration" yaml:"usernameConfiguration"`
	// User pool add-ons.
	//
	// Contains settings for activation of advanced security features. To log user security information but take no action, set to `AUDIT` . To configure automatic security responses to risky traffic to your user pool, set to `ENFORCED` .
	//
	// For more information, see [Adding advanced security to a user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooladdons
	//
	UserPoolAddOns interface{} `field:"optional" json:"userPoolAddOns" yaml:"userPoolAddOns"`
	// A string used to name the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpoolname
	//
	UserPoolName *string `field:"optional" json:"userPoolName" yaml:"userPoolName"`
	// The tag keys and values to assign to the user pool.
	//
	// A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooltags
	//
	UserPoolTags interface{} `field:"optional" json:"userPoolTags" yaml:"userPoolTags"`
	// The template for the verification message that the user sees when the app requests permission to access the user's information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-verificationmessagetemplate
	//
	VerificationMessageTemplate interface{} `field:"optional" json:"verificationMessageTemplate" yaml:"verificationMessageTemplate"`
}

