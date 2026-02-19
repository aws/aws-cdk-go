package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An integration test case stack. Allows the definition of test properties that should apply to this stack.
//
// This should be used if there are multiple stacks in the integration test
// and it is necessary to specify different test case option for each. Otherwise
// normal stacks should be added to IntegTest.
//
// Example:
//   var app App
//   var stackUnderTest Stack
//
//   testCaseWithAssets := awscdkintegtestsalpha.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &IntegTestCaseStackProps{
//   	DiffAssets: jsii.Boolean(true),
//   })
//
//   awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
//   	TestCases: []Stack{
//   		stackUnderTest,
//   		testCaseWithAssets,
//   	},
//   })
//
// Experimental.
type IntegTestCaseStack interface {
	awscdk.Stack
	// The AWS account into which this stack will be deployed.
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.account` when the stack is defined. This can
	//    either be a concrete account (e.g. `585695031111`) or the
	//    `Aws.ACCOUNT_ID` token.
	// 3. `Aws.ACCOUNT_ID`, which represents the CloudFormation intrinsic reference
	//    `{ "Ref": "AWS::AccountId" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concrete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.account)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into an **account-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other account-agnostic behavior.
	// Experimental.
	Account() *string
	// The ID of the cloud assembly artifact for this stack.
	// Experimental.
	ArtifactId() *string
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
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
	// Experimental.
	AvailabilityZones() *[]*string
	// Indicates whether the stack requires bundling or not.
	// Experimental.
	BundlingRequired() *bool
	// Return the stacks this stack depends on.
	// Experimental.
	Dependencies() *[]awscdk.Stack
	// The environment this Stack deploys to.
	// Experimental.
	Env() *interfaces.ResourceEnvironment
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
	// `Aws.ACCOUNT_ID` or `Aws.REGION`) the special strings `unknown-account` and/or
	// `unknown-region` will be used respectively to indicate this stack is
	// region/account-agnostic.
	// Experimental.
	Environment() *string
	// Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.
	// Experimental.
	Nested() *bool
	// If this is a nested stack, returns it's parent stack.
	// Experimental.
	NestedStackParent() awscdk.Stack
	// If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.
	//
	// `undefined` for top-level (non-nested) stacks.
	// Experimental.
	NestedStackResource() awscdk.CfnResource
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns the list of notification Amazon Resource Names (ARNs) for the current stack.
	// Experimental.
	NotificationArns() *[]*string
	// The partition in which this stack is defined.
	// Experimental.
	Partition() *string
	// The AWS region into which this stack will be deployed (e.g. `us-west-2`).
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.region` when the stack is defined. This can
	//    either be a concrete region (e.g. `us-west-2`) or the `Aws.REGION`
	//    token.
	// 3. `Aws.REGION`, which is represents the CloudFormation intrinsic reference
	//    `{ "Ref": "AWS::Region" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concrete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.region)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **region-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	// Experimental.
	Region() *string
	// The ID of the stack.
	//
	// Example:
	//   // After resolving, looks like
	//   'arn:aws:cloudformation:us-west-2:123456789012:stack/teststack/51af3dc0-da77-11e4-872e-1234567db123'
	//
	// Experimental.
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
	// you can use `Aws.STACK_NAME` directly.
	// Experimental.
	StackName() *string
	// Synthesis method for this stack.
	// Experimental.
	Synthesizer() awscdk.IStackSynthesizer
	// Tags to be applied to the stack.
	// Experimental.
	Tags() awscdk.TagManager
	// The name of the CloudFormation template file emitted to the output directory during synthesis.
	//
	// Example value: `MyStack.template.json`
	// Experimental.
	TemplateFile() *string
	// Options for CloudFormation template (like version, transform, description).
	// Experimental.
	TemplateOptions() awscdk.ITemplateOptions
	// Whether termination protection is enabled for this stack.
	// Experimental.
	TerminationProtection() *bool
	// Experimental.
	SetTerminationProtection(val *bool)
	// The Amazon domain suffix for the region in which this stack is defined.
	// Experimental.
	UrlSuffix() *string
	// Add a dependency between this stack and another stack.
	//
	// This can be used to define dependencies between any two stacks within an
	// app, and also supports nested stacks.
	// Experimental.
	AddDependency(target awscdk.Stack, reason *string)
	// Adds an arbitrary key-value pair, with information you want to record about the stack.
	//
	// These get translated to the Metadata section of the generated template.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Configure a stack tag.
	//
	// At deploy time, CloudFormation will automatically apply all stack tags to all resources in the stack.
	// Experimental.
	AddStackTag(tagName *string, tagValue *string)
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
	// Experimental.
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
	//   <path.join('')><md5(path.join('/')>
	//     "human"      "hash"
	//
	// If the "human" part of the ID exceeds 240 characters, we simply trim it so
	// the total ID doesn't exceed CloudFormation's 255 character limit.
	//
	// We only take 8 characters from the md5 hash (0.000005 chance of collision).
	//
	// Special cases:
	//
	// - If the path only contains a single component (i.e. it's a top-level
	//   resource), we won't add the hash to it. The hash is not needed for
	//   disambiguation and also, it allows for a more straightforward migration an
	//   existing CloudFormation template to a CDK stack without logical ID changes
	//   (or renames).
	// - For aesthetic reasons, if the last components of the path are the same
	//   (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
	//   resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
	//   instead of `L1L2PipelinePipeline<HASH>`
	// - If a component is named "Default" it will be omitted from the path. This
	//   allows refactoring higher level abstractions around constructs without affecting
	//   the IDs of already deployed resources.
	// - If a component is named "Resource" it will be omitted from the user-visible
	//   path, but included in the hash. This reduces visual noise in the human readable
	//   part of the identifier.
	// Experimental.
	AllocateLogicalId(cfnElement awscdk.CfnElement) *string
	// Create a CloudFormation Export for a string list value.
	//
	// Returns a string list representing the corresponding `Fn.importValue()`
	// expression for this Export. The export expression is automatically wrapped with an
	// `Fn::Join` and the import value with an `Fn::Split`, since CloudFormation can only
	// export strings. You can control the name for the export by passing the `name` option.
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
	// See `exportValue` for an example of this process.
	// Experimental.
	ExportStringListValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *[]*string
	// Create a CloudFormation Export for a string value.
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
	// **Deployment 1: break the relationship**:
	//
	// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
	//   stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
	//   remove the Lambda Function altogether).
	// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
	//   will make sure the CloudFormation Export continues to exist while the relationship
	//   between the two stacks is being broken.
	// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
	//
	// **Deployment 2: remove the bucket resource**:
	//
	// - You are now free to remove the `bucket` resource from `producerStack`.
	// - Don't forget to remove the `exportValue()` call as well.
	// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
	// Experimental.
	ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string
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
	//   arn:{partition}:{service}:{region}:{account}:{resource}{sep}{resource-name}
	//
	// The required ARN pieces that are omitted will be taken from the stack that
	// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
	// can be 'undefined'.
	// Experimental.
	FormatArn(components *awscdk.ArnComponents) *string
	// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
	//
	// This method is called when a `CfnElement` is created and used to render the
	// initial logical identity of resources. Logical ID renames are applied at
	// this stage.
	//
	// This method uses the protected method `allocateLogicalId` to render the
	// logical ID for an element. To modify the naming scheme, extend the `Stack`
	// class and override this method.
	// Experimental.
	GetLogicalId(element awscdk.CfnElement) *string
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
	// Experimental.
	RegionalFact(factName *string, defaultValue *string) *string
	// Remove a stack tag.
	//
	// At deploy time, CloudFormation will automatically apply all stack tags to all resources in the stack.
	// Experimental.
	RemoveStackTag(tagName *string)
	// Rename a generated logical identities.
	//
	// To modify the naming scheme strategy, extend the `Stack` class and
	// override the `allocateLogicalId` method.
	// Experimental.
	RenameLogicalId(oldId *string, newId *string)
	// Indicate that a context key was expected.
	//
	// Contains instructions which will be emitted into the cloud assembly on how
	// the key should be supplied.
	// Experimental.
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	// Resolve a tokenized value in the context of the current stack.
	// Experimental.
	Resolve(obj interface{}) interface{}
	// Splits the provided ARN into its components.
	//
	// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
	// and a Token representing a dynamic CloudFormation expression
	// (in which case the returned components will also be dynamic CloudFormation expressions,
	// encoded as Tokens).
	// Experimental.
	SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents
	// Convert an object, potentially containing tokens, to a JSON string.
	// Experimental.
	ToJsonString(obj interface{}, space *float64) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Convert an object, potentially containing tokens, to a YAML string.
	// Experimental.
	ToYamlString(obj interface{}) *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for IntegTestCaseStack
type jsiiProxy_IntegTestCaseStack struct {
	internal.Type__awscdkStack
}

func (j *jsiiProxy_IntegTestCaseStack) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Assertions() IDeployAssert {
	var returns IDeployAssert
	_jsii_.Get(
		j,
		"assertions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) BundlingRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bundlingRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Dependencies() *[]awscdk.Stack {
	var returns *[]awscdk.Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NestedStackParent() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NestedStackResource() awscdk.CfnResource {
	var returns awscdk.CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Synthesizer() awscdk.IStackSynthesizer {
	var returns awscdk.IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TemplateOptions() awscdk.ITemplateOptions {
	var returns awscdk.ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTestCaseStack(scope constructs.Construct, id *string, props *IntegTestCaseStackProps) IntegTestCaseStack {
	_init_.Initialize()

	if err := validateNewIntegTestCaseStackParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegTestCaseStack{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCaseStack_Override(i IntegTestCaseStack, scope constructs.Construct, id *string, props *IntegTestCaseStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		[]interface{}{scope, id, props},
		i,
	)
}

func (j *jsiiProxy_IntegTestCaseStack)SetTerminationProtection(val *bool) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
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
// Experimental.
func IntegTestCaseStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTestCaseStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns whether the construct is a IntegTestCaseStack.
// Experimental.
func IntegTestCaseStack_IsIntegTestCaseStack(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTestCaseStack_IsIntegTestCaseStackParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isIntegTestCaseStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
// Experimental.
func IntegTestCaseStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTestCaseStack_IsStackParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
// Experimental.
func IntegTestCaseStack_Of(construct constructs.IConstruct) awscdk.Stack {
	_init_.Initialize()

	if err := validateIntegTestCaseStack_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns awscdk.Stack

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) AddDependency(target awscdk.Stack, reason *string) {
	if err := i.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDependency",
		[]interface{}{target, reason},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AddMetadata(key *string, value interface{}) {
	if err := i.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AddStackTag(tagName *string, tagValue *string) {
	if err := i.validateAddStackTagParameters(tagName, tagValue); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addStackTag",
		[]interface{}{tagName, tagValue},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AddTransform(transform *string) {
	if err := i.validateAddTransformParameters(transform); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTransform",
		[]interface{}{transform},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AllocateLogicalId(cfnElement awscdk.CfnElement) *string {
	if err := i.validateAllocateLogicalIdParameters(cfnElement); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ExportStringListValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *[]*string {
	if err := i.validateExportStringListValueParameters(exportedValue, options); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"exportStringListValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string {
	if err := i.validateExportValueParameters(exportedValue, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) FormatArn(components *awscdk.ArnComponents) *string {
	if err := i.validateFormatArnParameters(components); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) GetLogicalId(element awscdk.CfnElement) *string {
	if err := i.validateGetLogicalIdParameters(element); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) RegionalFact(factName *string, defaultValue *string) *string {
	if err := i.validateRegionalFactParameters(factName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"regionalFact",
		[]interface{}{factName, defaultValue},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) RemoveStackTag(tagName *string) {
	if err := i.validateRemoveStackTagParameters(tagName); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"removeStackTag",
		[]interface{}{tagName},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) RenameLogicalId(oldId *string, newId *string) {
	if err := i.validateRenameLogicalIdParameters(oldId, newId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	if err := i.validateReportMissingContextKeyParameters(report); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) Resolve(obj interface{}) interface{} {
	if err := i.validateResolveParameters(obj); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents {
	if err := i.validateSplitArnParameters(arn, arnFormat); err != nil {
		panic(err)
	}
	var returns *awscdk.ArnComponents

	_jsii_.Invoke(
		i,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ToJsonString(obj interface{}, space *float64) *string {
	if err := i.validateToJsonStringParameters(obj); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ToYamlString(obj interface{}) *string {
	if err := i.validateToYamlStringParameters(obj); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"toYamlString",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

