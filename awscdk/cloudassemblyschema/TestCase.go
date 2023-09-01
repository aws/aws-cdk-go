package cloudassemblyschema


// Represents an integration test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   testCase := &TestCase{
//   	Stacks: []*string{
//   		jsii.String("stacks"),
//   	},
//
//   	// the properties below are optional
//   	AllowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	AssertionStack: jsii.String("assertionStack"),
//   	AssertionStackName: jsii.String("assertionStackName"),
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
//   				Concurrency: jsii.Number(123),
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
//   }
//
type TestCase struct {
	// List of CloudFormation resource types in this stack that can be destroyed as part of an update without failing the test.
	//
	// This list should only include resources that for this specific
	// integration test we are sure will not cause errors or an outage if
	// destroyed. For example, maybe we know that a new resource will be created
	// first before the old resource is destroyed which prevents any outage.
	//
	// e.g. ['AWS::IAM::Role']
	// Default: - do not allow destruction of any resources on update.
	//
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Default: - runner default options.
	//
	CdkCommandOptions *CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Default: false.
	//
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Default: - no hooks.
	//
	Hooks *Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Default: - can run in any region.
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Default: true.
	//
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// Stacks that should be tested as part of this test case The stackNames will be passed as args to the cdk commands so dependent stacks will be automatically deployed unless `exclusively` is passed.
	Stacks *[]*string `field:"required" json:"stacks" yaml:"stacks"`
	// The node id of the stack that contains assertions.
	//
	// This is the value that can be used to deploy the stack with the CDK CLI.
	// Default: - no assertion stack.
	//
	AssertionStack *string `field:"optional" json:"assertionStack" yaml:"assertionStack"`
	// The name of the stack that contains assertions.
	// Default: - no assertion stack.
	//
	AssertionStackName *string `field:"optional" json:"assertionStackName" yaml:"assertionStackName"`
}

