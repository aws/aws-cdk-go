package awssecurityhub


// Enables the creation of filtering criteria for security findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compositeFilterProperty := &CompositeFilterProperty{
//   	BooleanFilters: []interface{}{
//   		&OcsfBooleanFilterProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Filter: &BooleanFilterProperty{
//   				Value: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	DateFilters: []interface{}{
//   		&OcsfDateFilterProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Filter: &DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   	},
//   	MapFilters: []interface{}{
//   		&OcsfMapFilterProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Filter: &MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	NumberFilters: []interface{}{
//   		&OcsfNumberFilterProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Filter: &NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Operator: jsii.String("operator"),
//   	StringFilters: []interface{}{
//   		&OcsfStringFilterProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Filter: &StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html
//
type CfnAutomationRuleV2_CompositeFilterProperty struct {
	// Enables filtering based on boolean field values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-booleanfilters
	//
	BooleanFilters interface{} `field:"optional" json:"booleanFilters" yaml:"booleanFilters"`
	// Enables filtering based on date and timestamp fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-datefilters
	//
	DateFilters interface{} `field:"optional" json:"dateFilters" yaml:"dateFilters"`
	// Enables the creation of filtering criteria for security findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-mapfilters
	//
	MapFilters interface{} `field:"optional" json:"mapFilters" yaml:"mapFilters"`
	// Enables filtering based on numerical field values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-numberfilters
	//
	NumberFilters interface{} `field:"optional" json:"numberFilters" yaml:"numberFilters"`
	// The logical operator used to combine multiple filter conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Enables filtering based on string field values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-compositefilter.html#cfn-securityhub-automationrulev2-compositefilter-stringfilters
	//
	StringFilters interface{} `field:"optional" json:"stringFilters" yaml:"stringFilters"`
}

