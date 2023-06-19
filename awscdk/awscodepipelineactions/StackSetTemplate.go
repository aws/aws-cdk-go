package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// The source of a StackSet template.
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
// Experimental.
type StackSetTemplate interface {
}

// The jsii proxy struct for StackSetTemplate
type jsiiProxy_StackSetTemplate struct {
	_ byte // padding
}

// Experimental.
func NewStackSetTemplate_Override(s StackSetTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.StackSetTemplate",
		nil, // no parameters
		s,
	)
}

// Use a file in an artifact as Stack Template.
// Experimental.
func StackSetTemplate_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath) StackSetTemplate {
	_init_.Initialize()

	if err := validateStackSetTemplate_FromArtifactPathParameters(artifactPath); err != nil {
		panic(err)
	}
	var returns StackSetTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline_actions.StackSetTemplate",
		"fromArtifactPath",
		[]interface{}{artifactPath},
		&returns,
	)

	return returns
}

