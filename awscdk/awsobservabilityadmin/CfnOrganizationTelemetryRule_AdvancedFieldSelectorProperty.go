package awsobservabilityadmin


// Defines criteria for selecting resources based on field values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedFieldSelectorProperty := &AdvancedFieldSelectorProperty{
//   	EndsWith: []*string{
//   		jsii.String("endsWith"),
//   	},
//   	EqualTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	Field: jsii.String("field"),
//   	NotEndsWith: []*string{
//   		jsii.String("notEndsWith"),
//   	},
//   	NotEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   	NotStartsWith: []*string{
//   		jsii.String("notStartsWith"),
//   	},
//   	StartsWith: []*string{
//   		jsii.String("startsWith"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html
//
type CfnOrganizationTelemetryRule_AdvancedFieldSelectorProperty struct {
	// Matches if the field value ends with the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-endswith
	//
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Matches if the field value equals the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-equals
	//
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// The name of the field to use for selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Matches if the field value does not end with the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-notendswith
	//
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Matches if the field value does not equal the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-notequals
	//
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Matches if the field value does not start with the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-notstartswith
	//
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Matches if the field value starts with the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedfieldselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedfieldselector-startswith
	//
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

