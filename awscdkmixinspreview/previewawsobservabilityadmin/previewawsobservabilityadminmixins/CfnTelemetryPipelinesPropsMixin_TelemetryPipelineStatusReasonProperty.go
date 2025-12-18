package previewawsobservabilityadminmixins


// Provides detailed information about the status of a telemetry pipeline, including reasons for specific states.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   telemetryPipelineStatusReasonProperty := &TelemetryPipelineStatusReasonProperty{
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelinestatusreason.html
//
type CfnTelemetryPipelinesPropsMixin_TelemetryPipelineStatusReasonProperty struct {
	// A description of the pipeline status reason, providing additional context about the current state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelinestatusreason.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipelinestatusreason-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

