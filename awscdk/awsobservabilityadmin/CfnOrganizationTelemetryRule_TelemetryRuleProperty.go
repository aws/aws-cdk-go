package awsobservabilityadmin


// Defines how telemetry should be configured for specific AWS resources.
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
//   		CloudtrailParameters: &CloudtrailParametersProperty{
//   			AdvancedEventSelectors: []interface{}{
//   				&AdvancedEventSelectorProperty{
//   					FieldSelectors: []interface{}{
//   						&AdvancedFieldSelectorProperty{
//   							EndsWith: []*string{
//   								jsii.String("endsWith"),
//   							},
//   							EqualTo: []*string{
//   								jsii.String("equalTo"),
//   							},
//   							Field: jsii.String("field"),
//   							NotEndsWith: []*string{
//   								jsii.String("notEndsWith"),
//   							},
//   							NotEquals: []*string{
//   								jsii.String("notEquals"),
//   							},
//   							NotStartsWith: []*string{
//   								jsii.String("notStartsWith"),
//   							},
//   							StartsWith: []*string{
//   								jsii.String("startsWith"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		DestinationPattern: jsii.String("destinationPattern"),
//   		DestinationType: jsii.String("destinationType"),
//   		ElbLoadBalancerLoggingParameters: &ELBLoadBalancerLoggingParametersProperty{
//   			FieldDelimiter: jsii.String("fieldDelimiter"),
//   			OutputFormat: jsii.String("outputFormat"),
//   		},
//   		RetentionInDays: jsii.Number(123),
//   		VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   			LogFormat: jsii.String("logFormat"),
//   			MaxAggregationInterval: jsii.Number(123),
//   			TrafficType: jsii.String("trafficType"),
//   		},
//   		WafLoggingParameters: &WAFLoggingParametersProperty{
//   			LoggingFilter: &LoggingFilterProperty{
//   				DefaultBehavior: jsii.String("defaultBehavior"),
//   				Filters: []interface{}{
//   					&FilterProperty{
//   						Behavior: jsii.String("behavior"),
//   						Conditions: []interface{}{
//   							&ConditionProperty{
//   								ActionCondition: &ActionConditionProperty{
//   									Action: jsii.String("action"),
//   								},
//   								LabelNameCondition: &LabelNameConditionProperty{
//   									LabelName: jsii.String("labelName"),
//   								},
//   							},
//   						},
//   						Requirement: jsii.String("requirement"),
//   					},
//   				},
//   			},
//   			LogType: jsii.String("logType"),
//   			RedactedFields: []interface{}{
//   				&FieldToMatchProperty{
//   					Method: jsii.String("method"),
//   					QueryString: jsii.String("queryString"),
//   					SingleHeader: &SingleHeaderProperty{
//   						Name: jsii.String("name"),
//   					},
//   					UriPath: jsii.String("uriPath"),
//   				},
//   			},
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	SelectionCriteria: jsii.String("selectionCriteria"),
//   	TelemetrySourceTypes: []*string{
//   		jsii.String("telemetrySourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html
//
type CfnOrganizationTelemetryRule_TelemetryRuleProperty struct {
	// The type of AWS resource to configure telemetry for (e.g., "AWS::EC2::VPC", "AWS::EKS::Cluster", "AWS::WAFv2::WebACL").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The type of telemetry to collect (Logs, Metrics, or Traces).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-telemetrytype
	//
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Configuration specifying where and how the telemetry data should be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// The organizational scope to which the rule applies, specified using accounts or organizational units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Criteria for selecting which resources the rule applies to, such as resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-selectioncriteria
	//
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// The specific telemetry source types to configure for the resource, such as VPC_FLOW_LOGS or EKS_AUDIT_LOGS.
	//
	// TelemetrySourceTypes must be correlated with the specific resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-telemetryrule-telemetrysourcetypes
	//
	TelemetrySourceTypes *[]*string `field:"optional" json:"telemetrySourceTypes" yaml:"telemetrySourceTypes"`
}

