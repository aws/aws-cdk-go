package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPoolTags interface{}
//
//   cfnUserPool := awscdk.Aws_cognito.NewCfnUserPool(this, jsii.String("MyCfnUserPool"), &CfnUserPoolProps{
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
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
//
type CfnUserPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
	AccountRecoverySetting() interface{}
	SetAccountRecoverySetting(val interface{})
	// The configuration for creating a new user profile.
	AdminCreateUserConfig() interface{}
	SetAdminCreateUserConfig(val interface{})
	// Attributes supported as an alias for this user pool.
	//
	// Possible values: *phone_number* , *email* , or *preferred_username* .
	AliasAttributes() *[]*string
	SetAliasAttributes(val *[]*string)
	// The Amazon Resource Name (ARN) of the user pool, such as `arn:aws:cognito-idp:us-east-1:123412341234:userpool/us-east-1_123412341` .
	AttrArn() *string
	// The provider name of the Amazon Cognito user pool, specified as a `String` .
	AttrProviderName() *string
	// The URL of the provider of the Amazon Cognito user pool, specified as a `String` .
	AttrProviderUrl() *string
	// The ID of the user pool.
	AttrUserPoolId() *string
	// The attributes to be auto-verified.
	AutoVerifiedAttributes() *[]*string
	SetAutoVerifiedAttributes(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// When active, `DeletionProtection` prevents accidental deletion of your user pool.
	DeletionProtection() *string
	SetDeletionProtection(val *string)
	// The device-remembering configuration for a user pool.
	DeviceConfiguration() interface{}
	SetDeviceConfiguration(val interface{})
	// The email configuration of your user pool.
	EmailConfiguration() interface{}
	SetEmailConfiguration(val interface{})
	// This parameter is no longer used.
	EmailVerificationMessage() *string
	SetEmailVerificationMessage(val *string)
	// This parameter is no longer used.
	EmailVerificationSubject() *string
	SetEmailVerificationSubject(val *string)
	// Enables MFA on a specified user pool.
	EnabledMfas() *[]*string
	SetEnabledMfas(val *[]*string)
	// The Lambda trigger configuration information for the new user pool.
	LambdaConfig() interface{}
	SetLambdaConfig(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The multi-factor authentication (MFA) configuration.
	//
	// Valid values include:.
	MfaConfiguration() *string
	SetMfaConfiguration(val *string)
	// The tree node.
	Node() constructs.Node
	// The policy associated with a user pool.
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schema attributes for the new user pool.
	//
	// These attributes can be standard or custom attributes.
	Schema() interface{}
	SetSchema(val interface{})
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage() *string
	SetSmsAuthenticationMessage(val *string)
	// The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service.
	SmsConfiguration() interface{}
	SetSmsConfiguration(val interface{})
	// This parameter is no longer used.
	SmsVerificationMessage() *string
	SetSmsVerificationMessage(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The settings for updates to user attributes.
	UserAttributeUpdateSettings() interface{}
	SetUserAttributeUpdateSettings(val interface{})
	// Determines whether email addresses or phone numbers can be specified as user names when a user signs up.
	UsernameAttributes() *[]*string
	SetUsernameAttributes(val *[]*string)
	// You can choose to set case sensitivity on the username input for the selected sign-in option.
	UsernameConfiguration() interface{}
	SetUsernameConfiguration(val interface{})
	// User pool add-ons.
	UserPoolAddOns() interface{}
	SetUserPoolAddOns(val interface{})
	// A string used to name the user pool.
	UserPoolName() *string
	SetUserPoolName(val *string)
	// The tag keys and values to assign to the user pool.
	UserPoolTagsRaw() interface{}
	SetUserPoolTagsRaw(val interface{})
	// The template for the verification message that the user sees when the app requests permission to access the user's information.
	VerificationMessageTemplate() interface{}
	SetVerificationMessageTemplate(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPool
type jsiiProxy_CfnUserPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnUserPool) AccountRecoverySetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AdminCreateUserConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AttrProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AttrProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProviderUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AttrUserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) DeletionProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) DeviceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) EmailConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) EnabledMfas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMfas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) LambdaConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Schema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) SmsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UserAttributeUpdateSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAttributeUpdateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UsernameConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UserPoolAddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UserPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) UserPoolTagsRaw() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPoolTagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPool) VerificationMessageTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}


func NewCfnUserPool(scope constructs.Construct, id *string, props *CfnUserPoolProps) CfnUserPool {
	_init_.Initialize()

	if err := validateNewCfnUserPoolParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPool{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnUserPool_Override(c CfnUserPool, scope constructs.Construct, id *string, props *CfnUserPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPool)SetAccountRecoverySetting(val interface{}) {
	if err := j.validateSetAccountRecoverySettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountRecoverySetting",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetAdminCreateUserConfig(val interface{}) {
	if err := j.validateSetAdminCreateUserConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminCreateUserConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetAliasAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetAutoVerifiedAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetDeletionProtection(val *string) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetDeviceConfiguration(val interface{}) {
	if err := j.validateSetDeviceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetEmailConfiguration(val interface{}) {
	if err := j.validateSetEmailConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetEmailVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetEmailVerificationSubject(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetEnabledMfas(val *[]*string) {
	_jsii_.Set(
		j,
		"enabledMfas",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetLambdaConfig(val interface{}) {
	if err := j.validateSetLambdaConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetMfaConfiguration(val *string) {
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetPolicies(val interface{}) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetSchema(val interface{}) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetSmsAuthenticationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetSmsConfiguration(val interface{}) {
	if err := j.validateSetSmsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetSmsVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUserAttributeUpdateSettings(val interface{}) {
	if err := j.validateSetUserAttributeUpdateSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAttributeUpdateSettings",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUsernameAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUsernameConfiguration(val interface{}) {
	if err := j.validateSetUsernameConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUserPoolAddOns(val interface{}) {
	if err := j.validateSetUserPoolAddOnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolAddOns",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUserPoolName(val *string) {
	_jsii_.Set(
		j,
		"userPoolName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetUserPoolTagsRaw(val interface{}) {
	_jsii_.Set(
		j,
		"userPoolTagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool)SetVerificationMessageTemplate(val interface{}) {
	if err := j.validateSetVerificationMessageTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationMessageTemplate",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUserPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPool_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnUserPool_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPool_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPool_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.CfnUserPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPool) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPool) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPool) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPool) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPool) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPool) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPool) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPool) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPool) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnUserPool) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

