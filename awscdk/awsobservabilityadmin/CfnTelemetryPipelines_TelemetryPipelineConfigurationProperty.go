package awsobservabilityadmin


// Defines the configuration for a pipeline, including how data flows from sources through processors to destinations.
//
// The configuration is specified in YAML format and must include a valid pipeline definition with required source and sink components. This pipeline enables end-to-end telemetry data collection, transformation, and delivery while supporting optional processing steps and extensions for enhanced functionality.
//
// The primary pipeline configuration section are:
//
// - *Source:* Defines where log data originates from (S3 buckets, CloudWatch Logs, third-party APIs). Each pipeline must have exactly one source.
// - *Processors (optional):* Transform, parse, and enrich log data as it flows through the pipeline. Processors are applied sequentially in the order they are defined.
// - *Sink:* Defines the destination where processed log data is sent. Each pipeline must have exactly one sink.
// - *Extensions (optional):* Provide additional functionality such as AWS Secrets Manager integration for credential management.
//
// For more details on each configuration section see [CloudWatch pipelines User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-pipelines.html) . Additional comprehensive configuration examples can be found in the [CreateTelemetryPipeline API docs](https://docs.aws.amazon.com/cloudwatch/latest/observabilityadmin/API_CreateTelemetryPipeline.html#API_CreateTelemetryPipeline_Examples) .
//
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
	// The pipeline configuration body that defines the data processing rules and transformations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration-body
	//
	Body *string `field:"required" json:"body" yaml:"body"`
}

