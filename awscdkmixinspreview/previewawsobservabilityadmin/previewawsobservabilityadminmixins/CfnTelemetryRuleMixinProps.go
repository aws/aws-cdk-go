package previewawsobservabilityadminmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTelemetryRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTelemetryRuleMixinProps := &CfnTelemetryRuleMixinProps{
//   	Rule: &TelemetryRuleProperty{
//   		DestinationConfiguration: &TelemetryDestinationConfigurationProperty{
//   			DestinationPattern: jsii.String("destinationPattern"),
//   			DestinationType: jsii.String("destinationType"),
//   			RetentionInDays: jsii.Number(123),
//   			VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   				LogFormat: jsii.String("logFormat"),
//   				MaxAggregationInterval: jsii.Number(123),
//   				TrafficType: jsii.String("trafficType"),
//   			},
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		SelectionCriteria: jsii.String("selectionCriteria"),
//   		TelemetryType: jsii.String("telemetryType"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryrule.html
//
type CfnTelemetryRuleMixinProps struct {
	// Retrieves the details of a specific telemetry rule in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryrule.html#cfn-observabilityadmin-telemetryrule-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of the telemetry rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryrule.html#cfn-observabilityadmin-telemetryrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Lists all tags attached to the specified telemetry rule resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryrule.html#cfn-observabilityadmin-telemetryrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

