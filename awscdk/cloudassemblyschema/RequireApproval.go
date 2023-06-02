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
//   testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &IntegTestProps{
//   	TestCases: []stack{
//   		stackUnderTest,
//   	},
//   	DiffAssets: jsii.Boolean(true),
//   	StackUpdateWorkflow: jsii.Boolean(true),
//   	CdkCommandOptions: &CdkCommands{
//   		Deploy: &DeployCommand{
//   			Args: &DeployOptions{
//   				RequireApproval: awscdk.RequireApproval_NEVER,
//   				Json: jsii.Boolean(true),
//   			},
//   		},
//   		Destroy: &DestroyCommand{
//   			Args: &DestroyOptions{
//   				Force: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   })
//
type RequireApproval string

const (
	// Never ask for approval.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

