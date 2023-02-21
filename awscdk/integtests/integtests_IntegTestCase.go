package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/integtests/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// An integration test case. Allows the definition of test properties that apply to all stacks under this case.
//
// It is recommended that you use the IntegTest construct since that will create
// a default IntegTestCase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   integTestCase := awscdk.Integ_tests.NewIntegTestCase(this, jsii.String("MyIntegTestCase"), &integTestCaseProps{
//   	stacks: []*stack{
//   		stack,
//   	},
//
//   	// the properties below are optional
//   	allowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	cdkCommandOptions: &cdkCommands{
//   		deploy: &deployCommand{
//   			args: &deployOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				changeSetName: jsii.String("changeSetName"),
//   				ci: jsii.Boolean(false),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				execute: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				notificationArns: []*string{
//   					jsii.String("notificationArns"),
//   				},
//   				output: jsii.String("output"),
//   				outputsFile: jsii.String("outputsFile"),
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				requireApproval: awscdk.Cloud_assembly_schema.requireApproval_NEVER,
//   				reuseAssets: []*string{
//   					jsii.String("reuseAssets"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				rollback: jsii.Boolean(false),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				toolkitStackName: jsii.String("toolkitStackName"),
//   				trace: jsii.Boolean(false),
//   				usePreviousParameters: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   		destroy: &destroyCommand{
//   			args: &destroyOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				output: jsii.String("output"),
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				roleArn: jsii.String("roleArn"),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				trace: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   	},
//   	diffAssets: jsii.Boolean(false),
//   	hooks: &hooks{
//   		postDeploy: []*string{
//   			jsii.String("postDeploy"),
//   		},
//   		postDestroy: []*string{
//   			jsii.String("postDestroy"),
//   		},
//   		preDeploy: []*string{
//   			jsii.String("preDeploy"),
//   		},
//   		preDestroy: []*string{
//   			jsii.String("preDestroy"),
//   		},
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	stackUpdateWorkflow: jsii.Boolean(false),
//   })
//
// Experimental.
type IntegTestCase interface {
	awscdk.Construct
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
	// The integration test manifest for this test case.
	//
	// Manifests are used
	// by the integration test runner.
	// Experimental.
	Manifest() *cloudassemblyschema.IntegManifest
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for IntegTestCase
type jsiiProxy_IntegTestCase struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_IntegTestCase) Assertions() IDeployAssert {
	var returns IDeployAssert
	_jsii_.Get(
		j,
		"assertions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCase) Manifest() *cloudassemblyschema.IntegManifest {
	var returns *cloudassemblyschema.IntegManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTestCase(scope constructs.Construct, id *string, props *IntegTestCaseProps) IntegTestCase {
	_init_.Initialize()

	if err := validateNewIntegTestCaseParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegTestCase{}

	_jsii_.Create(
		"monocdk.integ_tests.IntegTestCase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCase_Override(i IntegTestCase, scope constructs.Construct, id *string, props *IntegTestCaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.IntegTestCase",
		[]interface{}{scope, id, props},
		i,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func IntegTestCase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTestCase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.IntegTestCase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCase) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegTestCase) OnSynthesize(session constructs.ISynthesisSession) {
	if err := i.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_IntegTestCase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCase) Prepare() {
	_jsii_.InvokeVoid(
		i,
		"prepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegTestCase) Synthesize(session awscdk.ISynthesisSession) {
	if err := i.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_IntegTestCase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

