package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-createdtimestamp
	//
	CreatedTimeStamp *float64 `field:"optional" json:"createdTimeStamp" yaml:"createdTimeStamp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-lastupdatetimestamp
	//
	LastUpdateTimeStamp *float64 `field:"optional" json:"lastUpdateTimeStamp" yaml:"lastUpdateTimeStamp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-statusreason
	//
	StatusReason interface{} `field:"optional" json:"statusReason" yaml:"statusReason"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipeline.html#cfn-observabilityadmin-telemetrypipelines-telemetrypipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

