package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   intermediateTableAnalysisRulePolicyProperty := &IntermediateTableAnalysisRulePolicyProperty{
//   	V1: &IntermediateTableAnalysisRulePolicyV1Property{
//   		Custom: &IntermediateTableAnalysisRuleCustomProperty{
//   			AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   			AggregationThresholds: []interface{}{
//   				&AggregationThresholdProperty{
//   					AllowedAggregateExpressionType: jsii.String("allowedAggregateExpressionType"),
//   					IdentityColumns: []*string{
//   						jsii.String("identityColumns"),
//   					},
//   					MinimumIdentityCount: jsii.Number(123),
//   					OutputColumnThresholds: []interface{}{
//   						&OutputColumnThresholdProperty{
//   							MinimumIdentityCount: jsii.Number(123),
//   							OutputColumnName: jsii.String("outputColumnName"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			AllowedAnalyses: []*string{
//   				jsii.String("allowedAnalyses"),
//   			},
//   			AllowedAnalysisProviders: []*string{
//   				jsii.String("allowedAnalysisProviders"),
//   			},
//   			AllowedResultReceivers: []*string{
//   				jsii.String("allowedResultReceivers"),
//   			},
//   			ComparisonControls: &ComparisonControlsProperty{
//   				AllowedColumnComparisonColumns: []*string{
//   					jsii.String("allowedColumnComparisonColumns"),
//   				},
//   				AllowedLiteralComparisonColumns: []*string{
//   					jsii.String("allowedLiteralComparisonColumns"),
//   				},
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulepolicy.html
//
type CfnIntermediateTablePropsMixin_IntermediateTableAnalysisRulePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulepolicy.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulepolicy-v1
	//
	V1 interface{} `field:"optional" json:"v1" yaml:"v1"`
}

