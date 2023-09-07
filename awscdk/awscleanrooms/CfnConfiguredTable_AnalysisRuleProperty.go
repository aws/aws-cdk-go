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
//   				AllowedJoinOperators: []*string{
//   					jsii.String("allowedJoinOperators"),
//   				},
//   				JoinRequired: jsii.String("joinRequired"),
//   			},
//   			Custom: &AnalysisRuleCustomProperty{
//   				AllowedAnalyses: []*string{
//   					jsii.String("allowedAnalyses"),
//   				},
//
//   				// the properties below are optional
//   				AllowedAnalysisProviders: []*string{
//   					jsii.String("allowedAnalysisProviders"),
//   				},
//   			},
//   			List: &AnalysisRuleListProperty{
//   				JoinColumns: []*string{
//   					jsii.String("joinColumns"),
//   				},
//   				ListColumns: []*string{
//   					jsii.String("listColumns"),
//   				},
//
//   				// the properties below are optional
//   				AllowedJoinOperators: []*string{
//   					jsii.String("allowedJoinOperators"),
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html
//
type CfnConfiguredTable_AnalysisRuleProperty struct {
	// A policy that describes the associated data usage limitations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html#cfn-cleanrooms-configuredtable-analysisrule-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The type of analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html#cfn-cleanrooms-configuredtable-analysisrule-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

