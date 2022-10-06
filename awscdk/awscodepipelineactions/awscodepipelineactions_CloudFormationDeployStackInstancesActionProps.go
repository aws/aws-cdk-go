package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for the CloudFormationDeployStackInstancesAction.
//
// Example:
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
// Experimental.
type CloudFormationDeployStackInstancesActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	// Experimental.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	// Experimental.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
	// Specify where to create or update Stack Instances.
	//
	// You can specify either AWS Accounts Ids or AWS Organizations Organizational Units.
	// Experimental.
	StackInstances StackInstances `field:"required" json:"stackInstances" yaml:"stackInstances"`
	// The name of the StackSet we are adding instances to.
	// Experimental.
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// Parameter values that only apply to the current Stack Instances.
	//
	// These parameters are shared between all instances added by this action.
	// Experimental.
	ParameterOverrides StackSetParameters `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
}

