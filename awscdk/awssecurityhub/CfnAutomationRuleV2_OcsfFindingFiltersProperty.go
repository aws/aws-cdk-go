package awssecurityhub


// Specifies the filtering criteria for security findings using OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocsfFindingFiltersProperty := &OcsfFindingFiltersProperty{
//   	CompositeFilters: []interface{}{
//   		&CompositeFilterProperty{
//   			BooleanFilters: []interface{}{
//   				&OcsfBooleanFilterProperty{
//   					FieldName: jsii.String("fieldName"),
//   					Filter: &BooleanFilterProperty{
//   						Value: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DateFilters: []interface{}{
//   				&OcsfDateFilterProperty{
//   					FieldName: jsii.String("fieldName"),
//   					Filter: &DateFilterProperty{
//   						DateRange: &DateRangeProperty{
//   							Unit: jsii.String("unit"),
//   							Value: jsii.Number(123),
//   						},
//   						End: jsii.String("end"),
//   						Start: jsii.String("start"),
//   					},
//   				},
//   			},
//   			MapFilters: []interface{}{
//   				&OcsfMapFilterProperty{
//   					FieldName: jsii.String("fieldName"),
//   					Filter: &MapFilterProperty{
//   						Comparison: jsii.String("comparison"),
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			NumberFilters: []interface{}{
//   				&OcsfNumberFilterProperty{
//   					FieldName: jsii.String("fieldName"),
//   					Filter: &NumberFilterProperty{
//   						Eq: jsii.Number(123),
//   						Gte: jsii.Number(123),
//   						Lte: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Operator: jsii.String("operator"),
//   			StringFilters: []interface{}{
//   				&OcsfStringFilterProperty{
//   					FieldName: jsii.String("fieldName"),
//   					Filter: &StringFilterProperty{
//   						Comparison: jsii.String("comparison"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CompositeOperator: jsii.String("compositeOperator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsffindingfilters.html
//
type CfnAutomationRuleV2_OcsfFindingFiltersProperty struct {
	// Enables the creation of complex filtering conditions by combining filter criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsffindingfilters.html#cfn-securityhub-automationrulev2-ocsffindingfilters-compositefilters
	//
	CompositeFilters interface{} `field:"optional" json:"compositeFilters" yaml:"compositeFilters"`
	// The logical operators used to combine the filtering on multiple `CompositeFilters` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsffindingfilters.html#cfn-securityhub-automationrulev2-ocsffindingfilters-compositeoperator
	//
	CompositeOperator *string `field:"optional" json:"compositeOperator" yaml:"compositeOperator"`
}

