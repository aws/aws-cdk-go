package awsdatapipeline


// A reference to a Pipeline resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineReference := &PipelineReference{
//   	PipelineId: jsii.String("pipelineId"),
//   }
//
type PipelineReference struct {
	// The PipelineId of the Pipeline resource.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
}

