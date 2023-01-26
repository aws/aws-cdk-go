package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the CloudFormationDeployStackSetAction.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("DeployStackSets"),
//   	actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
//   			actionName: jsii.String("UpdateStackSet"),
//   			runOrder: jsii.Number(1),
//   			stackSetName: jsii.String("MyStackSet"),
//   			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
//   			// This deploys to a set of accounts
//   			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
//   			actionName: jsii.String("AddMoreInstances"),
//   			runOrder: jsii.Number(2),
//   			stackSetName: jsii.String("MyStackSet"),
//   			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
	// The name to associate with the stack set.
	//
	// This name must be unique in the Region where it is created.
	//
	// The name may only contain alphanumeric and hyphen characters. It must begin with an alphabetic character and be 128 characters or fewer.
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// The location of the template that defines the resources in the stack set.
	//
	// This must point to a template with a maximum size of 460,800 bytes.
	//
	// Enter the path to the source artifact name and template file.
	Template StackSetTemplate `field:"required" json:"template" yaml:"template"`
	// Indicates that the template can create and update resources, depending on the types of resources in the template.
	//
	// You must use this property if you have IAM resources in your stack template or you create a stack directly from a template containing macros.
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// Determines how IAM roles are created and managed.
	//
	// The choices are:
	//
	// - Self Managed: you create IAM roles with the required permissions
	//    in the administration account and all target accounts.
	// - Service Managed: only available if the account and target accounts
	//    are part of an AWS Organization. The necessary roles will be created
	//    for you.
	//
	// If you want to deploy to all accounts that are a member of AWS
	// Organizations Organizational Units (OUs), you must select Service Managed
	// permissions.
	//
	// Note: This parameter can only be changed when no stack instances exist in
	// the stack set.
	DeploymentModel StackSetDeploymentModel `field:"optional" json:"deploymentModel" yaml:"deploymentModel"`
	// A description of the stack set.
	//
	// You can use this to describe the stack set’s purpose or other relevant information.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The template parameters for your stack set.
	//
	// These parameters are shared between all instances of the stack set.
	Parameters StackSetParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// Specify where to create or update Stack Instances.
	//
	// You can specify either AWS Accounts Ids or AWS Organizations Organizational Units.
	StackInstances StackInstances `field:"optional" json:"stackInstances" yaml:"stackInstances"`
}

