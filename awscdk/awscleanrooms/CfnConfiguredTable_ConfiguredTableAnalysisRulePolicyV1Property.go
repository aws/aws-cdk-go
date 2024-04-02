package awscleanrooms


// Controls on the query specifications that can be run on a configured table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredTableAnalysisRulePolicyV1Property := &ConfiguredTableAnalysisRulePolicyV1Property{
//   	Aggregation: &AnalysisRuleAggregationProperty{
//   		AggregateColumns: []interface{}{
//   			&AggregateColumnProperty{
//   				ColumnNames: []*string{
//   					jsii.String("columnNames"),
//   				},
//   				Function: jsii.String("function"),
//   			},
//   		},
//   		DimensionColumns: []*string{
//   			jsii.String("dimensionColumns"),
//   		},
//   		JoinColumns: []*string{
//   			jsii.String("joinColumns"),
//   		},
//   		OutputConstraints: []interface{}{
//   			&AggregationConstraintProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Minimum: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		ScalarFunctions: []*string{
//   			jsii.String("scalarFunctions"),
//   		},
//
//   		// the properties below are optional
//   		AllowedJoinOperators: []*string{
//   			jsii.String("allowedJoinOperators"),
//   		},
//   		JoinRequired: jsii.String("joinRequired"),
//   	},
//   	Custom: &AnalysisRuleCustomProperty{
//   		AllowedAnalyses: []*string{
//   			jsii.String("allowedAnalyses"),
//   		},
//
//   		// the properties below are optional
//   		AllowedAnalysisProviders: []*string{
//   			jsii.String("allowedAnalysisProviders"),
//   		},
//   		DifferentialPrivacy: &DifferentialPrivacyProperty{
//   			Columns: []interface{}{
//   				&DifferentialPrivacyColumnProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	List: &AnalysisRuleListProperty{
//   		JoinColumns: []*string{
//   			jsii.String("joinColumns"),
//   		},
//   		ListColumns: []*string{
//   			jsii.String("listColumns"),
//   		},
//
//   		// the properties below are optional
//   		AllowedJoinOperators: []*string{
//   			jsii.String("allowedJoinOperators"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1.html
//
type CfnConfiguredTable_ConfiguredTableAnalysisRulePolicyV1Property struct {
	// Analysis rule type that enables only aggregation queries on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1-aggregation
	//
	Aggregation interface{} `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Analysis rule type that enables custom SQL queries on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1-custom
	//
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// Analysis rule type that enables only list queries on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicyv1-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
}

