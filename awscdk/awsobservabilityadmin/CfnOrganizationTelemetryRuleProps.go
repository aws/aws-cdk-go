package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOrganizationTelemetryRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationTelemetryRuleProps := &CfnOrganizationTelemetryRuleProps{
//   	Rule: &TelemetryRuleProperty{
//   		ResourceType: jsii.String("resourceType"),
//   		TelemetryType: jsii.String("telemetryType"),
//
//   		// the properties below are optional
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
//   		Scope: jsii.String("scope"),
//   		SelectionCriteria: jsii.String("selectionCriteria"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html
//
type CfnOrganizationTelemetryRuleProps struct {
	// The telemetry rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The name of the organization telemetry rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationtelemetryrule.html#cfn-observabilityadmin-organizationtelemetryrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

