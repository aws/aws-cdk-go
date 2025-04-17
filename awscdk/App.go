package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
)

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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket iBucket
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"))
//
//   dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ImportSource: &ImportSourceSpecification{
//   		CompressionType: dynamodb.InputCompressionType_GZIP,
//   		InputFormat: dynamodb.InputFormat_Csv(&CsvOptions{
//   			Delimiter: jsii.String(","),
//   			HeaderList: []*string{
//   				jsii.String("id"),
//   				jsii.String("name"),
//   			},
//   		}),
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("prefix"),
//   	},
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
	// Validation plugins to run during synthesis.
	//
	// If any plugin reports any violation,
	// synthesis will be interrupted and the report displayed to the user.
	// Default: - no validation plugins are used.
	//
	PolicyValidationBeta1() *[]IPolicyValidationPluginBeta1
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

func (j *jsiiProxy_App) PolicyValidationBeta1() *[]IPolicyValidationPluginBeta1 {
	var returns *[]IPolicyValidationPluginBeta1
	_jsii_.Get(
		j,
		"policyValidationBeta1",
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

	if err := validateNewAppParameters(props); err != nil {
		panic(err)
	}
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

	if err := validateApp_IsAppParameters(obj); err != nil {
		panic(err)
	}
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
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateApp_IsStageParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateApp_OfParameters(construct); err != nil {
		panic(err)
	}
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
	if err := a.validateSynthParameters(options); err != nil {
		panic(err)
	}
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

