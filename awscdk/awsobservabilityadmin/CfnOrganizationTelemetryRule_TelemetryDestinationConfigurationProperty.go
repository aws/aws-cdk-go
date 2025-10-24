package awsobservabilityadmin


// Configuration specifying where and how telemetry data should be delivered for AWS resources.
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
	// The pattern used to generate the destination path or name, supporting macros like <resourceId> and <accountId>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-destinationpattern
	//
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// The type of destination for the telemetry data (e.g., "Amazon CloudWatch Logs", "S3").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-destinationtype
	//
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// The number of days to retain the telemetry data in the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Configuration parameters specific to VPC Flow Logs when VPC is the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration-vpcflowlogparameters
	//
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
}

