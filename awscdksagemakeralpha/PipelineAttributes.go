package awscdksagemakeralpha


// Attributes for importing a SageMaker Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   pipelineAttributes := &PipelineAttributes{
//   	Account: jsii.String("account"),
//   	PipelineArn: jsii.String("pipelineArn"),
//   	PipelineName: jsii.String("pipelineName"),
//   	Region: jsii.String("region"),
//   }
//
// Experimental.
type PipelineAttributes struct {
	// The account the pipeline is in.
	// Default: - same account as the stack.
	//
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The ARN of the pipeline.
	// Default: - Either this or pipelineName must be provided.
	//
	// Experimental.
	PipelineArn *string `field:"optional" json:"pipelineArn" yaml:"pipelineArn"`
	// The name of the pipeline.
	// Default: - Either this or pipelineArn must be provided.
	//
	// Experimental.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// The region the pipeline is in.
	// Default: - same region as the stack.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

