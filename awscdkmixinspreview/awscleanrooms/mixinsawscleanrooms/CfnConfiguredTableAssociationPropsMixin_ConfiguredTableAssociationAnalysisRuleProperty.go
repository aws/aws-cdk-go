package mixinsawscleanrooms


// An analysis rule for a configured table association.
//
// This analysis rule specifies how data from the table can be used within its associated collaboration. In the console, the `ConfiguredTableAssociationAnalysisRule` is referred to as the *collaboration analysis rule* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuredTableAssociationAnalysisRuleProperty := &ConfiguredTableAssociationAnalysisRuleProperty{
//   	Policy: &ConfiguredTableAssociationAnalysisRulePolicyProperty{
//   		V1: &ConfiguredTableAssociationAnalysisRulePolicyV1Property{
//   			Aggregation: &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   				AllowedAdditionalAnalyses: []*string{
//   					jsii.String("allowedAdditionalAnalyses"),
//   				},
//   				AllowedResultReceivers: []*string{
//   					jsii.String("allowedResultReceivers"),
//   				},
//   			},
//   			Custom: &ConfiguredTableAssociationAnalysisRuleCustomProperty{
//   				AllowedAdditionalAnalyses: []*string{
//   					jsii.String("allowedAdditionalAnalyses"),
//   				},
//   				AllowedResultReceivers: []*string{
//   					jsii.String("allowedResultReceivers"),
//   				},
//   			},
//   			List: &ConfiguredTableAssociationAnalysisRuleListProperty{
//   				AllowedAdditionalAnalyses: []*string{
//   					jsii.String("allowedAdditionalAnalyses"),
//   				},
//   				AllowedResultReceivers: []*string{
//   					jsii.String("allowedResultReceivers"),
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule.html
//
type CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleProperty struct {
	// The policy of the configured table association analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The type of the configured table association analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

