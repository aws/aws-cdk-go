package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Where Stack Instances will be created from the StackSet.
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
type StackInstances interface {
}

// The jsii proxy struct for StackInstances
type jsiiProxy_StackInstances struct {
	_ byte // padding
}

func NewStackInstances_Override(s StackInstances) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		nil, // no parameters
		s,
	)
}

// Create stack instances in a set of accounts or organizational units taken from the pipeline artifacts, and a set of regions  The file must be a JSON file containing a list of strings.
//
// For example:
//
// ```json
// [
//    "111111111111",
//    "222222222222",
//    "333333333333"
// ]
// ```
//
// Stack Instances will be created in every combination of region and account, or region and
// Organizational Units (OUs).
//
// If this is set of Organizational Units, you must have selected `StackSetDeploymentModel.organizations()`
// as deployment model.
func StackInstances_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath, regions *[]*string) StackInstances {
	_init_.Initialize()

	if err := validateStackInstances_FromArtifactPathParameters(artifactPath, regions); err != nil {
		panic(err)
	}
	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"fromArtifactPath",
		[]interface{}{artifactPath, regions},
		&returns,
	)

	return returns
}

// Create stack instances in a set of accounts and regions passed as literal lists.
//
// Stack Instances will be created in every combination of region and account.
//
// > NOTE: `StackInstances.inAccounts()` and `StackInstances.inOrganizationalUnits()`
// > have exactly the same behavior, and you can use them interchangeably if you want.
// > The only difference between them is that your code clearly indicates what entity
// > it's working with.
func StackInstances_InAccounts(accounts *[]*string, regions *[]*string) StackInstances {
	_init_.Initialize()

	if err := validateStackInstances_InAccountsParameters(accounts, regions); err != nil {
		panic(err)
	}
	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"inAccounts",
		[]interface{}{accounts, regions},
		&returns,
	)

	return returns
}

// Create stack instances in all accounts in a set of Organizational Units (OUs) and regions passed as literal lists.
//
// If you want to deploy to Organization Units, you must choose have created the StackSet
// with `deploymentModel: DeploymentModel.organizations()`.
//
// Stack Instances will be created in every combination of region and account.
//
// > NOTE: `StackInstances.inAccounts()` and `StackInstances.inOrganizationalUnits()`
// > have exactly the same behavior, and you can use them interchangeably if you want.
// > The only difference between them is that your code clearly indicates what entity
// > it's working with.
func StackInstances_InOrganizationalUnits(ous *[]*string, regions *[]*string) StackInstances {
	_init_.Initialize()

	if err := validateStackInstances_InOrganizationalUnitsParameters(ous, regions); err != nil {
		panic(err)
	}
	var returns StackInstances

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackInstances",
		"inOrganizationalUnits",
		[]interface{}{ous, regions},
		&returns,
	)

	return returns
}

