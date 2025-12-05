package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryPipelineConfigurationProperty := &TelemetryPipelineConfigurationProperty{
//   	Body: jsii.String("body"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html
//
type CfnTelemetryPipelines_TelemetryPipelineConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration-body
	//
	Body *string `field:"required" json:"body" yaml:"body"`
}

