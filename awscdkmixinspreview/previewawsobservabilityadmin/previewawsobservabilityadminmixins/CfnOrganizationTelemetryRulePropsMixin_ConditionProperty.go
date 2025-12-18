package previewawsobservabilityadminmixins


// A single condition that can match based on WAF rule action or label name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionProperty := &ConditionProperty{
//   	ActionCondition: &ActionConditionProperty{
//   		Action: jsii.String("action"),
//   	},
//   	LabelNameCondition: &LabelNameConditionProperty{
//   		LabelName: jsii.String("labelName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-condition.html
//
type CfnOrganizationTelemetryRulePropsMixin_ConditionProperty struct {
	// Matches log records based on the WAF rule action taken (ALLOW, BLOCK, COUNT, etc.).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-condition.html#cfn-observabilityadmin-organizationtelemetryrule-condition-actioncondition
	//
	ActionCondition interface{} `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// Matches log records based on WAF rule labels applied to the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-condition.html#cfn-observabilityadmin-organizationtelemetryrule-condition-labelnamecondition
	//
	LabelNameCondition interface{} `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

