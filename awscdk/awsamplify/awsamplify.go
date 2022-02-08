package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Amplify::App`.
//
// The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnApp) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnApp(scope constructs.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::App`.
func NewCfnApp_Override(c CfnApp, scope constructs.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnApp",
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
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnApp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnApp",
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
		"aws-cdk-lib.aws_amplify.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Use the AutoBranchCreationConfig property type to automatically create branches that match a certain pattern.
//
// TODO: EXAMPLE
//
type CfnApp_AutoBranchCreationConfigProperty struct {
	// Automated branch creation glob patterns for the Amplify app.
	AutoBranchCreationPatterns *[]*string `json:"autoBranchCreationPatterns" yaml:"autoBranchCreationPatterns"`
	// Sets password protection for your auto created branch.
	BasicAuthConfig interface{} `json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for the autocreated branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	BuildSpec *string `json:"buildSpec" yaml:"buildSpec"`
	// Enables automated branch creation for the Amplify app.
	EnableAutoBranchCreation interface{} `json:"enableAutoBranchCreation" yaml:"enableAutoBranchCreation"`
	// Enables auto building for the auto created branch.
	EnableAutoBuild interface{} `json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Sets whether pull request previews are enabled for each branch that Amplify Console automatically creates for your app.
	//
	// Amplify Console creates previews by deploying your app to a unique URL whenever a pull request is opened for the branch. Development and QA teams can use this preview to test the pull request before it's merged into a production or integration branch.
	//
	// To provide backend support for your preview, the Amplify Console automatically provisions a temporary backend environment that it deletes when the pull request is closed. If you want to specify a dedicated backend environment for your previews, use the `PullRequestEnvironmentName` property.
	//
	// For more information, see [Web Previews](https://docs.aws.amazon.com/amplify/latest/userguide/pr-previews.html) in the *AWS Amplify Console User Guide* .
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// Environment variables for the auto created branch.
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// If pull request previews are enabled, you can use this property to specify a dedicated backend environment for your previews.
	//
	// For example, you could specify an environment named `prod` , `test` , or `dev` that you initialized with the Amplify CLI.
	//
	// To enable pull request previews, set the `EnablePullRequestPreview` property to `true` .
	//
	// If you don't specify an environment, the Amplify Console provides backend support for each preview by automatically provisioning a temporary backend environment. Amplify Console deletes this environment when the pull request is closed.
	//
	// For more information about creating backend environments, see [Feature Branch Deployments and Team Workflows](https://docs.aws.amazon.com/amplify/latest/userguide/multi-environments.html) in the *AWS Amplify Console User Guide* .
	//
	// *Length Constraints:* Maximum length of 20.
	//
	// *Pattern:* (?s).*
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Stage for the auto created branch.
	Stage *string `json:"stage" yaml:"stage"`
}

// Use the BasicAuthConfig property type to set password protection at an app level to all your branches.
//
// TODO: EXAMPLE
//
type CfnApp_BasicAuthConfigProperty struct {
	// Enables basic authorization for the Amplify app's branches.
	EnableBasicAuth interface{} `json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// The password for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Password *string `json:"password" yaml:"password"`
	// The user name for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Username *string `json:"username" yaml:"username"`
}

// The CustomRule property type allows you to specify redirects, rewrites, and reverse proxies.
//
// Redirects enable a web app to reroute navigation from one URL to another.
//
// TODO: EXAMPLE
//
type CfnApp_CustomRuleProperty struct {
	// The source pattern for a URL rewrite or redirect rule.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 2048.
	//
	// *Pattern:* (?s).+
	Source *string `json:"source" yaml:"source"`
	// The target pattern for a URL rewrite or redirect rule.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 2048.
	//
	// *Pattern:* (?s).+
	Target *string `json:"target" yaml:"target"`
	// The condition for a URL rewrite or redirect rule, such as a country code.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 2048.
	//
	// *Pattern:* (?s).*
	Condition *string `json:"condition" yaml:"condition"`
	// The status code for a URL rewrite or redirect rule.
	//
	// - **200** - Represents a 200 rewrite rule.
	// - **301** - Represents a 301 (moved pemanently) redirect rule. This and all future requests should be directed to the target URL.
	// - **302** - Represents a 302 temporary redirect rule.
	// - **404** - Represents a 404 redirect rule.
	// - **404-200** - Represents a 404 rewrite rule.
	//
	// *Length Constraints:* Minimum length of 3. Maximum length of 7.
	//
	// *Pattern:* .{3,7}
	Status *string `json:"status" yaml:"status"`
}

// Environment variables are key-value pairs that are available at build time.
//
// Set environment variables for all branches in your app.
//
// TODO: EXAMPLE
//
type CfnApp_EnvironmentVariableProperty struct {
	// The environment variable name.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* (?s).*
	Name *string `json:"name" yaml:"name"`
	// The environment variable value.
	//
	// *Length Constraints:* Maximum length of 5500.
	//
	// *Pattern:* (?s).*
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnApp`.
//
// TODO: EXAMPLE
//
type CfnAppProps struct {
	// The name for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	Name *string `json:"name" yaml:"name"`
	// Personal Access token for 3rd party source control system for an Amplify app, used to create webhook and read-only deploy key.
	//
	// Token is not stored.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig interface{} `json:"autoBranchCreationConfig" yaml:"autoBranchCreationConfig"`
	// The credentials for basic authorization for an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig interface{} `json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	//
	// *Pattern:* (?s).+
	BuildSpec *string `json:"buildSpec" yaml:"buildSpec"`
	// The custom HTTP headers for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 25000.
	//
	// *Pattern:* (?s).*
	CustomHeaders *string `json:"customHeaders" yaml:"customHeaders"`
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules interface{} `json:"customRules" yaml:"customRules"`
	// The description for an Amplify app.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	Description *string `json:"description" yaml:"description"`
	// Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion interface{} `json:"enableBranchAutoDeletion" yaml:"enableBranchAutoDeletion"`
	// The environment variables map for an Amplify app.
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// The AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) of the Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	IamServiceRole *string `json:"iamServiceRole" yaml:"iamServiceRole"`
	// The OAuth token for a third-party source control system for an Amplify app.
	//
	// The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	OauthToken *string `json:"oauthToken" yaml:"oauthToken"`
	// The repository for an Amplify app.
	//
	// *Pattern:* (?s).*
	Repository *string `json:"repository" yaml:"repository"`
	// The tag for an Amplify app.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Amplify::Branch`.
//
// The AWS::Amplify::Branch resource creates a new branch within an app.
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnBranch) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBranch(scope constructs.Construct, id *string, props *CfnBranchProps) CfnBranch {
	_init_.Initialize()

	j := jsiiProxy_CfnBranch{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnBranch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::Branch`.
func NewCfnBranch_Override(c CfnBranch, scope constructs.Construct, id *string, props *CfnBranchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnBranch",
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
func CfnBranch_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnBranch",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBranch_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnBranch",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBranch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnBranch",
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
		"aws-cdk-lib.aws_amplify.CfnBranch",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnBranch) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBranch) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnBranch) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Use the BasicAuthConfig property type to set password protection for a specific branch.
//
// TODO: EXAMPLE
//
type CfnBranch_BasicAuthConfigProperty struct {
	// The password for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Password *string `json:"password" yaml:"password"`
	// The user name for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Username *string `json:"username" yaml:"username"`
	// Enables basic authorization for the branch.
	EnableBasicAuth interface{} `json:"enableBasicAuth" yaml:"enableBasicAuth"`
}

// The EnvironmentVariable property type sets environment variables for a specific branch.
//
// Environment variables are key-value pairs that are available at build time.
//
// TODO: EXAMPLE
//
type CfnBranch_EnvironmentVariableProperty struct {
	// The environment variable name.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* (?s).*
	Name *string `json:"name" yaml:"name"`
	// The environment variable value.
	//
	// *Length Constraints:* Maximum length of 5500.
	//
	// *Pattern:* (?s).*
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnBranch`.
//
// TODO: EXAMPLE
//
type CfnBranchProps struct {
	// The unique ID for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 20.
	//
	// *Pattern:* d[a-z0-9]+
	AppId *string `json:"appId" yaml:"appId"`
	// The name for the branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	BranchName *string `json:"branchName" yaml:"branchName"`
	// The basic authorization credentials for a branch of an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig interface{} `json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for the branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	//
	// *Pattern:* (?s).+
	BuildSpec *string `json:"buildSpec" yaml:"buildSpec"`
	// The description for the branch that is part of an Amplify app.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	Description *string `json:"description" yaml:"description"`
	// Enables auto building for the branch.
	EnableAutoBuild interface{} `json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Sets whether the Amplify Console creates a preview for each pull request that is made for this branch.
	//
	// If this property is enabled, the Amplify Console deploys your app to a unique preview URL after each pull request is opened. Development and QA teams can use this preview to test the pull request before it's merged into a production or integration branch.
	//
	// To provide backend support for your preview, the Amplify Console automatically provisions a temporary backend environment that it deletes when the pull request is closed. If you want to specify a dedicated backend environment for your previews, use the `PullRequestEnvironmentName` property.
	//
	// For more information, see [Web Previews](https://docs.aws.amazon.com/amplify/latest/userguide/pr-previews.html) in the *AWS Amplify Console User Guide* .
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// The environment variables for the branch.
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// If pull request previews are enabled for this branch, you can use this property to specify a dedicated backend environment for your previews.
	//
	// For example, you could specify an environment named `prod` , `test` , or `dev` that you initialized with the Amplify CLI and mapped to this branch.
	//
	// To enable pull request previews, set the `EnablePullRequestPreview` property to `true` .
	//
	// If you don't specify an environment, the Amplify Console provides backend support for each preview by automatically provisioning a temporary backend environment. Amplify Console deletes this environment when the pull request is closed.
	//
	// For more information about creating backend environments, see [Feature Branch Deployments and Team Workflows](https://docs.aws.amazon.com/amplify/latest/userguide/multi-environments.html) in the *AWS Amplify Console User Guide* .
	//
	// *Length Constraints:* Maximum length of 20.
	//
	// *Pattern:* (?s).*
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Describes the current stage for the branch.
	//
	// *Valid Values:* PRODUCTION | BETA | DEVELOPMENT | EXPERIMENTAL | PULL_REQUEST
	Stage *string `json:"stage" yaml:"stage"`
	// The tag for the branch.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Amplify::Domain`.
//
// The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnDomain) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnDomain",
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
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
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
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The SubDomainSetting property type enables you to connect a subdomain (for example, example.exampledomain.com) to a specific branch.
//
// TODO: EXAMPLE
//
type CfnDomain_SubDomainSettingProperty struct {
	// The branch name setting for the subdomain.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	BranchName *string `json:"branchName" yaml:"branchName"`
	// The prefix setting for the subdomain.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* (?s).*
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Properties for defining a `CfnDomain`.
//
// TODO: EXAMPLE
//
type CfnDomainProps struct {
	// The unique ID for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 20.
	//
	// *Pattern:* d[a-z0-9]+
	AppId *string `json:"appId" yaml:"appId"`
	// The domain name for the domain association.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* ^(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])(\.)?$
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The setting for the subdomain.
	SubDomainSettings interface{} `json:"subDomainSettings" yaml:"subDomainSettings"`
	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns *[]*string `json:"autoSubDomainCreationPatterns" yaml:"autoSubDomainCreationPatterns"`
	// The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* ^$|^arn:aws:iam::\d{12}:role.+
	AutoSubDomainIamRole *string `json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain interface{} `json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
}

