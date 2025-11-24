package mixinsawscleanrooms


// Controls on the query specifications that can be run on an associated configured table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuredTableAssociationAnalysisRulePolicyV1Property := &ConfiguredTableAssociationAnalysisRulePolicyV1Property{
//   	Aggregation: &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   		AllowedAdditionalAnalyses: []*string{
//   			jsii.String("allowedAdditionalAnalyses"),
//   		},
//   		AllowedResultReceivers: []*string{
//   			jsii.String("allowedResultReceivers"),
//   		},
//   	},
//   	Custom: &ConfiguredTableAssociationAnalysisRuleCustomProperty{
//   		AllowedAdditionalAnalyses: []*string{
//   			jsii.String("allowedAdditionalAnalyses"),
//   		},
//   		AllowedResultReceivers: []*string{
//   			jsii.String("allowedResultReceivers"),
//   		},
//   	},
//   	List: &ConfiguredTableAssociationAnalysisRuleListProperty{
//   		AllowedAdditionalAnalyses: []*string{
//   			jsii.String("allowedAdditionalAnalyses"),
//   		},
//   		AllowedResultReceivers: []*string{
//   			jsii.String("allowedResultReceivers"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1.html
//
type CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRulePolicyV1Property struct {
	// Analysis rule type that enables only aggregation queries on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1-aggregation
	//
	Aggregation interface{} `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Analysis rule type that enables the table owner to approve custom SQL queries on their configured tables.
	//
	// It supports differential privacy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1-custom
	//
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// Analysis rule type that enables only list queries on a configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicyv1-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
}

