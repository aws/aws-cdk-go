package awscdksagemakeralpha


// Properties for defining a SageMaker Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   pipelineProps := &PipelineProps{
//   	PipelineName: jsii.String("pipelineName"),
//   }
//
// Experimental.
type PipelineProps struct {
	// The physical name of the pipeline.
	// Default: - CDK generated name.
	//
	// Experimental.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
}

