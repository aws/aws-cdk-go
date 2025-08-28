package awsobservabilityadmin


// The destination configuration for telemetry data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryDestinationConfigurationProperty := &TelemetryDestinationConfigurationProperty{
//   	DestinationPattern: jsii.String("destinationPattern"),
//   	DestinationType: jsii.String("destinationType"),
//   	RetentionInDays: jsii.Number(123),
//   	VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   		LogFormat: jsii.String("logFormat"),
//   		MaxAggregationInterval: jsii.Number(123),
//   		TrafficType: jsii.String("trafficType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html
//
type CfnOrganizationTelemetryRule_TelemetryDestinationConfigurationProperty struct {
	// Pattern for telemetry data destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-destinationpattern
	//
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Type of telemetry destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-destinationtype
	//
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Number of days to retain the telemetry data in the specified destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Telemetry parameters for VPC Flow logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-vpcflowlogparameters
	//
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
}

