package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineDefinitionProperty := &PipelineDefinitionProperty{
//   	PipelineDefinitionBody: jsii.String("pipelineDefinitionBody"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-pipelinedefinition.html
//
type CfnPipeline_PipelineDefinitionProperty struct {
	// A specification that defines the pipeline in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-pipelinedefinition.html#cfn-sagemaker-pipeline-pipelinedefinition-pipelinedefinitionbody
	//
	PipelineDefinitionBody *string `field:"required" json:"pipelineDefinitionBody" yaml:"pipelineDefinitionBody"`
}

