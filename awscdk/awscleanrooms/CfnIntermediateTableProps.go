package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntermediateTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntermediateTableProps := &CfnIntermediateTableProps{
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	PopulationAnalysisConfiguration: &PopulationAnalysisConfigurationProperty{
//   		SqlParameters: &PopulationAnalysisSqlParametersProperty{
//   			AnalysisTemplateArn: jsii.String("analysisTemplateArn"),
//   			QueryString: jsii.String("queryString"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AnalysisRules: []interface{}{
//   		&IntermediateTableAnalysisRuleProperty{
//   			Policy: &IntermediateTableAnalysisRulePolicyProperty{
//   				V1: &IntermediateTableAnalysisRulePolicyV1Property{
//   					Custom: &IntermediateTableAnalysisRuleCustomProperty{
//   						AllowedAnalyses: []*string{
//   							jsii.String("allowedAnalyses"),
//   						},
//
//   						// the properties below are optional
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedAnalysisProviders: []*string{
//   							jsii.String("allowedAnalysisProviders"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   						DifferentialPrivacy: &DifferentialPrivacyProperty{
//   							Columns: []interface{}{
//   								&DifferentialPrivacyColumnProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   						},
//   						DisallowedOutputColumns: []*string{
//   							jsii.String("disallowedOutputColumns"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html
//
type CfnIntermediateTableProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-membershipidentifier
	//
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-populationanalysisconfiguration
	//
	PopulationAnalysisConfiguration interface{} `field:"required" json:"populationAnalysisConfiguration" yaml:"populationAnalysisConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-analysisrules
	//
	AnalysisRules interface{} `field:"optional" json:"analysisRules" yaml:"analysisRules"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html#cfn-cleanrooms-intermediatetable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

