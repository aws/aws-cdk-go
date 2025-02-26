package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines how IAM roles are created and managed.
//
// Example:
//   var pipeline pipeline
//   var sourceOutput artifact
//
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("DeployStackSets"),
//   	Actions: []iAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&CloudFormationDeployStackSetActionProps{
//   			ActionName: jsii.String("UpdateStackSet"),
//   			RunOrder: jsii.Number(1),
//   			StackSetName: jsii.String("MyStackSet"),
//   			Template: codepipeline_actions.StackSetTemplate_FromArtifactPath(sourceOutput.AtPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			DeploymentModel: codepipeline_actions.StackSetDeploymentModel_SelfManaged(),
//   			// This deploys to a set of accounts
//   			StackInstances: codepipeline_actions.StackInstances_InAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&CloudFormationDeployStackInstancesActionProps{
//   			ActionName: jsii.String("AddMoreInstances"),
//   			RunOrder: jsii.Number(2),
//   			StackSetName: jsii.String("MyStackSet"),
//   			StackInstances: codepipeline_actions.StackInstances_*InAccounts([]*string{
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
type StackSetDeploymentModel interface {
}

// The jsii proxy struct for StackSetDeploymentModel
type jsiiProxy_StackSetDeploymentModel struct {
	_ byte // padding
}

func NewStackSetDeploymentModel_Override(s StackSetDeploymentModel) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		nil, // no parameters
		s,
	)
}

// Deploy to AWS Organizations accounts.
//
// AWS CloudFormation StackSets automatically creates the IAM roles required
// to deploy to accounts managed by AWS Organizations. This requires an
// account to be a member of an Organization.
//
// Using this deployment model, you can specify either AWS Account Ids or
// Organization Unit Ids in the `stackInstances` parameter.
func StackSetDeploymentModel_Organizations(props *OrganizationsDeploymentProps) StackSetDeploymentModel {
	_init_.Initialize()

	if err := validateStackSetDeploymentModel_OrganizationsParameters(props); err != nil {
		panic(err)
	}
	var returns StackSetDeploymentModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		"organizations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Deploy to AWS Accounts not managed by AWS Organizations.
//
// You are responsible for creating Execution Roles in every account you will
// be deploying to in advance to create the actual stack instances. Unless you
// specify overrides, StackSets expects the execution roles you create to have
// the default name `AWSCloudFormationStackSetExecutionRole`. See the [Grant
// self-managed
// permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html)
// section of the CloudFormation documentation.
//
// The CDK will automatically create the central Administration Role in the
// Pipeline account which will be used to assume the Execution Role in each of
// the target accounts.
//
// If you wish to use a pre-created Administration Role, use `Role.fromRoleName()`
// or `Role.fromRoleArn()` to import it, and pass it to this function:
//
// ```ts
// const existingAdminRole = iam.Role.fromRoleName(this, 'AdminRole', 'AWSCloudFormationStackSetAdministrationRole');
//
// const deploymentModel = codepipeline_actions.StackSetDeploymentModel.selfManaged({
//   // Use an existing Role. Leave this out to create a new Role.
//   administrationRole: existingAdminRole,
// });
// ```
//
// Using this deployment model, you can only specify AWS Account Ids in the
// `stackInstances` parameter.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
//
func StackSetDeploymentModel_SelfManaged(props *SelfManagedDeploymentProps) StackSetDeploymentModel {
	_init_.Initialize()

	if err := validateStackSetDeploymentModel_SelfManagedParameters(props); err != nil {
		panic(err)
	}
	var returns StackSetDeploymentModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetDeploymentModel",
		"selfManaged",
		[]interface{}{props},
		&returns,
	)

	return returns
}

