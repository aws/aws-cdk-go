package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// SageMaker target properties.
//
// Example:
//   var sourceQueue queue
//   var targetPipeline iPipeline
//
//
//   pipelineTarget := targets.NewSageMakerTarget(targetPipeline, &SageMakerTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipelineTarget,
//   })
//
// Experimental.
type SageMakerTargetParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// List of parameter names and values for SageMaker Model Building Pipeline execution.
	//
	// The Name/Value pairs are passed to start execution of the pipeline.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsagemakerpipelineparameters.html#cfn-pipes-pipe-pipetargetsagemakerpipelineparameters-pipelineparameterlist
	//
	// Default: - none.
	//
	// Experimental.
	PipelineParameters *map[string]*string `field:"optional" json:"pipelineParameters" yaml:"pipelineParameters"`
}

