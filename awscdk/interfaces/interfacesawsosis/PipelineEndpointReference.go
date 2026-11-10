package interfacesawsosis


// A reference to a PipelineEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineEndpointReference := &PipelineEndpointReference{
//   	PipelineEndpointArn: jsii.String("pipelineEndpointArn"),
//   }
//
type PipelineEndpointReference struct {
	// The Arn of the PipelineEndpoint resource.
	PipelineEndpointArn *string `field:"required" json:"pipelineEndpointArn" yaml:"pipelineEndpointArn"`
}

