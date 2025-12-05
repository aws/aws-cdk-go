package interfacesawsobservabilityadmin


// A reference to a TelemetryPipelines resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryPipelinesReference := &TelemetryPipelinesReference{
//   	PipelineIdentifier: jsii.String("pipelineIdentifier"),
//   	TelemetryPipelinesArn: jsii.String("telemetryPipelinesArn"),
//   }
//
type TelemetryPipelinesReference struct {
	// The PipelineIdentifier of the TelemetryPipelines resource.
	PipelineIdentifier *string `field:"required" json:"pipelineIdentifier" yaml:"pipelineIdentifier"`
	// The ARN of the TelemetryPipelines resource.
	TelemetryPipelinesArn *string `field:"required" json:"telemetryPipelinesArn" yaml:"telemetryPipelinesArn"`
}

