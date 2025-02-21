package cloudassemblyschema


// Represents a cdk deploy command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployCommand := &DeployCommand{
//   	Args: &DeployOptions{
//   		All: jsii.Boolean(false),
//   		App: jsii.String("app"),
//   		AssetMetadata: jsii.Boolean(false),
//   		CaBundlePath: jsii.String("caBundlePath"),
//   		ChangeSetName: jsii.String("changeSetName"),
//   		Ci: jsii.Boolean(false),
//   		Color: jsii.Boolean(false),
//   		Concurrency: jsii.Number(123),
//   		Context: map[string]*string{
//   			"contextKey": jsii.String("context"),
//   		},
//   		Debug: jsii.Boolean(false),
//   		Ec2Creds: jsii.Boolean(false),
//   		Exclusively: jsii.Boolean(false),
//   		Execute: jsii.Boolean(false),
//   		Force: jsii.Boolean(false),
//   		IgnoreErrors: jsii.Boolean(false),
//   		Json: jsii.Boolean(false),
//   		Lookups: jsii.Boolean(false),
//   		Notices: jsii.Boolean(false),
//   		NotificationArns: []*string{
//   			jsii.String("notificationArns"),
//   		},
//   		Output: jsii.String("output"),
//   		OutputsFile: jsii.String("outputsFile"),
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		PathMetadata: jsii.Boolean(false),
//   		Profile: jsii.String("profile"),
//   		Proxy: jsii.String("proxy"),
//   		RequireApproval: awscdk.Cloud_assembly_schema.RequireApproval_NEVER,
//   		ReuseAssets: []*string{
//   			jsii.String("reuseAssets"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		Rollback: jsii.Boolean(false),
//   		Stacks: []*string{
//   			jsii.String("stacks"),
//   		},
//   		Staging: jsii.Boolean(false),
//   		Strict: jsii.Boolean(false),
//   		ToolkitStackName: jsii.String("toolkitStackName"),
//   		Trace: jsii.Boolean(false),
//   		UsePreviousParameters: jsii.Boolean(false),
//   		Verbose: jsii.Boolean(false),
//   		VersionReporting: jsii.Boolean(false),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	ExpectedMessage: jsii.String("expectedMessage"),
//   	ExpectError: jsii.Boolean(false),
//   }
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

