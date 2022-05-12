package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/integtests/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A collection of test cases.
//
// Each test case file should contain exactly one
// instance of this class.
//
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
//   app := awscdk.NewApp()
//
//   stack := awscdk.Newstack(app, jsii.String("stack"))
//
//   differentArchsCase := awscdk.NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
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
//   awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
//
//   	// Register as many test cases as you want here
//   	testCases: []integTestCase{
//   		differentArchsCase,
//   	},
//   })
//
// Experimental.
type IntegTest interface {
	awscdk.Construct
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

// The jsii proxy struct for IntegTest
type jsiiProxy_IntegTest struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_IntegTest) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTest(scope awscdk.Construct, id *string, props *IntegTestProps) IntegTest {
	_init_.Initialize()

	j := jsiiProxy_IntegTest{}

	_jsii_.Create(
		"monocdk.integ_tests.IntegTest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTest_Override(i IntegTest, scope awscdk.Construct, id *string, props *IntegTestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.IntegTest",
		[]interface{}{scope, id, props},
		i,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func IntegTest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.IntegTest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTest) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegTest) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_IntegTest) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTest) Prepare() {
	_jsii_.InvokeVoid(
		i,
		"prepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegTest) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_IntegTest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTest) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An integration test case.
//
// Allows the definition of test properties that
// apply to all stacks under this case.
//
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
//   app := awscdk.NewApp()
//
//   stack := awscdk.Newstack(app, jsii.String("stack"))
//
//   differentArchsCase := awscdk.NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
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
//   awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
//
//   	// Register as many test cases as you want here
//   	testCases: []integTestCase{
//   		differentArchsCase,
//   	},
//   })
//
// Experimental.
type IntegTestCase interface {
	awscdk.Construct
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
func NewIntegTestCase(scope awscdk.Construct, id *string, props *IntegTestCaseProps) IntegTestCase {
	_init_.Initialize()

	j := jsiiProxy_IntegTestCase{}

	_jsii_.Create(
		"monocdk.integ_tests.IntegTestCase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCase_Override(i IntegTestCase, scope awscdk.Construct, id *string, props *IntegTestCaseProps) {
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

// Properties of an integration test case.
//
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
//   app := awscdk.NewApp()
//
//   stack := awscdk.Newstack(app, jsii.String("stack"))
//
//   differentArchsCase := awscdk.NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
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
//   awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
//
//   	// Register as many test cases as you want here
//   	testCases: []integTestCase{
//   		differentArchsCase,
//   	},
//   })
//
// Experimental.
type IntegTestCaseProps struct {
	// List of CloudFormation resource types in this stack that can be destroyed as part of an update without failing the test.
	//
	// This list should only include resources that for this specific
	// integration test we are sure will not cause errors or an outage if
	// destroyed. For example, maybe we know that a new resource will be created
	// first before the old resource is destroyed which prevents any outage.
	//
	// e.g. ['AWS::IAM::Role']
	// Experimental.
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Experimental.
	CdkCommandOptions *cloudassemblyschema.CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Experimental.
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Experimental.
	Hooks *cloudassemblyschema.Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Experimental.
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// Stacks to be deployed during the test.
	// Experimental.
	Stacks *[]awscdk.Stack `field:"required" json:"stacks" yaml:"stacks"`
}

// Integration test properties.
//
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
//   app := awscdk.NewApp()
//
//   stack := awscdk.Newstack(app, jsii.String("stack"))
//
//   differentArchsCase := awscdk.NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
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
//   awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
//
//   	// Register as many test cases as you want here
//   	testCases: []integTestCase{
//   		differentArchsCase,
//   	},
//   })
//
// Experimental.
type IntegTestProps struct {
	// List of test cases that make up this test.
	// Experimental.
	TestCases *[]IntegTestCase `field:"required" json:"testCases" yaml:"testCases"`
}

