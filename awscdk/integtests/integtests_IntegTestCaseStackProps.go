package integtests

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

// Properties of an integration test case stack.
//
// Example:
//   var app app
//   var stackUnderTest stack
//
//   testCaseWithAssets := awscdk.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &IntegTestCaseStackProps{
//   	DiffAssets: jsii.Boolean(true),
//   })
//
//   awscdk.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
//   	TestCases: []*stack{
//   		stackUnderTest,
//   		testCaseWithAssets,
//   	},
//   })
//
// Experimental.
type IntegTestCaseStackProps struct {
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
	// Include runtime versioning information in this Stack.
	// Experimental.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// A description of the stack.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   new Stack(app, 'Stack1', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     },
	//   });
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   new Stack(app, 'Stack2', {
	//     env: {
	//       account: process.env.CDK_DEFAULT_ACCOUNT,
	//       region: process.env.CDK_DEFAULT_REGION
	//     },
	//   });
	//
	//   // Define multiple stacks stage associated with an environment
	//   const myStage = new Stage(app, 'MyStage', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     }
	//   });
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   new MyStack(myStage, 'Stack1');
	//   new YourStack(myStage, 'Stack2');
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   new MyStack(app, 'Stack1');
	//
	// Experimental.
	Env *awscdk.Environment `field:"optional" json:"env" yaml:"env"`
	// Name to deploy the stack with.
	// Experimental.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Synthesis method to use while deploying this stack.
	// Experimental.
	Synthesizer awscdk.IStackSynthesizer `field:"optional" json:"synthesizer" yaml:"synthesizer"`
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Experimental.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
}

