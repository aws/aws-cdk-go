package previewawscleanroomsmixins


// A specification about how data from the configured table can be used in a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisRuleProperty := &AnalysisRuleProperty{
//   	Policy: &ConfiguredTableAnalysisRulePolicyProperty{
//   		V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   			Aggregation: &AnalysisRuleAggregationProperty{
//   				AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   				AggregateColumns: []interface{}{
//   					&AggregateColumnProperty{
//   						ColumnNames: []*string{
//   							jsii.String("columnNames"),
//   						},
//   						Function: jsii.String("function"),
//   					},
//   				},
//   				AllowedJoinOperators: []*string{
//   					jsii.String("allowedJoinOperators"),
//   				},
//   				DimensionColumns: []*string{
//   					jsii.String("dimensionColumns"),
//   				},
//   				JoinColumns: []*string{
//   					jsii.String("joinColumns"),
//   				},
//   				JoinRequired: jsii.String("joinRequired"),
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
//   			},
//   			Custom: &AnalysisRuleCustomProperty{
//   				AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   				AllowedAnalyses: []*string{
//   					jsii.String("allowedAnalyses"),
//   				},
//   				AllowedAnalysisProviders: []*string{
//   					jsii.String("allowedAnalysisProviders"),
//   				},
//   				DifferentialPrivacy: &DifferentialPrivacyProperty{
//   					Columns: []interface{}{
//   						&DifferentialPrivacyColumnProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				DisallowedOutputColumns: []*string{
//   					jsii.String("disallowedOutputColumns"),
//   				},
//   			},
//   			List: &AnalysisRuleListProperty{
//   				AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   				AllowedJoinOperators: []*string{
//   					jsii.String("allowedJoinOperators"),
//   				},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html
//
type CfnConfiguredTablePropsMixin_AnalysisRuleProperty struct {
	// A policy that describes the associated data usage limitations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html#cfn-cleanrooms-configuredtable-analysisrule-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The type of analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrule.html#cfn-cleanrooms-configuredtable-analysisrule-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

