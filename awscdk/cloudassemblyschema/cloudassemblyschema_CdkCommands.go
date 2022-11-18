package cloudassemblyschema


// Options for specific cdk commands that are run as part of the integration test workflow.
//
// Example:
//   app := awscdk.NewApp()
//
//   stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))
//
//   stack := awscdk.NewStack(app, jsii.String("stack"))
//
//   testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &integTestProps{
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
type CdkCommands struct {
	// Options to for the cdk deploy command.
	Deploy *DeployCommand `field:"optional" json:"deploy" yaml:"deploy"`
	// Options to for the cdk destroy command.
	Destroy *DestroyCommand `field:"optional" json:"destroy" yaml:"destroy"`
}

