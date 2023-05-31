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
//   integTestCase := awscdk.Integ_tests.NewIntegTestCase(this, jsii.String("MyIntegTestCase"), &IntegTestCaseProps{
//   	Stacks: []*stack{
//   		stack,
//   	},
//
//   	// the properties below are optional
//   	AllowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	CdkCommandOptions: &CdkCommands{
//   		Deploy: &DeployCommand{
//   			Args: &DeployOptions{
//   				All: jsii.Boolean(false),
//   				App: jsii.String("app"),
//   				AssetMetadata: jsii.Boolean(false),
//   				CaBundlePath: jsii.String("caBundlePath"),
//   				ChangeSetName: jsii.String("changeSetName"),
//   				Ci: jsii.Boolean(false),
//   				Color: jsii.Boolean(false),
//   				Context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				Debug: jsii.Boolean(false),
//   				Ec2Creds: jsii.Boolean(false),
//   				Exclusively: jsii.Boolean(false),
//   				Execute: jsii.Boolean(false),
//   				Force: jsii.Boolean(false),
//   				IgnoreErrors: jsii.Boolean(false),
//   				Json: jsii.Boolean(false),
//   				Lookups: jsii.Boolean(false),
//   				Notices: jsii.Boolean(false),
//   				NotificationArns: []*string{
//   					jsii.String("notificationArns"),
//   				},
//   				Output: jsii.String("output"),
//   				OutputsFile: jsii.String("outputsFile"),
//   				Parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				PathMetadata: jsii.Boolean(false),
//   				Profile: jsii.String("profile"),
//   				Proxy: jsii.String("proxy"),
//   				RequireApproval: awscdk.Cloud_assembly_schema.RequireApproval_NEVER,
//   				ReuseAssets: []*string{
//   					jsii.String("reuseAssets"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				Rollback: jsii.Boolean(false),
//   				Stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				Staging: jsii.Boolean(false),
//   				Strict: jsii.Boolean(false),
//   				ToolkitStackName: jsii.String("toolkitStackName"),
//   				Trace: jsii.Boolean(false),
//   				UsePreviousParameters: jsii.Boolean(false),
//   				Verbose: jsii.Boolean(false),
//   				VersionReporting: jsii.Boolean(false),
//   			},
//   			Enabled: jsii.Boolean(false),
//   			ExpectedMessage: jsii.String("expectedMessage"),
//   			ExpectError: jsii.Boolean(false),
//   		},
//   		Destroy: &DestroyCommand{
//   			Args: &DestroyOptions{
//   				All: jsii.Boolean(false),
//   				App: jsii.String("app"),
//   				AssetMetadata: jsii.Boolean(false),
//   				CaBundlePath: jsii.String("caBundlePath"),
//   				Color: jsii.Boolean(false),
//   				Context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				Debug: jsii.Boolean(false),
//   				Ec2Creds: jsii.Boolean(false),
//   				Exclusively: jsii.Boolean(false),
//   				Force: jsii.Boolean(false),
//   				IgnoreErrors: jsii.Boolean(false),
//   				Json: jsii.Boolean(false),
//   				Lookups: jsii.Boolean(false),
//   				Notices: jsii.Boolean(false),
//   				Output: jsii.String("output"),
//   				PathMetadata: jsii.Boolean(false),
//   				Profile: jsii.String("profile"),
//   				Proxy: jsii.String("proxy"),
//   				RoleArn: jsii.String("roleArn"),
//   				Stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				Staging: jsii.Boolean(false),
//   				Strict: jsii.Boolean(false),
//   				Trace: jsii.Boolean(false),
//   				Verbose: jsii.Boolean(false),
//   				VersionReporting: jsii.Boolean(false),
//   			},
//   			Enabled: jsii.Boolean(false),
//   			ExpectedMessage: jsii.String("expectedMessage"),
//   			ExpectError: jsii.Boolean(false),
//   		},
//   	},
//   	DiffAssets: jsii.Boolean(false),
//   	Hooks: &Hooks{
//   		PostDeploy: []*string{
//   			jsii.String("postDeploy"),
//   		},
//   		PostDestroy: []*string{
//   			jsii.String("postDestroy"),
//   		},
//   		PreDeploy: []*string{
//   			jsii.String("preDeploy"),
//   		},
//   		PreDestroy: []*string{
//   			jsii.String("preDestroy"),
//   		},
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	StackUpdateWorkflow: jsii.Boolean(false),
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

