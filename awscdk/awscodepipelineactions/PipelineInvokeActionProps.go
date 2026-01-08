package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodepipeline"
)

// Construction properties of the `PipelineInvokeAction`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   targetPipeline := codepipeline.Pipeline_FromPipelineArn(this, jsii.String("Pipeline"), jsii.String("arn:aws:codepipeline:us-east-1:123456789012:InvokePipelineAction")) // If targetPipeline is not created by cdk, import from arn.
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("stageName"),
//   	Actions: []IAction{
//   		cpactions.NewPipelineInvokeAction(&PipelineInvokeActionProps{
//   			ActionName: jsii.String("Invoke"),
//   			TargetPipeline: *TargetPipeline,
//   			Variables: []Variable{
//   				&Variable{
//   					Name: jsii.String("name1"),
//   					Value: jsii.String("value1"),
//   				},
//   			},
//   			SourceRevisions: []SourceRevision{
//   				&SourceRevision{
//   					ActionName: jsii.String("Source"),
//   					RevisionType: cpactions.RevisionType_S3_OBJECT_VERSION_ID,
//   					RevisionValue: jsii.String("testRevisionValue"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
type PipelineInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The pipeline that will, upon running, start the current target pipeline.
	//
	// You must have already created the invoking pipeline.
	TargetPipeline interfacesawscodepipeline.IPipelineRef `field:"required" json:"targetPipeline" yaml:"targetPipeline"`
	// The source revisions that you want the target pipeline to use when it is started by the invoking pipeline.
	// Default: - no specific revisions.
	//
	SourceRevisions *[]*SourceRevision `field:"optional" json:"sourceRevisions" yaml:"sourceRevisions"`
	// The names and values of variables that you want the action to support.
	// Default: - no specific variable.
	//
	Variables *[]*Variable `field:"optional" json:"variables" yaml:"variables"`
}

