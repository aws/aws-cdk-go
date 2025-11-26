package previewawscleanroomsmixins


// The configured table association analysis rule applied to a configured table with the list analysis rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configuredTableAssociationAnalysisRuleListProperty := &ConfiguredTableAssociationAnalysisRuleListProperty{
//   	AllowedAdditionalAnalyses: []*string{
//   		jsii.String("allowedAdditionalAnalyses"),
//   	},
//   	AllowedResultReceivers: []*string{
//   		jsii.String("allowedResultReceivers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist.html
//
type CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleListProperty struct {
	// The list of resources or wildcards (ARNs) that are allowed to perform additional analysis on query output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist-allowedadditionalanalyses
	//
	AllowedAdditionalAnalyses *[]*string `field:"optional" json:"allowedAdditionalAnalyses" yaml:"allowedAdditionalAnalyses"`
	// The list of collaboration members who are allowed to receive results of queries run with this configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist-allowedresultreceivers
	//
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
}

