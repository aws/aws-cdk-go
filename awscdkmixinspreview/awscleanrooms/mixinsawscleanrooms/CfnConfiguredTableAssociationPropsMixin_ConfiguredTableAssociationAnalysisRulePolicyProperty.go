package mixinsawscleanrooms


// Controls on the query specifications that can be run on an associated configured table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuredTableAssociationAnalysisRulePolicyProperty := &ConfiguredTableAssociationAnalysisRulePolicyProperty{
//   	V1: &ConfiguredTableAssociationAnalysisRulePolicyV1Property{
//   		Aggregation: &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   			AllowedAdditionalAnalyses: []*string{
//   				jsii.String("allowedAdditionalAnalyses"),
//   			},
//   			AllowedResultReceivers: []*string{
//   				jsii.String("allowedResultReceivers"),
//   			},
//   		},
//   		Custom: &ConfiguredTableAssociationAnalysisRuleCustomProperty{
//   			AllowedAdditionalAnalyses: []*string{
//   				jsii.String("allowedAdditionalAnalyses"),
//   			},
//   			AllowedResultReceivers: []*string{
//   				jsii.String("allowedResultReceivers"),
//   			},
//   		},
//   		List: &ConfiguredTableAssociationAnalysisRuleListProperty{
//   			AllowedAdditionalAnalyses: []*string{
//   				jsii.String("allowedAdditionalAnalyses"),
//   			},
//   			AllowedResultReceivers: []*string{
//   				jsii.String("allowedResultReceivers"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicy.html
//
type CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRulePolicyProperty struct {
	// The policy for the configured table association analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicy.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicy-v1
	//
	V1 interface{} `field:"optional" json:"v1" yaml:"v1"`
}

