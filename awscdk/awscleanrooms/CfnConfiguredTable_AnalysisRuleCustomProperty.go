package awscleanrooms


// A type of analysis rule that enables the table owner to approve custom SQL queries on their configured tables.
//
// It supports differential privacy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRuleCustomProperty := &AnalysisRuleCustomProperty{
//   	AllowedAnalyses: []*string{
//   		jsii.String("allowedAnalyses"),
//   	},
//
//   	// the properties below are optional
//   	AllowedAnalysisProviders: []*string{
//   		jsii.String("allowedAnalysisProviders"),
//   	},
//   	DifferentialPrivacy: &DifferentialPrivacyProperty{
//   		Columns: []interface{}{
//   			&DifferentialPrivacyColumnProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulecustom.html
//
type CfnConfiguredTable_AnalysisRuleCustomProperty struct {
	// The ARN of the analysis templates that are allowed by the custom analysis rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulecustom.html#cfn-cleanrooms-configuredtable-analysisrulecustom-allowedanalyses
	//
	AllowedAnalyses *[]*string `field:"required" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// The IDs of the AWS accounts that are allowed to query by the custom analysis rule.
	//
	// Required when `allowedAnalyses` is `ANY_QUERY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulecustom.html#cfn-cleanrooms-configuredtable-analysisrulecustom-allowedanalysisproviders
	//
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
	// The differential privacy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulecustom.html#cfn-cleanrooms-configuredtable-analysisrulecustom-differentialprivacy
	//
	DifferentialPrivacy interface{} `field:"optional" json:"differentialPrivacy" yaml:"differentialPrivacy"`
}

