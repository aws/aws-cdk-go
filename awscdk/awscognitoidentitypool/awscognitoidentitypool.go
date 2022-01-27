package awscognitoidentitypool

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awscognitoidentitypool/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a Cognito IdentityPool.
// Experimental.
type IIdentityPool interface {
	awscdk.IResource
	// The ARN of the Identity Pool.
	// Experimental.
	IdentityPoolArn() *string
	// The id of the Identity Pool in the format REGION:GUID.
	// Experimental.
	IdentityPoolId() *string
	// Name of the Identity Pool.
	// Experimental.
	IdentityPoolName() *string
}

// The jsii proxy for IIdentityPool
type jsiiProxy_IIdentityPool struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

// Represents an Identity Pool Role Attachment.
// Experimental.
type IIdentityPoolRoleAttachment interface {
	awscdk.IResource
	// Id of the Attachments Underlying Identity Pool.
	// Experimental.
	IdentityPoolId() *string
}

// The jsii proxy for IIdentityPoolRoleAttachment
type jsiiProxy_IIdentityPoolRoleAttachment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIdentityPoolRoleAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

// Represents the concept of a User Pool Authentication Provider.
//
// You use user pool authentication providers to configure User Pools
// and User Pool Clients for use with Identity Pools
// Experimental.
type IUserPoolAuthenticationProvider interface {
	// The method called when a given User Pool Authentication Provider is added (for the first time) to an Identity Pool.
	// Experimental.
	Bind(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig
}

// The jsii proxy for IUserPoolAuthenticationProvider
type jsiiProxy_IUserPoolAuthenticationProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IUserPoolAuthenticationProvider) Bind(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig {
	var returns *UserPoolAuthenticationProviderBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, identityPool, options},
		&returns,
	)

	return returns
}

// Define a Cognito Identity Pool.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPool interface {
	awscdk.Resource
	IIdentityPool
	AuthenticatedRole() awsiam.IRole
	Env() *awscdk.ResourceEnvironment
	IdentityPoolArn() *string
	IdentityPoolId() *string
	IdentityPoolName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	UnauthenticatedRole() awsiam.IRole
	AddRoleMappings(roleMappings ...*IdentityPoolRoleMapping)
	AddUserPoolAuthentication(userPool IUserPoolAuthenticationProvider)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for IdentityPool
type jsiiProxy_IdentityPool struct {
	internal.Type__awscdkResource
	jsiiProxy_IIdentityPool
}

func (j *jsiiProxy_IdentityPool) AuthenticatedRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"authenticatedRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) IdentityPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPool) UnauthenticatedRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"unauthenticatedRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewIdentityPool(scope constructs.Construct, id *string, props *IdentityPoolProps) IdentityPool {
	_init_.Initialize()

	j := jsiiProxy_IdentityPool{}

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIdentityPool_Override(i IdentityPool, scope constructs.Construct, id *string, props *IdentityPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		[]interface{}{scope, id, props},
		i,
	)
}

// Import an existing Identity Pool from its Arn.
// Experimental.
func IdentityPool_FromIdentityPoolArn(scope constructs.Construct, id *string, identityPoolArn *string) IIdentityPool {
	_init_.Initialize()

	var returns IIdentityPool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		"fromIdentityPoolArn",
		[]interface{}{scope, id, identityPoolArn},
		&returns,
	)

	return returns
}

// Import an existing Identity Pool from its id.
// Experimental.
func IdentityPool_FromIdentityPoolId(scope constructs.Construct, id *string, identityPoolId *string) IIdentityPool {
	_init_.Initialize()

	var returns IIdentityPool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		"fromIdentityPoolId",
		[]interface{}{scope, id, identityPoolId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func IdentityPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func IdentityPool_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPool",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds Role Mappings to Identity Pool.
// Experimental.
func (i *jsiiProxy_IdentityPool) AddRoleMappings(roleMappings ...*IdentityPoolRoleMapping) {
	args := []interface{}{}
	for _, a := range roleMappings {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addRoleMappings",
		args,
	)
}

// Add a User Pool to the IdentityPool and configure User Pool Client to handle identities.
// Experimental.
func (i *jsiiProxy_IdentityPool) AddUserPoolAuthentication(userPool IUserPoolAuthenticationProvider) {
	_jsii_.InvokeVoid(
		i,
		"addUserPoolAuthentication",
		[]interface{}{userPool},
	)
}

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
func (i *jsiiProxy_IdentityPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (i *jsiiProxy_IdentityPool) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (i *jsiiProxy_IdentityPool) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (i *jsiiProxy_IdentityPool) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (i *jsiiProxy_IdentityPool) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (i *jsiiProxy_IdentityPool) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (i *jsiiProxy_IdentityPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (i *jsiiProxy_IdentityPool) Prepare() {
	_jsii_.InvokeVoid(
		i,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (i *jsiiProxy_IdentityPool) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (i *jsiiProxy_IdentityPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (i *jsiiProxy_IdentityPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Login Provider for Identity Federation using Amazon Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolAmazonLoginProvider struct {
	// App Id for Amazon Identity Federation.
	// Experimental.
	AppId *string `json:"appId" yaml:"appId"`
}

// Login Provider for Identity Federation using Apple Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolAppleLoginProvider struct {
	// App Id for Apple Identity Federation.
	// Experimental.
	ServicesId *string `json:"servicesId" yaml:"servicesId"`
}

// Authentication providers for using in identity pool.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/external-identity-providers.html
//
// Experimental.
type IdentityPoolAuthenticationProviders struct {
	// App Id for Amazon Identity Federation.
	// Experimental.
	Amazon *IdentityPoolAmazonLoginProvider `json:"amazon" yaml:"amazon"`
	// Services Id for Apple Identity Federation.
	// Experimental.
	Apple *IdentityPoolAppleLoginProvider `json:"apple" yaml:"apple"`
	// Consumer Key and Secret for Digits Identity Federation.
	// Experimental.
	Digits *IdentityPoolDigitsLoginProvider `json:"digits" yaml:"digits"`
	// App Id for Facebook Identity Federation.
	// Experimental.
	Facebook *IdentityPoolFacebookLoginProvider `json:"facebook" yaml:"facebook"`
	// Client Id for Google Identity Federation.
	// Experimental.
	Google *IdentityPoolGoogleLoginProvider `json:"google" yaml:"google"`
	// Consumer Key and Secret for Twitter Identity Federation.
	// Experimental.
	Twitter *IdentityPoolTwitterLoginProvider `json:"twitter" yaml:"twitter"`
	// The Developer Provider Name to associate with this Identity Pool.
	// Experimental.
	CustomProvider *string `json:"customProvider" yaml:"customProvider"`
	// The OpenIdConnect Provider associated with this Identity Pool.
	// Experimental.
	OpenIdConnectProviders *[]awsiam.IOpenIdConnectProvider `json:"openIdConnectProviders" yaml:"openIdConnectProviders"`
	// The Security Assertion Markup Language Provider associated with this Identity Pool.
	// Experimental.
	SamlProviders *[]awsiam.ISamlProvider `json:"samlProviders" yaml:"samlProviders"`
	// The User Pool Authentication Providers associated with this Identity Pool.
	// Experimental.
	UserPools *[]IUserPoolAuthenticationProvider `json:"userPools" yaml:"userPools"`
}

// Login Provider for Identity Federation using Digits Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolDigitsLoginProvider struct {
	// App Id for Twitter Identity Federation.
	// Experimental.
	ConsumerKey *string `json:"consumerKey" yaml:"consumerKey"`
	// App Secret for Twitter Identity Federation.
	// Experimental.
	ConsumerSecret *string `json:"consumerSecret" yaml:"consumerSecret"`
}

// Login Provider for Identity Federation using Facebook Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolFacebookLoginProvider struct {
	// App Id for Facebook Identity Federation.
	// Experimental.
	AppId *string `json:"appId" yaml:"appId"`
}

// Login Provider for Identity Federation using Google Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolGoogleLoginProvider struct {
	// App Id for Google Identity Federation.
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
}

// Props for the IdentityPool construct.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolProps struct {
	// Enables the Basic (Classic) authentication flow.
	// Experimental.
	AllowClassicFlow *bool `json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Wwhether the identity pool supports unauthenticated logins.
	// Experimental.
	AllowUnauthenticatedIdentities *bool `json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// The Default Role to be assumed by Authenticated Users.
	// Experimental.
	AuthenticatedRole awsiam.IRole `json:"authenticatedRole" yaml:"authenticatedRole"`
	// Authentication providers for using in identity pool.
	// Experimental.
	AuthenticationProviders *IdentityPoolAuthenticationProviders `json:"authenticationProviders" yaml:"authenticationProviders"`
	// The name of the Identity Pool.
	// Experimental.
	IdentityPoolName *string `json:"identityPoolName" yaml:"identityPoolName"`
	// Rules for mapping roles to users.
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `json:"roleMappings" yaml:"roleMappings"`
	// The Default Role to be assumed by Unauthenticated Users.
	// Experimental.
	UnauthenticatedRole awsiam.IRole `json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

// Types of Identity Pool Login Providers.
// Experimental.
type IdentityPoolProviderType string

const (
	IdentityPoolProviderType_FACEBOOK IdentityPoolProviderType = "FACEBOOK"
	IdentityPoolProviderType_GOOGLE IdentityPoolProviderType = "GOOGLE"
	IdentityPoolProviderType_AMAZON IdentityPoolProviderType = "AMAZON"
	IdentityPoolProviderType_APPLE IdentityPoolProviderType = "APPLE"
	IdentityPoolProviderType_TWITTER IdentityPoolProviderType = "TWITTER"
	IdentityPoolProviderType_DIGITS IdentityPoolProviderType = "DIGITS"
	IdentityPoolProviderType_OPEN_ID IdentityPoolProviderType = "OPEN_ID"
	IdentityPoolProviderType_SAML IdentityPoolProviderType = "SAML"
	IdentityPoolProviderType_USER_POOL IdentityPoolProviderType = "USER_POOL"
	IdentityPoolProviderType_CUSTOM IdentityPoolProviderType = "CUSTOM"
)

// Keys for Login Providers - correspond to client id's of respective federation identity providers.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolProviderUrl interface {
	Type() IdentityPoolProviderType
	Value() *string
}

// The jsii proxy struct for IdentityPoolProviderUrl
type jsiiProxy_IdentityPoolProviderUrl struct {
	_ byte // padding
}

func (j *jsiiProxy_IdentityPoolProviderUrl) Type() IdentityPoolProviderType {
	var returns IdentityPoolProviderType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolProviderUrl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewIdentityPoolProviderUrl(type_ IdentityPoolProviderType, value *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	j := jsiiProxy_IdentityPoolProviderUrl{}

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		[]interface{}{type_, value},
		&j,
	)

	return &j
}

// Experimental.
func NewIdentityPoolProviderUrl_Override(i IdentityPoolProviderUrl, type_ IdentityPoolProviderType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		[]interface{}{type_, value},
		i,
	)
}

// Custom Provider Url.
// Experimental.
func IdentityPoolProviderUrl_Custom(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"custom",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// OpenId Provider Url.
// Experimental.
func IdentityPoolProviderUrl_OpenId(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"openId",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// Saml Provider Url.
// Experimental.
func IdentityPoolProviderUrl_Saml(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"saml",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// User Pool Provider Url.
// Experimental.
func IdentityPoolProviderUrl_UserPool(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"userPool",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func IdentityPoolProviderUrl_AMAZON() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"AMAZON",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_APPLE() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"APPLE",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_DIGITS() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"DIGITS",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_FACEBOOK() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"FACEBOOK",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_GOOGLE() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"GOOGLE",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_TWITTER() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
		"TWITTER",
		&returns,
	)
	return returns
}

// External Identity Providers To Connect to User Pools and Identity Pools.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolProviders struct {
	// App Id for Amazon Identity Federation.
	// Experimental.
	Amazon *IdentityPoolAmazonLoginProvider `json:"amazon" yaml:"amazon"`
	// Services Id for Apple Identity Federation.
	// Experimental.
	Apple *IdentityPoolAppleLoginProvider `json:"apple" yaml:"apple"`
	// Consumer Key and Secret for Digits Identity Federation.
	// Experimental.
	Digits *IdentityPoolDigitsLoginProvider `json:"digits" yaml:"digits"`
	// App Id for Facebook Identity Federation.
	// Experimental.
	Facebook *IdentityPoolFacebookLoginProvider `json:"facebook" yaml:"facebook"`
	// Client Id for Google Identity Federation.
	// Experimental.
	Google *IdentityPoolGoogleLoginProvider `json:"google" yaml:"google"`
	// Consumer Key and Secret for Twitter Identity Federation.
	// Experimental.
	Twitter *IdentityPoolTwitterLoginProvider `json:"twitter" yaml:"twitter"`
}

// Defines an Identity Pool Role Attachment.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolRoleAttachment interface {
	awscdk.Resource
	IIdentityPoolRoleAttachment
	Env() *awscdk.ResourceEnvironment
	IdentityPoolId() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for IdentityPoolRoleAttachment
type jsiiProxy_IdentityPoolRoleAttachment struct {
	internal.Type__awscdkResource
	jsiiProxy_IIdentityPoolRoleAttachment
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewIdentityPoolRoleAttachment(scope constructs.Construct, id *string, props *IdentityPoolRoleAttachmentProps) IdentityPoolRoleAttachment {
	_init_.Initialize()

	j := jsiiProxy_IdentityPoolRoleAttachment{}

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIdentityPoolRoleAttachment_Override(i IdentityPoolRoleAttachment, scope constructs.Construct, id *string, props *IdentityPoolRoleAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		i,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func IdentityPoolRoleAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func IdentityPoolRoleAttachment_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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
func (i *jsiiProxy_IdentityPoolRoleAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) Prepare() {
	_jsii_.InvokeVoid(
		i,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (i *jsiiProxy_IdentityPoolRoleAttachment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Props for an Identity Pool Role Attachment.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolRoleAttachmentProps struct {
	// Id of the Attachments Underlying Identity Pool.
	// Experimental.
	IdentityPool IIdentityPool `json:"identityPool" yaml:"identityPool"`
	// Default Authenticated (User) Role.
	// Experimental.
	AuthenticatedRole awsiam.IRole `json:"authenticatedRole" yaml:"authenticatedRole"`
	// Rules for mapping roles to users.
	// Experimental.
	RoleMappings *[]*IdentityPoolRoleMapping `json:"roleMappings" yaml:"roleMappings"`
	// Default Unauthenticated (Guest) Role.
	// Experimental.
	UnauthenticatedRole awsiam.IRole `json:"unauthenticatedRole" yaml:"unauthenticatedRole"`
}

// Map roles to users in the identity pool based on claims from the Identity Provider.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
//
// Experimental.
type IdentityPoolRoleMapping struct {
	// The url of the provider of for which the role is mapped.
	// Experimental.
	ProviderUrl IdentityPoolProviderUrl `json:"providerUrl" yaml:"providerUrl"`
	// Allow for role assumption when results of role mapping are ambiguous.
	// Experimental.
	ResolveAmbiguousRoles *bool `json:"resolveAmbiguousRoles" yaml:"resolveAmbiguousRoles"`
	// The claim and value that must be matched in order to assume the role.
	//
	// Required if useToken is false
	// Experimental.
	Rules *[]*RoleMappingRule `json:"rules" yaml:"rules"`
	// If true then mapped roles must be passed through the cognito:roles or cognito:preferred_role claims from identity provider.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/role-based-access-control.html#using-tokens-to-assign-roles-to-users
	//
	// Experimental.
	UseToken *bool `json:"useToken" yaml:"useToken"`
}

// Login Provider for Identity Federation using Twitter Credentials.
//
// TODO: EXAMPLE
//
// Experimental.
type IdentityPoolTwitterLoginProvider struct {
	// App Id for Twitter Identity Federation.
	// Experimental.
	ConsumerKey *string `json:"consumerKey" yaml:"consumerKey"`
	// App Secret for Twitter Identity Federation.
	// Experimental.
	ConsumerSecret *string `json:"consumerSecret" yaml:"consumerSecret"`
}

// Types of matches allowed for Role Mapping.
// Experimental.
type RoleMappingMatchType string

const (
	RoleMappingMatchType_EQUALS RoleMappingMatchType = "EQUALS"
	RoleMappingMatchType_CONTAINS RoleMappingMatchType = "CONTAINS"
	RoleMappingMatchType_STARTS_WITH RoleMappingMatchType = "STARTS_WITH"
	RoleMappingMatchType_NOTEQUAL RoleMappingMatchType = "NOTEQUAL"
)

// Represents an Identity Pool Role Attachment Role Mapping Rule.
//
// TODO: EXAMPLE
//
// Experimental.
type RoleMappingRule struct {
	// The key sent in the token by the federated identity provider.
	// Experimental.
	Claim *string `json:"claim" yaml:"claim"`
	// The value of the claim that must be matched.
	// Experimental.
	ClaimValue *string `json:"claimValue" yaml:"claimValue"`
	// The Role to be assumed when Claim Value is matched.
	// Experimental.
	MappedRole awsiam.IRole `json:"mappedRole" yaml:"mappedRole"`
	// How to match with the Claim value.
	// Experimental.
	MatchType RoleMappingMatchType `json:"matchType" yaml:"matchType"`
}

// Defines a User Pool Authentication Provider.
//
// TODO: EXAMPLE
//
// Experimental.
type UserPoolAuthenticationProvider interface {
	IUserPoolAuthenticationProvider
	Bind(scope constructs.Construct, identityPool IIdentityPool, _options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig
}

// The jsii proxy struct for UserPoolAuthenticationProvider
type jsiiProxy_UserPoolAuthenticationProvider struct {
	jsiiProxy_IUserPoolAuthenticationProvider
}

// Experimental.
func NewUserPoolAuthenticationProvider(props *UserPoolAuthenticationProviderProps) UserPoolAuthenticationProvider {
	_init_.Initialize()

	j := jsiiProxy_UserPoolAuthenticationProvider{}

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolAuthenticationProvider_Override(u UserPoolAuthenticationProvider, props *UserPoolAuthenticationProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProvider",
		[]interface{}{props},
		u,
	)
}

// The method called when a given User Pool Authentication Provider is added (for the first time) to an Identity Pool.
// Experimental.
func (u *jsiiProxy_UserPoolAuthenticationProvider) Bind(scope constructs.Construct, identityPool IIdentityPool, _options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig {
	var returns *UserPoolAuthenticationProviderBindConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{scope, identityPool, _options},
		&returns,
	)

	return returns
}

// Represents a UserPoolAuthenticationProvider Bind Configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type UserPoolAuthenticationProviderBindConfig struct {
	// Client Id of the Associated User Pool Client.
	// Experimental.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The identity providers associated with the UserPool.
	// Experimental.
	ProviderName *string `json:"providerName" yaml:"providerName"`
	// Whether to enable the identity pool's server side token check.
	// Experimental.
	ServerSideTokenCheck *bool `json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

// Represents UserPoolAuthenticationProvider Bind Options.
//
// TODO: EXAMPLE
//
// Experimental.
type UserPoolAuthenticationProviderBindOptions struct {
}

// Props for the User Pool Authentication Provider.
//
// TODO: EXAMPLE
//
// Experimental.
type UserPoolAuthenticationProviderProps struct {
	// The User Pool of the Associated Identity Providers.
	// Experimental.
	UserPool awscognito.IUserPool `json:"userPool" yaml:"userPool"`
	// Setting this to true turns off identity pool checks for this user pool to make sure the user has not been globally signed out or deleted before the identity pool provides an OIDC token or AWS credentials for the user.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
	//
	// Experimental.
	DisableServerSideTokenCheck *bool `json:"disableServerSideTokenCheck" yaml:"disableServerSideTokenCheck"`
	// The User Pool Client for the provided User Pool.
	// Experimental.
	UserPoolClient awscognito.IUserPoolClient `json:"userPoolClient" yaml:"userPoolClient"`
}

