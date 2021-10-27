// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amplify Console application.
// Experimental.
type App interface {
	awscdk.Resource
	IApp
	awsiam.IGrantable
	AppId() *string
	AppName() *string
	Arn() *string
	DefaultDomain() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddAutoBranchEnvironment(name *string, value *string) App
	AddBranch(id *string, options *BranchOptions) Branch
	AddCustomRule(rule CustomRule) App
	AddDomain(id *string, options *DomainOptions) Domain
	AddEnvironment(name *string, value *string) App
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__awscdkResource
	jsiiProxy_IApp
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_App) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApp(scope constructs.Construct, id *string, props *AppProps) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.App",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApp_Override(a App, scope constructs.Construct, id *string, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.App",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing application.
// Experimental.
func App_FromAppId(scope constructs.Construct, id *string, appId *string) IApp {
	_init_.Initialize()

	var returns IApp

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"fromAppId",
		[]interface{}{scope, id, appId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func App_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an environment variable to the auto created branch.
//
// All environment variables that you add are encrypted to prevent rogue
// access so you can use them to store secret information.
// Experimental.
func (a *jsiiProxy_App) AddAutoBranchEnvironment(name *string, value *string) App {
	var returns App

	_jsii_.Invoke(
		a,
		"addAutoBranchEnvironment",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

// Adds a branch to this application.
// Experimental.
func (a *jsiiProxy_App) AddBranch(id *string, options *BranchOptions) Branch {
	var returns Branch

	_jsii_.Invoke(
		a,
		"addBranch",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a custom rewrite/redirect rule to this application.
// Experimental.
func (a *jsiiProxy_App) AddCustomRule(rule CustomRule) App {
	var returns App

	_jsii_.Invoke(
		a,
		"addCustomRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Adds a domain to this application.
// Experimental.
func (a *jsiiProxy_App) AddDomain(id *string, options *DomainOptions) Domain {
	var returns Domain

	_jsii_.Invoke(
		a,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds an environment variable to this application.
//
// All environment variables that you add are encrypted to prevent rogue
// access so you can use them to store secret information.
// Experimental.
func (a *jsiiProxy_App) AddEnvironment(name *string, value *string) App {
	var returns App

	_jsii_.Invoke(
		a,
		"addEnvironment",
		[]interface{}{name, value},
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
func (a *jsiiProxy_App) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_App) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_App) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_App) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an App.
// Experimental.
type AppProps struct {
	// The name for the application.
	// Experimental.
	AppName *string `json:"appName"`
	// The auto branch creation configuration.
	//
	// Use this to automatically create
	// branches that match a certain pattern.
	// Experimental.
	AutoBranchCreation *AutoBranchCreation `json:"autoBranchCreation"`
	// Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.
	// Experimental.
	AutoBranchDeletion *bool `json:"autoBranchDeletion"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection at an
	// app level to all your branches.
	// Experimental.
	BasicAuth BasicAuth `json:"basicAuth"`
	// BuildSpec for the application.
	//
	// Alternatively, add a `amplify.yml`
	// file to the repository.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// The custom HTTP response headers for an Amplify app.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/custom-headers.html
	//
	// Experimental.
	CustomResponseHeaders *[]*CustomResponseHeader `json:"customResponseHeaders"`
	// Custom rewrite/redirect rules for the application.
	// Experimental.
	CustomRules *[]CustomRule `json:"customRules"`
	// A description for the application.
	// Experimental.
	Description *string `json:"description"`
	// Environment variables for the application.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `json:"environmentVariables"`
	// The IAM service role to associate with the application.
	//
	// The App
	// implements IGrantable.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The source code provider for this application.
	// Experimental.
	SourceCodeProvider ISourceCodeProvider `json:"sourceCodeProvider"`
}

// Auto branch creation configuration.
// Experimental.
type AutoBranchCreation struct {
	// Whether to enable auto building for the auto created branch.
	// Experimental.
	AutoBuild *bool `json:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the auto created branch.
	// Experimental.
	BasicAuth BasicAuth `json:"basicAuth"`
	// Build spec for the auto created branch.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Environment variables for the auto created branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `json:"environmentVariables"`
	// Automated branch creation glob patterns.
	// Experimental.
	Patterns *[]*string `json:"patterns"`
	// The dedicated backend environment for the pull request previews of the auto created branch.
	// Experimental.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the auto created branch.
	// Experimental.
	PullRequestPreview *bool `json:"pullRequestPreview"`
	// Stage for the auto created branch.
	// Experimental.
	Stage *string `json:"stage"`
}

// Basic Auth configuration.
// Experimental.
type BasicAuth interface {
	Bind(scope constructs.Construct, id *string) *BasicAuthConfig
}

// The jsii proxy struct for BasicAuth
type jsiiProxy_BasicAuth struct {
	_ byte // padding
}

// Experimental.
func NewBasicAuth(props *BasicAuthProps) BasicAuth {
	_init_.Initialize()

	j := jsiiProxy_BasicAuth{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBasicAuth_Override(b BasicAuth, props *BasicAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		[]interface{}{props},
		b,
	)
}

// Creates a Basic Auth configuration from a username and a password.
// Experimental.
func BasicAuth_FromCredentials(username *string, password awscdk.SecretValue) BasicAuth {
	_init_.Initialize()

	var returns BasicAuth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		"fromCredentials",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Creates a Basic Auth configuration with a password generated in Secrets Manager.
// Experimental.
func BasicAuth_FromGeneratedPassword(username *string, encryptionKey awskms.IKey) BasicAuth {
	_init_.Initialize()

	var returns BasicAuth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		"fromGeneratedPassword",
		[]interface{}{username, encryptionKey},
		&returns,
	)

	return returns
}

// Binds this Basic Auth configuration to an App.
// Experimental.
func (b *jsiiProxy_BasicAuth) Bind(scope constructs.Construct, id *string) *BasicAuthConfig {
	var returns *BasicAuthConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// A Basic Auth configuration.
// Experimental.
type BasicAuthConfig struct {
	// Whether to enable Basic Auth.
	// Experimental.
	EnableBasicAuth *bool `json:"enableBasicAuth"`
	// The password.
	// Experimental.
	Password *string `json:"password"`
	// The username.
	// Experimental.
	Username *string `json:"username"`
}

// Properties for a BasicAuth.
// Experimental.
type BasicAuthProps struct {
	// The username.
	// Experimental.
	Username *string `json:"username"`
	// The encryption key to use to encrypt the password when it's generated in Secrets Manager.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The password.
	// Experimental.
	Password awscdk.SecretValue `json:"password"`
}

// An Amplify Console branch.
// Experimental.
type Branch interface {
	awscdk.Resource
	IBranch
	Arn() *string
	BranchName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEnvironment(name *string, value *string) Branch
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Branch
type jsiiProxy_Branch struct {
	internal.Type__awscdkResource
	jsiiProxy_IBranch
}

func (j *jsiiProxy_Branch) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Branch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Branch) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Branch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Branch) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Branch) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewBranch(scope constructs.Construct, id *string, props *BranchProps) Branch {
	_init_.Initialize()

	j := jsiiProxy_Branch{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.Branch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBranch_Override(b Branch, scope constructs.Construct, id *string, props *BranchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.Branch",
		[]interface{}{scope, id, props},
		b,
	)
}

// Import an existing branch.
// Experimental.
func Branch_FromBranchName(scope constructs.Construct, id *string, branchName *string) IBranch {
	_init_.Initialize()

	var returns IBranch

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.Branch",
		"fromBranchName",
		[]interface{}{scope, id, branchName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Branch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.Branch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Branch_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.Branch",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an environment variable to this branch.
//
// All environment variables that you add are encrypted to prevent rogue
// access so you can use them to store secret information.
// Experimental.
func (b *jsiiProxy_Branch) AddEnvironment(name *string, value *string) Branch {
	var returns Branch

	_jsii_.Invoke(
		b,
		"addEnvironment",
		[]interface{}{name, value},
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
func (b *jsiiProxy_Branch) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (b *jsiiProxy_Branch) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_Branch) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_Branch) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_Branch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to add a branch to an application.
// Experimental.
type BranchOptions struct {
	// Whether to enable auto building for the branch.
	// Experimental.
	AutoBuild *bool `json:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the branch
	// Experimental.
	BasicAuth BasicAuth `json:"basicAuth"`
	// The name of the branch.
	// Experimental.
	BranchName *string `json:"branchName"`
	// BuildSpec for the branch.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// A description for the branch.
	// Experimental.
	Description *string `json:"description"`
	// Environment variables for the branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `json:"environmentVariables"`
	// The dedicated backend environment for the pull request previews.
	// Experimental.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the branch.
	// Experimental.
	PullRequestPreview *bool `json:"pullRequestPreview"`
	// Stage for the branch.
	// Experimental.
	Stage *string `json:"stage"`
}

// Properties for a Branch.
// Experimental.
type BranchProps struct {
	// Whether to enable auto building for the branch.
	// Experimental.
	AutoBuild *bool `json:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the branch
	// Experimental.
	BasicAuth BasicAuth `json:"basicAuth"`
	// The name of the branch.
	// Experimental.
	BranchName *string `json:"branchName"`
	// BuildSpec for the branch.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// A description for the branch.
	// Experimental.
	Description *string `json:"description"`
	// Environment variables for the branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `json:"environmentVariables"`
	// The dedicated backend environment for the pull request previews.
	// Experimental.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the branch.
	// Experimental.
	PullRequestPreview *bool `json:"pullRequestPreview"`
	// Stage for the branch.
	// Experimental.
	Stage *string `json:"stage"`
	// The application within which the branch must be created.
	// Experimental.
	App IApp `json:"app"`
}

// CodeCommit source code provider.
// Experimental.
type CodeCommitSourceCodeProvider interface {
	ISourceCodeProvider
	Bind(app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for CodeCommitSourceCodeProvider
type jsiiProxy_CodeCommitSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewCodeCommitSourceCodeProvider(props *CodeCommitSourceCodeProviderProps) CodeCommitSourceCodeProvider {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitSourceCodeProvider_Override(c CodeCommitSourceCodeProvider, props *CodeCommitSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProvider",
		[]interface{}{props},
		c,
	)
}

// Binds the source code provider to an app.
// Experimental.
func (c *jsiiProxy_CodeCommitSourceCodeProvider) Bind(app App) *SourceCodeProviderConfig {
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Properties for a CodeCommit source code provider.
// Experimental.
type CodeCommitSourceCodeProviderProps struct {
	// The CodeCommit repository.
	// Experimental.
	Repository awscodecommit.IRepository `json:"repository"`
}

// Custom response header of an Amplify App.
// Experimental.
type CustomResponseHeader struct {
	// The map of custom headers to be applied.
	// Experimental.
	Headers *map[string]*string `json:"headers"`
	// These custom headers will be applied to all URL file paths that match this pattern.
	// Experimental.
	Pattern *string `json:"pattern"`
}

// Custom rewrite/redirect rule for an Amplify App.
// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
//
// Experimental.
type CustomRule interface {
	Condition() *string
	Source() *string
	Status() RedirectStatus
	Target() *string
}

// The jsii proxy struct for CustomRule
type jsiiProxy_CustomRule struct {
	_ byte // padding
}

func (j *jsiiProxy_CustomRule) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Status() RedirectStatus {
	var returns RedirectStatus
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomRule(options *CustomRuleOptions) CustomRule {
	_init_.Initialize()

	j := jsiiProxy_CustomRule{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CustomRule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomRule_Override(c CustomRule, options *CustomRuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CustomRule",
		[]interface{}{options},
		c,
	)
}

func CustomRule_SINGLE_PAGE_APPLICATION_REDIRECT() CustomRule {
	_init_.Initialize()
	var returns CustomRule
	_jsii_.StaticGet(
		"@aws-cdk/aws-amplify-alpha.CustomRule",
		"SINGLE_PAGE_APPLICATION_REDIRECT",
		&returns,
	)
	return returns
}

// Options for a custom rewrite/redirect rule for an Amplify App.
// Experimental.
type CustomRuleOptions struct {
	// The source pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Source *string `json:"source"`
	// The target pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Target *string `json:"target"`
	// The condition for a URL rewrite or redirect rule, e.g. country code.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Condition *string `json:"condition"`
	// The status code for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Status RedirectStatus `json:"status"`
}

// An Amplify Console domain.
// Experimental.
type Domain interface {
	awscdk.Resource
	Arn() *string
	CertificateRecord() *string
	DomainAutoSubDomainCreationPatterns() *[]*string
	DomainAutoSubDomainIamRole() *string
	DomainEnableAutoSubDomain() awscdk.IResolvable
	DomainName() *string
	DomainStatus() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	StatusReason() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MapRoot(branch IBranch) Domain
	MapSubDomain(branch IBranch, prefix *string) Domain
	ToString() *string
}

// The jsii proxy struct for Domain
type jsiiProxy_Domain struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Domain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) CertificateRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainAutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainAutoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainAutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAutoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainEnableAutoSubDomain() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"domainEnableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}


// Experimental.
func NewDomain(scope constructs.Construct, id *string, props *DomainProps) Domain {
	_init_.Initialize()

	j := jsiiProxy_Domain{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.Domain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Domain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.Domain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Domain_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.Domain",
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
func (d *jsiiProxy_Domain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_Domain) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_Domain) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_Domain) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Maps a branch to the domain root.
// Experimental.
func (d *jsiiProxy_Domain) MapRoot(branch IBranch) Domain {
	var returns Domain

	_jsii_.Invoke(
		d,
		"mapRoot",
		[]interface{}{branch},
		&returns,
	)

	return returns
}

// Maps a branch to a sub domain.
// Experimental.
func (d *jsiiProxy_Domain) MapSubDomain(branch IBranch, prefix *string) Domain {
	var returns Domain

	_jsii_.Invoke(
		d,
		"mapSubDomain",
		[]interface{}{branch, prefix},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_Domain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to add a domain to an application.
// Experimental.
type DomainOptions struct {
	// Branches which should automatically create subdomains.
	// Experimental.
	AutoSubdomainCreationPatterns *[]*string `json:"autoSubdomainCreationPatterns"`
	// The name of the domain.
	// Experimental.
	DomainName *string `json:"domainName"`
	// Automatically create subdomains for connected branches.
	// Experimental.
	EnableAutoSubdomain *bool `json:"enableAutoSubdomain"`
	// Subdomains.
	// Experimental.
	SubDomains *[]*SubDomain `json:"subDomains"`
}

// Properties for a Domain.
// Experimental.
type DomainProps struct {
	// Branches which should automatically create subdomains.
	// Experimental.
	AutoSubdomainCreationPatterns *[]*string `json:"autoSubdomainCreationPatterns"`
	// The name of the domain.
	// Experimental.
	DomainName *string `json:"domainName"`
	// Automatically create subdomains for connected branches.
	// Experimental.
	EnableAutoSubdomain *bool `json:"enableAutoSubdomain"`
	// Subdomains.
	// Experimental.
	SubDomains *[]*SubDomain `json:"subDomains"`
	// The application to which the domain must be connected.
	// Experimental.
	App IApp `json:"app"`
	// The IAM role with access to Route53 when using enableAutoSubdomain.
	// Experimental.
	AutoSubDomainIamRole awsiam.IRole `json:"autoSubDomainIamRole"`
}

// GitHub source code provider.
// Experimental.
type GitHubSourceCodeProvider interface {
	ISourceCodeProvider
	Bind(_app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for GitHubSourceCodeProvider
type jsiiProxy_GitHubSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewGitHubSourceCodeProvider(props *GitHubSourceCodeProviderProps) GitHubSourceCodeProvider {
	_init_.Initialize()

	j := jsiiProxy_GitHubSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubSourceCodeProvider_Override(g GitHubSourceCodeProvider, props *GitHubSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProvider",
		[]interface{}{props},
		g,
	)
}

// Binds the source code provider to an app.
// Experimental.
func (g *jsiiProxy_GitHubSourceCodeProvider) Bind(_app App) *SourceCodeProviderConfig {
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_app},
		&returns,
	)

	return returns
}

// Properties for a GitHub source code provider.
// Experimental.
type GitHubSourceCodeProviderProps struct {
	// A personal access token with the `repo` scope.
	// Experimental.
	OauthToken awscdk.SecretValue `json:"oauthToken"`
	// The user or organization owning the repository.
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the repository.
	// Experimental.
	Repository *string `json:"repository"`
}

// GitLab source code provider.
// Experimental.
type GitLabSourceCodeProvider interface {
	ISourceCodeProvider
	Bind(_app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for GitLabSourceCodeProvider
type jsiiProxy_GitLabSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewGitLabSourceCodeProvider(props *GitLabSourceCodeProviderProps) GitLabSourceCodeProvider {
	_init_.Initialize()

	j := jsiiProxy_GitLabSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitLabSourceCodeProvider_Override(g GitLabSourceCodeProvider, props *GitLabSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProvider",
		[]interface{}{props},
		g,
	)
}

// Binds the source code provider to an app.
// Experimental.
func (g *jsiiProxy_GitLabSourceCodeProvider) Bind(_app App) *SourceCodeProviderConfig {
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_app},
		&returns,
	)

	return returns
}

// Properties for a GitLab source code provider.
// Experimental.
type GitLabSourceCodeProviderProps struct {
	// A personal access token with the `repo` scope.
	// Experimental.
	OauthToken awscdk.SecretValue `json:"oauthToken"`
	// The user or organization owning the repository.
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the repository.
	// Experimental.
	Repository *string `json:"repository"`
}

// An Amplify Console application.
// Experimental.
type IApp interface {
	awscdk.IResource
	// The application id.
	// Experimental.
	AppId() *string
}

// The jsii proxy for IApp
type jsiiProxy_IApp struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

// A branch.
// Experimental.
type IBranch interface {
	awscdk.IResource
	// The name of the branch.
	// Experimental.
	BranchName() *string
}

// The jsii proxy for IBranch
type jsiiProxy_IBranch struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

// A source code provider.
// Experimental.
type ISourceCodeProvider interface {
	// Binds the source code provider to an app.
	// Experimental.
	Bind(app App) *SourceCodeProviderConfig
}

// The jsii proxy for ISourceCodeProvider
type jsiiProxy_ISourceCodeProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_ISourceCodeProvider) Bind(app App) *SourceCodeProviderConfig {
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// The status code for a URL rewrite or redirect rule.
// Experimental.
type RedirectStatus string

const (
	RedirectStatus_REWRITE RedirectStatus = "REWRITE"
	RedirectStatus_PERMANENT_REDIRECT RedirectStatus = "PERMANENT_REDIRECT"
	RedirectStatus_TEMPORARY_REDIRECT RedirectStatus = "TEMPORARY_REDIRECT"
	RedirectStatus_NOT_FOUND RedirectStatus = "NOT_FOUND"
	RedirectStatus_NOT_FOUND_REWRITE RedirectStatus = "NOT_FOUND_REWRITE"
)

// Configuration for the source code provider.
// Experimental.
type SourceCodeProviderConfig struct {
	// The repository for the application.
	//
	// Must use the `HTTPS` protocol.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Repository *string `json:"repository"`
	// Personal Access token for 3rd party source control system for an Amplify App, used to create webhook and read-only deploy key.
	//
	// Token is not stored.
	//
	// Either `accessToken` or `oauthToken` must be specified if `repository`
	// is sepcified.
	// Experimental.
	AccessToken awscdk.SecretValue `json:"accessToken"`
	// OAuth token for 3rd party source control system for an Amplify App, used to create webhook and read-only deploy key.
	//
	// OAuth token is not stored.
	//
	// Either `accessToken` or `oauthToken` must be specified if `repository`
	// is sepcified.
	// Experimental.
	OauthToken awscdk.SecretValue `json:"oauthToken"`
}

// Sub domain settings.
// Experimental.
type SubDomain struct {
	// The branch.
	// Experimental.
	Branch IBranch `json:"branch"`
	// The prefix.
	//
	// Use '' to map to the root of the domain
	// Experimental.
	Prefix *string `json:"prefix"`
}

