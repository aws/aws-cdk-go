# integ-tests

## Usage

Suppose you have a simple stack, that only encapsulates a Lambda function with a
certain handler:

```go
type stackUnderTestProps struct {
	stackProps
	architecture architecture
}

type stackUnderTest struct {
	stack
}

func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
	this := &stackUnderTest{}
	newStack_Override(this, scope, id, props)

	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
		runtime: lambda.runtime_NODEJS_12_X(),
		handler: jsii.String("index.handler"),
		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		architecture: props.architecture,
	})
	return this
}
```

You may want to test this stack under different conditions. For example, we want
this stack to be deployed correctly, regardless of the architecture we choose
for the Lambda function. In particular, it should work for both `ARM_64` and
`X86_64`. So you can create an `IntegTestCase` that exercises both scenarios:

```go
type stackUnderTestProps struct {
	stackProps
	architecture architecture
}

type stackUnderTest struct {
	stack
}

func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
	this := &stackUnderTest{}
	newStack_Override(this, scope, id, props)

	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
		runtime: lambda.runtime_NODEJS_12_X(),
		handler: jsii.String("index.handler"),
		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		architecture: props.architecture,
	})
	return this
}

// Beginning of the test suite
app := awscdk.NewApp()

stack := awscdk.Newstack(app, jsii.String("stack"))

differentArchsCase := awscdk.NewIntegTestCase(stack, jsii.String("DifferentArchitectures"), &integTestCaseProps{
	stacks: []*stack{
		NewStackUnderTest(app, jsii.String("Stack1"), &stackUnderTestProps{
			architecture: lambda.*architecture_ARM_64(),
		}),
		NewStackUnderTest(app, jsii.String("Stack2"), &stackUnderTestProps{
			architecture: lambda.*architecture_X86_64(),
		}),
	},
})

// There must be exactly one instance of TestCase per file
// There must be exactly one instance of TestCase per file
awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{

	// Register as many test cases as you want here
	testCases: []integTestCase{
		differentArchsCase,
	},
})
```

This is all the instruction you need for the integration test runner to know
which stacks to synthesize, deploy and destroy. But you may also need to
customize the behavior of the runner by changing its parameters. For example:

```go
app := awscdk.NewApp()

stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))

stack := awscdk.NewStack(app, jsii.String("stack"))

testCase := awscdk.NewIntegTestCase(stack, jsii.String("CustomizedDeploymentWorkflow"), &integTestCaseProps{
	stacks: []stack{
		stackUnderTest,
	},
	diffAssets: jsii.Boolean(true),
	stackUpdateWorkflow: jsii.Boolean(true),
	cdkCommandOptions: &cdkCommands{
		deploy: &deployCommand{
			args: &deployOptions{
				requireApproval: awscdk.RequireApproval_NEVER,
				json: jsii.Boolean(true),
			},
		},
		destroy: &destroyCommand{
			args: &destroyOptions{
				force: jsii.Boolean(true),
			},
		},
	},
})

awscdk.NewIntegTest(app, jsii.String("integ-test"), &integTestProps{
	testCases: []integTestCase{
		testCase,
	},
})
```
