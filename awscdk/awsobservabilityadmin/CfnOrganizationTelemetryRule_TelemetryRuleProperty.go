package awsobservabilityadmin


// The telemetry rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryRuleProperty := &TelemetryRuleProperty{
//   	ResourceType: jsii.String("resourceType"),
//   	TelemetryType: jsii.String("telemetryType"),
//
//   	// the properties below are optional
//   	DestinationConfiguration: &TelemetryDestinationConfigurationProperty{
//   		DestinationPattern: jsii.String("destinationPattern"),
//   		DestinationType: jsii.String("destinationType"),
//   		RetentionInDays: jsii.Number(123),
//   		VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   			LogFormat: jsii.String("logFormat"),
//   			MaxAggregationInterval: jsii.Number(123),
//   			TrafficType: jsii.String("trafficType"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	SelectionCriteria: jsii.String("selectionCriteria"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html
//
type CfnOrganizationTelemetryRule_TelemetryRuleProperty struct {
	// Resource Type associated with the Organization Telemetry Rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Telemetry Type associated with the Organization Telemetry Rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-telemetrytype
	//
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// The destination configuration for telemetry data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Selection Criteria on scope level for rule application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Selection Criteria on resource level for rule application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-selectioncriteria
	//
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

