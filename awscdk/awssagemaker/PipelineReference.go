package awssagemaker


// A reference to a Pipeline resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineReference := &PipelineReference{
//   	PipelineName: jsii.String("pipelineName"),
//   }
//
type PipelineReference struct {
	// The PipelineName of the Pipeline resource.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
}

