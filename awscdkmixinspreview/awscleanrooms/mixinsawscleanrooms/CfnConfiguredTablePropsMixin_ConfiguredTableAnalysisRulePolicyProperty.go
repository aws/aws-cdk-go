package mixinsawscleanrooms


// Controls on the query specifications that can be run on a configured table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuredTableAnalysisRulePolicyProperty := &ConfiguredTableAnalysisRulePolicyProperty{
//   	V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   		Aggregation: &AnalysisRuleAggregationProperty{
//   			AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   			AggregateColumns: []interface{}{
//   				&AggregateColumnProperty{
//   					ColumnNames: []*string{
//   						jsii.String("columnNames"),
//   					},
//   					Function: jsii.String("function"),
//   				},
//   			},
//   			AllowedJoinOperators: []*string{
//   				jsii.String("allowedJoinOperators"),
//   			},
//   			DimensionColumns: []*string{
//   				jsii.String("dimensionColumns"),
//   			},
//   			JoinColumns: []*string{
//   				jsii.String("joinColumns"),
//   			},
//   			JoinRequired: jsii.String("joinRequired"),
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
//   		},
//   		Custom: &AnalysisRuleCustomProperty{
//   			AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   			AllowedAnalyses: []*string{
//   				jsii.String("allowedAnalyses"),
//   			},
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
//   			DisallowedOutputColumns: []*string{
//   				jsii.String("disallowedOutputColumns"),
//   			},
//   		},
//   		List: &AnalysisRuleListProperty{
//   			AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   			AllowedJoinOperators: []*string{
//   				jsii.String("allowedJoinOperators"),
//   			},
//   			JoinColumns: []*string{
//   				jsii.String("joinColumns"),
//   			},
//   			ListColumns: []*string{
//   				jsii.String("listColumns"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html
//
type CfnConfiguredTablePropsMixin_ConfiguredTableAnalysisRulePolicyProperty struct {
	// Controls on the query specifications that can be run on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicy-v1
	//
	V1 interface{} `field:"optional" json:"v1" yaml:"v1"`
}

