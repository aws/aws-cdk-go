package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// How will a user be able to recover their account?
//
// When a user forgets their password, they can have a code sent to their verified email or verified phone to recover their account.
// You can choose the preferred way to send codes below.
// We recommend not allowing phone to be used for both password resets and multi-factor authentication (MFA).
//
// Example:
//   cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   	// ...
//   	accountRecovery: cognito.accountRecovery_EMAIL_ONLY,
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-recover-a-user-account.html
//
// Experimental.
type AccountRecovery string

const (
	// Email if available, otherwise phone, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	// Experimental.
	AccountRecovery_EMAIL_AND_PHONE_WITHOUT_MFA AccountRecovery = "EMAIL_AND_PHONE_WITHOUT_MFA"
	// Phone if available, otherwise email, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	// Experimental.
	AccountRecovery_PHONE_WITHOUT_MFA_AND_EMAIL AccountRecovery = "PHONE_WITHOUT_MFA_AND_EMAIL"
	// Email only.
	// Experimental.
	AccountRecovery_EMAIL_ONLY AccountRecovery = "EMAIL_ONLY"
	// Phone only, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	// Experimental.
	AccountRecovery_PHONE_ONLY_WITHOUT_MFA AccountRecovery = "PHONE_ONLY_WITHOUT_MFA"
	// (Not Recommended) Phone if available, otherwise email, and do allow a user to reset their password via phone if they are also using it for MFA.
	// Experimental.
	AccountRecovery_PHONE_AND_EMAIL AccountRecovery = "PHONE_AND_EMAIL"
	// None – users will have to contact an administrator to reset their passwords.
	// Experimental.
	AccountRecovery_NONE AccountRecovery = "NONE"
)

// The mapping of user pool attributes to the attributes provided by the identity providers.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   	userPool: userpool,
//   	attributeMapping: &attributeMapping{
//   		email: cognito.providerAttribute_AMAZON_EMAIL(),
//   		website: cognito.*providerAttribute.other(jsii.String("url")),
//   		 // use other() when an attribute is not pre-defined in the CDK
//   		custom: map[string]*providerAttribute{
//   			// custom user pool attributes go here
//   			"uniqueId": cognito.*providerAttribute_AMAZON_USER_ID(),
//   		},
//   	},
//   })
//
// Experimental.
type AttributeMapping struct {
	// The user's postal address is a required attribute.
	// Experimental.
	Address ProviderAttribute `json:"address" yaml:"address"`
	// The user's birthday.
	// Experimental.
	Birthdate ProviderAttribute `json:"birthdate" yaml:"birthdate"`
	// Specify custom attribute mapping here and mapping for any standard attributes not supported yet.
	// Experimental.
	Custom *map[string]ProviderAttribute `json:"custom" yaml:"custom"`
	// The user's e-mail address.
	// Experimental.
	Email ProviderAttribute `json:"email" yaml:"email"`
	// The surname or last name of user.
	// Experimental.
	FamilyName ProviderAttribute `json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form.
	// Experimental.
	Fullname ProviderAttribute `json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Experimental.
	Gender ProviderAttribute `json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Experimental.
	GivenName ProviderAttribute `json:"givenName" yaml:"givenName"`
	// Time, the user's information was last updated.
	// Experimental.
	LastUpdateTime ProviderAttribute `json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale.
	// Experimental.
	Locale ProviderAttribute `json:"locale" yaml:"locale"`
	// The user's middle name.
	// Experimental.
	MiddleName ProviderAttribute `json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Experimental.
	Nickname ProviderAttribute `json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Experimental.
	PhoneNumber ProviderAttribute `json:"phoneNumber" yaml:"phoneNumber"`
	// The user's preferred username.
	// Experimental.
	PreferredUsername ProviderAttribute `json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Experimental.
	ProfilePage ProviderAttribute `json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Experimental.
	ProfilePicture ProviderAttribute `json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Experimental.
	Timezone ProviderAttribute `json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Experimental.
	Website ProviderAttribute `json:"website" yaml:"website"`
}

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
	AdminUserPassword *bool `json:"adminUserPassword" yaml:"adminUserPassword"`
	// Enable custom authentication flow.
	// Experimental.
	Custom *bool `json:"custom" yaml:"custom"`
	// Enable auth using username & password.
	// Experimental.
	UserPassword *bool `json:"userPassword" yaml:"userPassword"`
	// Enable SRP based authentication.
	// Experimental.
	UserSrp *bool `json:"userSrp" yaml:"userSrp"`
}

// Attributes that can be automatically verified for users in a user pool.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	// ...
//   	signInAliases: &signInAliases{
//   		username: jsii.Boolean(true),
//   		email: jsii.Boolean(true),
//   	},
//   	autoVerify: &autoVerifiedAttrs{
//   		email: jsii.Boolean(true),
//   		phone: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type AutoVerifiedAttrs struct {
	// Whether the email address of the user should be auto verified at sign up.
	//
	// Note: If both `email` and `phone` is set, Cognito only verifies the phone number. To also verify email, see here -
	// https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html
	// Experimental.
	Email *bool `json:"email" yaml:"email"`
	// Whether the phone number of the user should be auto verified at sign up.
	// Experimental.
	Phone *bool `json:"phone" yaml:"phone"`
}

// The Boolean custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type BooleanAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	// Experimental.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for BooleanAttribute
type jsiiProxy_BooleanAttribute struct {
	jsiiProxy_ICustomAttribute
}

// Experimental.
func NewBooleanAttribute(props *CustomAttributeProps) BooleanAttribute {
	_init_.Initialize()

	j := jsiiProxy_BooleanAttribute{}

	_jsii_.Create(
		"monocdk.aws_cognito.BooleanAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBooleanAttribute_Override(b BooleanAttribute, props *CustomAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.BooleanAttribute",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_BooleanAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		b,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::Cognito::IdentityPool`.
//
// The `AWS::Cognito::IdentityPool` resource creates an Amazon Cognito identity pool.
//
// To avoid deleting the resource accidentally from AWS CloudFormation , use [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) and the [UpdateReplacePolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) to retain the resource on deletion or replacement.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var myProvider openIdConnectProvider
//   cognito.NewCfnIdentityPool(this, jsii.String("IdentityPool"), &cfnIdentityPoolProps{
//   	openIdConnectProviderArns: []*string{
//   		myProvider.openIdConnectProviderArn,
//   	},
//   	// And the other properties for your identity pool
//   	allowUnauthenticatedIdentities: jsii.Boolean(false),
//   })
//
type CfnIdentityPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Enables the Basic (Classic) authentication flow.
	AllowClassicFlow() interface{}
	SetAllowClassicFlow(val interface{})
	// Specifies whether the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities() interface{}
	SetAllowUnauthenticatedIdentities(val interface{})
	// The name of the Amazon Cognito identity pool, returned as a string.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The events to configure.
	CognitoEvents() interface{}
	SetCognitoEvents(val interface{})
	// The Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders() interface{}
	SetCognitoIdentityProviders(val interface{})
	// Configuration options for configuring Amazon Cognito streams.
	CognitoStreams() interface{}
	SetCognitoStreams(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The "domain" Amazon Cognito uses when referencing your users.
	//
	// This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 100.
	DeveloperProviderName() *string
	SetDeveloperProviderName(val *string)
	// The name of your Amazon Cognito identity pool.
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 128
	//
	// *Pattern* : `[\w\s+=,.@-]+`
	IdentityPoolName() *string
	SetIdentityPoolName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Amazon Resource Names (ARNs) of the OpenID connect providers.
	OpenIdConnectProviderArns() *[]*string
	SetOpenIdConnectProviderArns(val *[]*string)
	// The configuration options to be applied to the identity pool.
	PushSync() interface{}
	SetPushSync(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
	SamlProviderArns() *[]*string
	SetSamlProviderArns(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key-value pairs that map provider names to provider app IDs.
	SupportedLoginProviders() interface{}
	SetSupportedLoginProviders(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnIdentityPool
type jsiiProxy_CfnIdentityPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIdentityPool) AllowClassicFlow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) AllowUnauthenticatedIdentities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CognitoEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CognitoIdentityProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CognitoStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) DeveloperProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) OpenIdConnectProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openIdConnectProviderArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) PushSync() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) SamlProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) SupportedLoginProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedLoginProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPool) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::IdentityPool`.
func NewCfnIdentityPool(scope awscdk.Construct, id *string, props *CfnIdentityPoolProps) CfnIdentityPool {
	_init_.Initialize()

	j := jsiiProxy_CfnIdentityPool{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnIdentityPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::IdentityPool`.
func NewCfnIdentityPool_Override(c CfnIdentityPool, scope awscdk.Construct, id *string, props *CfnIdentityPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnIdentityPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetAllowClassicFlow(val interface{}) {
	_jsii_.Set(
		j,
		"allowClassicFlow",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetAllowUnauthenticatedIdentities(val interface{}) {
	_jsii_.Set(
		j,
		"allowUnauthenticatedIdentities",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetCognitoEvents(val interface{}) {
	_jsii_.Set(
		j,
		"cognitoEvents",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetCognitoIdentityProviders(val interface{}) {
	_jsii_.Set(
		j,
		"cognitoIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetCognitoStreams(val interface{}) {
	_jsii_.Set(
		j,
		"cognitoStreams",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetDeveloperProviderName(val *string) {
	_jsii_.Set(
		j,
		"developerProviderName",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetIdentityPoolName(val *string) {
	_jsii_.Set(
		j,
		"identityPoolName",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetOpenIdConnectProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"openIdConnectProviderArns",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetPushSync(val interface{}) {
	_jsii_.Set(
		j,
		"pushSync",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetSamlProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"samlProviderArns",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPool) SetSupportedLoginProviders(val interface{}) {
	_jsii_.Set(
		j,
		"supportedLoginProviders",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnIdentityPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIdentityPool_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPool",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnIdentityPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityPool_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnIdentityPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityPool) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIdentityPool) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIdentityPool) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIdentityPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIdentityPool) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIdentityPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIdentityPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIdentityPool) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIdentityPool) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityPool) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIdentityPool) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPool) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `CognitoIdentityProvider` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that represents an Amazon Cognito user pool and its client ID.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cognitoIdentityProviderProperty := &cognitoIdentityProviderProperty{
//   	clientId: jsii.String("clientId"),
//   	providerName: jsii.String("providerName"),
//   	serverSideTokenCheck: jsii.Boolean(false),
//   }
//
type CfnIdentityPool_CognitoIdentityProviderProperty struct {
	// The client ID for the Amazon Cognito user pool.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The provider name for an Amazon Cognito user pool.
	//
	// For example: `cognito-idp.us-east-2.amazonaws.com/us-east-2_123456789` .
	ProviderName *string `json:"providerName" yaml:"providerName"`
	// TRUE if server-side token validation is enabled for the identity provider’s token.
	//
	// After you set the `ServerSideTokenCheck` to TRUE for an identity pool, that identity pool checks with the integrated user pools to make sure the user has not been globally signed out or deleted before the identity pool provides an OIDC token or AWS credentials for the user.
	//
	// If the user is signed out or deleted, the identity pool returns a 400 Not Authorized error.
	ServerSideTokenCheck interface{} `json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

// `CognitoStreams` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines configuration options for Amazon Cognito streams.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cognitoStreamsProperty := &cognitoStreamsProperty{
//   	roleArn: jsii.String("roleArn"),
//   	streamingStatus: jsii.String("streamingStatus"),
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnIdentityPool_CognitoStreamsProperty struct {
	// The Amazon Resource Name (ARN) of the role Amazon Cognito can assume to publish to the stream.
	//
	// This role must grant access to Amazon Cognito (cognito-sync) to invoke `PutRecord` on your Amazon Cognito stream.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Status of the Amazon Cognito streams.
	//
	// Valid values are: `ENABLED` or `DISABLED` .
	StreamingStatus *string `json:"streamingStatus" yaml:"streamingStatus"`
	// The name of the Amazon Cognito stream to receive updates.
	//
	// This stream must be in the developer's account and in the same Region as the identity pool.
	StreamName *string `json:"streamName" yaml:"streamName"`
}

// `PushSync` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines the configuration options to be applied to an Amazon Cognito identity pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   pushSyncProperty := &pushSyncProperty{
//   	applicationArns: []*string{
//   		jsii.String("applicationArns"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnIdentityPool_PushSyncProperty struct {
	// The ARNs of the Amazon SNS platform applications that could be used by clients.
	ApplicationArns *[]*string `json:"applicationArns" yaml:"applicationArns"`
	// An IAM role configured to allow Amazon Cognito to call Amazon SNS on behalf of the developer.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Properties for defining a `CfnIdentityPool`.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var myProvider openIdConnectProvider
//   cognito.NewCfnIdentityPool(this, jsii.String("IdentityPool"), &cfnIdentityPoolProps{
//   	openIdConnectProviderArns: []*string{
//   		myProvider.openIdConnectProviderArn,
//   	},
//   	// And the other properties for your identity pool
//   	allowUnauthenticatedIdentities: jsii.Boolean(false),
//   })
//
type CfnIdentityPoolProps struct {
	// Specifies whether the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities interface{} `json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// Enables the Basic (Classic) authentication flow.
	AllowClassicFlow interface{} `json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// The events to configure.
	CognitoEvents interface{} `json:"cognitoEvents" yaml:"cognitoEvents"`
	// The Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders interface{} `json:"cognitoIdentityProviders" yaml:"cognitoIdentityProviders"`
	// Configuration options for configuring Amazon Cognito streams.
	CognitoStreams interface{} `json:"cognitoStreams" yaml:"cognitoStreams"`
	// The "domain" Amazon Cognito uses when referencing your users.
	//
	// This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 100.
	DeveloperProviderName *string `json:"developerProviderName" yaml:"developerProviderName"`
	// The name of your Amazon Cognito identity pool.
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 128
	//
	// *Pattern* : `[\w\s+=,.@-]+`
	IdentityPoolName *string `json:"identityPoolName" yaml:"identityPoolName"`
	// The Amazon Resource Names (ARNs) of the OpenID connect providers.
	OpenIdConnectProviderArns *[]*string `json:"openIdConnectProviderArns" yaml:"openIdConnectProviderArns"`
	// The configuration options to be applied to the identity pool.
	PushSync interface{} `json:"pushSync" yaml:"pushSync"`
	// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
	SamlProviderArns *[]*string `json:"samlProviderArns" yaml:"samlProviderArns"`
	// Key-value pairs that map provider names to provider app IDs.
	SupportedLoginProviders interface{} `json:"supportedLoginProviders" yaml:"supportedLoginProviders"`
}

// A CloudFormation `AWS::Cognito::IdentityPoolRoleAttachment`.
//
// The `AWS::Cognito::IdentityPoolRoleAttachment` resource manages the role configuration for an Amazon Cognito identity pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var roles interface{}
//   cfnIdentityPoolRoleAttachment := cognito.NewCfnIdentityPoolRoleAttachment(this, jsii.String("MyCfnIdentityPoolRoleAttachment"), &cfnIdentityPoolRoleAttachmentProps{
//   	identityPoolId: jsii.String("identityPoolId"),
//
//   	// the properties below are optional
//   	roleMappings: map[string]interface{}{
//   		"roleMappingsKey": &RoleMappingProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"ambiguousRoleResolution": jsii.String("ambiguousRoleResolution"),
//   			"identityProvider": jsii.String("identityProvider"),
//   			"rulesConfiguration": &RulesConfigurationTypeProperty{
//   				"rules": []interface{}{
//   					&MappingRuleProperty{
//   						"claim": jsii.String("claim"),
//   						"matchType": jsii.String("matchType"),
//   						"roleArn": jsii.String("roleArn"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roles: roles,
//   })
//
type CfnIdentityPoolRoleAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// An identity pool ID in the format `REGION:GUID` .
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// How users for a specific identity provider are mapped to roles.
	//
	// This is a string to the `RoleMapping` object map. The string identifies the identity provider. For example: `graph.facebook.com` or `cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id` .
	//
	// If the `IdentityProvider` field isn't provided in this object, the string is used as the identity provider name.
	//
	// For more information, see the [RoleMapping property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) .
	RoleMappings() interface{}
	SetRoleMappings(val interface{})
	// The map of the roles associated with this pool.
	//
	// For a given role, the key is either "authenticated" or "unauthenticated". The value is the role ARN.
	Roles() interface{}
	SetRoles(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnIdentityPoolRoleAttachment
type jsiiProxy_CfnIdentityPoolRoleAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) RoleMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) Roles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::IdentityPoolRoleAttachment`.
func NewCfnIdentityPoolRoleAttachment(scope awscdk.Construct, id *string, props *CfnIdentityPoolRoleAttachmentProps) CfnIdentityPoolRoleAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnIdentityPoolRoleAttachment{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::IdentityPoolRoleAttachment`.
func NewCfnIdentityPoolRoleAttachment_Override(c CfnIdentityPoolRoleAttachment, scope awscdk.Construct, id *string, props *CfnIdentityPoolRoleAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) SetIdentityPoolId(val *string) {
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) SetRoleMappings(val interface{}) {
	_jsii_.Set(
		j,
		"roleMappings",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachment) SetRoles(val interface{}) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnIdentityPoolRoleAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIdentityPoolRoleAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnIdentityPoolRoleAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityPoolRoleAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnIdentityPoolRoleAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Defines how to map a claim to a role ARN.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   mappingRuleProperty := &mappingRuleProperty{
//   	claim: jsii.String("claim"),
//   	matchType: jsii.String("matchType"),
//   	roleArn: jsii.String("roleArn"),
//   	value: jsii.String("value"),
//   }
//
type CfnIdentityPoolRoleAttachment_MappingRuleProperty struct {
	// The claim name that must be present in the token.
	//
	// For example: "isAdmin" or "paid".
	Claim *string `json:"claim" yaml:"claim"`
	// The match condition that specifies how closely the claim value in the IdP token must match `Value` .
	//
	// Valid values are: `Equals` , `Contains` , `StartsWith` , and `NotEqual` .
	MatchType *string `json:"matchType" yaml:"matchType"`
	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A brief string that the claim must match.
	//
	// For example, "paid" or "yes".
	Value *string `json:"value" yaml:"value"`
}

// `RoleMapping` is a property of the [AWS::Cognito::IdentityPoolRoleAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html) resource that defines the role-mapping attributes of an Amazon Cognito identity pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   roleMappingProperty := &roleMappingProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ambiguousRoleResolution: jsii.String("ambiguousRoleResolution"),
//   	identityProvider: jsii.String("identityProvider"),
//   	rulesConfiguration: &rulesConfigurationTypeProperty{
//   		rules: []interface{}{
//   			&mappingRuleProperty{
//   				claim: jsii.String("claim"),
//   				matchType: jsii.String("matchType"),
//   				roleArn: jsii.String("roleArn"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnIdentityPoolRoleAttachment_RoleMappingProperty struct {
	// The role-mapping type.
	//
	// `Token` uses `cognito:roles` and `cognito:preferred_role` claims from the Amazon Cognito identity provider token to map groups to roles. `Rules` attempts to match claims from the token to map to a role.
	//
	// Valid values are `Token` or `Rules` .
	Type *string `json:"type" yaml:"type"`
	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no `cognito:preferred_role` claim and there are multiple `cognito:roles` matches for the Token type.
	//
	// If you specify Token or Rules as the Type, AmbiguousRoleResolution is required.
	//
	// Valid values are `AuthenticatedRole` or `Deny` .
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution" yaml:"ambiguousRoleResolution"`
	// Identifier for the identity provider for which the role is mapped.
	//
	// For example: `graph.facebook.com` or `cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id (http://cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id)` . This is the identity provider that is used by the user for authentication.
	//
	// If the identity provider property isn't provided, the key of the entry in the `RoleMappings` map is used as the identity provider.
	IdentityProvider *string `json:"identityProvider" yaml:"identityProvider"`
	// The rules to be used for mapping users to roles.
	//
	// If you specify "Rules" as the role-mapping type, RulesConfiguration is required.
	RulesConfiguration interface{} `json:"rulesConfiguration" yaml:"rulesConfiguration"`
}

// `RulesConfigurationType` is a subproperty of the [RoleMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) property that defines the rules to be used for mapping users to roles.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   rulesConfigurationTypeProperty := &rulesConfigurationTypeProperty{
//   	rules: []interface{}{
//   		&mappingRuleProperty{
//   			claim: jsii.String("claim"),
//   			matchType: jsii.String("matchType"),
//   			roleArn: jsii.String("roleArn"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityPoolRoleAttachment_RulesConfigurationTypeProperty struct {
	// The rules.
	//
	// You can specify up to 25 rules per identity provider.
	Rules interface{} `json:"rules" yaml:"rules"`
}

// Properties for defining a `CfnIdentityPoolRoleAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var roles interface{}
//   cfnIdentityPoolRoleAttachmentProps := &cfnIdentityPoolRoleAttachmentProps{
//   	identityPoolId: jsii.String("identityPoolId"),
//
//   	// the properties below are optional
//   	roleMappings: map[string]interface{}{
//   		"roleMappingsKey": &RoleMappingProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"ambiguousRoleResolution": jsii.String("ambiguousRoleResolution"),
//   			"identityProvider": jsii.String("identityProvider"),
//   			"rulesConfiguration": &RulesConfigurationTypeProperty{
//   				"rules": []interface{}{
//   					&MappingRuleProperty{
//   						"claim": jsii.String("claim"),
//   						"matchType": jsii.String("matchType"),
//   						"roleArn": jsii.String("roleArn"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roles: roles,
//   }
//
type CfnIdentityPoolRoleAttachmentProps struct {
	// An identity pool ID in the format `REGION:GUID` .
	IdentityPoolId *string `json:"identityPoolId" yaml:"identityPoolId"`
	// How users for a specific identity provider are mapped to roles.
	//
	// This is a string to the `RoleMapping` object map. The string identifies the identity provider. For example: `graph.facebook.com` or `cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id` .
	//
	// If the `IdentityProvider` field isn't provided in this object, the string is used as the identity provider name.
	//
	// For more information, see the [RoleMapping property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) .
	RoleMappings interface{} `json:"roleMappings" yaml:"roleMappings"`
	// The map of the roles associated with this pool.
	//
	// For a given role, the key is either "authenticated" or "unauthenticated". The value is the role ARN.
	Roles interface{} `json:"roles" yaml:"roles"`
}

// A CloudFormation `AWS::Cognito::UserPool`.
//
// The `AWS::Cognito::UserPool` resource creates an Amazon Cognito user pool. For more information on working with Amazon Cognito user pools, see [Amazon Cognito User Pools](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html) and [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var userPoolTags interface{}
//   cfnUserPool := cognito.NewCfnUserPool(this, jsii.String("MyCfnUserPool"), &cfnUserPoolProps{
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
//   })
//
type CfnUserPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
	//
	// It allows you to define a preferred method when a user has more than one method available. With this setting, SMS does not qualify for a valid password recovery mechanism if the user also has SMS MFA enabled. In the absence of this setting, Cognito uses the legacy behavior to determine the recovery method where SMS is preferred over email.
	AccountRecoverySetting() interface{}
	SetAccountRecoverySetting(val interface{})
	// The configuration for creating a new user profile.
	AdminCreateUserConfig() interface{}
	SetAdminCreateUserConfig(val interface{})
	// Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
	//
	// > This user pool property cannot be updated.
	AliasAttributes() *[]*string
	SetAliasAttributes(val *[]*string)
	// The Amazon Resource Name (ARN) of the user pool, such as `arn:aws:cognito-idp:us-east-1:123412341234:userpool/us-east-1_123412341` .
	AttrArn() *string
	// The provider name of the Amazon Cognito user pool, specified as a `String` .
	AttrProviderName() *string
	// The URL of the provider of the Amazon Cognito user pool, specified as a `String` .
	AttrProviderUrl() *string
	// The attributes to be auto-verified.
	//
	// Possible values: *email* , *phone_number* .
	AutoVerifiedAttributes() *[]*string
	SetAutoVerifiedAttributes(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The device configuration.
	DeviceConfiguration() interface{}
	SetDeviceConfiguration(val interface{})
	// The email configuration of your user pool.
	//
	// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
	EmailConfiguration() interface{}
	SetEmailConfiguration(val interface{})
	// A string representing the email verification message.
	//
	// EmailVerificationMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationMessage() *string
	SetEmailVerificationMessage(val *string)
	// A string representing the email verification subject.
	//
	// EmailVerificationSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationSubject() *string
	SetEmailVerificationSubject(val *string)
	// Enables MFA on a specified user pool.
	//
	// To disable all MFAs after it has been enabled, set MfaConfiguration to “OFF” and remove EnabledMfas. MFAs can only be all disabled if MfaConfiguration is OFF. Once SMS_MFA is enabled, SMS_MFA can only be disabled by setting MfaConfiguration to “OFF”. Can be one of the following values:
	//
	// - `SMS_MFA` - Enables SMS MFA for the user pool. SMS_MFA can only be enabled if SMS configuration is provided.
	// - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
	//
	// Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA`.
	EnabledMfas() *[]*string
	SetEnabledMfas(val *[]*string)
	// The Lambda trigger configuration information for the new user pool.
	//
	// > In a push model, event sources (such as Amazon S3 and custom applications) need permission to invoke a function. So you must make an extra call to add permission for these event sources to invoke your Lambda function.
	// >
	// > For more information on using the Lambda API to add permission, see [AddPermission](https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html) .
	// >
	// > For adding permission using the AWS CLI , see [add-permission](https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html) .
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
	// Experimental.
	LogicalId() *string
	// The multi-factor (MFA) configuration. Valid values include:.
	//
	// - `OFF` MFA won't be used for any users.
	// - `ON` MFA is required for all users to sign in.
	// - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
	MfaConfiguration() *string
	SetMfaConfiguration(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The policy associated with a user pool.
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The schema attributes for the new user pool. These attributes can be standard or custom attributes.
	//
	// > During a user pool update, you can add new schema attributes but you cannot modify or delete an existing schema attribute.
	Schema() interface{}
	SetSchema(val interface{})
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage() *string
	SetSmsAuthenticationMessage(val *string)
	// The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service.
	//
	// To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
	SmsConfiguration() interface{}
	SetSmsConfiguration(val interface{})
	// A string representing the SMS verification message.
	SmsVerificationMessage() *string
	SetSmsVerificationMessage(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tag keys and values to assign to the user pool.
	//
	// A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Determines whether email addresses or phone numbers can be specified as user names when a user signs up.
	//
	// Possible values: `phone_number` or `email` .
	//
	// This user pool property cannot be updated.
	UsernameAttributes() *[]*string
	SetUsernameAttributes(val *[]*string)
	// You can choose to set case sensitivity on the username input for the selected sign-in option.
	//
	// For example, when this is set to `False` , users will be able to sign in using either "username" or "Username". This configuration is immutable once it has been set.
	UsernameConfiguration() interface{}
	SetUsernameConfiguration(val interface{})
	// Enables advanced security risk detection.
	//
	// Set the key `AdvancedSecurityMode` to the value "AUDIT".
	UserPoolAddOns() interface{}
	SetUserPoolAddOns(val interface{})
	// A string used to name the user pool.
	UserPoolName() *string
	SetUserPoolName(val *string)
	// The template for the verification message that the user sees when the app requests permission to access the user's information.
	VerificationMessageTemplate() interface{}
	SetVerificationMessageTemplate(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPool
type jsiiProxy_CfnUserPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnUserPool) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_CfnUserPool) VerificationMessageTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPool`.
func NewCfnUserPool(scope awscdk.Construct, id *string, props *CfnUserPoolProps) CfnUserPool {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPool{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPool`.
func NewCfnUserPool_Override(c CfnUserPool, scope awscdk.Construct, id *string, props *CfnUserPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPool) SetAccountRecoverySetting(val interface{}) {
	_jsii_.Set(
		j,
		"accountRecoverySetting",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetAdminCreateUserConfig(val interface{}) {
	_jsii_.Set(
		j,
		"adminCreateUserConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetAliasAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetAutoVerifiedAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetDeviceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"deviceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetEmailConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"emailConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetEmailVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetEmailVerificationSubject(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetEnabledMfas(val *[]*string) {
	_jsii_.Set(
		j,
		"enabledMfas",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetLambdaConfig(val interface{}) {
	_jsii_.Set(
		j,
		"lambdaConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetMfaConfiguration(val *string) {
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetSchema(val interface{}) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetSmsAuthenticationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetSmsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"smsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetSmsVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetUsernameAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetUsernameConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"usernameConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetUserPoolAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"userPoolAddOns",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetUserPoolName(val *string) {
	_jsii_.Set(
		j,
		"userPoolName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPool) SetVerificationMessageTemplate(val interface{}) {
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
// Experimental.
func CfnUserPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPool_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPool",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPool",
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
		"monocdk.aws_cognito.CfnUserPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPool) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPool) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPool) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPool) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPool) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPool) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPool) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPool) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnUserPool) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnUserPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPool) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
//
// It allows you to define a preferred method when a user has more than one method available. With this setting, SMS does not qualify for a valid password recovery mechanism if the user also has SMS MFA enabled. In the absence of this setting, Cognito uses the legacy behavior to determine the recovery method where SMS is preferred over email.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   accountRecoverySettingProperty := &accountRecoverySettingProperty{
//   	recoveryMechanisms: []interface{}{
//   		&recoveryOptionProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnUserPool_AccountRecoverySettingProperty struct {
	// The list of `RecoveryOptionTypes` .
	RecoveryMechanisms interface{} `json:"recoveryMechanisms" yaml:"recoveryMechanisms"`
}

// The configuration for `AdminCreateUser` requests.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   adminCreateUserConfigProperty := &adminCreateUserConfigProperty{
//   	allowAdminCreateUserOnly: jsii.Boolean(false),
//   	inviteMessageTemplate: &inviteMessageTemplateProperty{
//   		emailMessage: jsii.String("emailMessage"),
//   		emailSubject: jsii.String("emailSubject"),
//   		smsMessage: jsii.String("smsMessage"),
//   	},
//   	unusedAccountValidityDays: jsii.Number(123),
//   }
//
type CfnUserPool_AdminCreateUserConfigProperty struct {
	// Set to `True` if only the administrator is allowed to create user profiles.
	//
	// Set to `False` if users can sign themselves up via an app.
	AllowAdminCreateUserOnly interface{} `json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// The message template to be used for the welcome message to new users.
	//
	// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
	InviteMessageTemplate interface{} `json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
	// The user account expiration limit, in days, after which the account is no longer usable.
	//
	// To reset the account after that time limit, you must call `AdminCreateUser` again, specifying `"RESEND"` for the `MessageAction` parameter. The default value for this parameter is 7.
	//
	// > If you set a value for `TemporaryPasswordValidityDays` in `PasswordPolicy` , that value will be used, and `UnusedAccountValidityDays` will be no longer be an available parameter for that user pool.
	UnusedAccountValidityDays *float64 `json:"unusedAccountValidityDays" yaml:"unusedAccountValidityDays"`
}

// A custom email sender AWS Lambda trigger.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   customEmailSenderProperty := &customEmailSenderProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	lambdaVersion: jsii.String("lambdaVersion"),
//   }
//
type CfnUserPool_CustomEmailSenderProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon Cognito triggers to send email notifications to users.
	LambdaArn *string `json:"lambdaArn" yaml:"lambdaArn"`
	// The Lambda version represents the signature of the "request" attribute in the "event" information that Amazon Cognito passes to your custom email sender AWS Lambda function.
	//
	// The only supported value is `V1_0` .
	LambdaVersion *string `json:"lambdaVersion" yaml:"lambdaVersion"`
}

// A custom SMS sender AWS Lambda trigger.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   customSMSSenderProperty := &customSMSSenderProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	lambdaVersion: jsii.String("lambdaVersion"),
//   }
//
type CfnUserPool_CustomSMSSenderProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon Cognito triggers to send SMS notifications to users.
	LambdaArn *string `json:"lambdaArn" yaml:"lambdaArn"`
	// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS sender Lambda function.
	//
	// The only supported value is `V1_0` .
	LambdaVersion *string `json:"lambdaVersion" yaml:"lambdaVersion"`
}

// The device tracking configuration for a user pool. A user pool with device tracking deactivated returns a null value.
//
// > When you provide values for any DeviceConfiguration field, you activate device tracking.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   deviceConfigurationProperty := &deviceConfigurationProperty{
//   	challengeRequiredOnNewDevice: jsii.Boolean(false),
//   	deviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   }
//
type CfnUserPool_DeviceConfigurationProperty struct {
	// When true, device authentication can replace SMS and time-based one-time password (TOTP) factors for multi-factor authentication (MFA).
	//
	// > Users that sign in with devices that have not been confirmed or remembered will still have to provide a second factor, whether or not ChallengeRequiredOnNewDevice is true, when your user pool requires MFA.
	ChallengeRequiredOnNewDevice interface{} `json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// When true, users can opt in to remembering their device.
	//
	// Your app code must use callback functions to return the user's choice.
	DeviceOnlyRememberedOnUserPrompt interface{} `json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

// The email configuration of your user pool.
//
// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   emailConfigurationProperty := &emailConfigurationProperty{
//   	configurationSet: jsii.String("configurationSet"),
//   	emailSendingAccount: jsii.String("emailSendingAccount"),
//   	from: jsii.String("from"),
//   	replyToEmailAddress: jsii.String("replyToEmailAddress"),
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type CfnUserPool_EmailConfigurationProperty struct {
	// The set of configuration rules that can be applied to emails sent using Amazon SES.
	//
	// A configuration set is applied to an email by including a reference to the configuration set in the headers of the email. Once applied, all of the rules in that configuration set are applied to the email. Configuration sets can be used to apply the following types of rules to emails:
	//
	// - Event publishing – Amazon SES can track the number of send, delivery, open, click, bounce, and complaint events for each email sent. Use event publishing to send information about these events to other AWS services such as SNS and CloudWatch.
	// - IP pool management – When leasing dedicated IP addresses with Amazon SES, you can create groups of IP addresses, called dedicated IP pools. You can then associate the dedicated IP pools with configuration sets.
	ConfigurationSet *string `json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether Amazon Cognito uses its built-in functionality to send your users email messages, or uses your Amazon Simple Email Service email configuration.
	//
	// Specify one of the following values:
	//
	// - **COGNITO_DEFAULT** - When Amazon Cognito emails your users, it uses its built-in email functionality. When you use the default option, Amazon Cognito allows only a limited number of emails each day for your user pool. For typical production environments, the default email limit is less than the required delivery volume. To achieve a higher delivery volume, specify DEVELOPER to use your Amazon SES email configuration.
	//
	// To look up the email delivery limit for the default option, see [Limits in](https://docs.aws.amazon.com/cognito/latest/developerguide/limits.html) in the *Developer Guide* .
	//
	// The default FROM address is `no-reply@verificationemail.com` . To customize the FROM address, provide the Amazon Resource Name (ARN) of an Amazon SES verified email address for the `SourceArn` parameter.
	//
	// If EmailSendingAccount is COGNITO_DEFAULT, you can't use the following parameters:
	//
	// - EmailVerificationMessage
	// - EmailVerificationSubject
	// - InviteMessageTemplate.EmailMessage
	// - InviteMessageTemplate.EmailSubject
	// - VerificationMessageTemplate.EmailMessage
	// - VerificationMessageTemplate.EmailMessageByLink
	// - VerificationMessageTemplate.EmailSubject,
	// - VerificationMessageTemplate.EmailSubjectByLink
	//
	// > DEVELOPER EmailSendingAccount is required.
	// - **DEVELOPER** - When Amazon Cognito emails your users, it uses your Amazon SES configuration. Amazon Cognito calls Amazon SES on your behalf to send email from your verified email address. When you use this option, the email delivery limits are the same limits that apply to your Amazon SES verified email address in your AWS account .
	//
	// If you use this option, you must provide the ARN of an Amazon SES verified email address for the `SourceArn` parameter.
	//
	// Before Amazon Cognito can email your users, it requires additional permissions to call Amazon SES on your behalf. When you update your user pool with this option, Amazon Cognito creates a *service-linked role* , which is a type of role, in your AWS account . This role contains the permissions that allow to access Amazon SES and send email messages with your address. For more information about the service-linked role that Amazon Cognito creates, see [Using Service-Linked Roles for Amazon Cognito](https://docs.aws.amazon.com/cognito/latest/developerguide/using-service-linked-roles.html) in the *Amazon Cognito Developer Guide* .
	EmailSendingAccount *string `json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// Identifies either the sender's email address or the sender's name with their email address.
	//
	// For example, `testuser@example.com` or `Test User <testuser@example.com>` . This address appears before the body of the email.
	From *string `json:"from" yaml:"from"`
	// The destination to which the receiver of the email should reply.
	ReplyToEmailAddress *string `json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// The ARN of a verified email address in Amazon SES.
	//
	// Amazon Cognito uses this email address in one of the following ways, depending on the value that you specify for the `EmailSendingAccount` parameter:
	//
	// - If you specify `COGNITO_DEFAULT` , Amazon Cognito uses this address as the custom FROM address when it emails your users using its built-in email account.
	// - If you specify `DEVELOPER` , Amazon Cognito emails your users with this address by calling Amazon SES on your behalf.
	//
	// The Region value of the `SourceArn` parameter must indicate a supported AWS Region of your user pool. Typically, the Region in the `SourceArn` and the user pool Region are the same. For more information, see [Amazon SES email configuration regions](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-email.html#user-pool-email-developer-region-mapping) in the [Amazon Cognito Developer Guide](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html) .
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
}

// The message template to be used for the welcome message to new users.
//
// See also [Customizing User Invitation Messages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.html#cognito-user-pool-settings-user-invitation-message-customization) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   inviteMessageTemplateProperty := &inviteMessageTemplateProperty{
//   	emailMessage: jsii.String("emailMessage"),
//   	emailSubject: jsii.String("emailSubject"),
//   	smsMessage: jsii.String("smsMessage"),
//   }
//
type CfnUserPool_InviteMessageTemplateProperty struct {
	// The message template for email messages.
	//
	// EmailMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailMessage *string `json:"emailMessage" yaml:"emailMessage"`
	// The subject line for email messages.
	//
	// EmailSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailSubject *string `json:"emailSubject" yaml:"emailSubject"`
	// The message template for SMS messages.
	SmsMessage *string `json:"smsMessage" yaml:"smsMessage"`
}

// Specifies the configuration for AWS Lambda triggers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   lambdaConfigProperty := &lambdaConfigProperty{
//   	createAuthChallenge: jsii.String("createAuthChallenge"),
//   	customEmailSender: &customEmailSenderProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		lambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	customMessage: jsii.String("customMessage"),
//   	customSmsSender: &customSMSSenderProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		lambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	defineAuthChallenge: jsii.String("defineAuthChallenge"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	postAuthentication: jsii.String("postAuthentication"),
//   	postConfirmation: jsii.String("postConfirmation"),
//   	preAuthentication: jsii.String("preAuthentication"),
//   	preSignUp: jsii.String("preSignUp"),
//   	preTokenGeneration: jsii.String("preTokenGeneration"),
//   	userMigration: jsii.String("userMigration"),
//   	verifyAuthChallengeResponse: jsii.String("verifyAuthChallengeResponse"),
//   }
//
type CfnUserPool_LambdaConfigProperty struct {
	// Creates an authentication challenge.
	CreateAuthChallenge *string `json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// A custom email sender AWS Lambda trigger.
	CustomEmailSender interface{} `json:"customEmailSender" yaml:"customEmailSender"`
	// A custom Message AWS Lambda trigger.
	CustomMessage *string `json:"customMessage" yaml:"customMessage"`
	// A custom SMS sender AWS Lambda trigger.
	CustomSmsSender interface{} `json:"customSmsSender" yaml:"customSmsSender"`
	// Defines the authentication challenge.
	DefineAuthChallenge *string `json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// The Amazon Resource Name of a AWS Key Management Service ( AWS KMS ) key.
	//
	// Amazon Cognito uses the key to encrypt codes and temporary passwords sent to `CustomEmailSender` and `CustomSMSSender` .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// A post-authentication AWS Lambda trigger.
	PostAuthentication *string `json:"postAuthentication" yaml:"postAuthentication"`
	// A post-confirmation AWS Lambda trigger.
	PostConfirmation *string `json:"postConfirmation" yaml:"postConfirmation"`
	// A pre-authentication AWS Lambda trigger.
	PreAuthentication *string `json:"preAuthentication" yaml:"preAuthentication"`
	// A pre-registration AWS Lambda trigger.
	PreSignUp *string `json:"preSignUp" yaml:"preSignUp"`
	// A Lambda trigger that is invoked before token generation.
	PreTokenGeneration *string `json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// The user migration Lambda config type.
	UserMigration *string `json:"userMigration" yaml:"userMigration"`
	// Verifies the authentication challenge response.
	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

// The minimum and maximum values of an attribute that is of the number data type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   numberAttributeConstraintsProperty := &numberAttributeConstraintsProperty{
//   	maxValue: jsii.String("maxValue"),
//   	minValue: jsii.String("minValue"),
//   }
//
type CfnUserPool_NumberAttributeConstraintsProperty struct {
	// The maximum value of an attribute that is of the number data type.
	MaxValue *string `json:"maxValue" yaml:"maxValue"`
	// The minimum value of an attribute that is of the number data type.
	MinValue *string `json:"minValue" yaml:"minValue"`
}

// The password policy type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   passwordPolicyProperty := &passwordPolicyProperty{
//   	minimumLength: jsii.Number(123),
//   	requireLowercase: jsii.Boolean(false),
//   	requireNumbers: jsii.Boolean(false),
//   	requireSymbols: jsii.Boolean(false),
//   	requireUppercase: jsii.Boolean(false),
//   	temporaryPasswordValidityDays: jsii.Number(123),
//   }
//
type CfnUserPool_PasswordPolicyProperty struct {
	// The minimum length of the password in the policy that you have set.
	//
	// This value can't be less than 6.
	MinimumLength *float64 `json:"minimumLength" yaml:"minimumLength"`
	// In the password policy that you have set, refers to whether you have required users to use at least one lowercase letter in their password.
	RequireLowercase interface{} `json:"requireLowercase" yaml:"requireLowercase"`
	// In the password policy that you have set, refers to whether you have required users to use at least one number in their password.
	RequireNumbers interface{} `json:"requireNumbers" yaml:"requireNumbers"`
	// In the password policy that you have set, refers to whether you have required users to use at least one symbol in their password.
	RequireSymbols interface{} `json:"requireSymbols" yaml:"requireSymbols"`
	// In the password policy that you have set, refers to whether you have required users to use at least one uppercase letter in their password.
	RequireUppercase interface{} `json:"requireUppercase" yaml:"requireUppercase"`
	// The number of days a temporary password is valid in the password policy.
	//
	// If the user doesn't sign in during this time, an administrator must reset their password.
	//
	// > When you set `TemporaryPasswordValidityDays` for a user pool, you can no longer set the deprecated `UnusedAccountValidityDays` value for that user pool.
	TemporaryPasswordValidityDays *float64 `json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

// The policy associated with a user pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   policiesProperty := &policiesProperty{
//   	passwordPolicy: &passwordPolicyProperty{
//   		minimumLength: jsii.Number(123),
//   		requireLowercase: jsii.Boolean(false),
//   		requireNumbers: jsii.Boolean(false),
//   		requireSymbols: jsii.Boolean(false),
//   		requireUppercase: jsii.Boolean(false),
//   		temporaryPasswordValidityDays: jsii.Number(123),
//   	},
//   }
//
type CfnUserPool_PoliciesProperty struct {
	// The password policy.
	PasswordPolicy interface{} `json:"passwordPolicy" yaml:"passwordPolicy"`
}

// A map containing a priority as a key, and recovery method name as a value.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   recoveryOptionProperty := &recoveryOptionProperty{
//   	name: jsii.String("name"),
//   	priority: jsii.Number(123),
//   }
//
type CfnUserPool_RecoveryOptionProperty struct {
	// Specifies the recovery method for a user.
	Name *string `json:"name" yaml:"name"`
	// A positive integer specifying priority of a method with 1 being the highest priority.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Contains information about the schema attribute.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   schemaAttributeProperty := &schemaAttributeProperty{
//   	attributeDataType: jsii.String("attributeDataType"),
//   	developerOnlyAttribute: jsii.Boolean(false),
//   	mutable: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	numberAttributeConstraints: &numberAttributeConstraintsProperty{
//   		maxValue: jsii.String("maxValue"),
//   		minValue: jsii.String("minValue"),
//   	},
//   	required: jsii.Boolean(false),
//   	stringAttributeConstraints: &stringAttributeConstraintsProperty{
//   		maxLength: jsii.String("maxLength"),
//   		minLength: jsii.String("minLength"),
//   	},
//   }
//
type CfnUserPool_SchemaAttributeProperty struct {
	// The attribute data type.
	AttributeDataType *string `json:"attributeDataType" yaml:"attributeDataType"`
	// > We recommend that you use [WriteAttributes](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can be mutated for new use cases instead of using `DeveloperOnlyAttribute` .
	//
	// Specifies whether the attribute type is developer only. This attribute can only be modified by an administrator. Users will not be able to modify this attribute using their access token.
	DeveloperOnlyAttribute interface{} `json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that is mapped to an identity provider attribute, you must set this parameter to `true` . Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider. If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your User Pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	Mutable interface{} `json:"mutable" yaml:"mutable"`
	// A schema attribute of the name type.
	Name *string `json:"name" yaml:"name"`
	// Specifies the constraints for an attribute of the number type.
	NumberAttributeConstraints interface{} `json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Specifies whether a user pool attribute is required.
	//
	// If the attribute is required and the user doesn't provide a value, registration or sign-in will fail.
	Required interface{} `json:"required" yaml:"required"`
	// Specifies the constraints for an attribute of the string type.
	StringAttributeConstraints interface{} `json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

// The SMS configuration type that includes the settings the Cognito User Pool needs to call for the Amazon SNS service to send an SMS message from your AWS account .
//
// The Cognito User Pool makes the request to the Amazon SNS Service by using an IAM role that you provide for your AWS account .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   smsConfigurationProperty := &smsConfigurationProperty{
//   	externalId: jsii.String("externalId"),
//   	snsCallerArn: jsii.String("snsCallerArn"),
//   	snsRegion: jsii.String("snsRegion"),
//   }
//
type CfnUserPool_SmsConfigurationProperty struct {
	// The external ID is a value.
	//
	// We recommend you use `ExternalId` to add security to your IAM role, which is used to call Amazon SNS to send SMS messages for your user pool. If you provide an `ExternalId` , the Cognito User Pool uses it when attempting to assume your IAM role. You can also set your roles trust policy to require the `ExternalID` . If you use the Cognito Management Console to create a role for SMS MFA, Cognito creates a role with the required permissions and a trust policy that uses `ExternalId` .
	ExternalId *string `json:"externalId" yaml:"externalId"`
	// The Amazon Resource Name (ARN) of the Amazon SNS caller.
	//
	// This is the ARN of the IAM role in your AWS account that Amazon Cognito will use to send SMS messages. SMS messages are subject to a [spending limit](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html) .
	SnsCallerArn *string `json:"snsCallerArn" yaml:"snsCallerArn"`
	// The AWS Region to use with Amazon SNS integration.
	//
	// You can choose the same Region as your user pool, or a supported *Legacy Amazon SNS alternate Region* .
	//
	// Amazon Cognito resources in the Asia Pacific (Seoul) AWS Region must use your Amazon SNS configuration in the Asia Pacific (Tokyo) Region. For more information, see [SMS message settings for Amazon Cognito user pools](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html) .
	SnsRegion *string `json:"snsRegion" yaml:"snsRegion"`
}

// The `StringAttributeConstraints` property type defines the string attribute constraints of an Amazon Cognito user pool.
//
// `StringAttributeConstraints` is a subproperty of the [SchemaAttribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html) property type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   stringAttributeConstraintsProperty := &stringAttributeConstraintsProperty{
//   	maxLength: jsii.String("maxLength"),
//   	minLength: jsii.String("minLength"),
//   }
//
type CfnUserPool_StringAttributeConstraintsProperty struct {
	// The maximum length.
	MaxLength *string `json:"maxLength" yaml:"maxLength"`
	// The minimum length.
	MinLength *string `json:"minLength" yaml:"minLength"`
}

// The user pool add-ons type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   userPoolAddOnsProperty := &userPoolAddOnsProperty{
//   	advancedSecurityMode: jsii.String("advancedSecurityMode"),
//   }
//
type CfnUserPool_UserPoolAddOnsProperty struct {
	// The advanced security mode.
	AdvancedSecurityMode *string `json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

// The `UsernameConfiguration` property type specifies case sensitivity on the username input for the selected sign-in option.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   usernameConfigurationProperty := &usernameConfigurationProperty{
//   	caseSensitive: jsii.Boolean(false),
//   }
//
type CfnUserPool_UsernameConfigurationProperty struct {
	// Specifies whether username case sensitivity will be applied for all users in the user pool through Amazon Cognito APIs.
	//
	// Valid values include:
	//
	// - **True** - Enables case sensitivity for all username input. When this option is set to `True` , users must sign in using the exact capitalization of their given username, such as “UserName”. This is the default value.
	// - **False** - Enables case insensitivity for all username input. For example, when this option is set to `False` , users can sign in using either "username" or "Username". This option also enables both `preferred_username` and `email` alias to be case insensitive, in addition to the `username` attribute.
	CaseSensitive interface{} `json:"caseSensitive" yaml:"caseSensitive"`
}

// The template for verification messages.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   verificationMessageTemplateProperty := &verificationMessageTemplateProperty{
//   	defaultEmailOption: jsii.String("defaultEmailOption"),
//   	emailMessage: jsii.String("emailMessage"),
//   	emailMessageByLink: jsii.String("emailMessageByLink"),
//   	emailSubject: jsii.String("emailSubject"),
//   	emailSubjectByLink: jsii.String("emailSubjectByLink"),
//   	smsMessage: jsii.String("smsMessage"),
//   }
//
type CfnUserPool_VerificationMessageTemplateProperty struct {
	// The default email option.
	DefaultEmailOption *string `json:"defaultEmailOption" yaml:"defaultEmailOption"`
	// The email message template.
	//
	// EmailMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailMessage *string `json:"emailMessage" yaml:"emailMessage"`
	// The email message template for sending a confirmation link to the user.
	//
	// EmailMessageByLink is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailMessageByLink *string `json:"emailMessageByLink" yaml:"emailMessageByLink"`
	// The subject line for the email message template.
	//
	// EmailSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailSubject *string `json:"emailSubject" yaml:"emailSubject"`
	// The subject line for the email message template for sending a confirmation link to the user.
	//
	// EmailSubjectByLink is allowed only [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailSubjectByLink *string `json:"emailSubjectByLink" yaml:"emailSubjectByLink"`
	// The SMS message template.
	SmsMessage *string `json:"smsMessage" yaml:"smsMessage"`
}

// A CloudFormation `AWS::Cognito::UserPoolClient`.
//
// The `AWS::Cognito::UserPoolClient` resource specifies an Amazon Cognito user pool client.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk"import elbv2 "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"type App awscdk.App
//   type CfnOutput awscdk.CfnOutput
//   type Stack awscdk.Stackimport constructs "github.com/aws/constructs-go/constructs"type Construct constructs.Constructimport actions "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStack struct {
//   stack
//   }
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &userPoolClientProps{
//   	userPool: userPool,
//
//   	// Required minimal configuration for use with an ELB
//   	generateSecret: jsii.Boolean(true),
//   	authFlows: &authFlow{
//   		userPassword: jsii.Boolean(true),
//   	},
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_EMAIL(),
//   		},
//   		callbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.node.defaultChild.(cfnUserPoolClient)
//   cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &userPoolDomainProps{
//   	userPool: userPool,
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(443),
//   	certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	defaultAction: actions.NewAuthenticateCognitoAction(&authenticateCognitoActionProps{
//   		userPool: userPool,
//   		userPoolClient: userPoolClient,
//   		userPoolDomain: userPoolDomain,
//   		next: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   			contentType: jsii.String("text/plain"),
//   			messageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
//   	value: lb.loadBalancerDnsName,
//   })
//
//   app := NewApp()
//   NewCognitoStack(app, jsii.String("integ-cognito"))
//   app.synth()
//
type CfnUserPoolClient interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time limit, after which the access token is no longer valid and cannot be used.
	AccessTokenValidity() *float64
	SetAccessTokenValidity(val *float64)
	// The allowed OAuth flows.
	//
	// Set to `code` to initiate a code grant flow, which provides an authorization code as the response. This code can be exchanged for access tokens with the token endpoint.
	//
	// Set to `implicit` to specify that the client should get the access token (and, optionally, ID token, based on scopes) directly.
	//
	// Set to `client_credentials` to specify that the client should get the access token (and, optionally, ID token, based on scopes) from the token endpoint using a combination of client and client_secret.
	AllowedOAuthFlows() *[]*string
	SetAllowedOAuthFlows(val *[]*string)
	// Set to true if the client is allowed to follow the OAuth protocol when interacting with Amazon Cognito user pools.
	AllowedOAuthFlowsUserPoolClient() interface{}
	SetAllowedOAuthFlowsUserPoolClient(val interface{})
	// The allowed OAuth scopes.
	//
	// Possible values provided by OAuth are: `phone` , `email` , `openid` , and `profile` . Possible values provided by AWS are: `aws.cognito.signin.user.admin` . Custom scopes created in Resource Servers are also supported.
	AllowedOAuthScopes() *[]*string
	SetAllowedOAuthScopes(val *[]*string)
	// The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.
	//
	// > In AWS Regions where Amazon Pinpoint isn't available, user pools only support sending events to Amazon Pinpoint projects in AWS Region us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
	AnalyticsConfiguration() interface{}
	SetAnalyticsConfiguration(val interface{})
	AttrClientSecret() *string
	AttrName() *string
	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackUrLs() *[]*string
	SetCallbackUrLs(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The client name for the user pool client you would like to create.
	ClientName() *string
	SetClientName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The default redirect URI. Must be in the `CallbackURLs` list.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectUri() *string
	SetDefaultRedirectUri(val *string)
	// Activates or deactivates token revocation. For more information about revoking tokens, see [RevokeToken](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html) .
	//
	// If you don't include this parameter, token revocation is automatically activated for the new user pool client.
	EnableTokenRevocation() interface{}
	SetEnableTokenRevocation(val interface{})
	// The authentication flows that are supported by the user pool clients.
	//
	// Flow names without the `ALLOW_` prefix are no longer supported, in favor of new names with the `ALLOW_` prefix.
	//
	// > Values with `ALLOW_` prefix must be used only along with the `ALLOW_` prefix.
	//
	// Valid values include:
	//
	// - `ALLOW_ADMIN_USER_PASSWORD_AUTH` : Enable admin based user password authentication flow `ADMIN_USER_PASSWORD_AUTH` . This setting replaces the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, Amazon Cognito receives the password in the request instead of using the Secure Remote Password (SRP) protocol to verify passwords.
	// - `ALLOW_CUSTOM_AUTH` : Enable AWS Lambda trigger based authentication.
	// - `ALLOW_USER_PASSWORD_AUTH` : Enable user password-based authentication. In this flow, Amazon Cognito receives the password in the request instead of using the SRP protocol to verify passwords.
	// - `ALLOW_USER_SRP_AUTH` : Enable SRP-based authentication.
	// - `ALLOW_REFRESH_TOKEN_AUTH` : Enable authflow to refresh tokens.
	ExplicitAuthFlows() *[]*string
	SetExplicitAuthFlows(val *[]*string)
	// Boolean to specify whether you want to generate a secret for the user pool client being created.
	GenerateSecret() interface{}
	SetGenerateSecret(val interface{})
	// The time limit, after which the ID token is no longer valid and cannot be used.
	IdTokenValidity() *float64
	SetIdTokenValidity(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// A list of allowed logout URLs for the identity providers.
	LogoutUrLs() *[]*string
	SetLogoutUrLs(val *[]*string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Use this setting to choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool.
	//
	// When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY` , those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors() *string
	SetPreventUserExistenceErrors(val *string)
	// The read attributes.
	ReadAttributes() *[]*string
	SetReadAttributes(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The time limit, in days, after which the refresh token is no longer valid and can't be used.
	RefreshTokenValidity() *float64
	SetRefreshTokenValidity(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of provider names for the identity providers that are supported on this client.
	//
	// The following are supported: `COGNITO` , `Facebook` , `SignInWithApple` , `Google` and `LoginWithAmazon` .
	SupportedIdentityProviders() *[]*string
	SetSupportedIdentityProviders(val *[]*string)
	// The units in which the validity times are represented in.
	//
	// Default for RefreshToken is days, and default for ID and access tokens are hours.
	TokenValidityUnits() interface{}
	SetTokenValidityUnits(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID for the user pool where you want to create a user pool client.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an identity provider, this array must include all attributes that you have mapped to identity provider attributes. Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider. If your app client does not have write access to a mapped attribute, Amazon Cognito throws an error when it tries to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	WriteAttributes() *[]*string
	SetWriteAttributes(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolClient
type jsiiProxy_CfnUserPoolClient struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOAuthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOAuthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AnalyticsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AttrClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CallbackUrLs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrLs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) GenerateSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) LogoutUrLs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrLs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) TokenValidityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClient(scope awscdk.Construct, id *string, props *CfnUserPoolClientProps) CfnUserPoolClient {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolClient{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolClient",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClient_Override(c CfnUserPoolClient, scope awscdk.Construct, id *string, props *CfnUserPoolClientProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolClient",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAccessTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthFlowsUserPoolClient(val interface{}) {
	_jsii_.Set(
		j,
		"allowedOAuthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOAuthScopes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAnalyticsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"analyticsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetCallbackUrLs(val *[]*string) {
	_jsii_.Set(
		j,
		"callbackUrLs",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetClientName(val *string) {
	_jsii_.Set(
		j,
		"clientName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetDefaultRedirectUri(val *string) {
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetEnableTokenRevocation(val interface{}) {
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetExplicitAuthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetGenerateSecret(val interface{}) {
	_jsii_.Set(
		j,
		"generateSecret",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetIdTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetLogoutUrLs(val *[]*string) {
	_jsii_.Set(
		j,
		"logoutUrLs",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetPreventUserExistenceErrors(val *string) {
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetReadAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetRefreshTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetSupportedIdentityProviders(val *[]*string) {
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetTokenValidityUnits(val interface{}) {
	_jsii_.Set(
		j,
		"tokenValidityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetWriteAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"writeAttributes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolClient_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolClient",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolClient_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolClient",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolClient_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolClient",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolClient) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolClient) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Amazon Pinpoint analytics configuration for collecting metrics for a user pool.
//
// > In Regions where Amazon Pinpointisn't available, user pools only support sending events to Amazon Pinpoint projects in us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   analyticsConfigurationProperty := &analyticsConfigurationProperty{
//   	applicationArn: jsii.String("applicationArn"),
//   	applicationId: jsii.String("applicationId"),
//   	externalId: jsii.String("externalId"),
//   	roleArn: jsii.String("roleArn"),
//   	userDataShared: jsii.Boolean(false),
//   }
//
type CfnUserPoolClient_AnalyticsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon Pinpoint project.
	//
	// You can use the Amazon Pinpoint project for integration with the chosen user pool client. Amazon Cognito publishes events to the Amazon Pinpoint project that the app ARN declares.
	ApplicationArn *string `json:"applicationArn" yaml:"applicationArn"`
	// The application ID for an Amazon Pinpoint application.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// The external ID.
	ExternalId *string `json:"externalId" yaml:"externalId"`
	// The ARN of an AWS Identity and Access Management role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// If `UserDataShared` is `true` , Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
	UserDataShared interface{} `json:"userDataShared" yaml:"userDataShared"`
}

// The units in which the validity times are represented in.
//
// Default for RefreshToken is days, and default for ID and access tokens are hours.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   tokenValidityUnitsProperty := &tokenValidityUnitsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	idToken: jsii.String("idToken"),
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnUserPoolClient_TokenValidityUnitsProperty struct {
	// A time unit in “seconds”, “minutes”, “hours” or “days” for the value in AccessTokenValidity, defaults to hours.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// A time unit in “seconds”, “minutes”, “hours” or “days” for the value in IdTokenValidity, defaults to hours.
	IdToken *string `json:"idToken" yaml:"idToken"`
	// A time unit in “seconds”, “minutes”, “hours” or “days” for the value in RefreshTokenValidity, defaults to days.
	RefreshToken *string `json:"refreshToken" yaml:"refreshToken"`
}

// Properties for defining a `CfnUserPoolClient`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolClientProps := &cfnUserPoolClientProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accessTokenValidity: jsii.Number(123),
//   	allowedOAuthFlows: []*string{
//   		jsii.String("allowedOAuthFlows"),
//   	},
//   	allowedOAuthFlowsUserPoolClient: jsii.Boolean(false),
//   	allowedOAuthScopes: []*string{
//   		jsii.String("allowedOAuthScopes"),
//   	},
//   	analyticsConfiguration: &analyticsConfigurationProperty{
//   		applicationArn: jsii.String("applicationArn"),
//   		applicationId: jsii.String("applicationId"),
//   		externalId: jsii.String("externalId"),
//   		roleArn: jsii.String("roleArn"),
//   		userDataShared: jsii.Boolean(false),
//   	},
//   	callbackUrLs: []*string{
//   		jsii.String("callbackUrLs"),
//   	},
//   	clientName: jsii.String("clientName"),
//   	defaultRedirectUri: jsii.String("defaultRedirectUri"),
//   	enableTokenRevocation: jsii.Boolean(false),
//   	explicitAuthFlows: []*string{
//   		jsii.String("explicitAuthFlows"),
//   	},
//   	generateSecret: jsii.Boolean(false),
//   	idTokenValidity: jsii.Number(123),
//   	logoutUrLs: []*string{
//   		jsii.String("logoutUrLs"),
//   	},
//   	preventUserExistenceErrors: jsii.String("preventUserExistenceErrors"),
//   	readAttributes: []*string{
//   		jsii.String("readAttributes"),
//   	},
//   	refreshTokenValidity: jsii.Number(123),
//   	supportedIdentityProviders: []*string{
//   		jsii.String("supportedIdentityProviders"),
//   	},
//   	tokenValidityUnits: &tokenValidityUnitsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		idToken: jsii.String("idToken"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	writeAttributes: []*string{
//   		jsii.String("writeAttributes"),
//   	},
//   }
//
type CfnUserPoolClientProps struct {
	// The user pool ID for the user pool where you want to create a user pool client.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// The time limit, after which the access token is no longer valid and cannot be used.
	AccessTokenValidity *float64 `json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The allowed OAuth flows.
	//
	// Set to `code` to initiate a code grant flow, which provides an authorization code as the response. This code can be exchanged for access tokens with the token endpoint.
	//
	// Set to `implicit` to specify that the client should get the access token (and, optionally, ID token, based on scopes) directly.
	//
	// Set to `client_credentials` to specify that the client should get the access token (and, optionally, ID token, based on scopes) from the token endpoint using a combination of client and client_secret.
	AllowedOAuthFlows *[]*string `json:"allowedOAuthFlows" yaml:"allowedOAuthFlows"`
	// Set to true if the client is allowed to follow the OAuth protocol when interacting with Amazon Cognito user pools.
	AllowedOAuthFlowsUserPoolClient interface{} `json:"allowedOAuthFlowsUserPoolClient" yaml:"allowedOAuthFlowsUserPoolClient"`
	// The allowed OAuth scopes.
	//
	// Possible values provided by OAuth are: `phone` , `email` , `openid` , and `profile` . Possible values provided by AWS are: `aws.cognito.signin.user.admin` . Custom scopes created in Resource Servers are also supported.
	AllowedOAuthScopes *[]*string `json:"allowedOAuthScopes" yaml:"allowedOAuthScopes"`
	// The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.
	//
	// > In AWS Regions where Amazon Pinpoint isn't available, user pools only support sending events to Amazon Pinpoint projects in AWS Region us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
	AnalyticsConfiguration interface{} `json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackUrLs *[]*string `json:"callbackUrLs" yaml:"callbackUrLs"`
	// The client name for the user pool client you would like to create.
	ClientName *string `json:"clientName" yaml:"clientName"`
	// The default redirect URI. Must be in the `CallbackURLs` list.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectUri *string `json:"defaultRedirectUri" yaml:"defaultRedirectUri"`
	// Activates or deactivates token revocation. For more information about revoking tokens, see [RevokeToken](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html) .
	//
	// If you don't include this parameter, token revocation is automatically activated for the new user pool client.
	EnableTokenRevocation interface{} `json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// The authentication flows that are supported by the user pool clients.
	//
	// Flow names without the `ALLOW_` prefix are no longer supported, in favor of new names with the `ALLOW_` prefix.
	//
	// > Values with `ALLOW_` prefix must be used only along with the `ALLOW_` prefix.
	//
	// Valid values include:
	//
	// - `ALLOW_ADMIN_USER_PASSWORD_AUTH` : Enable admin based user password authentication flow `ADMIN_USER_PASSWORD_AUTH` . This setting replaces the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, Amazon Cognito receives the password in the request instead of using the Secure Remote Password (SRP) protocol to verify passwords.
	// - `ALLOW_CUSTOM_AUTH` : Enable AWS Lambda trigger based authentication.
	// - `ALLOW_USER_PASSWORD_AUTH` : Enable user password-based authentication. In this flow, Amazon Cognito receives the password in the request instead of using the SRP protocol to verify passwords.
	// - `ALLOW_USER_SRP_AUTH` : Enable SRP-based authentication.
	// - `ALLOW_REFRESH_TOKEN_AUTH` : Enable authflow to refresh tokens.
	ExplicitAuthFlows *[]*string `json:"explicitAuthFlows" yaml:"explicitAuthFlows"`
	// Boolean to specify whether you want to generate a secret for the user pool client being created.
	GenerateSecret interface{} `json:"generateSecret" yaml:"generateSecret"`
	// The time limit, after which the ID token is no longer valid and cannot be used.
	IdTokenValidity *float64 `json:"idTokenValidity" yaml:"idTokenValidity"`
	// A list of allowed logout URLs for the identity providers.
	LogoutUrLs *[]*string `json:"logoutUrLs" yaml:"logoutUrLs"`
	// Use this setting to choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool.
	//
	// When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY` , those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The read attributes.
	ReadAttributes *[]*string `json:"readAttributes" yaml:"readAttributes"`
	// The time limit, in days, after which the refresh token is no longer valid and can't be used.
	RefreshTokenValidity *float64 `json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// A list of provider names for the identity providers that are supported on this client.
	//
	// The following are supported: `COGNITO` , `Facebook` , `SignInWithApple` , `Google` and `LoginWithAmazon` .
	SupportedIdentityProviders *[]*string `json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// The units in which the validity times are represented in.
	//
	// Default for RefreshToken is days, and default for ID and access tokens are hours.
	TokenValidityUnits interface{} `json:"tokenValidityUnits" yaml:"tokenValidityUnits"`
	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an identity provider, this array must include all attributes that you have mapped to identity provider attributes. Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider. If your app client does not have write access to a mapped attribute, Amazon Cognito throws an error when it tries to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	WriteAttributes *[]*string `json:"writeAttributes" yaml:"writeAttributes"`
}

// A CloudFormation `AWS::Cognito::UserPoolDomain`.
//
// The AWS::Cognito::UserPoolDomain resource creates a new domain for a user pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolDomain := cognito.NewCfnUserPoolDomain(this, jsii.String("MyCfnUserPoolDomain"), &cfnUserPoolDomainProps{
//   	domain: jsii.String("domain"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	customDomainConfig: &customDomainConfigTypeProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   	},
//   })
//
type CfnUserPoolDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// Use this object to specify an SSL certificate that is managed by ACM.
	CustomDomainConfig() interface{}
	SetCustomDomainConfig(val interface{})
	// The domain name for the domain that hosts the sign-up and sign-in pages for your application.
	//
	// For example: `auth.example.com` . If you're using a prefix domain, this field denotes the first part of the domain before `.auth.[region].amazoncognito.com` .
	//
	// This string can include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the first or last character. Use periods to separate subdomain names.
	Domain() *string
	SetDomain(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID for the user pool where you want to associate a user pool domain.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolDomain
type jsiiProxy_CfnUserPoolDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) CustomDomainConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDomainConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolDomain) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolDomain`.
func NewCfnUserPoolDomain(scope awscdk.Construct, id *string, props *CfnUserPoolDomainProps) CfnUserPoolDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolDomain{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolDomain`.
func NewCfnUserPoolDomain_Override(c CfnUserPoolDomain, scope awscdk.Construct, id *string, props *CfnUserPoolDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolDomain) SetCustomDomainConfig(val interface{}) {
	_jsii_.Set(
		j,
		"customDomainConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolDomain) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolDomain) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration for a custom domain that hosts the sign-up and sign-in webpages for your application.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   customDomainConfigTypeProperty := &customDomainConfigTypeProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnUserPoolDomain_CustomDomainConfigTypeProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Certificate Manager SSL certificate.
	//
	// You use this certificate for the subdomain of your custom domain.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
}

// Properties for defining a `CfnUserPoolDomain`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolDomainProps := &cfnUserPoolDomainProps{
//   	domain: jsii.String("domain"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	customDomainConfig: &customDomainConfigTypeProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   	},
//   }
//
type CfnUserPoolDomainProps struct {
	// The domain name for the domain that hosts the sign-up and sign-in pages for your application.
	//
	// For example: `auth.example.com` . If you're using a prefix domain, this field denotes the first part of the domain before `.auth.[region].amazoncognito.com` .
	//
	// This string can include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the first or last character. Use periods to separate subdomain names.
	Domain *string `json:"domain" yaml:"domain"`
	// The user pool ID for the user pool where you want to associate a user pool domain.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// Use this object to specify an SSL certificate that is managed by ACM.
	CustomDomainConfig interface{} `json:"customDomainConfig" yaml:"customDomainConfig"`
}

// A CloudFormation `AWS::Cognito::UserPoolGroup`.
//
// Specifies a new group in the identified user pool.
//
// Calling this action requires developer credentials.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolGroup := cognito.NewCfnUserPoolGroup(this, jsii.String("MyCfnUserPoolGroup"), &cfnUserPoolGroupProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	groupName: jsii.String("groupName"),
//   	precedence: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnUserPoolGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A string containing the description of the group.
	Description() *string
	SetDescription(val *string)
	// The name of the group.
	//
	// Must be unique.
	GroupName() *string
	SetGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.
	//
	// Zero is the highest precedence value. Groups with lower `Precedence` values take precedence over groups with higher ornull `Precedence` values. If a user belongs to two or more groups, it is the group with the lowest precedence value whose role ARN is given in the user's tokens for the `cognito:roles` and `cognito:preferred_role` claims.
	//
	// Two groups can have the same `Precedence` value. If this happens, neither group takes precedence over the other. If two groups with the same `Precedence` have the same role ARN, that role is used in the `cognito:preferred_role` claim in tokens for users in each group. If the two groups have different role ARNs, the `cognito:preferred_role` claim isn't set in users' tokens.
	//
	// The default `Precedence` value is null.
	Precedence() *float64
	SetPrecedence(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The role Amazon Resource Name (ARN) for the group.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID for the user pool.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolGroup
type jsiiProxy_CfnUserPoolGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolGroup) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolGroup`.
func NewCfnUserPoolGroup(scope awscdk.Construct, id *string, props *CfnUserPoolGroupProps) CfnUserPoolGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolGroup{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolGroup`.
func NewCfnUserPoolGroup_Override(c CfnUserPoolGroup, scope awscdk.Construct, id *string, props *CfnUserPoolGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolGroup) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolGroup) SetPrecedence(val *float64) {
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolGroup) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolGroup) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnUserPoolGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolGroupProps := &cfnUserPoolGroupProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	groupName: jsii.String("groupName"),
//   	precedence: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnUserPoolGroupProps struct {
	// The user pool ID for the user pool.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// A string containing the description of the group.
	Description *string `json:"description" yaml:"description"`
	// The name of the group.
	//
	// Must be unique.
	GroupName *string `json:"groupName" yaml:"groupName"`
	// A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.
	//
	// Zero is the highest precedence value. Groups with lower `Precedence` values take precedence over groups with higher ornull `Precedence` values. If a user belongs to two or more groups, it is the group with the lowest precedence value whose role ARN is given in the user's tokens for the `cognito:roles` and `cognito:preferred_role` claims.
	//
	// Two groups can have the same `Precedence` value. If this happens, neither group takes precedence over the other. If two groups with the same `Precedence` have the same role ARN, that role is used in the `cognito:preferred_role` claim in tokens for users in each group. If the two groups have different role ARNs, the `cognito:preferred_role` claim isn't set in users' tokens.
	//
	// The default `Precedence` value is null.
	Precedence *float64 `json:"precedence" yaml:"precedence"`
	// The role Amazon Resource Name (ARN) for the group.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::Cognito::UserPoolIdentityProvider`.
//
// The `AWS::Cognito::UserPoolIdentityProvider` resource creates an identity provider for a user pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var attributeMapping interface{}
//   var providerDetails interface{}
//   cfnUserPoolIdentityProvider := cognito.NewCfnUserPoolIdentityProvider(this, jsii.String("MyCfnUserPoolIdentityProvider"), &cfnUserPoolIdentityProviderProps{
//   	providerName: jsii.String("providerName"),
//   	providerType: jsii.String("providerType"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	attributeMapping: attributeMapping,
//   	idpIdentifiers: []*string{
//   		jsii.String("idpIdentifiers"),
//   	},
//   	providerDetails: providerDetails,
//   })
//
type CfnUserPoolIdentityProvider interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A mapping of identity provider attributes to standard and custom user pool attributes.
	AttributeMapping() interface{}
	SetAttributeMapping(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A list of identity provider identifiers.
	IdpIdentifiers() *[]*string
	SetIdpIdentifiers(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The identity provider details. The following list describes the provider detail keys for each identity provider type.
	//
	// - For Google and Login with Amazon:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - For Facebook:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - api_version
	// - For Sign in with Apple:
	//
	// - client_id
	// - team_id
	// - key_id
	// - private_key
	// - authorize_scopes
	// - For OpenID Connect (OIDC) providers:
	//
	// - client_id
	// - client_secret
	// - attributes_request_method
	// - oidc_issuer
	// - authorize_scopes
	// - authorize_url *if not available from discovery URL specified by oidc_issuer key*
	// - token_url *if not available from discovery URL specified by oidc_issuer key*
	// - attributes_url *if not available from discovery URL specified by oidc_issuer key*
	// - jwks_uri *if not available from discovery URL specified by oidc_issuer key*
	// - attributes_url_add_attributes *a read-only property that is set automatically*
	// - For SAML providers:
	//
	// - MetadataFile OR MetadataURL
	// - IDPSignout (optional).
	ProviderDetails() interface{}
	SetProviderDetails(val interface{})
	// The identity provider name.
	ProviderName() *string
	SetProviderName(val *string)
	// The identity provider type.
	ProviderType() *string
	SetProviderType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolIdentityProvider
type jsiiProxy_CfnUserPoolIdentityProvider struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) AttributeMapping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) IdpIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) ProviderDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolIdentityProvider`.
func NewCfnUserPoolIdentityProvider(scope awscdk.Construct, id *string, props *CfnUserPoolIdentityProviderProps) CfnUserPoolIdentityProvider {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolIdentityProvider{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolIdentityProvider`.
func NewCfnUserPoolIdentityProvider_Override(c CfnUserPoolIdentityProvider, scope awscdk.Construct, id *string, props *CfnUserPoolIdentityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetAttributeMapping(val interface{}) {
	_jsii_.Set(
		j,
		"attributeMapping",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetIdpIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"idpIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetProviderDetails(val interface{}) {
	_jsii_.Set(
		j,
		"providerDetails",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetProviderName(val *string) {
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetProviderType(val *string) {
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolIdentityProvider) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolIdentityProvider_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolIdentityProvider_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolIdentityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolIdentityProvider_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolIdentityProvider",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolIdentityProvider) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnUserPoolIdentityProvider`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var attributeMapping interface{}
//   var providerDetails interface{}
//   cfnUserPoolIdentityProviderProps := &cfnUserPoolIdentityProviderProps{
//   	providerName: jsii.String("providerName"),
//   	providerType: jsii.String("providerType"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	attributeMapping: attributeMapping,
//   	idpIdentifiers: []*string{
//   		jsii.String("idpIdentifiers"),
//   	},
//   	providerDetails: providerDetails,
//   }
//
type CfnUserPoolIdentityProviderProps struct {
	// The identity provider name.
	ProviderName *string `json:"providerName" yaml:"providerName"`
	// The identity provider type.
	ProviderType *string `json:"providerType" yaml:"providerType"`
	// The user pool ID.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// A mapping of identity provider attributes to standard and custom user pool attributes.
	AttributeMapping interface{} `json:"attributeMapping" yaml:"attributeMapping"`
	// A list of identity provider identifiers.
	IdpIdentifiers *[]*string `json:"idpIdentifiers" yaml:"idpIdentifiers"`
	// The identity provider details. The following list describes the provider detail keys for each identity provider type.
	//
	// - For Google and Login with Amazon:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - For Facebook:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - api_version
	// - For Sign in with Apple:
	//
	// - client_id
	// - team_id
	// - key_id
	// - private_key
	// - authorize_scopes
	// - For OpenID Connect (OIDC) providers:
	//
	// - client_id
	// - client_secret
	// - attributes_request_method
	// - oidc_issuer
	// - authorize_scopes
	// - authorize_url *if not available from discovery URL specified by oidc_issuer key*
	// - token_url *if not available from discovery URL specified by oidc_issuer key*
	// - attributes_url *if not available from discovery URL specified by oidc_issuer key*
	// - jwks_uri *if not available from discovery URL specified by oidc_issuer key*
	// - attributes_url_add_attributes *a read-only property that is set automatically*
	// - For SAML providers:
	//
	// - MetadataFile OR MetadataURL
	// - IDPSignout (optional).
	ProviderDetails interface{} `json:"providerDetails" yaml:"providerDetails"`
}

// Properties for defining a `CfnUserPool`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var userPoolTags interface{}
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
	AccountRecoverySetting interface{} `json:"accountRecoverySetting" yaml:"accountRecoverySetting"`
	// The configuration for creating a new user profile.
	AdminCreateUserConfig interface{} `json:"adminCreateUserConfig" yaml:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
	//
	// > This user pool property cannot be updated.
	AliasAttributes *[]*string `json:"aliasAttributes" yaml:"aliasAttributes"`
	// The attributes to be auto-verified.
	//
	// Possible values: *email* , *phone_number* .
	AutoVerifiedAttributes *[]*string `json:"autoVerifiedAttributes" yaml:"autoVerifiedAttributes"`
	// The device configuration.
	DeviceConfiguration interface{} `json:"deviceConfiguration" yaml:"deviceConfiguration"`
	// The email configuration of your user pool.
	//
	// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
	EmailConfiguration interface{} `json:"emailConfiguration" yaml:"emailConfiguration"`
	// A string representing the email verification message.
	//
	// EmailVerificationMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationMessage *string `json:"emailVerificationMessage" yaml:"emailVerificationMessage"`
	// A string representing the email verification subject.
	//
	// EmailVerificationSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
	EmailVerificationSubject *string `json:"emailVerificationSubject" yaml:"emailVerificationSubject"`
	// Enables MFA on a specified user pool.
	//
	// To disable all MFAs after it has been enabled, set MfaConfiguration to “OFF” and remove EnabledMfas. MFAs can only be all disabled if MfaConfiguration is OFF. Once SMS_MFA is enabled, SMS_MFA can only be disabled by setting MfaConfiguration to “OFF”. Can be one of the following values:
	//
	// - `SMS_MFA` - Enables SMS MFA for the user pool. SMS_MFA can only be enabled if SMS configuration is provided.
	// - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
	//
	// Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA`.
	EnabledMfas *[]*string `json:"enabledMfas" yaml:"enabledMfas"`
	// The Lambda trigger configuration information for the new user pool.
	//
	// > In a push model, event sources (such as Amazon S3 and custom applications) need permission to invoke a function. So you must make an extra call to add permission for these event sources to invoke your Lambda function.
	// >
	// > For more information on using the Lambda API to add permission, see [AddPermission](https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html) .
	// >
	// > For adding permission using the AWS CLI , see [add-permission](https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html) .
	LambdaConfig interface{} `json:"lambdaConfig" yaml:"lambdaConfig"`
	// The multi-factor (MFA) configuration. Valid values include:.
	//
	// - `OFF` MFA won't be used for any users.
	// - `ON` MFA is required for all users to sign in.
	// - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
	MfaConfiguration *string `json:"mfaConfiguration" yaml:"mfaConfiguration"`
	// The policy associated with a user pool.
	Policies interface{} `json:"policies" yaml:"policies"`
	// The schema attributes for the new user pool. These attributes can be standard or custom attributes.
	//
	// > During a user pool update, you can add new schema attributes but you cannot modify or delete an existing schema attribute.
	Schema interface{} `json:"schema" yaml:"schema"`
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage" yaml:"smsAuthenticationMessage"`
	// The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service.
	//
	// To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
	SmsConfiguration interface{} `json:"smsConfiguration" yaml:"smsConfiguration"`
	// A string representing the SMS verification message.
	SmsVerificationMessage *string `json:"smsVerificationMessage" yaml:"smsVerificationMessage"`
	// Determines whether email addresses or phone numbers can be specified as user names when a user signs up.
	//
	// Possible values: `phone_number` or `email` .
	//
	// This user pool property cannot be updated.
	UsernameAttributes *[]*string `json:"usernameAttributes" yaml:"usernameAttributes"`
	// You can choose to set case sensitivity on the username input for the selected sign-in option.
	//
	// For example, when this is set to `False` , users will be able to sign in using either "username" or "Username". This configuration is immutable once it has been set.
	UsernameConfiguration interface{} `json:"usernameConfiguration" yaml:"usernameConfiguration"`
	// Enables advanced security risk detection.
	//
	// Set the key `AdvancedSecurityMode` to the value "AUDIT".
	UserPoolAddOns interface{} `json:"userPoolAddOns" yaml:"userPoolAddOns"`
	// A string used to name the user pool.
	UserPoolName *string `json:"userPoolName" yaml:"userPoolName"`
	// The tag keys and values to assign to the user pool.
	//
	// A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
	UserPoolTags interface{} `json:"userPoolTags" yaml:"userPoolTags"`
	// The template for the verification message that the user sees when the app requests permission to access the user's information.
	VerificationMessageTemplate interface{} `json:"verificationMessageTemplate" yaml:"verificationMessageTemplate"`
}

// A CloudFormation `AWS::Cognito::UserPoolResourceServer`.
//
// The `AWS::Cognito::UserPoolResourceServer` resource creates a new OAuth2.0 resource server and defines custom scopes in it.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolResourceServer := cognito.NewCfnUserPoolResourceServer(this, jsii.String("MyCfnUserPoolResourceServer"), &cfnUserPoolResourceServerProps{
//   	identifier: jsii.String("identifier"),
//   	name: jsii.String("name"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	scopes: []interface{}{
//   		&resourceServerScopeTypeProperty{
//   			scopeDescription: jsii.String("scopeDescription"),
//   			scopeName: jsii.String("scopeName"),
//   		},
//   	},
//   })
//
type CfnUserPoolResourceServer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A unique resource server identifier for the resource server.
	//
	// This could be an HTTPS endpoint where the resource server is located. For example: `https://my-weather-api.example.com` .
	Identifier() *string
	SetIdentifier(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// A friendly name for the resource server.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list of scopes.
	//
	// Each scope is a map with keys `ScopeName` and `ScopeDescription` .
	Scopes() interface{}
	SetScopes(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID for the user pool.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolResourceServer
type jsiiProxy_CfnUserPoolResourceServer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolResourceServer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Scopes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServer) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolResourceServer`.
func NewCfnUserPoolResourceServer(scope awscdk.Construct, id *string, props *CfnUserPoolResourceServerProps) CfnUserPoolResourceServer {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolResourceServer{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolResourceServer`.
func NewCfnUserPoolResourceServer_Override(c CfnUserPoolResourceServer, scope awscdk.Construct, id *string, props *CfnUserPoolResourceServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolResourceServer) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolResourceServer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolResourceServer) SetScopes(val interface{}) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolResourceServer) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolResourceServer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolResourceServer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolResourceServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolResourceServer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolResourceServer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A resource server scope.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   resourceServerScopeTypeProperty := &resourceServerScopeTypeProperty{
//   	scopeDescription: jsii.String("scopeDescription"),
//   	scopeName: jsii.String("scopeName"),
//   }
//
type CfnUserPoolResourceServer_ResourceServerScopeTypeProperty struct {
	// A description of the scope.
	ScopeDescription *string `json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	ScopeName *string `json:"scopeName" yaml:"scopeName"`
}

// Properties for defining a `CfnUserPoolResourceServer`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolResourceServerProps := &cfnUserPoolResourceServerProps{
//   	identifier: jsii.String("identifier"),
//   	name: jsii.String("name"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	scopes: []interface{}{
//   		&resourceServerScopeTypeProperty{
//   			scopeDescription: jsii.String("scopeDescription"),
//   			scopeName: jsii.String("scopeName"),
//   		},
//   	},
//   }
//
type CfnUserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	//
	// This could be an HTTPS endpoint where the resource server is located. For example: `https://my-weather-api.example.com` .
	Identifier *string `json:"identifier" yaml:"identifier"`
	// A friendly name for the resource server.
	Name *string `json:"name" yaml:"name"`
	// The user pool ID for the user pool.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// A list of scopes.
	//
	// Each scope is a map with keys `ScopeName` and `ScopeDescription` .
	Scopes interface{} `json:"scopes" yaml:"scopes"`
}

// A CloudFormation `AWS::Cognito::UserPoolRiskConfigurationAttachment`.
//
// The `AWS::Cognito::UserPoolRiskConfigurationAttachment` resource sets the risk configuration that is used for Amazon Cognito advanced security features.
//
// You can specify risk configuration for a single client (with a specific `clientId` ) or for all clients (by setting the `clientId` to `ALL` ). If you specify `ALL` , the default configuration is used for every client that has had no risk configuration set previously. If you specify risk configuration for a particular client, it no longer falls back to the `ALL` configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolRiskConfigurationAttachment := cognito.NewCfnUserPoolRiskConfigurationAttachment(this, jsii.String("MyCfnUserPoolRiskConfigurationAttachment"), &cfnUserPoolRiskConfigurationAttachmentProps{
//   	clientId: jsii.String("clientId"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accountTakeoverRiskConfiguration: &accountTakeoverRiskConfigurationTypeProperty{
//   		actions: &accountTakeoverActionsTypeProperty{
//   			highAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			lowAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			mediumAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		notifyConfiguration: &notifyConfigurationTypeProperty{
//   			sourceArn: jsii.String("sourceArn"),
//
//   			// the properties below are optional
//   			blockEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			from: jsii.String("from"),
//   			mfaEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			noActionEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			replyTo: jsii.String("replyTo"),
//   		},
//   	},
//   	compromisedCredentialsRiskConfiguration: &compromisedCredentialsRiskConfigurationTypeProperty{
//   		actions: &compromisedCredentialsActionsTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   		},
//
//   		// the properties below are optional
//   		eventFilter: []*string{
//   			jsii.String("eventFilter"),
//   		},
//   	},
//   	riskExceptionConfiguration: &riskExceptionConfigurationTypeProperty{
//   		blockedIpRangeList: []*string{
//   			jsii.String("blockedIpRangeList"),
//   		},
//   		skippedIpRangeList: []*string{
//   			jsii.String("skippedIpRangeList"),
//   		},
//   	},
//   })
//
type CfnUserPoolRiskConfigurationAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
	AccountTakeoverRiskConfiguration() interface{}
	SetAccountTakeoverRiskConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The app client ID.
	//
	// You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	ClientId() *string
	SetClientId(val *string)
	// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
	CompromisedCredentialsRiskConfiguration() interface{}
	SetCompromisedCredentialsRiskConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The configuration to override the risk decision.
	RiskExceptionConfiguration() interface{}
	SetRiskExceptionConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolRiskConfigurationAttachment
type jsiiProxy_CfnUserPoolRiskConfigurationAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AccountTakeoverRiskConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountTakeoverRiskConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) CompromisedCredentialsRiskConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compromisedCredentialsRiskConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) RiskExceptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"riskExceptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolRiskConfigurationAttachment`.
func NewCfnUserPoolRiskConfigurationAttachment(scope awscdk.Construct, id *string, props *CfnUserPoolRiskConfigurationAttachmentProps) CfnUserPoolRiskConfigurationAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolRiskConfigurationAttachment{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolRiskConfigurationAttachment`.
func NewCfnUserPoolRiskConfigurationAttachment_Override(c CfnUserPoolRiskConfigurationAttachment, scope awscdk.Construct, id *string, props *CfnUserPoolRiskConfigurationAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) SetAccountTakeoverRiskConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"accountTakeoverRiskConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) SetCompromisedCredentialsRiskConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"compromisedCredentialsRiskConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) SetRiskExceptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"riskExceptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolRiskConfigurationAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolRiskConfigurationAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolRiskConfigurationAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolRiskConfigurationAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolRiskConfigurationAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Account takeover action type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   accountTakeoverActionTypeProperty := &accountTakeoverActionTypeProperty{
//   	eventAction: jsii.String("eventAction"),
//   	notify: jsii.Boolean(false),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionTypeProperty struct {
	// The action to take in response to the account takeover action. Valid values are:.
	//
	// - `BLOCK` Choosing this action will block the request.
	// - `MFA_IF_CONFIGURED` Present an MFA challenge if user has configured it, else allow the request.
	// - `MFA_REQUIRED` Present an MFA challenge if user has configured it, else block the request.
	// - `NO_ACTION` Allow the user to sign in.
	EventAction *string `json:"eventAction" yaml:"eventAction"`
	// Flag specifying whether to send a notification.
	Notify interface{} `json:"notify" yaml:"notify"`
}

// Account takeover actions type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   accountTakeoverActionsTypeProperty := &accountTakeoverActionsTypeProperty{
//   	highAction: &accountTakeoverActionTypeProperty{
//   		eventAction: jsii.String("eventAction"),
//   		notify: jsii.Boolean(false),
//   	},
//   	lowAction: &accountTakeoverActionTypeProperty{
//   		eventAction: jsii.String("eventAction"),
//   		notify: jsii.Boolean(false),
//   	},
//   	mediumAction: &accountTakeoverActionTypeProperty{
//   		eventAction: jsii.String("eventAction"),
//   		notify: jsii.Boolean(false),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionsTypeProperty struct {
	// Action to take for a high risk.
	HighAction interface{} `json:"highAction" yaml:"highAction"`
	// Action to take for a low risk.
	LowAction interface{} `json:"lowAction" yaml:"lowAction"`
	// Action to take for a medium risk.
	MediumAction interface{} `json:"mediumAction" yaml:"mediumAction"`
}

// Configuration for mitigation actions and notification for different levels of risk detected for a potential account takeover.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   accountTakeoverRiskConfigurationTypeProperty := &accountTakeoverRiskConfigurationTypeProperty{
//   	actions: &accountTakeoverActionsTypeProperty{
//   		highAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   		lowAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   		mediumAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   	},
//
//   	// the properties below are optional
//   	notifyConfiguration: &notifyConfigurationTypeProperty{
//   		sourceArn: jsii.String("sourceArn"),
//
//   		// the properties below are optional
//   		blockEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		from: jsii.String("from"),
//   		mfaEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		noActionEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		replyTo: jsii.String("replyTo"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationTypeProperty struct {
	// Account takeover risk configuration actions.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The notify configuration used to construct email notifications.
	NotifyConfiguration interface{} `json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

// The compromised credentials actions type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   compromisedCredentialsActionsTypeProperty := &compromisedCredentialsActionsTypeProperty{
//   	eventAction: jsii.String("eventAction"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsTypeProperty struct {
	// The event action.
	EventAction *string `json:"eventAction" yaml:"eventAction"`
}

// The compromised credentials risk configuration type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   compromisedCredentialsRiskConfigurationTypeProperty := &compromisedCredentialsRiskConfigurationTypeProperty{
//   	actions: &compromisedCredentialsActionsTypeProperty{
//   		eventAction: jsii.String("eventAction"),
//   	},
//
//   	// the properties below are optional
//   	eventFilter: []*string{
//   		jsii.String("eventFilter"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationTypeProperty struct {
	// The compromised credentials risk configuration actions.
	Actions interface{} `json:"actions" yaml:"actions"`
	// Perform the action for these events.
	//
	// The default is to perform all events if no event filter is specified.
	EventFilter *[]*string `json:"eventFilter" yaml:"eventFilter"`
}

// The notify configuration type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   notifyConfigurationTypeProperty := &notifyConfigurationTypeProperty{
//   	sourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	blockEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	from: jsii.String("from"),
//   	mfaEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	noActionEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	replyTo: jsii.String("replyTo"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_NotifyConfigurationTypeProperty struct {
	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy.
	//
	// This identity permits Amazon Cognito to send for the email address specified in the `From` parameter.
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
	// Email template used when a detected risk event is blocked.
	BlockEmail interface{} `json:"blockEmail" yaml:"blockEmail"`
	// The email address that is sending the email.
	//
	// The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	From *string `json:"from" yaml:"from"`
	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk.
	MfaEmail interface{} `json:"mfaEmail" yaml:"mfaEmail"`
	// The email template used when a detected risk event is allowed.
	NoActionEmail interface{} `json:"noActionEmail" yaml:"noActionEmail"`
	// The destination to which the receiver of an email should reply to.
	ReplyTo *string `json:"replyTo" yaml:"replyTo"`
}

// The notify email type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   notifyEmailTypeProperty := &notifyEmailTypeProperty{
//   	subject: jsii.String("subject"),
//
//   	// the properties below are optional
//   	htmlBody: jsii.String("htmlBody"),
//   	textBody: jsii.String("textBody"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_NotifyEmailTypeProperty struct {
	// The email subject.
	Subject *string `json:"subject" yaml:"subject"`
	// The email HTML body.
	HtmlBody *string `json:"htmlBody" yaml:"htmlBody"`
	// The email text body.
	TextBody *string `json:"textBody" yaml:"textBody"`
}

// The type of the configuration to override the risk decision.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   riskExceptionConfigurationTypeProperty := &riskExceptionConfigurationTypeProperty{
//   	blockedIpRangeList: []*string{
//   		jsii.String("blockedIpRangeList"),
//   	},
//   	skippedIpRangeList: []*string{
//   		jsii.String("skippedIpRangeList"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_RiskExceptionConfigurationTypeProperty struct {
	// Overrides the risk decision to always block the pre-authentication requests.
	//
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	BlockedIpRangeList *[]*string `json:"blockedIpRangeList" yaml:"blockedIpRangeList"`
	// Risk detection isn't performed on the IP addresses in this range list.
	//
	// The IP range is in CIDR notation.
	SkippedIpRangeList *[]*string `json:"skippedIpRangeList" yaml:"skippedIpRangeList"`
}

// Properties for defining a `CfnUserPoolRiskConfigurationAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolRiskConfigurationAttachmentProps := &cfnUserPoolRiskConfigurationAttachmentProps{
//   	clientId: jsii.String("clientId"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accountTakeoverRiskConfiguration: &accountTakeoverRiskConfigurationTypeProperty{
//   		actions: &accountTakeoverActionsTypeProperty{
//   			highAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			lowAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			mediumAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		notifyConfiguration: &notifyConfigurationTypeProperty{
//   			sourceArn: jsii.String("sourceArn"),
//
//   			// the properties below are optional
//   			blockEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			from: jsii.String("from"),
//   			mfaEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			noActionEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			replyTo: jsii.String("replyTo"),
//   		},
//   	},
//   	compromisedCredentialsRiskConfiguration: &compromisedCredentialsRiskConfigurationTypeProperty{
//   		actions: &compromisedCredentialsActionsTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   		},
//
//   		// the properties below are optional
//   		eventFilter: []*string{
//   			jsii.String("eventFilter"),
//   		},
//   	},
//   	riskExceptionConfiguration: &riskExceptionConfigurationTypeProperty{
//   		blockedIpRangeList: []*string{
//   			jsii.String("blockedIpRangeList"),
//   		},
//   		skippedIpRangeList: []*string{
//   			jsii.String("skippedIpRangeList"),
//   		},
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachmentProps struct {
	// The app client ID.
	//
	// You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The user pool ID.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
	AccountTakeoverRiskConfiguration interface{} `json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
	CompromisedCredentialsRiskConfiguration interface{} `json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// The configuration to override the risk decision.
	RiskExceptionConfiguration interface{} `json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

// A CloudFormation `AWS::Cognito::UserPoolUICustomizationAttachment`.
//
// The `AWS::Cognito::UserPoolUICustomizationAttachment` resource sets the UI customization information for a user pool's built-in app UI.
//
// You can specify app UI customization settings for a single client (with a specific `clientId` ) or for all clients (by setting the `clientId` to `ALL` ). If you specify `ALL` , the default configuration is used for every client that has had no UI customization set previously. If you specify UI customization settings for a particular client, it no longer falls back to the `ALL` configuration.
//
// > Before you create this resource, your user pool must have a domain associated with it. You can create an `AWS::Cognito::UserPoolDomain` resource first in this user pool.
//
// Setting a logo image isn't supported from AWS CloudFormation . Use the Amazon Cognito [SetUICustomization](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUICustomization.html#API_SetUICustomization_RequestSyntax) API operation to set the image.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolUICustomizationAttachment := cognito.NewCfnUserPoolUICustomizationAttachment(this, jsii.String("MyCfnUserPoolUICustomizationAttachment"), &cfnUserPoolUICustomizationAttachmentProps{
//   	clientId: jsii.String("clientId"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	css: jsii.String("css"),
//   })
//
type CfnUserPoolUICustomizationAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The client ID for the client app.
	//
	// You can specify the UI customization settings for a single client (with a specific clientId) or for all clients (by setting the clientId to `ALL` ).
	ClientId() *string
	SetClientId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The CSS values in the UI customization.
	Css() *string
	SetCss(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user pool ID for the user pool.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolUICustomizationAttachment
type jsiiProxy_CfnUserPoolUICustomizationAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) Css() *string {
	var returns *string
	_jsii_.Get(
		j,
		"css",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolUICustomizationAttachment`.
func NewCfnUserPoolUICustomizationAttachment(scope awscdk.Construct, id *string, props *CfnUserPoolUICustomizationAttachmentProps) CfnUserPoolUICustomizationAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolUICustomizationAttachment{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolUICustomizationAttachment`.
func NewCfnUserPoolUICustomizationAttachment_Override(c CfnUserPoolUICustomizationAttachment, scope awscdk.Construct, id *string, props *CfnUserPoolUICustomizationAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) SetCss(val *string) {
	_jsii_.Set(
		j,
		"css",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUICustomizationAttachment) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolUICustomizationAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolUICustomizationAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolUICustomizationAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolUICustomizationAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolUICustomizationAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUICustomizationAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnUserPoolUICustomizationAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolUICustomizationAttachmentProps := &cfnUserPoolUICustomizationAttachmentProps{
//   	clientId: jsii.String("clientId"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	css: jsii.String("css"),
//   }
//
type CfnUserPoolUICustomizationAttachmentProps struct {
	// The client ID for the client app.
	//
	// You can specify the UI customization settings for a single client (with a specific clientId) or for all clients (by setting the clientId to `ALL` ).
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The user pool ID for the user pool.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// The CSS values in the UI customization.
	Css *string `json:"css" yaml:"css"`
}

// A CloudFormation `AWS::Cognito::UserPoolUser`.
//
// The `AWS::Cognito::UserPoolUser` resource creates an Amazon Cognito user pool user.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var clientMetadata interface{}
//   cfnUserPoolUser := cognito.NewCfnUserPoolUser(this, jsii.String("MyCfnUserPoolUser"), &cfnUserPoolUserProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	clientMetadata: clientMetadata,
//   	desiredDeliveryMediums: []*string{
//   		jsii.String("desiredDeliveryMediums"),
//   	},
//   	forceAliasCreation: jsii.Boolean(false),
//   	messageAction: jsii.String("messageAction"),
//   	userAttributes: []interface{}{
//   		&attributeTypeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	username: jsii.String("username"),
//   	validationData: []interface{}{
//   		&attributeTypeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnUserPoolUser interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A map of custom key-value pairs that you can provide as input for the custom workflow that is invoked by the *pre sign-up* trigger.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you create a `UserPoolUser` resource and include the `ClientMetadata` property, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `clientMetadata` attribute, which provides the data that you assigned to the ClientMetadata property. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
	//
	// For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
	//
	// > Take the following limitations into consideration when you use the ClientMetadata parameter:
	// >
	// > - Amazon Cognito does not store the ClientMetadata value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose.
	// > - Amazon Cognito does not validate the ClientMetadata value.
	// > - Amazon Cognito does not encrypt the the ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata() interface{}
	SetClientMetadata(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specify `"EMAIL"` if email will be used to send the welcome message.
	//
	// Specify `"SMS"` if the phone number will be used. The default value is `"SMS"` . You can specify more than one value.
	DesiredDeliveryMediums() *[]*string
	SetDesiredDeliveryMediums(val *[]*string)
	// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` .
	//
	// Otherwise, it is ignored.
	//
	// If this parameter is set to `True` and the phone number or email address specified in the UserAttributes parameter already exists as an alias with a different user, the API call will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias.
	//
	// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
	ForceAliasCreation() interface{}
	SetForceAliasCreation(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account.
	//
	// Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
	MessageAction() *string
	SetMessageAction(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user attributes and attribute values to be set for the user to be created.
	//
	// These are name-value pairs You can create a user without specifying any attributes other than `Username` . However, any attributes that you specify as required (in [](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) or in the *Attributes* tab of the console) must be supplied either by you (in your call to `AdminCreateUser` ) or by the user (when they sign up in response to your welcome message).
	//
	// For custom attributes, you must prepend the `custom:` prefix to the attribute name.
	//
	// To send a message inviting the user to sign up, you must specify the user's email address or phone number. This can be done in your call to AdminCreateUser or in the *Users* tab of the Amazon Cognito console for managing your user pools.
	//
	// In your call to `AdminCreateUser` , you can set the `email_verified` attribute to `True` , and you can set the `phone_number_verified` attribute to `True` . (You can also do this by calling [](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminUpdateUserAttributes.html) .)
	//
	// - *email* : The email address of the user to whom the message that contains the code and user name will be sent. Required if the `email_verified` attribute is set to `True` , or if `"EMAIL"` is specified in the `DesiredDeliveryMediums` parameter.
	// - *phone_number* : The phone number of the user to whom the message that contains the code and user name will be sent. Required if the `phone_number_verified` attribute is set to `True` , or if `"SMS"` is specified in the `DesiredDeliveryMediums` parameter.
	UserAttributes() interface{}
	SetUserAttributes(val interface{})
	// The username for the user.
	//
	// Must be unique within the user pool. Must be a UTF-8 string between 1 and 128 characters. After the user is created, the username can't be changed.
	Username() *string
	SetUsername(val *string)
	// The user pool ID for the user pool where the user will be created.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// The user's validation data.
	//
	// This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. For example, you might choose to allow or disallow user sign-up based on the user's domain.
	//
	// To configure custom validation, you must create a Pre Sign-up AWS Lambda trigger for the user pool as described in the Amazon Cognito Developer Guide. The Lambda trigger receives the validation data and uses it in the validation process.
	//
	// The user's validation data isn't persisted.
	ValidationData() interface{}
	SetValidationData(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolUser
type jsiiProxy_CfnUserPoolUser struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolUser) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) ClientMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) DesiredDeliveryMediums() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"desiredDeliveryMediums",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) ForceAliasCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceAliasCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) MessageAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) UserAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUser) ValidationData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationData",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolUser`.
func NewCfnUserPoolUser(scope awscdk.Construct, id *string, props *CfnUserPoolUserProps) CfnUserPoolUser {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolUser{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUser",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolUser`.
func NewCfnUserPoolUser_Override(c CfnUserPoolUser, scope awscdk.Construct, id *string, props *CfnUserPoolUserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUser",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetClientMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"clientMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetDesiredDeliveryMediums(val *[]*string) {
	_jsii_.Set(
		j,
		"desiredDeliveryMediums",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetForceAliasCreation(val interface{}) {
	_jsii_.Set(
		j,
		"forceAliasCreation",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetMessageAction(val *string) {
	_jsii_.Set(
		j,
		"messageAction",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetUserAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"userAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUser) SetValidationData(val interface{}) {
	_jsii_.Set(
		j,
		"validationData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolUser_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUser",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolUser_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUser",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolUser_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolUser",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUser) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUser) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUser) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies whether the attribute is standard or custom.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   attributeTypeProperty := &attributeTypeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnUserPoolUser_AttributeTypeProperty struct {
	// The name of the attribute.
	Name *string `json:"name" yaml:"name"`
	// The value of the attribute.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnUserPoolUser`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var clientMetadata interface{}
//   cfnUserPoolUserProps := &cfnUserPoolUserProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	clientMetadata: clientMetadata,
//   	desiredDeliveryMediums: []*string{
//   		jsii.String("desiredDeliveryMediums"),
//   	},
//   	forceAliasCreation: jsii.Boolean(false),
//   	messageAction: jsii.String("messageAction"),
//   	userAttributes: []interface{}{
//   		&attributeTypeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	username: jsii.String("username"),
//   	validationData: []interface{}{
//   		&attributeTypeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserPoolUserProps struct {
	// The user pool ID for the user pool where the user will be created.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// A map of custom key-value pairs that you can provide as input for the custom workflow that is invoked by the *pre sign-up* trigger.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you create a `UserPoolUser` resource and include the `ClientMetadata` property, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `clientMetadata` attribute, which provides the data that you assigned to the ClientMetadata property. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
	//
	// For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
	//
	// > Take the following limitations into consideration when you use the ClientMetadata parameter:
	// >
	// > - Amazon Cognito does not store the ClientMetadata value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose.
	// > - Amazon Cognito does not validate the ClientMetadata value.
	// > - Amazon Cognito does not encrypt the the ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata interface{} `json:"clientMetadata" yaml:"clientMetadata"`
	// Specify `"EMAIL"` if email will be used to send the welcome message.
	//
	// Specify `"SMS"` if the phone number will be used. The default value is `"SMS"` . You can specify more than one value.
	DesiredDeliveryMediums *[]*string `json:"desiredDeliveryMediums" yaml:"desiredDeliveryMediums"`
	// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` .
	//
	// Otherwise, it is ignored.
	//
	// If this parameter is set to `True` and the phone number or email address specified in the UserAttributes parameter already exists as an alias with a different user, the API call will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias.
	//
	// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
	ForceAliasCreation interface{} `json:"forceAliasCreation" yaml:"forceAliasCreation"`
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account.
	//
	// Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
	MessageAction *string `json:"messageAction" yaml:"messageAction"`
	// The user attributes and attribute values to be set for the user to be created.
	//
	// These are name-value pairs You can create a user without specifying any attributes other than `Username` . However, any attributes that you specify as required (in [](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) or in the *Attributes* tab of the console) must be supplied either by you (in your call to `AdminCreateUser` ) or by the user (when they sign up in response to your welcome message).
	//
	// For custom attributes, you must prepend the `custom:` prefix to the attribute name.
	//
	// To send a message inviting the user to sign up, you must specify the user's email address or phone number. This can be done in your call to AdminCreateUser or in the *Users* tab of the Amazon Cognito console for managing your user pools.
	//
	// In your call to `AdminCreateUser` , you can set the `email_verified` attribute to `True` , and you can set the `phone_number_verified` attribute to `True` . (You can also do this by calling [](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminUpdateUserAttributes.html) .)
	//
	// - *email* : The email address of the user to whom the message that contains the code and user name will be sent. Required if the `email_verified` attribute is set to `True` , or if `"EMAIL"` is specified in the `DesiredDeliveryMediums` parameter.
	// - *phone_number* : The phone number of the user to whom the message that contains the code and user name will be sent. Required if the `phone_number_verified` attribute is set to `True` , or if `"SMS"` is specified in the `DesiredDeliveryMediums` parameter.
	UserAttributes interface{} `json:"userAttributes" yaml:"userAttributes"`
	// The username for the user.
	//
	// Must be unique within the user pool. Must be a UTF-8 string between 1 and 128 characters. After the user is created, the username can't be changed.
	Username *string `json:"username" yaml:"username"`
	// The user's validation data.
	//
	// This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. For example, you might choose to allow or disallow user sign-up based on the user's domain.
	//
	// To configure custom validation, you must create a Pre Sign-up AWS Lambda trigger for the user pool as described in the Amazon Cognito Developer Guide. The Lambda trigger receives the validation data and uses it in the validation process.
	//
	// The user's validation data isn't persisted.
	ValidationData interface{} `json:"validationData" yaml:"validationData"`
}

// A CloudFormation `AWS::Cognito::UserPoolUserToGroupAttachment`.
//
// Adds the specified user to the specified group.
//
// Calling this action requires developer credentials.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolUserToGroupAttachment := cognito.NewCfnUserPoolUserToGroupAttachment(this, jsii.String("MyCfnUserPoolUserToGroupAttachment"), &cfnUserPoolUserToGroupAttachmentProps{
//   	groupName: jsii.String("groupName"),
//   	username: jsii.String("username"),
//   	userPoolId: jsii.String("userPoolId"),
//   })
//
type CfnUserPoolUserToGroupAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The group name.
	GroupName() *string
	SetGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The username for the user.
	Username() *string
	SetUsername(val *string)
	// The user pool ID for the user pool.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserPoolUserToGroupAttachment
type jsiiProxy_CfnUserPoolUserToGroupAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolUserToGroupAttachment`.
func NewCfnUserPoolUserToGroupAttachment(scope awscdk.Construct, id *string, props *CfnUserPoolUserToGroupAttachmentProps) CfnUserPoolUserToGroupAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolUserToGroupAttachment{}

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolUserToGroupAttachment`.
func NewCfnUserPoolUserToGroupAttachment_Override(c CfnUserPoolUserToGroupAttachment, scope awscdk.Construct, id *string, props *CfnUserPoolUserToGroupAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolUserToGroupAttachment) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnUserPoolUserToGroupAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserPoolUserToGroupAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserPoolUserToGroupAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolUserToGroupAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cognito.CfnUserPoolUserToGroupAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolUserToGroupAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnUserPoolUserToGroupAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   cfnUserPoolUserToGroupAttachmentProps := &cfnUserPoolUserToGroupAttachmentProps{
//   	groupName: jsii.String("groupName"),
//   	username: jsii.String("username"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnUserPoolUserToGroupAttachmentProps struct {
	// The group name.
	GroupName *string `json:"groupName" yaml:"groupName"`
	// The username for the user.
	Username *string `json:"username" yaml:"username"`
	// The user pool ID for the user pool.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
}

// A set of attributes, useful to set Read and Write attributes.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   clientWriteAttributes := (cognito.NewClientAttributes()).withStandardAttributes(&standardAttributesMask{
//   	fullname: jsii.Boolean(true),
//   	email: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("favouritePizza"), jsii.String("favouriteBeverage"))
//
//   clientReadAttributes := clientWriteAttributes.withStandardAttributes(&standardAttributesMask{
//   	emailVerified: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("pointsEarned"))
//
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	readAttributes: clientReadAttributes,
//   	writeAttributes: clientWriteAttributes,
//   })
//
// Experimental.
type ClientAttributes interface {
	// The list of attributes represented by this ClientAttributes.
	// Experimental.
	Attributes() *[]*string
	// Creates a custom ClientAttributes with the specified attributes.
	// Experimental.
	WithCustomAttributes(attributes ...*string) ClientAttributes
	// Creates a custom ClientAttributes with the specified attributes.
	// Experimental.
	WithStandardAttributes(attributes *StandardAttributesMask) ClientAttributes
}

// The jsii proxy struct for ClientAttributes
type jsiiProxy_ClientAttributes struct {
	_ byte // padding
}

// Creates a ClientAttributes with the specified attributes.
// Experimental.
func NewClientAttributes() ClientAttributes {
	_init_.Initialize()

	j := jsiiProxy_ClientAttributes{}

	_jsii_.Create(
		"monocdk.aws_cognito.ClientAttributes",
		nil, // no parameters
		&j,
	)

	return &j
}

// Creates a ClientAttributes with the specified attributes.
// Experimental.
func NewClientAttributes_Override(c ClientAttributes) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.ClientAttributes",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_ClientAttributes) Attributes() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"attributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClientAttributes) WithCustomAttributes(attributes ...*string) ClientAttributes {
	args := []interface{}{}
	for _, a := range attributes {
		args = append(args, a)
	}

	var returns ClientAttributes

	_jsii_.Invoke(
		c,
		"withCustomAttributes",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClientAttributes) WithStandardAttributes(attributes *StandardAttributesMask) ClientAttributes {
	var returns ClientAttributes

	_jsii_.Invoke(
		c,
		"withStandardAttributes",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

// Options while specifying a cognito prefix domain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   pool.addDomain(jsii.String("CognitoDomain"), &userPoolDomainOptions{
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("my-awesome-app"),
//   	},
//   })
//
//   certificateArn := "arn:aws:acm:us-east-1:123456789012:certificate/11-3336f1-44483d-adc7-9cd375c5169d"
//
//   domainCert := certificatemanager.certificate.fromCertificateArn(this, jsii.String("domainCert"), certificateArn)
//   pool.addDomain(jsii.String("CustomDomain"), &userPoolDomainOptions{
//   	customDomain: &customDomainOptions{
//   		domainName: jsii.String("user.myapp.com"),
//   		certificate: domainCert,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
//
// Experimental.
type CognitoDomainOptions struct {
	// The prefix to the Cognito hosted domain name that will be associated with the user pool.
	// Experimental.
	DomainPrefix *string `json:"domainPrefix" yaml:"domainPrefix"`
}

// Configuration that will be fed into CloudFormation for any custom attribute type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   customAttributeConfig := &customAttributeConfig{
//   	dataType: jsii.String("dataType"),
//
//   	// the properties below are optional
//   	mutable: jsii.Boolean(false),
//   	numberConstraints: &numberAttributeConstraints{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	stringConstraints: &stringAttributeConstraints{
//   		maxLen: jsii.Number(123),
//   		minLen: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type CustomAttributeConfig struct {
	// The data type of the custom attribute.
	// See: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SchemaAttributeType.html#CognitoUserPools-Type-SchemaAttributeType-AttributeDataType
	//
	// Experimental.
	DataType *string `json:"dataType" yaml:"dataType"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
	// The constraints for a custom attribute of the 'Number' data type.
	// Experimental.
	NumberConstraints *NumberAttributeConstraints `json:"numberConstraints" yaml:"numberConstraints"`
	// The constraints for a custom attribute of 'String' data type.
	// Experimental.
	StringConstraints *StringAttributeConstraints `json:"stringConstraints" yaml:"stringConstraints"`
}

// Constraints that can be applied to a custom attribute of any type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type CustomAttributeProps struct {
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
}

// Options while specifying custom domain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   pool.addDomain(jsii.String("CognitoDomain"), &userPoolDomainOptions{
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("my-awesome-app"),
//   	},
//   })
//
//   certificateArn := "arn:aws:acm:us-east-1:123456789012:certificate/11-3336f1-44483d-adc7-9cd375c5169d"
//
//   domainCert := certificatemanager.certificate.fromCertificateArn(this, jsii.String("domainCert"), certificateArn)
//   pool.addDomain(jsii.String("CustomDomain"), &userPoolDomainOptions{
//   	customDomain: &customDomainOptions{
//   		domainName: jsii.String("user.myapp.com"),
//   		certificate: domainCert,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
//
// Experimental.
type CustomDomainOptions struct {
	// The certificate to associate with this domain.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate" yaml:"certificate"`
	// The custom domain name that you would like to associate with this User Pool.
	// Experimental.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// The DateTime custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type DateTimeAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	// Experimental.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for DateTimeAttribute
type jsiiProxy_DateTimeAttribute struct {
	jsiiProxy_ICustomAttribute
}

// Experimental.
func NewDateTimeAttribute(props *CustomAttributeProps) DateTimeAttribute {
	_init_.Initialize()

	j := jsiiProxy_DateTimeAttribute{}

	_jsii_.Create(
		"monocdk.aws_cognito.DateTimeAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDateTimeAttribute_Override(d DateTimeAttribute, props *CustomAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.DateTimeAttribute",
		[]interface{}{props},
		d,
	)
}

func (d *jsiiProxy_DateTimeAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		d,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Device tracking settings.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	deviceTracking: &deviceTracking{
//   		challengeRequiredOnNewDevice: jsii.Boolean(true),
//   		deviceOnlyRememberedOnUserPrompt: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
//
// Experimental.
type DeviceTracking struct {
	// Indicates whether a challenge is required on a new device.
	//
	// Only applicable to a new device.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	//
	// Experimental.
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// If true, a device is only remembered on user prompt.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	//
	// Experimental.
	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

// Email settings for the user pool.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   emailSettings := &emailSettings{
//   	from: jsii.String("from"),
//   	replyTo: jsii.String("replyTo"),
//   }
//
// Experimental.
type EmailSettings struct {
	// The 'from' address on the emails received by the user.
	// Experimental.
	From *string `json:"from" yaml:"from"`
	// The 'replyTo' address on the emails received by the user as defined by IETF RFC-5322.
	//
	// When set, most email clients recognize to change 'to' line to this address when a reply is drafted.
	// Experimental.
	ReplyTo *string `json:"replyTo" yaml:"replyTo"`
}

// Represents a custom attribute type.
// Experimental.
type ICustomAttribute interface {
	// Bind this custom attribute type to the values as expected by CloudFormation.
	// Experimental.
	Bind() *CustomAttributeConfig
}

// The jsii proxy for ICustomAttribute
type jsiiProxy_ICustomAttribute struct {
	_ byte // padding
}

func (i *jsiiProxy_ICustomAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		i,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a Cognito UserPool.
// Experimental.
type IUserPool interface {
	awscdk.IResource
	// Add a new app client to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html
	//
	// Experimental.
	AddClient(id *string, options *UserPoolClientOptions) UserPoolClient
	// Associate a domain to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain.html
	//
	// Experimental.
	AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain
	// Add a new resource server to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-resource-servers.html
	//
	// Experimental.
	AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer
	// Register an identity provider with this user pool.
	// Experimental.
	RegisterIdentityProvider(provider IUserPoolIdentityProvider)
	// Get all identity providers registered with this user pool.
	// Experimental.
	IdentityProviders() *[]IUserPoolIdentityProvider
	// The ARN of this user pool resource.
	// Experimental.
	UserPoolArn() *string
	// The physical ID of this user pool resource.
	// Experimental.
	UserPoolId() *string
}

// The jsii proxy for IUserPool
type jsiiProxy_IUserPool struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IUserPool) AddClient(id *string, options *UserPoolClientOptions) UserPoolClient {
	var returns UserPoolClient

	_jsii_.Invoke(
		i,
		"addClient",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain {
	var returns UserPoolDomain

	_jsii_.Invoke(
		i,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer {
	var returns UserPoolResourceServer

	_jsii_.Invoke(
		i,
		"addResourceServer",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) RegisterIdentityProvider(provider IUserPoolIdentityProvider) {
	_jsii_.InvokeVoid(
		i,
		"registerIdentityProvider",
		[]interface{}{provider},
	)
}

func (j *jsiiProxy_IUserPool) IdentityProviders() *[]IUserPoolIdentityProvider {
	var returns *[]IUserPoolIdentityProvider
	_jsii_.Get(
		j,
		"identityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPool) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPool) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

// Represents a Cognito user pool client.
// Experimental.
type IUserPoolClient interface {
	awscdk.IResource
	// Name of the application client.
	// Experimental.
	UserPoolClientId() *string
}

// The jsii proxy for IUserPoolClient
type jsiiProxy_IUserPoolClient struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolClient) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

// Represents a user pool domain.
// Experimental.
type IUserPoolDomain interface {
	awscdk.IResource
	// The domain that was specified to be created.
	//
	// If `customDomain` was selected, this holds the full domain name that was specified.
	// If the `cognitoDomain` was used, it contains the prefix to the Cognito hosted domain.
	// Experimental.
	DomainName() *string
}

// The jsii proxy for IUserPoolDomain
type jsiiProxy_IUserPoolDomain struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

// Represents a UserPoolIdentityProvider.
// Experimental.
type IUserPoolIdentityProvider interface {
	awscdk.IResource
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
}

// The jsii proxy for IUserPoolIdentityProvider
type jsiiProxy_IUserPoolIdentityProvider struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolIdentityProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

// Represents a Cognito user pool resource server.
// Experimental.
type IUserPoolResourceServer interface {
	awscdk.IResource
	// Resource server id.
	// Experimental.
	UserPoolResourceServerId() *string
}

// The jsii proxy for IUserPoolResourceServer
type jsiiProxy_IUserPoolResourceServer struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolResourceServer) UserPoolResourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolResourceServerId",
		&returns,
	)
	return returns
}

// The different ways in which a user pool's MFA enforcement can be configured.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	mfa: cognito.mfa_REQUIRED,
//   	mfaSecondFactor: &mfaSecondFactor{
//   		sms: jsii.Boolean(true),
//   		otp: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html
//
// Experimental.
type Mfa string

const (
	// Users are not required to use MFA for sign in, and cannot configure one.
	// Experimental.
	Mfa_OFF Mfa = "OFF"
	// Users are not required to use MFA for sign in, but can configure one if they so choose to.
	// Experimental.
	Mfa_OPTIONAL Mfa = "OPTIONAL"
	// Users are required to configure an MFA, and have to use it to sign in.
	// Experimental.
	Mfa_REQUIRED Mfa = "REQUIRED"
)

// The different ways in which a user pool can obtain their MFA token for sign in.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	mfa: cognito.mfa_REQUIRED,
//   	mfaSecondFactor: &mfaSecondFactor{
//   		sms: jsii.Boolean(true),
//   		otp: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html
//
// Experimental.
type MfaSecondFactor struct {
	// The MFA token is a time-based one time password that is generated by a hardware or software token.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa-totp.html
	//
	// Experimental.
	Otp *bool `json:"otp" yaml:"otp"`
	// The MFA token is sent to the user via SMS to their verified phone numbers.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa-sms-text-message.html
	//
	// Experimental.
	Sms *bool `json:"sms" yaml:"sms"`
}

// The Number custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type NumberAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	// Experimental.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for NumberAttribute
type jsiiProxy_NumberAttribute struct {
	jsiiProxy_ICustomAttribute
}

// Experimental.
func NewNumberAttribute(props *NumberAttributeProps) NumberAttribute {
	_init_.Initialize()

	j := jsiiProxy_NumberAttribute{}

	_jsii_.Create(
		"monocdk.aws_cognito.NumberAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewNumberAttribute_Override(n NumberAttribute, props *NumberAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.NumberAttribute",
		[]interface{}{props},
		n,
	)
}

func (n *jsiiProxy_NumberAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		n,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Constraints that can be applied to a custom attribute of number type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   numberAttributeConstraints := &numberAttributeConstraints{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
// Experimental.
type NumberAttributeConstraints struct {
	// Maximum value of this attribute.
	// Experimental.
	Max *float64 `json:"max" yaml:"max"`
	// Minimum value of this attribute.
	// Experimental.
	Min *float64 `json:"min" yaml:"min"`
}

// Props for NumberAttr.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type NumberAttributeProps struct {
	// Maximum value of this attribute.
	// Experimental.
	Max *float64 `json:"max" yaml:"max"`
	// Minimum value of this attribute.
	// Experimental.
	Min *float64 `json:"min" yaml:"min"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
}

// Types of OAuth grant flows.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			implicitCodeGrant: jsii.Boolean(true),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &userPoolDomainOptions{
//   })
//   signInUrl := domain.signInUrl(client, &signInUrlOptions{
//   	redirectUri: jsii.String("https://myapp.com/home"),
//   })
//
// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
// Experimental.
type OAuthFlows struct {
	// Initiate an authorization code grant flow, which provides an authorization code as the response.
	// Experimental.
	AuthorizationCodeGrant *bool `json:"authorizationCodeGrant" yaml:"authorizationCodeGrant"`
	// Client should get the access token and ID token from the token endpoint using a combination of client and client_secret.
	// Experimental.
	ClientCredentials *bool `json:"clientCredentials" yaml:"clientCredentials"`
	// The client should get the access token and ID token directly.
	// Experimental.
	ImplicitCodeGrant *bool `json:"implicitCodeGrant" yaml:"implicitCodeGrant"`
}

// OAuth scopes that are allowed with this client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
// Experimental.
type OAuthScope interface {
	// The name of this scope as recognized by CloudFormation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-allowedoauthscopes
	//
	// Experimental.
	ScopeName() *string
}

// The jsii proxy struct for OAuthScope
type jsiiProxy_OAuthScope struct {
	_ byte // padding
}

func (j *jsiiProxy_OAuthScope) ScopeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeName",
		&returns,
	)
	return returns
}


// Custom scope is one that you define for your own resource server in the Resource Servers.
//
// The format is 'resource-server-identifier/scope'.
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
//
// Experimental.
func OAuthScope_Custom(name *string) OAuthScope {
	_init_.Initialize()

	var returns OAuthScope

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.OAuthScope",
		"custom",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Adds a custom scope that's tied to a resource server in your stack.
// Experimental.
func OAuthScope_ResourceServer(server IUserPoolResourceServer, scope ResourceServerScope) OAuthScope {
	_init_.Initialize()

	var returns OAuthScope

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.OAuthScope",
		"resourceServer",
		[]interface{}{server, scope},
		&returns,
	)

	return returns
}

func OAuthScope_COGNITO_ADMIN() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"monocdk.aws_cognito.OAuthScope",
		"COGNITO_ADMIN",
		&returns,
	)
	return returns
}

func OAuthScope_EMAIL() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"monocdk.aws_cognito.OAuthScope",
		"EMAIL",
		&returns,
	)
	return returns
}

func OAuthScope_OPENID() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"monocdk.aws_cognito.OAuthScope",
		"OPENID",
		&returns,
	)
	return returns
}

func OAuthScope_PHONE() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"monocdk.aws_cognito.OAuthScope",
		"PHONE",
		&returns,
	)
	return returns
}

func OAuthScope_PROFILE() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"monocdk.aws_cognito.OAuthScope",
		"PROFILE",
		&returns,
	)
	return returns
}

// OAuth settings to configure the interaction between the app and this client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type OAuthSettings struct {
	// List of allowed redirect URLs for the identity providers.
	// Experimental.
	CallbackUrls *[]*string `json:"callbackUrls" yaml:"callbackUrls"`
	// OAuth flows that are allowed with this client.
	// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Experimental.
	Flows *OAuthFlows `json:"flows" yaml:"flows"`
	// List of allowed logout URLs for the identity providers.
	// Experimental.
	LogoutUrls *[]*string `json:"logoutUrls" yaml:"logoutUrls"`
	// OAuth scopes that are allowed with this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Experimental.
	Scopes *[]OAuthScope `json:"scopes" yaml:"scopes"`
}

// Password policy for User Pools.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	passwordPolicy: &passwordPolicy{
//   		minLength: jsii.Number(12),
//   		requireLowercase: jsii.Boolean(true),
//   		requireUppercase: jsii.Boolean(true),
//   		requireDigits: jsii.Boolean(true),
//   		requireSymbols: jsii.Boolean(true),
//   		tempPasswordValidity: duration.days(jsii.Number(3)),
//   	},
//   })
//
// Experimental.
type PasswordPolicy struct {
	// Minimum length required for a user's password.
	// Experimental.
	MinLength *float64 `json:"minLength" yaml:"minLength"`
	// Whether the user is required to have digits in their password.
	// Experimental.
	RequireDigits *bool `json:"requireDigits" yaml:"requireDigits"`
	// Whether the user is required to have lowercase characters in their password.
	// Experimental.
	RequireLowercase *bool `json:"requireLowercase" yaml:"requireLowercase"`
	// Whether the user is required to have symbols in their password.
	// Experimental.
	RequireSymbols *bool `json:"requireSymbols" yaml:"requireSymbols"`
	// Whether the user is required to have uppercase characters in their password.
	// Experimental.
	RequireUppercase *bool `json:"requireUppercase" yaml:"requireUppercase"`
	// The length of time the temporary password generated by an admin is valid.
	//
	// This must be provided as whole days, like Duration.days(3) or Duration.hours(48).
	// Fractional days, such as Duration.hours(20), will generate an error.
	// Experimental.
	TempPasswordValidity awscdk.Duration `json:"tempPasswordValidity" yaml:"tempPasswordValidity"`
}

// An attribute available from a third party identity provider.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   	userPool: userpool,
//   	attributeMapping: &attributeMapping{
//   		email: cognito.providerAttribute_AMAZON_EMAIL(),
//   		website: cognito.*providerAttribute.other(jsii.String("url")),
//   		 // use other() when an attribute is not pre-defined in the CDK
//   		custom: map[string]*providerAttribute{
//   			// custom user pool attributes go here
//   			"uniqueId": cognito.*providerAttribute_AMAZON_USER_ID(),
//   		},
//   	},
//   })
//
// Experimental.
type ProviderAttribute interface {
	// The attribute value string as recognized by the provider.
	// Experimental.
	AttributeName() *string
}

// The jsii proxy struct for ProviderAttribute
type jsiiProxy_ProviderAttribute struct {
	_ byte // padding
}

func (j *jsiiProxy_ProviderAttribute) AttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeName",
		&returns,
	)
	return returns
}


// Use this to specify an attribute from the identity provider that is not pre-defined in the CDK.
// Experimental.
func ProviderAttribute_Other(attributeName *string) ProviderAttribute {
	_init_.Initialize()

	var returns ProviderAttribute

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.ProviderAttribute",
		"other",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func ProviderAttribute_AMAZON_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"AMAZON_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"AMAZON_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_POSTAL_CODE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"AMAZON_POSTAL_CODE",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_USER_ID() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"AMAZON_USER_ID",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"APPLE_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_FIRST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"APPLE_FIRST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_LAST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"APPLE_LAST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"APPLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_BIRTHDAY() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_BIRTHDAY",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_FIRST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_FIRST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_GENDER() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_GENDER",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_ID() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_ID",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_LAST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_LAST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_LOCALE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_LOCALE",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_MIDDLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_MIDDLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"FACEBOOK_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_BIRTHDAYS() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_BIRTHDAYS",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_FAMILY_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_FAMILY_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_GENDER() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_GENDER",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_GIVEN_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_GIVEN_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_NAMES() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_NAMES",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_PHONE_NUMBERS() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_PHONE_NUMBERS",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_PICTURE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"monocdk.aws_cognito.ProviderAttribute",
		"GOOGLE_PICTURE",
		&returns,
	)
	return returns
}

// A scope for ResourceServer.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type ResourceServerScope interface {
	// A description of the scope.
	// Experimental.
	ScopeDescription() *string
	// The name of the scope.
	// Experimental.
	ScopeName() *string
}

// The jsii proxy struct for ResourceServerScope
type jsiiProxy_ResourceServerScope struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourceServerScope) ScopeDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceServerScope) ScopeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewResourceServerScope(props *ResourceServerScopeProps) ResourceServerScope {
	_init_.Initialize()

	j := jsiiProxy_ResourceServerScope{}

	_jsii_.Create(
		"monocdk.aws_cognito.ResourceServerScope",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewResourceServerScope_Override(r ResourceServerScope, props *ResourceServerScopeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.ResourceServerScope",
		[]interface{}{props},
		r,
	)
}

// Props to initialize ResourceServerScope.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type ResourceServerScopeProps struct {
	// A description of the scope.
	// Experimental.
	ScopeDescription *string `json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	// Experimental.
	ScopeName *string `json:"scopeName" yaml:"scopeName"`
}

// The different ways in which users of this pool can sign up or sign in.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	// ...
//   	signInAliases: &signInAliases{
//   		username: jsii.Boolean(true),
//   		email: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type SignInAliases struct {
	// Whether a user is allowed to sign up or sign in with an email address.
	// Experimental.
	Email *bool `json:"email" yaml:"email"`
	// Whether a user is allowed to sign up or sign in with a phone number.
	// Experimental.
	Phone *bool `json:"phone" yaml:"phone"`
	// Whether a user is allowed to sign in with a secondary username, that can be set and modified after sign up.
	//
	// Can only be used in conjunction with `USERNAME`.
	// Experimental.
	PreferredUsername *bool `json:"preferredUsername" yaml:"preferredUsername"`
	// Whether user is allowed to sign up or sign in with a username.
	// Experimental.
	Username *bool `json:"username" yaml:"username"`
}

// Options to customize the behaviour of `signInUrl()`.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			implicitCodeGrant: jsii.Boolean(true),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &userPoolDomainOptions{
//   })
//   signInUrl := domain.signInUrl(client, &signInUrlOptions{
//   	redirectUri: jsii.String("https://myapp.com/home"),
//   })
//
// Experimental.
type SignInUrlOptions struct {
	// Where to redirect to after sign in.
	// Experimental.
	RedirectUri *string `json:"redirectUri" yaml:"redirectUri"`
	// The path in the URI where the sign-in page is located.
	// Experimental.
	SignInPath *string `json:"signInPath" yaml:"signInPath"`
}

// Standard attribute that can be marked as required or mutable.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes
//
// Experimental.
type StandardAttribute struct {
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, this must be set to `true`.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
	// Specifies whether the attribute is required upon user registration.
	//
	// If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	// Experimental.
	Required *bool `json:"required" yaml:"required"`
}

// The set of standard attributes that can be marked as required or mutable.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes
//
// Experimental.
type StandardAttributes struct {
	// The user's postal address.
	// Experimental.
	Address *StandardAttribute `json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	// Experimental.
	Birthdate *StandardAttribute `json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	// Experimental.
	Email *StandardAttribute `json:"email" yaml:"email"`
	// DEPRECATED.
	// Deprecated: this is not a standard attribute and was incorrectly added to the CDK.
	// It is a Cognito built-in attribute and cannot be controlled as part of user pool creation.
	EmailVerified *StandardAttribute `json:"emailVerified" yaml:"emailVerified"`
	// The surname or last name of the user.
	// Experimental.
	FamilyName *StandardAttribute `json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	// Experimental.
	Fullname *StandardAttribute `json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Experimental.
	Gender *StandardAttribute `json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Experimental.
	GivenName *StandardAttribute `json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	// Experimental.
	LastUpdateTime *StandardAttribute `json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	// Experimental.
	Locale *StandardAttribute `json:"locale" yaml:"locale"`
	// The user's middle name.
	// Experimental.
	MiddleName *StandardAttribute `json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Experimental.
	Nickname *StandardAttribute `json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Experimental.
	PhoneNumber *StandardAttribute `json:"phoneNumber" yaml:"phoneNumber"`
	// DEPRECATED.
	// Deprecated: this is not a standard attribute and was incorrectly added to the CDK.
	// It is a Cognito built-in attribute and cannot be controlled as part of user pool creation.
	PhoneNumberVerified *StandardAttribute `json:"phoneNumberVerified" yaml:"phoneNumberVerified"`
	// The user's preffered username, different from the immutable user name.
	// Experimental.
	PreferredUsername *StandardAttribute `json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Experimental.
	ProfilePage *StandardAttribute `json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Experimental.
	ProfilePicture *StandardAttribute `json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Experimental.
	Timezone *StandardAttribute `json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Experimental.
	Website *StandardAttribute `json:"website" yaml:"website"`
}

// This interface contains standard attributes recognized by Cognito from https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html including built-in attributes `email_verified` and `phone_number_verified`.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   clientWriteAttributes := (cognito.NewClientAttributes()).withStandardAttributes(&standardAttributesMask{
//   	fullname: jsii.Boolean(true),
//   	email: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("favouritePizza"), jsii.String("favouriteBeverage"))
//
//   clientReadAttributes := clientWriteAttributes.withStandardAttributes(&standardAttributesMask{
//   	emailVerified: jsii.Boolean(true),
//   }).withCustomAttributes(jsii.String("pointsEarned"))
//
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	readAttributes: clientReadAttributes,
//   	writeAttributes: clientWriteAttributes,
//   })
//
// Experimental.
type StandardAttributesMask struct {
	// The user's postal address.
	// Experimental.
	Address *bool `json:"address" yaml:"address"`
	// The user's birthday, represented as an ISO 8601:2004 format.
	// Experimental.
	Birthdate *bool `json:"birthdate" yaml:"birthdate"`
	// The user's e-mail address, represented as an RFC 5322 [RFC5322] addr-spec.
	// Experimental.
	Email *bool `json:"email" yaml:"email"`
	// Whether the email address has been verified.
	// Experimental.
	EmailVerified *bool `json:"emailVerified" yaml:"emailVerified"`
	// The surname or last name of the user.
	// Experimental.
	FamilyName *bool `json:"familyName" yaml:"familyName"`
	// The user's full name in displayable form, including all name parts, titles and suffixes.
	// Experimental.
	Fullname *bool `json:"fullname" yaml:"fullname"`
	// The user's gender.
	// Experimental.
	Gender *bool `json:"gender" yaml:"gender"`
	// The user's first name or give name.
	// Experimental.
	GivenName *bool `json:"givenName" yaml:"givenName"`
	// The time, the user's information was last updated.
	// Experimental.
	LastUpdateTime *bool `json:"lastUpdateTime" yaml:"lastUpdateTime"`
	// The user's locale, represented as a BCP47 [RFC5646] language tag.
	// Experimental.
	Locale *bool `json:"locale" yaml:"locale"`
	// The user's middle name.
	// Experimental.
	MiddleName *bool `json:"middleName" yaml:"middleName"`
	// The user's nickname or casual name.
	// Experimental.
	Nickname *bool `json:"nickname" yaml:"nickname"`
	// The user's telephone number.
	// Experimental.
	PhoneNumber *bool `json:"phoneNumber" yaml:"phoneNumber"`
	// Whether the phone number has been verified.
	// Experimental.
	PhoneNumberVerified *bool `json:"phoneNumberVerified" yaml:"phoneNumberVerified"`
	// The user's preffered username, different from the immutable user name.
	// Experimental.
	PreferredUsername *bool `json:"preferredUsername" yaml:"preferredUsername"`
	// The URL to the user's profile page.
	// Experimental.
	ProfilePage *bool `json:"profilePage" yaml:"profilePage"`
	// The URL to the user's profile picture.
	// Experimental.
	ProfilePicture *bool `json:"profilePicture" yaml:"profilePicture"`
	// The user's time zone.
	// Experimental.
	Timezone *bool `json:"timezone" yaml:"timezone"`
	// The URL to the user's web page or blog.
	// Experimental.
	Website *bool `json:"website" yaml:"website"`
}

// The String custom attribute type.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type StringAttribute interface {
	ICustomAttribute
	// Bind this custom attribute type to the values as expected by CloudFormation.
	// Experimental.
	Bind() *CustomAttributeConfig
}

// The jsii proxy struct for StringAttribute
type jsiiProxy_StringAttribute struct {
	jsiiProxy_ICustomAttribute
}

// Experimental.
func NewStringAttribute(props *StringAttributeProps) StringAttribute {
	_init_.Initialize()

	j := jsiiProxy_StringAttribute{}

	_jsii_.Create(
		"monocdk.aws_cognito.StringAttribute",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewStringAttribute_Override(s StringAttribute, props *StringAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.StringAttribute",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StringAttribute) Bind() *CustomAttributeConfig {
	var returns *CustomAttributeConfig

	_jsii_.Invoke(
		s,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Constraints that can be applied to a custom attribute of string type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//   stringAttributeConstraints := &stringAttributeConstraints{
//   	maxLen: jsii.Number(123),
//   	minLen: jsii.Number(123),
//   }
//
// Experimental.
type StringAttributeConstraints struct {
	// Maximum length of this attribute.
	// Experimental.
	MaxLen *float64 `json:"maxLen" yaml:"maxLen"`
	// Minimum length of this attribute.
	// Experimental.
	MinLen *float64 `json:"minLen" yaml:"minLen"`
}

// Props for constructing a StringAttr.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// Experimental.
type StringAttributeProps struct {
	// Maximum length of this attribute.
	// Experimental.
	MaxLen *float64 `json:"maxLen" yaml:"maxLen"`
	// Minimum length of this attribute.
	// Experimental.
	MinLen *float64 `json:"minLen" yaml:"minLen"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
}

// User pool configuration when administrators sign users up.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	userInvitation: &userInvitationConfig{
//   		emailSubject: jsii.String("Invite to join our awesome app!"),
//   		emailBody: jsii.String("Hello {username}, you have been invited to join our awesome app! Your temporary password is {####}"),
//   		smsMessage: jsii.String("Hello {username}, your temporary password for our awesome app is {####}"),
//   	},
//   })
//
// Experimental.
type UserInvitationConfig struct {
	// The template to the email body that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	EmailBody *string `json:"emailBody" yaml:"emailBody"`
	// The template to the email subject that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	EmailSubject *string `json:"emailSubject" yaml:"emailSubject"`
	// The template to the SMS message that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	SmsMessage *string `json:"smsMessage" yaml:"smsMessage"`
}

// Define a Cognito User Pool.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_OPENID(),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://my-app-domain.com/welcome"),
//   		},
//   		logoutUrls: []*string{
//   			jsii.String("https://my-app-domain.com/signin"),
//   		},
//   	},
//   })
//
// Experimental.
type UserPool interface {
	awscdk.Resource
	IUserPool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// Get all identity providers registered with this user pool.
	// Experimental.
	IdentityProviders() *[]IUserPoolIdentityProvider
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The ARN of the user pool.
	// Experimental.
	UserPoolArn() *string
	// The physical ID of this user pool resource.
	// Experimental.
	UserPoolId() *string
	// User pool provider name.
	// Experimental.
	UserPoolProviderName() *string
	// User pool provider URL.
	// Experimental.
	UserPoolProviderUrl() *string
	// Add a new app client to this user pool.
	// Experimental.
	AddClient(id *string, options *UserPoolClientOptions) UserPoolClient
	// Associate a domain to this user pool.
	// Experimental.
	AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain
	// Add a new resource server to this user pool.
	// Experimental.
	AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer
	// Add a lambda trigger to a user pool operation.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	//
	// Experimental.
	AddTrigger(operation UserPoolOperation, fn awslambda.IFunction)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Register an identity provider with this user pool.
	// Experimental.
	RegisterIdentityProvider(provider IUserPoolIdentityProvider)
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPool
type jsiiProxy_UserPool struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPool
}

func (j *jsiiProxy_UserPool) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) IdentityProviders() *[]IUserPoolIdentityProvider {
	var returns *[]IUserPoolIdentityProvider
	_jsii_.Get(
		j,
		"identityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolProviderUrl",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPool(scope constructs.Construct, id *string, props *UserPoolProps) UserPool {
	_init_.Initialize()

	j := jsiiProxy_UserPool{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPool_Override(u UserPool, scope constructs.Construct, id *string, props *UserPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPool",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import an existing user pool based on its ARN.
// Experimental.
func UserPool_FromUserPoolArn(scope constructs.Construct, id *string, userPoolArn *string) IUserPool {
	_init_.Initialize()

	var returns IUserPool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPool",
		"fromUserPoolArn",
		[]interface{}{scope, id, userPoolArn},
		&returns,
	)

	return returns
}

// Import an existing user pool based on its id.
// Experimental.
func UserPool_FromUserPoolId(scope constructs.Construct, id *string, userPoolId *string) IUserPool {
	_init_.Initialize()

	var returns IUserPool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPool",
		"fromUserPoolId",
		[]interface{}{scope, id, userPoolId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func UserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPool_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPool",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddClient(id *string, options *UserPoolClientOptions) UserPoolClient {
	var returns UserPoolClient

	_jsii_.Invoke(
		u,
		"addClient",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain {
	var returns UserPoolDomain

	_jsii_.Invoke(
		u,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer {
	var returns UserPoolResourceServer

	_jsii_.Invoke(
		u,
		"addResourceServer",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddTrigger(operation UserPoolOperation, fn awslambda.IFunction) {
	_jsii_.InvokeVoid(
		u,
		"addTrigger",
		[]interface{}{operation, fn},
	)
}

func (u *jsiiProxy_UserPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPool) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPool) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPool) RegisterIdentityProvider(provider IUserPoolIdentityProvider) {
	_jsii_.InvokeVoid(
		u,
		"registerIdentityProvider",
		[]interface{}{provider},
	)
}

func (u *jsiiProxy_UserPool) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Define a UserPool App Client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   provider := cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	userPool: pool,
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   })
//
//   client := pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	supportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   	},
//   })
//
//   client.node.addDependency(provider)
//
// Experimental.
type UserPoolClient interface {
	awscdk.Resource
	IUserPoolClient
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The OAuth flows enabled for this client.
	// Experimental.
	OAuthFlows() *OAuthFlows
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Name of the application client.
	// Experimental.
	UserPoolClientId() *string
	// The client name that was specified via the `userPoolClientName` property during initialization, throws an error otherwise.
	// Experimental.
	UserPoolClientName() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolClient
type jsiiProxy_UserPoolClient struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolClient
}

func (j *jsiiProxy_UserPoolClient) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) OAuthFlows() *OAuthFlows {
	var returns *OAuthFlows
	_jsii_.Get(
		j,
		"oAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) UserPoolClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientName",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolClient(scope constructs.Construct, id *string, props *UserPoolClientProps) UserPoolClient {
	_init_.Initialize()

	j := jsiiProxy_UserPoolClient{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolClient",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolClient_Override(u UserPoolClient, scope constructs.Construct, id *string, props *UserPoolClientProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolClient",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a user pool client given its id.
// Experimental.
func UserPoolClient_FromUserPoolClientId(scope constructs.Construct, id *string, userPoolClientId *string) IUserPoolClient {
	_init_.Initialize()

	var returns IUserPoolClient

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolClient",
		"fromUserPoolClientId",
		[]interface{}{scope, id, userPoolClientId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolClient_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolClient",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolClient) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolClient) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolClient) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolClient) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Identity providers supported by the UserPoolClient.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	supportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   		cognito.*userPoolClientIdentityProvider_COGNITO(),
//   	},
//   })
//
// Experimental.
type UserPoolClientIdentityProvider interface {
	// The name of the identity provider as recognized by CloudFormation property `SupportedIdentityProviders`.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for UserPoolClientIdentityProvider
type jsiiProxy_UserPoolClientIdentityProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_UserPoolClientIdentityProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Specify a provider not yet supported by the CDK.
// Experimental.
func UserPoolClientIdentityProvider_Custom(name *string) UserPoolClientIdentityProvider {
	_init_.Initialize()

	var returns UserPoolClientIdentityProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"custom",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func UserPoolClientIdentityProvider_AMAZON() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"AMAZON",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_APPLE() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"APPLE",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_COGNITO() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"COGNITO",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_FACEBOOK() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"FACEBOOK",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_GOOGLE() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolClientIdentityProvider",
		"GOOGLE",
		&returns,
	)
	return returns
}

// Options to create a UserPoolClient.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_OPENID(),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://my-app-domain.com/welcome"),
//   		},
//   		logoutUrls: []*string{
//   			jsii.String("https://my-app-domain.com/signin"),
//   		},
//   	},
//   })
//
// Experimental.
type UserPoolClientOptions struct {
	// Validity of the access token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-access-token
	//
	// Experimental.
	AccessTokenValidity awscdk.Duration `json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The set of OAuth authentication flows to enable on the client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
	//
	// Experimental.
	AuthFlows *AuthFlow `json:"authFlows" yaml:"authFlows"`
	// Turns off all OAuth interactions for this client.
	// Experimental.
	DisableOAuth *bool `json:"disableOAuth" yaml:"disableOAuth"`
	// Enable token revocation for this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html#enable-token-revocation
	//
	// Experimental.
	EnableTokenRevocation *bool `json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// Whether to generate a client secret.
	// Experimental.
	GenerateSecret *bool `json:"generateSecret" yaml:"generateSecret"`
	// Validity of the ID token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-id-token
	//
	// Experimental.
	IdTokenValidity awscdk.Duration `json:"idTokenValidity" yaml:"idTokenValidity"`
	// OAuth settings for this client to interact with the app.
	//
	// An error is thrown when this is specified and `disableOAuth` is set.
	// Experimental.
	OAuth *OAuthSettings `json:"oAuth" yaml:"oAuth"`
	// Whether Cognito returns a UserNotFoundException exception when the user does not exist in the user pool (false), or whether it returns another type of error that doesn't reveal the user's absence.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-managing-errors.html
	//
	// Experimental.
	PreventUserExistenceErrors *bool `json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The set of attributes this client will be able to read.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	ReadAttributes ClientAttributes `json:"readAttributes" yaml:"readAttributes"`
	// Validity of the refresh token.
	//
	// Values between 60 minutes and 10 years are valid.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-refresh-token
	//
	// Experimental.
	RefreshTokenValidity awscdk.Duration `json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// The list of identity providers that users should be able to use to sign in using this client.
	// Experimental.
	SupportedIdentityProviders *[]UserPoolClientIdentityProvider `json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// Name of the application client.
	// Experimental.
	UserPoolClientName *string `json:"userPoolClientName" yaml:"userPoolClientName"`
	// The set of attributes this client will be able to write.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	WriteAttributes ClientAttributes `json:"writeAttributes" yaml:"writeAttributes"`
}

// Properties for the UserPoolClient construct.
//
// Example:
//   importedPool := cognito.userPool.fromUserPoolId(this, jsii.String("imported-pool"), jsii.String("us-east-1_oiuR12Abd"))
//   cognito.NewUserPoolClient(this, jsii.String("customer-app-client"), &userPoolClientProps{
//   	userPool: importedPool,
//   })
//
// Experimental.
type UserPoolClientProps struct {
	// Validity of the access token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-access-token
	//
	// Experimental.
	AccessTokenValidity awscdk.Duration `json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The set of OAuth authentication flows to enable on the client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
	//
	// Experimental.
	AuthFlows *AuthFlow `json:"authFlows" yaml:"authFlows"`
	// Turns off all OAuth interactions for this client.
	// Experimental.
	DisableOAuth *bool `json:"disableOAuth" yaml:"disableOAuth"`
	// Enable token revocation for this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html#enable-token-revocation
	//
	// Experimental.
	EnableTokenRevocation *bool `json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// Whether to generate a client secret.
	// Experimental.
	GenerateSecret *bool `json:"generateSecret" yaml:"generateSecret"`
	// Validity of the ID token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-id-token
	//
	// Experimental.
	IdTokenValidity awscdk.Duration `json:"idTokenValidity" yaml:"idTokenValidity"`
	// OAuth settings for this client to interact with the app.
	//
	// An error is thrown when this is specified and `disableOAuth` is set.
	// Experimental.
	OAuth *OAuthSettings `json:"oAuth" yaml:"oAuth"`
	// Whether Cognito returns a UserNotFoundException exception when the user does not exist in the user pool (false), or whether it returns another type of error that doesn't reveal the user's absence.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-managing-errors.html
	//
	// Experimental.
	PreventUserExistenceErrors *bool `json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The set of attributes this client will be able to read.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	ReadAttributes ClientAttributes `json:"readAttributes" yaml:"readAttributes"`
	// Validity of the refresh token.
	//
	// Values between 60 minutes and 10 years are valid.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-refresh-token
	//
	// Experimental.
	RefreshTokenValidity awscdk.Duration `json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// The list of identity providers that users should be able to use to sign in using this client.
	// Experimental.
	SupportedIdentityProviders *[]UserPoolClientIdentityProvider `json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// Name of the application client.
	// Experimental.
	UserPoolClientName *string `json:"userPoolClientName" yaml:"userPoolClientName"`
	// The set of attributes this client will be able to write.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	WriteAttributes ClientAttributes `json:"writeAttributes" yaml:"writeAttributes"`
	// The UserPool resource this client will have access to.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
}

// Define a user pool domain.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			implicitCodeGrant: jsii.Boolean(true),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &userPoolDomainOptions{
//   })
//   signInUrl := domain.signInUrl(client, &signInUrlOptions{
//   	redirectUri: jsii.String("https://myapp.com/home"),
//   })
//
// Experimental.
type UserPoolDomain interface {
	awscdk.Resource
	IUserPoolDomain
	// The domain name of the CloudFront distribution associated with the user pool domain.
	// Experimental.
	CloudFrontDomainName() *string
	// The domain that was specified to be created.
	//
	// If `customDomain` was selected, this holds the full domain name that was specified.
	// If the `cognitoDomain` was used, it contains the prefix to the Cognito hosted domain.
	// Experimental.
	DomainName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// The URL to the hosted UI associated with this domain.
	// Experimental.
	BaseUrl() *string
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// The URL to the sign in page in this domain using a specific UserPoolClient.
	// Experimental.
	SignInUrl(client UserPoolClient, options *SignInUrlOptions) *string
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolDomain
type jsiiProxy_UserPoolDomain struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolDomain
}

func (j *jsiiProxy_UserPoolDomain) CloudFrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolDomain(scope constructs.Construct, id *string, props *UserPoolDomainProps) UserPoolDomain {
	_init_.Initialize()

	j := jsiiProxy_UserPoolDomain{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolDomain_Override(u UserPoolDomain, scope constructs.Construct, id *string, props *UserPoolDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolDomain",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a UserPoolDomain given its domain name.
// Experimental.
func UserPoolDomain_FromDomainName(scope constructs.Construct, id *string, userPoolDomainName *string) IUserPoolDomain {
	_init_.Initialize()

	var returns IUserPoolDomain

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolDomain",
		"fromDomainName",
		[]interface{}{scope, id, userPoolDomainName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolDomain_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolDomain",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolDomain) BaseUrl() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"baseUrl",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolDomain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolDomain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolDomain) SignInUrl(client UserPoolClient, options *SignInUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"signInUrl",
		[]interface{}{client, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to create a UserPoolDomain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   pool.addDomain(jsii.String("CognitoDomain"), &userPoolDomainOptions{
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("my-awesome-app"),
//   	},
//   })
//
//   certificateArn := "arn:aws:acm:us-east-1:123456789012:certificate/11-3336f1-44483d-adc7-9cd375c5169d"
//
//   domainCert := certificatemanager.certificate.fromCertificateArn(this, jsii.String("domainCert"), certificateArn)
//   pool.addDomain(jsii.String("CustomDomain"), &userPoolDomainOptions{
//   	customDomain: &customDomainOptions{
//   		domainName: jsii.String("user.myapp.com"),
//   		certificate: domainCert,
//   	},
//   })
//
// Experimental.
type UserPoolDomainOptions struct {
	// Associate a cognito prefix domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
	//
	// Experimental.
	CognitoDomain *CognitoDomainOptions `json:"cognitoDomain" yaml:"cognitoDomain"`
	// Associate a custom domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
	//
	// Experimental.
	CustomDomain *CustomDomainOptions `json:"customDomain" yaml:"customDomain"`
}

// Props for UserPoolDomain construct.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk"import elbv2 "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"type App awscdk.App
//   type CfnOutput awscdk.CfnOutput
//   type Stack awscdk.Stackimport constructs "github.com/aws/constructs-go/constructs"type Construct constructs.Constructimport actions "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStack struct {
//   stack
//   }
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &userPoolClientProps{
//   	userPool: userPool,
//
//   	// Required minimal configuration for use with an ELB
//   	generateSecret: jsii.Boolean(true),
//   	authFlows: &authFlow{
//   		userPassword: jsii.Boolean(true),
//   	},
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_EMAIL(),
//   		},
//   		callbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.node.defaultChild.(cfnUserPoolClient)
//   cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &userPoolDomainProps{
//   	userPool: userPool,
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(443),
//   	certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	defaultAction: actions.NewAuthenticateCognitoAction(&authenticateCognitoActionProps{
//   		userPool: userPool,
//   		userPoolClient: userPoolClient,
//   		userPoolDomain: userPoolDomain,
//   		next: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   			contentType: jsii.String("text/plain"),
//   			messageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
//   	value: lb.loadBalancerDnsName,
//   })
//
//   app := NewApp()
//   NewCognitoStack(app, jsii.String("integ-cognito"))
//   app.synth()
//
// Experimental.
type UserPoolDomainProps struct {
	// Associate a cognito prefix domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
	//
	// Experimental.
	CognitoDomain *CognitoDomainOptions `json:"cognitoDomain" yaml:"cognitoDomain"`
	// Associate a custom domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
	//
	// Experimental.
	CustomDomain *CustomDomainOptions `json:"customDomain" yaml:"customDomain"`
	// The user pool to which this domain should be associated.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
}

// Configure how Cognito sends emails.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	email: cognito.userPoolEmail.withSES(&userPoolSESOptions{
//   		fromEmail: jsii.String("noreply@myawesomeapp.com"),
//   		fromName: jsii.String("Awesome App"),
//   		replyTo: jsii.String("support@myawesomeapp.com"),
//   	}),
//   })
//
// Experimental.
type UserPoolEmail interface {
}

// The jsii proxy struct for UserPoolEmail
type jsiiProxy_UserPoolEmail struct {
	_ byte // padding
}

// Experimental.
func NewUserPoolEmail_Override(u UserPoolEmail) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolEmail",
		nil, // no parameters
		u,
	)
}

// Send email using Cognito.
// Experimental.
func UserPoolEmail_WithCognito(replyTo *string) UserPoolEmail {
	_init_.Initialize()

	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolEmail",
		"withCognito",
		[]interface{}{replyTo},
		&returns,
	)

	return returns
}

// Send email using SES.
// Experimental.
func UserPoolEmail_WithSES(options *UserPoolSESOptions) UserPoolEmail {
	_init_.Initialize()

	var returns UserPoolEmail

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolEmail",
		"withSES",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// User pool third-party identity providers.
// Experimental.
type UserPoolIdentityProvider interface {
}

// The jsii proxy struct for UserPoolIdentityProvider
type jsiiProxy_UserPoolIdentityProvider struct {
	_ byte // padding
}

// Import an existing UserPoolIdentityProvider.
// Experimental.
func UserPoolIdentityProvider_FromProviderName(scope constructs.Construct, id *string, providerName *string) IUserPoolIdentityProvider {
	_init_.Initialize()

	var returns IUserPoolIdentityProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProvider",
		"fromProviderName",
		[]interface{}{scope, id, providerName},
		&returns,
	)

	return returns
}

// Represents a identity provider that integrates with 'Login with Amazon'.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   provider := cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	userPool: pool,
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   })
//
//   client := pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	supportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   	},
//   })
//
//   client.node.addDependency(provider)
//
// Experimental.
type UserPoolIdentityProviderAmazon interface {
	awscdk.Resource
	IUserPoolIdentityProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	ConfigureAttributeMapping() interface{}
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolIdentityProviderAmazon
type jsiiProxy_UserPoolIdentityProviderAmazon struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolIdentityProvider
}

func (j *jsiiProxy_UserPoolIdentityProviderAmazon) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderAmazon) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderAmazon) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderAmazon) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderAmazon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolIdentityProviderAmazon(scope constructs.Construct, id *string, props *UserPoolIdentityProviderAmazonProps) UserPoolIdentityProviderAmazon {
	_init_.Initialize()

	j := jsiiProxy_UserPoolIdentityProviderAmazon{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderAmazon",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolIdentityProviderAmazon_Override(u UserPoolIdentityProviderAmazon, scope constructs.Construct, id *string, props *UserPoolIdentityProviderAmazonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderAmazon",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolIdentityProviderAmazon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderAmazon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolIdentityProviderAmazon_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderAmazon",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) ConfigureAttributeMapping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"configureAttributeMapping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to initialize UserPoolAmazonIdentityProvider.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   provider := cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	userPool: pool,
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   })
//
//   client := pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	supportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   	},
//   })
//
//   client.node.addDependency(provider)
//
// Experimental.
type UserPoolIdentityProviderAmazonProps struct {
	// The user pool to which this construct provides identities.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Experimental.
	AttributeMapping *AttributeMapping `json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by 'Login with Amazon' APIs.
	// See: https://developer.amazon.com/docs/login-with-amazon/security-profile.html#client-identifier
	//
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for 'Login with Amazon' APIs to authenticate the client.
	// See: https://developer.amazon.com/docs/login-with-amazon/security-profile.html#client-identifier
	//
	// Experimental.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The types of user profile data to obtain for the Amazon profile.
	// See: https://developer.amazon.com/docs/login-with-amazon/customer-profile.html
	//
	// Experimental.
	Scopes *[]*string `json:"scopes" yaml:"scopes"`
}

// Represents a identity provider that integrates with 'Apple'.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderApple := cognito.NewUserPoolIdentityProviderApple(this, jsii.String("MyUserPoolIdentityProviderApple"), &userPoolIdentityProviderAppleProps{
//   	clientId: jsii.String("clientId"),
//   	keyId: jsii.String("keyId"),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   })
//
// Experimental.
type UserPoolIdentityProviderApple interface {
	awscdk.Resource
	IUserPoolIdentityProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	ConfigureAttributeMapping() interface{}
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolIdentityProviderApple
type jsiiProxy_UserPoolIdentityProviderApple struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolIdentityProvider
}

func (j *jsiiProxy_UserPoolIdentityProviderApple) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderApple) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderApple) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderApple) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderApple) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolIdentityProviderApple(scope constructs.Construct, id *string, props *UserPoolIdentityProviderAppleProps) UserPoolIdentityProviderApple {
	_init_.Initialize()

	j := jsiiProxy_UserPoolIdentityProviderApple{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderApple",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolIdentityProviderApple_Override(u UserPoolIdentityProviderApple, scope constructs.Construct, id *string, props *UserPoolIdentityProviderAppleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderApple",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolIdentityProviderApple_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderApple",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolIdentityProviderApple_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderApple",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) ConfigureAttributeMapping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"configureAttributeMapping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to initialize UserPoolAppleIdentityProvider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderAppleProps := &userPoolIdentityProviderAppleProps{
//   	clientId: jsii.String("clientId"),
//   	keyId: jsii.String("keyId"),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// Experimental.
type UserPoolIdentityProviderAppleProps struct {
	// The user pool to which this construct provides identities.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Experimental.
	AttributeMapping *AttributeMapping `json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Apple APIs.
	// See: https://developer.apple.com/documentation/sign_in_with_apple/clientconfigi/3230948-clientid
	//
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The keyId (of the same key, which content has to be later supplied as `privateKey`) for Apple APIs to authenticate the client.
	// Experimental.
	KeyId *string `json:"keyId" yaml:"keyId"`
	// The privateKey content for Apple APIs to authenticate the client.
	// Experimental.
	PrivateKey *string `json:"privateKey" yaml:"privateKey"`
	// The teamId for Apple APIs to authenticate the client.
	// Experimental.
	TeamId *string `json:"teamId" yaml:"teamId"`
	// The list of apple permissions to obtain for getting access to the apple profile.
	// See: https://developer.apple.com/documentation/sign_in_with_apple/clientconfigi/3230955-scope
	//
	// Experimental.
	Scopes *[]*string `json:"scopes" yaml:"scopes"`
}

// Represents a identity provider that integrates with 'Facebook Login'.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderFacebook := cognito.NewUserPoolIdentityProviderFacebook(this, jsii.String("MyUserPoolIdentityProviderFacebook"), &userPoolIdentityProviderFacebookProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   })
//
// Experimental.
type UserPoolIdentityProviderFacebook interface {
	awscdk.Resource
	IUserPoolIdentityProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	ConfigureAttributeMapping() interface{}
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolIdentityProviderFacebook
type jsiiProxy_UserPoolIdentityProviderFacebook struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolIdentityProvider
}

func (j *jsiiProxy_UserPoolIdentityProviderFacebook) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderFacebook) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderFacebook) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderFacebook) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderFacebook) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolIdentityProviderFacebook(scope constructs.Construct, id *string, props *UserPoolIdentityProviderFacebookProps) UserPoolIdentityProviderFacebook {
	_init_.Initialize()

	j := jsiiProxy_UserPoolIdentityProviderFacebook{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderFacebook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolIdentityProviderFacebook_Override(u UserPoolIdentityProviderFacebook, scope constructs.Construct, id *string, props *UserPoolIdentityProviderFacebookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderFacebook",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolIdentityProviderFacebook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderFacebook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolIdentityProviderFacebook_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderFacebook",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) ConfigureAttributeMapping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"configureAttributeMapping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to initialize UserPoolFacebookIdentityProvider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderFacebookProps := &userPoolIdentityProviderFacebookProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// Experimental.
type UserPoolIdentityProviderFacebookProps struct {
	// The user pool to which this construct provides identities.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Experimental.
	AttributeMapping *AttributeMapping `json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Facebook APIs.
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientUd for Facebook to authenticate the client.
	// See: https://developers.facebook.com/docs/facebook-login/security#appsecret
	//
	// Experimental.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The Facebook API version to use.
	// Experimental.
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	// The list of facebook permissions to obtain for getting access to the Facebook profile.
	// See: https://developers.facebook.com/docs/facebook-login/permissions
	//
	// Experimental.
	Scopes *[]*string `json:"scopes" yaml:"scopes"`
}

// Represents a identity provider that integrates with 'Google'.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderGoogle := cognito.NewUserPoolIdentityProviderGoogle(this, jsii.String("MyUserPoolIdentityProviderGoogle"), &userPoolIdentityProviderGoogleProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   })
//
// Experimental.
type UserPoolIdentityProviderGoogle interface {
	awscdk.Resource
	IUserPoolIdentityProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The primary identifier of this identity provider.
	// Experimental.
	ProviderName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	ConfigureAttributeMapping() interface{}
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolIdentityProviderGoogle
type jsiiProxy_UserPoolIdentityProviderGoogle struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolIdentityProvider
}

func (j *jsiiProxy_UserPoolIdentityProviderGoogle) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderGoogle) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderGoogle) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderGoogle) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderGoogle) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolIdentityProviderGoogle(scope constructs.Construct, id *string, props *UserPoolIdentityProviderGoogleProps) UserPoolIdentityProviderGoogle {
	_init_.Initialize()

	j := jsiiProxy_UserPoolIdentityProviderGoogle{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderGoogle",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolIdentityProviderGoogle_Override(u UserPoolIdentityProviderGoogle, scope constructs.Construct, id *string, props *UserPoolIdentityProviderGoogleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolIdentityProviderGoogle",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolIdentityProviderGoogle_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderGoogle",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolIdentityProviderGoogle_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolIdentityProviderGoogle",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) ConfigureAttributeMapping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"configureAttributeMapping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to initialize UserPoolGoogleIdentityProvider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderGoogleProps := &userPoolIdentityProviderGoogleProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// Experimental.
type UserPoolIdentityProviderGoogleProps struct {
	// The user pool to which this construct provides identities.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Experimental.
	AttributeMapping *AttributeMapping `json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Google APIs.
	// See: https://developers.google.com/identity/sign-in/web/sign-in#specify_your_apps_client_id
	//
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for Google APIs to authenticate the client.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Experimental.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The list of google permissions to obtain for getting access to the google profile.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Experimental.
	Scopes *[]*string `json:"scopes" yaml:"scopes"`
}

// Properties to create a new instance of UserPoolIdentityProvider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   userPoolIdentityProviderProps := &userPoolIdentityProviderProps{
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   }
//
// Experimental.
type UserPoolIdentityProviderProps struct {
	// The user pool to which this construct provides identities.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Experimental.
	AttributeMapping *AttributeMapping `json:"attributeMapping" yaml:"attributeMapping"`
}

// User pool operations to which lambda triggers can be attached.
//
// Example:
//   authChallengeFn := lambda.NewFunction(this, jsii.String("authChallengeFn"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   })
//
//   userpool := cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	lambdaTriggers: &userPoolTriggers{
//   		createAuthChallenge: authChallengeFn,
//   	},
//   })
//
//   userpool.addTrigger(cognito.userPoolOperation_USER_MIGRATION(), lambda.NewFunction(this, jsii.String("userMigrationFn"), &functionProps{
//   	runtime: lambda.*runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.*code.fromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   }))
//
// Experimental.
type UserPoolOperation interface {
	// The key to use in `CfnUserPool.LambdaConfigProperty`.
	// Experimental.
	OperationName() *string
}

// The jsii proxy struct for UserPoolOperation
type jsiiProxy_UserPoolOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_UserPoolOperation) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}


// A custom user pool operation.
// Experimental.
func UserPoolOperation_Of(name *string) UserPoolOperation {
	_init_.Initialize()

	var returns UserPoolOperation

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolOperation",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func UserPoolOperation_CREATE_AUTH_CHALLENGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"CREATE_AUTH_CHALLENGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_EMAIL_SENDER() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"CUSTOM_EMAIL_SENDER",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_MESSAGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"CUSTOM_MESSAGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_CUSTOM_SMS_SENDER() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"CUSTOM_SMS_SENDER",
		&returns,
	)
	return returns
}

func UserPoolOperation_DEFINE_AUTH_CHALLENGE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"DEFINE_AUTH_CHALLENGE",
		&returns,
	)
	return returns
}

func UserPoolOperation_POST_AUTHENTICATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"POST_AUTHENTICATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_POST_CONFIRMATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"POST_CONFIRMATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_AUTHENTICATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"PRE_AUTHENTICATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_SIGN_UP() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"PRE_SIGN_UP",
		&returns,
	)
	return returns
}

func UserPoolOperation_PRE_TOKEN_GENERATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"PRE_TOKEN_GENERATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_USER_MIGRATION() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"USER_MIGRATION",
		&returns,
	)
	return returns
}

func UserPoolOperation_VERIFY_AUTH_CHALLENGE_RESPONSE() UserPoolOperation {
	_init_.Initialize()
	var returns UserPoolOperation
	_jsii_.StaticGet(
		"monocdk.aws_cognito.UserPoolOperation",
		"VERIFY_AUTH_CHALLENGE_RESPONSE",
		&returns,
	)
	return returns
}

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
// Experimental.
type UserPoolProps struct {
	// How will a user be able to recover their account?
	// Experimental.
	AccountRecovery AccountRecovery `json:"accountRecovery" yaml:"accountRecovery"`
	// Attributes which Cognito will look to verify automatically upon user sign up.
	//
	// EMAIL and PHONE are the only available options.
	// Experimental.
	AutoVerify *AutoVerifiedAttrs `json:"autoVerify" yaml:"autoVerify"`
	// Define a set of custom attributes that can be configured for each user in the user pool.
	// Experimental.
	CustomAttributes *map[string]ICustomAttribute `json:"customAttributes" yaml:"customAttributes"`
	// This key will be used to encrypt temporary passwords and authorization codes that Amazon Cognito generates.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-sender-triggers.html
	//
	// Experimental.
	CustomSenderKmsKey awskms.IKey `json:"customSenderKmsKey" yaml:"customSenderKmsKey"`
	// Device tracking settings.
	// Experimental.
	DeviceTracking *DeviceTracking `json:"deviceTracking" yaml:"deviceTracking"`
	// Email settings for a user pool.
	// Experimental.
	Email UserPoolEmail `json:"email" yaml:"email"`
	// Email settings for a user pool.
	// Deprecated: Use 'email' instead.
	EmailSettings *EmailSettings `json:"emailSettings" yaml:"emailSettings"`
	// Setting this would explicitly enable or disable SMS role creation.
	//
	// When left unspecified, CDK will determine based on other properties if a role is needed or not.
	// Experimental.
	EnableSmsRole *bool `json:"enableSmsRole" yaml:"enableSmsRole"`
	// Lambda functions to use for supported Cognito triggers.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	//
	// Experimental.
	LambdaTriggers *UserPoolTriggers `json:"lambdaTriggers" yaml:"lambdaTriggers"`
	// Configure whether users of this user pool can or are required use MFA to sign in.
	// Experimental.
	Mfa Mfa `json:"mfa" yaml:"mfa"`
	// The SMS message template sent during MFA verification.
	//
	// Use '{####}' in the template where Cognito should insert the verification code.
	// Experimental.
	MfaMessage *string `json:"mfaMessage" yaml:"mfaMessage"`
	// Configure the MFA types that users can use in this user pool.
	//
	// Ignored if `mfa` is set to `OFF`.
	// Experimental.
	MfaSecondFactor *MfaSecondFactor `json:"mfaSecondFactor" yaml:"mfaSecondFactor"`
	// Password policy for this user pool.
	// Experimental.
	PasswordPolicy *PasswordPolicy `json:"passwordPolicy" yaml:"passwordPolicy"`
	// Policy to apply when the user pool is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Whether self sign up should be enabled.
	//
	// This can be further configured via the `selfSignUp` property.
	// Experimental.
	SelfSignUpEnabled *bool `json:"selfSignUpEnabled" yaml:"selfSignUpEnabled"`
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
	// Experimental.
	SignInAliases *SignInAliases `json:"signInAliases" yaml:"signInAliases"`
	// Whether sign-in aliases should be evaluated with case sensitivity.
	//
	// For example, when this option is set to false, users will be able to sign in using either `MyUsername` or `myusername`.
	// Experimental.
	SignInCaseSensitive *bool `json:"signInCaseSensitive" yaml:"signInCaseSensitive"`
	// The IAM role that Cognito will assume while sending SMS messages.
	// Experimental.
	SmsRole awsiam.IRole `json:"smsRole" yaml:"smsRole"`
	// The 'ExternalId' that Cognito service must using when assuming the `smsRole`, if the role is restricted with an 'sts:ExternalId' conditional.
	//
	// Learn more about ExternalId here - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	//
	// This property will be ignored if `smsRole` is not specified.
	// Experimental.
	SmsRoleExternalId *string `json:"smsRoleExternalId" yaml:"smsRoleExternalId"`
	// The region to integrate with SNS to send SMS messages.
	//
	// This property will do nothing if SMS configuration is not configured.
	// Experimental.
	SnsRegion *string `json:"snsRegion" yaml:"snsRegion"`
	// The set of attributes that are required for every user in the user pool.
	//
	// Read more on attributes here - https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html
	// Experimental.
	StandardAttributes *StandardAttributes `json:"standardAttributes" yaml:"standardAttributes"`
	// Configuration around admins signing up users into a user pool.
	// Experimental.
	UserInvitation *UserInvitationConfig `json:"userInvitation" yaml:"userInvitation"`
	// Name of the user pool.
	// Experimental.
	UserPoolName *string `json:"userPoolName" yaml:"userPoolName"`
	// Configuration around users signing themselves up to the user pool.
	//
	// Enable or disable self sign-up via the `selfSignUpEnabled` property.
	// Experimental.
	UserVerification *UserVerificationConfig `json:"userVerification" yaml:"userVerification"`
}

// Defines a User Pool OAuth2.0 Resource Server.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type UserPoolResourceServer interface {
	awscdk.Resource
	IUserPoolResourceServer
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Resource server id.
	// Experimental.
	UserPoolResourceServerId() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for UserPoolResourceServer
type jsiiProxy_UserPoolResourceServer struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolResourceServer
}

func (j *jsiiProxy_UserPoolResourceServer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolResourceServer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolResourceServer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolResourceServer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolResourceServer) UserPoolResourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolResourceServerId",
		&returns,
	)
	return returns
}


// Experimental.
func NewUserPoolResourceServer(scope constructs.Construct, id *string, props *UserPoolResourceServerProps) UserPoolResourceServer {
	_init_.Initialize()

	j := jsiiProxy_UserPoolResourceServer{}

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolResourceServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolResourceServer_Override(u UserPoolResourceServer, scope constructs.Construct, id *string, props *UserPoolResourceServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito.UserPoolResourceServer",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a user pool resource client given its id.
// Experimental.
func UserPoolResourceServer_FromUserPoolResourceServerId(scope constructs.Construct, id *string, userPoolResourceServerId *string) IUserPoolResourceServer {
	_init_.Initialize()

	var returns IUserPoolResourceServer

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolResourceServer",
		"fromUserPoolResourceServerId",
		[]interface{}{scope, id, userPoolResourceServerId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func UserPoolResourceServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolResourceServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func UserPoolResourceServer_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito.UserPoolResourceServer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolResourceServer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolResourceServer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolResourceServer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPoolResourceServer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UserPoolResourceServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolResourceServer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to create a UserPoolResourceServer.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type UserPoolResourceServerOptions struct {
	// A unique resource server identifier for the resource server.
	// Experimental.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	// Experimental.
	Scopes *[]ResourceServerScope `json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	// Experimental.
	UserPoolResourceServerName *string `json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
}

// Properties for the UserPoolResourceServer construct.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var resourceServerScope resourceServerScope
//   var userPool userPool
//   userPoolResourceServerProps := &userPoolResourceServerProps{
//   	identifier: jsii.String("identifier"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	scopes: []*resourceServerScope{
//   		resourceServerScope,
//   	},
//   	userPoolResourceServerName: jsii.String("userPoolResourceServerName"),
//   }
//
// Experimental.
type UserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	// Experimental.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	// Experimental.
	Scopes *[]ResourceServerScope `json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	// Experimental.
	UserPoolResourceServerName *string `json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
	// The user pool to add this resource server to.
	// Experimental.
	UserPool IUserPool `json:"userPool" yaml:"userPool"`
}

// Configuration for Cognito sending emails via Amazon SES.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	email: cognito.userPoolEmail.withSES(&userPoolSESOptions{
//   		fromEmail: jsii.String("noreply@myawesomeapp.com"),
//   		fromName: jsii.String("Awesome App"),
//   		replyTo: jsii.String("support@myawesomeapp.com"),
//   	}),
//   })
//
// Experimental.
type UserPoolSESOptions struct {
	// The verified Amazon SES email address that Cognito should use to send emails.
	//
	// The email address used must be a verified email address
	// in Amazon SES and must be configured to allow Cognito to
	// send emails.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-email.html
	//
	// Experimental.
	FromEmail *string `json:"fromEmail" yaml:"fromEmail"`
	// The name of a configuration set in Amazon SES that should be applied to emails sent via Cognito.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-configurationset
	//
	// Experimental.
	ConfigurationSetName *string `json:"configurationSetName" yaml:"configurationSetName"`
	// An optional name that should be used as the sender's name along with the email.
	// Experimental.
	FromName *string `json:"fromName" yaml:"fromName"`
	// The destination to which the receiver of the email should reploy to.
	// Experimental.
	ReplyTo *string `json:"replyTo" yaml:"replyTo"`
	// Required if the UserPool region is different than the SES region.
	//
	// If sending emails with a Amazon SES verified email address,
	// and the region that SES is configured is different than the
	// region in which the UserPool is deployed, you must specify that
	// region here.
	//
	// Must be 'us-east-1', 'us-west-2', or 'eu-west-1'.
	// Experimental.
	SesRegion *string `json:"sesRegion" yaml:"sesRegion"`
}

// Triggers for a user pool.
//
// Example:
//   authChallengeFn := lambda.NewFunction(this, jsii.String("authChallengeFn"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   })
//
//   userpool := cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	lambdaTriggers: &userPoolTriggers{
//   		createAuthChallenge: authChallengeFn,
//   	},
//   })
//
//   userpool.addTrigger(cognito.userPoolOperation_USER_MIGRATION(), lambda.NewFunction(this, jsii.String("userMigrationFn"), &functionProps{
//   	runtime: lambda.*runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.*code.fromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   }))
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
//
// Experimental.
type UserPoolTriggers struct {
	// Creates an authentication challenge.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-create-auth-challenge.html
	//
	// Experimental.
	CreateAuthChallenge awslambda.IFunction `json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// Amazon Cognito invokes this trigger to send email notifications to users.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-email-sender.html
	//
	// Experimental.
	CustomEmailSender awslambda.IFunction `json:"customEmailSender" yaml:"customEmailSender"`
	// A custom Message AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-message.html
	//
	// Experimental.
	CustomMessage awslambda.IFunction `json:"customMessage" yaml:"customMessage"`
	// Amazon Cognito invokes this trigger to send SMS notifications to users.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-sms-sender.html
	//
	// Experimental.
	CustomSmsSender awslambda.IFunction `json:"customSmsSender" yaml:"customSmsSender"`
	// Defines the authentication challenge.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-define-auth-challenge.html
	//
	// Experimental.
	DefineAuthChallenge awslambda.IFunction `json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// A post-authentication AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-authentication.html
	//
	// Experimental.
	PostAuthentication awslambda.IFunction `json:"postAuthentication" yaml:"postAuthentication"`
	// A post-confirmation AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-confirmation.html
	//
	// Experimental.
	PostConfirmation awslambda.IFunction `json:"postConfirmation" yaml:"postConfirmation"`
	// A pre-authentication AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-authentication.html
	//
	// Experimental.
	PreAuthentication awslambda.IFunction `json:"preAuthentication" yaml:"preAuthentication"`
	// A pre-registration AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html
	//
	// Experimental.
	PreSignUp awslambda.IFunction `json:"preSignUp" yaml:"preSignUp"`
	// A pre-token-generation AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html
	//
	// Experimental.
	PreTokenGeneration awslambda.IFunction `json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// A user-migration AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-migrate-user.html
	//
	// Experimental.
	UserMigration awslambda.IFunction `json:"userMigration" yaml:"userMigration"`
	// Verifies the authentication challenge response.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-verify-auth-challenge-response.html
	//
	// Experimental.
	VerifyAuthChallengeResponse awslambda.IFunction `json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

// User pool configuration for user self sign up.
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
// Experimental.
type UserVerificationConfig struct {
	// The email body template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	EmailBody *string `json:"emailBody" yaml:"emailBody"`
	// Emails can be verified either using a code or a link.
	//
	// Learn more at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-email-verification-message-customization.html
	// Experimental.
	EmailStyle VerificationEmailStyle `json:"emailStyle" yaml:"emailStyle"`
	// The email subject template for the verification email sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	EmailSubject *string `json:"emailSubject" yaml:"emailSubject"`
	// The message template for the verification SMS sent to the user upon sign up.
	//
	// See https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-message-templates.html to
	// learn more about message templates.
	// Experimental.
	SmsMessage *string `json:"smsMessage" yaml:"smsMessage"`
}

// The email verification style.
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
// Experimental.
type VerificationEmailStyle string

const (
	// Verify email via code.
	// Experimental.
	VerificationEmailStyle_CODE VerificationEmailStyle = "CODE"
	// Verify email via link.
	// Experimental.
	VerificationEmailStyle_LINK VerificationEmailStyle = "LINK"
)

