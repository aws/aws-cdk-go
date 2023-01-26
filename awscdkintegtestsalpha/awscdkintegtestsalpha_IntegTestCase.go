// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An integration test case. Allows the definition of test properties that apply to all stacks under this case.
//
// It is recommended that you use the IntegTest construct since that will create
// a default IntegTestCase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   integTestCase := integ_tests_alpha.NewIntegTestCase(this, jsii.String("MyIntegTestCase"), &integTestCaseProps{
//   	stacks: []*stack{
//   		stack,
//   	},
//
//   	// the properties below are optional
//   	allowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	assertionStack: stack,
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
	constructs.Construct
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
	// The integration test manifest for this test case.
	//
	// Manifests are used
	// by the integration test runner.
	// Experimental.
	Manifest() *cloudassemblyschema.IntegManifest
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegTestCase
type jsiiProxy_IntegTestCase struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_IntegTestCase) Node() constructs.Node {
	var returns constructs.Node
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
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCase_Override(i IntegTestCase, scope constructs.Construct, id *string, props *IntegTestCaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		[]interface{}{scope, id, props},
		i,
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
func IntegTestCase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTestCase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

