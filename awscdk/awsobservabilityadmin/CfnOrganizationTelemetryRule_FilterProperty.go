package awsobservabilityadmin


// A single filter condition that specifies behavior, requirement, and matching conditions for WAF log records.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Behavior: jsii.String("behavior"),
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			ActionCondition: &ActionConditionProperty{
//   				Action: jsii.String("action"),
//   			},
//   			LabelNameCondition: &LabelNameConditionProperty{
//   				LabelName: jsii.String("labelName"),
//   			},
//   		},
//   	},
//   	Requirement: jsii.String("requirement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-filter.html
//
type CfnOrganizationTelemetryRule_FilterProperty struct {
	// The action to take for log records matching this filter (KEEP or DROP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-filter.html#cfn-observabilityadmin-organizationtelemetryrule-filter-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The list of conditions that determine if a log record matches this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-filter.html#cfn-observabilityadmin-organizationtelemetryrule-filter-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether the log record must meet all conditions (MEETS_ALL) or any condition (MEETS_ANY) to match this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-filter.html#cfn-observabilityadmin-organizationtelemetryrule-filter-requirement
	//
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

