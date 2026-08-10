package interfacesawsosis


// A reference to a PipelineBlueprint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineBlueprintReference := &PipelineBlueprintReference{
//   	PipelineBlueprintArn: jsii.String("pipelineBlueprintArn"),
//   }
//
type PipelineBlueprintReference struct {
	// The Arn of the PipelineBlueprint resource.
	PipelineBlueprintArn *string `field:"required" json:"pipelineBlueprintArn" yaml:"pipelineBlueprintArn"`
}

