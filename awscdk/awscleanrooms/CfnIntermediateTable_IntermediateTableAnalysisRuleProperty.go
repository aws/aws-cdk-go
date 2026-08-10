package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intermediateTableAnalysisRuleProperty := &IntermediateTableAnalysisRuleProperty{
//   	Policy: &IntermediateTableAnalysisRulePolicyProperty{
//   		V1: &IntermediateTableAnalysisRulePolicyV1Property{
//   			Custom: &IntermediateTableAnalysisRuleCustomProperty{
//   				AllowedAnalyses: []*string{
//   					jsii.String("allowedAnalyses"),
//   				},
//
//   				// the properties below are optional
//   				AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   				AllowedAnalysisProviders: []*string{
//   					jsii.String("allowedAnalysisProviders"),
//   				},
//   				AllowedResultReceivers: []*string{
//   					jsii.String("allowedResultReceivers"),
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
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrule.html
//
type CfnIntermediateTable_IntermediateTableAnalysisRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrule.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrule-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrule.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrule-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

