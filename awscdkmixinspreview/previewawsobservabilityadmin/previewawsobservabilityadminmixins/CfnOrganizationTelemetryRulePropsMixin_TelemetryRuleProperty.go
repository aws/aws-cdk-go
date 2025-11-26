package previewawsobservabilityadminmixins


// Defines how telemetry should be configured for specific AWS resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   telemetryRuleProperty := &TelemetryRuleProperty{
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
//   	ResourceType: jsii.String("resourceType"),
//   	Scope: jsii.String("scope"),
//   	SelectionCriteria: jsii.String("selectionCriteria"),
//   	TelemetryType: jsii.String("telemetryType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html
//
type CfnOrganizationTelemetryRulePropsMixin_TelemetryRuleProperty struct {
	// Configuration specifying where and how the telemetry data should be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// The type of AWS resource to configure telemetry for (e.g., "AWS::EC2::VPC").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The organizational scope to which the rule applies, specified using accounts or organizational units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Criteria for selecting which resources the rule applies to, such as resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-selectioncriteria
	//
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// The type of telemetry to collect (Logs, Metrics, or Traces).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-telemetrytype
	//
	TelemetryType *string `field:"optional" json:"telemetryType" yaml:"telemetryType"`
}

