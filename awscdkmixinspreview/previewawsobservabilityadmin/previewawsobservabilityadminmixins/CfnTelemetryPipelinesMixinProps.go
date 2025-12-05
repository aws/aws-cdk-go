package previewawsobservabilityadminmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTelemetryPipelinesPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTelemetryPipelinesMixinProps := &CfnTelemetryPipelinesMixinProps{
//   	Configuration: &TelemetryPipelineConfigurationProperty{
//   		Body: jsii.String("body"),
//   	},
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
type CfnTelemetryPipelinesMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetrypipelines.html#cfn-observabilityadmin-telemetrypipelines-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

