package previewawsobservabilityadminmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOrganizationTelemetryRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationTelemetryRuleMixinProps := &CfnOrganizationTelemetryRuleMixinProps{
//   	Rule: &TelemetryRuleProperty{
//   		DestinationConfiguration: &TelemetryDestinationConfigurationProperty{
//   			CloudtrailParameters: &CloudtrailParametersProperty{
//   				AdvancedEventSelectors: []interface{}{
//   					&AdvancedEventSelectorProperty{
//   						FieldSelectors: []interface{}{
//   							&AdvancedFieldSelectorProperty{
//   								EndsWith: []*string{
//   									jsii.String("endsWith"),
//   								},
//   								EqualTo: []*string{
//   									jsii.String("equalTo"),
//   								},
//   								Field: jsii.String("field"),
//   								NotEndsWith: []*string{
//   									jsii.String("notEndsWith"),
//   								},
//   								NotEquals: []*string{
//   									jsii.String("notEquals"),
//   								},
//   								NotStartsWith: []*string{
//   									jsii.String("notStartsWith"),
//   								},
//   								StartsWith: []*string{
//   									jsii.String("startsWith"),
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			DestinationPattern: jsii.String("destinationPattern"),
//   			DestinationType: jsii.String("destinationType"),
//   			ElbLoadBalancerLoggingParameters: &ELBLoadBalancerLoggingParametersProperty{
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				OutputFormat: jsii.String("outputFormat"),
//   			},
//   			RetentionInDays: jsii.Number(123),
//   			VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   				LogFormat: jsii.String("logFormat"),
//   				MaxAggregationInterval: jsii.Number(123),
//   				TrafficType: jsii.String("trafficType"),
//   			},
//   			WafLoggingParameters: &WAFLoggingParametersProperty{
//   				LoggingFilter: &LoggingFilterProperty{
//   					DefaultBehavior: jsii.String("defaultBehavior"),
//   					Filters: []interface{}{
//   						&FilterProperty{
//   							Behavior: jsii.String("behavior"),
//   							Conditions: []interface{}{
//   								&ConditionProperty{
//   									ActionCondition: &ActionConditionProperty{
//   										Action: jsii.String("action"),
//   									},
//   									LabelNameCondition: &LabelNameConditionProperty{
//   										LabelName: jsii.String("labelName"),
//   									},
//   								},
//   							},
//   							Requirement: jsii.String("requirement"),
//   						},
//   					},
//   				},
//   				LogType: jsii.String("logType"),
//   				RedactedFields: []interface{}{
//   					&FieldToMatchProperty{
//   						Method: jsii.String("method"),
//   						QueryString: jsii.String("queryString"),
//   						SingleHeader: &SingleHeaderProperty{
//   							Name: jsii.String("name"),
//   						},
//   						UriPath: jsii.String("uriPath"),
//   					},
//   				},
//   			},
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		Scope: jsii.String("scope"),
//   		SelectionCriteria: jsii.String("selectionCriteria"),
//   		TelemetrySourceTypes: []*string{
//   			jsii.String("telemetrySourceTypes"),
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html
//
type CfnOrganizationTelemetryRuleMixinProps struct {
	// The name of the organization telemetry rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of the organization centralization rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Lists all tags attached to the specified resource.
	//
	// Supports telemetry rule resources and telemetry pipeline resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

