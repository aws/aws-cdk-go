package previewawsobservabilityadminmixins


// Defines the configuration for a telemetry pipeline, including how data flows from sources through processors to destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   telemetryPipelineConfigurationProperty := &TelemetryPipelineConfigurationProperty{
//   	Body: jsii.String("body"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html
//
type CfnTelemetryPipelinesPropsMixin_TelemetryPipelineConfigurationProperty struct {
	// The pipeline configuration body that defines the data processing rules and transformations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
}

