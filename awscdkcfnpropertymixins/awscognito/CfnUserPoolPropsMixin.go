package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPool` resource creates an Amazon Cognito user pool.
//
// For more information on working with Amazon Cognito user pools, see [Amazon Cognito User Pools](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html) and [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) .
//
// > If you don't specify a value for a parameter, Amazon Cognito sets it to a default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var userPoolTags interface{}
//
//   cfnUserPoolPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnUserPoolPropsMixin(&CfnUserPoolMixinProps{
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
//   	EmailAuthenticationMessage: jsii.String("emailAuthenticationMessage"),
//   	EmailAuthenticationSubject: jsii.String("emailAuthenticationSubject"),
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
//   		InboundFederation: &InboundFederationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			LambdaVersion: jsii.String("lambdaVersion"),
//   		},
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
//   			PasswordHistorySize: jsii.Number(123),
//   			RequireLowercase: jsii.Boolean(false),
//   			RequireNumbers: jsii.Boolean(false),
//   			RequireSymbols: jsii.Boolean(false),
//   			RequireUppercase: jsii.Boolean(false),
//   			TemporaryPasswordValidityDays: jsii.Number(123),
//   		},
//   		SignInPolicy: &SignInPolicyProperty{
//   			AllowedFirstAuthFactors: []*string{
//   				jsii.String("allowedFirstAuthFactors"),
//   			},
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
//   		AdvancedSecurityAdditionalFlows: &AdvancedSecurityAdditionalFlowsProperty{
//   			CustomAuthMode: jsii.String("customAuthMode"),
//   		},
//   		AdvancedSecurityMode: jsii.String("advancedSecurityMode"),
//   	},
//   	UserPoolName: jsii.String("userPoolName"),
//   	UserPoolTags: userPoolTags,
//   	UserPoolTier: jsii.String("userPoolTier"),
//   	VerificationMessageTemplate: &VerificationMessageTemplateProperty{
//   		DefaultEmailOption: jsii.String("defaultEmailOption"),
//   		EmailMessage: jsii.String("emailMessage"),
//   		EmailMessageByLink: jsii.String("emailMessageByLink"),
//   		EmailSubject: jsii.String("emailSubject"),
//   		EmailSubjectByLink: jsii.String("emailSubjectByLink"),
//   		SmsMessage: jsii.String("smsMessage"),
//   	},
//   	WebAuthnRelyingPartyId: jsii.String("webAuthnRelyingPartyId"),
//   	WebAuthnUserVerification: jsii.String("webAuthnUserVerification"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
//
type CfnUserPoolPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserPoolMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolPropsMixin
type jsiiProxy_CfnUserPoolPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserPoolPropsMixin) Props() *CfnUserPoolMixinProps {
	var returns *CfnUserPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPool`.
func NewCfnUserPoolPropsMixin(props *CfnUserPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPool`.
func NewCfnUserPoolPropsMixin_Override(c CfnUserPoolPropsMixin, props *CfnUserPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

