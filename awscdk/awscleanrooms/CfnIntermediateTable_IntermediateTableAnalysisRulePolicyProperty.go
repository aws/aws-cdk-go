package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intermediateTableAnalysisRulePolicyProperty := &IntermediateTableAnalysisRulePolicyProperty{
//   	V1: &IntermediateTableAnalysisRulePolicyV1Property{
//   		Custom: &IntermediateTableAnalysisRuleCustomProperty{
//   			AllowedAnalyses: []*string{
//   				jsii.String("allowedAnalyses"),
//   			},
//
//   			// the properties below are optional
//   			AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   			AllowedAnalysisProviders: []*string{
//   				jsii.String("allowedAnalysisProviders"),
//   			},
//   			AllowedResultReceivers: []*string{
//   				jsii.String("allowedResultReceivers"),
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
type CfnIntermediateTable_IntermediateTableAnalysisRulePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulepolicy.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulepolicy-v1
	//
	V1 interface{} `field:"required" json:"v1" yaml:"v1"`
}

