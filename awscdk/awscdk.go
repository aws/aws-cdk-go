// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Includes API for attaching annotations such as warning messages to constructs.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"import _ "github.com/aws-samples/dummy/constructs"
//
//   type myAspect struct {
//   }
//
//   func (this *myAspect) visit(node iConstruct) {
//   	if *node instanceof cdk.cfnResource && *node.cfnResourceType == "Foo::Bar" {
//   		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
//   	}
//   }
//
//   func (this *myAspect) error(node iConstruct, message *string) {
//   	cdk.annotations.of(*node).addError(*message)
//   }
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string) *myStack {
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id)
//
//   	stack := cdk.NewStack()
//   	cdk.NewCfnResource(stack, jsii.String("Foo"), &cfnResourceProps{
//   		type: jsii.String("Foo::Bar"),
//   		properties: map[string]interface{}{
//   			"Fred": jsii.String("Thud"),
//   		},
//   	})
//   	cdk.aspects.of(stack).add(NewMyAspect())
//   	return this
//   }
//
type Annotations interface {
	// Adds a deprecation warning for a specific API.
	//
	// Deprecations will be added only once per construct as a warning and will be
	// deduplicated based on the `api`.
	//
	// If the environment variable `CDK_BLOCK_DEPRECATIONS` is set, this method
	// will throw an error instead with the deprecation message.
	AddDeprecation(api *string, message *string)
	// Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail deployment of any stack that has errors reported against it.
	AddError(message *string)
	// Adds an info metadata entry to this construct.
	//
	// The CLI will display the info message when apps are synthesized.
	AddInfo(message *string)
	// Adds a warning metadata entry to this construct.
	//
	// The CLI will display the warning when an app is synthesized, or fail if run
	// in --strict mode.
	AddWarning(message *string)
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Returns the annotations API for a construct scope.
func Annotations_Of(scope constructs.IConstruct) Annotations {
	_init_.Initialize()

	var returns Annotations

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Annotations",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) AddDeprecation(api *string, message *string) {
	_jsii_.InvokeVoid(
		a,
		"addDeprecation",
		[]interface{}{api, message},
	)
}

func (a *jsiiProxy_Annotations) AddError(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addError",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddInfo(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addInfo",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddWarning(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addWarning",
		[]interface{}{message},
	)
}

// A construct which represents an entire CDK app. This construct is normally the root of the construct tree.
//
// You would normally define an `App` instance in your program's entrypoint,
// then define constructs where the app is used as the parent scope.
//
// After all the child constructs are defined within the app, you should call
// `app.synth()` which will emit a "cloud assembly" from this app into the
// directory specified by `outdir`. Cloud assemblies includes artifacts such as
// CloudFormation templates and assets that are needed to deploy this app into
// the AWS cloud.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"import lambda "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"type App awscdk.App
//   type Stack awscdk.Stackimport awscdk "github.com/aws/aws-cdk-go/awscdk"type MockIntegration awscdk.MockIntegration
//   type PassthroughBehavior awscdk.PassthroughBehavior
//   type RestApi awscdk.RestApi
//   type TokenAuthorizer awscdk.TokenAuthorizer
//
//   /*
//    * Stack verification steps:
//    * * `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//    * * `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>` should return 403
//    * * `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>` should return 200
//    */
//
//   app := NewApp()
//   stack := NewStack(app, jsii.String("TokenAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.assetCode.fromAsset(path.join(__dirname, jsii.String("integ.token-authorizer.handler"))),
//   })
//
//   restapi := NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := NewTokenAuthorizer(stack, jsii.String("MyAuthorizer"), &tokenAuthorizerProps{
//   	handler: authorizerFn,
//   })
//
//   restapi.root.addMethod(jsii.String("ANY"), NewMockIntegration(&integrationOptions{
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	passthroughBehavior: passthroughBehavior_NEVER,
//   	requestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &methodOptions{
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	authorizer: authorizer,
//   })
//
// See: https://docs.aws.amazon.com/cdk/latest/guide/apps.html
//
type App interface {
	Stage
	// The default account for all resources defined within this stage.
	Account() *string
	// Artifact ID of the assembly if it is a nested stage. The root stage (app) will return an empty string.
	//
	// Derived from the construct path.
	ArtifactId() *string
	// The cloud assembly asset output directory.
	AssetOutdir() *string
	// The tree node.
	Node() constructs.Node
	// The cloud assembly output directory.
	Outdir() *string
	// The parent stage or `undefined` if this is the app.
	//
	// *.
	ParentStage() Stage
	// The default region for all resources defined within this stage.
	Region() *string
	// The name of the stage.
	//
	// Based on names of the parent stages separated by
	// hypens.
	StageName() *string
	// Synthesize this stage into a cloud assembly.
	//
	// Once an assembly has been synthesized, it cannot be modified. Subsequent
	// calls will return the same assembly.
	Synth(options *StageSynthesisOptions) cxapi.CloudAssembly
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	jsiiProxy_Stage
}

func (j *jsiiProxy_App) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) AssetOutdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetOutdir",
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

func (j *jsiiProxy_App) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ParentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"parentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


// Initializes a CDK application.
func NewApp(props *AppProps) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"aws-cdk-lib.App",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Initializes a CDK application.
func NewApp_Override(a App, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.App",
		[]interface{}{props},
		a,
	)
}

// Checks if an object is an instance of the `App` class.
//
// Returns: `true` if `obj` is an `App`.
func App_IsApp(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.App",
		"isApp",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a stage.
func App_IsStage(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.App",
		"isStage",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return the stage this construct is contained with, if available.
//
// If called
// on a nested stage, returns its parent.
func App_Of(construct constructs.IConstruct) Stage {
	_init_.Initialize()

	var returns Stage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.App",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) Synth(options *StageSynthesisOptions) cxapi.CloudAssembly {
	var returns cxapi.CloudAssembly

	_jsii_.Invoke(
		a,
		"synth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

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

// Initialization props for apps.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var context interface{}
//   appProps := &appProps{
//   	analyticsReporting: jsii.Boolean(false),
//   	autoSynth: jsii.Boolean(false),
//   	context: map[string]interface{}{
//   		"contextKey": context,
//   	},
//   	outdir: jsii.String("outdir"),
//   	stackTraces: jsii.Boolean(false),
//   	treeMetadata: jsii.Boolean(false),
//   }
//
type AppProps struct {
	// Include runtime versioning information in the Stacks of this app.
	AnalyticsReporting *bool `json:"analyticsReporting" yaml:"analyticsReporting"`
	// Automatically call `synth()` before the program exits.
	//
	// If you set this, you don't have to call `synth()` explicitly. Note that
	// this feature is only available for certain programming languages, and
	// calling `synth()` is still recommended.
	AutoSynth *bool `json:"autoSynth" yaml:"autoSynth"`
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdk.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	Context *map[string]interface{} `json:"context" yaml:"context"`
	// The output directory into which to emit synthesized artifacts.
	//
	// You should never need to set this value. By default, the value you pass to
	// the CLI's `--output` flag will be used, and if you change it to a different
	// directory the CLI will fail to pick up the generated Cloud Assembly.
	//
	// This property is intended for internal and testing use.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	StackTraces *bool `json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	TreeMetadata *bool `json:"treeMetadata" yaml:"treeMetadata"`
}

type Arn interface {
}

// The jsii proxy struct for Arn
type jsiiProxy_Arn struct {
	_ byte // padding
}

// Extract the full resource name from an ARN.
//
// Necessary for resource names (paths) that may contain the separator, like
// `arn:aws:iam::111111111111:role/path/to/role/name`.
//
// Only works if we statically know the expected `resourceType` beforehand, since we're going
// to use that to split the string on ':<resourceType>/' (and take the right-hand side).
//
// We can't extract the 'resourceType' from the ARN at hand, because CloudFormation Expressions
// only allow literals in the 'separator' argument to `{ Fn::Split }`, and so it can't be
// `{ Fn::Select: [5, { Fn::Split: [':', ARN] }}`.
//
// Only necessary for ARN formats for which the type-name separator is `/`.
func Arn_ExtractResourceName(arn *string, resourceType *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
		"extractResourceName",
		[]interface{}{arn, resourceType},
		&returns,
	)

	return returns
}

// Creates an ARN from components.
//
// If `partition`, `region` or `account` are not specified, the stack's
// partition, region and account will be used.
//
// If any component is the empty string, an empty string will be inserted
// into the generated ARN at the location that component corresponds to.
//
// The ARN will be formatted as follows:
//
//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}{resource-name}
//
// The required ARN pieces that are omitted will be taken from the stack that
// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
// can be 'undefined'.
func Arn_Format(components *ArnComponents, stack Stack) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
		"format",
		[]interface{}{components, stack},
		&returns,
	)

	return returns
}

// Splits the provided ARN into its components.
//
// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
// and a Token representing a dynamic CloudFormation expression
// (in which case the returned components will also be dynamic CloudFormation expressions,
// encoded as Tokens).
func Arn_Split(arn *string, arnFormat ArnFormat) *ArnComponents {
	_init_.Initialize()

	var returns *ArnComponents

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
		"split",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

// Example:
//   subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &publicHostedZoneProps{
//   	zoneName: jsii.String("sub.someexample.com"),
//   })
//
//   // import the delegation role by constructing the roleArn
//   delegationRoleArn := stack.of(this).formatArn(&arnComponents{
//   	region: jsii.String(""),
//   	 // IAM is global in each partition
//   	service: jsii.String("iam"),
//   	account: jsii.String("parent-account-id"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("MyDelegationRole"),
//   })
//   delegationRole := iam.role.fromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)
//
//   // create the record
//   // create the record
//   route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &crossAccountZoneDelegationRecordProps{
//   	delegatedZone: subZone,
//   	parentHostedZoneName: jsii.String("someexample.com"),
//   	 // or you can use parentHostedZoneId
//   	delegationRole: delegationRole,
//   })
//
type ArnComponents struct {
	// Resource type (e.g. "table", "autoScalingGroup", "certificate"). For some resource types, e.g. S3 buckets, this field defines the bucket name.
	Resource *string `json:"resource" yaml:"resource"`
	// The service namespace that identifies the AWS product (for example, 's3', 'iam', 'codepipline').
	Service *string `json:"service" yaml:"service"`
	// The ID of the AWS account that owns the resource, without the hyphens.
	//
	// For example, 123456789012. Note that the ARNs for some resources don't
	// require an account number, so this component might be omitted.
	Account *string `json:"account" yaml:"account"`
	// The specific ARN format to use for this ARN value.
	ArnFormat ArnFormat `json:"arnFormat" yaml:"arnFormat"`
	// The partition that the resource is in.
	//
	// For standard AWS regions, the
	// partition is aws. If you have resources in other partitions, the
	// partition is aws-partitionname. For example, the partition for resources
	// in the China (Beijing) region is aws-cn.
	Partition *string `json:"partition" yaml:"partition"`
	// The region the resource resides in.
	//
	// Note that the ARNs for some resources
	// do not require a region, so this component might be omitted.
	Region *string `json:"region" yaml:"region"`
	// Resource name or path within the resource (i.e. S3 bucket object key) or a wildcard such as ``"*"``. This is service-dependent.
	ResourceName *string `json:"resourceName" yaml:"resourceName"`
}

// An enum representing the various ARN formats that different services use.
type ArnFormat string

const (
	// This represents a format where there is no 'resourceName' part.
	//
	// This format is used for S3 resources,
	// like 'arn:aws:s3:::bucket'.
	// Everything after the last colon is considered the 'resource',
	// even if it contains slashes,
	// like in 'arn:aws:s3:::bucket/object.zip'.
	ArnFormat_NO_RESOURCE_NAME ArnFormat = "NO_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are separated with a colon.
	//
	// Like in: 'arn:aws:service:region:account:resource:resourceName'.
	// Everything after the last colon is considered the 'resourceName',
	// even if it contains slashes,
	// like in 'arn:aws:apigateway:region:account:resource:/test/mydemoresource/*'.
	ArnFormat_COLON_RESOURCE_NAME ArnFormat = "COLON_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are separated with a slash.
	//
	// Like in: 'arn:aws:service:region:account:resource/resourceName'.
	// Everything after the separating slash is considered the 'resourceName',
	// even if it contains colons,
	// like in 'arn:aws:cognito-sync:region:account:identitypool/us-east-1:1a1a1a1a-ffff-1111-9999-12345678:bla'.
	ArnFormat_SLASH_RESOURCE_NAME ArnFormat = "SLASH_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are seperated with a slash, but there is also an additional slash after the colon separating 'account' from 'resource'.
	//
	// Like in: 'arn:aws:service:region:account:/resource/resourceName'.
	// Note that the leading slash is _not_ included in the parsed 'resource' part.
	ArnFormat_SLASH_RESOURCE_SLASH_RESOURCE_NAME ArnFormat = "SLASH_RESOURCE_SLASH_RESOURCE_NAME"
)

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"import _ "github.com/aws-samples/dummy/constructs"
//
//   type myAspect struct {
//   }
//
//   func (this *myAspect) visit(node iConstruct) {
//   	if *node instanceof cdk.cfnResource && *node.cfnResourceType == "Foo::Bar" {
//   		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
//   	}
//   }
//
//   func (this *myAspect) error(node iConstruct, message *string) {
//   	cdk.annotations.of(*node).addError(*message)
//   }
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string) *myStack {
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id)
//
//   	stack := cdk.NewStack()
//   	cdk.NewCfnResource(stack, jsii.String("Foo"), &cfnResourceProps{
//   		type: jsii.String("Foo::Bar"),
//   		properties: map[string]interface{}{
//   			"Fred": jsii.String("Thud"),
//   		},
//   	})
//   	cdk.aspects.of(stack).add(NewMyAspect())
//   	return this
//   }
//
type Aspects interface {
	// The list of aspects which were directly applied on this scope.
	All() *[]IAspect
	// Adds an aspect to apply this scope before synthesis.
	Add(aspect IAspect)
}

// The jsii proxy struct for Aspects
type jsiiProxy_Aspects struct {
	_ byte // padding
}

func (j *jsiiProxy_Aspects) All() *[]IAspect {
	var returns *[]IAspect
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}


// Returns the `Aspects` object associated with a construct scope.
func Aspects_Of(scope constructs.IConstruct) Aspects {
	_init_.Initialize()

	var returns Aspects

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Aspects",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Aspects) Add(aspect IAspect) {
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{aspect},
	)
}

// The type of asset hash.
//
// NOTE: the hash is used in order to identify a specific revision of the asset, and
// used for optimizing and caching deployment activities related to this asset such as
// packaging, uploading to Amazon S3, etc.
type AssetHashType string

const (
	// Based on the content of the source path.
	//
	// When bundling, use `SOURCE` when the content of the bundling output is not
	// stable across repeated bundling operations.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Based on the content of the bundling output.
	//
	// Use `OUTPUT` when the source of the asset is a top level folder containing
	// code and/or dependencies that are not directly linked to the asset.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

// Asset hash options.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
//   assetOptions := &assetOptions{
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   }
//
type AssetOptions struct {
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	AssetHash *string `json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	AssetHashType AssetHashType `json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	Bundling *BundlingOptions `json:"bundling" yaml:"bundling"`
}

// Stages a file or directory from a location on the file system into a staging directory.
//
// This is controlled by the context key 'aws:cdk:asset-staging' and enabled
// by the CLI by default in order to ensure that when the CDK app exists, all
// assets are available for deployment. Otherwise, if an app references assets
// in temporary locations, those will not be available when it exists (see
// https://github.com/aws/aws-cdk/issues/1716).
//
// The `stagedPath` property is a stringified token that represents the location
// of the file or directory after staging. It will be resolved only during the
// "prepare" stage and may be either the original path or the staged path
// depending on the context setting.
//
// The file/directory are staged based on their content hash (fingerprint). This
// means that only if content was changed, copy will happen.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
//   assetStaging := cdk.NewAssetStaging(this, jsii.String("MyAssetStaging"), &assetStagingProps{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   })
//
type AssetStaging interface {
	constructs.Construct
	// Absolute path to the asset data.
	//
	// If asset staging is disabled, this will just be the source path or
	// a temporary directory used for bundling.
	//
	// If asset staging is enabled it will be the staged path.
	//
	// IMPORTANT: If you are going to call `addFileAsset()`, use
	// `relativeStagedPath()` instead.
	AbsoluteStagedPath() *string
	// A cryptographic hash of the asset.
	AssetHash() *string
	// Whether this asset is an archive (zip or jar).
	IsArchive() *bool
	// The tree node.
	Node() constructs.Node
	// How this asset should be packaged.
	Packaging() FileAssetPackaging
	// The absolute path of the asset as it was referenced by the user.
	SourcePath() *string
	// Return the path to the staged asset, relative to the Cloud Assembly (manifest) directory of the given stack.
	//
	// Only returns a relative path if the asset was staged, returns an absolute path if
	// it was not staged.
	//
	// A bundled asset might end up in the outDir and still not count as
	// "staged"; if asset staging is disabled we're technically expected to
	// reference source directories, but we don't have a source directory for the
	// bundled outputs (as the bundle output is written to a temporary
	// directory). Nevertheless, we will still return an absolute path.
	//
	// A non-obvious directory layout may look like this:
	//
	// ```
	//    CLOUD ASSEMBLY ROOT
	//      +-- asset.12345abcdef/
	//      +-- assembly-Stage
	//            +-- MyStack.template.json
	//            +-- MyStack.assets.json <- will contain { "path": "../asset.12345abcdef" }
	// ```.
	RelativeStagedPath(stack Stack) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AssetStaging
type jsiiProxy_AssetStaging struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AssetStaging) AbsoluteStagedPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteStagedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) IsArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Packaging() FileAssetPackaging {
	var returns FileAssetPackaging
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) SourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePath",
		&returns,
	)
	return returns
}


func NewAssetStaging(scope constructs.Construct, id *string, props *AssetStagingProps) AssetStaging {
	_init_.Initialize()

	j := jsiiProxy_AssetStaging{}

	_jsii_.Create(
		"aws-cdk-lib.AssetStaging",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAssetStaging_Override(a AssetStaging, scope constructs.Construct, id *string, props *AssetStagingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.AssetStaging",
		[]interface{}{scope, id, props},
		a,
	)
}

// Clears the asset hash cache.
func AssetStaging_ClearAssetHashCache() {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.AssetStaging",
		"clearAssetHashCache",
		nil, // no parameters
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AssetStaging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.AssetStaging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AssetStaging_BUNDLING_INPUT_DIR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.AssetStaging",
		"BUNDLING_INPUT_DIR",
		&returns,
	)
	return returns
}

func AssetStaging_BUNDLING_OUTPUT_DIR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.AssetStaging",
		"BUNDLING_OUTPUT_DIR",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AssetStaging) RelativeStagedPath(stack Stack) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"relativeStagedPath",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetStaging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization properties for `AssetStaging`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
//   assetStagingProps := &assetStagingProps{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type AssetStagingProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	Follow SymlinkFollowMode `json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	AssetHash *string `json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	AssetHashType AssetHashType `json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	Bundling *BundlingOptions `json:"bundling" yaml:"bundling"`
	// The source file or directory to copy from.
	SourcePath *string `json:"sourcePath" yaml:"sourcePath"`
}

// Accessor for pseudo parameters.
//
// Since pseudo parameters need to be anchored to a stack somewhere in the
// construct tree, this class takes an scope parameter; the pseudo parameter
// values can be obtained as properties from an scoped object.
type Aws interface {
}

// The jsii proxy struct for Aws
type jsiiProxy_Aws struct {
	_ byte // padding
}

func Aws_ACCOUNT_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"ACCOUNT_ID",
		&returns,
	)
	return returns
}

func Aws_NO_VALUE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"NO_VALUE",
		&returns,
	)
	return returns
}

func Aws_NOTIFICATION_ARNS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"NOTIFICATION_ARNS",
		&returns,
	)
	return returns
}

func Aws_PARTITION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"PARTITION",
		&returns,
	)
	return returns
}

func Aws_REGION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"REGION",
		&returns,
	)
	return returns
}

func Aws_STACK_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"STACK_ID",
		&returns,
	)
	return returns
}

func Aws_STACK_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"STACK_NAME",
		&returns,
	)
	return returns
}

func Aws_URL_SUFFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.Aws",
		"URL_SUFFIX",
		&returns,
	)
	return returns
}

// Synthesizer that reuses bootstrap roles from a different region.
//
// A special synthesizer that behaves similarly to `DefaultStackSynthesizer`,
// but doesn't require bootstrapping the environment it operates in. Instead,
// it will re-use the Roles that were created for a different region (which
// is possible because IAM is a global service).
//
// However, it will not assume asset buckets or repositories have been created,
// and therefore does not support assets.
//
// Used by the CodePipeline construct for the support stacks needed for
// cross-region replication S3 buckets. App builders do not need to use this
// synthesizer directly.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   bootstraplessSynthesizer := cdk.NewBootstraplessSynthesizer(&bootstraplessSynthesizerProps{
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	deployRoleArn: jsii.String("deployRoleArn"),
//   })
//
type BootstraplessSynthesizer interface {
	DefaultStackSynthesizer
	// Returns the ARN of the CFN execution Role.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	DeployRoleArn() *string
	Stack() Stack
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(_asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for BootstraplessSynthesizer
type jsiiProxy_BootstraplessSynthesizer struct {
	jsiiProxy_DefaultStackSynthesizer
}

func (j *jsiiProxy_BootstraplessSynthesizer) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) DeployRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBootstraplessSynthesizer(props *BootstraplessSynthesizerProps) BootstraplessSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_BootstraplessSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBootstraplessSynthesizer_Override(b BootstraplessSynthesizer, props *BootstraplessSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		b,
	)
}

func BootstraplessSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		b,
		"addDockerImageAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddFileAsset(_asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		b,
		"addFileAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		b,
		"bind",
		[]interface{}{stack},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		b,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// Construction properties of {@link BootstraplessSynthesizer}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   bootstraplessSynthesizerProps := &bootstraplessSynthesizerProps{
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	deployRoleArn: jsii.String("deployRoleArn"),
//   }
//
type BootstraplessSynthesizerProps struct {
	// The CFN execution Role ARN to use.
	CloudFormationExecutionRoleArn *string `json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The deploy Role ARN to use.
	DeployRoleArn *string `json:"deployRoleArn" yaml:"deployRoleArn"`
}

// Bundling options.
//
// Example:
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &assetProps{
//   	path: path.join(__dirname, jsii.String("markdown-asset")),
//   	 // /asset-input and working directory in the container
//   	bundling: &bundlingOptions{
//   		image: dockerImage.fromBuild(path.join(__dirname, jsii.String("alpine-markdown"))),
//   		 // Build an image
//   		command: []*string{
//   			jsii.String("sh"),
//   			jsii.String("-c"),
//   			jsii.String("\n            markdown index.md > /asset-output/index.html\n          "),
//   		},
//   	},
//   })
//
type BundlingOptions struct {
	// The Docker image where the command will run.
	Image DockerImage `json:"image" yaml:"image"`
	// The command to run in the Docker container.
	//
	// Example value: `['npm', 'install']`.
	// See: https://docs.docker.com/engine/reference/run/
	//
	Command *[]*string `json:"command" yaml:"command"`
	// The entrypoint to run in the Docker container.
	//
	// Example value: `['/bin/sh', '-c']`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the Docker container.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// Local bundling provider.
	//
	// The provider implements a method `tryBundle()` which should return `true`
	// if local bundling was performed. If `false` is returned, docker bundling
	// will be done.
	Local ILocalBundling `json:"local" yaml:"local"`
	// The type of output that this bundling operation is producing.
	OutputType BundlingOutput `json:"outputType" yaml:"outputType"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	SecurityOpt *string `json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the Docker container.
	//
	// user | user:group | uid | uid:gid | user:gid | uid:group.
	// See: https://docs.docker.com/engine/reference/run/#user
	//
	User *string `json:"user" yaml:"user"`
	// Additional Docker volumes to mount.
	Volumes *[]*DockerVolume `json:"volumes" yaml:"volumes"`
	// Working directory inside the Docker container.
	WorkingDirectory *string `json:"workingDirectory" yaml:"workingDirectory"`
}

// The type of output that a bundling operation is producing.
//
// Example:
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &assetProps{
//   	path: jsii.String("/path/to/asset"),
//   	bundling: &bundlingOptions{
//   		image: dockerImage.fromRegistry(jsii.String("alpine")),
//   		command: []*string{
//   			jsii.String("command-that-produces-an-archive.sh"),
//   		},
//   		outputType: bundlingOutput_NOT_ARCHIVED,
//   	},
//   })
//
type BundlingOutput string

const (
	// The bundling output directory includes a single .zip or .jar file which will be used as the final bundle. If the output directory does not include exactly a single archive, bundling will fail.
	BundlingOutput_ARCHIVED BundlingOutput = "ARCHIVED"
	// The bundling output directory contains one or more files which will be archived and uploaded as a .zip file to S3.
	BundlingOutput_NOT_ARCHIVED BundlingOutput = "NOT_ARCHIVED"
	// If the bundling output directory contains a single archive file (zip or jar) it will be used as the bundle output as-is.
	//
	// Otherwise all the files in the bundling output directory will be zipped.
	BundlingOutput_AUTO_DISCOVER BundlingOutput = "AUTO_DISCOVER"
)

// Specifies whether an Auto Scaling group and the instances it contains are replaced during an update.
//
// During replacement,
// AWS CloudFormation retains the old group until it finishes creating the new one. If the update fails, AWS CloudFormation
// can roll back to the old Auto Scaling group and delete the new Auto Scaling group.
//
// While AWS CloudFormation creates the new group, it doesn't detach or attach any instances. After successfully creating
// the new Auto Scaling group, AWS CloudFormation deletes the old Auto Scaling group during the cleanup process.
//
// When you set the WillReplace parameter, remember to specify a matching CreationPolicy. If the minimum number of
// instances (specified by the MinSuccessfulInstancesPercent property) don't signal success within the Timeout period
// (specified in the CreationPolicy policy), the replacement update fails and AWS CloudFormation rolls back to the old
// Auto Scaling group.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnAutoScalingReplacingUpdate := &cfnAutoScalingReplacingUpdate{
//   	willReplace: jsii.Boolean(false),
//   }
//
type CfnAutoScalingReplacingUpdate struct {
	WillReplace *bool `json:"willReplace" yaml:"willReplace"`
}

// To specify how AWS CloudFormation handles rolling updates for an Auto Scaling group, use the AutoScalingRollingUpdate policy.
//
// Rolling updates enable you to specify whether AWS CloudFormation updates instances that are in an Auto Scaling
// group in batches or all at once.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnAutoScalingRollingUpdate := &cfnAutoScalingRollingUpdate{
//   	maxBatchSize: jsii.Number(123),
//   	minInstancesInService: jsii.Number(123),
//   	minSuccessfulInstancesPercent: jsii.Number(123),
//   	pauseTime: jsii.String("pauseTime"),
//   	suspendProcesses: []*string{
//   		jsii.String("suspendProcesses"),
//   	},
//   	waitOnResourceSignals: jsii.Boolean(false),
//   }
//
type CfnAutoScalingRollingUpdate struct {
	// Specifies the maximum number of instances that AWS CloudFormation updates.
	MaxBatchSize *float64 `json:"maxBatchSize" yaml:"maxBatchSize"`
	// Specifies the minimum number of instances that must be in service within the Auto Scaling group while AWS CloudFormation updates old instances.
	MinInstancesInService *float64 `json:"minInstancesInService" yaml:"minInstancesInService"`
	// Specifies the percentage of instances in an Auto Scaling rolling update that must signal success for an update to succeed.
	//
	// You can specify a value from 0 to 100. AWS CloudFormation rounds to the nearest tenth of a percent. For example, if you
	// update five instances with a minimum successful percentage of 50, three instances must signal success.
	//
	// If an instance doesn't send a signal within the time specified in the PauseTime property, AWS CloudFormation assumes
	// that the instance wasn't updated.
	//
	// If you specify this property, you must also enable the WaitOnResourceSignals and PauseTime properties.
	MinSuccessfulInstancesPercent *float64 `json:"minSuccessfulInstancesPercent" yaml:"minSuccessfulInstancesPercent"`
	// The amount of time that AWS CloudFormation pauses after making a change to a batch of instances to give those instances time to start software applications.
	//
	// For example, you might need to specify PauseTime when scaling up the number of
	// instances in an Auto Scaling group.
	//
	// If you enable the WaitOnResourceSignals property, PauseTime is the amount of time that AWS CloudFormation should wait
	// for the Auto Scaling group to receive the required number of valid signals from added or replaced instances. If the
	// PauseTime is exceeded before the Auto Scaling group receives the required number of signals, the update fails. For best
	// results, specify a time period that gives your applications sufficient time to get started. If the update needs to be
	// rolled back, a short PauseTime can cause the rollback to fail.
	//
	// Specify PauseTime in the ISO8601 duration format (in the format PT#H#M#S, where each # is the number of hours, minutes,
	// and seconds, respectively). The maximum PauseTime is one hour (PT1H).
	PauseTime *string `json:"pauseTime" yaml:"pauseTime"`
	// Specifies the Auto Scaling processes to suspend during a stack update.
	//
	// Suspending processes prevents Auto Scaling from
	// interfering with a stack update. For example, you can suspend alarming so that Auto Scaling doesn't execute scaling
	// policies associated with an alarm. For valid values, see the ScalingProcesses.member.N parameter for the SuspendProcesses
	// action in the Auto Scaling API Reference.
	SuspendProcesses *[]*string `json:"suspendProcesses" yaml:"suspendProcesses"`
	// Specifies whether the Auto Scaling group waits on signals from new instances during an update.
	//
	// Use this property to
	// ensure that instances have completed installing and configuring applications before the Auto Scaling group update proceeds.
	// AWS CloudFormation suspends the update of an Auto Scaling group after new EC2 instances are launched into the group.
	// AWS CloudFormation must receive a signal from each new instance within the specified PauseTime before continuing the update.
	// To signal the Auto Scaling group, use the cfn-signal helper script or SignalResource API.
	//
	// To have instances wait for an Elastic Load Balancing health check before they signal success, add a health-check
	// verification by using the cfn-init helper script. For an example, see the verify_instance_health command in the Auto Scaling
	// rolling updates sample template.
	WaitOnResourceSignals *bool `json:"waitOnResourceSignals" yaml:"waitOnResourceSignals"`
}

// With scheduled actions, the group size properties of an Auto Scaling group can change at any time.
//
// When you update a
// stack with an Auto Scaling group and scheduled action, AWS CloudFormation always sets the group size property values of
// your Auto Scaling group to the values that are defined in the AWS::AutoScaling::AutoScalingGroup resource of your template,
// even if a scheduled action is in effect.
//
// If you do not want AWS CloudFormation to change any of the group size property values when you have a scheduled action in
// effect, use the AutoScalingScheduledAction update policy to prevent AWS CloudFormation from changing the MinSize, MaxSize,
// or DesiredCapacity properties unless you have modified these values in your template.\
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnAutoScalingScheduledAction := &cfnAutoScalingScheduledAction{
//   	ignoreUnmodifiedGroupSizeProperties: jsii.Boolean(false),
//   }
//
type CfnAutoScalingScheduledAction struct {
	IgnoreUnmodifiedGroupSizeProperties *bool `json:"ignoreUnmodifiedGroupSizeProperties" yaml:"ignoreUnmodifiedGroupSizeProperties"`
}

// Capabilities that affect whether CloudFormation is allowed to change IAM resources.
type CfnCapabilities string

const (
	// No IAM Capabilities.
	//
	// Pass this capability if you wish to block the creation IAM resources.
	CfnCapabilities_NONE CfnCapabilities = "NONE"
	// Capability to create anonymous IAM resources.
	//
	// Pass this capability if you're only creating anonymous resources.
	CfnCapabilities_ANONYMOUS_IAM CfnCapabilities = "ANONYMOUS_IAM"
	// Capability to create named IAM resources.
	//
	// Pass this capability if you're creating IAM resources that have physical
	// names.
	//
	// `CloudFormationCapabilities.NamedIAM` implies `CloudFormationCapabilities.IAM`; you don't have to pass both.
	CfnCapabilities_NAMED_IAM CfnCapabilities = "NAMED_IAM"
	// Capability to run CloudFormation macros.
	//
	// Pass this capability if your template includes macros, for example AWS::Include or AWS::Serverless.
	CfnCapabilities_AUTO_EXPAND CfnCapabilities = "AUTO_EXPAND"
)

// Additional options for the blue/green deployment.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.additionalOptions} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenAdditionalOptions := &cfnCodeDeployBlueGreenAdditionalOptions{
//   	terminationWaitTimeInMinutes: jsii.Number(123),
//   }
//
type CfnCodeDeployBlueGreenAdditionalOptions struct {
	// Specifies time to wait, in minutes, before terminating the blue resources.
	TerminationWaitTimeInMinutes *float64 `json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

// The application actually being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenHookProps.applications} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenApplication := &cfnCodeDeployBlueGreenApplication{
//   	ecsAttributes: &cfnCodeDeployBlueGreenEcsAttributes{
//   		taskDefinitions: []*string{
//   			jsii.String("taskDefinitions"),
//   		},
//   		taskSets: []*string{
//   			jsii.String("taskSets"),
//   		},
//   		trafficRouting: &cfnTrafficRouting{
//   			prodTrafficRoute: &cfnTrafficRoute{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   			targetGroups: []*string{
//   				jsii.String("targetGroups"),
//   			},
//   			testTrafficRoute: &cfnTrafficRoute{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	target: &cfnCodeDeployBlueGreenApplicationTarget{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnCodeDeployBlueGreenApplication struct {
	// The detailed attributes of the deployed target.
	EcsAttributes *CfnCodeDeployBlueGreenEcsAttributes `json:"ecsAttributes" yaml:"ecsAttributes"`
	// The target that is being deployed.
	Target *CfnCodeDeployBlueGreenApplicationTarget `json:"target" yaml:"target"`
}

// Type of the {@link CfnCodeDeployBlueGreenApplication.target} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenApplicationTarget := &cfnCodeDeployBlueGreenApplicationTarget{
//   	logicalId: jsii.String("logicalId"),
//   	type: jsii.String("type"),
//   }
//
type CfnCodeDeployBlueGreenApplicationTarget struct {
	// The logical id of the target resource.
	LogicalId *string `json:"logicalId" yaml:"logicalId"`
	// The resource type of the target being deployed.
	//
	// Right now, the only allowed value is 'AWS::ECS::Service'.
	Type *string `json:"type" yaml:"type"`
}

// The attributes of the ECS Service being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenApplication.ecsAttributes} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenEcsAttributes := &cfnCodeDeployBlueGreenEcsAttributes{
//   	taskDefinitions: []*string{
//   		jsii.String("taskDefinitions"),
//   	},
//   	taskSets: []*string{
//   		jsii.String("taskSets"),
//   	},
//   	trafficRouting: &cfnTrafficRouting{
//   		prodTrafficRoute: &cfnTrafficRoute{
//   			logicalId: jsii.String("logicalId"),
//   			type: jsii.String("type"),
//   		},
//   		targetGroups: []*string{
//   			jsii.String("targetGroups"),
//   		},
//   		testTrafficRoute: &cfnTrafficRoute{
//   			logicalId: jsii.String("logicalId"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnCodeDeployBlueGreenEcsAttributes struct {
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskDefinition task definitions.
	TaskDefinitions *[]*string `json:"taskDefinitions" yaml:"taskDefinitions"`
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskSet task sets.
	TaskSets *[]*string `json:"taskSets" yaml:"taskSets"`
	// The traffic routing configuration.
	TrafficRouting *CfnTrafficRouting `json:"trafficRouting" yaml:"trafficRouting"`
}

// A CloudFormation Hook for CodeDeploy blue-green ECS deployments.
//
// Example:
//   var cfnTemplate cfnInclude
//
//   // mutating the hook
//   var myRole role
//   hook := cfnTemplate.getHook(jsii.String("MyOutput"))
//   codeDeployHook := hook.(cfnCodeDeployBlueGreenHook)
//   codeDeployHook.serviceRole = myRole.roleArn
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html#blue-green-template-reference
//
type CfnCodeDeployBlueGreenHook interface {
	CfnHook
	// Additional options for the blue/green deployment.
	AdditionalOptions() *CfnCodeDeployBlueGreenAdditionalOptions
	SetAdditionalOptions(val *CfnCodeDeployBlueGreenAdditionalOptions)
	// Properties of the Amazon ECS applications being deployed.
	Applications() *[]*CfnCodeDeployBlueGreenApplication
	SetApplications(val *[]*CfnCodeDeployBlueGreenApplication)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda {@link CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic}
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	LifecycleEventHooks() *CfnCodeDeployBlueGreenLifecycleEventHooks
	SetLifecycleEventHooks(val *CfnCodeDeployBlueGreenLifecycleEventHooks)
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
	// The tree node.
	Node() constructs.Node
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Traffic routing configuration settings.
	TrafficRoutingConfig() *CfnTrafficRoutingConfig
	SetTrafficRoutingConfig(val *CfnTrafficRoutingConfig)
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type() *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(_props *map[string]interface{}) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnCodeDeployBlueGreenHook
type jsiiProxy_CfnCodeDeployBlueGreenHook struct {
	jsiiProxy_CfnHook
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) AdditionalOptions() *CfnCodeDeployBlueGreenAdditionalOptions {
	var returns *CfnCodeDeployBlueGreenAdditionalOptions
	_jsii_.Get(
		j,
		"additionalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Applications() *[]*CfnCodeDeployBlueGreenApplication {
	var returns *[]*CfnCodeDeployBlueGreenApplication
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) LifecycleEventHooks() *CfnCodeDeployBlueGreenLifecycleEventHooks {
	var returns *CfnCodeDeployBlueGreenLifecycleEventHooks
	_jsii_.Get(
		j,
		"lifecycleEventHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) TrafficRoutingConfig() *CfnTrafficRoutingConfig {
	var returns *CfnTrafficRoutingConfig
	_jsii_.Get(
		j,
		"trafficRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new CodeDeploy blue-green ECS Hook.
func NewCfnCodeDeployBlueGreenHook(scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) CfnCodeDeployBlueGreenHook {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeDeployBlueGreenHook{}

	_jsii_.Create(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new CodeDeploy blue-green ECS Hook.
func NewCfnCodeDeployBlueGreenHook_Override(c CfnCodeDeployBlueGreenHook, scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) SetAdditionalOptions(val *CfnCodeDeployBlueGreenAdditionalOptions) {
	_jsii_.Set(
		j,
		"additionalOptions",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) SetApplications(val *[]*CfnCodeDeployBlueGreenApplication) {
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) SetLifecycleEventHooks(val *CfnCodeDeployBlueGreenLifecycleEventHooks) {
	_jsii_.Set(
		j,
		"lifecycleEventHooks",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) SetTrafficRoutingConfig(val *CfnTrafficRoutingConfig) {
	_jsii_.Set(
		j,
		"trafficRoutingConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCodeDeployBlueGreenHook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCodeDeployBlueGreenHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) RenderProperties(_props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{_props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link CfnCodeDeployBlueGreenHook}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenHookProps := &cfnCodeDeployBlueGreenHookProps{
//   	applications: []cfnCodeDeployBlueGreenApplication{
//   		&cfnCodeDeployBlueGreenApplication{
//   			ecsAttributes: &cfnCodeDeployBlueGreenEcsAttributes{
//   				taskDefinitions: []*string{
//   					jsii.String("taskDefinitions"),
//   				},
//   				taskSets: []*string{
//   					jsii.String("taskSets"),
//   				},
//   				trafficRouting: &cfnTrafficRouting{
//   					prodTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   					targetGroups: []*string{
//   						jsii.String("targetGroups"),
//   					},
//   					testTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			target: &cfnCodeDeployBlueGreenApplicationTarget{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
//   	additionalOptions: &cfnCodeDeployBlueGreenAdditionalOptions{
//   		terminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   	lifecycleEventHooks: &cfnCodeDeployBlueGreenLifecycleEventHooks{
//   		afterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   		afterAllowTraffic: jsii.String("afterAllowTraffic"),
//   		afterInstall: jsii.String("afterInstall"),
//   		beforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   		beforeInstall: jsii.String("beforeInstall"),
//   	},
//   	trafficRoutingConfig: &cfnTrafficRoutingConfig{
//   		type: cdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   		// the properties below are optional
//   		timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
//   		},
//   		timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnCodeDeployBlueGreenHookProps struct {
	// Properties of the Amazon ECS applications being deployed.
	Applications *[]*CfnCodeDeployBlueGreenApplication `json:"applications" yaml:"applications"`
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// Additional options for the blue/green deployment.
	AdditionalOptions *CfnCodeDeployBlueGreenAdditionalOptions `json:"additionalOptions" yaml:"additionalOptions"`
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda {@link CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic}
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	LifecycleEventHooks *CfnCodeDeployBlueGreenLifecycleEventHooks `json:"lifecycleEventHooks" yaml:"lifecycleEventHooks"`
	// Traffic routing configuration settings.
	TrafficRoutingConfig *CfnTrafficRoutingConfig `json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}

// Lifecycle events for blue-green deployments.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.lifecycleEventHooks} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployBlueGreenLifecycleEventHooks := &cfnCodeDeployBlueGreenLifecycleEventHooks{
//   	afterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   	afterAllowTraffic: jsii.String("afterAllowTraffic"),
//   	afterInstall: jsii.String("afterInstall"),
//   	beforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   	beforeInstall: jsii.String("beforeInstall"),
//   }
//
type CfnCodeDeployBlueGreenLifecycleEventHooks struct {
	// Function to use to run tasks after the test listener serves traffic to the replacement task set.
	AfterAllowTestTraffic *string `json:"afterAllowTestTraffic" yaml:"afterAllowTestTraffic"`
	// Function to use to run tasks after the second target group serves traffic to the replacement task set.
	AfterAllowTraffic *string `json:"afterAllowTraffic" yaml:"afterAllowTraffic"`
	// Function to use to run tasks after the replacement task set is created and one of the target groups is associated with it.
	AfterInstall *string `json:"afterInstall" yaml:"afterInstall"`
	// Function to use to run tasks after the second target group is associated with the replacement task set, but before traffic is shifted to the replacement task set.
	BeforeAllowTraffic *string `json:"beforeAllowTraffic" yaml:"beforeAllowTraffic"`
	// Function to use to run tasks before the replacement task set is created.
	BeforeInstall *string `json:"beforeInstall" yaml:"beforeInstall"`
}

// To perform an AWS CodeDeploy deployment when the version changes on an AWS::Lambda::Alias resource, use the CodeDeployLambdaAliasUpdate update policy.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCodeDeployLambdaAliasUpdate := &cfnCodeDeployLambdaAliasUpdate{
//   	applicationName: jsii.String("applicationName"),
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	afterAllowTrafficHook: jsii.String("afterAllowTrafficHook"),
//   	beforeAllowTrafficHook: jsii.String("beforeAllowTrafficHook"),
//   }
//
type CfnCodeDeployLambdaAliasUpdate struct {
	// The name of the AWS CodeDeploy application.
	ApplicationName *string `json:"applicationName" yaml:"applicationName"`
	// The name of the AWS CodeDeploy deployment group.
	//
	// This is where the traffic-shifting policy is set.
	DeploymentGroupName *string `json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The name of the Lambda function to run after traffic routing completes.
	AfterAllowTrafficHook *string `json:"afterAllowTrafficHook" yaml:"afterAllowTrafficHook"`
	// The name of the Lambda function to run before traffic routing starts.
	BeforeAllowTrafficHook *string `json:"beforeAllowTrafficHook" yaml:"beforeAllowTrafficHook"`
}

// Represents a CloudFormation condition, for resources which must be conditionally created and the determination must be made at deploy time.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &cfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.cfnOptions.condition = NewCfnCondition(this, jsii.String("EnableBucket"), &cfnConditionProps{
//   })
//   rawBucket.cfnOptions.metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnCondition interface {
	CfnElement
	ICfnConditionExpression
	IResolvable
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The condition statement.
	Expression() ICfnConditionExpression
	SetExpression(val ICfnConditionExpression)
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
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Synthesizes the condition.
	Resolve(_context IResolveContext) interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnCondition
type jsiiProxy_CfnCondition struct {
	jsiiProxy_CfnElement
	jsiiProxy_ICfnConditionExpression
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_CfnCondition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Expression() ICfnConditionExpression {
	var returns ICfnConditionExpression
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCondition) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Build a new condition.
//
// The condition must be constructed with a condition token,
// that the condition is based on.
func NewCfnCondition(scope constructs.Construct, id *string, props *CfnConditionProps) CfnCondition {
	_init_.Initialize()

	j := jsiiProxy_CfnCondition{}

	_jsii_.Create(
		"aws-cdk-lib.CfnCondition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Build a new condition.
//
// The condition must be constructed with a condition token,
// that the condition is based on.
func NewCfnCondition_Override(c CfnCondition, scope constructs.Construct, id *string, props *CfnConditionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnCondition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCondition) SetExpression(val ICfnConditionExpression) {
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCondition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCondition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCondition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCondition) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &cfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.cfnOptions.condition = NewCfnCondition(this, jsii.String("EnableBucket"), &cfnConditionProps{
//   })
//   rawBucket.cfnOptions.metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnConditionProps struct {
	// The expression that the condition will evaluate.
	Expression ICfnConditionExpression `json:"expression" yaml:"expression"`
}

// Associate the CreationPolicy attribute with a resource to prevent its status from reaching create complete until AWS CloudFormation receives a specified number of success signals or the timeout period is exceeded.
//
// To signal a
// resource, you can use the cfn-signal helper script or SignalResource API. AWS CloudFormation publishes valid signals
// to the stack events so that you track the number of signals sent.
//
// The creation policy is invoked only when AWS CloudFormation creates the associated resource. Currently, the only
// AWS CloudFormation resources that support creation policies are AWS::AutoScaling::AutoScalingGroup, AWS::EC2::Instance,
// and AWS::CloudFormation::WaitCondition.
//
// Use the CreationPolicy attribute when you want to wait on resource configuration actions before stack creation proceeds.
// For example, if you install and configure software applications on an EC2 instance, you might want those applications to
// be running before proceeding. In such cases, you can add a CreationPolicy attribute to the instance, and then send a success
// signal to the instance after the applications are installed and configured. For a detailed example, see Deploying Applications
// on Amazon EC2 with AWS CloudFormation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCreationPolicy := &cfnCreationPolicy{
//   	autoScalingCreationPolicy: &cfnResourceAutoScalingCreationPolicy{
//   		minSuccessfulInstancesPercent: jsii.Number(123),
//   	},
//   	resourceSignal: &cfnResourceSignal{
//   		count: jsii.Number(123),
//   		timeout: jsii.String("timeout"),
//   	},
//   }
//
type CfnCreationPolicy struct {
	// For an Auto Scaling group replacement update, specifies how many instances must signal success for the update to succeed.
	AutoScalingCreationPolicy *CfnResourceAutoScalingCreationPolicy `json:"autoScalingCreationPolicy" yaml:"autoScalingCreationPolicy"`
	// When AWS CloudFormation creates the associated resource, configures the number of required success signals and the length of time that AWS CloudFormation waits for those signals.
	ResourceSignal *CfnResourceSignal `json:"resourceSignal" yaml:"resourceSignal"`
}

// A CloudFormation `AWS::CloudFormation::CustomResource`.
//
// In a CloudFormation template, you use the `AWS::CloudFormation::CustomResource` or `Custom:: *String*` resource type to specify custom resources.
//
// Custom resources provide a way for you to write custom provisioning logic in CloudFormation template and have CloudFormation run it during a stack operation, such as when you create, update or delete a stack. For more information, see [Custom resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) .
//
// > If you use the [VPC endpoints](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints.html) feature, custom resources in the VPC must have access to CloudFormation -specific Amazon Simple Storage Service ( Amazon S3 ) buckets. Custom resources must send responses to a presigned Amazon S3 URL. If they can't send responses to Amazon S3 , CloudFormation won't receive a response and the stack operation fails. For more information, see [Setting up VPC endpoints for AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-vpce-bucketnames.html) .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCustomResource := cdk.NewCfnCustomResource(this, jsii.String("MyCfnCustomResource"), &cfnCustomResourceProps{
//   	serviceToken: jsii.String("serviceToken"),
//   })
//
type CfnCustomResource interface {
	CfnResource
	IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// > Only one property is defined by AWS for a custom resource: `ServiceToken` .
	//
	// All other properties are defined by the service provider.
	//
	// The service token that was given to the template developer by the service provider to access the service, such as an Amazon SNS topic ARN or Lambda function ARN. The service token must be from the same Region in which you are creating the stack.
	//
	// Updates aren't supported.
	ServiceToken() *string
	SetServiceToken(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnCustomResource
type jsiiProxy_CfnCustomResource struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnCustomResource) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomResource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::CustomResource`.
func NewCfnCustomResource(scope constructs.Construct, id *string, props *CfnCustomResourceProps) CfnCustomResource {
	_init_.Initialize()

	j := jsiiProxy_CfnCustomResource{}

	_jsii_.Create(
		"aws-cdk-lib.CfnCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::CustomResource`.
func NewCfnCustomResource_Override(c CfnCustomResource, scope constructs.Construct, id *string, props *CfnCustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnCustomResource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomResource) SetServiceToken(val *string) {
	_jsii_.Set(
		j,
		"serviceToken",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCustomResource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCustomResource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCustomResource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCustomResource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnCustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomResource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnCustomResource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomResource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomResource) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomResource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomResource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomResource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomResource) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCustomResource) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomResource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomResource) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCustomResource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomResource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomResource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCustomResource`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnCustomResourceProps := &cfnCustomResourceProps{
//   	serviceToken: jsii.String("serviceToken"),
//   }
//
type CfnCustomResourceProps struct {
	// > Only one property is defined by AWS for a custom resource: `ServiceToken` .
	//
	// All other properties are defined by the service provider.
	//
	// The service token that was given to the template developer by the service provider to access the service, such as an Amazon SNS topic ARN or Lambda function ARN. The service token must be from the same Region in which you are creating the stack.
	//
	// Updates aren't supported.
	ServiceToken *string `json:"serviceToken" yaml:"serviceToken"`
}

// With the DeletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted.
//
// You specify a DeletionPolicy attribute for each resource that you want to control. If a resource has no DeletionPolicy
// attribute, AWS CloudFormation deletes the resource by default. Note that this capability also applies to update operations
// that lead to resources being removed.
type CfnDeletionPolicy string

const (
	// AWS CloudFormation deletes the resource and all its content if applicable during stack deletion.
	//
	// You can add this
	// deletion policy to any resource type. By default, if you don't specify a DeletionPolicy, AWS CloudFormation deletes
	// your resources. However, be aware of the following considerations:
	CfnDeletionPolicy_DELETE CfnDeletionPolicy = "DELETE"
	// AWS CloudFormation keeps the resource without deleting the resource or its contents when its stack is deleted.
	//
	// You can add this deletion policy to any resource type. Note that when AWS CloudFormation completes the stack deletion,
	// the stack will be in Delete_Complete state; however, resources that are retained continue to exist and continue to incur
	// applicable charges until you delete those resources.
	CfnDeletionPolicy_RETAIN CfnDeletionPolicy = "RETAIN"
	// For resources that support snapshots (AWS::EC2::Volume, AWS::ElastiCache::CacheCluster, AWS::ElastiCache::ReplicationGroup, AWS::RDS::DBInstance, AWS::RDS::DBCluster, and AWS::Redshift::Cluster), AWS CloudFormation creates a snapshot for the resource before deleting it.
	//
	// Note that when AWS CloudFormation completes the stack deletion, the stack will be in the
	// Delete_Complete state; however, the snapshots that are created with this policy continue to exist and continue to
	// incur applicable charges until you delete those snapshots.
	CfnDeletionPolicy_SNAPSHOT CfnDeletionPolicy = "SNAPSHOT"
)

// References a dynamically retrieved value.
//
// This is a Construct so that subclasses will (eventually) be able to attach
// metadata to themselves without having to change call signatures.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewCfnDynamicReference(cfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html
//
type CfnDynamicReference interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Produce the Token's value at resolution time.
	Resolve(_context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
}

// The jsii proxy struct for CfnDynamicReference
type jsiiProxy_CfnDynamicReference struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_CfnDynamicReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


func NewCfnDynamicReference(service CfnDynamicReferenceService, key *string) CfnDynamicReference {
	_init_.Initialize()

	j := jsiiProxy_CfnDynamicReference{}

	_jsii_.Create(
		"aws-cdk-lib.CfnDynamicReference",
		[]interface{}{service, key},
		&j,
	)

	return &j
}

func NewCfnDynamicReference_Override(c CfnDynamicReference, service CfnDynamicReferenceService, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnDynamicReference",
		[]interface{}{service, key},
		c,
	)
}

func (c *jsiiProxy_CfnDynamicReference) NewError(message *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDynamicReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Dynamic Reference.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnDynamicReferenceProps := &cfnDynamicReferenceProps{
//   	referenceKey: jsii.String("referenceKey"),
//   	service: cdk.cfnDynamicReferenceService_SSM,
//   }
//
type CfnDynamicReferenceProps struct {
	// The reference key of the dynamic reference.
	ReferenceKey *string `json:"referenceKey" yaml:"referenceKey"`
	// The service to retrieve the dynamic reference from.
	Service CfnDynamicReferenceService `json:"service" yaml:"service"`
}

// The service to retrieve the dynamic reference from.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewCfnDynamicReference(cfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
//
type CfnDynamicReferenceService string

const (
	// Plaintext value stored in AWS Systems Manager Parameter Store.
	CfnDynamicReferenceService_SSM CfnDynamicReferenceService = "SSM"
	// Secure string stored in AWS Systems Manager Parameter Store.
	CfnDynamicReferenceService_SSM_SECURE CfnDynamicReferenceService = "SSM_SECURE"
	// Secret stored in AWS Secrets Manager.
	CfnDynamicReferenceService_SECRETS_MANAGER CfnDynamicReferenceService = "SECRETS_MANAGER"
)

// An element of a CloudFormation stack.
type CfnElement interface {
	constructs.Construct
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnElement
type jsiiProxy_CfnElement struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CfnElement) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElement) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnElement) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates an entity and binds it to a tree.
//
// Note that the root of the tree must be a Stack object (not just any Root).
func NewCfnElement_Override(c CfnElement, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnElement",
		[]interface{}{scope, id},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnElement_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnElement",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnElement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnElement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnElement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnElement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a CloudFormation resource.
//
// Example:
//   var cfnTemplate cfnInclude
//
//   // mutating the hook
//   var myRole role
//   hook := cfnTemplate.getHook(jsii.String("MyOutput"))
//   codeDeployHook := hook.(cfnCodeDeployBlueGreenHook)
//   codeDeployHook.serviceRole = myRole.roleArn
//
type CfnHook interface {
	CfnElement
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type() *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnHook
type jsiiProxy_CfnHook struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnHook) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHook) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHook) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHook) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new Hook object.
func NewCfnHook(scope constructs.Construct, id *string, props *CfnHookProps) CfnHook {
	_init_.Initialize()

	j := jsiiProxy_CfnHook{}

	_jsii_.Create(
		"aws-cdk-lib.CfnHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new Hook object.
func NewCfnHook_Override(c CfnHook, scope constructs.Construct, id *string, props *CfnHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnHook",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHook) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::CloudFormation::HookDefaultVersion`.
//
// The `HookDefaultVersion` resource specifies the default version of the hook. The default version of the hook is used in CloudFormation operations for this AWS account and AWS Region .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookDefaultVersion := cdk.NewCfnHookDefaultVersion(this, jsii.String("MyCfnHookDefaultVersion"), &cfnHookDefaultVersionProps{
//   	typeName: jsii.String("typeName"),
//   	typeVersionArn: jsii.String("typeVersionArn"),
//   	versionId: jsii.String("versionId"),
//   })
//
type CfnHookDefaultVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the activated extension, in this account and Region.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The name of the hook.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName() *string
	SetTypeName(val *string)
	// The version ID of the type configuration.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn() *string
	SetTypeVersionArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The version ID of the type specified.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId() *string
	SetVersionId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnHookDefaultVersion
type jsiiProxy_CfnHookDefaultVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnHookDefaultVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) TypeVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::HookDefaultVersion`.
func NewCfnHookDefaultVersion(scope constructs.Construct, id *string, props *CfnHookDefaultVersionProps) CfnHookDefaultVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnHookDefaultVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnHookDefaultVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::HookDefaultVersion`.
func NewCfnHookDefaultVersion_Override(c CfnHookDefaultVersion, scope constructs.Construct, id *string, props *CfnHookDefaultVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnHookDefaultVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHookDefaultVersion) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CfnHookDefaultVersion) SetTypeVersionArn(val *string) {
	_jsii_.Set(
		j,
		"typeVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnHookDefaultVersion) SetVersionId(val *string) {
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHookDefaultVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookDefaultVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHookDefaultVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookDefaultVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnHookDefaultVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookDefaultVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHookDefaultVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnHookDefaultVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHookDefaultVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnHookDefaultVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookDefaultVersionProps := &cfnHookDefaultVersionProps{
//   	typeName: jsii.String("typeName"),
//   	typeVersionArn: jsii.String("typeVersionArn"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnHookDefaultVersionProps struct {
	// The name of the hook.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The version ID of the type configuration.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn *string `json:"typeVersionArn" yaml:"typeVersionArn"`
	// The version ID of the type specified.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId *string `json:"versionId" yaml:"versionId"`
}

// Construction properties of {@link CfnHook}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var properties interface{}
//   cfnHookProps := &cfnHookProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	properties: map[string]interface{}{
//   		"propertiesKey": properties,
//   	},
//   }
//
type CfnHookProps struct {
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type *string `json:"type" yaml:"type"`
	// The properties of the hook.
	Properties *map[string]interface{} `json:"properties" yaml:"properties"`
}

// A CloudFormation `AWS::CloudFormation::HookTypeConfig`.
//
// The `HookTypeConfig` resource specifies the configuration of a hook.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookTypeConfig := cdk.NewCfnHookTypeConfig(this, jsii.String("MyCfnHookTypeConfig"), &cfnHookTypeConfigProps{
//   	configuration: jsii.String("configuration"),
//
//   	// the properties below are optional
//   	configurationAlias: jsii.String("configurationAlias"),
//   	typeArn: jsii.String("typeArn"),
//   	typeName: jsii.String("typeName"),
//   })
//
type CfnHookTypeConfig interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the activated hook type configuration, in this account and Region.
	AttrConfigurationArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	Configuration() *string
	SetConfiguration(val *string)
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// Defaults to `default` alias. Hook types currently support default configuration alias.
	ConfigurationAlias() *string
	SetConfigurationAlias(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The Amazon Resource Number (ARN) for the hook to set `Configuration` for.
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	TypeArn() *string
	SetTypeArn(val *string)
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	TypeName() *string
	SetTypeName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnHookTypeConfig
type jsiiProxy_CfnHookTypeConfig struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnHookTypeConfig) AttrConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) ConfigurationAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) TypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::HookTypeConfig`.
func NewCfnHookTypeConfig(scope constructs.Construct, id *string, props *CfnHookTypeConfigProps) CfnHookTypeConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnHookTypeConfig{}

	_jsii_.Create(
		"aws-cdk-lib.CfnHookTypeConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::HookTypeConfig`.
func NewCfnHookTypeConfig_Override(c CfnHookTypeConfig, scope constructs.Construct, id *string, props *CfnHookTypeConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnHookTypeConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHookTypeConfig) SetConfiguration(val *string) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnHookTypeConfig) SetConfigurationAlias(val *string) {
	_jsii_.Set(
		j,
		"configurationAlias",
		val,
	)
}

func (j *jsiiProxy_CfnHookTypeConfig) SetTypeArn(val *string) {
	_jsii_.Set(
		j,
		"typeArn",
		val,
	)
}

func (j *jsiiProxy_CfnHookTypeConfig) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHookTypeConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookTypeConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHookTypeConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookTypeConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnHookTypeConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookTypeConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHookTypeConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnHookTypeConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHookTypeConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookTypeConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnHookTypeConfig`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookTypeConfigProps := &cfnHookTypeConfigProps{
//   	configuration: jsii.String("configuration"),
//
//   	// the properties below are optional
//   	configurationAlias: jsii.String("configurationAlias"),
//   	typeArn: jsii.String("typeArn"),
//   	typeName: jsii.String("typeName"),
//   }
//
type CfnHookTypeConfigProps struct {
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	Configuration *string `json:"configuration" yaml:"configuration"`
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// Defaults to `default` alias. Hook types currently support default configuration alias.
	ConfigurationAlias *string `json:"configurationAlias" yaml:"configurationAlias"`
	// The Amazon Resource Number (ARN) for the hook to set `Configuration` for.
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	TypeArn *string `json:"typeArn" yaml:"typeArn"`
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	TypeName *string `json:"typeName" yaml:"typeName"`
}

// A CloudFormation `AWS::CloudFormation::HookVersion`.
//
// The `HookVersion` resource publishes new or first hook version to the AWS CloudFormation registry.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookVersion := cdk.NewCfnHookVersion(this, jsii.String("MyCfnHookVersion"), &cfnHookVersionProps{
//   	schemaHandlerPackage: jsii.String("schemaHandlerPackage"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   })
//
type CfnHookVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Name (ARN) of the hook.
	AttrArn() *string
	// Whether the specified hook version is set as the default version.
	AttrIsDefaultVersion() IResolvable
	// The Amazon Resource Number (ARN) assigned to this version of the hook.
	AttrTypeArn() *string
	// The ID of this version of the hook.
	AttrVersionId() *string
	// The scope at which the resource is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// - `PRIVATE` : The resource is only visible and usable within the account in which it's registered. CloudFormation marks any resources you register as `PRIVATE` .
	// - `PUBLIC` : The resource is publicly visible and usable within any Amazon account.
	AttrVisibility() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the task execution role that grants the hook permission.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// Contains logging configuration information for an extension.
	LoggingConfig() interface{}
	SetLoggingConfig(val interface{})
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A URL to the Amazon S3 bucket containing the hook project package that contains the necessary files for the hook you want to register.
	//
	// For information on generating a schema handler package for the resource you want to register, see [submit](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) in the *CloudFormation CLI User Guide for Extension Development* .
	//
	// > The user registering the resource must be able to access the package in the S3 bucket. That's, the user must have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the schema handler package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	SchemaHandlerPackage() *string
	SetSchemaHandlerPackage(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// > The following organization namespaces are reserved and can't be used in your hook type names:
	// >
	// > - `Alexa`
	// > - `AMZN`
	// > - `Amazon`
	// > - `ASK`
	// > - `AWS`
	// > - `Custom`
	// > - `Dev`.
	TypeName() *string
	SetTypeName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnHookVersion
type jsiiProxy_CfnHookVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnHookVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) AttrIsDefaultVersion() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"attrIsDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) AttrTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) AttrVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) AttrVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) LoggingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) SchemaHandlerPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaHandlerPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::HookVersion`.
func NewCfnHookVersion(scope constructs.Construct, id *string, props *CfnHookVersionProps) CfnHookVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnHookVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnHookVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::HookVersion`.
func NewCfnHookVersion_Override(c CfnHookVersion, scope constructs.Construct, id *string, props *CfnHookVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnHookVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHookVersion) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnHookVersion) SetLoggingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"loggingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnHookVersion) SetSchemaHandlerPackage(val *string) {
	_jsii_.Set(
		j,
		"schemaHandlerPackage",
		val,
	)
}

func (j *jsiiProxy_CfnHookVersion) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHookVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHookVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnHookVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnHookVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHookVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnHookVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHookVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHookVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHookVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHookVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHookVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHookVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHookVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHookVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHookVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHookVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHookVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `LoggingConfig` property type specifies logging configuration information for an extension.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnHookVersion_LoggingConfigProperty struct {
	// The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the extension's handlers.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log entries to CloudWatch Logs.
	LogRoleArn *string `json:"logRoleArn" yaml:"logRoleArn"`
}

// Properties for defining a `CfnHookVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnHookVersionProps := &cfnHookVersionProps{
//   	schemaHandlerPackage: jsii.String("schemaHandlerPackage"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   }
//
type CfnHookVersionProps struct {
	// A URL to the Amazon S3 bucket containing the hook project package that contains the necessary files for the hook you want to register.
	//
	// For information on generating a schema handler package for the resource you want to register, see [submit](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) in the *CloudFormation CLI User Guide for Extension Development* .
	//
	// > The user registering the resource must be able to access the package in the S3 bucket. That's, the user must have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the schema handler package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	SchemaHandlerPackage *string `json:"schemaHandlerPackage" yaml:"schemaHandlerPackage"`
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// > The following organization namespaces are reserved and can't be used in your hook type names:
	// >
	// > - `Alexa`
	// > - `AMZN`
	// > - `Amazon`
	// > - `ASK`
	// > - `AWS`
	// > - `Custom`
	// > - `Dev`.
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the hook permission.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Contains logging configuration information for an extension.
	LoggingConfig interface{} `json:"loggingConfig" yaml:"loggingConfig"`
}

// Captures a synthesis-time JSON object a CloudFormation reference which resolves during deployment to the resolved values of the JSON object.
//
// The main use case for this is to overcome a limitation in CloudFormation that
// does not allow using intrinsic functions as dictionary keys (because
// dictionary keys in JSON must be strings). Specifically this is common in IAM
// conditions such as `StringEquals: { lhs: "rhs" }` where you want "lhs" to be
// a reference.
//
// This object is resolvable, so it can be used as a value.
//
// This construct is backed by a custom resource.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   tagParam := NewCfnParameter(this, jsii.String("TagName"))
//
//   stringEquals := NewCfnJson(this, jsii.String("ConditionJson"), &cfnJsonProps{
//   	value: map[string]*bool{
//   		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
//   	},
//   })
//
//   principal := iam.NewAccountRootPrincipal().withConditions(map[string]interface{}{
//   	"StringEquals": stringEquals,
//   })
//
//   iam.NewRole(this, jsii.String("MyRole"), &roleProps{
//   	assumedBy: principal,
//   })
//
type CfnJson interface {
	constructs.Construct
	IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// The tree node.
	Node() constructs.Node
	// An Fn::GetAtt to the JSON object passed through `value` and resolved during synthesis.
	//
	// Normally there is no need to use this property since `CfnJson` is an
	// IResolvable, so it can be simply used as a value.
	Value() Reference
	// Produce the Token's value at resolution time.
	Resolve(_arg IResolveContext) interface{}
	// This is required in case someone JSON.stringifys an object which refrences this object. Otherwise, we'll get a cyclic JSON reference.
	ToJSON() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnJson
type jsiiProxy_CfnJson struct {
	internal.Type__constructsConstruct
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_CfnJson) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJson) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJson) Value() Reference {
	var returns Reference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewCfnJson(scope constructs.Construct, id *string, props *CfnJsonProps) CfnJson {
	_init_.Initialize()

	j := jsiiProxy_CfnJson{}

	_jsii_.Create(
		"aws-cdk-lib.CfnJson",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnJson_Override(c CfnJson, scope constructs.Construct, id *string, props *CfnJsonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnJson",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnJson_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnJson",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) Resolve(_arg IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJson) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   tagParam := NewCfnParameter(this, jsii.String("TagName"))
//
//   stringEquals := NewCfnJson(this, jsii.String("ConditionJson"), &cfnJsonProps{
//   	value: map[string]*bool{
//   		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
//   	},
//   })
//
//   principal := iam.NewAccountRootPrincipal().withConditions(map[string]interface{}{
//   	"StringEquals": stringEquals,
//   })
//
//   iam.NewRole(this, jsii.String("MyRole"), &roleProps{
//   	assumedBy: principal,
//   })
//
type CfnJsonProps struct {
	// The value to resolve.
	//
	// Can be any JavaScript object, including tokens and
	// references in keys or values.
	Value interface{} `json:"value" yaml:"value"`
}

// A CloudFormation `AWS::CloudFormation::Macro`.
//
// The `AWS::CloudFormation::Macro` resource is a CloudFormation resource type that creates a CloudFormation macro to perform custom processing on CloudFormation templates. For more information, see [Using AWS CloudFormation macros to perform custom processing on templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html) .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnMacro := cdk.NewCfnMacro(this, jsii.String("MyCfnMacro"), &cfnMacroProps{
//   	functionName: jsii.String("functionName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   })
//
type CfnMacro interface {
	CfnResource
	IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the macro.
	Description() *string
	SetDescription(val *string)
	// The Amazon Resource Name (ARN) of the underlying AWS Lambda function that you want AWS CloudFormation to invoke when the macro is run.
	FunctionName() *string
	SetFunctionName(val *string)
	// The CloudWatch Logs group to which AWS CloudFormation sends error logging information when invoking the macro's underlying AWS Lambda function.
	LogGroupName() *string
	SetLogGroupName(val *string)
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
	// The ARN of the role AWS CloudFormation should assume when sending log entries to CloudWatch Logs .
	LogRoleArn() *string
	SetLogRoleArn(val *string)
	// The name of the macro.
	//
	// The name of the macro must be unique across all macros in the account.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnMacro
type jsiiProxy_CfnMacro struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnMacro) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) LogRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacro) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::Macro`.
func NewCfnMacro(scope constructs.Construct, id *string, props *CfnMacroProps) CfnMacro {
	_init_.Initialize()

	j := jsiiProxy_CfnMacro{}

	_jsii_.Create(
		"aws-cdk-lib.CfnMacro",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::Macro`.
func NewCfnMacro_Override(c CfnMacro, scope constructs.Construct, id *string, props *CfnMacroProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnMacro",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMacro) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMacro) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnMacro) SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnMacro) SetLogRoleArn(val *string) {
	_jsii_.Set(
		j,
		"logRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnMacro) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMacro_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMacro",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMacro_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMacro",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnMacro_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMacro",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMacro_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnMacro",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMacro) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMacro) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMacro) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMacro) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMacro) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMacro) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMacro) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMacro) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMacro) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMacro) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMacro) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMacro) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMacro) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMacro) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMacro) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnMacro`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnMacroProps := &cfnMacroProps{
//   	functionName: jsii.String("functionName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnMacroProps struct {
	// The Amazon Resource Name (ARN) of the underlying AWS Lambda function that you want AWS CloudFormation to invoke when the macro is run.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// The name of the macro.
	//
	// The name of the macro must be unique across all macros in the account.
	Name *string `json:"name" yaml:"name"`
	// A description of the macro.
	Description *string `json:"description" yaml:"description"`
	// The CloudWatch Logs group to which AWS CloudFormation sends error logging information when invoking the macro's underlying AWS Lambda function.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role AWS CloudFormation should assume when sending log entries to CloudWatch Logs .
	LogRoleArn *string `json:"logRoleArn" yaml:"logRoleArn"`
}

// Represents a CloudFormation mapping.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   regionTable := NewCfnMapping(this, jsii.String("RegionTable"), &cfnMappingProps{
//   	mapping: map[string]map[string]interface{}{
//   		"us-east-1": map[string]interface{}{
//   			"regionName": jsii.String("US East (N. Virginia)"),
//   		},
//   		"us-east-2": map[string]interface{}{
//   			"regionName": jsii.String("US East (Ohio)"),
//   		},
//   	},
//   })
//
//   regionTable.findInMap(aws_REGION(), jsii.String("regionName"))
//
type CfnMapping interface {
	CfnRefElement
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Returns: A reference to a value in the map based on the two keys.
	FindInMap(key1 *string, key2 *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Sets a value in the map based on the two keys.
	SetValue(key1 *string, key2 *string, value interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnMapping
type jsiiProxy_CfnMapping struct {
	jsiiProxy_CfnRefElement
}

func (j *jsiiProxy_CfnMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMapping) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCfnMapping(scope constructs.Construct, id *string, props *CfnMappingProps) CfnMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnMapping{}

	_jsii_.Create(
		"aws-cdk-lib.CfnMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnMapping_Override(c CfnMapping, scope constructs.Construct, id *string, props *CfnMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMapping) FindInMap(key1 *string, key2 *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"findInMap",
		[]interface{}{key1, key2},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMapping) SetValue(key1 *string, key2 *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"setValue",
		[]interface{}{key1, key2, value},
	)
}

func (c *jsiiProxy_CfnMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   regionTable := NewCfnMapping(this, jsii.String("RegionTable"), &cfnMappingProps{
//   	mapping: map[string]map[string]interface{}{
//   		"us-east-1": map[string]interface{}{
//   			"regionName": jsii.String("US East (N. Virginia)"),
//   		},
//   		"us-east-2": map[string]interface{}{
//   			"regionName": jsii.String("US East (Ohio)"),
//   		},
//   	},
//   })
//
//   regionTable.findInMap(aws_REGION(), jsii.String("regionName"))
//
type CfnMappingProps struct {
	Lazy *bool `json:"lazy" yaml:"lazy"`
	// Mapping of key to a set of corresponding set of named values.
	//
	// The key identifies a map of name-value pairs and must be unique within the mapping.
	//
	// For example, if you want to set values based on a region, you can create a mapping
	// that uses the region name as a key and contains the values you want to specify for
	// each specific region.
	Mapping *map[string]*map[string]interface{} `json:"mapping" yaml:"mapping"`
}

// A CloudFormation `AWS::CloudFormation::ModuleDefaultVersion`.
//
// Specifies the default version of a module. The default version of the module will be used in CloudFormation operations for this account and Region.
//
// To register a module version, use the `[AWS::CloudFormation::ModuleVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html)` resource.
//
// For more information using modules, see [Using modules to encapsulate and reuse resource configurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/modules.html) and [Registering extensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-register) in the *CloudFormation User Guide* . For information on developing modules, see [Developing modules](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/modules.html) in the *CloudFormation CLI User Guide* .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnModuleDefaultVersion := cdk.NewCfnModuleDefaultVersion(this, jsii.String("MyCfnModuleDefaultVersion"), &cfnModuleDefaultVersionProps{
//   	arn: jsii.String("arn"),
//   	moduleName: jsii.String("moduleName"),
//   	versionId: jsii.String("versionId"),
//   })
//
type CfnModuleDefaultVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Name (ARN) of the module version to set as the default version.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	Arn() *string
	SetArn(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	ModuleName() *string
	SetModuleName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The ID for the specific version of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	VersionId() *string
	SetVersionId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnModuleDefaultVersion
type jsiiProxy_CfnModuleDefaultVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnModuleDefaultVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) ModuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::ModuleDefaultVersion`.
func NewCfnModuleDefaultVersion(scope constructs.Construct, id *string, props *CfnModuleDefaultVersionProps) CfnModuleDefaultVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnModuleDefaultVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::ModuleDefaultVersion`.
func NewCfnModuleDefaultVersion_Override(c CfnModuleDefaultVersion, scope constructs.Construct, id *string, props *CfnModuleDefaultVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModuleDefaultVersion) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_CfnModuleDefaultVersion) SetModuleName(val *string) {
	_jsii_.Set(
		j,
		"moduleName",
		val,
	)
}

func (j *jsiiProxy_CfnModuleDefaultVersion) SetVersionId(val *string) {
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnModuleDefaultVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnModuleDefaultVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnModuleDefaultVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModuleDefaultVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModuleDefaultVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnModuleDefaultVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnModuleDefaultVersionProps := &cfnModuleDefaultVersionProps{
//   	arn: jsii.String("arn"),
//   	moduleName: jsii.String("moduleName"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnModuleDefaultVersionProps struct {
	// The Amazon Resource Name (ARN) of the module version to set as the default version.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	Arn *string `json:"arn" yaml:"arn"`
	// The name of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	ModuleName *string `json:"moduleName" yaml:"moduleName"`
	// The ID for the specific version of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	VersionId *string `json:"versionId" yaml:"versionId"`
}

// A CloudFormation `AWS::CloudFormation::ModuleVersion`.
//
// Registers the specified version of the module with the CloudFormation service. Registering a module makes it available for use in CloudFormation templates in your AWS account and Region.
//
// To specify a module version as the default version, use the `[AWS::CloudFormation::ModuleDefaultVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html)` resource.
//
// For more information using modules, see [Using modules to encapsulate and reuse resource configurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/modules.html) and [Registering extensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-register) in the *CloudFormation User Guide* . For information on developing modules, see [Developing modules](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/modules.html) in the *CloudFormation CLI User Guide* .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnModuleVersion := cdk.NewCfnModuleVersion(this, jsii.String("MyCfnModuleVersion"), &cfnModuleVersionProps{
//   	moduleName: jsii.String("moduleName"),
//   	modulePackage: jsii.String("modulePackage"),
//   })
//
type CfnModuleVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Name (ARN) of the module.
	AttrArn() *string
	// The description of the module.
	AttrDescription() *string
	// The URL of a page providing detailed documentation for this module.
	AttrDocumentationUrl() *string
	// Whether the specified module version is set as the default version.
	AttrIsDefaultVersion() IResolvable
	// The schema that defines the module.
	AttrSchema() *string
	// When the specified module version was registered.
	AttrTimeCreated() *string
	// The ID of this version of the module.
	AttrVersionId() *string
	// The scope at which the module is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// - `PRIVATE` : The module is only visible and usable within the account in which it's registered.
	// - `PUBLIC` : The module is publicly visible and usable within any Amazon account.
	AttrVisibility() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of the module being registered.
	ModuleName() *string
	SetModuleName(val *string)
	// A URL to the S3 bucket containing the package that contains the template fragment and schema files for the module version to register.
	//
	// > The user registering the module version must be able to access the module package in the S3 bucket. That's, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	ModulePackage() *string
	SetModulePackage(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnModuleVersion
type jsiiProxy_CfnModuleVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnModuleVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrDocumentationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDocumentationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrIsDefaultVersion() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"attrIsDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrTimeCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTimeCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) AttrVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) ModuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) ModulePackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modulePackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::ModuleVersion`.
func NewCfnModuleVersion(scope constructs.Construct, id *string, props *CfnModuleVersionProps) CfnModuleVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnModuleVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnModuleVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::ModuleVersion`.
func NewCfnModuleVersion_Override(c CfnModuleVersion, scope constructs.Construct, id *string, props *CfnModuleVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnModuleVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModuleVersion) SetModuleName(val *string) {
	_jsii_.Set(
		j,
		"moduleName",
		val,
	)
}

func (j *jsiiProxy_CfnModuleVersion) SetModulePackage(val *string) {
	_jsii_.Set(
		j,
		"modulePackage",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnModuleVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnModuleVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnModuleVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnModuleVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModuleVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnModuleVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModuleVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModuleVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModuleVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModuleVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModuleVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModuleVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModuleVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModuleVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModuleVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModuleVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModuleVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnModuleVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnModuleVersionProps := &cfnModuleVersionProps{
//   	moduleName: jsii.String("moduleName"),
//   	modulePackage: jsii.String("modulePackage"),
//   }
//
type CfnModuleVersionProps struct {
	// The name of the module being registered.
	ModuleName *string `json:"moduleName" yaml:"moduleName"`
	// A URL to the S3 bucket containing the package that contains the template fragment and schema files for the module version to register.
	//
	// > The user registering the module version must be able to access the module package in the S3 bucket. That's, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	ModulePackage *string `json:"modulePackage" yaml:"modulePackage"`
}

// Example:
//   var cluster cluster
//   // add service account
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
//
//   mypod := cluster.addManifest(jsii.String("mypod"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Pod"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("mypod"),
//   	},
//   	"spec": map[string]interface{}{
//   		"serviceAccountName": serviceAccount.serviceAccountName,
//   		"containers": []map[string]interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("hello"),
//   				"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   				"ports": []map[string]*f64{
//   					map[string]*f64{
//   						"containerPort": jsii.Number(8080),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // create the resource after the service account.
//   mypod.node.addDependency(serviceAccount)
//
//   // print the IAM role arn for this service account
//   // print the IAM role arn for this service account
//   NewCfnOutput(this, jsii.String("ServiceAccountIamRole"), &cfnOutputProps{
//   	value: serviceAccount.role.roleArn,
//   })
//
type CfnOutput interface {
	CfnElement
	// A condition to associate with this output value.
	//
	// If the condition evaluates
	// to `false`, this output value will not be included in the stack.
	Condition() CfnCondition
	SetCondition(val CfnCondition)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A String type that describes the output value.
	//
	// The description can be a maximum of 4 K in length.
	Description() *string
	SetDescription(val *string)
	// The name used to export the value of this output across stacks.
	//
	// To use the value in another stack, pass the value of
	// `output.importValue` to it.
	ExportName() *string
	SetExportName(val *string)
	// Return the `Fn.importValue` expression to import this value into another stack.
	//
	// The returned value should not be used in the same stack, but in a
	// different one. It must be deployed to the same environment, as
	// CloudFormation exports can only be imported in the same Region and
	// account.
	//
	// The is no automatic registration of dependencies between stacks when using
	// this mechanism, so you should make sure to deploy them in the right order
	// yourself.
	//
	// You can use this mechanism to share values across Stacks in different
	// Stages. If you intend to share the value to another Stack inside the same
	// Stage, the automatic cross-stack referencing mechanism is more convenient.
	ImportValue() *string
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
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The value of the property returned by the aws cloudformation describe-stacks command.
	//
	// The value of an output can include literals, parameter references, pseudo-parameters,
	// a mapping value, or intrinsic functions.
	Value() interface{}
	SetValue(val interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnOutput
type jsiiProxy_CfnOutput struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnOutput) Condition() CfnCondition {
	var returns CfnCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) ExportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) ImportValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an CfnOutput value for this stack.
func NewCfnOutput(scope constructs.Construct, id *string, props *CfnOutputProps) CfnOutput {
	_init_.Initialize()

	j := jsiiProxy_CfnOutput{}

	_jsii_.Create(
		"aws-cdk-lib.CfnOutput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an CfnOutput value for this stack.
func NewCfnOutput_Override(c CfnOutput, scope constructs.Construct, id *string, props *CfnOutputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnOutput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOutput) SetCondition(val CfnCondition) {
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_CfnOutput) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOutput) SetExportName(val *string) {
	_jsii_.Set(
		j,
		"exportName",
		val,
	)
}

func (j *jsiiProxy_CfnOutput) SetValue(val interface{}) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOutput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnOutput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var cluster cluster
//   // add service account
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
//
//   mypod := cluster.addManifest(jsii.String("mypod"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Pod"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("mypod"),
//   	},
//   	"spec": map[string]interface{}{
//   		"serviceAccountName": serviceAccount.serviceAccountName,
//   		"containers": []map[string]interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("hello"),
//   				"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   				"ports": []map[string]*f64{
//   					map[string]*f64{
//   						"containerPort": jsii.Number(8080),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // create the resource after the service account.
//   mypod.node.addDependency(serviceAccount)
//
//   // print the IAM role arn for this service account
//   // print the IAM role arn for this service account
//   NewCfnOutput(this, jsii.String("ServiceAccountIamRole"), &cfnOutputProps{
//   	value: serviceAccount.role.roleArn,
//   })
//
type CfnOutputProps struct {
	// The value of the property returned by the aws cloudformation describe-stacks command.
	//
	// The value of an output can include literals, parameter references, pseudo-parameters,
	// a mapping value, or intrinsic functions.
	Value *string `json:"value" yaml:"value"`
	// A condition to associate with this output value.
	//
	// If the condition evaluates
	// to `false`, this output value will not be included in the stack.
	Condition CfnCondition `json:"condition" yaml:"condition"`
	// A String type that describes the output value.
	//
	// The description can be a maximum of 4 K in length.
	Description *string `json:"description" yaml:"description"`
	// The name used to export the value of this output across stacks.
	//
	// To import the value from another stack, use `Fn.importValue(exportName)`.
	ExportName *string `json:"exportName" yaml:"exportName"`
}

// A CloudFormation parameter.
//
// Use the optional Parameters section to customize your templates.
// Parameters enable you to input custom values to your template each time you create or
// update a stack.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//   url := NewCfnParameter(this, jsii.String("url-param"))
//
//   myTopic.addSubscription(subscriptions.NewUrlSubscription(url.valueAsString))
//
type CfnParameter interface {
	CfnElement
	// A regular expression that represents the patterns to allow for String types.
	AllowedPattern() *string
	SetAllowedPattern(val *string)
	// An array containing the list of values allowed for the parameter.
	AllowedValues() *[]*string
	SetAllowedValues(val *[]*string)
	// A string that explains a constraint when the constraint is violated.
	//
	// For example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user specifies
	// an invalid value:.
	ConstraintDescription() *string
	SetConstraintDescription(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A value of the appropriate type for the template to use if no value is specified when a stack is created.
	//
	// If you define constraints for the parameter, you must specify
	// a value that adheres to those constraints.
	Default() interface{}
	SetDefault(val interface{})
	// A string of up to 4000 characters that describes the parameter.
	Description() *string
	SetDescription(val *string)
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
	// An integer value that determines the largest number of characters you want to allow for String types.
	MaxLength() *float64
	SetMaxLength(val *float64)
	// A numeric value that determines the largest numeric value you want to allow for Number types.
	MaxValue() *float64
	SetMaxValue(val *float64)
	// An integer value that determines the smallest number of characters you want to allow for String types.
	MinLength() *float64
	SetMinLength(val *float64)
	// A numeric value that determines the smallest numeric value you want to allow for Number types.
	MinValue() *float64
	SetMinValue(val *float64)
	// The tree node.
	Node() constructs.Node
	// Indicates if this parameter is configured with "NoEcho" enabled.
	NoEcho() *bool
	SetNoEcho(val *bool)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The data type for the parameter (DataType).
	Type() *string
	SetType(val *string)
	// The parameter value as a Token.
	Value() IResolvable
	// The parameter value, if it represents a string list.
	ValueAsList() *[]*string
	// The parameter value, if it represents a number.
	ValueAsNumber() *float64
	// The parameter value, if it represents a string.
	ValueAsString() *string
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	Resolve(_context IResolveContext) interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnParameter
type jsiiProxy_CfnParameter struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnParameter) AllowedPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) AllowedValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ConstraintDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraintDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MaxValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) MinValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) NoEcho() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"noEcho",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) Value() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueAsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueAsNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameter) ValueAsString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueAsString",
		&returns,
	)
	return returns
}


// Creates a parameter construct.
//
// Note that the name (logical ID) of the parameter will derive from it's `coname` and location
// within the stack. Therefore, it is recommended that parameters are defined at the stack level.
func NewCfnParameter(scope constructs.Construct, id *string, props *CfnParameterProps) CfnParameter {
	_init_.Initialize()

	j := jsiiProxy_CfnParameter{}

	_jsii_.Create(
		"aws-cdk-lib.CfnParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a parameter construct.
//
// Note that the name (logical ID) of the parameter will derive from it's `coname` and location
// within the stack. Therefore, it is recommended that parameters are defined at the stack level.
func NewCfnParameter_Override(c CfnParameter, scope constructs.Construct, id *string, props *CfnParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnParameter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnParameter) SetAllowedPattern(val *string) {
	_jsii_.Set(
		j,
		"allowedPattern",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetAllowedValues(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedValues",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetConstraintDescription(val *string) {
	_jsii_.Set(
		j,
		"constraintDescription",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetDefault(val interface{}) {
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetMaxLength(val *float64) {
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetMaxValue(val *float64) {
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetMinLength(val *float64) {
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetMinValue(val *float64) {
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetNoEcho(val *bool) {
	_jsii_.Set(
		j,
		"noEcho",
		val,
	)
}

func (j *jsiiProxy_CfnParameter) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnParameter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnParameter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnParameter) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewCfnParameter(this, jsii.String("MyParameter"), &cfnParameterProps{
//   	type: jsii.String("Number"),
//   	default: jsii.Number(1337),
//   })
//
type CfnParameterProps struct {
	// A regular expression that represents the patterns to allow for String types.
	AllowedPattern *string `json:"allowedPattern" yaml:"allowedPattern"`
	// An array containing the list of values allowed for the parameter.
	AllowedValues *[]*string `json:"allowedValues" yaml:"allowedValues"`
	// A string that explains a constraint when the constraint is violated.
	//
	// For example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user specifies
	// an invalid value:.
	ConstraintDescription *string `json:"constraintDescription" yaml:"constraintDescription"`
	// A value of the appropriate type for the template to use if no value is specified when a stack is created.
	//
	// If you define constraints for the parameter, you must specify
	// a value that adheres to those constraints.
	Default interface{} `json:"default" yaml:"default"`
	// A string of up to 4000 characters that describes the parameter.
	Description *string `json:"description" yaml:"description"`
	// An integer value that determines the largest number of characters you want to allow for String types.
	MaxLength *float64 `json:"maxLength" yaml:"maxLength"`
	// A numeric value that determines the largest numeric value you want to allow for Number types.
	MaxValue *float64 `json:"maxValue" yaml:"maxValue"`
	// An integer value that determines the smallest number of characters you want to allow for String types.
	MinLength *float64 `json:"minLength" yaml:"minLength"`
	// A numeric value that determines the smallest numeric value you want to allow for Number types.
	MinValue *float64 `json:"minValue" yaml:"minValue"`
	// Whether to mask the parameter value when anyone makes a call that describes the stack.
	//
	// If you set the value to ``true``, the parameter value is masked with asterisks (``*****``).
	NoEcho *bool `json:"noEcho" yaml:"noEcho"`
	// The data type for the parameter (DataType).
	Type *string `json:"type" yaml:"type"`
}

// A CloudFormation `AWS::CloudFormation::PublicTypeVersion`.
//
// Tests and publishes a registered extension as a public, third-party extension.
//
// CloudFormation first tests the extension to make sure it meets all necessary requirements for being published in the CloudFormation registry. If it does, CloudFormation then publishes it to the registry as a public third-party extension in this Region. Public extensions are available for use by all CloudFormation users.
//
// - For resource types, testing includes passing all contracts tests defined for the type.
// - For modules, testing includes determining if the module's model meets all necessary requirements.
//
// For more information, see [Testing your public extension prior to publishing](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-testing) in the *CloudFormation CLI User Guide* .
//
// If you don't specify a version, CloudFormation uses the default version of the extension in your account and Region for testing.
//
// To perform testing, CloudFormation assumes the execution role specified when the type was registered.
//
// An extension must have a test status of `PASSED` before it can be published. For more information, see [Publishing extensions to make them available for public use](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-publish.html) in the *CloudFormation CLI User Guide* .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnPublicTypeVersion := cdk.NewCfnPublicTypeVersion(this, jsii.String("MyCfnPublicTypeVersion"), &cfnPublicTypeVersionProps{
//   	arn: jsii.String("arn"),
//   	logDeliveryBucket: jsii.String("logDeliveryBucket"),
//   	publicVersionNumber: jsii.String("publicVersionNumber"),
//   	type: jsii.String("type"),
//   	typeName: jsii.String("typeName"),
//   })
//
type CfnPublicTypeVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the extension.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Arn() *string
	SetArn(val *string)
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication.
	AttrPublicTypeArn() *string
	// The publisher ID of the extension publisher.
	AttrPublisherId() *string
	// The Amazon Resource Number (ARN) assigned to this version of the extension.
	AttrTypeVersionArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The S3 bucket to which CloudFormation delivers the contract test execution logs.
	//
	// CloudFormation delivers the logs by the time contract testing has completed and the extension has been assigned a test type status of `PASSED` or `FAILED` .
	//
	// The user initiating the stack operation must be able to access items in the specified S3 bucket. Specifically, the user needs the following permissions:
	//
	// - GetObject
	// - PutObject
	//
	// For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	LogDeliveryBucket() *string
	SetLogDeliveryBucket(val *string)
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
	// The tree node.
	Node() constructs.Node
	// The version number to assign to this version of the extension.
	//
	// Use the following format, and adhere to semantic versioning when assigning a version number to your extension:
	//
	// `MAJOR.MINOR.PATCH`
	//
	// For more information, see [Semantic Versioning 2.0.0](https://docs.aws.amazon.com/https://semver.org/) .
	//
	// If you don't specify a version number, CloudFormation increments the version number by one minor version release.
	//
	// You cannot specify a version number the first time you publish a type. AWS CloudFormation automatically sets the first version number to be `1.0.0` .
	PublicVersionNumber() *string
	SetPublicVersionNumber(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The type of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Type() *string
	SetType(val *string)
	// The name of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	TypeName() *string
	SetTypeName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnPublicTypeVersion
type jsiiProxy_CfnPublicTypeVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnPublicTypeVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrPublicTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrPublisherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublisherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrTypeVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) LogDeliveryBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDeliveryBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) PublicVersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::PublicTypeVersion`.
func NewCfnPublicTypeVersion(scope constructs.Construct, id *string, props *CfnPublicTypeVersionProps) CfnPublicTypeVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnPublicTypeVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnPublicTypeVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::PublicTypeVersion`.
func NewCfnPublicTypeVersion_Override(c CfnPublicTypeVersion, scope constructs.Construct, id *string, props *CfnPublicTypeVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnPublicTypeVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion) SetLogDeliveryBucket(val *string) {
	_jsii_.Set(
		j,
		"logDeliveryBucket",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion) SetPublicVersionNumber(val *string) {
	_jsii_.Set(
		j,
		"publicVersionNumber",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPublicTypeVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublicTypeVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPublicTypeVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublicTypeVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPublicTypeVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublicTypeVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicTypeVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnPublicTypeVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPublicTypeVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnPublicTypeVersionProps := &cfnPublicTypeVersionProps{
//   	arn: jsii.String("arn"),
//   	logDeliveryBucket: jsii.String("logDeliveryBucket"),
//   	publicVersionNumber: jsii.String("publicVersionNumber"),
//   	type: jsii.String("type"),
//   	typeName: jsii.String("typeName"),
//   }
//
type CfnPublicTypeVersionProps struct {
	// The Amazon Resource Number (ARN) of the extension.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Arn *string `json:"arn" yaml:"arn"`
	// The S3 bucket to which CloudFormation delivers the contract test execution logs.
	//
	// CloudFormation delivers the logs by the time contract testing has completed and the extension has been assigned a test type status of `PASSED` or `FAILED` .
	//
	// The user initiating the stack operation must be able to access items in the specified S3 bucket. Specifically, the user needs the following permissions:
	//
	// - GetObject
	// - PutObject
	//
	// For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	LogDeliveryBucket *string `json:"logDeliveryBucket" yaml:"logDeliveryBucket"`
	// The version number to assign to this version of the extension.
	//
	// Use the following format, and adhere to semantic versioning when assigning a version number to your extension:
	//
	// `MAJOR.MINOR.PATCH`
	//
	// For more information, see [Semantic Versioning 2.0.0](https://docs.aws.amazon.com/https://semver.org/) .
	//
	// If you don't specify a version number, CloudFormation increments the version number by one minor version release.
	//
	// You cannot specify a version number the first time you publish a type. AWS CloudFormation automatically sets the first version number to be `1.0.0` .
	PublicVersionNumber *string `json:"publicVersionNumber" yaml:"publicVersionNumber"`
	// The type of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Type *string `json:"type" yaml:"type"`
	// The name of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	TypeName *string `json:"typeName" yaml:"typeName"`
}

// A CloudFormation `AWS::CloudFormation::Publisher`.
//
// Registers your account as a publisher of public extensions in the CloudFormation registry. Public extensions are available for use by all CloudFormation users.
//
// For information on requirements for registering as a public extension publisher, see [Registering your account to publish CloudFormation extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs) in the *CloudFormation CLI User Guide* .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnPublisher := cdk.NewCfnPublisher(this, jsii.String("MyCfnPublisher"), &cfnPublisherProps{
//   	acceptTermsAndConditions: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	connectionArn: jsii.String("connectionArn"),
//   })
//
type CfnPublisher interface {
	CfnResource
	IInspectable
	// Whether you accept the [Terms and Conditions](https://docs.aws.amazon.com/https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to register to publish public extensions to the CloudFormation registry.
	//
	// The default is `false` .
	AcceptTermsAndConditions() interface{}
	SetAcceptTermsAndConditions(val interface{})
	// The type of account used as the identity provider when registering this publisher with CloudFormation .
	//
	// Values include: `AWS_Marketplace` | `Bitbucket` | `GitHub` .
	AttrIdentityProvider() *string
	// The ID of the extension publisher.
	//
	// This publisher ID applies to your account in all AWS Regions .
	AttrPublisherId() *string
	// The URL to the publisher's profile with the identity provider.
	AttrPublisherProfile() *string
	// Whether the publisher is verified.
	AttrPublisherStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	//
	// For more information, see [Registering your account to publish CloudFormation extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs) in the *CloudFormation CLI User Guide* .
	ConnectionArn() *string
	SetConnectionArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnPublisher
type jsiiProxy_CfnPublisher struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnPublisher) AcceptTermsAndConditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptTermsAndConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) AttrIdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) AttrPublisherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublisherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) AttrPublisherProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublisherProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) AttrPublisherStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublisherStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublisher) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::Publisher`.
func NewCfnPublisher(scope constructs.Construct, id *string, props *CfnPublisherProps) CfnPublisher {
	_init_.Initialize()

	j := jsiiProxy_CfnPublisher{}

	_jsii_.Create(
		"aws-cdk-lib.CfnPublisher",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::Publisher`.
func NewCfnPublisher_Override(c CfnPublisher, scope constructs.Construct, id *string, props *CfnPublisherProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnPublisher",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublisher) SetAcceptTermsAndConditions(val interface{}) {
	_jsii_.Set(
		j,
		"acceptTermsAndConditions",
		val,
	)
}

func (j *jsiiProxy_CfnPublisher) SetConnectionArn(val *string) {
	_jsii_.Set(
		j,
		"connectionArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPublisher_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublisher",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPublisher_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublisher",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPublisher_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnPublisher",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublisher_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnPublisher",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublisher) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPublisher) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPublisher) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPublisher) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPublisher) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPublisher) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPublisher) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPublisher) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublisher) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublisher) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPublisher) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublisher) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublisher) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublisher) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublisher) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPublisher`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnPublisherProps := &cfnPublisherProps{
//   	acceptTermsAndConditions: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	connectionArn: jsii.String("connectionArn"),
//   }
//
type CfnPublisherProps struct {
	// Whether you accept the [Terms and Conditions](https://docs.aws.amazon.com/https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to register to publish public extensions to the CloudFormation registry.
	//
	// The default is `false` .
	AcceptTermsAndConditions interface{} `json:"acceptTermsAndConditions" yaml:"acceptTermsAndConditions"`
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	//
	// For more information, see [Registering your account to publish CloudFormation extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs) in the *CloudFormation CLI User Guide* .
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
}

// Base class for referenceable CloudFormation constructs which are not Resources.
//
// These constructs are things like Conditions and Parameters, can be
// referenced by taking the `.ref` attribute.
//
// Resource constructs do not inherit from CfnRefElement because they have their
// own, more specific types returned from the .ref attribute. Also, some
// resources aren't referenceable at all (such as BucketPolicies or GatewayAttachments).
type CfnRefElement interface {
	CfnElement
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnRefElement
type jsiiProxy_CfnRefElement struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnRefElement) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRefElement) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRefElement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRefElement) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRefElement) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates an entity and binds it to a tree.
//
// Note that the root of the tree must be a Stack object (not just any Root).
func NewCfnRefElement_Override(c CfnRefElement, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnRefElement",
		[]interface{}{scope, id},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRefElement_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnRefElement",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnRefElement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnRefElement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRefElement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRefElement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a CloudFormation resource.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var bucket bucket
//
//
//   cfnBucket := bucket.node.findChild(jsii.String("Resource")).(cfnResource)
//   cfnBucket.applyRemovalPolicy(removalPolicy_DESTROY)
//
type CfnResource interface {
	CfnRefElement
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnResource
type jsiiProxy_CfnResource struct {
	jsiiProxy_CfnRefElement
}

func (j *jsiiProxy_CfnResource) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Creates a resource construct.
func NewCfnResource(scope constructs.Construct, id *string, props *CfnResourceProps) CfnResource {
	_init_.Initialize()

	j := jsiiProxy_CfnResource{}

	_jsii_.Create(
		"aws-cdk-lib.CfnResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a resource construct.
func NewCfnResource_Override(c CfnResource, scope constructs.Construct, id *string, props *CfnResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnResource",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResource) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResource) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResource) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// For an Auto Scaling group replacement update, specifies how many instances must signal success for the update to succeed.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceAutoScalingCreationPolicy := &cfnResourceAutoScalingCreationPolicy{
//   	minSuccessfulInstancesPercent: jsii.Number(123),
//   }
//
type CfnResourceAutoScalingCreationPolicy struct {
	// Specifies the percentage of instances in an Auto Scaling replacement update that must signal success for the update to succeed.
	//
	// You can specify a value from 0 to 100. AWS CloudFormation rounds to the nearest tenth of a percent.
	// For example, if you update five instances with a minimum successful percentage of 50, three instances must signal success.
	// If an instance doesn't send a signal within the time specified by the Timeout property, AWS CloudFormation assumes that the
	// instance wasn't created.
	MinSuccessfulInstancesPercent *float64 `json:"minSuccessfulInstancesPercent" yaml:"minSuccessfulInstancesPercent"`
}

// A CloudFormation `AWS::CloudFormation::ResourceDefaultVersion`.
//
// Specifies the default version of a resource. The default version of a resource will be used in CloudFormation operations.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceDefaultVersion := cdk.NewCfnResourceDefaultVersion(this, jsii.String("MyCfnResourceDefaultVersion"), &cfnResourceDefaultVersionProps{
//   	typeName: jsii.String("typeName"),
//   	typeVersionArn: jsii.String("typeVersionArn"),
//   	versionId: jsii.String("versionId"),
//   })
//
type CfnResourceDefaultVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Name (ARN) of the resource.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The name of the resource.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName() *string
	SetTypeName(val *string)
	// The Amazon Resource Name (ARN) of the resource version.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn() *string
	SetTypeVersionArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The ID of a specific version of the resource.
	//
	// The version ID is the value at the end of the Amazon Resource Name (ARN) assigned to the resource version when it's registered.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId() *string
	SetVersionId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnResourceDefaultVersion
type jsiiProxy_CfnResourceDefaultVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnResourceDefaultVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) TypeVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::ResourceDefaultVersion`.
func NewCfnResourceDefaultVersion(scope constructs.Construct, id *string, props *CfnResourceDefaultVersionProps) CfnResourceDefaultVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceDefaultVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::ResourceDefaultVersion`.
func NewCfnResourceDefaultVersion_Override(c CfnResourceDefaultVersion, scope constructs.Construct, id *string, props *CfnResourceDefaultVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceDefaultVersion) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDefaultVersion) SetTypeVersionArn(val *string) {
	_jsii_.Set(
		j,
		"typeVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDefaultVersion) SetVersionId(val *string) {
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceDefaultVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceDefaultVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnResourceDefaultVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDefaultVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourceDefaultVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceDefaultVersionProps := &cfnResourceDefaultVersionProps{
//   	typeName: jsii.String("typeName"),
//   	typeVersionArn: jsii.String("typeVersionArn"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnResourceDefaultVersionProps struct {
	// The name of the resource.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the resource version.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn *string `json:"typeVersionArn" yaml:"typeVersionArn"`
	// The ID of a specific version of the resource.
	//
	// The version ID is the value at the end of the Amazon Resource Name (ARN) assigned to the resource version when it's registered.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId *string `json:"versionId" yaml:"versionId"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   resourceA := NewCfnResource(this, jsii.String("ResourceA"), resourceProps)
//   resourceB := NewCfnResource(this, jsii.String("ResourceB"), resourceProps)
//
//   resourceB.addDependsOn(resourceA)
//
type CfnResourceProps struct {
	// CloudFormation resource type (e.g. `AWS::S3::Bucket`).
	Type *string `json:"type" yaml:"type"`
	// Resource properties.
	Properties *map[string]interface{} `json:"properties" yaml:"properties"`
}

// When AWS CloudFormation creates the associated resource, configures the number of required success signals and the length of time that AWS CloudFormation waits for those signals.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceSignal := &cfnResourceSignal{
//   	count: jsii.Number(123),
//   	timeout: jsii.String("timeout"),
//   }
//
type CfnResourceSignal struct {
	// The number of success signals AWS CloudFormation must receive before it sets the resource status as CREATE_COMPLETE.
	//
	// If the resource receives a failure signal or doesn't receive the specified number of signals before the timeout period
	// expires, the resource creation fails and AWS CloudFormation rolls the stack back.
	Count *float64 `json:"count" yaml:"count"`
	// The length of time that AWS CloudFormation waits for the number of signals that was specified in the Count property.
	//
	// The timeout period starts after AWS CloudFormation starts creating the resource, and the timeout expires no sooner
	// than the time you specify but can occur shortly thereafter. The maximum time that you can specify is 12 hours.
	Timeout *string `json:"timeout" yaml:"timeout"`
}

// A CloudFormation `AWS::CloudFormation::ResourceVersion`.
//
// Registers a resource version with the CloudFormation service. Registering a resource version makes it available for use in CloudFormation templates in your AWS account , and includes:
//
// - Validating the resource schema.
// - Determining which handlers, if any, have been specified for the resource.
// - Making the resource available for use in your account.
//
// For more information on how to develop resources and ready them for registration, see [Creating Resource Providers](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html) in the *CloudFormation CLI User Guide* .
//
// You can have a maximum of 50 resource versions registered at a time. This maximum is per account and per Region.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceVersion := cdk.NewCfnResourceVersion(this, jsii.String("MyCfnResourceVersion"), &cfnResourceVersionProps{
//   	schemaHandlerPackage: jsii.String("schemaHandlerPackage"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   })
//
type CfnResourceVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Name (ARN) of the resource version.
	AttrArn() *string
	// Whether the resource version is set as the default version.
	AttrIsDefaultVersion() IResolvable
	// The provisioning behavior of the resource type.
	//
	// CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.
	//
	// Valid values include:
	//
	// - `FULLY_MUTABLE` : The resource type includes an update handler to process updates to the type during stack update operations.
	// - `IMMUTABLE` : The resource type doesn't include an update handler, so the type can't be updated and must instead be replaced during stack update operations.
	// - `NON_PROVISIONABLE` : The resource type doesn't include all the following handlers, and therefore can't actually be provisioned.
	//
	// - create
	// - read
	// - delete.
	AttrProvisioningType() *string
	// The Amazon Resource Name (ARN) of the resource.
	AttrTypeArn() *string
	// The ID of a specific version of the resource.
	//
	// The version ID is the value at the end of the Amazon Resource Name (ARN) assigned to the resource version when it is registered.
	AttrVersionId() *string
	// The scope at which the resource is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// - `PRIVATE` : The resource is only visible and usable within the account in which it's registered. CloudFormation marks any resources you register as `PRIVATE` .
	// - `PUBLIC` : The resource is publicly visible and usable within any Amazon account.
	AttrVisibility() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume when invoking the resource.
	//
	// If your resource calls AWS APIs in any of its handlers, you must create an *[IAM execution role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)* that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the resource type handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the resource type handler, thereby supplying your resource type with the appropriate credentials.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// Logging configuration information for a resource.
	LoggingConfig() interface{}
	SetLoggingConfig(val interface{})
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A URL to the S3 bucket containing the resource project package that contains the necessary files for the resource you want to register.
	//
	// For information on generating a schema handler package for the resource you want to register, see [submit](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) in the *CloudFormation CLI User Guide* .
	//
	// > The user registering the resource must be able to access the package in the S3 bucket. That is, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the schema handler package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	SchemaHandlerPackage() *string
	SetSchemaHandlerPackage(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The name of the resource being registered.
	//
	// We recommend that resource names adhere to the following pattern: *company_or_organization* :: *service* :: *type* .
	//
	// > The following organization namespaces are reserved and can't be used in your resource names:
	// >
	// > - `Alexa`
	// > - `AMZN`
	// > - `Amazon`
	// > - `AWS`
	// > - `Custom`
	// > - `Dev`.
	TypeName() *string
	SetTypeName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnResourceVersion
type jsiiProxy_CfnResourceVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnResourceVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) AttrIsDefaultVersion() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"attrIsDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) AttrProvisioningType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProvisioningType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) AttrTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) AttrVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) AttrVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) LoggingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) SchemaHandlerPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaHandlerPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::ResourceVersion`.
func NewCfnResourceVersion(scope constructs.Construct, id *string, props *CfnResourceVersionProps) CfnResourceVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceVersion{}

	_jsii_.Create(
		"aws-cdk-lib.CfnResourceVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::ResourceVersion`.
func NewCfnResourceVersion_Override(c CfnResourceVersion, scope constructs.Construct, id *string, props *CfnResourceVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnResourceVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceVersion) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnResourceVersion) SetLoggingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"loggingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResourceVersion) SetSchemaHandlerPackage(val *string) {
	_jsii_.Set(
		j,
		"schemaHandlerPackage",
		val,
	)
}

func (j *jsiiProxy_CfnResourceVersion) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnResourceVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnResourceVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceVersion) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceVersion) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceVersion) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Logging configuration information for a resource.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnResourceVersion_LoggingConfigProperty struct {
	// The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.
	LogRoleArn *string `json:"logRoleArn" yaml:"logRoleArn"`
}

// Properties for defining a `CfnResourceVersion`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnResourceVersionProps := &cfnResourceVersionProps{
//   	schemaHandlerPackage: jsii.String("schemaHandlerPackage"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   }
//
type CfnResourceVersionProps struct {
	// A URL to the S3 bucket containing the resource project package that contains the necessary files for the resource you want to register.
	//
	// For information on generating a schema handler package for the resource you want to register, see [submit](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) in the *CloudFormation CLI User Guide* .
	//
	// > The user registering the resource must be able to access the package in the S3 bucket. That is, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the schema handler package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	SchemaHandlerPackage *string `json:"schemaHandlerPackage" yaml:"schemaHandlerPackage"`
	// The name of the resource being registered.
	//
	// We recommend that resource names adhere to the following pattern: *company_or_organization* :: *service* :: *type* .
	//
	// > The following organization namespaces are reserved and can't be used in your resource names:
	// >
	// > - `Alexa`
	// > - `AMZN`
	// > - `Amazon`
	// > - `AWS`
	// > - `Custom`
	// > - `Dev`.
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume when invoking the resource.
	//
	// If your resource calls AWS APIs in any of its handlers, you must create an *[IAM execution role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)* that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the resource type handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the resource type handler, thereby supplying your resource type with the appropriate credentials.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Logging configuration information for a resource.
	LoggingConfig interface{} `json:"loggingConfig" yaml:"loggingConfig"`
}

// The Rules that define template constraints in an AWS Service Catalog portfolio describe when end users can use the template and which values they can specify for parameters that are declared in the AWS CloudFormation template used to create the product they are attempting to use.
//
// Rules
// are useful for preventing end users from inadvertently specifying an incorrect value.
// For example, you can add a rule to verify whether end users specified a valid subnet in a
// given VPC or used m1.small instance types for test environments. AWS CloudFormation uses
// rules to validate parameter values before it creates the resources for the product.
//
// A rule can include a RuleCondition property and must include an Assertions property.
// For each rule, you can define only one rule condition; you can define one or more asserts within the Assertions property.
// You define a rule condition and assertions by using rule-specific intrinsic functions.
//
// Example:
//   var cfnTemplate cfnInclude
//
//   // mutating the rule
//   var myParameter cfnParameter
//   rule := cfnTemplate.getRule(jsii.String("MyRule"))
//   rule.addAssertion(core.fn.conditionContains([]*string{
//   	jsii.String("m1.small"),
//   }, myParameter.valueAsString), jsii.String("MyParameter has to be m1.small"))
//
type CfnRule interface {
	CfnRefElement
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Adds an assertion to the rule.
	AddAssertion(condition ICfnConditionExpression, description *string)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnRule
type jsiiProxy_CfnRule struct {
	jsiiProxy_CfnRefElement
}

func (j *jsiiProxy_CfnRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates and adds a rule.
func NewCfnRule(scope constructs.Construct, id *string, props *CfnRuleProps) CfnRule {
	_init_.Initialize()

	j := jsiiProxy_CfnRule{}

	_jsii_.Create(
		"aws-cdk-lib.CfnRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates and adds a rule.
func NewCfnRule_Override(c CfnRule, scope constructs.Construct, id *string, props *CfnRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnRule",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRule) AddAssertion(condition ICfnConditionExpression, description *string) {
	_jsii_.InvokeVoid(
		c,
		"addAssertion",
		[]interface{}{condition, description},
	)
}

func (c *jsiiProxy_CfnRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A rule assertion.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//   cfnRuleAssertion := &cfnRuleAssertion{
//   	assert: cfnConditionExpression,
//   	assertDescription: jsii.String("assertDescription"),
//   }
//
type CfnRuleAssertion struct {
	// The assertion.
	Assert ICfnConditionExpression `json:"assert" yaml:"assert"`
	// The assertion description.
	AssertDescription *string `json:"assertDescription" yaml:"assertDescription"`
}

// A rule can include a RuleCondition property and must include an Assertions property.
//
// For each rule, you can define only one rule condition; you can define one or more asserts within the Assertions property.
// You define a rule condition and assertions by using rule-specific intrinsic functions.
//
// You can use the following rule-specific intrinsic functions to define rule conditions and assertions:
//
//   Fn::And
//   Fn::Contains
//   Fn::EachMemberEquals
//   Fn::EachMemberIn
//   Fn::Equals
//   Fn::If
//   Fn::Not
//   Fn::Or
//   Fn::RefAll
//   Fn::ValueOf
//   Fn::ValueOfAll
//
// https://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//   cfnRuleProps := &cfnRuleProps{
//   	assertions: []cfnRuleAssertion{
//   		&cfnRuleAssertion{
//   			assert: cfnConditionExpression,
//   			assertDescription: jsii.String("assertDescription"),
//   		},
//   	},
//   	ruleCondition: cfnConditionExpression,
//   }
//
type CfnRuleProps struct {
	// Assertions which define the rule.
	Assertions *[]*CfnRuleAssertion `json:"assertions" yaml:"assertions"`
	// If the rule condition evaluates to false, the rule doesn't take effect.
	//
	// If the function in the rule condition evaluates to true, expressions in each assert are evaluated and applied.
	RuleCondition ICfnConditionExpression `json:"ruleCondition" yaml:"ruleCondition"`
}

// A CloudFormation `AWS::CloudFormation::Stack`.
//
// The `AWS::CloudFormation::Stack` resource nests a stack as a resource in a top-level template.
//
// You can add output values from a nested stack within the containing template. You use the [GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html) function with the nested stack's logical name and the name of the output value in the nested stack in the format `Outputs. *NestedStackOutputName*` .
//
// > We strongly recommend that updates to nested stacks are run from the parent stack.
//
// When you apply template changes to update a top-level stack, CloudFormation updates the top-level stack and initiates an update to its nested stacks. CloudFormation updates the resources of modified nested stacks, but doesn't update the resources of unmodified nested stacks. For more information, see [CloudFormation stack updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// > You must acknowledge IAM capabilities for nested stacks that contain IAM resources. Also, verify that you have cancel update stack permissions, which is required if an update rolls back. For more information about IAM and CloudFormation , see [Controlling access with AWS Identity and Access Management](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html) .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnStack := cdk.NewCfnStack(this, jsii.String("MyCfnStack"), &cfnStackProps{
//   	templateUrl: jsii.String("templateUrl"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   })
//
type CfnStack interface {
	CfnResource
	IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The Amazon Simple Notification Service (Amazon SNS) topic ARNs to publish stack related events.
	//
	// You can find your Amazon SNS topic ARNs using the Amazon SNS console or your Command Line Interface (CLI).
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding to a parameter defined in the embedded template and a value representing the value that you want to set for the parameter.
	//
	// > If you use the `Ref` function to pass a parameter value to a nested stack, comma-delimited list parameters must be of type `String` . In other words, you can't pass values that are of type `CommaDelimitedList` to nested stacks.
	//
	// Conditional. Required if the nested stack requires input parameters.
	//
	// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
	Parameters() interface{}
	SetParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Key-value pairs to associate with this stack.
	//
	// AWS CloudFormation also propagates these tags to the resources created in the stack. A maximum number of 50 tags can be specified.
	Tags() TagManager
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket. For more information, see [Template anatomy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html) .
	//
	// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
	TemplateUrl() *string
	SetTemplateUrl(val *string)
	// The length of time, in minutes, that CloudFormation waits for the nested stack to reach the `CREATE_COMPLETE` state.
	//
	// The default is no timeout. When CloudFormation detects that the nested stack has reached the `CREATE_COMPLETE` state, it marks the nested stack resource as `CREATE_COMPLETE` in the parent stack and resumes creating the parent stack. If the timeout period expires before the nested stack reaches `CREATE_COMPLETE` , CloudFormation marks the nested stack as failed and rolls back both the nested stack and parent stack.
	//
	// Updates aren't supported.
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStack
type jsiiProxy_CfnStack struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnStack) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::Stack`.
func NewCfnStack(scope constructs.Construct, id *string, props *CfnStackProps) CfnStack {
	_init_.Initialize()

	j := jsiiProxy_CfnStack{}

	_jsii_.Create(
		"aws-cdk-lib.CfnStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::Stack`.
func NewCfnStack_Override(c CfnStack, scope constructs.Construct, id *string, props *CfnStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnStack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStack) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetTemplateUrl(val *string) {
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

func (j *jsiiProxy_CfnStack) SetTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStack_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStack",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnStack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStack) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStack) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStack) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStack) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStack) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStack`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnStackProps := &cfnStackProps{
//   	templateUrl: jsii.String("templateUrl"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   }
//
type CfnStackProps struct {
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket. For more information, see [Template anatomy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html) .
	//
	// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
	TemplateUrl *string `json:"templateUrl" yaml:"templateUrl"`
	// The Amazon Simple Notification Service (Amazon SNS) topic ARNs to publish stack related events.
	//
	// You can find your Amazon SNS topic ARNs using the Amazon SNS console or your Command Line Interface (CLI).
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding to a parameter defined in the embedded template and a value representing the value that you want to set for the parameter.
	//
	// > If you use the `Ref` function to pass a parameter value to a nested stack, comma-delimited list parameters must be of type `String` . In other words, you can't pass values that are of type `CommaDelimitedList` to nested stacks.
	//
	// Conditional. Required if the nested stack requires input parameters.
	//
	// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Key-value pairs to associate with this stack.
	//
	// AWS CloudFormation also propagates these tags to the resources created in the stack. A maximum number of 50 tags can be specified.
	Tags *[]*CfnTag `json:"tags" yaml:"tags"`
	// The length of time, in minutes, that CloudFormation waits for the nested stack to reach the `CREATE_COMPLETE` state.
	//
	// The default is no timeout. When CloudFormation detects that the nested stack has reached the `CREATE_COMPLETE` state, it marks the nested stack resource as `CREATE_COMPLETE` in the parent stack and resumes creating the parent stack. If the timeout period expires before the nested stack reaches `CREATE_COMPLETE` , CloudFormation marks the nested stack as failed and rolls back both the nested stack and parent stack.
	//
	// Updates aren't supported.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

// A CloudFormation `AWS::CloudFormation::StackSet`.
//
// The `AWS::CloudFormation::StackSet` enables you to provision stacks into AWS accounts and across Regions by using a single CloudFormation template. In the stack set, you specify the template to use, in addition to any parameters and capabilities that the template requires.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedExecution interface{}
//   cfnStackSet := cdk.NewCfnStackSet(this, jsii.String("MyCfnStackSet"), &cfnStackSetProps{
//   	permissionModel: jsii.String("permissionModel"),
//   	stackSetName: jsii.String("stackSetName"),
//
//   	// the properties below are optional
//   	administrationRoleArn: jsii.String("administrationRoleArn"),
//   	autoDeployment: &autoDeploymentProperty{
//   		enabled: jsii.Boolean(false),
//   		retainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	callAs: jsii.String("callAs"),
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	executionRoleName: jsii.String("executionRoleName"),
//   	managedExecution: managedExecution,
//   	operationPreferences: &operationPreferencesProperty{
//   		failureToleranceCount: jsii.Number(123),
//   		failureTolerancePercentage: jsii.Number(123),
//   		maxConcurrentCount: jsii.Number(123),
//   		maxConcurrentPercentage: jsii.Number(123),
//   		regionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		regionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	stackInstancesGroup: []interface{}{
//   		&stackInstancesProperty{
//   			deploymentTargets: &deploymentTargetsProperty{
//   				accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				organizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			parameterOverrides: []interface{}{
//   				&parameterProperty{
//   					parameterKey: jsii.String("parameterKey"),
//   					parameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateBody: jsii.String("templateBody"),
//   	templateUrl: jsii.String("templateUrl"),
//   })
//
type CfnStackSet interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	//
	// Use customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account. For more information, see [Prerequisites: Granting Permissions for Stack Set Operations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) in the *AWS CloudFormation User Guide* .
	//
	// *Minimum* : `20`
	//
	// *Maximum* : `2048`.
	AdministrationRoleArn() *string
	SetAdministrationRoleArn(val *string)
	// The ID of the stack that you're creating.
	AttrStackSetId() *string
	// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
	AutoDeployment() interface{}
	SetAutoDeployment(val interface{})
	// [Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.
	//
	// By default, `SELF` is specified. Use `SELF` for stack sets with self-managed permissions.
	//
	// - To create a stack set with service-managed permissions while signed in to the management account, specify `SELF` .
	// - To create a stack set with service-managed permissions while signed in to a delegated administrator account, specify `DELEGATED_ADMIN` .
	//
	// Your AWS account must be registered as a delegated admin in the management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the *AWS CloudFormation User Guide* .
	//
	// Stack sets with service-managed permissions are created in the management account, including stack sets that are created by delegated administrators.
	//
	// *Valid Values* : `SELF` | `DELEGATED_ADMIN`.
	CallAs() *string
	SetCallAs(val *string)
	// The capabilities that are allowed in the stack set.
	//
	// Some stack set templates might include resources that can affect permissions in your AWS account —for example, by creating new AWS Identity and Access Management ( IAM ) users. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities) .
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the stack set.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	Description() *string
	SetDescription(val *string)
	// The name of the IAM execution role to use to create the stack set.
	//
	// If you don't specify an execution role, AWS CloudFormation uses the `AWSCloudFormationStackSetExecutionRole` role for the stack set operation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `[a-zA-Z_0-9+=,.@-]+`
	ExecutionRoleName() *string
	SetExecutionRoleName(val *string)
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
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// When active, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// > If there are already running or queued operations, StackSets queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your stack set's execution configuration while there are running or queued operations for that stack set.
	//
	// When inactive (default), StackSets performs one operation at a time in request order.
	ManagedExecution() interface{}
	SetManagedExecution(val interface{})
	// The tree node.
	Node() constructs.Node
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences() interface{}
	SetOperationPreferences(val interface{})
	// The input parameters for the stack set template.
	Parameters() interface{}
	SetParameters(val interface{})
	// Describes how the IAM roles required for stack set operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant Self-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations . For more information, see [Grant Service-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html) .
	//
	// *Allowed Values* : `SERVICE_MANAGED` | `SELF_MANAGED`
	//
	// > The `PermissionModel` property is required.
	PermissionModel() *string
	SetPermissionModel(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// A group of stack instances with parameters in some specific accounts and Regions.
	StackInstancesGroup() interface{}
	SetStackInstancesGroup(val interface{})
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// *Maximum* : `128`
	//
	// *Pattern* : `^[a-zA-Z][a-zA-Z0-9-]{0,127}$`
	//
	// > The `StackSetName` property is required.
	StackSetName() *string
	SetStackSetName(val *string)
	// The key-value pairs to associate with this stack set and the stacks created from it.
	//
	// AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	Tags() TagManager
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates containing dynamic references through `TemplateUrl` instead.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `51200`.
	TemplateBody() *string
	SetTemplateBody(val *string)
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	TemplateUrl() *string
	SetTemplateUrl(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStackSet
type jsiiProxy_CfnStackSet struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnStackSet) AdministrationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) AttrStackSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStackSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) AutoDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CallAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) ExecutionRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) ManagedExecution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) OperationPreferences() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) PermissionModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) StackInstancesGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackInstancesGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::StackSet`.
func NewCfnStackSet(scope constructs.Construct, id *string, props *CfnStackSetProps) CfnStackSet {
	_init_.Initialize()

	j := jsiiProxy_CfnStackSet{}

	_jsii_.Create(
		"aws-cdk-lib.CfnStackSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::StackSet`.
func NewCfnStackSet_Override(c CfnStackSet, scope constructs.Construct, id *string, props *CfnStackSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnStackSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStackSet) SetAdministrationRoleArn(val *string) {
	_jsii_.Set(
		j,
		"administrationRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetAutoDeployment(val interface{}) {
	_jsii_.Set(
		j,
		"autoDeployment",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetCallAs(val *string) {
	_jsii_.Set(
		j,
		"callAs",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetExecutionRoleName(val *string) {
	_jsii_.Set(
		j,
		"executionRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetManagedExecution(val interface{}) {
	_jsii_.Set(
		j,
		"managedExecution",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetOperationPreferences(val interface{}) {
	_jsii_.Set(
		j,
		"operationPreferences",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetPermissionModel(val *string) {
	_jsii_.Set(
		j,
		"permissionModel",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetStackInstancesGroup(val interface{}) {
	_jsii_.Set(
		j,
		"stackInstancesGroup",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetStackSetName(val *string) {
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet) SetTemplateUrl(val *string) {
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStackSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStackSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStackSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStackSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStackSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnStackSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnStackSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStackSet) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStackSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStackSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStackSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStackSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStackSet) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStackSet) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStackSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStackSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organizational unit (OU).
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   autoDeploymentProperty := &autoDeploymentProperty{
//   	enabled: jsii.Boolean(false),
//   	retainStacksOnAccountRemoval: jsii.Boolean(false),
//   }
//
type CfnStackSet_AutoDeploymentProperty struct {
	// If set to `true` , StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// If set to `true` , stack resources are retained when an account is removed from a target organization or OU.
	//
	// If set to `false` , stack resources are deleted. Specify only if `Enabled` is set to `True` .
	RetainStacksOnAccountRemoval interface{} `json:"retainStacksOnAccountRemoval" yaml:"retainStacksOnAccountRemoval"`
}

// The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   deploymentTargetsProperty := &deploymentTargetsProperty{
//   	accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	organizationalUnitIds: []*string{
//   		jsii.String("organizationalUnitIds"),
//   	},
//   }
//
type CfnStackSet_DeploymentTargetsProperty struct {
	// The names of one or more AWS accounts for which you want to deploy stack set updates.
	//
	// *Pattern* : `^[0-9]{12}$`.
	Accounts *[]*string `json:"accounts" yaml:"accounts"`
	// The organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	//
	// *Pattern* : `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`.
	OrganizationalUnitIds *[]*string `json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
//
// For more information on maximum concurrent accounts and failure tolerance, see [Stack set operation options](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options) .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   operationPreferencesProperty := &operationPreferencesProperty{
//   	failureToleranceCount: jsii.Number(123),
//   	failureTolerancePercentage: jsii.Number(123),
//   	maxConcurrentCount: jsii.Number(123),
//   	maxConcurrentPercentage: jsii.Number(123),
//   	regionConcurrencyType: jsii.String("regionConcurrencyType"),
//   	regionOrder: []*string{
//   		jsii.String("regionOrder"),
//   	},
//   }
//
type CfnStackSet_OperationPreferencesProperty struct {
	// The number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in any subsequent Regions.
	//
	// Conditional: You must specify either `FailureToleranceCount` or `FailureTolerancePercentage` (but not both).
	FailureToleranceCount *float64 `json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// The percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in any subsequent Regions.
	//
	// When calculating the number of accounts based on the specified percentage, AWS CloudFormation rounds *down* to the next whole number.
	//
	// Conditional: You must specify either `FailureToleranceCount` or `FailureTolerancePercentage` , but not both.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum number of accounts in which to perform this operation at one time.
	//
	// This is dependent on the value of `FailureToleranceCount` . `MaxConcurrentCount` is at most one more than the `FailureToleranceCount` .
	//
	// Note that this setting lets you specify the *maximum* for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Conditional: You must specify either `MaxConcurrentCount` or `MaxConcurrentPercentage` , but not both.
	MaxConcurrentCount *float64 `json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, CloudFormation sets the number as one instead.
	//
	// Note that this setting lets you specify the *maximum* for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Conditional: You must specify either `MaxConcurrentCount` or `MaxConcurrentPercentage` , but not both.
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	//
	// *Allowed values* : `SEQUENTIAL` | `PARALLEL`.
	RegionConcurrencyType *string `json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// The order of the Regions where you want to perform the stack operation.
	RegionOrder *[]*string `json:"regionOrder" yaml:"regionOrder"`
}

// The Parameter data type.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   parameterProperty := &parameterProperty{
//   	parameterKey: jsii.String("parameterKey"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnStackSet_ParameterProperty struct {
	// The key associated with the parameter.
	//
	// If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that's specified in your template.
	ParameterKey *string `json:"parameterKey" yaml:"parameterKey"`
	// The input value associated with the parameter.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Stack instances in some specific accounts and Regions.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   stackInstancesProperty := &stackInstancesProperty{
//   	deploymentTargets: &deploymentTargetsProperty{
//   		accounts: []*string{
//   			jsii.String("accounts"),
//   		},
//   		organizationalUnitIds: []*string{
//   			jsii.String("organizationalUnitIds"),
//   		},
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//
//   	// the properties below are optional
//   	parameterOverrides: []interface{}{
//   		&parameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
type CfnStackSet_StackInstancesProperty struct {
	// The AWS `OrganizationalUnitIds` or `Accounts` for which to create stack instances in the specified Regions.
	DeploymentTargets interface{} `json:"deploymentTargets" yaml:"deploymentTargets"`
	// The names of one or more Regions where you want to create stack instances using the specified AWS accounts .
	Regions *[]*string `json:"regions" yaml:"regions"`
	// A list of stack set parameters whose values you want to override in the selected stack instances.
	ParameterOverrides interface{} `json:"parameterOverrides" yaml:"parameterOverrides"`
}

// Properties for defining a `CfnStackSet`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedExecution interface{}
//   cfnStackSetProps := &cfnStackSetProps{
//   	permissionModel: jsii.String("permissionModel"),
//   	stackSetName: jsii.String("stackSetName"),
//
//   	// the properties below are optional
//   	administrationRoleArn: jsii.String("administrationRoleArn"),
//   	autoDeployment: &autoDeploymentProperty{
//   		enabled: jsii.Boolean(false),
//   		retainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	callAs: jsii.String("callAs"),
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	executionRoleName: jsii.String("executionRoleName"),
//   	managedExecution: managedExecution,
//   	operationPreferences: &operationPreferencesProperty{
//   		failureToleranceCount: jsii.Number(123),
//   		failureTolerancePercentage: jsii.Number(123),
//   		maxConcurrentCount: jsii.Number(123),
//   		maxConcurrentPercentage: jsii.Number(123),
//   		regionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		regionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	stackInstancesGroup: []interface{}{
//   		&stackInstancesProperty{
//   			deploymentTargets: &deploymentTargetsProperty{
//   				accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				organizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			parameterOverrides: []interface{}{
//   				&parameterProperty{
//   					parameterKey: jsii.String("parameterKey"),
//   					parameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateBody: jsii.String("templateBody"),
//   	templateUrl: jsii.String("templateUrl"),
//   }
//
type CfnStackSetProps struct {
	// Describes how the IAM roles required for stack set operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant Self-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations . For more information, see [Grant Service-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html) .
	//
	// *Allowed Values* : `SERVICE_MANAGED` | `SELF_MANAGED`
	//
	// > The `PermissionModel` property is required.
	PermissionModel *string `json:"permissionModel" yaml:"permissionModel"`
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// *Maximum* : `128`
	//
	// *Pattern* : `^[a-zA-Z][a-zA-Z0-9-]{0,127}$`
	//
	// > The `StackSetName` property is required.
	StackSetName *string `json:"stackSetName" yaml:"stackSetName"`
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	//
	// Use customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account. For more information, see [Prerequisites: Granting Permissions for Stack Set Operations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) in the *AWS CloudFormation User Guide* .
	//
	// *Minimum* : `20`
	//
	// *Maximum* : `2048`.
	AdministrationRoleArn *string `json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
	AutoDeployment interface{} `json:"autoDeployment" yaml:"autoDeployment"`
	// [Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.
	//
	// By default, `SELF` is specified. Use `SELF` for stack sets with self-managed permissions.
	//
	// - To create a stack set with service-managed permissions while signed in to the management account, specify `SELF` .
	// - To create a stack set with service-managed permissions while signed in to a delegated administrator account, specify `DELEGATED_ADMIN` .
	//
	// Your AWS account must be registered as a delegated admin in the management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the *AWS CloudFormation User Guide* .
	//
	// Stack sets with service-managed permissions are created in the management account, including stack sets that are created by delegated administrators.
	//
	// *Valid Values* : `SELF` | `DELEGATED_ADMIN`.
	CallAs *string `json:"callAs" yaml:"callAs"`
	// The capabilities that are allowed in the stack set.
	//
	// Some stack set templates might include resources that can affect permissions in your AWS account —for example, by creating new AWS Identity and Access Management ( IAM ) users. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities) .
	Capabilities *[]*string `json:"capabilities" yaml:"capabilities"`
	// A description of the stack set.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	Description *string `json:"description" yaml:"description"`
	// The name of the IAM execution role to use to create the stack set.
	//
	// If you don't specify an execution role, AWS CloudFormation uses the `AWSCloudFormationStackSetExecutionRole` role for the stack set operation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `[a-zA-Z_0-9+=,.@-]+`
	ExecutionRoleName *string `json:"executionRoleName" yaml:"executionRoleName"`
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// When active, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// > If there are already running or queued operations, StackSets queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your stack set's execution configuration while there are running or queued operations for that stack set.
	//
	// When inactive (default), StackSets performs one operation at a time in request order.
	ManagedExecution interface{} `json:"managedExecution" yaml:"managedExecution"`
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences interface{} `json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the stack set template.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// A group of stack instances with parameters in some specific accounts and Regions.
	StackInstancesGroup interface{} `json:"stackInstancesGroup" yaml:"stackInstancesGroup"`
	// The key-value pairs to associate with this stack set and the stacks created from it.
	//
	// AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	Tags *[]*CfnTag `json:"tags" yaml:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates containing dynamic references through `TemplateUrl` instead.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `51200`.
	TemplateBody *string `json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	TemplateUrl *string `json:"templateUrl" yaml:"templateUrl"`
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTag := &cfnTag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTag struct {
	Key *string `json:"key" yaml:"key"`
	Value *string `json:"value" yaml:"value"`
}

// A traffic route, representing where the traffic is being directed to.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTrafficRoute := &cfnTrafficRoute{
//   	logicalId: jsii.String("logicalId"),
//   	type: jsii.String("type"),
//   }
//
type CfnTrafficRoute struct {
	// The logical id of the target resource.
	LogicalId *string `json:"logicalId" yaml:"logicalId"`
	// The resource type of the route.
	//
	// Today, the only allowed value is 'AWS::ElasticLoadBalancingV2::Listener'.
	Type *string `json:"type" yaml:"type"`
}

// Type of the {@link CfnCodeDeployBlueGreenEcsAttributes.trafficRouting} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTrafficRouting := &cfnTrafficRouting{
//   	prodTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   	targetGroups: []*string{
//   		jsii.String("targetGroups"),
//   	},
//   	testTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnTrafficRouting struct {
	// The listener to be used by your load balancer to direct traffic to your target groups.
	ProdTrafficRoute *CfnTrafficRoute `json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// The logical IDs of the blue and green, respectively, AWS::ElasticLoadBalancingV2::TargetGroup target groups.
	TargetGroups *[]*string `json:"targetGroups" yaml:"targetGroups"`
	// The listener to be used by your load balancer to direct traffic to your target groups.
	TestTrafficRoute *CfnTrafficRoute `json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

// Traffic routing configuration settings.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.trafficRoutingConfig} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTrafficRoutingConfig := &cfnTrafficRoutingConfig{
//   	type: cdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   	// the properties below are optional
//   	timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnTrafficRoutingConfig struct {
	// The type of traffic shifting used by the blue-green deployment configuration.
	Type CfnTrafficRoutingType `json:"type" yaml:"type"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_CANARY}.
	TimeBasedCanary *CfnTrafficRoutingTimeBasedCanary `json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
	TimeBasedLinear *CfnTrafficRoutingTimeBasedLinear `json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

// The traffic routing configuration if {@link CfnTrafficRoutingConfig.type} is {@link CfnTrafficRoutingType.TIME_BASED_CANARY}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTrafficRoutingTimeBasedCanary := &cfnTrafficRoutingTimeBasedCanary{
//   	bakeTimeMins: jsii.Number(123),
//   	stepPercentage: jsii.Number(123),
//   }
//
type CfnTrafficRoutingTimeBasedCanary struct {
	// The number of minutes between the first and second traffic shifts of a time-based canary deployment.
	BakeTimeMins *float64 `json:"bakeTimeMins" yaml:"bakeTimeMins"`
	// The percentage of traffic to shift in the first increment of a time-based canary deployment.
	//
	// The step percentage must be 14% or greater.
	StepPercentage *float64 `json:"stepPercentage" yaml:"stepPercentage"`
}

// The traffic routing configuration if {@link CfnTrafficRoutingConfig.type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTrafficRoutingTimeBasedLinear := &cfnTrafficRoutingTimeBasedLinear{
//   	bakeTimeMins: jsii.Number(123),
//   	stepPercentage: jsii.Number(123),
//   }
//
type CfnTrafficRoutingTimeBasedLinear struct {
	// The number of minutes between the first and second traffic shifts of a time-based linear deployment.
	BakeTimeMins *float64 `json:"bakeTimeMins" yaml:"bakeTimeMins"`
	// The percentage of traffic that is shifted at the start of each increment of a time-based linear deployment.
	//
	// The step percentage must be 14% or greater.
	StepPercentage *float64 `json:"stepPercentage" yaml:"stepPercentage"`
}

// The possible types of traffic shifting for the blue-green deployment configuration.
//
// The type of the {@link CfnTrafficRoutingConfig.type} property.
type CfnTrafficRoutingType string

const (
	// Switch from blue to green at once.
	CfnTrafficRoutingType_ALL_AT_ONCE CfnTrafficRoutingType = "ALL_AT_ONCE"
	// Specifies a configuration that shifts traffic from blue to green in two increments.
	CfnTrafficRoutingType_TIME_BASED_CANARY CfnTrafficRoutingType = "TIME_BASED_CANARY"
	// Specifies a configuration that shifts traffic from blue to green in equal increments, with an equal number of minutes between each increment.
	CfnTrafficRoutingType_TIME_BASED_LINEAR CfnTrafficRoutingType = "TIME_BASED_LINEAR"
)

// A CloudFormation `AWS::CloudFormation::TypeActivation`.
//
// Activates a public third-party extension, making it available for use in stack templates. For more information, see [Using public extensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html) in the *AWS CloudFormation User Guide* .
//
// Once you have activated a public third-party extension in your account and region, use [SetTypeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html) to specify configuration properties for the extension. For more information, see [Configuring extensions at the account level](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration) in the *CloudFormation User Guide* .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTypeActivation := cdk.NewCfnTypeActivation(this, jsii.String("MyCfnTypeActivation"), &cfnTypeActivationProps{
//   	autoUpdate: jsii.Boolean(false),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   	majorVersion: jsii.String("majorVersion"),
//   	publicTypeArn: jsii.String("publicTypeArn"),
//   	publisherId: jsii.String("publisherId"),
//   	type: jsii.String("type"),
//   	typeName: jsii.String("typeName"),
//   	typeNameAlias: jsii.String("typeNameAlias"),
//   	versionBump: jsii.String("versionBump"),
//   })
//
type CfnTypeActivation interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the activated extension, in this account and Region.
	AttrArn() *string
	// Whether to automatically update the extension in this account and region when a new *minor* version is published by the extension publisher.
	//
	// Major versions released by the publisher must be manually updated.
	//
	// The default is `true` .
	AutoUpdate() interface{}
	SetAutoUpdate(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the IAM execution role to use to activate the extension.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// Specifies logging configuration information for an extension.
	LoggingConfig() interface{}
	SetLoggingConfig(val interface{})
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
	// The major version of this extension you want to activate, if multiple major versions are available.
	//
	// The default is the latest major version. CloudFormation uses the latest available *minor* version of the major version selected.
	//
	// You can specify `MajorVersion` or `VersionBump` , but not both.
	MajorVersion() *string
	SetMajorVersion(val *string)
	// The tree node.
	Node() constructs.Node
	// The Amazon Resource Number (ARN) of the public extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublicTypeArn() *string
	SetPublicTypeArn(val *string)
	// The ID of the extension publisher.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublisherId() *string
	SetPublisherId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The extension type.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	Type() *string
	SetType(val *string)
	// The name of the extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	TypeName() *string
	SetTypeName(val *string)
	// An alias to assign to the public extension, in this account and region.
	//
	// If you specify an alias for the extension, CloudFormation treats the alias as the extension type name within this account and region. You must use the alias to refer to the extension in your templates, API calls, and CloudFormation console.
	//
	// An extension alias must be unique within a given account and region. You can activate the same public resource multiple times in the same account and region, using different type name aliases.
	TypeNameAlias() *string
	SetTypeNameAlias(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Manually updates a previously-activated type to a new major or minor version, if available.
	//
	// You can also use this parameter to update the value of `AutoUpdate` .
	//
	// - `MAJOR` : CloudFormation updates the extension to the newest major version, if one is available.
	// - `MINOR` : CloudFormation updates the extension to the newest minor version, if one is available.
	VersionBump() *string
	SetVersionBump(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnTypeActivation
type jsiiProxy_CfnTypeActivation struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnTypeActivation) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) AutoUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) LoggingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) MajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) PublicTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) PublisherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) TypeNameAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) VersionBump() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionBump",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::TypeActivation`.
func NewCfnTypeActivation(scope constructs.Construct, id *string, props *CfnTypeActivationProps) CfnTypeActivation {
	_init_.Initialize()

	j := jsiiProxy_CfnTypeActivation{}

	_jsii_.Create(
		"aws-cdk-lib.CfnTypeActivation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::TypeActivation`.
func NewCfnTypeActivation_Override(c CfnTypeActivation, scope constructs.Construct, id *string, props *CfnTypeActivationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnTypeActivation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetAutoUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"autoUpdate",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetLoggingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"loggingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetMajorVersion(val *string) {
	_jsii_.Set(
		j,
		"majorVersion",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetPublicTypeArn(val *string) {
	_jsii_.Set(
		j,
		"publicTypeArn",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetPublisherId(val *string) {
	_jsii_.Set(
		j,
		"publisherId",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetTypeNameAlias(val *string) {
	_jsii_.Set(
		j,
		"typeNameAlias",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation) SetVersionBump(val *string) {
	_jsii_.Set(
		j,
		"versionBump",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTypeActivation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTypeActivation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnTypeActivation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTypeActivation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnTypeActivation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTypeActivation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTypeActivation) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTypeActivation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTypeActivation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains logging configuration information for an extension.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnTypeActivation_LoggingConfigProperty struct {
	// The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the extension's handlers.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log entries to CloudWatch Logs.
	LogRoleArn *string `json:"logRoleArn" yaml:"logRoleArn"`
}

// Properties for defining a `CfnTypeActivation`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnTypeActivationProps := &cfnTypeActivationProps{
//   	autoUpdate: jsii.Boolean(false),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   	majorVersion: jsii.String("majorVersion"),
//   	publicTypeArn: jsii.String("publicTypeArn"),
//   	publisherId: jsii.String("publisherId"),
//   	type: jsii.String("type"),
//   	typeName: jsii.String("typeName"),
//   	typeNameAlias: jsii.String("typeNameAlias"),
//   	versionBump: jsii.String("versionBump"),
//   }
//
type CfnTypeActivationProps struct {
	// Whether to automatically update the extension in this account and region when a new *minor* version is published by the extension publisher.
	//
	// Major versions released by the publisher must be manually updated.
	//
	// The default is `true` .
	AutoUpdate interface{} `json:"autoUpdate" yaml:"autoUpdate"`
	// The name of the IAM execution role to use to activate the extension.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies logging configuration information for an extension.
	LoggingConfig interface{} `json:"loggingConfig" yaml:"loggingConfig"`
	// The major version of this extension you want to activate, if multiple major versions are available.
	//
	// The default is the latest major version. CloudFormation uses the latest available *minor* version of the major version selected.
	//
	// You can specify `MajorVersion` or `VersionBump` , but not both.
	MajorVersion *string `json:"majorVersion" yaml:"majorVersion"`
	// The Amazon Resource Number (ARN) of the public extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublicTypeArn *string `json:"publicTypeArn" yaml:"publicTypeArn"`
	// The ID of the extension publisher.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublisherId *string `json:"publisherId" yaml:"publisherId"`
	// The extension type.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	Type *string `json:"type" yaml:"type"`
	// The name of the extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	TypeName *string `json:"typeName" yaml:"typeName"`
	// An alias to assign to the public extension, in this account and region.
	//
	// If you specify an alias for the extension, CloudFormation treats the alias as the extension type name within this account and region. You must use the alias to refer to the extension in your templates, API calls, and CloudFormation console.
	//
	// An extension alias must be unique within a given account and region. You can activate the same public resource multiple times in the same account and region, using different type name aliases.
	TypeNameAlias *string `json:"typeNameAlias" yaml:"typeNameAlias"`
	// Manually updates a previously-activated type to a new major or minor version, if available.
	//
	// You can also use this parameter to update the value of `AutoUpdate` .
	//
	// - `MAJOR` : CloudFormation updates the extension to the newest major version, if one is available.
	// - `MINOR` : CloudFormation updates the extension to the newest minor version, if one is available.
	VersionBump *string `json:"versionBump" yaml:"versionBump"`
}

// Use the UpdatePolicy attribute to specify how AWS CloudFormation handles updates to the AWS::AutoScaling::AutoScalingGroup resource.
//
// AWS CloudFormation invokes one of three update policies depending on the type of change you make or whether a
// scheduled action is associated with the Auto Scaling group.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnUpdatePolicy := &cfnUpdatePolicy{
//   	autoScalingReplacingUpdate: &cfnAutoScalingReplacingUpdate{
//   		willReplace: jsii.Boolean(false),
//   	},
//   	autoScalingRollingUpdate: &cfnAutoScalingRollingUpdate{
//   		maxBatchSize: jsii.Number(123),
//   		minInstancesInService: jsii.Number(123),
//   		minSuccessfulInstancesPercent: jsii.Number(123),
//   		pauseTime: jsii.String("pauseTime"),
//   		suspendProcesses: []*string{
//   			jsii.String("suspendProcesses"),
//   		},
//   		waitOnResourceSignals: jsii.Boolean(false),
//   	},
//   	autoScalingScheduledAction: &cfnAutoScalingScheduledAction{
//   		ignoreUnmodifiedGroupSizeProperties: jsii.Boolean(false),
//   	},
//   	codeDeployLambdaAliasUpdate: &cfnCodeDeployLambdaAliasUpdate{
//   		applicationName: jsii.String("applicationName"),
//   		deploymentGroupName: jsii.String("deploymentGroupName"),
//
//   		// the properties below are optional
//   		afterAllowTrafficHook: jsii.String("afterAllowTrafficHook"),
//   		beforeAllowTrafficHook: jsii.String("beforeAllowTrafficHook"),
//   	},
//   	enableVersionUpgrade: jsii.Boolean(false),
//   	useOnlineResharding: jsii.Boolean(false),
//   }
//
type CfnUpdatePolicy struct {
	// Specifies whether an Auto Scaling group and the instances it contains are replaced during an update.
	//
	// During replacement,
	// AWS CloudFormation retains the old group until it finishes creating the new one. If the update fails, AWS CloudFormation
	// can roll back to the old Auto Scaling group and delete the new Auto Scaling group.
	AutoScalingReplacingUpdate *CfnAutoScalingReplacingUpdate `json:"autoScalingReplacingUpdate" yaml:"autoScalingReplacingUpdate"`
	// To specify how AWS CloudFormation handles rolling updates for an Auto Scaling group, use the AutoScalingRollingUpdate policy.
	//
	// Rolling updates enable you to specify whether AWS CloudFormation updates instances that are in an Auto Scaling
	// group in batches or all at once.
	AutoScalingRollingUpdate *CfnAutoScalingRollingUpdate `json:"autoScalingRollingUpdate" yaml:"autoScalingRollingUpdate"`
	// To specify how AWS CloudFormation handles updates for the MinSize, MaxSize, and DesiredCapacity properties when the AWS::AutoScaling::AutoScalingGroup resource has an associated scheduled action, use the AutoScalingScheduledAction policy.
	AutoScalingScheduledAction *CfnAutoScalingScheduledAction `json:"autoScalingScheduledAction" yaml:"autoScalingScheduledAction"`
	// To perform an AWS CodeDeploy deployment when the version changes on an AWS::Lambda::Alias resource, use the CodeDeployLambdaAliasUpdate update policy.
	CodeDeployLambdaAliasUpdate *CfnCodeDeployLambdaAliasUpdate `json:"codeDeployLambdaAliasUpdate" yaml:"codeDeployLambdaAliasUpdate"`
	// To upgrade an Amazon ES domain to a new version of Elasticsearch rather than replacing the entire AWS::Elasticsearch::Domain resource, use the EnableVersionUpgrade update policy.
	EnableVersionUpgrade *bool `json:"enableVersionUpgrade" yaml:"enableVersionUpgrade"`
	// To modify a replication group's shards by adding or removing shards, rather than replacing the entire AWS::ElastiCache::ReplicationGroup resource, use the UseOnlineResharding update policy.
	UseOnlineResharding *bool `json:"useOnlineResharding" yaml:"useOnlineResharding"`
}

// A CloudFormation `AWS::CloudFormation::WaitCondition`.
//
// > For Amazon EC2 and Auto Scaling resources, we recommend that you use a `CreationPolicy` attribute instead of wait conditions. Add a CreationPolicy attribute to those resources, and use the cfn-signal helper script to signal when an instance creation process has completed successfully.
//
// You can use a wait condition for situations like the following:
//
// - To coordinate stack resource creation with configuration actions that are external to the stack creation.
// - To track the status of a configuration process.
//
// For these situations, we recommend that you associate a [CreationPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-creationpolicy.html) attribute with the wait condition so that you don't have to use a wait condition handle. For more information and an example, see [Creating wait conditions in a template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-waitcondition.html) . If you use a CreationPolicy with a wait condition, don't specify any of the wait condition's properties.
//
// > If you use the [VPC endpoints](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints.html) feature, resources in the VPC that respond to wait conditions must have access to CloudFormation , specific Amazon Simple Storage Service ( Amazon S3 ) buckets. Resources must send wait condition responses to a presigned Amazon S3 URL. If they can't send responses to Amazon S3 , CloudFormation won't receive a response and the stack operation fails. For more information, see [Setting up VPC endpoints for AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-vpce-bucketnames.html) .
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnWaitCondition := cdk.NewCfnWaitCondition(this, jsii.String("MyCfnWaitCondition"), &cfnWaitConditionProps{
//   	count: jsii.Number(123),
//   	handle: jsii.String("handle"),
//   	timeout: jsii.String("timeout"),
//   })
//
type CfnWaitCondition interface {
	CfnResource
	IInspectable
	// A JSON object that contains the `UniqueId` and `Data` values from the wait condition signal(s) for the specified wait condition.
	//
	// For more information about wait condition signals, see [Wait condition signal JSON format](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-waitcondition.html#using-cfn-waitcondition-signaljson) .
	//
	// Example return value for a wait condition with 2 signals:
	//
	// `{ "Signal1" : "Step 1 complete." , "Signal2" : "Step 2 complete." }`
	AttrData() IResolvable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The number of success signals that CloudFormation must receive before it continues the stack creation process.
	//
	// When the wait condition receives the requisite number of success signals, CloudFormation resumes the creation of the stack. If the wait condition doesn't receive the specified number of success signals before the Timeout period expires, CloudFormation assumes that the wait condition has failed and rolls the stack back.
	//
	// Updates aren't supported.
	Count() *float64
	SetCount(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to the wait condition handle used to signal this wait condition.
	//
	// Use the `Ref` intrinsic function to specify an [AWS::CloudFormation::WaitConditionHandle](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html) resource.
	//
	// Anytime you add a WaitCondition resource during a stack update, you must associate the wait condition with a new WaitConditionHandle resource. Don't reuse an old wait condition handle that has already been defined in the template. If you reuse a wait condition handle, the wait condition might evaluate old signals from a previous create or update stack command.
	//
	// Updates aren't supported.
	Handle() *string
	SetHandle(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The length of time (in seconds) to wait for the number of signals that the `Count` property specifies.
	//
	// `Timeout` is a minimum-bound property, meaning the timeout occurs no sooner than the time you specify, but can occur shortly thereafter. The maximum time that can be specified for this property is 12 hours (43200 seconds).
	//
	// Updates aren't supported.
	Timeout() *string
	SetTimeout(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnWaitCondition
type jsiiProxy_CfnWaitCondition struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnWaitCondition) AttrData() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"attrData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Handle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitCondition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::WaitCondition`.
func NewCfnWaitCondition(scope constructs.Construct, id *string, props *CfnWaitConditionProps) CfnWaitCondition {
	_init_.Initialize()

	j := jsiiProxy_CfnWaitCondition{}

	_jsii_.Create(
		"aws-cdk-lib.CfnWaitCondition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::WaitCondition`.
func NewCfnWaitCondition_Override(c CfnWaitCondition, scope constructs.Construct, id *string, props *CfnWaitConditionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnWaitCondition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWaitCondition) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CfnWaitCondition) SetHandle(val *string) {
	_jsii_.Set(
		j,
		"handle",
		val,
	)
}

func (j *jsiiProxy_CfnWaitCondition) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWaitCondition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitCondition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWaitCondition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitCondition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnWaitCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWaitCondition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnWaitCondition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWaitCondition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWaitCondition) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWaitCondition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWaitCondition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWaitCondition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWaitCondition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWaitCondition) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWaitCondition) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitCondition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitCondition) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWaitCondition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWaitCondition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitCondition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitCondition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::CloudFormation::WaitConditionHandle`.
//
// > For Amazon EC2 and Auto Scaling resources, we recommend that you use a `CreationPolicy` attribute instead of wait conditions. Add a `CreationPolicy` attribute to those resources, and use the cfn-signal helper script to signal when an instance creation process has completed successfully.
// >
// > For more information, see [Deploying applications on Amazon EC2 with AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/deploying.applications.html) .
//
// The `AWS::CloudFormation::WaitConditionHandle` type has no properties. When you reference the `WaitConditionHandle` resource by using the Ref function, AWS CloudFormation returns a presigned URL. You pass this URL to applications or scripts that are running on your Amazon EC2 instances to send signals to that URL. An associated `AWS::CloudFormation::WaitCondition` resource checks the URL for the required number of success signals or for a failure signal.
//
// > Anytime you add a `WaitCondition` resource during a stack update or update a resource with a wait condition, you must associate the wait condition with a new `WaitConditionHandle` resource. Don't reuse an old wait condition handle that has already been defined in the template. If you reuse a wait condition handle, the wait condition might evaluate old signals from a previous create or update stack command. > Updates aren't supported for this resource.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnWaitConditionHandle := cdk.NewCfnWaitConditionHandle(this, jsii.String("MyCfnWaitConditionHandle"))
//
type CfnWaitConditionHandle interface {
	CfnResource
	IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnWaitConditionHandle
type jsiiProxy_CfnWaitConditionHandle struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnWaitConditionHandle) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWaitConditionHandle) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::WaitConditionHandle`.
func NewCfnWaitConditionHandle(scope constructs.Construct, id *string) CfnWaitConditionHandle {
	_init_.Initialize()

	j := jsiiProxy_CfnWaitConditionHandle{}

	_jsii_.Create(
		"aws-cdk-lib.CfnWaitConditionHandle",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::WaitConditionHandle`.
func NewCfnWaitConditionHandle_Override(c CfnWaitConditionHandle, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnWaitConditionHandle",
		[]interface{}{scope, id},
		c,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWaitConditionHandle_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitConditionHandle",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWaitConditionHandle_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitConditionHandle",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnWaitConditionHandle_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnWaitConditionHandle",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWaitConditionHandle_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnWaitConditionHandle",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddDependsOn(target CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWaitConditionHandle) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWaitConditionHandle) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWaitCondition`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cfnWaitConditionProps := &cfnWaitConditionProps{
//   	count: jsii.Number(123),
//   	handle: jsii.String("handle"),
//   	timeout: jsii.String("timeout"),
//   }
//
type CfnWaitConditionProps struct {
	// The number of success signals that CloudFormation must receive before it continues the stack creation process.
	//
	// When the wait condition receives the requisite number of success signals, CloudFormation resumes the creation of the stack. If the wait condition doesn't receive the specified number of success signals before the Timeout period expires, CloudFormation assumes that the wait condition has failed and rolls the stack back.
	//
	// Updates aren't supported.
	Count *float64 `json:"count" yaml:"count"`
	// A reference to the wait condition handle used to signal this wait condition.
	//
	// Use the `Ref` intrinsic function to specify an [AWS::CloudFormation::WaitConditionHandle](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html) resource.
	//
	// Anytime you add a WaitCondition resource during a stack update, you must associate the wait condition with a new WaitConditionHandle resource. Don't reuse an old wait condition handle that has already been defined in the template. If you reuse a wait condition handle, the wait condition might evaluate old signals from a previous create or update stack command.
	//
	// Updates aren't supported.
	Handle *string `json:"handle" yaml:"handle"`
	// The length of time (in seconds) to wait for the number of signals that the `Count` property specifies.
	//
	// `Timeout` is a minimum-bound property, meaning the timeout occurs no sooner than the time you specify, but can occur shortly thereafter. The maximum time that can be specified for this property is 12 hours (43200 seconds).
	//
	// Updates aren't supported.
	Timeout *string `json:"timeout" yaml:"timeout"`
}

// A synthesizer that uses conventional asset locations, but not conventional deployment roles.
//
// Instead of assuming the bootstrapped deployment roles, all stack operations will be performed
// using the CLI's current credentials.
//
// - This synthesizer does not support deploying to accounts to which the CLI does not have
//    credentials. It also does not support deploying using **CDK Pipelines**. For either of those
//    features, use `DefaultStackSynthesizer`.
// - This synthesizer requires an S3 bucket and ECR repository with well-known names. To
//    not depend on those, use `LegacyStackSynthesizer`.
//
// Be aware that your CLI credentials must be valid for the duration of the
// entire deployment. If you are using session credentials, make sure the
// session lifetime is long enough.
//
// By default, expects the environment to have been bootstrapped with just the staging resources
// of the Bootstrap Stack V2 (also known as "modern bootstrap stack"). You can override
// the default names using the synthesizer's construction properties.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cliCredentialsStackSynthesizer := cdk.NewCliCredentialsStackSynthesizer(&cliCredentialsStackSynthesizerProps{
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	dockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	fileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	imageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	qualifier: jsii.String("qualifier"),
//   })
//
type CliCredentialsStackSynthesizer interface {
	StackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for CliCredentialsStackSynthesizer
type jsiiProxy_CliCredentialsStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func NewCliCredentialsStackSynthesizer(props *CliCredentialsStackSynthesizerProps) CliCredentialsStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_CliCredentialsStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCliCredentialsStackSynthesizer_Override(c CliCredentialsStackSynthesizer, props *CliCredentialsStackSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CliCredentialsStackSynthesizer",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		c,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		c,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		c,
		"bind",
		[]interface{}{stack},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		c,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// Properties for the CliCredentialsStackSynthesizer.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   cliCredentialsStackSynthesizerProps := &cliCredentialsStackSynthesizerProps{
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	dockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	fileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	imageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	qualifier: jsii.String("qualifier"),
//   }
//
type CliCredentialsStackSynthesizerProps struct {
	// bucketPrefix to use while storing S3 Assets.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// A prefix to use while tagging and uploading Docker images to ECR.
	//
	// This does not add any separators - the source hash will be appended to
	// this string directly.
	DockerTagPrefix *string `json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// Name of the S3 bucket to hold file assets.
	//
	// You must supply this if you have given a non-standard name to the staging bucket.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	FileAssetsBucketName *string `json:"fileAssetsBucketName" yaml:"fileAssetsBucketName"`
	// Name of the ECR repository to hold Docker Image assets.
	//
	// You must supply this if you have given a non-standard name to the ECR repository.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	ImageAssetsRepositoryName *string `json:"imageAssetsRepositoryName" yaml:"imageAssetsRepositoryName"`
	// Qualifier to disambiguate multiple environments in the same account.
	//
	// You can use this and leave the other naming properties empty if you have deployed
	// the bootstrap environment with standard names but only differnet qualifiers.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
}

// Base class for the model side of context providers.
//
// Instances of this class communicate with context provider plugins in the 'cdk
// toolkit' via context variables (input), outputting specialized queries for
// more context variables (output).
//
// ContextProvider needs access to a Construct to hook into the context mechanism.
type ContextProvider interface {
}

// The jsii proxy struct for ContextProvider
type jsiiProxy_ContextProvider struct {
	_ byte // padding
}

// Returns: the context key or undefined if a key cannot be rendered (due to tokens used in any of the props).
func ContextProvider_GetKey(scope constructs.Construct, options *GetContextKeyOptions) *GetContextKeyResult {
	_init_.Initialize()

	var returns *GetContextKeyResult

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ContextProvider",
		"getKey",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func ContextProvider_GetValue(scope constructs.Construct, options *GetContextValueOptions) *GetContextValueResult {
	_init_.Initialize()

	var returns *GetContextValueResult

	_jsii_.StaticInvoke(
		"aws-cdk-lib.ContextProvider",
		"getValue",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Options applied when copying directories.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   copyOptions := &copyOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type CopyOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	Follow SymlinkFollowMode `json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
}

// Instantiation of a custom resource, whose implementation is provided a Provider.
//
// This class is intended to be used by construct library authors. Application
// builder should not be able to tell whether or not a construct is backed by
// a custom resource, and so the use of this class should be invisible.
//
// Instead, construct library authors declare a custom construct that hides the
// choice of provider, and accepts a strongly-typed properties object with the
// properties your provider accepts.
//
// Your custom resource provider (identified by the `serviceToken` property)
// can be one of 4 constructs:
//
// - If you are authoring a construct library or application, we recommend you
//    use the `Provider` class in the `custom-resources` module.
// - If you are authoring a construct for the CDK's AWS Construct Library,
//    you should use the `CustomResourceProvider` construct in this package.
// - If you want full control over the provider, you can always directly use
//    a Lambda Function or SNS Topic by passing the ARN into `serviceToken`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := customresources.NewProvider(this, jsii.String("MyProvider"), &providerProps{
//   	onEventHandler: onEventHandler,
//   	isCompleteHandler: isCompleteHandler,
//   })
//
//   NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
//   	serviceToken: provider.serviceToken,
//   })
//
type CustomResource interface {
	Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The physical name of this custom resource.
	Ref() *string
	// The stack in which this resource is defined.
	Stack() Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt`. Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getAttString` for string attributes.
	GetAtt(attributeName *string) Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt` encoded as a string.
	GetAttString(attributeName *string) *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomResource
type jsiiProxy_CustomResource struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_CustomResource) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCustomResource(scope constructs.Construct, id *string, props *CustomResourceProps) CustomResource {
	_init_.Initialize()

	j := jsiiProxy_CustomResource{}

	_jsii_.Create(
		"aws-cdk-lib.CustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomResource_Override(c CustomResource, scope constructs.Construct, id *string, props *CustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CustomResource",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CustomResource_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) ApplyRemovalPolicy(policy RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CustomResource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetAtt(attributeName *string) Reference {
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetAttString(attributeName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetResourceArnAttribute(arnAttr *string, arnComponents *ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to provide a Lambda-backed custom resource.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := customresources.NewProvider(this, jsii.String("MyProvider"), &providerProps{
//   	onEventHandler: onEventHandler,
//   	isCompleteHandler: isCompleteHandler,
//   })
//
//   NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
//   	serviceToken: provider.serviceToken,
//   })
//
type CustomResourceProps struct {
	// The ARN of the provider which implements this custom resource type.
	//
	// You can implement a provider by listening to raw AWS CloudFormation events
	// and specify the ARN of an SNS topic (`topic.topicArn`) or the ARN of an AWS
	// Lambda function (`lambda.functionArn`) or use the CDK's custom [resource
	// provider framework] which makes it easier to implement robust providers.
	//
	// [resource provider framework]:
	// https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html
	//
	// Provider framework:
	//
	// ```ts
	// // use the provider framework from aws-cdk/custom-resources:
	// const provider = new customresources.Provider(this, 'ResourceProvider', {
	//    onEventHandler,
	//    isCompleteHandler, // optional
	// });
	//
	// new CustomResource(this, 'MyResource', {
	//    serviceToken: provider.serviceToken,
	// });
	// ```
	//
	// AWS Lambda function:
	//
	// ```ts
	// // invoke an AWS Lambda function when a lifecycle event occurs:
	// new CustomResource(this, 'MyResource', {
	//    serviceToken: myFunction.functionArn,
	// });
	// ```
	//
	// SNS topic:
	//
	// ```ts
	// // publish lifecycle events to an SNS topic:
	// new CustomResource(this, 'MyResource', {
	//    serviceToken: myTopic.topicArn,
	// });
	// ```.
	ServiceToken *string `json:"serviceToken" yaml:"serviceToken"`
	// Convert all property keys to pascal case.
	PascalCaseProperties *bool `json:"pascalCaseProperties" yaml:"pascalCaseProperties"`
	// Properties to pass to the Lambda.
	Properties *map[string]interface{} `json:"properties" yaml:"properties"`
	// The policy to apply when this resource is removed from the application.
	RemovalPolicy RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// For custom resources, you can specify AWS::CloudFormation::CustomResource (the default) as the resource type, or you can specify your own resource type name.
	//
	// For example, you can use "Custom::MyCustomResourceTypeName".
	//
	// Custom resource type names must begin with "Custom::" and can include
	// alphanumeric characters and the following characters: _@-. You can specify
	// a custom resource type name up to a maximum length of 60 characters. You
	// cannot change the type during an update.
	//
	// Using your own resource type names helps you quickly differentiate the
	// types of custom resources in your stack. For example, if you had two custom
	// resources that conduct two different ping tests, you could name their type
	// as Custom::PingTester to make them easily identifiable as ping testers
	// (instead of using AWS::CloudFormation::CustomResource).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#aws-cfn-resource-type-name
	//
	ResourceType *string `json:"resourceType" yaml:"resourceType"`
}

// An AWS-Lambda backed custom resource provider, for CDK Construct Library constructs.
//
// This is a provider for `CustomResource` constructs, backed by an AWS Lambda
// Function. It only supports NodeJS runtimes.
//
// **This is not a generic custom resource provider class**. It is specifically
// intended to be used only by constructs in the AWS CDK Construct Library, and
// only exists here because of reverse dependency issues (for example, it cannot
// use `iam.PolicyStatement` objects, since the `iam` library already depends on
// the CDK `core` library and we cannot have cyclic dependencies).
//
// If you are not writing constructs for the AWS Construct Library, you should
// use the `Provider` class in the `custom-resources` module instead, which has
// a better API and supports all Lambda runtimes, not just Node.
//
// N.B.: When you are writing Custom Resource Providers, there are a number of
// lifecycle events you have to pay attention to. These are documented in the
// README of the `custom-resources` module. Be sure to give the documentation
// in that module a read, regardless of whether you end up using the Provider
// class in there or this one.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := customResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	runtime: customResourceProviderRuntime_NODEJS_12_X,
//   })
//
//   roleArn := provider.roleArn
//
type CustomResourceProvider interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The ARN of the provider's AWS Lambda function role.
	RoleArn() *string
	// The ARN of the provider's AWS Lambda function which should be used as the `serviceToken` when defining a custom resource.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   var myProvider customResourceProvider
	//
	//
	//   NewCustomResource(this, jsii.String("MyCustomResource"), &customResourceProps{
	//   	serviceToken: myProvider.serviceToken,
	//   	properties: map[string]interface{}{
	//   		"myPropertyOne": jsii.String("one"),
	//   		"myPropertyTwo": jsii.String("two"),
	//   	},
	//   })
	//
	ServiceToken() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomResourceProvider
type jsiiProxy_CustomResourceProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomResourceProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


func NewCustomResourceProvider(scope constructs.Construct, id *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	j := jsiiProxy_CustomResourceProvider{}

	_jsii_.Create(
		"aws-cdk-lib.CustomResourceProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomResourceProvider_Override(c CustomResourceProvider, scope constructs.Construct, id *string, props *CustomResourceProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CustomResourceProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns a stack-level singleton ARN (service token) for the custom resource provider.
//
// Returns: the service token of the custom resource provider, which should be
// used when defining a `CustomResource`.
func CustomResourceProvider_GetOrCreate(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
		"getOrCreate",
		[]interface{}{scope, uniqueid, props},
		&returns,
	)

	return returns
}

// Returns a stack-level singleton for the custom resource provider.
//
// Returns: the service token of the custom resource provider, which should be
// used when defining a `CustomResource`.
func CustomResourceProvider_GetOrCreateProvider(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
		"getOrCreateProvider",
		[]interface{}{scope, uniqueid, props},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomResourceProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization properties for `CustomResourceProvider`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := customResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	runtime: customResourceProviderRuntime_NODEJS_12_X,
//   	policyStatements: []interface{}{
//   		map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("s3:PutObject*"),
//   			"Resource": jsii.String("*"),
//   		},
//   	},
//   })
//
type CustomResourceProviderProps struct {
	// A local file system directory with the provider's code.
	//
	// The code will be
	// bundled into a zip asset and wired to the provider's AWS Lambda function.
	CodeDirectory *string `json:"codeDirectory" yaml:"codeDirectory"`
	// The AWS Lambda runtime and version to use for the provider.
	Runtime CustomResourceProviderRuntime `json:"runtime" yaml:"runtime"`
	// A description of the function.
	Description *string `json:"description" yaml:"description"`
	// Key-value pairs that are passed to Lambda as Environment.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The amount of memory that your function has access to.
	//
	// Increasing the
	// function's memory also increases its CPU allocation.
	MemorySize Size `json:"memorySize" yaml:"memorySize"`
	// A set of IAM policy statements to include in the inline policy of the provider's lambda function.
	//
	// **Please note**: these are direct IAM JSON policy blobs, *not* `iam.PolicyStatement`
	// objects like you will see in the rest of the CDK.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   provider := customResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
	//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	//   	runtime: customResourceProviderRuntime_NODEJS_12_X,
	//   	policyStatements: []interface{}{
	//   		map[string]*string{
	//   			"Effect": jsii.String("Allow"),
	//   			"Action": jsii.String("s3:PutObject*"),
	//   			"Resource": jsii.String("*"),
	//   		},
	//   	},
	//   })
	//
	PolicyStatements *[]interface{} `json:"policyStatements" yaml:"policyStatements"`
	// AWS Lambda timeout for the provider.
	Timeout Duration `json:"timeout" yaml:"timeout"`
}

// The lambda runtime to use for the resource provider.
//
// This also indicates
// which language is used for the handler.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := customResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	runtime: customResourceProviderRuntime_NODEJS_12_X,
//   	policyStatements: []interface{}{
//   		map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("s3:PutObject*"),
//   			"Resource": jsii.String("*"),
//   		},
//   	},
//   })
//
type CustomResourceProviderRuntime string

const (
	// Node.js 12.x.
	CustomResourceProviderRuntime_NODEJS_12_X CustomResourceProviderRuntime = "NODEJS_12_X"
	// Node.js 14.x.
	CustomResourceProviderRuntime_NODEJS_14_X CustomResourceProviderRuntime = "NODEJS_14_X"
)

// Uses conventionally named roles and asset storage locations.
//
// This synthesizer:
//
// - Supports cross-account deployments (the CLI can have credentials to one
//    account, and you can still deploy to another account by assuming roles with
//    well-known names in the other account).
// - Supports the **CDK Pipelines** library.
//
// Requires the environment to have been bootstrapped with Bootstrap Stack V2
// (also known as "modern bootstrap stack"). The synthesizer adds a version
// check to the template, to make sure the bootstrap stack is recent enough
// to support all features expected by this synthesizer.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewMyStack(app, jsii.String("MyStack"), &stackProps{
//   	synthesizer: NewDefaultStackSynthesizer(&defaultStackSynthesizerProps{
//   		fileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
//   	}),
//   })
//
type DefaultStackSynthesizer interface {
	StackSynthesizer
	// Returns the ARN of the CFN execution Role.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	DeployRoleArn() *string
	Stack() Stack
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for DefaultStackSynthesizer
type jsiiProxy_DefaultStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func (j *jsiiProxy_DefaultStackSynthesizer) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultStackSynthesizer) DeployRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultStackSynthesizer) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDefaultStackSynthesizer(props *DefaultStackSynthesizerProps) DefaultStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_DefaultStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.DefaultStackSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDefaultStackSynthesizer_Override(d DefaultStackSynthesizer, props *DefaultStackSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DefaultStackSynthesizer",
		[]interface{}{props},
		d,
	)
}

func DefaultStackSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		d,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		d,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		d,
		"bind",
		[]interface{}{stack},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		d,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// Configuration properties for DefaultStackSynthesizer.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewMyStack(app, jsii.String("MyStack"), &stackProps{
//   	synthesizer: NewDefaultStackSynthesizer(&defaultStackSynthesizerProps{
//   		fileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
//   	}),
//   })
//
type DefaultStackSynthesizerProps struct {
	// Bootstrap stack version SSM parameter.
	//
	// The placeholder `${Qualifier}` will be replaced with the value of qualifier.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// bucketPrefix to use while storing S3 Assets.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
	// The role CloudFormation will assume when deploying the Stack.
	//
	// You must supply this if you have given a non-standard name to the execution role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	CloudFormationExecutionRole *string `json:"cloudFormationExecutionRole" yaml:"cloudFormationExecutionRole"`
	// The role to assume to initiate a deployment in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	DeployRoleArn *string `json:"deployRoleArn" yaml:"deployRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	DeployRoleExternalId *string `json:"deployRoleExternalId" yaml:"deployRoleExternalId"`
	// A prefix to use while tagging and uploading Docker images to ECR.
	//
	// This does not add any separators - the source hash will be appended to
	// this string directly.
	DockerTagPrefix *string `json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// External ID to use when assuming role for file asset publishing.
	FileAssetPublishingExternalId *string `json:"fileAssetPublishingExternalId" yaml:"fileAssetPublishingExternalId"`
	// The role to use to publish file assets to the S3 bucket in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	FileAssetPublishingRoleArn *string `json:"fileAssetPublishingRoleArn" yaml:"fileAssetPublishingRoleArn"`
	// Name of the S3 bucket to hold file assets.
	//
	// You must supply this if you have given a non-standard name to the staging bucket.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	FileAssetsBucketName *string `json:"fileAssetsBucketName" yaml:"fileAssetsBucketName"`
	// Whether to add a Rule to the stack template verifying the bootstrap stack version.
	//
	// This generally should be left set to `true`, unless you explicitly
	// want to be able to deploy to an unbootstrapped environment.
	GenerateBootstrapVersionRule *bool `json:"generateBootstrapVersionRule" yaml:"generateBootstrapVersionRule"`
	// External ID to use when assuming role for image asset publishing.
	ImageAssetPublishingExternalId *string `json:"imageAssetPublishingExternalId" yaml:"imageAssetPublishingExternalId"`
	// The role to use to publish image assets to the ECR repository in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	ImageAssetPublishingRoleArn *string `json:"imageAssetPublishingRoleArn" yaml:"imageAssetPublishingRoleArn"`
	// Name of the ECR repository to hold Docker Image assets.
	//
	// You must supply this if you have given a non-standard name to the ECR repository.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	ImageAssetsRepositoryName *string `json:"imageAssetsRepositoryName" yaml:"imageAssetsRepositoryName"`
	// The role to use to look up values from the target AWS account during synthesis.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// External ID to use when assuming lookup role.
	LookupRoleExternalId *string `json:"lookupRoleExternalId" yaml:"lookupRoleExternalId"`
	// Qualifier to disambiguate multiple environments in the same account.
	//
	// You can use this and leave the other naming properties empty if you have deployed
	// the bootstrap environment with standard names but only differnet qualifiers.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
	// Use the bootstrapped lookup role for (read-only) stack operations.
	//
	// Use the lookup role when performing a `cdk diff`. If set to `false`, the
	// `deploy role` credentials will be used to perform a `cdk diff`.
	//
	// Requires bootstrap stack version 8.
	UseLookupRoleForStackOperations *bool `json:"useLookupRoleForStackOperations" yaml:"useLookupRoleForStackOperations"`
}

// Default resolver implementation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var fragmentConcatenator iFragmentConcatenator
//   defaultTokenResolver := cdk.NewDefaultTokenResolver(fragmentConcatenator)
//
type DefaultTokenResolver interface {
	ITokenResolver
	// Resolve a tokenized list.
	ResolveList(xs *[]*string, context IResolveContext) interface{}
	// Resolve string fragments to Tokens.
	ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{}
	// Default Token resolution.
	//
	// Resolve the Token, recurse into whatever it returns,
	// then finally post-process it.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy struct for DefaultTokenResolver
type jsiiProxy_DefaultTokenResolver struct {
	jsiiProxy_ITokenResolver
}

func NewDefaultTokenResolver(concat IFragmentConcatenator) DefaultTokenResolver {
	_init_.Initialize()

	j := jsiiProxy_DefaultTokenResolver{}

	_jsii_.Create(
		"aws-cdk-lib.DefaultTokenResolver",
		[]interface{}{concat},
		&j,
	)

	return &j
}

func NewDefaultTokenResolver_Override(d DefaultTokenResolver, concat IFragmentConcatenator) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DefaultTokenResolver",
		[]interface{}{concat},
		d,
	)
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveList(xs *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveList",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveString",
		[]interface{}{fragments, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Docker build options.
//
// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: lambda.code.fromAsset(jsii.String("/path/to/handler"), &assetOptions{
//   		bundling: &bundlingOptions{
//   			image: dockerImage.fromBuild(jsii.String("/path/to/dir/with/DockerFile"), &dockerBuildOptions{
//   				buildArgs: map[string]*string{
//   					"ARG1": jsii.String("value1"),
//   				},
//   			}),
//   			command: []*string{
//   				jsii.String("my"),
//   				jsii.String("cool"),
//   				jsii.String("command"),
//   			},
//   		},
//   	}),
//   	runtime: lambda.runtime_PYTHON_3_9(),
//   	handler: jsii.String("index.handler"),
//   })
//
type DockerBuildOptions struct {
	// Build args.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	File *string `json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	Platform *string `json:"platform" yaml:"platform"`
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   dockerIgnoreStrategy := cdk.NewDockerIgnoreStrategy(jsii.String("absoluteRootPath"), []*string{
//   	jsii.String("patterns"),
//   })
//
type DockerIgnoreStrategy interface {
	IgnoreStrategy
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for DockerIgnoreStrategy
type jsiiProxy_DockerIgnoreStrategy struct {
	jsiiProxy_IgnoreStrategy
}

func NewDockerIgnoreStrategy(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	j := jsiiProxy_DockerIgnoreStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.DockerIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		&j,
	)

	return &j
}

func NewDockerIgnoreStrategy_Override(d DockerIgnoreStrategy, absoluteRootPath *string, patterns *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		d,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func DockerIgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerIgnoreStrategy) Add(pattern *string) {
	_jsii_.InvokeVoid(
		d,
		"add",
		[]interface{}{pattern},
	)
}

func (d *jsiiProxy_DockerIgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

// A Docker image.
//
// Example:
//   entry := "/path/to/function"
//   image := dockerImage.fromBuild(entry)
//
//   lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
//   	entry: jsii.String(entry),
//   	runtime: runtime_PYTHON_3_8(),
//   	bundling: &bundlingOptions{
//   		buildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
type DockerImage interface {
	// The Docker image.
	Image() *string
	// Copies a file or directory out of the Docker image to the local filesystem.
	//
	// If `outputPath` is omitted the destination path is a temporary directory.
	//
	// Returns: the destination path.
	Cp(imagePath *string, outputPath *string) *string
	// Runs a Docker image.
	Run(options *DockerRunOptions)
	// Provides a stable representation of this image for JSON serialization.
	//
	// Returns: The overridden image name if set or image hash name in that order.
	ToJSON() *string
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_DockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


func NewDockerImage(image *string, _imageHash *string) DockerImage {
	_init_.Initialize()

	j := jsiiProxy_DockerImage{}

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		&j,
	)

	return &j
}

func NewDockerImage_Override(d DockerImage, image *string, _imageHash *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		d,
	)
}

// Builds a Docker image.
func DockerImage_FromBuild(path *string, options *DockerBuildOptions) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func DockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Cp(imagePath *string, outputPath *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"cp",
		[]interface{}{imagePath, outputPath},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Run(options *DockerRunOptions) {
	_jsii_.InvokeVoid(
		d,
		"run",
		[]interface{}{options},
	)
}

func (d *jsiiProxy_DockerImage) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The location of the published docker image.
//
// This is where the image can be
// consumed at runtime.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   dockerImageAssetLocation := &dockerImageAssetLocation{
//   	imageUri: jsii.String("imageUri"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type DockerImageAssetLocation struct {
	// The URI of the image in Amazon ECR.
	ImageUri *string `json:"imageUri" yaml:"imageUri"`
	// The name of the ECR repository.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   dockerImageAssetSource := &dockerImageAssetSource{
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	directoryName: jsii.String("directoryName"),
//   	dockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	dockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	dockerFile: jsii.String("dockerFile"),
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	networkMode: jsii.String("networkMode"),
//   }
//
type DockerImageAssetSource struct {
	// The hash of the contents of the docker build context.
	//
	// This hash is used
	// throughout the system to identify this image and avoid duplicate work
	// in case the source did not change.
	//
	// NOTE: this means that if you wish to update your docker image, you
	// must make a modification to the source (e.g. add some metadata to your Dockerfile).
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
	// The directory where the Dockerfile is stored, must be relative to the cloud assembly root.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	//
	// Only allowed when `directoryName` is specified.
	DockerBuildArgs *map[string]*string `json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Docker target to build to.
	//
	// Only allowed when `directoryName` is specified.
	DockerBuildTarget *string `json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// Path to the Dockerfile (relative to the directory).
	//
	// Only allowed when `directoryName` is specified.
	DockerFile *string `json:"dockerFile" yaml:"dockerFile"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the name of a local Docker image on `stdout`.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	NetworkMode *string `json:"networkMode" yaml:"networkMode"`
}

// Docker run options.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   dockerRunOptions := &dockerRunOptions{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	securityOpt: jsii.String("securityOpt"),
//   	user: jsii.String("user"),
//   	volumes: []dockerVolume{
//   		&dockerVolume{
//   			containerPath: jsii.String("containerPath"),
//   			hostPath: jsii.String("hostPath"),
//
//   			// the properties below are optional
//   			consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   		},
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type DockerRunOptions struct {
	// The command to run in the container.
	Command *[]*string `json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	Entrypoint *[]*string `json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	SecurityOpt *string `json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	User *string `json:"user" yaml:"user"`
	// Docker volumes to mount.
	Volumes *[]*DockerVolume `json:"volumes" yaml:"volumes"`
	// Working directory inside the container.
	WorkingDirectory *string `json:"workingDirectory" yaml:"workingDirectory"`
}

// A Docker volume.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   dockerVolume := &dockerVolume{
//   	containerPath: jsii.String("containerPath"),
//   	hostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   }
//
type DockerVolume struct {
	// The path where the file or directory is mounted in the container.
	ContainerPath *string `json:"containerPath" yaml:"containerPath"`
	// The path to the file or directory on the host machine.
	HostPath *string `json:"hostPath" yaml:"hostPath"`
	// Mount consistency.
	//
	// Only applicable for macOS.
	// See: https://docs.docker.com/storage/bind-mounts/#configure-mount-consistency-for-macos
	//
	Consistency DockerVolumeConsistency `json:"consistency" yaml:"consistency"`
}

// Supported Docker volume consistency types.
//
// Only valid on macOS due to the way file storage works on Mac.
type DockerVolumeConsistency string

const (
	// Read/write operations inside the Docker container are applied immediately on the mounted host machine volumes.
	DockerVolumeConsistency_CONSISTENT DockerVolumeConsistency = "CONSISTENT"
	// Read/write operations on mounted Docker volumes are first written inside the container and then synchronized to the host machine.
	DockerVolumeConsistency_DELEGATED DockerVolumeConsistency = "DELEGATED"
	// Read/write operations on mounted Docker volumes are first applied on the host machine and then synchronized to the container.
	DockerVolumeConsistency_CACHED DockerVolumeConsistency = "CACHED"
)

// Represents a length of time.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative, or as an unresolved number token.
//
// When the amount is passed as a token, unit conversion is not possible.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
//   	deadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	maxEventAge: cdk.duration.hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	retryAttempts: jsii.Number(2),
//   }))
//
type Duration interface {
	// Returns stringified number of duration.
	FormatTokenToNumber() *string
	// Checks if duration is a token or a resolvable object.
	IsUnresolved() *bool
	// Substract two Durations together.
	Minus(rhs Duration) Duration
	// Add two Durations together.
	Plus(rhs Duration) Duration
	// Return the total number of days in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Days.
	ToDays(opts *TimeConversionOptions) *float64
	// Return the total number of hours in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Hours.
	ToHours(opts *TimeConversionOptions) *float64
	// Turn this duration into a human-readable string.
	ToHumanString() *string
	// Return an ISO 8601 representation of this period.
	//
	// Returns: a string starting with 'P' describing the period.
	// See: https://www.iso.org/fr/standard/70907.html
	//
	ToIsoString() *string
	// Return the total number of milliseconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Milliseconds.
	ToMilliseconds(opts *TimeConversionOptions) *float64
	// Return the total number of minutes in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Minutes.
	ToMinutes(opts *TimeConversionOptions) *float64
	// Return the total number of seconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Seconds.
	ToSeconds(opts *TimeConversionOptions) *float64
	// Returns a string representation of this `Duration`.
	//
	// This is is never the right function to use when you want to use the `Duration`
	// object in a template. Use `toSeconds()`, `toMinutes()`, `toDays()`, etc. instead.
	ToString() *string
	// Returns unit of the duration.
	UnitLabel() *string
}

// The jsii proxy struct for Duration
type jsiiProxy_Duration struct {
	_ byte // padding
}

// Create a Duration representing an amount of days.
//
// Returns: a new `Duration` representing `amount` Days.
func Duration_Days(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"days",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of hours.
//
// Returns: a new `Duration` representing `amount` Hours.
func Duration_Hours(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"hours",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of milliseconds.
//
// Returns: a new `Duration` representing `amount` ms.
func Duration_Millis(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of minutes.
//
// Returns: a new `Duration` representing `amount` Minutes.
func Duration_Minutes(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"minutes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Parse a period formatted according to the ISO 8601 standard.
//
// Returns: the parsed `Duration`.
// See: https://www.iso.org/fr/standard/70907.html
//
func Duration_Parse(duration *string) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"parse",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of seconds.
//
// Returns: a new `Duration` representing `amount` Seconds.
func Duration_Seconds(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"seconds",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) FormatTokenToNumber() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"formatTokenToNumber",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) IsUnresolved() *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"isUnresolved",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) Minus(rhs Duration) Duration {
	var returns Duration

	_jsii_.Invoke(
		d,
		"minus",
		[]interface{}{rhs},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) Plus(rhs Duration) Duration {
	var returns Duration

	_jsii_.Invoke(
		d,
		"plus",
		[]interface{}{rhs},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToDays(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toDays",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToHours(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toHours",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToHumanString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toHumanString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToIsoString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toIsoString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToMilliseconds(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMilliseconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToMinutes(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMinutes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToSeconds(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toSeconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) UnitLabel() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"unitLabel",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to string encodings.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   encodingOptions := &encodingOptions{
//   	displayHint: jsii.String("displayHint"),
//   }
//
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	DisplayHint *string `json:"displayHint" yaml:"displayHint"`
}

// The deployment environment for a stack.
//
// Example:
//   // Passing a replication bucket created in a different stack.
//   app := NewApp()
//   replicationStack := NewStack(app, jsii.String("ReplicationStack"), &stackProps{
//   	env: &environment{
//   		region: jsii.String("us-west-1"),
//   	},
//   })
//   key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
//   replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &bucketProps{
//   	// like was said above - replication buckets need a set physical name
//   	bucketName: physicalName_GENERATE_IF_NEEDED(),
//   	encryptionKey: key,
//   })
//
//   // later...
//   // later...
//   codepipeline.NewPipeline(replicationStack, jsii.String("Pipeline"), &pipelineProps{
//   	crossRegionReplicationBuckets: map[string]iBucket{
//   		"us-west-1": replicationBucket,
//   	},
//   })
//
type Environment struct {
	// The AWS account ID for this environment.
	//
	// This can be either a concrete value such as `585191031104` or `Aws.accountId` which
	// indicates that account ID will only be determined during deployment (it
	// will resolve to the CloudFormation intrinsic `{"Ref":"AWS::AccountId"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concerete region information and
	// will cause this stack to emit synthesis errors.
	Account *string `json:"account" yaml:"account"`
	// The AWS region for this environment.
	//
	// This can be either a concrete value such as `eu-west-2` or `Aws.region`
	// which indicates that account ID will only be determined during deployment
	// (it will resolve to the CloudFormation intrinsic `{"Ref":"AWS::Region"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concerete region information and
	// will cause this stack to emit synthesis errors.
	Region *string `json:"region" yaml:"region"`
}

// Represents a date of expiration.
//
// The amount can be specified either as a Date object, timestamp, Duration or string.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   expiration := cdk.expiration.after(cdk.duration.minutes(jsii.Number(30)))
//
type Expiration interface {
	// Expiration value as a Date object.
	Date() *time.Time
	// Check if Exipiration expires after input.
	IsAfter(t Duration) *bool
	// Check if Exipiration expires before input.
	IsBefore(t Duration) *bool
	// Exipration Value in a formatted Unix Epoch Time in seconds.
	ToEpoch() *float64
}

// The jsii proxy struct for Expiration
type jsiiProxy_Expiration struct {
	_ byte // padding
}

func (j *jsiiProxy_Expiration) Date() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}


// Expire once the specified duration has passed since deployment time.
func Expiration_After(t Duration) Expiration {
	_init_.Initialize()

	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"after",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at the specified date.
func Expiration_AtDate(d *time.Time) Expiration {
	_init_.Initialize()

	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"atDate",
		[]interface{}{d},
		&returns,
	)

	return returns
}

// Expire at the specified timestamp.
func Expiration_AtTimestamp(t *float64) Expiration {
	_init_.Initialize()

	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"atTimestamp",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at specified date, represented as a string.
func Expiration_FromString(s *string) Expiration {
	_init_.Initialize()

	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) IsAfter(t Duration) *bool {
	var returns *bool

	_jsii_.Invoke(
		e,
		"isAfter",
		[]interface{}{t},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) IsBefore(t Duration) *bool {
	var returns *bool

	_jsii_.Invoke(
		e,
		"isBefore",
		[]interface{}{t},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) ToEpoch() *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"toEpoch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for the `stack.exportValue()` method.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   exportValueOptions := &exportValueOptions{
//   	name: jsii.String("name"),
//   }
//
type ExportValueOptions struct {
	// The name of the export to create.
	Name *string `json:"name" yaml:"name"`
}

// Features that are implemented behind a flag in order to preserve backwards compatibility for existing apps.
//
// The list of flags are available in the
// `@aws-cdk/cx-api` module.
//
// The state of the flag for this application is stored as a CDK context variable.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   featureFlags := cdk.featureFlags.of(this)
//
type FeatureFlags interface {
	// Check whether a feature flag is enabled.
	//
	// If configured, the flag is present in
	// the construct node context. Falls back to the defaults defined in the `cx-api`
	// module.
	IsEnabled(featureFlag *string) *bool
}

// The jsii proxy struct for FeatureFlags
type jsiiProxy_FeatureFlags struct {
	_ byte // padding
}

// Inspect feature flags on the construct node's context.
func FeatureFlags_Of(scope constructs.IConstruct) FeatureFlags {
	_init_.Initialize()

	var returns FeatureFlags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.FeatureFlags",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlags) IsEnabled(featureFlag *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		f,
		"isEnabled",
		[]interface{}{featureFlag},
		&returns,
	)

	return returns
}

// The location of the published file asset.
//
// This is where the asset
// can be consumed at runtime.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fileAssetLocation := &fileAssetLocation{
//   	bucketName: jsii.String("bucketName"),
//   	httpUrl: jsii.String("httpUrl"),
//   	objectKey: jsii.String("objectKey"),
//   	s3ObjectUrl: jsii.String("s3ObjectUrl"),
//
//   	// the properties below are optional
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	s3ObjectUrlWithPlaceholders: jsii.String("s3ObjectUrlWithPlaceholders"),
//   }
//
type FileAssetLocation struct {
	// The name of the Amazon S3 bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The HTTP URL of this asset on Amazon S3.
	//
	// This value suitable for inclusion in a CloudFormation template, and
	// may be an encoded token.
	//
	// Example value: `https://s3-us-east-1.amazonaws.com/mybucket/myobject`
	HttpUrl *string `json:"httpUrl" yaml:"httpUrl"`
	// The Amazon S3 object key.
	ObjectKey *string `json:"objectKey" yaml:"objectKey"`
	// The S3 URL of this asset on Amazon S3.
	//
	// This value suitable for inclusion in a CloudFormation template, and
	// may be an encoded token.
	//
	// Example value: `s3://mybucket/myobject`.
	S3ObjectUrl *string `json:"s3ObjectUrl" yaml:"s3ObjectUrl"`
	// The ARN of the KMS key used to encrypt the file asset bucket, if any.
	//
	// The CDK bootstrap stack comes with a key policy that does not require
	// setting this property, so you only need to set this property if you
	// have customized the bootstrap stack to require it.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Like `s3ObjectUrl`, but not suitable for CloudFormation consumption.
	//
	// If there are placeholders in the S3 URL, they will be returned unreplaced
	// and un-evaluated.
	S3ObjectUrlWithPlaceholders *string `json:"s3ObjectUrlWithPlaceholders" yaml:"s3ObjectUrlWithPlaceholders"`
}

// Packaging modes for file assets.
type FileAssetPackaging string

const (
	// The asset source path points to a directory, which should be archived using zip and and then uploaded to Amazon S3.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
	// The asset source path points to a single file, which should be uploaded to Amazon S3.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
)

// Represents the source for a file asset.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fileAssetSource := &fileAssetSource{
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	fileName: jsii.String("fileName"),
//   	packaging: cdk.fileAssetPackaging_ZIP_DIRECTORY,
//   }
//
type FileAssetSource struct {
	// A hash on the content source.
	//
	// This hash is used to uniquely identify this
	// asset throughout the system. If this value doesn't change, the asset will
	// not be rebuilt or republished.
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the location of a ZIP file on `stdout`.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// The path, relative to the root of the cloud assembly, in which this asset source resides.
	//
	// This can be a path to a file or a directory, depending on the
	// packaging type.
	FileName *string `json:"fileName" yaml:"fileName"`
	// Which type of packaging to perform.
	Packaging FileAssetPackaging `json:"packaging" yaml:"packaging"`
}

// Options applied when copying directories into the staging location.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fileCopyOptions := &fileCopyOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type FileCopyOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
}

// Options related to calculating source hash.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fileFingerprintOptions := &fileFingerprintOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type FileFingerprintOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
}

// File system utilities.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fileSystem := cdk.NewFileSystem()
//
type FileSystem interface {
}

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	_ byte // padding
}

func NewFileSystem() FileSystem {
	_init_.Initialize()

	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.FileSystem",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFileSystem_Override(f FileSystem) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.FileSystem",
		nil, // no parameters
		f,
	)
}

// Copies an entire directory structure.
func FileSystem_CopyDirectory(srcDir *string, destDir *string, options *CopyOptions, rootDir *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.FileSystem",
		"copyDirectory",
		[]interface{}{srcDir, destDir, options, rootDir},
	)
}

// Produces fingerprint based on the contents of a single file or an entire directory tree.
//
// The fingerprint will also include:
// 1. An extra string if defined in `options.extra`.
// 2. The set of exclude patterns, if defined in `options.exclude`
// 3. The symlink follow mode value.
func FileSystem_Fingerprint(fileOrDirectory *string, options *FingerprintOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.FileSystem",
		"fingerprint",
		[]interface{}{fileOrDirectory, options},
		&returns,
	)

	return returns
}

// Checks whether a directory is empty.
func FileSystem_IsEmpty(dir *string) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.FileSystem",
		"isEmpty",
		[]interface{}{dir},
		&returns,
	)

	return returns
}

// Creates a unique temporary directory in the **system temp directory**.
func FileSystem_Mkdtemp(prefix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.FileSystem",
		"mkdtemp",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func FileSystem_Tmpdir() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.FileSystem",
		"tmpdir",
		&returns,
	)
	return returns
}

// Options related to calculating source hash.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   fingerprintOptions := &fingerprintOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type FingerprintOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	Follow SymlinkFollowMode `json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
}

// CloudFormation intrinsic functions.
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//   portfolio.constrainCloudFormationParameters(product, &cloudFormationRuleConstraintOptions{
//   	rule: &templateRule{
//   		ruleName: jsii.String("testInstanceType"),
//   		condition: cdk.fn.conditionEquals(cdk.*fn.ref(jsii.String("Environment")), jsii.String("test")),
//   		assertions: []templateRuleAssertion{
//   			&templateRuleAssertion{
//   				assert: cdk.*fn.conditionContains([]*string{
//   					jsii.String("t2.micro"),
//   					jsii.String("t2.small"),
//   				}, cdk.*fn.ref(jsii.String("InstanceType"))),
//   				description: jsii.String("For test environment, the instance type should be small"),
//   			},
//   		},
//   	},
//   })
//
type Fn interface {
}

// The jsii proxy struct for Fn
type jsiiProxy_Fn struct {
	_ byte // padding
}

// The intrinsic function ``Fn::Base64`` returns the Base64 representation of the input string.
//
// This function is typically used to pass encoded data to
// Amazon EC2 instances by way of the UserData property.
//
// Returns: a token represented as a string.
func Fn_Base64(data *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"base64",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Cidr`` returns the specified Cidr address block.
//
// Returns: a token represented as a string.
func Fn_Cidr(ipBlock *string, count *float64, sizeMask *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"cidr",
		[]interface{}{ipBlock, count, sizeMask},
		&returns,
	)

	return returns
}

// Returns true if all the specified conditions evaluate to true, or returns false if any one of the conditions evaluates to false.
//
// ``Fn::And`` acts as
// an AND operator. The minimum number of conditions that you can include is
// 1.
//
// Returns: an FnCondition token.
func Fn_ConditionAnd(conditions ...ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionAnd",
		args,
		&returns,
	)

	return returns
}

// Returns true if a specified string matches at least one value in a list of strings.
//
// Returns: an FnCondition token.
func Fn_ConditionContains(listOfStrings *[]*string, value *string) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionContains",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

// Returns true if a specified string matches all values in a list.
//
// Returns: an FnCondition token.
func Fn_ConditionEachMemberEquals(listOfStrings *[]*string, value *string) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEachMemberEquals",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

// Returns true if each member in a list of strings matches at least one value in a second list of strings.
//
// Returns: an FnCondition token.
func Fn_ConditionEachMemberIn(stringsToCheck *[]*string, stringsToMatch *[]*string) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEachMemberIn",
		[]interface{}{stringsToCheck, stringsToMatch},
		&returns,
	)

	return returns
}

// Compares if two values are equal.
//
// Returns true if the two values are equal
// or false if they aren't.
//
// Returns: an FnCondition token.
func Fn_ConditionEquals(lhs interface{}, rhs interface{}) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEquals",
		[]interface{}{lhs, rhs},
		&returns,
	)

	return returns
}

// Returns one value if the specified condition evaluates to true and another value if the specified condition evaluates to false.
//
// Currently, AWS
// CloudFormation supports the ``Fn::If`` intrinsic function in the metadata
// attribute, update policy attribute, and property values in the Resources
// section and Outputs sections of a template. You can use the AWS::NoValue
// pseudo parameter as a return value to remove the corresponding property.
//
// Returns: an FnCondition token.
func Fn_ConditionIf(conditionId *string, valueIfTrue interface{}, valueIfFalse interface{}) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionIf",
		[]interface{}{conditionId, valueIfTrue, valueIfFalse},
		&returns,
	)

	return returns
}

// Returns true for a condition that evaluates to false or returns false for a condition that evaluates to true.
//
// ``Fn::Not`` acts as a NOT operator.
//
// Returns: an FnCondition token.
func Fn_ConditionNot(condition ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionNot",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

// Returns true if any one of the specified conditions evaluate to true, or returns false if all of the conditions evaluates to false.
//
// ``Fn::Or`` acts
// as an OR operator. The minimum number of conditions that you can include is
// 1.
//
// Returns: an FnCondition token.
func Fn_ConditionOr(conditions ...ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionOr",
		args,
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::FindInMap`` returns the value corresponding to keys in a two-level map that is declared in the Mappings section.
//
// Returns: a token represented as a string.
func Fn_FindInMap(mapName *string, topLevelKey *string, secondLevelKey *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"findInMap",
		[]interface{}{mapName, topLevelKey, secondLevelKey},
		&returns,
	)

	return returns
}

// The ``Fn::GetAtt`` intrinsic function returns the value of an attribute from a resource in the template.
//
// Returns: an IResolvable object.
func Fn_GetAtt(logicalNameOfResource *string, attributeName *string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"getAtt",
		[]interface{}{logicalNameOfResource, attributeName},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::GetAZs`` returns an array that lists Availability Zones for a specified region.
//
// Because customers have access to
// different Availability Zones, the intrinsic function ``Fn::GetAZs`` enables
// template authors to write templates that adapt to the calling user's
// access. That way you don't have to hard-code a full list of Availability
// Zones for a specified region.
//
// Returns: a token represented as a string array.
func Fn_GetAzs(region *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"getAzs",
		[]interface{}{region},
		&returns,
	)

	return returns
}

// Like `Fn.importValue`, but import a list with a known length.
//
// If you explicitly want a list with an unknown length, call `Fn.split(',',
// Fn.importValue(exportName))`. See the documentation of `Fn.split` to read
// more about the limitations of using lists of unknown length.
//
// `Fn.importListValue(exportName, assumedLength)` is the same as
// `Fn.split(',', Fn.importValue(exportName), assumedLength)`,
// but easier to read and impossible to forget to pass `assumedLength`.
func Fn_ImportListValue(sharedValueToImport *string, assumedLength *float64, delimiter *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"importListValue",
		[]interface{}{sharedValueToImport, assumedLength, delimiter},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::ImportValue`` returns the value of an output exported by another stack.
//
// You typically use this function to create
// cross-stack references. In the following example template snippets, Stack A
// exports VPC security group values and Stack B imports them.
//
// Returns: a token represented as a string.
func Fn_ImportValue(sharedValueToImport *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"importValue",
		[]interface{}{sharedValueToImport},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Join`` appends a set of values into a single value, separated by the specified delimiter.
//
// If a delimiter is the empty
// string, the set of values are concatenated with no delimiter.
//
// Returns: a token represented as a string.
func Fn_Join(delimiter *string, listOfValues *[]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"join",
		[]interface{}{delimiter, listOfValues},
		&returns,
	)

	return returns
}

// Given an url, parse the domain name.
func Fn_ParseDomainName(url *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"parseDomainName",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// The ``Ref`` intrinsic function returns the value of the specified parameter or resource.
//
// Note that it doesn't validate the logicalName, it mainly serves paremeter/resource reference defined in a ``CfnInclude`` template.
func Fn_Ref(logicalName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"ref",
		[]interface{}{logicalName},
		&returns,
	)

	return returns
}

// Returns all values for a specified parameter type.
//
// Returns: a token represented as a string array.
func Fn_RefAll(parameterType *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"refAll",
		[]interface{}{parameterType},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Select`` returns a single object from a list of objects by index.
//
// Returns: a token represented as a string.
func Fn_Select(index *float64, array *[]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"select",
		[]interface{}{index, array},
		&returns,
	)

	return returns
}

// Split a string token into a token list of string values.
//
// Specify the location of splits with a delimiter such as ',' (a comma).
// Renders to the `Fn::Split` intrinsic function.
//
// Lists with unknown lengths (default)
// -------------------------------------
//
// Since this function is used to work with deploy-time values, if `assumedLength`
// is not given the CDK cannot know the length of the resulting list at synthesis time.
// This brings the following restrictions:
//
// - You must use `Fn.select(i, list)` to pick elements out of the list (you must not use
//    `list[i]`).
// - You cannot add elements to the list, remove elements from the list,
//    combine two such lists together, or take a slice of the list.
// - You cannot pass the list to constructs that do any of the above.
//
// The only valid operation with such a tokenized list is to pass it unmodified to a
// CloudFormation Resource construct.
//
// Lists with assumed lengths
// --------------------------
//
// Pass `assumedLength` if you know the length of the list that will be
// produced by splitting. The actual list length at deploy time may be
// *longer* than the number you pass, but not *shorter*.
//
// The returned list will look like:
//
// ```
// [Fn.select(0, split), Fn.select(1, split), Fn.select(2, split), ...]
// ```
//
// The restrictions from the section "Lists with unknown lengths" will now be lifted,
// at the expense of having to know and fix the length of the list.
//
// Returns: a token represented as a string array.
func Fn_Split(delimiter *string, source *string, assumedLength *float64) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"split",
		[]interface{}{delimiter, source, assumedLength},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Sub`` substitutes variables in an input string with values that you specify.
//
// In your templates, you can use this function
// to construct commands or outputs that include values that aren't available
// until you create or update a stack.
//
// Returns: a token represented as a string.
func Fn_Sub(body *string, variables *map[string]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"sub",
		[]interface{}{body, variables},
		&returns,
	)

	return returns
}

// Creates a token representing the ``Fn::Transform`` expression.
//
// Returns: a token representing the transform expression.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-transform.html
//
func Fn_Transform(macroName *string, parameters *map[string]interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"transform",
		[]interface{}{macroName, parameters},
		&returns,
	)

	return returns
}

// Returns an attribute value or list of values for a specific parameter and attribute.
//
// Returns: a token represented as a string.
func Fn_ValueOf(parameterOrLogicalId *string, attribute *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"valueOf",
		[]interface{}{parameterOrLogicalId, attribute},
		&returns,
	)

	return returns
}

// Returns a list of all attribute values for a given parameter type and attribute.
//
// Returns: a token represented as a string array.
func Fn_ValueOfAll(parameterType *string, attribute *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"valueOfAll",
		[]interface{}{parameterType, attribute},
		&returns,
	)

	return returns
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//   getContextKeyOptions := &getContextKeyOptions{
//   	provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	includeEnvironment: jsii.Boolean(false),
//   	props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextKeyOptions struct {
	// The context provider to query.
	Provider *string `json:"provider" yaml:"provider"`
	// Whether to include the stack's account and region automatically.
	IncludeEnvironment *bool `json:"includeEnvironment" yaml:"includeEnvironment"`
	// Provider-specific properties.
	Props *map[string]interface{} `json:"props" yaml:"props"`
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//   getContextKeyResult := &getContextKeyResult{
//   	key: jsii.String("key"),
//   	props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextKeyResult struct {
	Key *string `json:"key" yaml:"key"`
	Props *map[string]interface{} `json:"props" yaml:"props"`
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dummyValue interface{}
//   var props interface{}
//   getContextValueOptions := &getContextValueOptions{
//   	dummyValue: dummyValue,
//   	provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	includeEnvironment: jsii.Boolean(false),
//   	props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextValueOptions struct {
	// The context provider to query.
	Provider *string `json:"provider" yaml:"provider"`
	// Whether to include the stack's account and region automatically.
	IncludeEnvironment *bool `json:"includeEnvironment" yaml:"includeEnvironment"`
	// Provider-specific properties.
	Props *map[string]interface{} `json:"props" yaml:"props"`
	// The value to return if the context value was not found and a missing context is reported.
	//
	// This should be a dummy value that should preferably
	// fail during deployment since it represents an invalid state.
	DummyValue interface{} `json:"dummyValue" yaml:"dummyValue"`
}

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//   getContextValueResult := &getContextValueResult{
//   	value: value,
//   }
//
type GetContextValueResult struct {
	Value interface{} `json:"value" yaml:"value"`
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   gitIgnoreStrategy := cdk.NewGitIgnoreStrategy(jsii.String("absoluteRootPath"), []*string{
//   	jsii.String("patterns"),
//   })
//
type GitIgnoreStrategy interface {
	IgnoreStrategy
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for GitIgnoreStrategy
type jsiiProxy_GitIgnoreStrategy struct {
	jsiiProxy_IgnoreStrategy
}

func NewGitIgnoreStrategy(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	j := jsiiProxy_GitIgnoreStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.GitIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		&j,
	)

	return &j
}

func NewGitIgnoreStrategy_Override(g GitIgnoreStrategy, absoluteRootPath *string, patterns *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.GitIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		g,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func GitIgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitIgnoreStrategy) Add(pattern *string) {
	_jsii_.InvokeVoid(
		g,
		"add",
		[]interface{}{pattern},
	)
}

func (g *jsiiProxy_GitIgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		g,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   globIgnoreStrategy := cdk.NewGlobIgnoreStrategy(jsii.String("absoluteRootPath"), []*string{
//   	jsii.String("patterns"),
//   })
//
type GlobIgnoreStrategy interface {
	IgnoreStrategy
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for GlobIgnoreStrategy
type jsiiProxy_GlobIgnoreStrategy struct {
	jsiiProxy_IgnoreStrategy
}

func NewGlobIgnoreStrategy(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	j := jsiiProxy_GlobIgnoreStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.GlobIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		&j,
	)

	return &j
}

func NewGlobIgnoreStrategy_Override(g GlobIgnoreStrategy, absoluteRootPath *string, patterns *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.GlobIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		g,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func GlobIgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GlobIgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func GlobIgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GlobIgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func GlobIgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GlobIgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func GlobIgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GlobIgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobIgnoreStrategy) Add(pattern *string) {
	_jsii_.InvokeVoid(
		g,
		"add",
		[]interface{}{pattern},
	)
}

func (g *jsiiProxy_GlobIgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		g,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

// Interface for lazy untyped value producers.
type IAnyProducer interface {
	// Produce the value.
	Produce(context IResolveContext) interface{}
}

// The jsii proxy for IAnyProducer
type jsiiProxy_IAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnyProducer) Produce(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Represents an Aspect.
type IAspect interface {
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy for IAspect
type jsiiProxy_IAspect struct {
	_ byte // padding
}

func (i *jsiiProxy_IAspect) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

// Common interface for all assets.
type IAsset interface {
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
}

// The jsii proxy for IAsset
type jsiiProxy_IAsset struct {
	_ byte // padding
}

func (j *jsiiProxy_IAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

// Represents a CloudFormation element that can be used within a Condition.
//
// You can use intrinsic functions, such as ``Fn.conditionIf``,
// ``Fn.conditionEquals``, and ``Fn.conditionNot``, to conditionally create
// stack resources. These conditions are evaluated based on input parameters
// that you declare when you create or update a stack. After you define all your
// conditions, you can associate them with resources or resource properties in
// the Resources and Outputs sections of a template.
//
// You define all conditions in the Conditions section of a template except for
// ``Fn.conditionIf`` conditions. You can use the ``Fn.conditionIf`` condition
// in the metadata attribute, update policy attribute, and property values in
// the Resources section and Outputs sections of a template.
//
// You might use conditions when you want to reuse a template that can create
// resources in different contexts, such as a test environment versus a
// production environment. In your template, you can add an EnvironmentType
// input parameter, which accepts either prod or test as inputs. For the
// production environment, you might include Amazon EC2 instances with certain
// capabilities; however, for the test environment, you want to use less
// capabilities to save costs. With conditions, you can define which resources
// are created and how they're configured for each environment type.
//
// You can use `toString` when you wish to embed a condition expression
// in a property value that accepts a `string`. For example:
//
// ```ts
// new sqs.Queue(this, 'MyQueue', {
//    queueName: Fn.conditionIf('Condition', 'Hello', 'World').toString()
// });
// ```.
type ICfnConditionExpression interface {
	IResolvable
}

// The jsii proxy for ICfnConditionExpression
type jsiiProxy_ICfnConditionExpression struct {
	jsiiProxy_IResolvable
}

type ICfnResourceOptions interface {
	// A condition to associate with this resource.
	//
	// This means that only if the condition evaluates to 'true' when the stack
	// is deployed, the resource will be included. This is provided to allow CDK projects to produce legacy templates, but noramlly
	// there is no need to use it in CDK projects.
	Condition() CfnCondition
	SetCondition(c CfnCondition)
	// Associate the CreationPolicy attribute with a resource to prevent its status from reaching create complete until AWS CloudFormation receives a specified number of success signals or the timeout period is exceeded.
	//
	// To signal a
	// resource, you can use the cfn-signal helper script or SignalResource API. AWS CloudFormation publishes valid signals
	// to the stack events so that you track the number of signals sent.
	CreationPolicy() *CfnCreationPolicy
	SetCreationPolicy(c *CfnCreationPolicy)
	// With the DeletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted.
	//
	// You specify a DeletionPolicy attribute for each resource that you want to control. If a resource has no DeletionPolicy
	// attribute, AWS CloudFormation deletes the resource by default. Note that this capability also applies to update operations
	// that lead to resources being removed.
	DeletionPolicy() CfnDeletionPolicy
	SetDeletionPolicy(d CfnDeletionPolicy)
	// The description of this resource.
	//
	// Used for informational purposes only, is not processed in any way
	// (and stays with the CloudFormation template, is not passed to the underlying resource,
	// even if it does have a 'description' property).
	Description() *string
	SetDescription(d *string)
	// Metadata associated with the CloudFormation resource.
	//
	// This is not the same as the construct metadata which can be added
	// using construct.addMetadata(), but would not appear in the CloudFormation template automatically.
	Metadata() *map[string]interface{}
	SetMetadata(m *map[string]interface{})
	// Use the UpdatePolicy attribute to specify how AWS CloudFormation handles updates to the AWS::AutoScaling::AutoScalingGroup resource.
	//
	// AWS CloudFormation invokes one of three update policies depending on the type of change you make or whether a
	// scheduled action is associated with the Auto Scaling group.
	UpdatePolicy() *CfnUpdatePolicy
	SetUpdatePolicy(u *CfnUpdatePolicy)
	// Use the UpdateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy() CfnDeletionPolicy
	SetUpdateReplacePolicy(u CfnDeletionPolicy)
	// The version of this resource.
	//
	// Used only for custom CloudFormation resources.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
	//
	Version() *string
	SetVersion(v *string)
}

// The jsii proxy for ICfnResourceOptions
type jsiiProxy_ICfnResourceOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ICfnResourceOptions) Condition() CfnCondition {
	var returns CfnCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetCondition(val CfnCondition) {
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) CreationPolicy() *CfnCreationPolicy {
	var returns *CfnCreationPolicy
	_jsii_.Get(
		j,
		"creationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetCreationPolicy(val *CfnCreationPolicy) {
	_jsii_.Set(
		j,
		"creationPolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) DeletionPolicy() CfnDeletionPolicy {
	var returns CfnDeletionPolicy
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetDeletionPolicy(val CfnDeletionPolicy) {
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Metadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetMetadata(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) UpdatePolicy() *CfnUpdatePolicy {
	var returns *CfnUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetUpdatePolicy(val *CfnUpdatePolicy) {
	_jsii_.Set(
		j,
		"updatePolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) UpdateReplacePolicy() CfnDeletionPolicy {
	var returns CfnDeletionPolicy
	_jsii_.Get(
		j,
		"updateReplacePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetUpdateReplacePolicy(val CfnDeletionPolicy) {
	_jsii_.Set(
		j,
		"updateReplacePolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Interface to specify certain functions as Service Catalog rule-specifc.
//
// These functions can only be used in ``Rules`` section of template.
type ICfnRuleConditionExpression interface {
	ICfnConditionExpression
	// This field is only needed to defeat TypeScript's structural typing.
	//
	// It is never used.
	Disambiguator() *bool
}

// The jsii proxy for ICfnRuleConditionExpression
type jsiiProxy_ICfnRuleConditionExpression struct {
	jsiiProxy_ICfnConditionExpression
}

func (j *jsiiProxy_ICfnRuleConditionExpression) Disambiguator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disambiguator",
		&returns,
	)
	return returns
}

// Function used to concatenate symbols in the target document language.
//
// Interface so it could potentially be exposed over jsii.
type IFragmentConcatenator interface {
	// Join the fragment on the left and on the right.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy for IFragmentConcatenator
type jsiiProxy_IFragmentConcatenator struct {
	_ byte // padding
}

func (i *jsiiProxy_IFragmentConcatenator) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Interface for examining a construct and exposing metadata.
type IInspectable interface {
	// Examines construct.
	Inspect(inspector TreeInspector)
}

// The jsii proxy for IInspectable
type jsiiProxy_IInspectable struct {
	_ byte // padding
}

func (i *jsiiProxy_IInspectable) Inspect(inspector TreeInspector) {
	_jsii_.InvokeVoid(
		i,
		"inspect",
		[]interface{}{inspector},
	)
}

// Interface for lazy list producers.
type IListProducer interface {
	// Produce the list value.
	Produce(context IResolveContext) *[]*string
}

// The jsii proxy for IListProducer
type jsiiProxy_IListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IListProducer) Produce(context IResolveContext) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Local bundling.
type ILocalBundling interface {
	// This method is called before attempting docker bundling to allow the bundler to be executed locally.
	//
	// If the local bundler exists, and bundling
	// was performed locally, return `true`. Otherwise, return `false`.
	TryBundle(outputDir *string, options *BundlingOptions) *bool
}

// The jsii proxy for ILocalBundling
type jsiiProxy_ILocalBundling struct {
	_ byte // padding
}

func (i *jsiiProxy_ILocalBundling) TryBundle(outputDir *string, options *BundlingOptions) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"tryBundle",
		[]interface{}{outputDir, options},
		&returns,
	)

	return returns
}

// Interface for lazy number producers.
type INumberProducer interface {
	// Produce the number value.
	Produce(context IResolveContext) *float64
}

// The jsii proxy for INumberProducer
type jsiiProxy_INumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_INumberProducer) Produce(context IResolveContext) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// A Token that can post-process the complete resolved value, after resolve() has recursed over it.
type IPostProcessor interface {
	// Process the completely resolved value, after full recursion/resolution has happened.
	PostProcess(input interface{}, context IResolveContext) interface{}
}

// The jsii proxy for IPostProcessor
type jsiiProxy_IPostProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IPostProcessor) PostProcess(input interface{}, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"postProcess",
		[]interface{}{input, context},
		&returns,
	)

	return returns
}

// Interface for values that can be resolvable later.
//
// Tokens are special objects that participate in synthesis.
type IResolvable interface {
	// Produce the Token's value at resolution time.
	Resolve(context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	ToString() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
}

// The jsii proxy for IResolvable
type jsiiProxy_IResolvable struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolvable) Resolve(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResolvable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolvable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

// Current resolution context for tokens.
type IResolveContext interface {
	// Use this postprocessor after the entire token structure has been resolved.
	RegisterPostProcessor(postProcessor IPostProcessor)
	// Resolve an inner object.
	Resolve(x interface{}, options *ResolveChangeContextOptions) interface{}
	// Path in the JSON document that is being constructed.
	DocumentPath() *[]*string
	// True when we are still preparing, false if we're rendering the final output.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	Scope() constructs.IConstruct
}

// The jsii proxy for IResolveContext
type jsiiProxy_IResolveContext struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolveContext) RegisterPostProcessor(postProcessor IPostProcessor) {
	_jsii_.InvokeVoid(
		i,
		"registerPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (i *jsiiProxy_IResolveContext) Resolve(x interface{}, options *ResolveChangeContextOptions) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{x, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolveContext) DocumentPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"documentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Preparing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preparing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

// Interface for the Resource construct.
type IResource interface {
	constructs.IConstruct
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *ResourceEnvironment
	// The stack in which this resource is defined.
	Stack() Stack
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IResource) ApplyRemovalPolicy(policy RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IResource) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface for (stable) lazy untyped value producers.
type IStableAnyProducer interface {
	// Produce the value.
	Produce() interface{}
}

// The jsii proxy for IStableAnyProducer
type jsiiProxy_IStableAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableAnyProducer) Produce() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface for (stable) lazy list producers.
type IStableListProducer interface {
	// Produce the list value.
	Produce() *[]*string
}

// The jsii proxy for IStableListProducer
type jsiiProxy_IStableListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableListProducer) Produce() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface for (stable) lazy number producers.
type IStableNumberProducer interface {
	// Produce the number value.
	Produce() *float64
}

// The jsii proxy for IStableNumberProducer
type jsiiProxy_IStableNumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableNumberProducer) Produce() *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface for (stable) lazy string producers.
type IStableStringProducer interface {
	// Produce the string value.
	Produce() *string
}

// The jsii proxy for IStableStringProducer
type jsiiProxy_IStableStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableStringProducer) Produce() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Encodes information how a certain Stack should be deployed.
type IStackSynthesizer interface {
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
}

// The jsii proxy for IStackSynthesizer
type jsiiProxy_IStackSynthesizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		i,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		i,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		i,
		"bind",
		[]interface{}{stack},
	)
}

func (i *jsiiProxy_IStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

// Interface for lazy string producers.
type IStringProducer interface {
	// Produce the string value.
	Produce(context IResolveContext) *string
}

// The jsii proxy for IStringProducer
type jsiiProxy_IStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStringProducer) Produce(context IResolveContext) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Represents a single session of synthesis.
//
// Passed into `Construct.synthesize()` methods.
type ISynthesisSession interface {
	// Cloud assembly builder.
	Assembly() cxapi.CloudAssemblyBuilder
	SetAssembly(a cxapi.CloudAssemblyBuilder)
	// The output directory for this synthesis session.
	Outdir() *string
	SetOutdir(o *string)
	// Whether the stack should be validated after synthesis to check for error metadata.
	ValidateOnSynth() *bool
	SetValidateOnSynth(v *bool)
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Assembly() cxapi.CloudAssemblyBuilder {
	var returns cxapi.CloudAssemblyBuilder
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SetAssembly(val cxapi.CloudAssemblyBuilder) {
	_jsii_.Set(
		j,
		"assembly",
		val,
	)
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SetOutdir(val *string) {
	_jsii_.Set(
		j,
		"outdir",
		val,
	)
}

func (j *jsiiProxy_ISynthesisSession) ValidateOnSynth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"validateOnSynth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SetValidateOnSynth(val *bool) {
	_jsii_.Set(
		j,
		"validateOnSynth",
		val,
	)
}

// Interface to implement tags.
type ITaggable interface {
	// TagManager to set, remove and format tags.
	Tags() TagManager
}

// The jsii proxy for ITaggable
type jsiiProxy_ITaggable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITaggable) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

// CloudFormation template options for a stack.
type ITemplateOptions interface {
	// Gets or sets the description of this stack.
	//
	// If provided, it will be included in the CloudFormation template's "Description" attribute.
	Description() *string
	SetDescription(d *string)
	// Metadata associated with the CloudFormation template.
	Metadata() *map[string]interface{}
	SetMetadata(m *map[string]interface{})
	// Gets or sets the AWSTemplateFormatVersion field of the CloudFormation template.
	TemplateFormatVersion() *string
	SetTemplateFormatVersion(t *string)
	// Gets or sets the top-level template transform(s) for this stack (e.g. `["AWS::Serverless-2016-10-31"]`).
	Transforms() *[]*string
	SetTransforms(t *[]*string)
}

// The jsii proxy for ITemplateOptions
type jsiiProxy_ITemplateOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ITemplateOptions) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) Metadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions) SetMetadata(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) TemplateFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions) SetTemplateFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"templateFormatVersion",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

// Interface to apply operation to tokens in a string.
//
// Interface so it can be exported via jsii.
type ITokenMapper interface {
	// Replace a single token.
	MapToken(t IResolvable) interface{}
}

// The jsii proxy for ITokenMapper
type jsiiProxy_ITokenMapper struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenMapper) MapToken(t IResolvable) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"mapToken",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// How to resolve tokens.
type ITokenResolver interface {
	// Resolve a tokenized list.
	ResolveList(l *[]*string, context IResolveContext) interface{}
	// Resolve a string with at least one stringified token in it.
	//
	// (May use concatenation).
	ResolveString(s TokenizedStringFragments, context IResolveContext) interface{}
	// Resolve a single token.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy for ITokenResolver
type jsiiProxy_ITokenResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenResolver) ResolveList(l *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveList",
		[]interface{}{l, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveString(s TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveString",
		[]interface{}{s, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Determines the ignore behavior to use.
type IgnoreMode string

const (
	// Ignores file paths based on simple glob patterns.
	//
	// This is the default for file assets.
	//
	// It is also the default for Docker image assets, unless the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_GLOB IgnoreMode = "GLOB"
	// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
	IgnoreMode_GIT IgnoreMode = "GIT"
	// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
	//
	// This is the default for Docker image assets if the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_DOCKER IgnoreMode = "DOCKER"
)

// Represents file path ignoring behavior.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   ignoreStrategy := cdk.ignoreStrategy.fromCopyOptions(&copyOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }, jsii.String("absoluteRootPath"))
//
type IgnoreStrategy interface {
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for IgnoreStrategy
type jsiiProxy_IgnoreStrategy struct {
	_ byte // padding
}

func NewIgnoreStrategy_Override(i IgnoreStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.IgnoreStrategy",
		nil, // no parameters
		i,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func IgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.IgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func IgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.IgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func IgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.IgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func IgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.IgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IgnoreStrategy) Add(pattern *string) {
	_jsii_.InvokeVoid(
		i,
		"add",
		[]interface{}{pattern},
	)
}

func (i *jsiiProxy_IgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

// Token subclass that represents values intrinsic to the target document language.
//
// WARNING: this class should not be externally exposed, but is currently visible
// because of a limitation of jsii (https://github.com/aws/jsii/issues/524).
//
// This class will disappear in a future release and should not be used.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//   intrinsic := cdk.NewIntrinsic(value, &intrinsicProps{
//   	stackTrace: jsii.Boolean(false),
//   })
//
type Intrinsic interface {
	IResolvable
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Produce the Token's value at resolution time.
	Resolve(_context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
}

// The jsii proxy struct for Intrinsic
type jsiiProxy_Intrinsic struct {
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_Intrinsic) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


func NewIntrinsic(value interface{}, options *IntrinsicProps) Intrinsic {
	_init_.Initialize()

	j := jsiiProxy_Intrinsic{}

	_jsii_.Create(
		"aws-cdk-lib.Intrinsic",
		[]interface{}{value, options},
		&j,
	)

	return &j
}

func NewIntrinsic_Override(i Intrinsic, value interface{}, options *IntrinsicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Intrinsic",
		[]interface{}{value, options},
		i,
	)
}

func (i *jsiiProxy_Intrinsic) NewError(message *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Intrinsic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Customization properties for an Intrinsic token.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   intrinsicProps := &intrinsicProps{
//   	stackTrace: jsii.Boolean(false),
//   }
//
type IntrinsicProps struct {
	// Capture the stack trace of where this token is created.
	StackTrace *bool `json:"stackTrace" yaml:"stackTrace"`
}

// Lazily produce a value.
//
// Can be used to return a string, list or numeric value whose actual value
// will only be calculated later, during synthesis.
type Lazy interface {
}

// The jsii proxy struct for Lazy
type jsiiProxy_Lazy struct {
	_ byte // padding
}

// Defer the one-time calculation of an arbitrarily typed value to synthesis time.
//
// Use this if you want to render an object to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// The inner function will only be invoked one time and cannot depend on
// resolution context.
func Lazy_Any(producer IStableAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"any",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a list value to synthesis time.
//
// Use this if you want to render a list to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `string[]` type and don't need
// the calculation to be deferred, use `Token.asList()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_List(producer IStableListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"list",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a number value to synthesis time.
//
// Use this if you want to render a number to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `number` type and don't need
// the calculation to be deferred, use `Token.asNumber()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_Number(producer IStableNumberProducer) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"number",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Defer the one-time calculation of a string value to synthesis time.
//
// Use this if you want to render a string to a template whose actual value depends on
// some state mutation that may happen after the construct has been created.
//
// If you are simply looking to force a value to a `string` type and don't need
// the calculation to be deferred, use `Token.asString()` instead.
//
// The inner function will only be invoked once, and the resolved value
// cannot depend on the Stack the Token is used in.
func Lazy_String(producer IStableStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"string",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of an untyped value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.any()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedAny(producer IAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedAny",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of a list value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.list()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedList(producer IListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedList",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Defer the calculation of a number value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.number()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedNumber(producer INumberProducer) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedNumber",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Defer the calculation of a string value to synthesis time.
//
// Use of this function is not recommended; unless you know you need it for sure, you
// probably don't. Use `Lazy.string()` instead.
//
// The inner function may be invoked multiple times during synthesis. You
// should only use this method if the returned value depends on variables
// that may change during the Aspect application phase of synthesis, or if
// the value depends on the Stack the value is being used in. Both of these
// cases are rare, and only ever occur for AWS Construct Library authors.
func Lazy_UncachedString(producer IStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Lazy",
		"uncachedString",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Options for creating lazy untyped tokens.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   lazyAnyValueOptions := &lazyAnyValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   	omitEmptyArray: jsii.Boolean(false),
//   }
//
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `json:"displayHint" yaml:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	OmitEmptyArray *bool `json:"omitEmptyArray" yaml:"omitEmptyArray"`
}

// Options for creating a lazy list token.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   lazyListValueOptions := &lazyListValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   	omitEmpty: jsii.Boolean(false),
//   }
//
type LazyListValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `json:"displayHint" yaml:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// Options for creating a lazy string token.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   lazyStringValueOptions := &lazyStringValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   }
//
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `json:"displayHint" yaml:"displayHint"`
}

// Use the CDK classic way of referencing assets.
//
// This synthesizer will generate CloudFormation parameters for every referenced
// asset, and use the CLI's current credentials to deploy the stack.
//
// - It does not support cross-account deployment (the CLI must have credentials
//    to the account you are trying to deploy to).
// - It cannot be used with **CDK Pipelines**. To deploy using CDK Pipelines,
//    you must use the `DefaultStackSynthesizer`.
// - Each asset will take up a CloudFormation Parameter in your template. Keep in
//    mind that there is a maximum of 200 parameters in a CloudFormation template.
//    To use determinstic asset locations instead, use `CliCredentialsStackSynthesizer`.
//
// Be aware that your CLI credentials must be valid for the duration of the
// entire deployment. If you are using session credentials, make sure the
// session lifetime is long enough.
//
// This is the only StackSynthesizer that supports customizing asset behavior
// by overriding `Stack.addFileAsset()` and `Stack.addDockerImageAsset()`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   legacyStackSynthesizer := cdk.NewLegacyStackSynthesizer()
//
type LegacyStackSynthesizer interface {
	StackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for LegacyStackSynthesizer
type jsiiProxy_LegacyStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func NewLegacyStackSynthesizer() LegacyStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_LegacyStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.LegacyStackSynthesizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLegacyStackSynthesizer_Override(l LegacyStackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.LegacyStackSynthesizer",
		nil, // no parameters
		l,
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		l,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		l,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		l,
		"bind",
		[]interface{}{stack},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		l,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LegacyStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// Functions for devising unique names for constructs.
//
// For example, those can be
// used to allocate unique physical names for resources.
type Names interface {
}

// The jsii proxy struct for Names
type jsiiProxy_Names struct {
	_ byte // padding
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// The identifier includes a human readable portion rendered
// from the path components and a hash suffix.
//
// TODO (v2): replace with API to use `constructs.Node`.
//
// Returns: a unique id based on the construct path.
func Names_NodeUniqueId(node constructs.Node) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Names",
		"nodeUniqueId",
		[]interface{}{node},
		&returns,
	)

	return returns
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// The identifier includes a human readable portion rendered
// from the path components and a hash suffix.
//
// Returns: a unique id based on the construct path.
func Names_UniqueId(construct constructs.Construct) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Names",
		"uniqueId",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// A CloudFormation nested stack.
//
// When you apply template changes to update a top-level stack, CloudFormation
// updates the top-level stack and initiates an update to its nested stacks.
// CloudFormation updates the resources of modified nested stacks, but does not
// update the resources of unmodified nested stacks.
//
// Furthermore, this stack will not be treated as an independent deployment
// artifact (won't be listed in "cdk list" or deployable through "cdk deploy"),
// but rather only synthesized as a template and uploaded as an asset to S3.
//
// Cross references of resource attributes between the parent stack and the
// nested stack will automatically be translated to stack parameters and
// outputs.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type App awscdk.App
//   type CfnOutput awscdk.CfnOutput
//   type NestedStack awscdk.NestedStack
//   type NestedStackProps awscdk.NestedStackProps
//   type Stack awscdk.Stackimport _ "github.com/aws-samples/dummy/constructs"import awscdk "github.com/aws/aws-cdk-go/awscdk"type Deployment awscdk.Deployment
//   type Method awscdk.Method
//   type MockIntegration awscdk.MockIntegration
//   type PassthroughBehavior awscdk.PassthroughBehavior
//   type RestApi awscdk.RestApi
//   type Stage awscdk.Stage
//
//   /**
//    * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
//    *
//    * The root stack 'RootStack' first defines a RestApi.
//    * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
//    * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
//    *
//    * To verify this worked, go to the APIGateway
//    */
//
//   type rootStack struct {
//   	stack
//   }
//
//   func newRootStack(scope construct) *rootStack {
//   	this := &rootStack{}
//   	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))
//
//   	restApi := NewRestApi(this, jsii.String("RestApi"), &restApiProps{
//   		deploy: jsii.Boolean(false),
//   	})
//   	restApi.root.addMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.restApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	NewCfnOutput(this, jsii.String("PetsURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.restApiId, this.region),
//   	})
//
//   	NewCfnOutput(this, jsii.String("BooksURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.restApiId, this.region),
//   	})
//   	return this
//   }
//
//   type resourceNestedStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	rootResourceId *string
//   }
//
//   type petsStack struct {
//   	nestedStack
//   	methods []method
//   }
//
//   func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
//   	this := &petsStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)
//
//   	api := restApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("pets")).addMethod(jsii.String("GET"), NewMockIntegration(&integrationOptions{
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: passthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type booksStack struct {
//   	nestedStack
//   	methods []*method
//   }
//
//   func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
//   	this := &booksStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)
//
//   	api := restApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("books")).addMethod(jsii.String("GET"), NewMockIntegration(&integrationOptions{
//   		integrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: *passthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []*methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type deployStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	methods []*method
//   }
//
//   type deployStack struct {
//   	nestedStack
//   }
//
//   func newDeployStack(scope construct, props deployStackProps) *deployStack {
//   	this := &deployStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)
//
//   	deployment := NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   		api: *restApi.fromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.node.addDependency(method)
//   		}
//   	}
//   	NewStage(this, jsii.String("Stage"), &stageProps{
//   		deployment: deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(NewApp())
//
type NestedStack interface {
	Stack
	// The AWS account into which this stack will be deployed.
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.account` when the stack is defined. This can
	//     either be a concerete account (e.g. `585695031111`) or the
	//     `Aws.accountId` token.
	// 3. `Aws.accountId`, which represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::AccountId" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.account)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **account-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	Account() *string
	// The ID of the cloud assembly artifact for this stack.
	ArtifactId() *string
	// Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack.
	//
	// If the stack is environment-agnostic (either account and/or region are
	// tokens), this property will return an array with 2 tokens that will resolve
	// at deploy-time to the first two availability zones returned from CloudFormation's
	// `Fn::GetAZs` intrinsic function.
	//
	// If they are not available in the context, returns a set of dummy values and
	// reports them as missing, and let the CLI resolve them by calling EC2
	// `DescribeAvailabilityZones` on the target environment.
	//
	// To specify a different strategy for selecting availability zones override this method.
	AvailabilityZones() *[]*string
	// Indicates whether the stack requires bundling or not.
	BundlingRequired() *bool
	// Return the stacks this stack depends on.
	Dependencies() *[]Stack
	// The environment coordinates in which this stack is deployed.
	//
	// In the form
	// `aws://account/region`. Use `stack.account` and `stack.region` to obtain
	// the specific values, no need to parse.
	//
	// You can use this value to determine if two stacks are targeting the same
	// environment.
	//
	// If either `stack.account` or `stack.region` are not concrete values (e.g.
	// `Aws.account` or `Aws.region`) the special strings `unknown-account` and/or
	// `unknown-region` will be used respectively to indicate this stack is
	// region/account-agnostic.
	Environment() *string
	// Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.
	Nested() *bool
	// If this is a nested stack, returns it's parent stack.
	NestedStackParent() Stack
	// If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.
	//
	// `undefined` for top-level (non-nested) stacks.
	NestedStackResource() CfnResource
	// The tree node.
	Node() constructs.Node
	// Returns the list of notification Amazon Resource Names (ARNs) for the current stack.
	NotificationArns() *[]*string
	// The partition in which this stack is defined.
	Partition() *string
	// The AWS region into which this stack will be deployed (e.g. `us-west-2`).
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.region` when the stack is defined. This can
	//     either be a concerete region (e.g. `us-west-2`) or the `Aws.region`
	//     token.
	// 3. `Aws.region`, which is represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::Region" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.region)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **region-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	Region() *string
	// An attribute that represents the ID of the stack.
	//
	// This is a context aware attribute:
	// - If this is referenced from the parent stack, it will return `{ "Ref": "LogicalIdOfNestedStackResource" }`.
	// - If this is referenced from the context of the nested stack, it will return `{ "Ref": "AWS::StackId" }`
	//
	// Example value: `arn:aws:cloudformation:us-east-2:123456789012:stack/mystack-mynestedstack-sggfrhxhum7w/f449b250-b969-11e0-a185-5081d0136786`.
	StackId() *string
	// An attribute that represents the name of the nested stack.
	//
	// This is a context aware attribute:
	// - If this is referenced from the parent stack, it will return a token that parses the name from the stack ID.
	// - If this is referenced from the context of the nested stack, it will return `{ "Ref": "AWS::StackName" }`
	//
	// Example value: `mystack-mynestedstack-sggfrhxhum7w`.
	StackName() *string
	// Synthesis method for this stack.
	Synthesizer() IStackSynthesizer
	// Tags to be applied to the stack.
	Tags() TagManager
	// The name of the CloudFormation template file emitted to the output directory during synthesis.
	//
	// Example value: `MyStack.template.json`
	TemplateFile() *string
	// Options for CloudFormation template (like version, transform, description).
	TemplateOptions() ITemplateOptions
	// Whether termination protection is enabled for this stack.
	TerminationProtection() *bool
	// The Amazon domain suffix for the region in which this stack is defined.
	UrlSuffix() *string
	// Add a dependency between this stack and another stack.
	//
	// This can be used to define dependencies between any two stacks within an
	// app, and also supports nested stacks.
	AddDependency(target Stack, reason *string)
	// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
	//
	// Duplicate values are removed when stack is synthesized.
	//
	// Example:
	//   declare const stack: Stack;
	//
	//   stack.addTransform('AWS::Serverless-2016-10-31')
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
	//
	AddTransform(transform *string)
	// Returns the naming scheme used to allocate logical IDs.
	//
	// By default, uses
	// the `HashedAddressingScheme` but this method can be overridden to customize
	// this behavior.
	//
	// In order to make sure logical IDs are unique and stable, we hash the resource
	// construct tree path (i.e. toplevel/secondlevel/.../myresource) and add it as
	// a suffix to the path components joined without a separator (CloudFormation
	// IDs only allow alphanumeric characters).
	//
	// The result will be:
	//
	//    <path.join('')><md5(path.join('/')>
	//      "human"      "hash"
	//
	// If the "human" part of the ID exceeds 240 characters, we simply trim it so
	// the total ID doesn't exceed CloudFormation's 255 character limit.
	//
	// We only take 8 characters from the md5 hash (0.000005 chance of collision).
	//
	// Special cases:
	//
	// - If the path only contains a single component (i.e. it's a top-level
	//    resource), we won't add the hash to it. The hash is not needed for
	//    disamiguation and also, it allows for a more straightforward migration an
	//    existing CloudFormation template to a CDK stack without logical ID changes
	//    (or renames).
	// - For aesthetic reasons, if the last components of the path are the same
	//    (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
	//    resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
	//    instead of `L1L2PipelinePipeline<HASH>`
	// - If a component is named "Default" it will be omitted from the path. This
	//    allows refactoring higher level abstractions around constructs without affecting
	//    the IDs of already deployed resources.
	// - If a component is named "Resource" it will be omitted from the user-visible
	//    path, but included in the hash. This reduces visual noise in the human readable
	//    part of the identifier.
	AllocateLogicalId(cfnElement CfnElement) *string
	// Create a CloudFormation Export for a value.
	//
	// Returns a string representing the corresponding `Fn.importValue()`
	// expression for this Export. You can control the name for the export by
	// passing the `name` option.
	//
	// If you don't supply a value for `name`, the value you're exporting must be
	// a Resource attribute (for example: `bucket.bucketName`) and it will be
	// given the same name as the automatic cross-stack reference that would be created
	// if you used the attribute in another Stack.
	//
	// One of the uses for this method is to *remove* the relationship between
	// two Stacks established by automatic cross-stack references. It will
	// temporarily ensure that the CloudFormation Export still exists while you
	// remove the reference from the consuming stack. After that, you can remove
	// the resource and the manual export.
	//
	// ## Example
	//
	// Here is how the process works. Let's say there are two stacks,
	// `producerStack` and `consumerStack`, and `producerStack` has a bucket
	// called `bucket`, which is referenced by `consumerStack` (perhaps because
	// an AWS Lambda Function writes into it, or something like that).
	//
	// It is not safe to remove `producerStack.bucket` because as the bucket is being
	// deleted, `consumerStack` might still be using it.
	//
	// Instead, the process takes two deployments:
	//
	// ### Deployment 1: break the relationship
	//
	// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
	//    stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
	//    remove the Lambda Function altogether).
	// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
	//    will make sure the CloudFormation Export continues to exist while the relationship
	//    between the two stacks is being broken.
	// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
	//
	// ### Deployment 2: remove the bucket resource
	//
	// - You are now free to remove the `bucket` resource from `producerStack`.
	// - Don't forget to remove the `exportValue()` call as well.
	// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
	ExportValue(exportedValue interface{}, options *ExportValueOptions) *string
	// Creates an ARN from components.
	//
	// If `partition`, `region` or `account` are not specified, the stack's
	// partition, region and account will be used.
	//
	// If any component is the empty string, an empty string will be inserted
	// into the generated ARN at the location that component corresponds to.
	//
	// The ARN will be formatted as follows:
	//
	//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}}{resource-name}
	//
	// The required ARN pieces that are omitted will be taken from the stack that
	// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
	// can be 'undefined'.
	FormatArn(components *ArnComponents) *string
	// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
	//
	// This method is called when a `CfnElement` is created and used to render the
	// initial logical identity of resources. Logical ID renames are applied at
	// this stage.
	//
	// This method uses the protected method `allocateLogicalId` to render the
	// logical ID for an element. To modify the naming scheme, extend the `Stack`
	// class and override this method.
	GetLogicalId(element CfnElement) *string
	// Look up a fact value for the given fact for the region of this stack.
	//
	// Will return a definite value only if the region of the current stack is resolved.
	// If not, a lookup map will be added to the stack and the lookup will be done at
	// CDK deployment time.
	//
	// What regions will be included in the lookup map is controlled by the
	// `@aws-cdk/core:target-partitions` context value: it must be set to a list
	// of partitions, and only regions from the given partitions will be included.
	// If no such context key is set, all regions will be included.
	//
	// This function is intended to be used by construct library authors. Application
	// builders can rely on the abstractions offered by construct libraries and do
	// not have to worry about regional facts.
	//
	// If `defaultValue` is not given, it is an error if the fact is unknown for
	// the given region.
	RegionalFact(factName *string, defaultValue *string) *string
	// Rename a generated logical identities.
	//
	// To modify the naming scheme strategy, extend the `Stack` class and
	// override the `allocateLogicalId` method.
	RenameLogicalId(oldId *string, newId *string)
	// Indicate that a context key was expected.
	//
	// Contains instructions which will be emitted into the cloud assembly on how
	// the key should be supplied.
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	// Resolve a tokenized value in the context of the current stack.
	Resolve(obj interface{}) interface{}
	// Assign a value to one of the nested stack parameters.
	SetParameter(name *string, value *string)
	// Splits the provided ARN into its components.
	//
	// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
	// and a Token representing a dynamic CloudFormation expression
	// (in which case the returned components will also be dynamic CloudFormation expressions,
	// encoded as Tokens).
	SplitArn(arn *string, arnFormat ArnFormat) *ArnComponents
	// Convert an object, potentially containing tokens, to a JSON string.
	ToJsonString(obj interface{}, space *float64) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NestedStack
type jsiiProxy_NestedStack struct {
	jsiiProxy_Stack
}

func (j *jsiiProxy_NestedStack) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) BundlingRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bundlingRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Dependencies() *[]Stack {
	var returns *[]Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) NestedStackParent() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) NestedStackResource() CfnResource {
	var returns CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Synthesizer() IStackSynthesizer {
	var returns IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) TemplateOptions() ITemplateOptions {
	var returns ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedStack) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


func NewNestedStack(scope constructs.Construct, id *string, props *NestedStackProps) NestedStack {
	_init_.Initialize()

	j := jsiiProxy_NestedStack{}

	_jsii_.Create(
		"aws-cdk-lib.NestedStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNestedStack_Override(n NestedStack, scope constructs.Construct, id *string, props *NestedStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.NestedStack",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NestedStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.NestedStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is an object of type `NestedStack`.
func NestedStack_IsNestedStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.NestedStack",
		"isNestedStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func NestedStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.NestedStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
func NestedStack_Of(construct constructs.IConstruct) Stack {
	_init_.Initialize()

	var returns Stack

	_jsii_.StaticInvoke(
		"aws-cdk-lib.NestedStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) AddDependency(target Stack, reason *string) {
	_jsii_.InvokeVoid(
		n,
		"addDependency",
		[]interface{}{target, reason},
	)
}

func (n *jsiiProxy_NestedStack) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		n,
		"addTransform",
		[]interface{}{transform},
	)
}

func (n *jsiiProxy_NestedStack) AllocateLogicalId(cfnElement CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) ExportValue(exportedValue interface{}, options *ExportValueOptions) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) FormatArn(components *ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) GetLogicalId(element CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) RegionalFact(factName *string, defaultValue *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"regionalFact",
		[]interface{}{factName, defaultValue},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		n,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

func (n *jsiiProxy_NestedStack) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		n,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

func (n *jsiiProxy_NestedStack) Resolve(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) SetParameter(name *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"setParameter",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NestedStack) SplitArn(arn *string, arnFormat ArnFormat) *ArnComponents {
	var returns *ArnComponents

	_jsii_.Invoke(
		n,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) ToJsonString(obj interface{}, space *float64) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization props for the `NestedStack` construct.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type App awscdk.App
//   type CfnOutput awscdk.CfnOutput
//   type NestedStack awscdk.NestedStack
//   type NestedStackProps awscdk.NestedStackProps
//   type Stack awscdk.Stackimport _ "github.com/aws-samples/dummy/constructs"import awscdk "github.com/aws/aws-cdk-go/awscdk"type Deployment awscdk.Deployment
//   type Method awscdk.Method
//   type MockIntegration awscdk.MockIntegration
//   type PassthroughBehavior awscdk.PassthroughBehavior
//   type RestApi awscdk.RestApi
//   type Stage awscdk.Stage
//
//   /**
//    * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
//    *
//    * The root stack 'RootStack' first defines a RestApi.
//    * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
//    * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
//    *
//    * To verify this worked, go to the APIGateway
//    */
//
//   type rootStack struct {
//   	stack
//   }
//
//   func newRootStack(scope construct) *rootStack {
//   	this := &rootStack{}
//   	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))
//
//   	restApi := NewRestApi(this, jsii.String("RestApi"), &restApiProps{
//   		deploy: jsii.Boolean(false),
//   	})
//   	restApi.root.addMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.restApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	NewCfnOutput(this, jsii.String("PetsURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.restApiId, this.region),
//   	})
//
//   	NewCfnOutput(this, jsii.String("BooksURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.restApiId, this.region),
//   	})
//   	return this
//   }
//
//   type resourceNestedStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	rootResourceId *string
//   }
//
//   type petsStack struct {
//   	nestedStack
//   	methods []method
//   }
//
//   func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
//   	this := &petsStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)
//
//   	api := restApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("pets")).addMethod(jsii.String("GET"), NewMockIntegration(&integrationOptions{
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: passthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type booksStack struct {
//   	nestedStack
//   	methods []*method
//   }
//
//   func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
//   	this := &booksStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)
//
//   	api := restApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("books")).addMethod(jsii.String("GET"), NewMockIntegration(&integrationOptions{
//   		integrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: *passthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []*methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type deployStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	methods []*method
//   }
//
//   type deployStack struct {
//   	nestedStack
//   }
//
//   func newDeployStack(scope construct, props deployStackProps) *deployStack {
//   	this := &deployStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)
//
//   	deployment := NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   		api: *restApi.fromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.node.addDependency(method)
//   		}
//   	}
//   	NewStage(this, jsii.String("Stage"), &stageProps{
//   		deployment: deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(NewApp())
//
type NestedStackProps struct {
	// The Simple Notification Service (SNS) topics to publish stack related events.
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding
	// to a parameter defined in the embedded template and a value representing
	// the value that you want to set for the parameter.
	//
	// The nested stack construct will automatically synthesize parameters in order
	// to bind references from the parent stack(s) into the nested stack.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Policy to apply when the nested stack is removed.
	//
	// The default is `Destroy`, because all Removal Policies of resources inside the
	// Nested Stack should already have been set correctly. You normally should
	// not need to set this value.
	RemovalPolicy RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// The length of time that CloudFormation waits for the nested stack to reach the CREATE_COMPLETE state.
	//
	// When CloudFormation detects that the nested stack has reached the
	// CREATE_COMPLETE state, it marks the nested stack resource as
	// CREATE_COMPLETE in the parent stack and resumes creating the parent stack.
	// If the timeout period expires before the nested stack reaches
	// CREATE_COMPLETE, CloudFormation marks the nested stack as failed and rolls
	// back both the nested stack and parent stack.
	Timeout Duration `json:"timeout" yaml:"timeout"`
}

// Synthesizer for a nested stack.
//
// Forwards all calls to the parent stack's synthesizer.
//
// This synthesizer is automatically used for `NestedStack` constructs.
// App builder do not need to use this class directly.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var stackSynthesizer stackSynthesizer
//   nestedStackSynthesizer := cdk.NewNestedStackSynthesizer(stackSynthesizer)
//
type NestedStackSynthesizer interface {
	StackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for NestedStackSynthesizer
type jsiiProxy_NestedStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func NewNestedStackSynthesizer(parentDeployment IStackSynthesizer) NestedStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_NestedStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.NestedStackSynthesizer",
		[]interface{}{parentDeployment},
		&j,
	)

	return &j
}

func NewNestedStackSynthesizer_Override(n NestedStackSynthesizer, parentDeployment IStackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.NestedStackSynthesizer",
		[]interface{}{parentDeployment},
		n,
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		n,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		n,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		n,
		"bind",
		[]interface{}{stack},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		n,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NestedStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// Includes special markers for automatic generation of physical names.
type PhysicalName interface {
}

// The jsii proxy struct for PhysicalName
type jsiiProxy_PhysicalName struct {
	_ byte // padding
}

func PhysicalName_GENERATE_IF_NEEDED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.PhysicalName",
		"GENERATE_IF_NEEDED",
		&returns,
	)
	return returns
}

// An intrinsic Token that represents a reference to a construct.
//
// References are recorded.
type Reference interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	DisplayName() *string
	Target() constructs.IConstruct
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Produce the Token's value at resolution time.
	Resolve(_context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
}

// The jsii proxy struct for Reference
type jsiiProxy_Reference struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_Reference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Reference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Reference) Target() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


func NewReference_Override(r Reference, value interface{}, target constructs.IConstruct, displayName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Reference",
		[]interface{}{value, target, displayName},
		r,
	)
}

// Check whether this is actually a Reference.
func Reference_IsReference(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Reference",
		"isReference",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) NewError(message *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Reference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Possible values for a resource's Removal Policy.
//
// The removal policy controls what happens to the resource if it stops being
// managed by CloudFormation. This can happen in one of three situations:
//
// - The resource is removed from the template, so CloudFormation stops managing it;
// - A change to the resource is made that requires it to be replaced, so CloudFormation stops
//    managing it;
// - The stack is deleted, so CloudFormation stops managing all resources in it.
//
// The Removal Policy applies to all above cases.
//
// Many stateful resources in the AWS Construct Library will accept a
// `removalPolicy` as a property, typically defaulting it to `RETAIN`.
//
// If the AWS Construct Library resource does not accept a `removalPolicy`
// argument, you can always configure it by using the escape hatch mechanism,
// as shown in the following example:
//
// ```ts
// declare const bucket: s3.Bucket;
//
// const cfnBucket = bucket.node.findChild('Resource') as CfnResource;
// cfnBucket.applyRemovalPolicy(RemovalPolicy.DESTROY);
// ```.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//
//   var api graphqlApi
//
//   user := iam.NewUser(this, jsii.String("User"))
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: opensearch.engineVersion_OPENSEARCH_1_2(),
//   	removalPolicy: removalPolicy_DESTROY,
//   	fineGrainedAccessControl: &advancedSecurityOptions{
//   		masterUserArn: user.userArn,
//   	},
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	enforceHttps: jsii.Boolean(true),
//   })
//   ds := api.addOpenSearchDataSource(jsii.String("ds"), domain)
//
//   ds.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getTests"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromString(jSON.stringify(map[string]interface{}{
//   		"version": jsii.String("2017-02-28"),
//   		"operation": jsii.String("GET"),
//   		"path": jsii.String("/id/post/_search"),
//   		"params": map[string]map[string]interface{}{
//   			"headers": map[string]interface{}{
//   			},
//   			"queryString": map[string]interface{}{
//   			},
//   			"body": map[string]*f64{
//   				"from": jsii.Number(0),
//   				"size": jsii.Number(50),
//   			},
//   		},
//   	})),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("[\n    #foreach($entry in $context.result.hits.hits)\n    #if( $velocityCount > 1 ) , #end\n    $utils.toJson($entry.get(\"_source\"))\n    #end\n  ]")),
//   })
//
type RemovalPolicy string

const (
	// This is the default removal policy.
	//
	// It means that when the resource is
	// removed from the app, it will be physically destroyed.
	RemovalPolicy_DESTROY RemovalPolicy = "DESTROY"
	// This uses the 'Retain' DeletionPolicy, which will cause the resource to be retained in the account, but orphaned from the stack.
	RemovalPolicy_RETAIN RemovalPolicy = "RETAIN"
	// This retention policy deletes the resource, but saves a snapshot of its data before deleting, so that it can be re-created later.
	//
	// Only available for some stateful resources,
	// like databases, EFS volumes, etc.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	RemovalPolicy_SNAPSHOT RemovalPolicy = "SNAPSHOT"
)

// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   removalPolicyOptions := &removalPolicyOptions{
//   	applyToUpdateReplacePolicy: jsii.Boolean(false),
//   	default: cdk.removalPolicy_DESTROY,
//   }
//
type RemovalPolicyOptions struct {
	// Apply the same deletion policy to the resource's "UpdateReplacePolicy".
	ApplyToUpdateReplacePolicy *bool `json:"applyToUpdateReplacePolicy" yaml:"applyToUpdateReplacePolicy"`
	// The default policy to apply in case the removal policy is not defined.
	Default RemovalPolicy `json:"default" yaml:"default"`
}

// The RemoveTag Aspect will handle removing tags from this node and children.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   removeTag := cdk.NewRemoveTag(jsii.String("key"), &tagProps{
//   	applyToLaunchedInstances: jsii.Boolean(false),
//   	excludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	includeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	priority: jsii.Number(123),
//   })
//
type RemoveTag interface {
	IAspect
	// The string key for the tag.
	Key() *string
	Props() *TagProps
	ApplyTag(resource ITaggable)
	// All aspects can visit an IConstruct.
	Visit(construct constructs.IConstruct)
}

// The jsii proxy struct for RemoveTag
type jsiiProxy_RemoveTag struct {
	jsiiProxy_IAspect
}

func (j *jsiiProxy_RemoveTag) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoveTag) Props() *TagProps {
	var returns *TagProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewRemoveTag(key *string, props *TagProps) RemoveTag {
	_init_.Initialize()

	j := jsiiProxy_RemoveTag{}

	_jsii_.Create(
		"aws-cdk-lib.RemoveTag",
		[]interface{}{key, props},
		&j,
	)

	return &j
}

func NewRemoveTag_Override(r RemoveTag, key *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.RemoveTag",
		[]interface{}{key, props},
		r,
	)
}

func (r *jsiiProxy_RemoveTag) ApplyTag(resource ITaggable) {
	_jsii_.InvokeVoid(
		r,
		"applyTag",
		[]interface{}{resource},
	)
}

func (r *jsiiProxy_RemoveTag) Visit(construct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		r,
		"visit",
		[]interface{}{construct},
	)
}

// Options that can be changed while doing a recursive resolve.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   resolveChangeContextOptions := &resolveChangeContextOptions{
//   	allowIntrinsicKeys: jsii.Boolean(false),
//   }
//
type ResolveChangeContextOptions struct {
	// Change the 'allowIntrinsicKeys' option.
	AllowIntrinsicKeys *bool `json:"allowIntrinsicKeys" yaml:"allowIntrinsicKeys"`
}

// Options to the resolve() operation.
//
// NOT the same as the ResolveContext; ResolveContext is exposed to Token
// implementors and resolution hooks, whereas this struct is just to bundle
// a number of things that would otherwise be arguments to resolve() in a
// readable way.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var tokenResolver iTokenResolver
//   resolveOptions := &resolveOptions{
//   	resolver: tokenResolver,
//   	scope: construct,
//
//   	// the properties below are optional
//   	preparing: jsii.Boolean(false),
//   	removeEmpty: jsii.Boolean(false),
//   }
//
type ResolveOptions struct {
	// The resolver to apply to any resolvable tokens found.
	Resolver ITokenResolver `json:"resolver" yaml:"resolver"`
	// The scope from which resolution is performed.
	Scope constructs.IConstruct `json:"scope" yaml:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	Preparing *bool `json:"preparing" yaml:"preparing"`
	// Whether to remove undefined elements from arrays and objects when resolving.
	RemoveEmpty *bool `json:"removeEmpty" yaml:"removeEmpty"`
}

// A construct which represents an AWS resource.
type Resource interface {
	constructs.Construct
	IResource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__constructsConstruct
	jsiiProxy_IResource
}

func (j *jsiiProxy_Resource) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewResource_Override(r Resource, scope constructs.Construct, id *string, props *ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Resource",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Resource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Resource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Resource_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Resource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ApplyRemovalPolicy(policy RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Resource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetResourceArnAttribute(arnAttr *string, arnComponents *ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents the environment a given resource lives in.
//
// Used as the return value for the {@link IResource.env} property.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   resourceEnvironment := &resourceEnvironment{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//   }
//
type ResourceEnvironment struct {
	// The AWS account ID that this resource belongs to.
	//
	// Since this can be a Token
	// (for example, when the account is CloudFormation's AWS::AccountId intrinsic),
	// make sure to use Token.compareStrings()
	// instead of just comparing the values for equality.
	Account *string `json:"account" yaml:"account"`
	// The AWS region that this resource belongs to.
	//
	// Since this can be a Token
	// (for example, when the region is CloudFormation's AWS::Region intrinsic),
	// make sure to use Token.compareStrings()
	// instead of just comparing the values for equality.
	Region *string `json:"region" yaml:"region"`
}

// Construction properties for {@link Resource}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   resourceProps := &resourceProps{
//   	account: jsii.String("account"),
//   	environmentFromArn: jsii.String("environmentFromArn"),
//   	physicalName: jsii.String("physicalName"),
//   	region: jsii.String("region"),
//   }
//
type ResourceProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `json:"region" yaml:"region"`
}

// Options for the 'reverse()' operation.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   reverseOptions := &reverseOptions{
//   	failConcat: jsii.Boolean(false),
//   }
//
type ReverseOptions struct {
	// Fail if the given string is a concatenation.
	//
	// If `false`, just return `undefined`.
	FailConcat *bool `json:"failConcat" yaml:"failConcat"`
}

// Accessor for scoped pseudo parameters.
//
// These pseudo parameters are anchored to a stack somewhere in the construct
// tree, and their values will be exported automatically.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   scopedAws := cdk.NewScopedAws(this)
//
type ScopedAws interface {
	AccountId() *string
	NotificationArns() *[]*string
	Partition() *string
	Region() *string
	StackId() *string
	StackName() *string
	UrlSuffix() *string
}

// The jsii proxy struct for ScopedAws
type jsiiProxy_ScopedAws struct {
	_ byte // padding
}

func (j *jsiiProxy_ScopedAws) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


func NewScopedAws(scope constructs.Construct) ScopedAws {
	_init_.Initialize()

	j := jsiiProxy_ScopedAws{}

	_jsii_.Create(
		"aws-cdk-lib.ScopedAws",
		[]interface{}{scope},
		&j,
	)

	return &j
}

func NewScopedAws_Override(s ScopedAws, scope constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.ScopedAws",
		[]interface{}{scope},
		s,
	)
}

// Work with secret values in the CDK.
//
// Constructs that need secrets will declare parameters of type `SecretValue`.
//
// The actual values of these secrets should not be committed to your
// repository, or even end up in the synthesized CloudFormation template. Instead, you should
// store them in an external system like AWS Secrets Manager or SSM Parameter
// Store, and you can reference them by calling `SecretValue.secretsManager()` or
// `SecretValue.ssmSecure()`.
//
// You can use `SecretValue.unsafePlainText()` to construct a `SecretValue` from a
// literal string, but doing so is highly discouraged.
//
// To make sure secret values don't accidentally end up in readable parts
// of your infrastructure definition (such as the environment variables
// of an AWS Lambda Function, where everyone who can read the function
// definition has access to the secret), using secret values directly is not
// allowed. You must pass them to constructs that accept `SecretValue`
// properties, which are guaranteed to use the value only in CloudFormation
// properties that are write-only.
//
// If you are sure that what you are doing is safe, you can call
// `secretValue.unsafeUnwrap()` to access the protected string of the secret
// value.
//
// (If you are writing something like an AWS Lambda Function and need to access
// a secret inside it, make the API call to `GetSecretValue` directly inside
// your Lamba's code, instead of using environment variables.)
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
//   	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
//   		owner: jsii.String("<user>"),
//   		repository: jsii.String("<repo>"),
//   		oauthToken: secretValue.secretsManager(jsii.String("my-github-token")),
//   	}),
//   	basicAuth: amplify.basicAuth.fromCredentials(jsii.String("username"), *secretValue.secretsManager(jsii.String("my-github-token"))),
//   })
//
type SecretValue interface {
	Intrinsic
	// The captured stack trace which represents the location in which this token was created.
	CreationStack() *[]*string
	// Creates a throwable Error object that contains the token creation stack trace.
	NewError(message *string) interface{}
	// Resolve the secret.
	//
	// If the feature flag is not set, resolve as normal. Otherwise, throw a descriptive
	// error that the usage guard is missing.
	Resolve(context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	ToJSON() interface{}
	// Convert an instance of this Token to a string.
	//
	// This method will be called implicitly by language runtimes if the object
	// is embedded into a string. We treat it the same as an explicit
	// stringification.
	ToString() *string
	// Disable usage protection on this secret.
	//
	// Call this to indicate that you want to use the secret value held by this
	// object in an unchecked way. If you don't call this method, using the secret
	// value directly in a string context or as a property value somewhere will
	// produce an error.
	//
	// This method has 'unsafe' in the name on purpose! Make sure that the
	// construct property you are using the returned value in is does not end up
	// in a place in your AWS infrastructure where it could be read by anyone
	// unexpected.
	//
	// When in doubt, don't call this method and only pass the object to constructs that
	// accept `SecretValue` parameters.
	UnsafeUnwrap() *string
}

// The jsii proxy struct for SecretValue
type jsiiProxy_SecretValue struct {
	jsiiProxy_Intrinsic
}

func (j *jsiiProxy_SecretValue) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
func NewSecretValue(protectedValue interface{}, options *IntrinsicProps) SecretValue {
	_init_.Initialize()

	j := jsiiProxy_SecretValue{}

	_jsii_.Create(
		"aws-cdk-lib.SecretValue",
		[]interface{}{protectedValue, options},
		&j,
	)

	return &j
}

// Construct a SecretValue (do not use!).
//
// Do not use the constructor directly: use one of the factory functions on the class
// instead.
func NewSecretValue_Override(s SecretValue, protectedValue interface{}, options *IntrinsicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.SecretValue",
		[]interface{}{protectedValue, options},
		s,
	)
}

// Obtain the secret value through a CloudFormation dynamic reference.
//
// If possible, use `SecretValue.ssmSecure` or `SecretValue.secretsManager` directly.
func SecretValue_CfnDynamicReference(ref CfnDynamicReference) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"cfnDynamicReference",
		[]interface{}{ref},
		&returns,
	)

	return returns
}

// Obtain the secret value through a CloudFormation parameter.
//
// Generally, this is not a recommended approach. AWS Secrets Manager is the
// recommended way to reference secrets.
func SecretValue_CfnParameter(param CfnParameter) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"cfnParameter",
		[]interface{}{param},
		&returns,
	)

	return returns
}

// Test whether an object is a SecretValue.
func SecretValue_IsSecretValue(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"isSecretValue",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Construct a literal secret value for use with secret-aware constructs.
//
// Do not use this method for any secrets that you care about! The value
// will be visible to anyone who has access to the CloudFormation template
// (via the AWS Console, SDKs, or CLI).
//
// The only reasonable use case for using this method is when you are testing.
// Deprecated: Use `unsafePlainText()` instead.
func SecretValue_PlainText(secret *string) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"plainText",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

// Use a resource's output as secret value.
func SecretValue_ResourceAttribute(attr *string) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"resourceAttribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Creates a `SecretValue` with a value which is dynamically loaded from AWS Secrets Manager.
func SecretValue_SecretsManager(secretId *string, options *SecretsManagerSecretOptions) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"secretsManager",
		[]interface{}{secretId, options},
		&returns,
	)

	return returns
}

// Use a secret value stored from a Systems Manager (SSM) parameter.
func SecretValue_SsmSecure(parameterName *string, version *string) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"ssmSecure",
		[]interface{}{parameterName, version},
		&returns,
	)

	return returns
}

// Construct a literal secret value for use with secret-aware constructs.
//
// Do not use this method for any secrets that you care about! The value
// will be visible to anyone who has access to the CloudFormation template
// (via the AWS Console, SDKs, or CLI).
//
// The only reasonable use case for using this method is when you are testing.
func SecretValue_UnsafePlainText(secret *string) SecretValue {
	_init_.Initialize()

	var returns SecretValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.SecretValue",
		"unsafePlainText",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) NewError(message *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"newError",
		[]interface{}{message},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) Resolve(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretValue) UnsafeUnwrap() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"unsafeUnwrap",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for referencing a secret value from Secrets Manager.
//
// Example:
//   codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &bitBucketSourceCredentialsProps{
//   	username: secretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("username"),
//   	}),
//   	password: *secretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("password"),
//   	}),
//   })
//
type SecretsManagerSecretOptions struct {
	// The key of a JSON field to retrieve.
	//
	// This can only be used if the secret
	// stores a JSON object.
	JsonField *string `json:"jsonField" yaml:"jsonField"`
	// Specifies the unique identifier of the version of the secret you want to use.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionId *string `json:"versionId" yaml:"versionId"`
	// Specifies the secret version that you want to retrieve by the staging label attached to the version.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionStage *string `json:"versionStage" yaml:"versionStage"`
}

// Represents the amount of digital storage.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative, or as an unresolved number token.
//
// When the amount is passed as a token, unit conversion is not possible.
//
// Example:
//   var bucket bucket// Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &dataProcessorProps{
//   	bufferInterval: duration.minutes(jsii.Number(5)),
//   	bufferSize: size.mebibytes(jsii.Number(5)),
//   	retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
type Size interface {
	// Checks if size is a token or a resolvable object.
	IsUnresolved() *bool
	// Return this storage as a total number of gibibytes.
	//
	// Returns: the quantity of bytes expressed in gibibytes.
	ToGibibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of kibibytes.
	//
	// Returns: the quantity of bytes expressed in kibibytes.
	ToKibibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of mebibytes.
	//
	// Returns: the quantity of bytes expressed in mebibytes.
	ToMebibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of pebibytes.
	//
	// Returns: the quantity of bytes expressed in pebibytes.
	ToPebibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of tebibytes.
	//
	// Returns: the quantity of bytes expressed in tebibytes.
	ToTebibytes(opts *SizeConversionOptions) *float64
}

// The jsii proxy struct for Size
type jsiiProxy_Size struct {
	_ byte // padding
}

// Create a Storage representing an amount gibibytes.
//
// 1 GiB = 1024 MiB.
//
// Returns: a new `Size` instance.
func Size_Gibibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"gibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount kibibytes.
//
// 1 KiB = 1024 bytes.
//
// Returns: a new `Size` instance.
func Size_Kibibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"kibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount mebibytes.
//
// 1 MiB = 1024 KiB.
//
// Returns: a new `Size` instance.
func Size_Mebibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"mebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount pebibytes.
//
// 1 PiB = 1024 TiB.
//
// Returns: a new `Size` instance.
func Size_Pebibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"pebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount tebibytes.
//
// 1 TiB = 1024 GiB.
//
// Returns: a new `Size` instance.
func Size_Tebibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"tebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) IsUnresolved() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isUnresolved",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToGibibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toGibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToKibibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toKibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToMebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toMebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToPebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toPebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToTebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toTebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Options for how to convert time to a different unit.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   size.mebibytes(jsii.Number(2)).toKibibytes() // yields 2048
//   size.kibibytes(jsii.Number(2050)).toMebibytes(&sizeConversionOptions{
//   	rounding: sizeRoundingBehavior_FLOOR,
//   })
//
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	Rounding SizeRoundingBehavior `json:"rounding" yaml:"rounding"`
}

// Rounding behaviour when converting between units of `Size`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   size.mebibytes(jsii.Number(2)).toKibibytes() // yields 2048
//   size.kibibytes(jsii.Number(2050)).toMebibytes(&sizeConversionOptions{
//   	rounding: sizeRoundingBehavior_FLOOR,
//   })
//
type SizeRoundingBehavior string

const (
	// Fail the conversion if the result is not an integer.
	SizeRoundingBehavior_FAIL SizeRoundingBehavior = "FAIL"
	// If the result is not an integer, round it to the closest integer less than the result.
	SizeRoundingBehavior_FLOOR SizeRoundingBehavior = "FLOOR"
	// Don't round.
	//
	// Return even if the result is a fraction.
	SizeRoundingBehavior_NONE SizeRoundingBehavior = "NONE"
)

// A root construct which represents a single CloudFormation stack.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
//   	name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	name: jsii.String("foo"),
//   	dnsRecordType: servicediscovery.dnsRecordType_A,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTPS,
//   		resourcePath: jsii.String("/healthcheck"),
//   		failureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   	port: jsii.Number(443),
//   })
//
//   app.synth()
//
type Stack interface {
	constructs.Construct
	ITaggable
	// The AWS account into which this stack will be deployed.
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.account` when the stack is defined. This can
	//     either be a concerete account (e.g. `585695031111`) or the
	//     `Aws.accountId` token.
	// 3. `Aws.accountId`, which represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::AccountId" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.account)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **account-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	Account() *string
	// The ID of the cloud assembly artifact for this stack.
	ArtifactId() *string
	// Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack.
	//
	// If the stack is environment-agnostic (either account and/or region are
	// tokens), this property will return an array with 2 tokens that will resolve
	// at deploy-time to the first two availability zones returned from CloudFormation's
	// `Fn::GetAZs` intrinsic function.
	//
	// If they are not available in the context, returns a set of dummy values and
	// reports them as missing, and let the CLI resolve them by calling EC2
	// `DescribeAvailabilityZones` on the target environment.
	//
	// To specify a different strategy for selecting availability zones override this method.
	AvailabilityZones() *[]*string
	// Indicates whether the stack requires bundling or not.
	BundlingRequired() *bool
	// Return the stacks this stack depends on.
	Dependencies() *[]Stack
	// The environment coordinates in which this stack is deployed.
	//
	// In the form
	// `aws://account/region`. Use `stack.account` and `stack.region` to obtain
	// the specific values, no need to parse.
	//
	// You can use this value to determine if two stacks are targeting the same
	// environment.
	//
	// If either `stack.account` or `stack.region` are not concrete values (e.g.
	// `Aws.account` or `Aws.region`) the special strings `unknown-account` and/or
	// `unknown-region` will be used respectively to indicate this stack is
	// region/account-agnostic.
	Environment() *string
	// Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.
	Nested() *bool
	// If this is a nested stack, returns it's parent stack.
	NestedStackParent() Stack
	// If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.
	//
	// `undefined` for top-level (non-nested) stacks.
	NestedStackResource() CfnResource
	// The tree node.
	Node() constructs.Node
	// Returns the list of notification Amazon Resource Names (ARNs) for the current stack.
	NotificationArns() *[]*string
	// The partition in which this stack is defined.
	Partition() *string
	// The AWS region into which this stack will be deployed (e.g. `us-west-2`).
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.region` when the stack is defined. This can
	//     either be a concerete region (e.g. `us-west-2`) or the `Aws.region`
	//     token.
	// 3. `Aws.region`, which is represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::Region" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.region)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **region-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	Region() *string
	// The ID of the stack.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // After resolving, looks like
	//   "arn:aws:cloudformation:us-west-2:123456789012:stack/teststack/51af3dc0-da77-11e4-872e-1234567db123"
	//
	StackId() *string
	// The concrete CloudFormation physical stack name.
	//
	// This is either the name defined explicitly in the `stackName` prop or
	// allocated based on the stack's location in the construct tree. Stacks that
	// are directly defined under the app use their construct `id` as their stack
	// name. Stacks that are defined deeper within the tree will use a hashed naming
	// scheme based on the construct path to ensure uniqueness.
	//
	// If you wish to obtain the deploy-time AWS::StackName intrinsic,
	// you can use `Aws.stackName` directly.
	StackName() *string
	// Synthesis method for this stack.
	Synthesizer() IStackSynthesizer
	// Tags to be applied to the stack.
	Tags() TagManager
	// The name of the CloudFormation template file emitted to the output directory during synthesis.
	//
	// Example value: `MyStack.template.json`
	TemplateFile() *string
	// Options for CloudFormation template (like version, transform, description).
	TemplateOptions() ITemplateOptions
	// Whether termination protection is enabled for this stack.
	TerminationProtection() *bool
	// The Amazon domain suffix for the region in which this stack is defined.
	UrlSuffix() *string
	// Add a dependency between this stack and another stack.
	//
	// This can be used to define dependencies between any two stacks within an
	// app, and also supports nested stacks.
	AddDependency(target Stack, reason *string)
	// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
	//
	// Duplicate values are removed when stack is synthesized.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   var stack stack
	//
	//
	//   stack.addTransform(jsii.String("AWS::Serverless-2016-10-31"))
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
	//
	AddTransform(transform *string)
	// Returns the naming scheme used to allocate logical IDs.
	//
	// By default, uses
	// the `HashedAddressingScheme` but this method can be overridden to customize
	// this behavior.
	//
	// In order to make sure logical IDs are unique and stable, we hash the resource
	// construct tree path (i.e. toplevel/secondlevel/.../myresource) and add it as
	// a suffix to the path components joined without a separator (CloudFormation
	// IDs only allow alphanumeric characters).
	//
	// The result will be:
	//
	//    <path.join('')><md5(path.join('/')>
	//      "human"      "hash"
	//
	// If the "human" part of the ID exceeds 240 characters, we simply trim it so
	// the total ID doesn't exceed CloudFormation's 255 character limit.
	//
	// We only take 8 characters from the md5 hash (0.000005 chance of collision).
	//
	// Special cases:
	//
	// - If the path only contains a single component (i.e. it's a top-level
	//    resource), we won't add the hash to it. The hash is not needed for
	//    disamiguation and also, it allows for a more straightforward migration an
	//    existing CloudFormation template to a CDK stack without logical ID changes
	//    (or renames).
	// - For aesthetic reasons, if the last components of the path are the same
	//    (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
	//    resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
	//    instead of `L1L2PipelinePipeline<HASH>`
	// - If a component is named "Default" it will be omitted from the path. This
	//    allows refactoring higher level abstractions around constructs without affecting
	//    the IDs of already deployed resources.
	// - If a component is named "Resource" it will be omitted from the user-visible
	//    path, but included in the hash. This reduces visual noise in the human readable
	//    part of the identifier.
	AllocateLogicalId(cfnElement CfnElement) *string
	// Create a CloudFormation Export for a value.
	//
	// Returns a string representing the corresponding `Fn.importValue()`
	// expression for this Export. You can control the name for the export by
	// passing the `name` option.
	//
	// If you don't supply a value for `name`, the value you're exporting must be
	// a Resource attribute (for example: `bucket.bucketName`) and it will be
	// given the same name as the automatic cross-stack reference that would be created
	// if you used the attribute in another Stack.
	//
	// One of the uses for this method is to *remove* the relationship between
	// two Stacks established by automatic cross-stack references. It will
	// temporarily ensure that the CloudFormation Export still exists while you
	// remove the reference from the consuming stack. After that, you can remove
	// the resource and the manual export.
	//
	// ## Example
	//
	// Here is how the process works. Let's say there are two stacks,
	// `producerStack` and `consumerStack`, and `producerStack` has a bucket
	// called `bucket`, which is referenced by `consumerStack` (perhaps because
	// an AWS Lambda Function writes into it, or something like that).
	//
	// It is not safe to remove `producerStack.bucket` because as the bucket is being
	// deleted, `consumerStack` might still be using it.
	//
	// Instead, the process takes two deployments:
	//
	// ### Deployment 1: break the relationship
	//
	// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
	//    stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
	//    remove the Lambda Function altogether).
	// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
	//    will make sure the CloudFormation Export continues to exist while the relationship
	//    between the two stacks is being broken.
	// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
	//
	// ### Deployment 2: remove the bucket resource
	//
	// - You are now free to remove the `bucket` resource from `producerStack`.
	// - Don't forget to remove the `exportValue()` call as well.
	// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
	ExportValue(exportedValue interface{}, options *ExportValueOptions) *string
	// Creates an ARN from components.
	//
	// If `partition`, `region` or `account` are not specified, the stack's
	// partition, region and account will be used.
	//
	// If any component is the empty string, an empty string will be inserted
	// into the generated ARN at the location that component corresponds to.
	//
	// The ARN will be formatted as follows:
	//
	//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}}{resource-name}
	//
	// The required ARN pieces that are omitted will be taken from the stack that
	// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
	// can be 'undefined'.
	FormatArn(components *ArnComponents) *string
	// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
	//
	// This method is called when a `CfnElement` is created and used to render the
	// initial logical identity of resources. Logical ID renames are applied at
	// this stage.
	//
	// This method uses the protected method `allocateLogicalId` to render the
	// logical ID for an element. To modify the naming scheme, extend the `Stack`
	// class and override this method.
	GetLogicalId(element CfnElement) *string
	// Look up a fact value for the given fact for the region of this stack.
	//
	// Will return a definite value only if the region of the current stack is resolved.
	// If not, a lookup map will be added to the stack and the lookup will be done at
	// CDK deployment time.
	//
	// What regions will be included in the lookup map is controlled by the
	// `@aws-cdk/core:target-partitions` context value: it must be set to a list
	// of partitions, and only regions from the given partitions will be included.
	// If no such context key is set, all regions will be included.
	//
	// This function is intended to be used by construct library authors. Application
	// builders can rely on the abstractions offered by construct libraries and do
	// not have to worry about regional facts.
	//
	// If `defaultValue` is not given, it is an error if the fact is unknown for
	// the given region.
	RegionalFact(factName *string, defaultValue *string) *string
	// Rename a generated logical identities.
	//
	// To modify the naming scheme strategy, extend the `Stack` class and
	// override the `allocateLogicalId` method.
	RenameLogicalId(oldId *string, newId *string)
	// Indicate that a context key was expected.
	//
	// Contains instructions which will be emitted into the cloud assembly on how
	// the key should be supplied.
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	// Resolve a tokenized value in the context of the current stack.
	Resolve(obj interface{}) interface{}
	// Splits the provided ARN into its components.
	//
	// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
	// and a Token representing a dynamic CloudFormation expression
	// (in which case the returned components will also be dynamic CloudFormation expressions,
	// encoded as Tokens).
	SplitArn(arn *string, arnFormat ArnFormat) *ArnComponents
	// Convert an object, potentially containing tokens, to a JSON string.
	ToJsonString(obj interface{}, space *float64) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Stack
type jsiiProxy_Stack struct {
	internal.Type__constructsConstruct
	jsiiProxy_ITaggable
}

func (j *jsiiProxy_Stack) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) BundlingRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bundlingRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Dependencies() *[]Stack {
	var returns *[]Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) NestedStackParent() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) NestedStackResource() CfnResource {
	var returns CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Synthesizer() IStackSynthesizer {
	var returns IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) TemplateOptions() ITemplateOptions {
	var returns ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stack) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


// Creates a new stack.
func NewStack(scope constructs.Construct, id *string, props *StackProps) Stack {
	_init_.Initialize()

	j := jsiiProxy_Stack{}

	_jsii_.Create(
		"aws-cdk-lib.Stack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new stack.
func NewStack_Override(s Stack, scope constructs.Construct, id *string, props *StackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Stack",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Stack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func Stack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
func Stack_Of(construct constructs.IConstruct) Stack {
	_init_.Initialize()

	var returns Stack

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) AddDependency(target Stack, reason *string) {
	_jsii_.InvokeVoid(
		s,
		"addDependency",
		[]interface{}{target, reason},
	)
}

func (s *jsiiProxy_Stack) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		s,
		"addTransform",
		[]interface{}{transform},
	)
}

func (s *jsiiProxy_Stack) AllocateLogicalId(cfnElement CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) ExportValue(exportedValue interface{}, options *ExportValueOptions) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) FormatArn(components *ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) GetLogicalId(element CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) RegionalFact(factName *string, defaultValue *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"regionalFact",
		[]interface{}{factName, defaultValue},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		s,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

func (s *jsiiProxy_Stack) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		s,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

func (s *jsiiProxy_Stack) Resolve(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) SplitArn(arn *string, arnFormat ArnFormat) *ArnComponents {
	var returns *ArnComponents

	_jsii_.Invoke(
		s,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) ToJsonString(obj interface{}, space *float64) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   type stackUnderTestProps struct {
//   	stackProps
//   	architecture architecture
//   }
//
//   type stackUnderTest struct {
//   	stack
//   }
//
//   func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
//   	this := &stackUnderTest{}
//   	newStack_Override(this, scope, id, props)
//
//   	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
//   		runtime: lambda.runtime_NODEJS_12_X(),
//   		handler: jsii.String("index.handler"),
//   		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   		architecture: props.architecture,
//   	})
//   	return this
//   }
//
//   // Beginning of the test suite
//   app := NewApp()
//
//   stack := NewStack(app, jsii.String("stack"))
//
//   differentArchsCase := NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
//   	stacks: []*stack{
//   		NewStackUnderTest(app, jsii.String("Stack1"), &stackUnderTestProps{
//   			architecture: lambda.*architecture_ARM_64(),
//   		}),
//   		NewStackUnderTest(app, jsii.String("Stack2"), &stackUnderTestProps{
//   			architecture: lambda.*architecture_X86_64(),
//   		}),
//   	},
//   })
//
//   // There must be exactly one instance of TestCase per file
//   // There must be exactly one instance of TestCase per file
//   NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
//
//   	// Register as many test cases as you want here
//   	testCases: []integTestCase{
//   		differentArchsCase,
//   	},
//   })
//
type StackProps struct {
	// Include runtime versioning information in this Stack.
	AnalyticsReporting *bool `json:"analyticsReporting" yaml:"analyticsReporting"`
	// A description of the stack.
	Description *string `json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   NewStack(app, jsii.String("Stack1"), &stackProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   NewStack(app, jsii.String("Stack2"), &stackProps{
	//   	env: &environment{
	//   		account: process.env.cDK_DEFAULT_ACCOUNT,
	//   		region: process.env.cDK_DEFAULT_REGION,
	//   	},
	//   })
	//
	//   // Define multiple stacks stage associated with an environment
	//   myStage := NewStage(app, jsii.String("MyStage"), &stageProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   NewMyStack(myStage, jsii.String("Stack1"))
	//   NewYourStack(myStage, jsii.String("Stack2"))
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   NewMyStack(app, jsii.String("Stack1"))
	//
	Env *Environment `json:"env" yaml:"env"`
	// Name to deploy the stack with.
	StackName *string `json:"stackName" yaml:"stackName"`
	// Synthesis method to use while deploying this stack.
	Synthesizer IStackSynthesizer `json:"synthesizer" yaml:"synthesizer"`
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `json:"terminationProtection" yaml:"terminationProtection"`
}

// Base class for implementing an IStackSynthesizer.
//
// This class needs to exist to provide public surface area for external
// implementations of stack synthesizers. The protected methods give
// access to functions that are otherwise @_internal to the framework
// and could not be accessed by external implementors.
type StackSynthesizer interface {
	IStackSynthesizer
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for StackSynthesizer
type jsiiProxy_StackSynthesizer struct {
	jsiiProxy_IStackSynthesizer
}

func NewStackSynthesizer_Override(s StackSynthesizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.StackSynthesizer",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		s,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		s,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{stack},
	)
}

func (s *jsiiProxy_StackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		s,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (s *jsiiProxy_StackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// An abstract application modeling unit consisting of Stacks that should be deployed together.
//
// Derive a subclass of `Stage` and use it to model a single instance of your
// application.
//
// You can then instantiate your subclass multiple times to model multiple
// copies of your application which should be be deployed to different
// environments.
//
// Example:
//   var pipeline codePipelinetype myOutputStage struct {
//   	stage
//   	loadBalancerAddress cfnOutput
//   }
//
//   func newMyOutputStage(scope construct, id *string, props stageProps) *myOutputStage {
//   	this := &myOutputStage{}
//   	newStage_Override(this, scope, id, props)
//   	this.loadBalancerAddress = NewCfnOutput(this, jsii.String("Output"), &cfnOutputProps{
//   		value: jsii.String("value"),
//   	})
//   	return this
//   }
//
//   lbApp := NewMyOutputStage(this, jsii.String("MyApp"))
//   pipeline.addStage(lbApp, &addStageOpts{
//   	post: []step{
//   		pipelines.NewShellStep(jsii.String("HitEndpoint"), &shellStepProps{
//   			envFromCfnOutputs: map[string]*cfnOutput{
//   				// Make the load balancer address available as $URL inside the commands
//   				"URL": lbApp.loadBalancerAddress,
//   			},
//   			commands: []*string{
//   				jsii.String("curl -Ssf $URL"),
//   			},
//   		}),
//   	},
//   })
//
type Stage interface {
	constructs.Construct
	// The default account for all resources defined within this stage.
	Account() *string
	// Artifact ID of the assembly if it is a nested stage. The root stage (app) will return an empty string.
	//
	// Derived from the construct path.
	ArtifactId() *string
	// The cloud assembly asset output directory.
	AssetOutdir() *string
	// The tree node.
	Node() constructs.Node
	// The cloud assembly output directory.
	Outdir() *string
	// The parent stage or `undefined` if this is the app.
	//
	// *.
	ParentStage() Stage
	// The default region for all resources defined within this stage.
	Region() *string
	// The name of the stage.
	//
	// Based on names of the parent stages separated by
	// hypens.
	StageName() *string
	// Synthesize this stage into a cloud assembly.
	//
	// Once an assembly has been synthesized, it cannot be modified. Subsequent
	// calls will return the same assembly.
	Synth(options *StageSynthesisOptions) cxapi.CloudAssembly
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Stage
type jsiiProxy_Stage struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Stage) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) AssetOutdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetOutdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) ParentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"parentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


func NewStage(scope constructs.Construct, id *string, props *StageProps) Stage {
	_init_.Initialize()

	j := jsiiProxy_Stage{}

	_jsii_.Create(
		"aws-cdk-lib.Stage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStage_Override(s Stage, scope constructs.Construct, id *string, props *StageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Stage",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Stage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a stage.
func Stage_IsStage(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
		"isStage",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return the stage this construct is contained with, if available.
//
// If called
// on a nested stage, returns its parent.
func Stage_Of(construct constructs.IConstruct) Stage {
	_init_.Initialize()

	var returns Stage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) Synth(options *StageSynthesisOptions) cxapi.CloudAssembly {
	var returns cxapi.CloudAssembly

	_jsii_.Invoke(
		s,
		"synth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization props for a stage.
//
// Example:
//   var pipeline codePipelinetype myOutputStage struct {
//   	stage
//   	loadBalancerAddress cfnOutput
//   }
//
//   func newMyOutputStage(scope construct, id *string, props stageProps) *myOutputStage {
//   	this := &myOutputStage{}
//   	newStage_Override(this, scope, id, props)
//   	this.loadBalancerAddress = NewCfnOutput(this, jsii.String("Output"), &cfnOutputProps{
//   		value: jsii.String("value"),
//   	})
//   	return this
//   }
//
//   lbApp := NewMyOutputStage(this, jsii.String("MyApp"))
//   pipeline.addStage(lbApp, &addStageOpts{
//   	post: []step{
//   		pipelines.NewShellStep(jsii.String("HitEndpoint"), &shellStepProps{
//   			envFromCfnOutputs: map[string]*cfnOutput{
//   				// Make the load balancer address available as $URL inside the commands
//   				"URL": lbApp.loadBalancerAddress,
//   			},
//   			commands: []*string{
//   				jsii.String("curl -Ssf $URL"),
//   			},
//   		}),
//   	},
//   })
//
type StageProps struct {
	// Default AWS environment (account/region) for `Stack`s in this `Stage`.
	//
	// Stacks defined inside this `Stage` with either `region` or `account` missing
	// from its env will use the corresponding field given here.
	//
	// If either `region` or `account`is is not configured for `Stack` (either on
	// the `Stack` itself or on the containing `Stage`), the Stack will be
	// *environment-agnostic*.
	//
	// Environment-agnostic stacks can be deployed to any environment, may not be
	// able to take advantage of all features of the CDK. For example, they will
	// not be able to use environmental context lookups, will not automatically
	// translate Service Principals to the right format based on the environment's
	// AWS partition, and other such enhancements.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // Use a concrete account and region to deploy this Stage to
	//   // Use a concrete account and region to deploy this Stage to
	//   NewStage(app, jsii.String("Stage1"), &stageProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // Use the CLI's current credentials to determine the target environment
	//   // Use the CLI's current credentials to determine the target environment
	//   NewStage(app, jsii.String("Stage2"), &stageProps{
	//   	env: &environment{
	//   		account: process.env.cDK_DEFAULT_ACCOUNT,
	//   		region: process.env.cDK_DEFAULT_REGION,
	//   	},
	//   })
	//
	Env *Environment `json:"env" yaml:"env"`
	// The output directory into which to emit synthesized artifacts.
	//
	// Can only be specified if this stage is the root stage (the app). If this is
	// specified and this stage is nested within another stage, an error will be
	// thrown.
	Outdir *string `json:"outdir" yaml:"outdir"`
}

// Options for assembly synthesis.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   stageSynthesisOptions := &stageSynthesisOptions{
//   	force: jsii.Boolean(false),
//   	skipValidation: jsii.Boolean(false),
//   	validateOnSynthesis: jsii.Boolean(false),
//   }
//
type StageSynthesisOptions struct {
	// Force a re-synth, even if the stage has already been synthesized.
	//
	// This is used by tests to allow for incremental verification of the output.
	// Do not use in production.
	Force *bool `json:"force" yaml:"force"`
	// Should we skip construct validation.
	SkipValidation *bool `json:"skipValidation" yaml:"skipValidation"`
	// Whether the stack should be validated after synthesis to check for error metadata.
	ValidateOnSynthesis *bool `json:"validateOnSynthesis" yaml:"validateOnSynthesis"`
}

// Converts all fragments to strings and concats those.
//
// Drops 'undefined's.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   stringConcat := cdk.NewStringConcat()
//
type StringConcat interface {
	IFragmentConcatenator
	// Join the fragment on the left and on the right.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy struct for StringConcat
type jsiiProxy_StringConcat struct {
	jsiiProxy_IFragmentConcatenator
}

func NewStringConcat() StringConcat {
	_init_.Initialize()

	j := jsiiProxy_StringConcat{}

	_jsii_.Create(
		"aws-cdk-lib.StringConcat",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStringConcat_Override(s StringConcat) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.StringConcat",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StringConcat) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Determines how symlinks are followed.
type SymlinkFollowMode string

const (
	// Never follow symlinks.
	SymlinkFollowMode_NEVER SymlinkFollowMode = "NEVER"
	// Materialize all symlinks, whether they are internal or external to the source directory.
	SymlinkFollowMode_ALWAYS SymlinkFollowMode = "ALWAYS"
	// Only follows symlinks that are external to the source directory.
	SymlinkFollowMode_EXTERNAL SymlinkFollowMode = "EXTERNAL"
	// Forbids source from having any symlinks pointing outside of the source tree.
	//
	// This is the safest mode of operation as it ensures that copy operations
	// won't materialize files from the user's file system. Internal symlinks are
	// not followed.
	//
	// If the copy operation runs into an external symlink, it will fail.
	SymlinkFollowMode_BLOCK_EXTERNAL SymlinkFollowMode = "BLOCK_EXTERNAL"
)

// Stack artifact options.
//
// A subset of `cxschema.AwsCloudFormationStackProperties` of optional settings that need to be
// configurable by synthesizers, plus `additionalDependencies`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   synthesizeStackArtifactOptions := &synthesizeStackArtifactOptions{
//   	additionalDependencies: []*string{
//   		jsii.String("additionalDependencies"),
//   	},
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	lookupRole: &bootstrapRole{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		requiresBootstrapStackVersion: jsii.Number(123),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	requiresBootstrapStackVersion: jsii.Number(123),
//   	stackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   }
//
type SynthesizeStackArtifactOptions struct {
	// Identifiers of additional dependencies.
	AdditionalDependencies *[]*string `json:"additionalDependencies" yaml:"additionalDependencies"`
	// The role that needs to be assumed to deploy the stack.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The externalID to use with the assumeRoleArn.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// Only used if `requiresBootstrapStackVersion` is set.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	CloudFormationExecutionRoleArn *string `json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	LookupRole *cloudassemblyschema.BootstrapRole `json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	StackTemplateAssetObjectUrl *string `json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
}

// The Tag Aspect will handle adding a tag to this node and cascading tags to children.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   tag := cdk.NewTag(jsii.String("key"), jsii.String("value"), &tagProps{
//   	applyToLaunchedInstances: jsii.Boolean(false),
//   	excludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	includeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	priority: jsii.Number(123),
//   })
//
type Tag interface {
	IAspect
	// The string key for the tag.
	Key() *string
	Props() *TagProps
	// The string value of the tag.
	Value() *string
	ApplyTag(resource ITaggable)
	// All aspects can visit an IConstruct.
	Visit(construct constructs.IConstruct)
}

// The jsii proxy struct for Tag
type jsiiProxy_Tag struct {
	jsiiProxy_IAspect
}

func (j *jsiiProxy_Tag) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tag) Props() *TagProps {
	var returns *TagProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tag) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewTag(key *string, value *string, props *TagProps) Tag {
	_init_.Initialize()

	j := jsiiProxy_Tag{}

	_jsii_.Create(
		"aws-cdk-lib.Tag",
		[]interface{}{key, value, props},
		&j,
	)

	return &j
}

func NewTag_Override(t Tag, key *string, value *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Tag",
		[]interface{}{key, value, props},
		t,
	)
}

func (t *jsiiProxy_Tag) ApplyTag(resource ITaggable) {
	_jsii_.InvokeVoid(
		t,
		"applyTag",
		[]interface{}{resource},
	)
}

func (t *jsiiProxy_Tag) Visit(construct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		t,
		"visit",
		[]interface{}{construct},
	)
}

// TagManager facilitates a common implementation of tagging for Constructs.
//
// Normally, you do not need to use this class, as the CloudFormation specification
// will indicate which resources are taggable. However, sometimes you will need this
// to make custom resources taggable. Used `tagManager.renderedTags` to obtain a
// value that will resolve to the tags at synthesis time.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cdk "github.com/aws-samples/dummy/awscdkcore"
//
//
//   type myConstruct struct {
//   	resource
//   	tags
//   }
//
//   func newMyConstruct(scope cdk.Construct, id *string) *myConstruct {
//   	this := &myConstruct{}
//   	cdk.NewResource_Override(this, scope, id)
//
//   	cdk.NewCfnResource(this, jsii.String("Resource"), &cfnResourceProps{
//   		type: jsii.String("Whatever::The::Type"),
//   		properties: map[string]interface{}{
//   			// ...
//   			"Tags": this.tags.renderedTags,
//   		},
//   	})
//   	return this
//   }
//
type TagManager interface {
	// A lazy value that represents the rendered tags at synthesis time.
	//
	// If you need to make a custom construct taggable, use the value of this
	// property to pass to the `tags` property of the underlying construct.
	RenderedTags() IResolvable
	// The property name for tag values.
	//
	// Normally this is `tags` but some resources choose a different name. Cognito
	// UserPool uses UserPoolTags.
	TagPropertyName() *string
	// Determine if the aspect applies here.
	//
	// Looks at the include and exclude resourceTypeName arrays to determine if
	// the aspect applies here.
	ApplyTagAspectHere(include *[]*string, exclude *[]*string) *bool
	// Returns true if there are any tags defined.
	HasTags() *bool
	// Removes the specified tag from the array if it exists.
	RemoveTag(key *string, priority *float64)
	// Renders tags into the proper format based on TagType.
	//
	// This method will eagerly render the tags currently applied. In
	// most cases, you should be using `tagManager.renderedTags` instead,
	// which will return a `Lazy` value that will resolve to the correct
	// tags at synthesis time.
	RenderTags() interface{}
	// Adds the specified tag to the array of tags.
	SetTag(key *string, value *string, priority *float64, applyToLaunchedInstances *bool)
	// Render the tags in a readable format.
	TagValues() *map[string]*string
}

// The jsii proxy struct for TagManager
type jsiiProxy_TagManager struct {
	_ byte // padding
}

func (j *jsiiProxy_TagManager) RenderedTags() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"renderedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagManager) TagPropertyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPropertyName",
		&returns,
	)
	return returns
}


func NewTagManager(tagType TagType, resourceTypeName *string, tagStructure interface{}, options *TagManagerOptions) TagManager {
	_init_.Initialize()

	j := jsiiProxy_TagManager{}

	_jsii_.Create(
		"aws-cdk-lib.TagManager",
		[]interface{}{tagType, resourceTypeName, tagStructure, options},
		&j,
	)

	return &j
}

func NewTagManager_Override(t TagManager, tagType TagType, resourceTypeName *string, tagStructure interface{}, options *TagManagerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.TagManager",
		[]interface{}{tagType, resourceTypeName, tagStructure, options},
		t,
	)
}

// Check whether the given construct is Taggable.
func TagManager_IsTaggable(construct interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.TagManager",
		"isTaggable",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) ApplyTagAspectHere(include *[]*string, exclude *[]*string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"applyTagAspectHere",
		[]interface{}{include, exclude},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) HasTags() *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"hasTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) RemoveTag(key *string, priority *float64) {
	_jsii_.InvokeVoid(
		t,
		"removeTag",
		[]interface{}{key, priority},
	)
}

func (t *jsiiProxy_TagManager) RenderTags() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) SetTag(key *string, value *string, priority *float64, applyToLaunchedInstances *bool) {
	_jsii_.InvokeVoid(
		t,
		"setTag",
		[]interface{}{key, value, priority, applyToLaunchedInstances},
	)
}

func (t *jsiiProxy_TagManager) TagValues() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"tagValues",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to configure TagManager behavior.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   tagManagerOptions := &tagManagerOptions{
//   	tagPropertyName: jsii.String("tagPropertyName"),
//   }
//
type TagManagerOptions struct {
	// The name of the property in CloudFormation for these tags.
	//
	// Normally this is `tags`, but Cognito UserPool uses UserPoolTags.
	TagPropertyName *string `json:"tagPropertyName" yaml:"tagPropertyName"`
}

// Properties for a tag.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   tagProps := &tagProps{
//   	applyToLaunchedInstances: jsii.Boolean(false),
//   	excludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	includeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	priority: jsii.Number(123),
//   }
//
type TagProps struct {
	// Whether the tag should be applied to instances in an AutoScalingGroup.
	ApplyToLaunchedInstances *bool `json:"applyToLaunchedInstances" yaml:"applyToLaunchedInstances"`
	// An array of Resource Types that will not receive this tag.
	//
	// An empty array will allow this tag to be applied to all resources. A
	// non-empty array will apply this tag only if the Resource type is not in
	// this array.
	ExcludeResourceTypes *[]*string `json:"excludeResourceTypes" yaml:"excludeResourceTypes"`
	// An array of Resource Types that will receive this tag.
	//
	// An empty array will match any Resource. A non-empty array will apply this
	// tag only to Resource types that are included in this array.
	IncludeResourceTypes *[]*string `json:"includeResourceTypes" yaml:"includeResourceTypes"`
	// Priority of the tag operation.
	//
	// Higher or equal priority tags will take precedence.
	//
	// Setting priority will enable the user to control tags when they need to not
	// follow the default precedence pattern of last applied and closest to the
	// construct in the tree.
	Priority *float64 `json:"priority" yaml:"priority"`
}

type TagType string

const (
	TagType_STANDARD TagType = "STANDARD"
	TagType_AUTOSCALING_GROUP TagType = "AUTOSCALING_GROUP"
	TagType_MAP TagType = "MAP"
	TagType_KEY_VALUE TagType = "KEY_VALUE"
	TagType_NOT_TAGGABLE TagType = "NOT_TAGGABLE"
)

// Manages AWS tags for all resources within a construct scope.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   var service service
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8080),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				path: jsii.String("/ping"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   			timeout: &httpTimeout{
//   				idle: cdk.*duration.seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.file(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.tags.of(node).add(jsii.String("Environment"), jsii.String("Dev"))
//
type Tags interface {
	// add tags to the node of a construct and all its the taggable children.
	Add(key *string, value *string, props *TagProps)
	// remove tags to the node of a construct and all its the taggable children.
	Remove(key *string, props *TagProps)
}

// The jsii proxy struct for Tags
type jsiiProxy_Tags struct {
	_ byte // padding
}

// Returns the tags API for this scope.
func Tags_Of(scope constructs.IConstruct) Tags {
	_init_.Initialize()

	var returns Tags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tags",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tags) Add(key *string, value *string, props *TagProps) {
	_jsii_.InvokeVoid(
		t,
		"add",
		[]interface{}{key, value, props},
	)
}

func (t *jsiiProxy_Tags) Remove(key *string, props *TagProps) {
	_jsii_.InvokeVoid(
		t,
		"remove",
		[]interface{}{key, props},
	)
}

// Options for how to convert time to a different unit.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   timeConversionOptions := &timeConversionOptions{
//   	integral: jsii.Boolean(false),
//   }
//
type TimeConversionOptions struct {
	// If `true`, conversions into a larger time unit (e.g. `Seconds` to `Minutes`) will fail if the result is not an integer.
	Integral *bool `json:"integral" yaml:"integral"`
}

// Represents a special or lazily-evaluated value.
//
// Can be used to delay evaluation of a certain value in case, for example,
// that it requires some context or late-bound data. Can also be used to
// mark values that need special processing at document rendering time.
//
// Tokens can be embedded into strings while retaining their original
// semantics.
type Token interface {
}

// The jsii proxy struct for Token
type jsiiProxy_Token struct {
	_ byte // padding
}

// Return a resolvable representation of the given value.
func Token_AsAny(value interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asAny",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible list representation of this token.
func Token_AsList(value interface{}, options *EncodingOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asList",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible number representation of this token.
func Token_AsNumber(value interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible string representation of this token.
//
// If the Token is initialized with a literal, the stringified value of the
// literal is returned. Otherwise, a special quoted string representation
// of the Token is returned that can be embedded into other strings.
//
// Strings with quoted Tokens in them can be restored back into
// complex values with the Tokens restored by calling `resolve()`
// on the string.
func Token_AsString(value interface{}, options *EncodingOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"asString",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Compare two strings that might contain Tokens with each other.
func Token_CompareStrings(possibleToken1 *string, possibleToken2 *string) TokenComparison {
	_init_.Initialize()

	var returns TokenComparison

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"compareStrings",
		[]interface{}{possibleToken1, possibleToken2},
		&returns,
	)

	return returns
}

// Returns true if obj represents an unresolved value.
//
// One of these must be true:
//
// - `obj` is an IResolvable
// - `obj` is a string containing at least one encoded `IResolvable`
// - `obj` is either an encoded number or list
//
// This does NOT recurse into lists or objects to see if they
// containing resolvables.
func Token_IsUnresolved(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Token",
		"isUnresolved",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// An enum-like class that represents the result of comparing two Tokens.
//
// The return type of {@link Token.compareStrings}.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   tokenComparison := cdk.tokenComparison_BOTH_UNRESOLVED()
//
type TokenComparison interface {
}

// The jsii proxy struct for TokenComparison
type jsiiProxy_TokenComparison struct {
	_ byte // padding
}

func TokenComparison_BOTH_UNRESOLVED() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"aws-cdk-lib.TokenComparison",
		"BOTH_UNRESOLVED",
		&returns,
	)
	return returns
}

func TokenComparison_DIFFERENT() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"aws-cdk-lib.TokenComparison",
		"DIFFERENT",
		&returns,
	)
	return returns
}

func TokenComparison_ONE_UNRESOLVED() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"aws-cdk-lib.TokenComparison",
		"ONE_UNRESOLVED",
		&returns,
	)
	return returns
}

func TokenComparison_SAME() TokenComparison {
	_init_.Initialize()
	var returns TokenComparison
	_jsii_.StaticGet(
		"aws-cdk-lib.TokenComparison",
		"SAME",
		&returns,
	)
	return returns
}

// Less oft-needed functions to manipulate Tokens.
type Tokenization interface {
}

// The jsii proxy struct for Tokenization
type jsiiProxy_Tokenization struct {
	_ byte // padding
}

// Return whether the given object is an IResolvable object.
//
// This is different from Token.isUnresolved() which will also check for
// encoded Tokens, whereas this method will only do a type check on the given
// object.
func Tokenization_IsResolvable(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"isResolvable",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Resolves an object by evaluating all tokens and removing any undefined or empty objects or arrays.
//
// Values can only be primitives, arrays or tokens. Other objects (i.e. with methods) will be rejected.
func Tokenization_Resolve(obj interface{}, options *ResolveOptions) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"resolve",
		[]interface{}{obj, options},
		&returns,
	)

	return returns
}

// Reverse any value into a Resolvable, if possible.
//
// In case of a string, the string must not be a concatenation.
func Tokenization_Reverse(x interface{}, options *ReverseOptions) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverse",
		[]interface{}{x, options},
		&returns,
	)

	return returns
}

// Un-encode a string which is either a complete encoded token, or doesn't contain tokens at all.
//
// It's illegal for the string to be a concatenation of an encoded token and something else.
func Tokenization_ReverseCompleteString(s *string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseCompleteString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a list.
func Tokenization_ReverseList(l *[]*string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseList",
		[]interface{}{l},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a number.
func Tokenization_ReverseNumber(n *float64) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseNumber",
		[]interface{}{n},
		&returns,
	)

	return returns
}

// Un-encode a string potentially containing encoded tokens.
func Tokenization_ReverseString(s *string) TokenizedStringFragments {
	_init_.Initialize()

	var returns TokenizedStringFragments

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"reverseString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Stringify a number directly or lazily if it's a Token.
//
// If it is an object (i.e., { Ref: 'SomeLogicalId' }), return it as-is.
func Tokenization_StringifyNumber(x *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tokenization",
		"stringifyNumber",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Fragments of a concatenated string containing stringified Tokens.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   tokenizedStringFragments := cdk.NewTokenizedStringFragments()
//
type TokenizedStringFragments interface {
	FirstToken() IResolvable
	FirstValue() interface{}
	Length() *float64
	// Return all Tokens from this string.
	Tokens() *[]IResolvable
	AddIntrinsic(value interface{})
	AddLiteral(lit interface{})
	AddToken(token IResolvable)
	// Combine the string fragments using the given joiner.
	//
	// If there are any.
	Join(concat IFragmentConcatenator) interface{}
	// Apply a transformation function to all tokens in the string.
	MapTokens(mapper ITokenMapper) TokenizedStringFragments
}

// The jsii proxy struct for TokenizedStringFragments
type jsiiProxy_TokenizedStringFragments struct {
	_ byte // padding
}

func (j *jsiiProxy_TokenizedStringFragments) FirstToken() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"firstToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) FirstValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Tokens() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


func NewTokenizedStringFragments() TokenizedStringFragments {
	_init_.Initialize()

	j := jsiiProxy_TokenizedStringFragments{}

	_jsii_.Create(
		"aws-cdk-lib.TokenizedStringFragments",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTokenizedStringFragments_Override(t TokenizedStringFragments) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.TokenizedStringFragments",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) Join(concat IFragmentConcatenator) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"join",
		[]interface{}{concat},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenizedStringFragments) MapTokens(mapper ITokenMapper) TokenizedStringFragments {
	var returns TokenizedStringFragments

	_jsii_.Invoke(
		t,
		"mapTokens",
		[]interface{}{mapper},
		&returns,
	)

	return returns
}

// Inspector that maintains an attribute bag.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   treeInspector := cdk.NewTreeInspector()
//
type TreeInspector interface {
	// Represents the bag of attributes as key-value pairs.
	Attributes() *map[string]interface{}
	// Adds attribute to bag.
	//
	// Keys should be added by convention to prevent conflicts
	// i.e. L1 constructs will contain attributes with keys prefixed with aws:cdk:cloudformation
	AddAttribute(key *string, value interface{})
}

// The jsii proxy struct for TreeInspector
type jsiiProxy_TreeInspector struct {
	_ byte // padding
}

func (j *jsiiProxy_TreeInspector) Attributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}


func NewTreeInspector() TreeInspector {
	_init_.Initialize()

	j := jsiiProxy_TreeInspector{}

	_jsii_.Create(
		"aws-cdk-lib.TreeInspector",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTreeInspector_Override(t TreeInspector) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.TreeInspector",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TreeInspector) AddAttribute(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addAttribute",
		[]interface{}{key, value},
	)
}

// Representation of validation results.
//
// Models a tree of validation errors so that we have as much information as possible
// about the failure that occurred.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var validationResults validationResults
//   validationResult := cdk.NewValidationResult(jsii.String("errorMessage"), validationResults)
//
type ValidationResult interface {
	ErrorMessage() *string
	IsSuccess() *bool
	Results() ValidationResults
	// Turn a failed validation into an exception.
	AssertSuccess()
	// Return a string rendering of the tree of validation failures.
	ErrorTree() *string
	// Wrap this result with an error message, if it concerns an error.
	Prefix(message *string) ValidationResult
}

// The jsii proxy struct for ValidationResult
type jsiiProxy_ValidationResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationResult) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResult) IsSuccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResult) Results() ValidationResults {
	var returns ValidationResults
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}


func NewValidationResult(errorMessage *string, results ValidationResults) ValidationResult {
	_init_.Initialize()

	j := jsiiProxy_ValidationResult{}

	_jsii_.Create(
		"aws-cdk-lib.ValidationResult",
		[]interface{}{errorMessage, results},
		&j,
	)

	return &j
}

func NewValidationResult_Override(v ValidationResult, errorMessage *string, results ValidationResults) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.ValidationResult",
		[]interface{}{errorMessage, results},
		v,
	)
}

func (v *jsiiProxy_ValidationResult) AssertSuccess() {
	_jsii_.InvokeVoid(
		v,
		"assertSuccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidationResult) ErrorTree() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"errorTree",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidationResult) Prefix(message *string) ValidationResult {
	var returns ValidationResult

	_jsii_.Invoke(
		v,
		"prefix",
		[]interface{}{message},
		&returns,
	)

	return returns
}

// A collection of validation results.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var validationResult validationResult
//   validationResults := cdk.NewValidationResults([]validationResult{
//   	validationResult,
//   })
//
type ValidationResults interface {
	IsSuccess() *bool
	Results() *[]ValidationResult
	SetResults(val *[]ValidationResult)
	Collect(result ValidationResult)
	ErrorTreeList() *string
	// Wrap up all validation results into a single tree node.
	//
	// If there are failures in the collection, add a message, otherwise
	// return a success.
	Wrap(message *string) ValidationResult
}

// The jsii proxy struct for ValidationResults
type jsiiProxy_ValidationResults struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationResults) IsSuccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResults) Results() *[]ValidationResult {
	var returns *[]ValidationResult
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}


func NewValidationResults(results *[]ValidationResult) ValidationResults {
	_init_.Initialize()

	j := jsiiProxy_ValidationResults{}

	_jsii_.Create(
		"aws-cdk-lib.ValidationResults",
		[]interface{}{results},
		&j,
	)

	return &j
}

func NewValidationResults_Override(v ValidationResults, results *[]ValidationResult) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.ValidationResults",
		[]interface{}{results},
		v,
	)
}

func (j *jsiiProxy_ValidationResults) SetResults(val *[]ValidationResult) {
	_jsii_.Set(
		j,
		"results",
		val,
	)
}

func (v *jsiiProxy_ValidationResults) Collect(result ValidationResult) {
	_jsii_.InvokeVoid(
		v,
		"collect",
		[]interface{}{result},
	)
}

func (v *jsiiProxy_ValidationResults) ErrorTreeList() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"errorTreeList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidationResults) Wrap(message *string) ValidationResult {
	var returns ValidationResult

	_jsii_.Invoke(
		v,
		"wrap",
		[]interface{}{message},
		&returns,
	)

	return returns
}

