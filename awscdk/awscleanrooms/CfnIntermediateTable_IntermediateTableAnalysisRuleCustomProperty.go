package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intermediateTableAnalysisRuleCustomProperty := &IntermediateTableAnalysisRuleCustomProperty{
//   	AllowedAnalyses: []*string{
//   		jsii.String("allowedAnalyses"),
//   	},
//
//   	// the properties below are optional
//   	AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   	AllowedAnalysisProviders: []*string{
//   		jsii.String("allowedAnalysisProviders"),
//   	},
//   	AllowedResultReceivers: []*string{
//   		jsii.String("allowedResultReceivers"),
//   	},
//   	DifferentialPrivacy: &DifferentialPrivacyProperty{
//   		Columns: []interface{}{
//   			&DifferentialPrivacyColumnProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	DisallowedOutputColumns: []*string{
//   		jsii.String("disallowedOutputColumns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html
//
type CfnIntermediateTable_IntermediateTableAnalysisRuleCustomProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedanalyses
	//
	AllowedAnalyses *[]*string `field:"required" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-additionalanalyses
	//
	AdditionalAnalyses *string `field:"optional" json:"additionalAnalyses" yaml:"additionalAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedanalysisproviders
	//
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedresultreceivers
	//
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-differentialprivacy
	//
	DifferentialPrivacy interface{} `field:"optional" json:"differentialPrivacy" yaml:"differentialPrivacy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-disallowedoutputcolumns
	//
	DisallowedOutputColumns *[]*string `field:"optional" json:"disallowedOutputColumns" yaml:"disallowedOutputColumns"`
}

