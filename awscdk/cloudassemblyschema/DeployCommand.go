package cloudassemblyschema


// Represents a cdk deploy command.
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
type DeployCommand struct {
	// Whether or not to run this command as part of the workflow This can be used if you only want to test some of the workflow for example enable `synth` and disable `deploy` & `destroy` in order to limit the test to synthesis.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// This can be used in combination with `expectedError` to validate that a specific message is returned.
	// Default: - do not validate message.
	//
	ExpectedMessage *string `field:"optional" json:"expectedMessage" yaml:"expectedMessage"`
	// If the runner should expect this command to fail.
	// Default: false.
	//
	ExpectError *bool `field:"optional" json:"expectError" yaml:"expectError"`
	// Additional arguments to pass to the command This can be used to test specific CLI functionality.
	// Default: - only default args are used.
	//
	Args *DeployOptions `field:"optional" json:"args" yaml:"args"`
}

