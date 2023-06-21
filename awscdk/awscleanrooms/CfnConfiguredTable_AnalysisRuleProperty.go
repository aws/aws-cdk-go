package awscleanrooms


// A specification about how data from the configured table can be used in a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRuleProperty := &AnalysisRuleProperty{
//   	Policy: &ConfiguredTableAnalysisRulePolicyProperty{
//   		V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   			Aggregation: &AnalysisRuleAggregationProperty{
//   				AggregateColumns: []interface{}{
//   					&AggregateColumnProperty{
//   						ColumnNames: []*string{
//   							jsii.String("columnNames"),
//   						},
//   						Function: jsii.String("function"),
//   					},
//   				},
//   				DimensionColumns: []*string{
//   					jsii.String("dimensionColumns"),
//   				},
//   				JoinColumns: []*string{
//   					jsii.String("joinColumns"),
//   				},
//   				OutputConstraints: []interface{}{
//   					&AggregationConstraintProperty{
//   						ColumnName: jsii.String("columnName"),
//   						Minimum: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				ScalarFunctions: []*string{
//   					jsii.String("scalarFunctions"),
//   				},
//
//   				// the properties below are optional
//   				JoinRequired: jsii.String("joinRequired"),
//   			},
//   			List: &AnalysisRuleListProperty{
//   				JoinColumns: []*string{
//   					jsii.String("joinColumns"),
//   				},
//   				ListColumns: []*string{
//   					jsii.String("listColumns"),
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnConfiguredTable_AnalysisRuleProperty struct {
	// A policy that describes the associated data usage limitations.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The type of analysis rule.
	//
	// Valid values are `AGGREGATION` and `LIST`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

