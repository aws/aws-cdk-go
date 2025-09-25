package awsosis


// A reference to a Pipeline resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineReference := &PipelineReference{
//   	PipelineArn: jsii.String("pipelineArn"),
//   }
//
type PipelineReference struct {
	// The PipelineArn of the Pipeline resource.
	PipelineArn *string `field:"required" json:"pipelineArn" yaml:"pipelineArn"`
}

