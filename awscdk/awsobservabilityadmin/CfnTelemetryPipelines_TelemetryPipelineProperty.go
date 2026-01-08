package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents a complete telemetry pipeline resource with configuration, status, and metadata for data processing and transformation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryPipelineProperty := &TelemetryPipelineProperty{
//   	Arn: jsii.String("arn"),
//   	Configuration: &TelemetryPipelineConfigurationProperty{
//   		Body: jsii.String("body"),
//   	},
//   	CreatedTimeStamp: jsii.Number(123),
//   	LastUpdateTimeStamp: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   	StatusReason: &TelemetryPipelineStatusReasonProperty{
//   		Description: jsii.String("description"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html
//
type CfnTelemetryPipelines_TelemetryPipelineProperty struct {
	// The Amazon Resource Name (ARN) of the telemetry pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The configuration that defines how the telemetry pipeline processes data.
	//
	// For more information, see the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Creating-pipelines.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The timestamp when the telemetry pipeline was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-createdtimestamp
	//
	CreatedTimeStamp *float64 `field:"optional" json:"createdTimeStamp" yaml:"createdTimeStamp"`
	// The timestamp when the telemetry pipeline was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-lastupdatetimestamp
	//
	LastUpdateTimeStamp *float64 `field:"optional" json:"lastUpdateTimeStamp" yaml:"lastUpdateTimeStamp"`
	// The name of the telemetry pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current status of the telemetry pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Additional information about the pipeline status, including reasons for failure states.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-statusreason
	//
	StatusReason interface{} `field:"optional" json:"statusReason" yaml:"statusReason"`
	// The key-value pairs associated with the telemetry pipeline resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

