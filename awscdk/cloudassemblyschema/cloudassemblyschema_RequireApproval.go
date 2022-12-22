package cloudassemblyschema


// In what scenarios should the CLI ask for approval.
//
// Example:
//   app := awscdk.NewApp()
//
//   stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))
//
//   stack := awscdk.NewStack(app, jsii.String("stack"))
//
//   testCase := awscdk.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &integTestProps{
//   	testCases: []stack{
//   		stackUnderTest,
//   	},
//   	diffAssets: jsii.Boolean(true),
//   	stackUpdateWorkflow: jsii.Boolean(true),
//   	cdkCommandOptions: &cdkCommands{
//   		deploy: &deployCommand{
//   			args: &deployOptions{
//   				requireApproval: awscdk.RequireApproval_NEVER,
//   				json: jsii.Boolean(true),
//   			},
//   		},
//   		destroy: &destroyCommand{
//   			args: &destroyOptions{
//   				force: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type RequireApproval string

const (
	// Never ask for approval.
	// Experimental.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	// Experimental.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	// Experimental.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

