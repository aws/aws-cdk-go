package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Integration test properties.
//
// Example:
//   var lambdaFunction iFunction
//   var app app
//
//
//   stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
//   	TestCases: []stack{
//   		stack,
//   	},
//   })
//
//   invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
//   	FunctionName: lambdaFunction.FunctionName,
//   })
//   invoke.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"Payload": jsii.String("200"),
//   }))
//
// Experimental.
type IntegTestProps struct {
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
	// Experimental.
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Default: - runner default options.
	//
	// Experimental.
	CdkCommandOptions *cloudassemblyschema.CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Default: false.
	//
	// Experimental.
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Default: - no hooks.
	//
	// Experimental.
	Hooks *cloudassemblyschema.Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Default: - can run in any region.
	//
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Default: true.
	//
	// Experimental.
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// List of test cases that make up this test.
	// Experimental.
	TestCases *[]awscdk.Stack `field:"required" json:"testCases" yaml:"testCases"`
	// Specify a stack to use for assertions.
	// Default: - a stack is created for you.
	//
	// Experimental.
	AssertionStack awscdk.Stack `field:"optional" json:"assertionStack" yaml:"assertionStack"`
	// Enable lookups for this test.
	//
	// If lookups are enabled
	// then `stackUpdateWorkflow` must be set to false.
	// Lookups should only be enabled when you are explicitly testing
	// lookups.
	// Default: false.
	//
	// Experimental.
	EnableLookups *bool `field:"optional" json:"enableLookups" yaml:"enableLookups"`
}

