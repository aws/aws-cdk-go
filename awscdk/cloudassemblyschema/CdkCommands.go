package cloudassemblyschema


// Options for specific cdk commands that are run as part of the integration test workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdkCommands := &CdkCommands{
//   	Deploy: &DeployCommand{
//   		Args: &DeployOptions{
//   			All: jsii.Boolean(false),
//   			App: jsii.String("app"),
//   			AssetMetadata: jsii.Boolean(false),
//   			CaBundlePath: jsii.String("caBundlePath"),
//   			ChangeSetName: jsii.String("changeSetName"),
//   			Ci: jsii.Boolean(false),
//   			Color: jsii.Boolean(false),
//   			Concurrency: jsii.Number(123),
//   			Context: map[string]*string{
//   				"contextKey": jsii.String("context"),
//   			},
//   			Debug: jsii.Boolean(false),
//   			Ec2Creds: jsii.Boolean(false),
//   			Exclusively: jsii.Boolean(false),
//   			Execute: jsii.Boolean(false),
//   			Force: jsii.Boolean(false),
//   			IgnoreErrors: jsii.Boolean(false),
//   			Json: jsii.Boolean(false),
//   			Lookups: jsii.Boolean(false),
//   			Notices: jsii.Boolean(false),
//   			NotificationArns: []*string{
//   				jsii.String("notificationArns"),
//   			},
//   			Output: jsii.String("output"),
//   			OutputsFile: jsii.String("outputsFile"),
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			PathMetadata: jsii.Boolean(false),
//   			Profile: jsii.String("profile"),
//   			Proxy: jsii.String("proxy"),
//   			RequireApproval: awscdk.Cloud_assembly_schema.RequireApproval_NEVER,
//   			ReuseAssets: []*string{
//   				jsii.String("reuseAssets"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			Rollback: jsii.Boolean(false),
//   			Stacks: []*string{
//   				jsii.String("stacks"),
//   			},
//   			Staging: jsii.Boolean(false),
//   			Strict: jsii.Boolean(false),
//   			ToolkitStackName: jsii.String("toolkitStackName"),
//   			Trace: jsii.Boolean(false),
//   			UsePreviousParameters: jsii.Boolean(false),
//   			Verbose: jsii.Boolean(false),
//   			VersionReporting: jsii.Boolean(false),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		ExpectedMessage: jsii.String("expectedMessage"),
//   		ExpectError: jsii.Boolean(false),
//   	},
//   	Destroy: &DestroyCommand{
//   		Args: &DestroyOptions{
//   			All: jsii.Boolean(false),
//   			App: jsii.String("app"),
//   			AssetMetadata: jsii.Boolean(false),
//   			CaBundlePath: jsii.String("caBundlePath"),
//   			Color: jsii.Boolean(false),
//   			Context: map[string]*string{
//   				"contextKey": jsii.String("context"),
//   			},
//   			Debug: jsii.Boolean(false),
//   			Ec2Creds: jsii.Boolean(false),
//   			Exclusively: jsii.Boolean(false),
//   			Force: jsii.Boolean(false),
//   			IgnoreErrors: jsii.Boolean(false),
//   			Json: jsii.Boolean(false),
//   			Lookups: jsii.Boolean(false),
//   			Notices: jsii.Boolean(false),
//   			Output: jsii.String("output"),
//   			PathMetadata: jsii.Boolean(false),
//   			Profile: jsii.String("profile"),
//   			Proxy: jsii.String("proxy"),
//   			RoleArn: jsii.String("roleArn"),
//   			Stacks: []*string{
//   				jsii.String("stacks"),
//   			},
//   			Staging: jsii.Boolean(false),
//   			Strict: jsii.Boolean(false),
//   			Trace: jsii.Boolean(false),
//   			Verbose: jsii.Boolean(false),
//   			VersionReporting: jsii.Boolean(false),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		ExpectedMessage: jsii.String("expectedMessage"),
//   		ExpectError: jsii.Boolean(false),
//   	},
//   }
//
type CdkCommands struct {
	// Options to for the cdk deploy command.
	// Default: - default deploy options.
	//
	Deploy *DeployCommand `field:"optional" json:"deploy" yaml:"deploy"`
	// Options to for the cdk destroy command.
	// Default: - default destroy options.
	//
	Destroy *DestroyCommand `field:"optional" json:"destroy" yaml:"destroy"`
}

