package awsobservabilityadmin


// Configuration that determines which WAF log records to keep or drop based on specified conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFilterProperty := &LoggingFilterProperty{
//   	DefaultBehavior: jsii.String("defaultBehavior"),
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Behavior: jsii.String("behavior"),
//   			Conditions: []interface{}{
//   				&ConditionProperty{
//   					ActionCondition: &ActionConditionProperty{
//   						Action: jsii.String("action"),
//   					},
//   					LabelNameCondition: &LabelNameConditionProperty{
//   						LabelName: jsii.String("labelName"),
//   					},
//   				},
//   			},
//   			Requirement: jsii.String("requirement"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-loggingfilter.html
//
type CfnOrganizationTelemetryRule_LoggingFilterProperty struct {
	// The default action (KEEP or DROP) for log records that don't match any filter conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-loggingfilter.html#cfn-observabilityadmin-organizationtelemetryrule-loggingfilter-defaultbehavior
	//
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// A list of filter conditions that determine log record handling behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-loggingfilter.html#cfn-observabilityadmin-organizationtelemetryrule-loggingfilter-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

