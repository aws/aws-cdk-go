package awscleanrooms


// Controls on the query specifications that can be run on a configured table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredTableAnalysisRulePolicyProperty := &ConfiguredTableAnalysisRulePolicyProperty{
//   	V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   		Aggregation: &AnalysisRuleAggregationProperty{
//   			AggregateColumns: []interface{}{
//   				&AggregateColumnProperty{
//   					ColumnNames: []*string{
//   						jsii.String("columnNames"),
//   					},
//   					Function: jsii.String("function"),
//   				},
//   			},
//   			DimensionColumns: []*string{
//   				jsii.String("dimensionColumns"),
//   			},
//   			JoinColumns: []*string{
//   				jsii.String("joinColumns"),
//   			},
//   			OutputConstraints: []interface{}{
//   				&AggregationConstraintProperty{
//   					ColumnName: jsii.String("columnName"),
//   					Minimum: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			ScalarFunctions: []*string{
//   				jsii.String("scalarFunctions"),
//   			},
//
//   			// the properties below are optional
//   			AllowedJoinOperators: []*string{
//   				jsii.String("allowedJoinOperators"),
//   			},
//   			JoinRequired: jsii.String("joinRequired"),
//   		},
//   		Custom: &AnalysisRuleCustomProperty{
//   			AllowedAnalyses: []*string{
//   				jsii.String("allowedAnalyses"),
//   			},
//
//   			// the properties below are optional
//   			AllowedAnalysisProviders: []*string{
//   				jsii.String("allowedAnalysisProviders"),
//   			},
//   			DifferentialPrivacy: &DifferentialPrivacyProperty{
//   				Columns: []interface{}{
//   					&DifferentialPrivacyColumnProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		List: &AnalysisRuleListProperty{
//   			JoinColumns: []*string{
//   				jsii.String("joinColumns"),
//   			},
//   			ListColumns: []*string{
//   				jsii.String("listColumns"),
//   			},
//
//   			// the properties below are optional
//   			AllowedJoinOperators: []*string{
//   				jsii.String("allowedJoinOperators"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html
//
type CfnConfiguredTable_ConfiguredTableAnalysisRulePolicyProperty struct {
	// Controls on the query specifications that can be run on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicy-v1
	//
	V1 interface{} `field:"required" json:"v1" yaml:"v1"`
}

