package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTelemetryPipelines`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTelemetryPipelinesProps := &CfnTelemetryPipelinesProps{
//   	Configuration: &TelemetryPipelineConfigurationProperty{
//   		Body: jsii.String("body"),
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html
//
type CfnTelemetryPipelinesProps struct {
	// The configuration that defines how the telemetry pipeline processes data, including sources, processors, and destinations.
	//
	// For more information, see the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Creating-pipelines.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The name of the telemetry pipeline to create.
	//
	// The name must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The key-value pairs to associate with the telemetry pipeline resource for categorization and management purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

