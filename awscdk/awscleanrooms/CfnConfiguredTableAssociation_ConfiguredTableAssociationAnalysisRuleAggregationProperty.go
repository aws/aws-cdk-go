package awscleanrooms


// The configured table association analysis rule applied to a configured table with the aggregation analysis rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredTableAssociationAnalysisRuleAggregationProperty := &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   	AllowedAdditionalAnalyses: []*string{
//   		jsii.String("allowedAdditionalAnalyses"),
//   	},
//   	AllowedResultReceivers: []*string{
//   		jsii.String("allowedResultReceivers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation.html
//
type CfnConfiguredTableAssociation_ConfiguredTableAssociationAnalysisRuleAggregationProperty struct {
	// The list of resources or wildcards (ARNs) that are allowed to perform additional analysis on query output.
	//
	// The `allowedAdditionalAnalyses` parameter is currently supported for the list analysis rule ( `AnalysisRuleList` ) and the custom analysis rule ( `AnalysisRuleCustom` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation-allowedadditionalanalyses
	//
	AllowedAdditionalAnalyses *[]*string `field:"optional" json:"allowedAdditionalAnalyses" yaml:"allowedAdditionalAnalyses"`
	// The list of collaboration members who are allowed to receive results of queries run with this configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation-allowedresultreceivers
	//
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
}

