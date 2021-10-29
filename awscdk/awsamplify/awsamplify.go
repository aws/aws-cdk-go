package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsamplify/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/constructs-go/constructs/v3"
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_App) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_amplify.App",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApp_Override(a App, scope constructs.Construct, id *string, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.App",
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
		"monocdk.aws_amplify.App",
		"fromAppId",
		[]interface{}{scope, id, appId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func App_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.App",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_App) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_App) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_App) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_App) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_App) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_App) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
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
	Bind(scope awscdk.Construct, id *string) *BasicAuthConfig
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
		"monocdk.aws_amplify.BasicAuth",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBasicAuth_Override(b BasicAuth, props *BasicAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.BasicAuth",
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
		"monocdk.aws_amplify.BasicAuth",
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
		"monocdk.aws_amplify.BasicAuth",
		"fromGeneratedPassword",
		[]interface{}{username, encryptionKey},
		&returns,
	)

	return returns
}

// Binds this Basic Auth configuration to an App.
// Experimental.
func (b *jsiiProxy_BasicAuth) Bind(scope awscdk.Construct, id *string) *BasicAuthConfig {
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
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEnvironment(name *string, value *string) Branch
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

func (j *jsiiProxy_Branch) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_amplify.Branch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBranch_Override(b Branch, scope constructs.Construct, id *string, props *BranchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.Branch",
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
		"monocdk.aws_amplify.Branch",
		"fromBranchName",
		[]interface{}{scope, id, branchName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Branch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.Branch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Branch_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.Branch",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (b *jsiiProxy_Branch) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_Branch) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_Branch) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_Branch) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_Branch) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (b *jsiiProxy_Branch) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
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

// A CloudFormation `AWS::Amplify::App`.
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessToken() *string
	SetAccessToken(val *string)
	AttrAppId() *string
	AttrAppName() *string
	AttrArn() *string
	AttrDefaultDomain() *string
	AutoBranchCreationConfig() interface{}
	SetAutoBranchCreationConfig(val interface{})
	BasicAuthConfig() interface{}
	SetBasicAuthConfig(val interface{})
	BuildSpec() *string
	SetBuildSpec(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CustomHeaders() *string
	SetCustomHeaders(val *string)
	CustomRules() interface{}
	SetCustomRules(val interface{})
	Description() *string
	SetDescription(val *string)
	EnableBranchAutoDeletion() interface{}
	SetEnableBranchAutoDeletion(val interface{})
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	IamServiceRole() *string
	SetIamServiceRole(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	OauthToken() *string
	SetOauthToken(val *string)
	Ref() *string
	Repository() *string
	SetRepository(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrDefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDefaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AutoBranchCreationConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBranchCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) BasicAuthConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CustomHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CustomRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) EnableBranchAutoDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) IamServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Amplify::App`.
func NewCfnApp(scope awscdk.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"monocdk.aws_amplify.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::App`.
func NewCfnApp_Override(c CfnApp, scope awscdk.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp) SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetAutoBranchCreationConfig(val interface{}) {
	_jsii_.Set(
		j,
		"autoBranchCreationConfig",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetBasicAuthConfig(val interface{}) {
	_jsii_.Set(
		j,
		"basicAuthConfig",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetBuildSpec(val *string) {
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetCustomHeaders(val *string) {
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetCustomRules(val interface{}) {
	_jsii_.Set(
		j,
		"customRules",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetEnableBranchAutoDeletion(val interface{}) {
	_jsii_.Set(
		j,
		"enableBranchAutoDeletion",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetIamServiceRole(val *string) {
	_jsii_.Set(
		j,
		"iamServiceRole",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetOauthToken(val *string) {
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
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
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
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
func (c *jsiiProxy_CfnApp) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApp) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_CfnApp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
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
func (c *jsiiProxy_CfnApp) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApp) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CfnApp) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApp_AutoBranchCreationConfigProperty struct {
	// `CfnApp.AutoBranchCreationConfigProperty.AutoBranchCreationPatterns`.
	AutoBranchCreationPatterns *[]*string `json:"autoBranchCreationPatterns"`
	// `CfnApp.AutoBranchCreationConfigProperty.BasicAuthConfig`.
	BasicAuthConfig interface{} `json:"basicAuthConfig"`
	// `CfnApp.AutoBranchCreationConfigProperty.BuildSpec`.
	BuildSpec *string `json:"buildSpec"`
	// `CfnApp.AutoBranchCreationConfigProperty.EnableAutoBranchCreation`.
	EnableAutoBranchCreation interface{} `json:"enableAutoBranchCreation"`
	// `CfnApp.AutoBranchCreationConfigProperty.EnableAutoBuild`.
	EnableAutoBuild interface{} `json:"enableAutoBuild"`
	// `CfnApp.AutoBranchCreationConfigProperty.EnablePerformanceMode`.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode"`
	// `CfnApp.AutoBranchCreationConfigProperty.EnablePullRequestPreview`.
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview"`
	// `CfnApp.AutoBranchCreationConfigProperty.EnvironmentVariables`.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// `CfnApp.AutoBranchCreationConfigProperty.PullRequestEnvironmentName`.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// `CfnApp.AutoBranchCreationConfigProperty.Stage`.
	Stage *string `json:"stage"`
}

type CfnApp_BasicAuthConfigProperty struct {
	// `CfnApp.BasicAuthConfigProperty.EnableBasicAuth`.
	EnableBasicAuth interface{} `json:"enableBasicAuth"`
	// `CfnApp.BasicAuthConfigProperty.Password`.
	Password *string `json:"password"`
	// `CfnApp.BasicAuthConfigProperty.Username`.
	Username *string `json:"username"`
}

type CfnApp_CustomRuleProperty struct {
	// `CfnApp.CustomRuleProperty.Source`.
	Source *string `json:"source"`
	// `CfnApp.CustomRuleProperty.Target`.
	Target *string `json:"target"`
	// `CfnApp.CustomRuleProperty.Condition`.
	Condition *string `json:"condition"`
	// `CfnApp.CustomRuleProperty.Status`.
	Status *string `json:"status"`
}

type CfnApp_EnvironmentVariableProperty struct {
	// `CfnApp.EnvironmentVariableProperty.Name`.
	Name *string `json:"name"`
	// `CfnApp.EnvironmentVariableProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::Amplify::App`.
type CfnAppProps struct {
	// `AWS::Amplify::App.Name`.
	Name *string `json:"name"`
	// `AWS::Amplify::App.AccessToken`.
	AccessToken *string `json:"accessToken"`
	// `AWS::Amplify::App.AutoBranchCreationConfig`.
	AutoBranchCreationConfig interface{} `json:"autoBranchCreationConfig"`
	// `AWS::Amplify::App.BasicAuthConfig`.
	BasicAuthConfig interface{} `json:"basicAuthConfig"`
	// `AWS::Amplify::App.BuildSpec`.
	BuildSpec *string `json:"buildSpec"`
	// `AWS::Amplify::App.CustomHeaders`.
	CustomHeaders *string `json:"customHeaders"`
	// `AWS::Amplify::App.CustomRules`.
	CustomRules interface{} `json:"customRules"`
	// `AWS::Amplify::App.Description`.
	Description *string `json:"description"`
	// `AWS::Amplify::App.EnableBranchAutoDeletion`.
	EnableBranchAutoDeletion interface{} `json:"enableBranchAutoDeletion"`
	// `AWS::Amplify::App.EnvironmentVariables`.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// `AWS::Amplify::App.IAMServiceRole`.
	IamServiceRole *string `json:"iamServiceRole"`
	// `AWS::Amplify::App.OauthToken`.
	OauthToken *string `json:"oauthToken"`
	// `AWS::Amplify::App.Repository`.
	Repository *string `json:"repository"`
	// `AWS::Amplify::App.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Amplify::Branch`.
type CfnBranch interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AppId() *string
	SetAppId(val *string)
	AttrArn() *string
	AttrBranchName() *string
	BasicAuthConfig() interface{}
	SetBasicAuthConfig(val interface{})
	BranchName() *string
	SetBranchName(val *string)
	BuildSpec() *string
	SetBuildSpec(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EnableAutoBuild() interface{}
	SetEnableAutoBuild(val interface{})
	EnablePerformanceMode() interface{}
	SetEnablePerformanceMode(val interface{})
	EnablePullRequestPreview() interface{}
	SetEnablePullRequestPreview(val interface{})
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	PullRequestEnvironmentName() *string
	SetPullRequestEnvironmentName(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Stage() *string
	SetStage(val *string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBranch
type jsiiProxy_CfnBranch struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBranch) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) AttrBranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBranchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) BasicAuthConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) EnableAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) EnablePerformanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) EnablePullRequestPreview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranch) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Amplify::Branch`.
func NewCfnBranch(scope awscdk.Construct, id *string, props *CfnBranchProps) CfnBranch {
	_init_.Initialize()

	j := jsiiProxy_CfnBranch{}

	_jsii_.Create(
		"monocdk.aws_amplify.CfnBranch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::Branch`.
func NewCfnBranch_Override(c CfnBranch, scope awscdk.Construct, id *string, props *CfnBranchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CfnBranch",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBranch) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetBasicAuthConfig(val interface{}) {
	_jsii_.Set(
		j,
		"basicAuthConfig",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetBranchName(val *string) {
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetBuildSpec(val *string) {
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetEnableAutoBuild(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoBuild",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetEnablePerformanceMode(val interface{}) {
	_jsii_.Set(
		j,
		"enablePerformanceMode",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetEnablePullRequestPreview(val interface{}) {
	_jsii_.Set(
		j,
		"enablePullRequestPreview",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetPullRequestEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"pullRequestEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_CfnBranch) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
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
func CfnBranch_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnBranch",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBranch_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnBranch",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBranch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnBranch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBranch_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CfnBranch",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBranch) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnBranch) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnBranch) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnBranch) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBranch) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnBranch) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnBranch) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnBranch) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnBranch) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnBranch) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
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
func (c *jsiiProxy_CfnBranch) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnBranch) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_CfnBranch) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnBranch) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
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
func (c *jsiiProxy_CfnBranch) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBranch) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnBranch) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnBranch) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnBranch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CfnBranch) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnBranch) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBranch_BasicAuthConfigProperty struct {
	// `CfnBranch.BasicAuthConfigProperty.Password`.
	Password *string `json:"password"`
	// `CfnBranch.BasicAuthConfigProperty.Username`.
	Username *string `json:"username"`
	// `CfnBranch.BasicAuthConfigProperty.EnableBasicAuth`.
	EnableBasicAuth interface{} `json:"enableBasicAuth"`
}

type CfnBranch_EnvironmentVariableProperty struct {
	// `CfnBranch.EnvironmentVariableProperty.Name`.
	Name *string `json:"name"`
	// `CfnBranch.EnvironmentVariableProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::Amplify::Branch`.
type CfnBranchProps struct {
	// `AWS::Amplify::Branch.AppId`.
	AppId *string `json:"appId"`
	// `AWS::Amplify::Branch.BranchName`.
	BranchName *string `json:"branchName"`
	// `AWS::Amplify::Branch.BasicAuthConfig`.
	BasicAuthConfig interface{} `json:"basicAuthConfig"`
	// `AWS::Amplify::Branch.BuildSpec`.
	BuildSpec *string `json:"buildSpec"`
	// `AWS::Amplify::Branch.Description`.
	Description *string `json:"description"`
	// `AWS::Amplify::Branch.EnableAutoBuild`.
	EnableAutoBuild interface{} `json:"enableAutoBuild"`
	// `AWS::Amplify::Branch.EnablePerformanceMode`.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode"`
	// `AWS::Amplify::Branch.EnablePullRequestPreview`.
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview"`
	// `AWS::Amplify::Branch.EnvironmentVariables`.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// `AWS::Amplify::Branch.PullRequestEnvironmentName`.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// `AWS::Amplify::Branch.Stage`.
	Stage *string `json:"stage"`
	// `AWS::Amplify::Branch.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Amplify::Domain`.
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AppId() *string
	SetAppId(val *string)
	AttrArn() *string
	AttrAutoSubDomainCreationPatterns() *[]*string
	AttrAutoSubDomainIamRole() *string
	AttrCertificateRecord() *string
	AttrDomainName() *string
	AttrDomainStatus() *string
	AttrEnableAutoSubDomain() awscdk.IResolvable
	AttrStatusReason() *string
	AutoSubDomainCreationPatterns() *[]*string
	SetAutoSubDomainCreationPatterns(val *[]*string)
	AutoSubDomainIamRole() *string
	SetAutoSubDomainIamRole(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	EnableAutoSubDomain() interface{}
	SetEnableAutoSubDomain(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	SubDomainSettings() interface{}
	SetSubDomainSettings(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomain) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrAutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAutoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrAutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAutoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrCertificateRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrEnableAutoSubDomain() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEnableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EnableAutoSubDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SubDomainSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Amplify::Domain`.
func NewCfnDomain(scope awscdk.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"monocdk.aws_amplify.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope awscdk.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAutoSubDomainCreationPatterns(val *[]*string) {
	_jsii_.Set(
		j,
		"autoSubDomainCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAutoSubDomainIamRole(val *string) {
	_jsii_.Set(
		j,
		"autoSubDomainIamRole",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetEnableAutoSubDomain(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoSubDomain",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetSubDomainSettings(val interface{}) {
	_jsii_.Set(
		j,
		"subDomainSettings",
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
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
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
func (c *jsiiProxy_CfnDomain) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDomain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_CfnDomain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
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
func (c *jsiiProxy_CfnDomain) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDomain) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_CfnDomain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDomain_SubDomainSettingProperty struct {
	// `CfnDomain.SubDomainSettingProperty.BranchName`.
	BranchName *string `json:"branchName"`
	// `CfnDomain.SubDomainSettingProperty.Prefix`.
	Prefix *string `json:"prefix"`
}

// Properties for defining a `AWS::Amplify::Domain`.
type CfnDomainProps struct {
	// `AWS::Amplify::Domain.AppId`.
	AppId *string `json:"appId"`
	// `AWS::Amplify::Domain.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::Amplify::Domain.SubDomainSettings`.
	SubDomainSettings interface{} `json:"subDomainSettings"`
	// `AWS::Amplify::Domain.AutoSubDomainCreationPatterns`.
	AutoSubDomainCreationPatterns *[]*string `json:"autoSubDomainCreationPatterns"`
	// `AWS::Amplify::Domain.AutoSubDomainIAMRole`.
	AutoSubDomainIamRole *string `json:"autoSubDomainIamRole"`
	// `AWS::Amplify::Domain.EnableAutoSubDomain`.
	EnableAutoSubDomain interface{} `json:"enableAutoSubDomain"`
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
		"monocdk.aws_amplify.CodeCommitSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitSourceCodeProvider_Override(c CodeCommitSourceCodeProvider, props *CodeCommitSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CodeCommitSourceCodeProvider",
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
		"monocdk.aws_amplify.CustomRule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomRule_Override(c CustomRule, options *CustomRuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CustomRule",
		[]interface{}{options},
		c,
	)
}

func CustomRule_SINGLE_PAGE_APPLICATION_REDIRECT() CustomRule {
	_init_.Initialize()
	var returns CustomRule
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CustomRule",
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
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	StatusReason() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MapRoot(branch IBranch) Domain
	MapSubDomain(branch IBranch, prefix *string) Domain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_Domain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_amplify.Domain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Domain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.Domain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Domain_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.Domain",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_Domain) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_Domain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
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
func (d *jsiiProxy_Domain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_Domain) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_Domain) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
// Experimental.
func (d *jsiiProxy_Domain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
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
		"monocdk.aws_amplify.GitHubSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubSourceCodeProvider_Override(g GitHubSourceCodeProvider, props *GitHubSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.GitHubSourceCodeProvider",
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
		"monocdk.aws_amplify.GitLabSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitLabSourceCodeProvider_Override(g GitLabSourceCodeProvider, props *GitLabSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.GitLabSourceCodeProvider",
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

